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

package query

import (
	"testing"

	"github.com/hyperledger/burrow/logging/structure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryBuilder(t *testing.T) {
	qb := NewBuilder()
	qry, err := qb.Query()
	require.NoError(t, err)
	assert.Equal(t, emptyString, qry.String())

	qb = qb.AndGreaterThanOrEqual("foo.size", 45)
	qry, err = qb.Query()
	require.NoError(t, err)
	assert.Equal(t, "foo.size >= 45", qry.String())

	qb = qb.AndEquals("bar.name", "marmot")
	qry, err = qb.Query()
	require.NoError(t, err)
	assert.Equal(t, "foo.size >= 45 AND bar.name = 'marmot'", qry.String())

	assert.True(t, qry.Matches(makeTagMap("foo.size", 80, "bar.name", "marmot")))
	assert.False(t, qry.Matches(makeTagMap("foo.size", 8, "bar.name", "marmot")))
	assert.False(t, qry.Matches(makeTagMap("foo.size", 80, "bar.name", "marot")))

	qb = qb.AndContains("bar.desc", "burrow")
	qry, err = qb.Query()
	require.NoError(t, err)
	assert.Equal(t, "foo.size >= 45 AND bar.name = 'marmot' AND bar.desc CONTAINS 'burrow'", qry.String())

	assert.True(t, qry.Matches(makeTagMap("foo.size", 80, "bar.name", "marmot", "bar.desc", "lives in a burrow")))
	assert.False(t, qry.Matches(makeTagMap("foo.size", 80, "bar.name", "marmot", "bar.desc", "lives in a shoe")))

	qb = NewBuilder().AndEquals("foo", "bar")
	qb = qb.And(NewBuilder().AndGreaterThanOrEqual("frogs", 4))
	qry, err = qb.Query()
	require.NoError(t, err)
	assert.Equal(t, "foo = 'bar' AND frogs >= 4", qry.String())
}

func makeTagMap(keyvals ...interface{}) TagMap {
	tmap := make(TagMap)
	for i := 0; i < len(keyvals); i += 2 {
		tmap[keyvals[i].(string)] = structure.StringifyKey(keyvals[i+1])
	}
	return tmap
}
