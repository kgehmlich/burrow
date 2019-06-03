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
	"github.com/stretchr/testify/require"
)

func Test_sortKeyvals(t *testing.T) {
	keyvals := []interface{}{"foo", 3, "bar", 5}
	indices := map[string]int{"foo": 1, "bar": 0}
	sortKeyvals(indices, keyvals)
	assert.Equal(t, []interface{}{"bar", 5, "foo", 3}, keyvals)
}

func TestSortLogger(t *testing.T) {
	testLogger := newTestLogger()
	sortLogger := SortLogger(testLogger, "foo", "bar", "baz")
	sortLogger.Log([][]int{}, "bar", "foo", 3, "baz", "horse", "crabs", "cycle", "bar", 4, "ALL ALONE")
	sortLogger.Log("foo", 0)
	sortLogger.Log("bar", "foo", "foo", "baz")
	lines, err := testLogger.logLines(3)
	require.NoError(t, err)
	// non string keys sort after string keys, specified keys sort before unspecifed keys, specified key sort in order
	assert.Equal(t, [][]interface{}{
		{"foo", 3, "bar", 4, "baz", "horse", [][]int{}, "bar", "crabs", "cycle", "ALL ALONE"},
		{"foo", 0},
		{"foo", "baz", "bar", "foo"},
	}, lines)
}
