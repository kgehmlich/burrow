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
	"fmt"

	"github.com/hyperledger/burrow/crypto"
)

type LacksSNativePermission struct {
	Address crypto.Address
	SNative string
}

func (e LacksSNativePermission) Error() string {
	return fmt.Sprintf("account %s does not have SNative function call permission: %s", e.Address, e.SNative)
}

func (e LacksSNativePermission) ErrorCode() Code {
	return ErrorCodeNativeFunction
}
