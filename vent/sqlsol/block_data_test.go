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
	"testing"

	"github.com/hyperledger/burrow/vent/sqlsol"
	"github.com/hyperledger/burrow/vent/types"
	"github.com/stretchr/testify/require"
)

func TestSetRow(t *testing.T) {
	t.Run("successfully sets a new data row", func(t *testing.T) {
		values := make(map[string]interface{})
		values["c1"] = "v1"
		values["c2"] = "v2"

		blockData := sqlsol.NewBlockData(44)
		blockData.AddRow("TEST_TABLE", types.EventDataRow{Action: types.ActionUpsert, RowData: values})

		rows, err := blockData.GetRows("TEST_TABLE")
		require.NoError(t, err)
		require.Equal(t, 1, len(rows))
		require.Equal(t, "v1", rows[0].RowData["c1"])
		require.Equal(t, "v2", rows[0].RowData["c2"])
	})
}

func TestGetBlockData(t *testing.T) {
	t.Run("successfully gets block data", func(t *testing.T) {
		blockData := sqlsol.NewBlockData(2)
		blk := blockData.Data
		require.EqualValues(t, 2, blk.BlockHeight)
	})
}

func TestPendingRows(t *testing.T) {
	t.Run("successfully returns true if a given block has pending rows to upsert", func(t *testing.T) {
		values := make(map[string]interface{})
		values["c1"] = "v1"
		values["c2"] = "v2"

		blockData := sqlsol.NewBlockData(99)
		blockData.AddRow("TEST_TABLE", types.EventDataRow{Action: types.ActionUpsert, RowData: values})

		hasRows := blockData.PendingRows(99)

		require.Equal(t, true, hasRows)
	})

	t.Run("successfully returns false if a given block does not have pending rows to upsert", func(t *testing.T) {
		values := make(map[string]interface{})
		values["c1"] = "v1"
		values["c2"] = "v2"

		blockData := sqlsol.NewBlockData(99)
		blockData.AddRow("TEST_TABLE", types.EventDataRow{Action: types.ActionUpsert, RowData: values})

		hasRows := blockData.PendingRows(88)

		require.Equal(t, false, hasRows)
	})

	t.Run("successfully returns false if a given block does not exists", func(t *testing.T) {
		blockData := sqlsol.NewBlockData(0)
		hasRows := blockData.PendingRows(999)

		require.Equal(t, false, hasRows)
	})
}
