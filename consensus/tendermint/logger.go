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

package tendermint

import (
	"github.com/hyperledger/burrow/logging"
	"github.com/tendermint/tendermint/libs/log"
)

type tendermintLogger struct {
	logger *logging.Logger
}

func NewLogger(logger *logging.Logger) log.Logger {
	return &tendermintLogger{
		logger: logger,
	}
}

func (tml *tendermintLogger) Info(msg string, keyvals ...interface{}) {
	tml.logger.InfoMsg(msg, keyvals...)
}

func (tml *tendermintLogger) Error(msg string, keyvals ...interface{}) {
	tml.logger.InfoMsg(msg, keyvals...)
}

func (tml *tendermintLogger) Debug(msg string, keyvals ...interface{}) {
	tml.logger.TraceMsg(msg, keyvals...)
}

func (tml *tendermintLogger) With(keyvals ...interface{}) log.Logger {
	return &tendermintLogger{
		logger: tml.logger.With(keyvals...),
	}
}
