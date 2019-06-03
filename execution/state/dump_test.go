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

package state

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/genesis"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tendermint/libs/db"

	"github.com/hyperledger/burrow/binary"
	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/dump"
	"github.com/hyperledger/burrow/execution/exec"
	"github.com/hyperledger/burrow/execution/names"
)

type MockDumpReader struct {
	accounts int
	storage  int
	names    int
	events   int
}

func (m *MockDumpReader) Next() (*dump.Dump, error) {
	// acccounts
	row := dump.Dump{Height: 102}

	if m.accounts > 0 {
		var addr crypto.Address
		binary.PutUint64BE(addr.Bytes(), uint64(m.accounts))

		row.Account = &acm.Account{
			Address: addr,
			Balance: 102,
		}

		if m.accounts%2 > 0 {
			row.Account.Code = make([]byte, rand.Int()%10000)
		} else {
			row.Account.PublicKey = crypto.PublicKey{}
		}
		m.accounts--
	} else if m.storage > 0 {
		var addr crypto.Address
		binary.PutUint64BE(addr.Bytes(), uint64(m.storage))
		storagelen := rand.Int() % 25

		row.AccountStorage = &dump.AccountStorage{
			Address: addr,
			Storage: make([]*dump.Storage, storagelen),
		}

		for i := 0; i < storagelen; i++ {
			row.AccountStorage.Storage[i] = &dump.Storage{}
		}

		m.storage--
	} else if m.names > 0 {
		row.Name = &names.Entry{
			Name:    fmt.Sprintf("name%d", m.names),
			Data:    fmt.Sprintf("data%x", m.names),
			Owner:   crypto.ZeroAddress,
			Expires: 1337,
		}
		m.names--
	} else if m.events > 0 {
		datalen := rand.Int() % 10
		data := make([]byte, datalen*32)
		topiclen := rand.Int() % 5
		topics := make([]binary.Word256, topiclen)
		row.EVMEvent = &dump.EVMEvent{
			ChainID: "MockyChain",
			Event: &exec.LogEvent{
				Address: crypto.ZeroAddress,
				Data:    data,
				Topics:  topics,
			},
		}
		m.events--
	} else {
		return nil, nil
	}

	return &row, nil
}

func BenchmarkLoadDump(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mock := MockDumpReader{
			accounts: 2000,
			storage:  1000,
			names:    100,
			events:   100000,
		}
		st, err := MakeGenesisState(dbm.NewMemDB(), &genesis.GenesisDoc{})
		require.NoError(b, err)
		err = st.LoadDump(&mock)
		require.NoError(b, err)
		err = st.InitialCommit()
		require.NoError(b, err)
	}
}
