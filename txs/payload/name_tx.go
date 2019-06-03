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

package payload

import (
	"fmt"

	"github.com/hyperledger/burrow/acm/acmstate"
	"github.com/hyperledger/burrow/crypto"
)

func NewNameTx(st acmstate.AccountGetter, from crypto.PublicKey, name, data string, amt, fee uint64) (*NameTx, error) {
	addr := from.GetAddress()
	acc, err := st.GetAccount(addr)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, fmt.Errorf("NewNameTx: could not find account with address %v", addr)
	}

	sequence := acc.Sequence + 1
	return NewNameTxWithSequence(from, name, data, amt, fee, sequence), nil
}

func NewNameTxWithSequence(from crypto.PublicKey, name, data string, amt, fee, sequence uint64) *NameTx {
	input := &TxInput{
		Address:  from.GetAddress(),
		Amount:   amt,
		Sequence: sequence,
	}

	return &NameTx{
		Input: input,
		Name:  name,
		Data:  data,
		Fee:   fee,
	}
}

func (tx *NameTx) Type() Type {
	return TypeName
}

func (tx *NameTx) GetInputs() []*TxInput {
	return []*TxInput{tx.Input}
}

func (tx *NameTx) String() string {
	return fmt.Sprintf("NameTx{%v -> %s: %s}", tx.Input, tx.Name, tx.Data)
}

func (tx *NameTx) Any() *Any {
	return &Any{
		NameTx: tx,
	}
}
