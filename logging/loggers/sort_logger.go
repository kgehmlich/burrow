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

package loggers

import (
	"sort"

	"github.com/go-kit/kit/log"
)

type sortableKeyvals struct {
	indices map[string]int
	keyvals []interface{}
	len     int
}

func sortKeyvals(indices map[string]int, keyvals []interface{}) {
	sort.Stable(sortable(indices, keyvals))
}

func sortable(indices map[string]int, keyvals []interface{}) *sortableKeyvals {
	return &sortableKeyvals{
		indices: indices,
		keyvals: keyvals,
		len:     len(keyvals) / 2,
	}
}

func (skv *sortableKeyvals) Len() int {
	return skv.len
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (skv *sortableKeyvals) Less(i, j int) bool {
	return skv.keyRank(i) < skv.keyRank(j)
}

// Swap swaps the elements with indexes i and j.
func (skv *sortableKeyvals) Swap(i, j int) {
	keyIdx, keyJdx := i*2, j*2
	valIdx, valJdx := keyIdx+1, keyJdx+1
	keyI, valI := skv.keyvals[keyIdx], skv.keyvals[valIdx]
	skv.keyvals[keyIdx], skv.keyvals[valIdx] = skv.keyvals[keyJdx], skv.keyvals[valJdx]
	skv.keyvals[keyJdx], skv.keyvals[valJdx] = keyI, valI
}

func (skv *sortableKeyvals) keyRank(i int) int {
	// Check there is a key at this index
	key, ok := skv.keyvals[i*2].(string)
	if !ok {
		// Sort keys not provided after those that have been but maintain relative order
		return len(skv.indices) + i
	}
	// See if we have been provided an explicit rank/order for the key
	idx, ok := skv.indices[key]
	if !ok {
		// Sort keys not provided after those that have been but maintain relative order
		return len(skv.indices) + i
	}
	return idx
}

// Provides a logger that sorts key-values with keys in keys before other key-values
func SortLogger(outputLogger log.Logger, keys ...string) log.Logger {
	indices := make(map[string]int, len(keys))
	for i, k := range keys {
		indices[k] = i
	}
	return log.LoggerFunc(func(keyvals ...interface{}) error {
		sortKeyvals(indices, keyvals)
		return outputLogger.Log(keyvals...)
	})
}
