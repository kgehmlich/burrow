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
	"fmt"
	"sort"
)

type Balances []Balance

func (b Balance) String() string {
	return fmt.Sprintf("{%v: %d}", b.Type, b.Amount)
}

func New() Balances {
	return []Balance{}
}

func (bs Balances) Sort() Balances {
	sort.Stable(bs)
	return bs
}

func (bs Balances) Len() int {
	return len(bs)
}

func (bs Balances) Less(i, j int) bool {
	if bs[i].Type < bs[j].Type {
		return true
	}
	return bs[i].Type == bs[j].Type && bs[i].Amount < bs[j].Amount
}

func (bs Balances) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func (bs Balances) Add(ty Type, amount uint64) Balances {
	return append(bs, Balance{
		Type:   ty,
		Amount: amount,
	})
}

func (bs Balances) Native(amount uint64) Balances {
	return bs.Add(TypeNative, amount)
}

func (bs Balances) Power(amount uint64) Balances {
	return bs.Add(TypePower, amount)
}

func (bs Balances) Sum(bss ...Balances) Balances {
	return Sum(append(bss, bs)...)
}

func Sum(bss ...Balances) Balances {
	sum := New()
	sumMap := make(map[Type]uint64)
	for _, bs := range bss {
		for _, b := range bs {
			sumMap[b.Type] += b.Amount
		}
	}
	for k, v := range sumMap {
		sum = sum.Add(k, v)
	}
	sort.Stable(sum)
	return sum
}

func Native(native uint64) Balance {
	return Balance{
		Type:   TypeNative,
		Amount: native,
	}
}

func Power(power uint64) Balance {
	return Balance{
		Type:   TypePower,
		Amount: power,
	}
}

func (bs Balances) Has(ty Type) bool {
	for _, b := range bs {
		if b.Type == ty {
			return true
		}
	}
	return false
}

func (bs Balances) Get(ty Type) *uint64 {
	for _, b := range bs {
		if b.Type == ty {
			return &b.Amount
		}
	}
	return nil
}

func (bs Balances) GetFallback(ty Type, fallback uint64) uint64 {
	for _, b := range bs {
		if b.Type == ty {
			return b.Amount
		}
	}
	return fallback
}

func (bs Balances) GetNative(fallback uint64) uint64 {
	return bs.GetFallback(TypeNative, fallback)
}

func (bs Balances) GetPower(fallback uint64) uint64 {
	return bs.GetFallback(TypePower, fallback)
}

func (bs Balances) HasNative() bool {
	return bs.Has(TypeNative)
}

func (bs Balances) HasPower() bool {
	return bs.Has(TypePower)
}
