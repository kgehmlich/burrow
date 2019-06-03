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
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/hyperledger/burrow/logging/structure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFileLogger(t *testing.T) {
	f, err := ioutil.TempFile("", "TestNewFileLogger.log")
	require.NoError(t, err)
	logPath := f.Name()
	f.Close()
	fileLogger, err := NewFileLogger(logPath, JSONFormat)
	require.NoError(t, err)

	err = fileLogger.Log("foo", "bar")
	require.NoError(t, err)

	err = structure.Sync(fileLogger)
	require.NoError(t, err)

	bs, err := ioutil.ReadFile(logPath)

	require.NoError(t, err)
	assert.Equal(t, "{\"foo\":\"bar\"}\n", string(bs))
}

func TestFileTemplateParams(t *testing.T) {
	ftp := FileTemplateParams{
		Date: time.Now(),
	}
	fmt.Println(ftp.Timestamp())
}
