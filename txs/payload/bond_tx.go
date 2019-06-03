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

func NewBondTx(pubkey crypto.PublicKey) (*BondTx, error) {
	return &BondTx{
		Inputs:   []*TxInput{},
		UnbondTo: []*TxOutput{},
	}, nil
}

func (tx *BondTx) Type() Type {
	return TypeBond
}

func (tx *BondTx) GetInputs() []*TxInput {
	return tx.Inputs
}

func (tx *BondTx) String() string {
	return fmt.Sprintf("BondTx{%v -> %v}", tx.Inputs, tx.UnbondTo)
}

func (tx *BondTx) AddInput(st acmstate.AccountGetter, pubkey crypto.PublicKey, amt uint64) error {
	addr := pubkey.GetAddress()
	acc, err := st.GetAccount(addr)
	if err != nil {
		return err
	}
	if acc == nil {
		return fmt.Errorf("Invalid address %s from pubkey %s", addr, pubkey)
	}
	return tx.AddInputWithSequence(pubkey, amt, acc.Sequence+uint64(1))
}

func (tx *BondTx) AddInputWithSequence(pubkey crypto.PublicKey, amt uint64, sequence uint64) error {
	tx.Inputs = append(tx.Inputs, &TxInput{
		Address:  pubkey.GetAddress(),
		Amount:   amt,
		Sequence: sequence,
	})
	return nil
}

func (tx *BondTx) AddOutput(addr crypto.Address, amt uint64) error {
	tx.UnbondTo = append(tx.UnbondTo, &TxOutput{
		Address: addr,
		Amount:  amt,
	})
	return nil
}

func (tx *BondTx) Any() *Any {
	return &Any{
		BondTx: tx,
	}
}
