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

var Bytecode_EventEmitter = hex.MustDecodeString("6080604052348015600f57600080fd5b506101908061001f6000396000f3fe608060405234801561001057600080fd5b5060043610610048576000357c010000000000000000000000000000000000000000000000000000000090048063e8e49a711461004d575b600080fd5b610055610057565b005b60405180807f68617368000000000000000000000000000000000000000000000000000000008152506004019050604051809103902060667f446f776e736965210000000000000000000000000000000000000000000000007f20aec2a3bcd8050a3a9e852e9d424805bad75ba33b57077464c73ae98d0582696001602a6040518083151515158152602001806020018381526020018281038252605181526020018061011460519139606001935050505060405180910390a456fe446f6e617564616d7066736368696666666168727473656c656b7472697a6974c3a474656e686175707462657472696562737765726b626175756e7465726265616d74656e676573656c6c736368616674a165627a7a7230582043472c03b2946767b21150a9f581b7cee0c585db6817446b4fa045bff32809450029")
var Abi_EventEmitter = []byte(`[{"constant":false,"inputs":[],"name":"EmitOne","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"direction","type":"bytes32"},{"indexed":false,"name":"trueism","type":"bool"},{"indexed":false,"name":"german","type":"string"},{"indexed":true,"name":"newDepth","type":"int64"},{"indexed":false,"name":"bignum","type":"int256"},{"indexed":true,"name":"hash","type":"string"}],"name":"ManyTypes","type":"event"}]`)
