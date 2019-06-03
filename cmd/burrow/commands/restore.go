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
	"github.com/hyperledger/burrow/core"
	cli "github.com/jawher/mow.cli"
)

// Restore reads a state file and saves into a runnable dir
func Restore(output Output) func(cmd *cli.Cmd) {
	return func(cmd *cli.Cmd) {
		configOpts := addConfigOptions(cmd)
		silentOpt := cmd.BoolOpt("s silent", false, "If state already exists don't throw error")
		filename := cmd.StringArg("FILE", "", "Restore from this dump")
		cmd.Spec += "[--silent] [FILE]"

		cmd.Action = func() {
			conf, err := configOpts.obtainBurrowConfig()
			if err != nil {
				output.Fatalf("could not set up config: %v", err)
			}

			if err := conf.Verify(); err != nil {
				output.Fatalf("cannot continue with config: %v", err)
			}

			output.Logf("Using validator address: %s", *conf.Address)

			kern, err := core.NewKernel(conf.BurrowDir)
			if err != nil {
				output.Fatalf("could not create Burrow kernel: %v", err)
			}

			if err = kern.LoadLoggerFromConfig(conf.Logging); err != nil {
				output.Fatalf("could not create Burrow kernel: %v", err)
			}

			if err = kern.LoadDump(conf.GenesisDoc, *filename, *silentOpt); err != nil {
				output.Fatalf("could not create Burrow kernel: %v", err)
			}

			kern.ShutdownAndExit()
		}
	}
}
