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

var Bytecode_ZeroReset = hex.MustDecodeString("608060405234801561001057600080fd5b50610195806100206000396000f3fe608060405234801561001057600080fd5b506004361061007e576000357c0100000000000000000000000000000000000000000000000000000000900480620267a4146100835780634ef65c3b146100a157806362738998146100cf578063747586b8146100ed578063987dc8201461011b578063b15a0d5f14610125575b600080fd5b61008b61012f565b6040518082815260200191505060405180910390f35b6100cd600480360360208110156100b757600080fd5b8101908080359060200190929190505050610139565b005b6100d7610143565b6040518082815260200191505060405180910390f35b6101196004803603602081101561010357600080fd5b810190808035906020019092919050505061014c565b005b610123610156565b005b61012d61015f565b005b6000600154905090565b8060018190555050565b60008054905090565b8060008190555050565b60008081905550565b600060018190555056fea165627a7a72305820f425681ba5df6ad8326c87681bfe7f8a84f407dc25e79a4cb790063ac3a8ba1f0029")
var Abi_ZeroReset = []byte(`[{"constant":true,"inputs":[],"name":"getUint","outputs":[{"name":"retUint","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"x","type":"uint256"}],"name":"setUint","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getInt","outputs":[{"name":"retInt","type":"int256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"x","type":"int256"}],"name":"setInt","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"setIntToZero","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"setUintToZero","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`)
