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
	"testing"

	"github.com/hyperledger/burrow/crypto"
	"github.com/stretchr/testify/require"
)

func TestPackage_Validate(t *testing.T) {
	address := crypto.Address{3, 4}.String()
	pkgs := &Playbook{
		Jobs: []*Job{{
			Name: "CallJob",
			Call: &Call{
				Sequence:    "13",
				Destination: address,
			},
		}},
	}
	err := pkgs.Validate()
	require.NoError(t, err)

	pkgs.Jobs = append(pkgs.Jobs, &Job{
		Name: "Foo",
		Account: &Account{
			Address: address,
		},
	})
	err = pkgs.Validate()
	require.NoError(t, err)

	// cannot set two job fields
	pkgs.Jobs[1].QueryAccount = &QueryAccount{
		Account: address,
		Field:   "Foo",
	}
	err = pkgs.Validate()
	require.Error(t, err)

	pkgs = &Playbook{
		Jobs: []*Job{{
			Name: "UpdateAccount",
			UpdateAccount: &UpdateAccount{
				Target:   address,
				Sequence: "13",
				Native:   "333",
			},
		}},
	}
	err = pkgs.Validate()
	require.NoError(t, err)
}
