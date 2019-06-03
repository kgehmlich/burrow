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

package exec

type CallType uint32

const (
	CallTypeCall     = CallType(0x00)
	CallTypeCode     = CallType(0x01)
	CallTypeDelegate = CallType(0x02)
	CallTypeStatic   = CallType(0x03)
	CallTypeSNative  = CallType(0x04)
)

var nameFromCallType = map[CallType]string{
	CallTypeCall:     "Call",
	CallTypeCode:     "CallCode",
	CallTypeDelegate: "DelegateCall",
	CallTypeStatic:   "StaticCall",
	CallTypeSNative:  "SNativeCall",
}

var callTypeFromName = make(map[string]CallType)

func init() {
	for t, n := range nameFromCallType {
		callTypeFromName[n] = t
	}
}

func CallTypeFromString(name string) CallType {
	return callTypeFromName[name]
}

func (ct CallType) String() string {
	name, ok := nameFromCallType[ct]
	if ok {
		return name
	}
	return "UnknownCallType"
}

func (ct CallType) MarshalText() ([]byte, error) {
	return []byte(ct.String()), nil
}

func (ct *CallType) UnmarshalText(data []byte) error {
	*ct = CallTypeFromString(string(data))
	return nil
}
