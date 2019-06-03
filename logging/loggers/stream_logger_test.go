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
	"bytes"
	"testing"

	"github.com/hyperledger/burrow/logging/structure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStreamLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	logger, err := NewStreamLogger(buf, LogfmtFormat)
	require.NoError(t, err)
	err = logger.Log("oh", "my")
	require.NoError(t, err)

	err = structure.Sync(logger)
	require.NoError(t, err)

	assert.Equal(t, "oh=my\n", string(buf.Bytes()))
}

func TestNewTemplateLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	logger, err := NewTemplateLogger(buf, "Why Hello {{.name}}", []byte{'\n'})
	require.NoError(t, err)
	err = logger.Log("name", "Marjorie Stewart-Baxter", "fingertip_width_cm", float32(1.34))
	require.NoError(t, err)
	err = logger.Log("name", "Fred")
	require.NoError(t, err)
	assert.Equal(t, "Why Hello Marjorie Stewart-Baxter\nWhy Hello Fred\n", buf.String())
}
