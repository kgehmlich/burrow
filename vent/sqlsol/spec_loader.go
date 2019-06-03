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

	"github.com/hyperledger/burrow/txs"
	"github.com/hyperledger/burrow/vent/types"
)

// SpecLoader loads spec files and parses them
func SpecLoader(specFileOrDirs []string, createBlkTxTables bool) (*Projection, error) {
	var projection *Projection
	var err error

	if len(specFileOrDirs) == 0 {
		return nil, fmt.Errorf("please provide a spec file or directory")
	}

	projection, err = NewProjectionFromFolder(specFileOrDirs...)
	if err != nil {
		return nil, fmt.Errorf("error parsing spec: %v", err)
	}

	if createBlkTxTables {
		// add block & tx to tables definition
		blkTxTables := getBlockTxTablesDefinition()

		for k, v := range blkTxTables {
			projection.Tables[k] = v
		}

	}

	return projection, nil
}

// getBlockTxTablesDefinition returns block & transaction structures
func getBlockTxTablesDefinition() types.EventTables {
	return types.EventTables{
		types.SQLBlockTableName: &types.SQLTable{
			Name: types.SQLBlockTableName,
			Columns: []*types.SQLTableColumn{
				{
					Name:    types.SQLColumnLabelHeight,
					Type:    types.SQLColumnTypeVarchar,
					Length:  100,
					Primary: true,
				},
				{
					Name:    types.SQLColumnLabelBlockHeader,
					Type:    types.SQLColumnTypeJSON,
					Primary: false,
				},
			},
		},

		types.SQLTxTableName: &types.SQLTable{
			Name: types.SQLTxTableName,
			Columns: []*types.SQLTableColumn{
				// transaction table
				{
					Name:    types.SQLColumnLabelHeight,
					Type:    types.SQLColumnTypeVarchar,
					Length:  100,
					Primary: true,
				},
				{
					Name:    types.SQLColumnLabelTxHash,
					Type:    types.SQLColumnTypeVarchar,
					Length:  txs.HashLengthHex,
					Primary: true,
				},
				{
					Name:    types.SQLColumnLabelIndex,
					Type:    types.SQLColumnTypeNumeric,
					Length:  0,
					Primary: false,
				},
				{
					Name:    types.SQLColumnLabelTxType,
					Type:    types.SQLColumnTypeVarchar,
					Length:  100,
					Primary: false,
				},
				{
					Name:    types.SQLColumnLabelEnvelope,
					Type:    types.SQLColumnTypeJSON,
					Primary: false,
				},
				{
					Name:    types.SQLColumnLabelEvents,
					Type:    types.SQLColumnTypeJSON,
					Primary: false,
				},
				{
					Name:    types.SQLColumnLabelResult,
					Type:    types.SQLColumnTypeJSON,
					Primary: false,
				},
				{
					Name:    types.SQLColumnLabelReceipt,
					Type:    types.SQLColumnTypeJSON,
					Primary: false,
				},
				{
					Name:    types.SQLColumnLabelException,
					Type:    types.SQLColumnTypeJSON,
					Primary: false,
				},
			},
		},
	}
}
