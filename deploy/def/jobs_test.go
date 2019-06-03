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

package def

import (
	"testing"

	"github.com/hyperledger/burrow/acm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGovernAccount_Validate(t *testing.T) {
	sourceAddress := acm.GeneratePrivateAccountFromSecret("frogs").GetAddress()
	targetAddress := acm.GeneratePrivateAccountFromSecret("logs").GetAddress()
	job := &UpdateAccount{
		Target:      targetAddress.String(),
		Source:      sourceAddress.String(),
		Sequence:    "34",
		Native:      "1033",
		Power:       "324324322",
		Roles:       []string{"foo"},
		Permissions: []PermissionString{"root", "send"},
	}
	err := job.Validate()
	require.NoError(t, err)
}

func TestKeyNameCurveType(t *testing.T) {
	match := NewKeyRegex.FindStringSubmatch("new()")
	keyName, curveType := KeyNameCurveType(match)
	assert.Equal(t, "", keyName)
	assert.Equal(t, "", curveType)

	match = NewKeyRegex.FindStringSubmatch("new(mySpecialKey)")
	keyName, curveType = KeyNameCurveType(match)
	assert.Equal(t, "mySpecialKey", keyName)
	assert.Equal(t, "", curveType)

	match = NewKeyRegex.FindStringSubmatch("new(,secp256k1)")
	keyName, curveType = KeyNameCurveType(match)
	assert.Equal(t, "", keyName)
	assert.Equal(t, "secp256k1", curveType)

	match = NewKeyRegex.FindStringSubmatch("new(myLessSpecialKey0,ed25519)")
	keyName, curveType = KeyNameCurveType(match)
	assert.Equal(t, "myLessSpecialKey0", keyName)
	assert.Equal(t, "ed25519", curveType)

	assert.False(t, NewKeyRegex.MatchString("new"))
}
