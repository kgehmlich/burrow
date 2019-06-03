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

// SQLColumnType to store generic SQL column types
type SQLColumnType int

// generic SQL column types
const (
	SQLColumnTypeBool SQLColumnType = iota
	SQLColumnTypeByteA
	SQLColumnTypeInt
	SQLColumnTypeSerial
	SQLColumnTypeText
	SQLColumnTypeVarchar
	SQLColumnTypeTimeStamp
	SQLColumnTypeNumeric
	SQLColumnTypeJSON
	SQLColumnTypeBigInt
)

func (ct SQLColumnType) String() string {
	switch ct {
	case SQLColumnTypeBool:
		return "bool"
	case SQLColumnTypeByteA:
		return "bytea"
	case SQLColumnTypeInt:
		return "int"
	case SQLColumnTypeSerial:
		return "serial"
	case SQLColumnTypeText:
		return "text"
	case SQLColumnTypeVarchar:
		return "varchar"
	case SQLColumnTypeTimeStamp:
		return "timestamp"
	case SQLColumnTypeNumeric:
		return "numeric"
	case SQLColumnTypeJSON:
		return "json"
	case SQLColumnTypeBigInt:
		return "bigint"
	}
	return "unknown SQL type"
}

// IsNumeric determines if an sqlColumnType is numeric
func (ct SQLColumnType) IsNumeric() bool {
	return ct == SQLColumnTypeInt || ct == SQLColumnTypeSerial || ct == SQLColumnTypeNumeric || ct == SQLColumnTypeBigInt
}
