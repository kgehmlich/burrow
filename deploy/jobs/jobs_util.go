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

package jobs

import (
	"github.com/hyperledger/burrow/deploy/def"
	"github.com/hyperledger/burrow/logging"
)

func SetAccountJob(account *def.Account, do *def.DeployArgs, script *def.Playbook, logger *logging.Logger) (string, error) {
	var result string

	// Set the Account in the Package & Announce
	script.Account = account.Address
	logger.InfoMsg("Setting Account", "account", script.Account)

	// Set result and return
	result = account.Address
	return result, nil
}

func SetValJob(set *def.Set, do *def.DeployArgs, logger *logging.Logger) (string, error) {
	var result string
	logger.InfoMsg("Setting Variable", "result", set.Value)
	result = set.Value
	return result, nil
}
