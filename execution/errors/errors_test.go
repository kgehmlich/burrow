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

package errors

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorCode_MarshalJSON(t *testing.T) {
	ec := NewException(ErrorCodeDataStackOverflow, "arrgh")
	bs, err := json.Marshal(ec)
	require.NoError(t, err)

	ecOut := new(Exception)
	err = json.Unmarshal(bs, ecOut)
	require.NoError(t, err)

	assert.Equal(t, ec, ecOut)
}

func TestCode_String(t *testing.T) {
	err := ErrorCodeCodeOutOfBounds
	fmt.Println(err.Error())
}

func TestFirstOnly(t *testing.T) {
	err := FirstOnly()
	// This will be a wrapped nil - it should not register as first error
	var ex CodedError = (*Exception)(nil)
	err.PushError(ex)
	// This one should
	realErr := ErrorCodef(ErrorCodeInsufficientBalance, "real error")
	err.PushError(realErr)
	assert.True(t, realErr.Equal(err.Error()))
}
