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

package keys

import (
	"net"
	"os"
	"testing"

	"github.com/hyperledger/burrow/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitKeyClient(t *testing.T) {
	dirTest := "test_scratch/.keys"
	os.RemoveAll(dirTest)
	server := keys.StandAloneServer(dirTest, true)
	listener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	address := listener.Addr().String()
	go server.Serve(listener)
	localKeyClient, err := InitKeyClient(address)
	require.NoError(t, err)
	err = localKeyClient.HealthCheck()
	assert.NoError(t, err)
}
