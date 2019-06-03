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

package errors

import (
	"bytes"
	"fmt"

	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/permission"
)

type PermissionDenied struct {
	Address crypto.Address
	Perm    permission.PermFlag
}

func (err PermissionDenied) ErrorCode() Code {
	return ErrorCodePermissionDenied
}

func (err PermissionDenied) Error() string {
	return fmt.Sprintf("Account/contract %v does not have permission %v", err.Address, err.Perm)
}

type NestedCallError struct {
	CodedError
	Caller     crypto.Address
	Callee     crypto.Address
	StackDepth uint64
}

func (err NestedCallError) Error() string {
	return fmt.Sprintf("error in nested call at depth %v: %s (callee) -> %s (caller): %v",
		err.StackDepth, err.Callee, err.Caller, err.CodedError)
}

type CallError struct {
	// The error from the original call which defines the overall error code
	CodedError
	// Errors from nested sub-calls of the original call that may have also occurred
	NestedErrors []NestedCallError
}

func (err CallError) Error() string {
	buf := new(bytes.Buffer)
	buf.WriteString("Call error: ")
	buf.WriteString(err.CodedError.String())
	if len(err.NestedErrors) > 0 {
		buf.WriteString(", nested call errors:\n")
		for _, nestedErr := range err.NestedErrors {
			buf.WriteString(nestedErr.Error())
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}
