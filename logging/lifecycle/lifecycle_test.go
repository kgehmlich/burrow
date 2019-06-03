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

package lifecycle

import (
	"os"
	"testing"

	"bufio"

	"github.com/stretchr/testify/assert"
)

func TestNewLoggerFromLoggingConfig(t *testing.T) {
	reader := CaptureStderr(t, func() {
		logger, err := NewLoggerFromLoggingConfig(nil)
		assert.NoError(t, err)
		logger.Info.Log("Quick", "Test")
	})
	line, _, err := reader.ReadLine()
	assert.NoError(t, err)
	lineString := string(line)
	assert.NotEmpty(t, lineString)
}

func CaptureStderr(t *testing.T, runner func()) *bufio.Reader {
	stderr := os.Stderr
	defer func() {
		os.Stderr = stderr
	}()
	r, w, err := os.Pipe()
	assert.NoError(t, err, "Couldn't make fifo")
	os.Stderr = w

	runner()

	return bufio.NewReader(r)
}
