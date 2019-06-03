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

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBurrow(t *testing.T) {
	var outputCount int
	out := &output{
		PrintfFunc: func(format string, args ...interface{}) {
			outputCount++
		},
		LogfFunc: func(format string, args ...interface{}) {
			outputCount++
		},
		FatalfFunc: func(format string, args ...interface{}) {
			t.Fatalf("fatalf called by burrow cmd: %s", fmt.Sprintf(format, args...))
		},
	}
	app := burrow(out)
	// Basic smoke test for cli config
	err := app.Run([]string{"burrow", "--version"})
	assert.NoError(t, err)
	err = app.Run([]string{"burrow", "spec", "--name-prefix", "foo", "-f1"})
	assert.NoError(t, err)
	err = app.Run([]string{"burrow", "configure"})
	assert.NoError(t, err)
	err = app.Run([]string{"burrow", "start", "--help"})
	assert.NoError(t, err)
	assert.True(t, outputCount > 0)
}
