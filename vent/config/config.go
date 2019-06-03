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

package config

import (
	"time"

	"github.com/hyperledger/burrow/vent/types"
)

const DefaultPostgresDBURL = "postgres://postgres@localhost:5432/postgres?sslmode=disable"

// VentConfig is a set of configuration parameters
type VentConfig struct {
	DBAdapter      string
	DBURL          string
	DBSchema       string
	GRPCAddr       string
	HTTPAddr       string
	LogLevel       string
	SpecFileOrDirs []string
	AbiFileOrDirs  []string
	DBBlockTx      bool
	// Announce status every AnnouncePeriod
	AnnounceEvery time.Duration
}

// DefaultFlags returns a configuration with default values
func DefaultVentConfig() *VentConfig {
	return &VentConfig{
		DBAdapter:     types.PostgresDB,
		DBURL:         DefaultPostgresDBURL,
		DBSchema:      "vent",
		GRPCAddr:      "localhost:10997",
		HTTPAddr:      "0.0.0.0:8080",
		LogLevel:      "debug",
		DBBlockTx:     false,
		AnnounceEvery: time.Second * 5,
	}
}
