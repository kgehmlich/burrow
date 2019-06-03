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

import "github.com/hyperledger/burrow/vent/logger"

// SQLConnection stores parameters to build a new db connection & initialize the database
type SQLConnection struct {
	DBAdapter string
	DBURL     string
	DBSchema  string
	Log       *logger.Logger
}

// SQLCleanDBQuery stores queries needed to clean the database
type SQLCleanDBQuery struct {
	SelectChainIDQry    string
	DeleteChainIDQry    string
	InsertChainIDQry    string
	SelectDictionaryQry string
	DeleteDictionaryQry string
	DeleteLogQry        string
}
