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

package sqlsol

import (
	"fmt"

	"github.com/hyperledger/burrow/vent/types"
)

// BlockData contains EventData definition
type BlockData struct {
	Data types.EventData
}

// NewBlockData returns a pointer to an empty BlockData structure
func NewBlockData(height uint64) *BlockData {
	data := types.EventData{
		Tables:      make(map[string]types.EventDataTable),
		BlockHeight: height,
	}

	return &BlockData{
		Data: data,
	}
}

// AddRow appends a row to a specific table name in structure
func (b *BlockData) AddRow(tableName string, row types.EventDataRow) {
	if _, ok := b.Data.Tables[tableName]; !ok {
		b.Data.Tables[tableName] = types.EventDataTable{}
	}
	b.Data.Tables[tableName] = append(b.Data.Tables[tableName], row)
}

// GetRows gets data rows for a given table name from structure
func (b *BlockData) GetRows(tableName string) (types.EventDataTable, error) {
	if table, ok := b.Data.Tables[tableName]; ok {
		return table, nil
	}
	return nil, fmt.Errorf("GetRows: tableName does not exists as a table in data structure: %s ", tableName)
}

// PendingRows returns true if the given block has at least one pending row to upsert
func (b *BlockData) PendingRows(height uint64) bool {
	hasRows := false
	// TODO: understand why the guard on height is needed - what does it prevent?
	if b.Data.BlockHeight == height && len(b.Data.Tables) > 0 {
		hasRows = true
	}
	return hasRows
}
