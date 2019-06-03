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

package types

import (
	"fmt"
)

// SQLTable contains the structure of a SQL table,
type SQLTable struct {
	Name    string
	Columns []*SQLTableColumn
	// Map of channel name -> columns to be sent as payload on that channel
	NotifyChannels map[string][]string
	columns        map[string]*SQLTableColumn
}

func (table *SQLTable) GetColumn(columnName string) *SQLTableColumn {
	if table.columns == nil {
		table.columns = make(map[string]*SQLTableColumn, len(table.Columns))
		for _, column := range table.Columns {
			table.columns[column.Name] = column
		}
	}
	return table.columns[columnName]
}

// SQLTableColumn contains the definition of a SQL table column,
// the Order is given to be able to sort the columns to be created
type SQLTableColumn struct {
	Name    string
	Type    SQLColumnType
	Primary bool
	Length  int
}

func (col *SQLTableColumn) String() string {
	primaryString := ""
	if col.Primary {
		primaryString = " (primary)"
	}
	lengthString := ""
	if col.Length != 0 {
		lengthString = fmt.Sprintf(" (length %d)", col.Length)
	}
	return fmt.Sprintf("SQLTableColumn{%s%s: %v%s}",
		col.Name, primaryString, col.Type, lengthString)
}

func (col *SQLTableColumn) Equals(otherCol *SQLTableColumn) bool {
	columnA := *col
	columnB := *otherCol
	return columnA == columnB
}

// UpsertDeleteQuery contains query and values to upsert or delete row data
type UpsertDeleteQuery struct {
	Query    string
	Values   string
	Pointers []interface{}
}

// SQL log & dictionary tables
const (
	SQLLogTableName        = "_vent_log"
	SQLDictionaryTableName = "_vent_dictionary"
	SQLBlockTableName      = "_vent_block"
	SQLTxTableName         = "_vent_tx"
	SQLChainInfoTableName  = "_vent_chain"
)

// fixed sql column names in tables
const (
	// log
	SQLColumnLabelId          = "_id"
	SQLColumnLabelTimeStamp   = "_timestamp"
	SQLColumnLabelTableName   = "_tablename"
	SQLColumnLabelEventName   = "_eventname"
	SQLColumnLabelEventFilter = "_eventfilter"
	SQLColumnLabelHeight      = "_height"
	SQLColumnLabelTxHash      = "_txhash"
	SQLColumnLabelAction      = "_action"
	SQLColumnLabelDataRow     = "_datarow"
	SQLColumnLabelSqlStmt     = "_sqlstmt"
	SQLColumnLabelSqlValues   = "_sqlvalues"

	// dictionary
	SQLColumnLabelColumnName   = "_columnname"
	SQLColumnLabelColumnType   = "_columntype"
	SQLColumnLabelColumnLength = "_columnlength"
	SQLColumnLabelPrimaryKey   = "_primarykey"
	SQLColumnLabelColumnOrder  = "_columnorder"

	// chain info
	SQLColumnLabelBurrowVer = "_burrowversion"
	SQLColumnLabelChainID   = "_chainid"

	// context
	SQLColumnLabelIndex       = "_index"
	SQLColumnLabelEventType   = "_eventtype"
	SQLColumnLabelBlockHeader = "_blockheader"
	SQLColumnLabelTxType      = "_txtype"
	SQLColumnLabelEnvelope    = "_envelope"
	SQLColumnLabelEvents      = "_events"
	SQLColumnLabelResult      = "_result"
	SQLColumnLabelReceipt     = "_receipt"
	SQLColumnLabelException   = "_exception"
)

// labels for column mapping
const (
	// event related
	EventNameLabel = "eventName"
	EventTypeLabel = "eventType"

	// block related
	ChainIDLabel     = "chainid"
	BlockHeightLabel = "height"
	BlockHeaderLabel = "blockHeader"

	// transaction related
	TxTxTypeLabel    = "txType"
	TxTxHashLabel    = "txHash"
	TxIndexLabel     = "index"
	TxEnvelopeLabel  = "envelope"
	TxEventsLabel    = "events"
	TxResultLabel    = "result"
	TxReceiptLabel   = "receipt"
	TxExceptionLabel = "exception"
)
