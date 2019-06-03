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
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorsRing_AlterPower(t *testing.T) {
	vsBase := NewSet()
	powAInitial := int64(10000)
	vsBase.ChangePower(pubA, big.NewInt(powAInitial))

	vs := Copy(vsBase)
	vw := NewRing(vs, 3)

	// Just allowable validator tide
	var powA, powB, powC int64 = 7000, 23, 309
	powerChange, totalFlow, err := alterPowers(t, vw, powA, powB, powC)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(powA+powB+powC-powAInitial), powerChange)
	assert.Equal(t, big.NewInt(powAInitial/3-1), totalFlow)

	// This one is not
	vs = Copy(vsBase)
	vw = NewRing(vs, 5)
	powA, powB, powC = 7000, 23, 310
	powerChange, totalFlow, err = alterPowers(t, vw, powA, powB, powC)
	require.Error(t, err)

	powA, powB, powC = 7000, 23, 309
	powerChange, totalFlow, err = alterPowers(t, vw, powA, powB, powC)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(powA+powB+powC-powAInitial), powerChange)
	assert.Equal(t, big.NewInt(powAInitial/3-1), totalFlow)

	powA, powB, powC = 7000, 23, 309
	powerChange, totalFlow, err = alterPowers(t, vw, powA, powB, powC)
	require.NoError(t, err)
	assertZero(t, powerChange)
	assertZero(t, totalFlow)

	_, err = vw.AlterPower(pubA, big.NewInt(8000))
	assert.NoError(t, err)

	// Should fail - not enough flow left
	_, err = vw.AlterPower(pubB, big.NewInt(2000))
	assert.Error(t, err)

	// Take a bit off should work
	_, err = vw.AlterPower(pubA, big.NewInt(7000))
	assert.NoError(t, err)

	_, err = vw.AlterPower(pubB, big.NewInt(2000))
	assert.NoError(t, err)
	_, _, err = vw.Rotate()
	require.NoError(t, err)

	powerChange, totalFlow, err = alterPowers(t, vw, powA, powB, powC)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(-1977), powerChange)
	assert.Equal(t, big.NewInt(1977), totalFlow)

	powerChange, totalFlow, err = alterPowers(t, vw, powA, powB, powC)
	require.NoError(t, err)
	assertZero(t, powerChange)
	assert.Equal(t, big0, totalFlow)

	powerChange, totalFlow, err = alterPowers(t, vw, powA, powB, powC)
	require.NoError(t, err)
	assertZero(t, powerChange)
	assert.Equal(t, big0, totalFlow)
}

func TestRing_Rotate(t *testing.T) {
	ring := NewRing(nil, 3)
	err := ring.SetPower(pubA, big.NewInt(234))
	require.NoError(t, err)
	fmt.Println(printBuckets(ring))
	_, _, err = ring.Rotate()
	require.NoError(t, err)

	err = ring.SetPower(pubA, big.NewInt(111))
	require.NoError(t, err)
	fmt.Println(printBuckets(ring))
	_, _, err = ring.Rotate()
	require.NoError(t, err)
	fmt.Println(printBuckets(ring))

	err = ring.SetPower(pubB, big.NewInt(40))
	require.NoError(t, err)
	fmt.Println(printBuckets(ring))
	_, _, err = ring.Rotate()
	require.NoError(t, err)
	fmt.Println(printBuckets(ring))

	err = ring.SetPower(pubC, big.NewInt(99990))
	require.NoError(t, err)
	fmt.Println(printBuckets(ring))
	_, _, err = ring.Rotate()
	require.NoError(t, err)
	fmt.Println(printBuckets(ring))

	fmt.Println(ring.ValidatorChanges(1))
}

func printBuckets(ring *Ring) string {
	buf := new(bytes.Buffer)
	for i, b := range ring.OrderedBuckets() {
		buf.WriteString(fmt.Sprintf("%d: ", i))
		buf.WriteString(b.String())
		buf.WriteString("\n")
	}
	return buf.String()
}

func alterPowers(t testing.TB, vw *Ring, powA, powB, powC int64) (powerChange, totalFlow *big.Int, err error) {
	_, err = vw.AlterPower(pubA, big.NewInt(powA))
	if err != nil {
		return nil, nil, err
	}
	_, err = vw.AlterPower(pubB, big.NewInt(powB))
	if err != nil {
		return nil, nil, err
	}
	_, err = vw.AlterPower(pubC, big.NewInt(powC))
	if err != nil {
		return nil, nil, err
	}
	maxFlow := vw.Head().Previous.MaxFlow()
	powerChange, totalFlow, err = vw.Rotate()
	require.NoError(t, err)
	// totalFlow > maxFlow
	if totalFlow.Cmp(maxFlow) == 1 {
		return powerChange, totalFlow, fmt.Errorf("totalFlow (%v) exceeds maxFlow (%v)", totalFlow, maxFlow)
	}

	return powerChange, totalFlow, nil
}

// Since we have -0 and 0 with big.Int due to its representation with a neg flag
func assertZero(t testing.TB, i *big.Int) {
	assert.True(t, big0.Cmp(i) == 0, "expected 0 but got %v", i)
}
