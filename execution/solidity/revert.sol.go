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

package solidity

import hex "github.com/tmthrgd/go-hex"

var Bytecode_Revert = hex.MustDecodeString("608060405234801561001057600080fd5b50610216806100206000396000f3fe608060405234801561001057600080fd5b5060043610610053576000357c0100000000000000000000000000000000000000000000000000000000900480635b202afb146100585780636037b04c1461008c575b600080fd5b61008a6004803603602081101561006e57600080fd5b81019080803563ffffffff169060200190929190505050610096565b005b6100946101e5565b005b60008163ffffffff161415610113576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f492068617665207265766572746564000000000000000000000000000000000081525060200191505060405180910390fd5b8080600190039150508063ffffffff167ff7f0feb5b4ac5276c55faa8936d962de931ebe8333a2efdc0506878de3979ba960405160405180910390a23073ffffffffffffffffffffffffffffffffffffffff16635b202afb826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b1580156101ca57600080fd5b505af11580156101de573d6000803e3d6000fd5b5050505050565b600080fdfea165627a7a7230582021dc462608e362c1ab12bfca9c326514ec64ff6bbe411b6a3b1c41e20a82a8ac0029")
var Abi_Revert = []byte(`[{"constant":false,"inputs":[{"name":"i","type":"uint32"}],"name":"RevertAt","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"RevertNoReason","outputs":[],"payable":false,"stateMutability":"pure","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"i","type":"uint32"}],"name":"NotReverting","type":"event"}]`)
