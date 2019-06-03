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

package evm

import "github.com/hyperledger/burrow/execution/errors"

func MemoryProvider(memoryProvider func(errors.Sink) Memory) func(*VM) {
	return func(vm *VM) {
		vm.memoryProvider = memoryProvider
	}
}

func DebugOpcodes(vm *VM) {
	vm.debugOpcodes = true
}

func DumpTokens(vm *VM) {
	vm.dumpTokens = true
}

func StackOptions(callStackMaxDepth uint64, dataStackInitialCapacity uint64, dataStackMaxDepth uint64) func(*VM) {
	return func(vm *VM) {
		vm.params.CallStackMaxDepth = callStackMaxDepth
		vm.params.DataStackInitialCapacity = dataStackInitialCapacity
		vm.params.DataStackMaxDepth = dataStackMaxDepth
	}
}
