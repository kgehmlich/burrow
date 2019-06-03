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

package mock

import (
	"testing"

	"encoding/json"

	"github.com/hyperledger/burrow/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMockKey_MonaxKeyJSON(t *testing.T) {
	key, err := newKey("monax-key-test")
	require.NoError(t, err)
	monaxKey := key.MonaxKeysJSON()
	t.Logf("key is: %v", monaxKey)
	keyJSON := &plainKeyJSON{}
	err = json.Unmarshal([]byte(monaxKey), keyJSON)
	require.NoError(t, err)
	// byte length of UUID string = 16 * 2 + 4 = 36
	assert.Equal(t, key.Address.String(), keyJSON.Address)
	assert.Equal(t, key.PrivateKey, keyJSON.PrivateKey.Plain)
	assert.Equal(t, string(crypto.CurveTypeEd25519.String()), keyJSON.Type)
}
