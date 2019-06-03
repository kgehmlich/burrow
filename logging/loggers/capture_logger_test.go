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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlushCaptureLogger(t *testing.T) {
	outputLogger := newTestLogger()
	cl := NewCaptureLogger(outputLogger, 100, false)
	buffered := 50
	for i := 0; i < buffered; i++ {
		cl.Log("Foo", "Bar", "Index", i)
	}
	assert.True(t, outputLogger.empty())

	// Flush the ones we bufferred
	cl.Flush()
	_, err := outputLogger.logLines(buffered)
	assert.NoError(t, err)
}

func TestTeeCaptureLogger(t *testing.T) {
	outputLogger := newTestLogger()
	cl := NewCaptureLogger(outputLogger, 100, true)
	buffered := 50
	for i := 0; i < buffered; i++ {
		cl.Log("Foo", "Bar", "Index", i)
	}
	// Check passthrough to output
	ll, err := outputLogger.logLines(buffered)
	assert.NoError(t, err)
	assert.Equal(t, ll, cl.BufferLogger().FlushLogLines())

	cl.SetPassthrough(false)
	buffered = 110
	for i := 0; i < buffered; i++ {
		cl.Log("Foo", "Bar", "Index", i)
	}
	assert.True(t, outputLogger.empty())

	cl.Flush()
	_, err = outputLogger.logLines(100)
	assert.NoError(t, err)
	_, err = outputLogger.logLines(1)
	// Expect timeout
	assert.Error(t, err)
}
