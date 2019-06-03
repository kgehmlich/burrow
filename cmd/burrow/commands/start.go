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

// Start launches the burrow daemon
func Start(output Output) func(cmd *cli.Cmd) {
	return func(cmd *cli.Cmd) {
		configOpts := addConfigOptions(cmd)

		cmd.Action = func() {
			conf, err := configOpts.obtainBurrowConfig()
			if err != nil {
				output.Fatalf("could not set up config: %v", err)
			}

			if err := conf.Verify(); err != nil {
				output.Fatalf("cannot continue with config: %v", err)
			}

			output.Logf("Using validator address: %s", *conf.Address)

			kern, err := core.LoadKernelFromConfig(conf)
			if err != nil {
				output.Fatalf("could not configure Burrow kernel: %v", err)
			}

			if err = kern.Boot(); err != nil {
				output.Fatalf("could not boot Burrow kernel: %v", err)
			}

			kern.WaitForShutdown()
		}
	}
}
