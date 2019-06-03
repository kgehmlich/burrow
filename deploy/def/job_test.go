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

package def

import (
	"strings"
	"testing"

	"github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/execution/evm/abi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJob_Validate(t *testing.T) {
	address := acm.GeneratePrivateAccountFromSecret("blah").GetAddress()
	job := &Job{
		Result: "brian",
		// This should pass emptiness validation
		Variables: []*abi.Variable{},
		QueryAccount: &QueryAccount{
			Account: address.String(),
			Field:   "bar",
		},
	}
	err := job.Validate()
	require.Error(t, err)
	errs := strings.Split(err.Error(), ";")
	if !assert.Len(t, errs, 2, "Should have two validation error from omitted name and included result") {
		t.Logf("Validation error was: %v", err)
	}

	job = &Job{
		Name: "Any kind of job",
		Account: &Account{
			Address: address.String(),
		},
	}
	err = job.Validate()
	require.NoError(t, err)

	job.Account.Address = "blah"
	err = job.Validate()
	require.NoError(t, err)
}
