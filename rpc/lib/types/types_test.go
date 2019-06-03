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

package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type SampleResult struct {
	Value string
}

func TestResponses(t *testing.T) {
	assert := assert.New(t)

	a := NewRPCSuccessResponse("1", &SampleResult{"hello"})
	b, _ := json.Marshal(a)
	s := `{"jsonrpc":"2.0","id":"1","result":{"Value":"hello"}}`
	assert.Equal(string(s), string(b))

	d := RPCParseError("1", errors.New("Hello world"))
	e, _ := json.Marshal(d)
	f := `{"jsonrpc":"2.0","id":"1","error":{"code":-32700,"message":"Parse Error","data":"Hello world"}}`
	assert.Equal(string(f), string(e))

	g := RPCMethodNotFoundError("2")
	h, _ := json.Marshal(g)
	i := `{"jsonrpc":"2.0","id":"2","error":{"code":-32601,"message":"Method Not Found"}}`
	assert.Equal(string(h), string(i))
}

func TestRPCError(t *testing.T) {
	assert.Equal(t, "RPC error 12 - Badness: One worse than a code 11",
		fmt.Sprintf("%v", &RPCError{
			Code:    12,
			Message: "Badness",
			Data:    "One worse than a code 11",
		}))

	assert.Equal(t, "RPC error 12 - Badness",
		fmt.Sprintf("%v", &RPCError{
			Code:    12,
			Message: "Badness",
		}))
}
