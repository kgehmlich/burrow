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

	"github.com/hyperledger/burrow/crypto"
)

func NewUnbondTx(address crypto.Address, height uint64) *UnbondTx {
	return &UnbondTx{
		Address: address,
		Height:  height,
	}
}

func (tx *UnbondTx) Type() Type {
	return TypeUnbond
}

func (tx *UnbondTx) GetInputs() []*TxInput {
	return []*TxInput{tx.Input}
}

func (tx *UnbondTx) String() string {
	return fmt.Sprintf("UnbondTx{%v -> %s,%v}", tx.Input, tx.Address, tx.Height)
}

func (tx *UnbondTx) Any() *Any {
	return &Any{
		UnbondTx: tx,
	}
}
