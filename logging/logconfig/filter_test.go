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

package logconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildKeyValuesPredicateMatchAll(t *testing.T) {
	conf := []*KeyValuePredicateConfig{
		{
			KeyRegex:   "Foo",
			ValueRegex: "bar",
		},
	}
	kvp, err := BuildKeyValuesPredicate(conf, true)
	assert.NoError(t, err)
	assert.True(t, kvp([]interface{}{"Foo", "bar", "Bosh", "Bish"}))
}

func TestBuildKeyValuesPredicateMatchAny(t *testing.T) {
	conf := []*KeyValuePredicateConfig{
		{
			KeyRegex:   "Bosh",
			ValueRegex: "Bish",
		},
	}
	kvp, err := BuildKeyValuesPredicate(conf, false)
	assert.NoError(t, err)
	assert.True(t, kvp([]interface{}{"Foo", "bar", "Bosh", "Bish"}))
}

func TestExcludeAllFilterPredicate(t *testing.T) {
	fc := &FilterConfig{
		FilterMode: ExcludeWhenAllMatch,
		Predicates: []*KeyValuePredicateConfig{
			{
				KeyRegex:   "Bosh",
				ValueRegex: "Bish",
			},
			{
				KeyRegex:   "Bosh",
				ValueRegex: "Bash",
			},
		},
	}
	fp, err := BuildFilterPredicate(fc)
	assert.NoError(t, err)
	assert.False(t, fp([]interface{}{"Bosh", "Bash", "Shoes", 42}))
	assert.True(t, fp([]interface{}{"Bosh", "Bash", "Foo", "bar", "Shoes", 42, "Bosh", "Bish"}))
	assert.False(t, fp([]interface{}{"Food", 0.2, "Shoes", 42}))

}
func TestExcludeAnyFilterPredicate(t *testing.T) {
	fc := &FilterConfig{
		FilterMode: ExcludeWhenAnyMatches,
		Predicates: []*KeyValuePredicateConfig{
			{
				KeyRegex:   "Bosh",
				ValueRegex: "Bish",
			},
			{
				KeyRegex:   "Bosh",
				ValueRegex: "Bash",
			},
		},
	}
	fp, err := BuildFilterPredicate(fc)
	assert.NoError(t, err)
	assert.False(t, fp([]interface{}{"Foo", "bar", "Shoes", 42}))
	assert.True(t, fp([]interface{}{"Foo", "bar", "Shoes", 42, "Bosh", "Bish"}))
	assert.True(t, fp([]interface{}{"Food", 0.2, "Shoes", 42, "Bosh", "Bish"}))

}

func TestIncludeAllFilterPredicate(t *testing.T) {
	fc := &FilterConfig{
		FilterMode: IncludeWhenAllMatch,
		Predicates: []*KeyValuePredicateConfig{
			{
				KeyRegex:   "Bosh",
				ValueRegex: "Bish",
			},
			{
				KeyRegex:   "Planks",
				ValueRegex: "^0.2$",
			},
		},
	}
	fp, err := BuildFilterPredicate(fc)
	assert.NoError(t, err)
	assert.True(t, fp([]interface{}{"Foo", "bar", "Shoes", 42}))
	// Don't filter, it has all the required key values
	assert.False(t, fp([]interface{}{"Foo", "bar", "Planks", 0.2, "Shoes", 42, "imBoshy", "unBishy"}))
	assert.True(t, fp([]interface{}{"Foo", "bar", "Planks", 0.23, "Shoes", 42, "imBoshy", "unBishy"}))
	assert.True(t, fp([]interface{}{"Food", 0.2, "Shoes", 42}))
}

func TestIncludeAnyFilterPredicate(t *testing.T) {
	fc := &FilterConfig{
		FilterMode: IncludeWhenAnyMatches,
		Predicates: []*KeyValuePredicateConfig{
			{
				KeyRegex:   "Bosh",
				ValueRegex: "Bish",
			},
			{
				KeyRegex:   "^Shoes$",
				ValueRegex: "42",
			},
		},
	}
	fp, err := BuildFilterPredicate(fc)
	assert.NoError(t, err)
	assert.False(t, fp([]interface{}{"Foo", "bar", "Shoes", 3427}))
	assert.False(t, fp([]interface{}{"Foo", "bar", "Shoes", 42, "Bosh", "Bish"}))
	assert.False(t, fp([]interface{}{"Food", 0.2, "Shoes", 42}))
}

func TestKeyOnlyPredicate(t *testing.T) {

	fc := &FilterConfig{
		FilterMode: IncludeWhenAnyMatches,
		Predicates: []*KeyValuePredicateConfig{
			{
				KeyRegex: "Bosh",
			},
		},
	}
	fp, err := BuildFilterPredicate(fc)
	assert.NoError(t, err)
	assert.True(t, fp([]interface{}{"Foo", "bar", "Shoes", 3427}))
	assert.False(t, fp([]interface{}{"Foo", "bar", "Shoes", 42, "Bosh", "Bish"}))
	assert.True(t, fp([]interface{}{"Food", 0.2, "Shoes", 42}))
}
