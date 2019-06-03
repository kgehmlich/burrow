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

package acmstate

import (
	"bytes"
	"encoding/hex"
	"encoding/json"

	"github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/binary"
	"github.com/hyperledger/burrow/crypto"
)

type DumpState struct {
	bytes.Buffer
}

func (dw *DumpState) UpdateAccount(updatedAccount *acm.Account) error {
	dw.WriteString("UpdateAccount\n")
	bs, err := json.Marshal(updatedAccount)
	if err != nil {
		return err
	}
	dw.Write(bs)
	dw.WriteByte('\n')
	return nil
}

func (dw *DumpState) RemoveAccount(address crypto.Address) error {
	dw.WriteString("RemoveAccount\n")
	dw.WriteString(address.String())
	dw.WriteByte('\n')
	return nil
}

func (dw *DumpState) SetStorage(address crypto.Address, key, value binary.Word256) error {
	dw.WriteString("SetStorage\n")
	dw.WriteString(address.String())
	dw.WriteByte('/')
	dw.WriteString(hex.EncodeToString(key[:]))
	dw.WriteByte('/')
	dw.WriteString(hex.EncodeToString(value[:]))
	dw.WriteByte('\n')
	return nil
}
