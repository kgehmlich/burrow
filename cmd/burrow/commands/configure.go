// Copyright 2019 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/hyperledger/burrow/config/source"
	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/deployment"
	"github.com/hyperledger/burrow/execution"
	"github.com/hyperledger/burrow/execution/state"
	"github.com/hyperledger/burrow/genesis"
	"github.com/hyperledger/burrow/genesis/spec"
	"github.com/hyperledger/burrow/keys"
	"github.com/hyperledger/burrow/logging"
	"github.com/hyperledger/burrow/logging/logconfig"
	"github.com/hyperledger/burrow/logging/logconfig/presets"
	"github.com/hyperledger/burrow/rpc"
	cli "github.com/jawher/mow.cli"
	amino "github.com/tendermint/go-amino"
	tmEd25519 "github.com/tendermint/tendermint/crypto/ed25519"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/p2p"
)

// Configure generates burrow configuration(s)
func Configure(output Output) func(cmd *cli.Cmd) {
	return func(cmd *cli.Cmd) {

		genesisSpecOpt := cmd.StringOpt("s genesis-spec", "",
			"A GenesisSpec to use as a template for a GenesisDoc that will be created along with keys")

		jsonOutOpt := cmd.BoolOpt("j json", false, "Emit config in JSON rather than TOML "+
			"suitable for further processing")

		keysURLOpt := cmd.StringOpt("k keys-url", "", fmt.Sprintf("Provide keys GRPC address, default: %s",
			keys.DefaultKeysConfig().RemoteAddress))

		keysDir := cmd.StringOpt("keys-dir", "", "Directory where keys are stored")

		generateNodeKeys := cmd.BoolOpt("generate-node-keys", false, "Generate node keys for validators")

		configTemplateIn := cmd.StringsOpt("config-template-in", nil,
			fmt.Sprintf("Go text/template template input filename (left delim: %s right delim: %s) to generate config "+
				"file specified with --config-out", deployment.LeftTemplateDelim, deployment.RightTemplateDelim))

		configOut := cmd.StringsOpt("config-out", nil,
			"Go text/template template output file. Template filename specified with --config-template-in "+
				"file specified with --config-out")

		separateGenesisDoc := cmd.StringOpt("w separate-genesis-doc", "", "Emit a separate genesis doc as JSON or TOML")

		loggingOpt := cmd.StringOpt("l logging", "",
			"Comma separated list of logging instructions which form a 'program' which is a depth-first "+
				"pre-order of instructions that will build the root logging sink. See 'burrow help' for more information.")

		describeLoggingOpt := cmd.BoolOpt("describe-logging", false,
			"Print an exhaustive list of logging instructions available with the --logging option")

		debugOpt := cmd.BoolOpt("d debug", false, "Include maximal debug options in config "+
			"including logging opcodes and dumping EVM tokens to disk these can be later pruned from the "+
			"generated config.")

		chainNameOpt := cmd.StringOpt("n chain-name", "", "Default chain name")

		restoreDumpOpt := cmd.StringOpt("restore-dump", "", "Including AppHash for restored file")

		pool := cmd.BoolOpt("pool", false, "Write config files for all the validators called burrowNNN.toml")

		cmd.Spec = "[--keys-url=<keys URL> | --keys-dir=<keys directory>] " +
			"[ --config-template-in=<text template> --config-out=<output file>]... " +
			"[--genesis-spec=<GenesisSpec file>] [--separate-genesis-doc=<genesis JSON file>] " +
			"[--chain-name=<chain name>] [--generate-node-keys] [--restore-dump=<dump file>] " +
			"[--logging=<logging program>] [--describe-logging] [--json] [--debug] [--pool]"

		// no sourcing logs
		source.LogWriter = ioutil.Discard
		// TODO: this has default, only set if explicit?
		configOpts := addConfigOptions(cmd)

		cmd.Action = func() {
			conf, err := configOpts.obtainBurrowConfig()
			if err != nil {
				output.Fatalf("could not obtain config: %v", err)
			}

			if *describeLoggingOpt {
				output.Logf("Usage:\n  burrow configure -l INSTRUCTION[,...]\n\nBuilds a logging " +
					"configuration by constructing a tree of logging sinks assembled from preset instructions " +
					"that generate the tree while traversing it.\n\nLogging Instructions:\n")
				for _, instruction := range presets.Instructons() {
					output.Logf("  %-15s\t%s\n", instruction.Name(), instruction.Description())
				}
				output.Logf("\nExample Usage:\n  burrow configure -l include-any,info,stderr\n")
				return
			}

			if *keysURLOpt != "" {
				conf.Keys.RemoteAddress = *keysURLOpt
			}

			if len(*configTemplateIn) != len(*configOut) {
				output.Fatalf("--config-template-in and --config-out must be specified the same number of times")
			}

			pkg := deployment.Config{}

			// Genesis Spec
			if *genesisSpecOpt != "" {
				genesisSpec := new(spec.GenesisSpec)
				err := source.FromFile(*genesisSpecOpt, genesisSpec)
				if err != nil {
					output.Fatalf("Could not read GenesisSpec: %v", err)
				}
				if *restoreDumpOpt != "" {
					output.Fatalf("Cannot restore using GenesisSpec, please provide GenesisDoc")
				}
				if conf.Keys.RemoteAddress == "" {
					dir := conf.Keys.KeysDirectory
					if *keysDir != "" {
						dir = *keysDir
					}
					keyStore := keys.NewKeyStore(dir, conf.Keys.AllowBadFilePermissions)

					keyClient := keys.NewLocalKeyClient(keyStore, logging.NewNoopLogger())
					conf.GenesisDoc, err = genesisSpec.GenesisDoc(keyClient, *generateNodeKeys || *pool)
					if err != nil {
						output.Fatalf("Could not generate GenesisDoc from GenesisSpec using MockKeyClient: %v", err)
					}

					allNames, err := keyStore.GetAllNames()
					if err != nil {
						output.Fatalf("could get all keys: %v", err)
					}

					cdc := amino.NewCodec()
					cryptoAmino.RegisterAmino(cdc)

					pkg = deployment.Config{Keys: make(map[crypto.Address]deployment.Key)}

					for k := range allNames {
						addr, err := crypto.AddressFromHexString(allNames[k])
						if err != nil {
							output.Fatalf("Address %s not valid: %v", k, err)
						}
						key, err := keyStore.GetKey("", addr[:])
						if err != nil {
							output.Fatalf("Failed to get key: %s: %v", k, err)
						}

						// Is this is a validator node key?
						nodeKey := false
						for _, a := range conf.GenesisDoc.Validators {
							if a.NodeAddress != nil && addr == *a.NodeAddress {
								nodeKey = true
								break
							}
						}

						if nodeKey {
							privKey := tmEd25519.GenPrivKey()
							copy(privKey[:], key.PrivateKey.PrivateKey)
							nodeKey := &p2p.NodeKey{
								PrivKey: privKey,
							}

							json, err := cdc.MarshalJSON(nodeKey)
							if err != nil {
								output.Fatalf("go-amino failed to json marshall private key: %v", err)
							}
							pkg.Keys[addr] = deployment.Key{Name: k, Address: addr, KeyJSON: json}
						} else {
							json, err := json.Marshal(key)
							if err != nil {
								output.Fatalf("Failed to json marshal key: %s: %v", k, err)
							}
							pkg.Keys[addr] = deployment.Key{Name: k, Address: addr, KeyJSON: json}
						}
					}
				} else {
					keyClient, err := keys.NewRemoteKeyClient(conf.Keys.RemoteAddress, logging.NewNoopLogger())
					if err != nil {
						output.Fatalf("Could not create remote key client: %v", err)
					}
					conf.GenesisDoc, err = genesisSpec.GenesisDoc(keyClient, *generateNodeKeys || *pool)
				}

				if err != nil {
					output.Fatalf("could not realise GenesisSpec: %v", err)
				}
			}

			if *chainNameOpt != "" {
				if conf.GenesisDoc == nil {
					output.Fatalf("Unable to set ChainName since no GenesisDoc/GenesisSpec provided.")
				}
				conf.GenesisDoc.ChainName = *chainNameOpt
			}

			if *restoreDumpOpt != "" {
				if conf.GenesisDoc == nil {
					output.Fatalf("no GenesisDoc provided, cannot restore dump")
				}

				if len(conf.GenesisDoc.Validators) == 0 {
					output.Fatalf("On restore, validators must be provided in GenesisDoc or GenesisSpec")
				}

				reader, err := state.NewFileDumpReader(*restoreDumpOpt)
				if err != nil {
					output.Fatalf("Failed to read restore dump: %v", err)
				}

				st, err := state.MakeGenesisState(db.NewMemDB(), conf.GenesisDoc)
				if err != nil {
					output.Fatalf("could not generate state from genesis: %v", err)
				}

				err = st.LoadDump(reader)
				if err != nil {
					output.Fatalf("could not restore dump %s: %v", *restoreDumpOpt, err)
				}

				err = st.InitialCommit()
				if err != nil {
					output.Fatalf("could not commit: %v", err)
				}

				conf.GenesisDoc.AppHash = st.Hash()
			}

			if conf.GenesisDoc != nil {
				pkg.Config = conf.GenesisDoc

				for _, v := range conf.GenesisDoc.Validators {
					tmplV := deployment.Validator{
						Name:    v.Name,
						Address: v.Address,
					}

					if v.NodeAddress != nil {
						tmplV.NodeAddress = *v.NodeAddress
					}

					pkg.Validators = append(pkg.Validators, tmplV)
				}

				for ind := range *configTemplateIn {
					err := processTemplate(&pkg, conf.GenesisDoc, (*configTemplateIn)[ind], (*configOut)[ind])
					if err != nil {
						output.Fatalf("could not template from %s to %s: %v", (*configTemplateIn)[ind], (*configOut)[ind], err)
					}
				}
			}

			// Logging
			if *loggingOpt != "" {
				ops := strings.Split(*loggingOpt, ",")
				sinkConfig, err := presets.BuildSinkConfig(ops...)
				if err != nil {
					output.Fatalf("could not build logging configuration: %v\n\nTo see possible logging "+
						"instructions run:\n  burrow configure --describe-logging", err)
				}
				conf.Logging = &logconfig.LoggingConfig{
					RootSink: sinkConfig,
				}
			}

			if *debugOpt {
				conf.Execution = &execution.ExecutionConfig{
					VMOptions: []execution.VMOption{execution.DumpTokens, execution.DebugOpcodes},
				}
			}

			if *separateGenesisDoc != "" {
				if conf.GenesisDoc == nil {
					output.Fatalf("Cannot write separate genesis doc since no GenesisDoc/GenesisSpec provided.")
				}
				genesisDocJSON, err := conf.GenesisDoc.JSONBytes()
				if err != nil {
					output.Fatalf("Could not form GenesisDoc JSON: %v", err)
				}
				err = ioutil.WriteFile(*separateGenesisDoc, genesisDocJSON, 0644)
				if err != nil {
					output.Fatalf("Could not write GenesisDoc JSON: %v", err)
				}
				conf.GenesisDoc = nil
			}

			if *pool {
				peers := make([]string, 0)
				for i, val := range conf.GenesisDoc.Validators {
					if val.NodeAddress == nil {
						continue
					}
					peers = append(peers, fmt.Sprintf("tcp://%s@127.0.0.1:%d", strings.ToLower(val.NodeAddress.String()), 26656+i))
				}
				for i, acc := range conf.GenesisDoc.Accounts {
					// set stuff
					conf.Address = &acc.Address
					conf.Tendermint.PersistentPeers = strings.Join(peers, ",")

					conf.BurrowDir = fmt.Sprintf(".burrow%03d", i)
					conf.Tendermint.ListenHost = rpc.LocalHost
					conf.Tendermint.ListenPort = fmt.Sprint(26656 + i)
					conf.RPC.Info.ListenHost = rpc.LocalHost
					conf.RPC.Info.ListenPort = fmt.Sprint(26758 + i)
					conf.RPC.GRPC.ListenHost = rpc.LocalHost
					conf.RPC.GRPC.ListenPort = fmt.Sprint(10997 + i)
					conf.RPC.Metrics.ListenHost = rpc.LocalHost
					conf.RPC.Metrics.ListenPort = fmt.Sprint(9102 + i)
					conf.Logging.RootSink.Output.OutputType = "file"
					conf.Logging.RootSink.Output.FileConfig = &logconfig.FileConfig{Path: fmt.Sprintf("burrow%03d.log", i)}

					if *jsonOutOpt {
						ioutil.WriteFile(fmt.Sprintf("burrow%03d.json", i), []byte(conf.JSONString()), 0644)
					} else {
						ioutil.WriteFile(fmt.Sprintf("burrow%03d.toml", i), []byte(conf.TOMLString()), 0644)
					}
				}
			} else if *jsonOutOpt {
				output.Printf(conf.JSONString())
			} else {
				output.Printf(conf.TOMLString())
			}

		}
	}
}

func processTemplate(pkg *deployment.Config, config *genesis.GenesisDoc, templateIn, templateOut string) error {
	data, err := ioutil.ReadFile(templateIn)
	if err != nil {
		return err
	}
	output, err := pkg.Dump(templateIn, string(data))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(templateOut, []byte(output), 0644)
}
