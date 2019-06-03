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
	"errors"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
)

const logLineTimeout = time.Second

type testLogger struct {
	channelLogger *ChannelLogger
	logLineCh     chan ([]interface{})
	err           error
}

func (tl *testLogger) empty() bool {
	return tl.channelLogger.BufferLength() == 0
}

func (tl *testLogger) logLines(numberOfLines int) ([][]interface{}, error) {
	logLines := make([][]interface{}, numberOfLines)
	for i := 0; i < numberOfLines; i++ {
		select {
		case logLine := <-tl.logLineCh:
			logLines[i] = logLine
		case <-time.After(logLineTimeout):
			return logLines, fmt.Errorf("timed out waiting for log line "+
				"(waited %s)", logLineTimeout)
		}
	}
	return logLines, nil
}

func (tl *testLogger) Log(keyvals ...interface{}) error {
	tl.channelLogger.Log(keyvals...)
	return tl.err
}

func newErrorLogger(errMessage string) *testLogger {
	return makeTestLogger(errors.New(errMessage))
}

func newTestLogger() *testLogger {
	return makeTestLogger(nil)
}

func makeTestLogger(err error) *testLogger {
	cl := NewChannelLogger(100)
	logLineCh := make(chan []interface{})
	go cl.DrainForever(log.LoggerFunc(func(keyvals ...interface{}) error {
		logLineCh <- keyvals
		return nil
	}), nil)
	return &testLogger{
		channelLogger: cl,
		logLineCh:     logLineCh,
		err:           err,
	}
}

// Utility function that returns a slice of log lines.
// Takes a variadic argument of log lines as a list of key value pairs delimited
// by the empty string and splits
func logLines(keyvals ...string) [][]interface{} {
	llines := make([][]interface{}, 0)
	line := make([]interface{}, 0)
	for _, kv := range keyvals {
		if kv == "" {
			llines = append(llines, line)
			line = make([]interface{}, 0)
		} else {
			line = append(line, kv)
		}
	}
	if len(line) > 0 {
		llines = append(llines, line)
	}
	return llines
}
