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

package jobs

import "testing"

func Test_matchInstanceName(t *testing.T) {
	type args struct {
		objectName     string
		deployInstance string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{
				objectName:     "contracts/storage.sol:SimpleConstructorArray",
				deployInstance: "SimpleConstructorArray",
			},
			true,
		},
		{
			"",
			args{
				objectName:     "storage.sol:SimpleConstructorArray",
				deployInstance: "simpleConstructorArray",
			},
			true,
		},
		{
			"",
			args{
				objectName:     "SimpleConstructorArray",
				deployInstance: "simpleconstructorarray",
			},
			true,
		},
		{
			"",
			args{
				objectName:     "",
				deployInstance: "Simpleconstructorarray",
			},
			false,
		},
		{
			"",
			args{
				objectName:     "SimpleConstructorArray:",
				deployInstance: "SimpleConstructorArray",
			},
			false,
		},
		{
			"",
			args{
				objectName:     ":",
				deployInstance: "SimpleConstructorArray",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchInstanceName(tt.args.objectName, tt.args.deployInstance); got != tt.want {
				t.Errorf("matchInstanceName() = %v, want %v", got, tt.want)
			}
		})
	}
}
