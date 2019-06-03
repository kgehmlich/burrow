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
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

var pubA = pubKey(1)
var pubB = pubKey(2)
var pubC = pubKey(3)
var big2 = big.NewInt(2)

func TestBucket_AlterPower(t *testing.T) {
	base := NewBucket()
	err := base.SetPower(pubA, new(big.Int).Sub(maxTotalVotingPower, big3))
	require.NoError(t, err)

	bucket := NewBucket(base.Next)

	flow, err := bucket.AlterPower(pubA, new(big.Int).Sub(maxTotalVotingPower, big2))
	require.NoError(t, err)
	require.Equal(t, big1.Int64(), flow.Int64())

	flow, err = bucket.AlterPower(pubA, new(big.Int).Sub(maxTotalVotingPower, big1))
	require.NoError(t, err)
	require.Equal(t, big2.Int64(), flow.Int64())

	flow, err = bucket.AlterPower(pubA, maxTotalVotingPower)
	require.NoError(t, err)
	require.Equal(t, big3.Int64(), flow.Int64())

	flow, err = bucket.AlterPower(pubA, new(big.Int).Add(maxTotalVotingPower, big1))
	require.Error(t, err, "should fail as we would breach total power")

	flow, err = bucket.AlterPower(pubB, big1)
	require.Error(t, err, "should fail as we would breach total power")

	// Drop A and raise B - should now succeed
	flow, err = bucket.AlterPower(pubA, new(big.Int).Sub(maxTotalVotingPower, big1))
	require.NoError(t, err)
	require.Equal(t, big2.Int64(), flow.Int64())

	flow, err = bucket.AlterPower(pubB, big1)
	require.NoError(t, err)
	require.Equal(t, big1.Int64(), flow.Int64())
}

//func setPower(t *testing.T, id crypto.PublicKey, bucket *Bucket, power int64) {
//	err := bucket.SetPower(id, power)
//}
