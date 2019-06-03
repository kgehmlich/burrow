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

package sqlsol_test

import (
	"os"
	"testing"

	"github.com/hyperledger/burrow/vent/sqlsol"
	"github.com/hyperledger/burrow/vent/types"
	"github.com/stretchr/testify/require"
)

func TestSpecLoader(t *testing.T) {
	specFile := []string{os.Getenv("GOPATH") + "/src/github.com/hyperledger/burrow/vent/test/sqlsol_example.json"}
	dBBlockTx := true
	t.Run("successfully add block and transaction tables to event structures", func(t *testing.T) {
		projection, err := sqlsol.SpecLoader(specFile, dBBlockTx)
		require.NoError(t, err)

		require.Equal(t, 4, len(projection.Tables))

		require.Equal(t, types.SQLBlockTableName, projection.Tables[types.SQLBlockTableName].Name)

		require.Equal(t, types.SQLColumnLabelHeight,
			projection.Tables[types.SQLBlockTableName].GetColumn(types.SQLColumnLabelHeight).Name)

		require.Equal(t, types.SQLTxTableName, projection.Tables[types.SQLTxTableName].Name)

		require.Equal(t, types.SQLColumnLabelTxHash,
			projection.Tables[types.SQLTxTableName].GetColumn(types.SQLColumnLabelTxHash).Name)
	})
}
