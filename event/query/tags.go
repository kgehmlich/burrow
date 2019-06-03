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

package query

import (
	"sort"
	"strings"
)

type Tagged interface {
	Keys() []string
	Get(key string) (value string, ok bool)
	// Len returns the number of tags.
	Len() int
}

type TagMap map[string]interface{}

func MapFromTagged(tagged Tagged) map[string]interface{} {
	tags := make(map[string]interface{})
	for _, k := range tagged.Keys() {
		v, ok := tagged.Get(k)
		if ok {
			tags[k] = v
		}
	}
	return tags
}

func (ts TagMap) Get(key string) (value string, ok bool) {
	var vint interface{}
	vint, ok = ts[key]
	if !ok {
		return "", false
	}
	return StringFromValue(vint), true
}

func (ts TagMap) Len() int {
	return len(ts)
}

func (ts TagMap) Map() map[string]interface{} {
	return ts
}

func (ts TagMap) Keys() []string {
	keys := make([]string, 0, len(ts))
	for k := range ts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

type CombinedTags struct {
	keys   []string
	ks     map[string][]Tagged
	concat bool
}

func NewCombinedTags() *CombinedTags {
	return &CombinedTags{
		ks: make(map[string][]Tagged),
	}
}

func MergeTags(tags ...Tagged) *CombinedTags {
	ct := NewCombinedTags()
	ct.AddTags(false, tags...)
	return ct
}

func ConcatTags(tags ...Tagged) *CombinedTags {
	ct := NewCombinedTags()
	ct.AddTags(true, tags...)
	return ct
}

// Adds each of tagsList to CombinedTags - choosing whether values for the same key should
// be concatenated or whether the first should value should stick
func (ct *CombinedTags) AddTags(concat bool, tagsList ...Tagged) {
	for _, t := range tagsList {
		for _, k := range t.Keys() {
			if len(ct.ks[k]) == 0 {
				ct.keys = append(ct.keys, k)
				// Store reference to key-holder amongst Taggeds
				ct.ks[k] = append(ct.ks[k], t)
			} else if concat {
				// Store additional tag reference only if concat, otherwise first key-holder wins
				ct.ks[k] = append(ct.ks[k], t)
			}
		}
	}
}

func (ct *CombinedTags) Get(key string) (string, bool) {
	ts := ct.ks[key]
	if len(ts) == 0 {
		return "", false
	}
	values := make([]string, 0, len(ts))
	for _, t := range ts {
		value, ok := t.Get(key)
		if ok {
			values = append(values, value)
		}
	}
	if len(values) == 0 {
		return "", false
	}
	return strings.Join(values, MultipleValueTagSeparator), true
}

func (ct *CombinedTags) Len() (length int) {
	return len(ct.keys)
}

func (ct *CombinedTags) Keys() []string {
	return ct.keys
}
