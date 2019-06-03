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

// +build integration sqlite

package service_test

import (
	"testing"
	"time"

	"github.com/hyperledger/burrow/integration"
	"github.com/hyperledger/burrow/integration/rpctest"

	"github.com/hyperledger/burrow/vent/test"
)

func TestSqliteConsumer(t *testing.T) {
	privateAccounts := rpctest.PrivateAccounts
	kern, shutdown := integration.RunNode(t, rpctest.GenesisDoc, rpctest.PrivateAccounts)
	defer shutdown()
	inputAddress := privateAccounts[0].GetAddress()
	grpcAddress := kern.GRPCListenAddress().String()
	tcli := test.NewTransactClient(t, grpcAddress)

	t.Parallel()
	time.Sleep(2 * time.Second)

	t.Run("Group", func(t *testing.T) {
		t.Run("Consume", func(t *testing.T) {
			testConsumer(t, kern.Blockchain.ChainID(), test.SqliteVentConfig(grpcAddress), tcli, inputAddress)
		})

		t.Run("SqliteInvalidUTF8", func(t *testing.T) {
			testInvalidUTF8(t, kern.Blockchain.ChainID(), test.SqliteVentConfig(grpcAddress), tcli, inputAddress)
		})

		t.Run("SqliteDeleteEvent", func(t *testing.T) {
			testDeleteEvent(t, kern.Blockchain.ChainID(), test.SqliteVentConfig(grpcAddress), tcli, inputAddress)
		})

		t.Run("SqliteResume", func(t *testing.T) {
			testResume(t, kern.Blockchain.ChainID(), test.SqliteVentConfig(grpcAddress))
		})
	})
}
