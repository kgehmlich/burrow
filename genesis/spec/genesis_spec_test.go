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

package spec

import (
	"testing"

	"github.com/hyperledger/burrow/acm/balance"
	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/keys/mock"
	"github.com/hyperledger/burrow/permission"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenesisSpec_GenesisDoc(t *testing.T) {
	keyClient := mock.NewKeyClient()

	// Try a spec with a single account/validator
	amtBonded := uint64(100)
	genesisSpec := GenesisSpec{
		Accounts: []TemplateAccount{{
			Amounts: balance.New().Power(amtBonded),
		}},
	}

	genesisDoc, err := genesisSpec.GenesisDoc(keyClient, false)
	require.NoError(t, err)
	require.Len(t, genesisDoc.Accounts, 1)
	// Should create validator
	require.Len(t, genesisDoc.Validators, 1)
	assert.NotZero(t, genesisDoc.Accounts[0].Address)
	assert.NotZero(t, genesisDoc.Accounts[0].PublicKey)
	assert.Equal(t, genesisDoc.Accounts[0].Address, genesisDoc.Validators[0].Address)
	assert.Equal(t, genesisDoc.Accounts[0].PublicKey, genesisDoc.Validators[0].PublicKey)
	assert.Equal(t, amtBonded, genesisDoc.Validators[0].Amount)
	assert.NotEmpty(t, genesisDoc.ChainName, "Chain name should not be empty")

	address, err := keyClient.Generate("test-lookup-of-key", crypto.CurveTypeEd25519)
	require.NoError(t, err)
	pubKey, err := keyClient.PublicKey(address)
	require.NoError(t, err)

	// Try a spec with two accounts and no validators
	amt := uint64(99299299)
	genesisSpec = GenesisSpec{
		Accounts: []TemplateAccount{
			{
				Address: &address,
			},
			{
				Amounts:     balance.New().Native(amt),
				Permissions: []string{permission.CreateAccountString, permission.CallString},
			}},
	}

	genesisDoc, err = genesisSpec.GenesisDoc(keyClient, false)
	require.NoError(t, err)

	require.Len(t, genesisDoc.Accounts, 2)
	// Nothing bonded so no validators
	require.Len(t, genesisDoc.Validators, 0)
	assert.Equal(t, pubKey, genesisDoc.Accounts[0].PublicKey)
	assert.Equal(t, amt, genesisDoc.Accounts[1].Amount)
	permFlag := permission.CreateAccount | permission.Call
	assert.Equal(t, permFlag, genesisDoc.Accounts[1].Permissions.Base.Perms)
	assert.Equal(t, permFlag, genesisDoc.Accounts[1].Permissions.Base.SetBit)

	// Try an empty spec
	genesisSpec = GenesisSpec{}

	genesisDoc, err = genesisSpec.GenesisDoc(keyClient, false)
	require.NoError(t, err)

	// Similar assersions to first case - should generate our default single identity chain
	require.Len(t, genesisDoc.Accounts, 1)
	// Should create validator
	require.Len(t, genesisDoc.Validators, 1)
	assert.NotZero(t, genesisDoc.Accounts[0].Address)
	assert.NotZero(t, genesisDoc.Accounts[0].PublicKey)
	assert.Equal(t, genesisDoc.Accounts[0].Address, genesisDoc.Validators[0].Address)
	assert.Equal(t, genesisDoc.Accounts[0].PublicKey, genesisDoc.Validators[0].PublicKey)
}

func TestTemplateAccount_AccountPermissions(t *testing.T) {
}
