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

package rpcevents

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockRange_Bounds(t *testing.T) {
	latestHeight := uint64(2344)
	br := &BlockRange{}
	start, end, streaming := br.Bounds(latestHeight)
	assert.Equal(t, latestHeight, start)
	assert.Equal(t, latestHeight+1, end)
	assert.False(t, streaming)
}
