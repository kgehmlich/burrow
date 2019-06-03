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

type Type uint32

const (
	TypeNative Type = 1
	TypePower  Type = 2
)

var nameFromType = map[Type]string{
	TypeNative: "Native",
	TypePower:  "Power",
}

var typeFromName = make(map[string]Type)

func init() {
	for t, n := range nameFromType {
		typeFromName[n] = t
	}
}

func TypeFromString(name string) Type {
	return typeFromName[name]
}

func (typ Type) String() string {
	name, ok := nameFromType[typ]
	if ok {
		return name
	}
	return "UnknownBalanceType"
}

func (typ Type) MarshalText() ([]byte, error) {
	return []byte(typ.String()), nil
}

func (typ *Type) UnmarshalText(data []byte) error {
	*typ = TypeFromString(string(data))
	return nil
}

// Protobuf support
func (typ Type) Marshal() ([]byte, error) {
	return typ.MarshalText()
}

func (typ *Type) Unmarshal(data []byte) error {
	return typ.UnmarshalText(data)
}
