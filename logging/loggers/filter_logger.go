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

package loggers

import (
	"github.com/go-kit/kit/log"
	"github.com/hyperledger/burrow/logging/structure"
)

// Filter logger allows us to filter lines logged to it before passing on to underlying
// output logger
// Creates a logger that removes lines from output when the predicate evaluates true
func FilterLogger(outputLogger log.Logger, predicate func(keyvals []interface{}) bool) log.Logger {
	return log.LoggerFunc(func(keyvals ...interface{}) error {
		// Always forward signals
		if structure.Signal(keyvals) != "" || !predicate(keyvals) {
			return outputLogger.Log(keyvals...)
		}
		return nil
	})
}
