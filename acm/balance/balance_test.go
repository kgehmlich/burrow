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

package balance

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	one := New().Power(23223).Native(34).Native(1111)
	two := New().Power(3).Native(22)
	sum := one.Sum(two)
	assert.Equal(t, New().Power(23226).Native(1167).Sort(), sum)
}

func TestSort(t *testing.T) {
	balances := New().Power(232).Native(2523543).Native(232).Power(2).Power(4).Native(1)
	sortedBalances := New().Native(1).Native(232).Native(2523543).Power(2).Power(4).Power(232)
	sort.Sort(balances)
	assert.Equal(t, sortedBalances, balances)
}
