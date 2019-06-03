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

package validator

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiff(t *testing.T) {
	before := makeSet(
		1, 10,
		2, 20,
		3, 30,
		4, 40,
		5, 50,
	)
	after := makeSet(
		2, 22,
		3, 30,
		4, 40,
		5, 50,
		6, 60,
		7, 70,
	)
	expectedDiff := makeSet(
		1, 0,
		2, 22,
		6, 60,
		7, 70,
	)
	diff, err := Diff(before, after)
	require.NoError(t, err)
	assert.NoError(t, expectedDiff.Equal(diff))

	// And in reverse
	expectedDiff = makeSet(
		1, 10,
		2, 20,
		6, 0,
		7, 0,
	)
	diff, err = Diff(after, before)
	require.NoError(t, err)

	assert.NoError(t, expectedDiff.Equal(diff))
	fmt.Println(diff)
}

func makeSet(keyvals ...int) *Set {
	set := NewSet()
	if len(keyvals)%2 != 0 {
		panic(fmt.Errorf("cannot make set with odd number of keyvals: %d", len(keyvals)))
	}
	for i := 0; i < len(keyvals); i += 2 {
		set.ChangePower(pubKey(keyvals[i]), big.NewInt(int64(keyvals[i+1])))
	}
	return set
}
