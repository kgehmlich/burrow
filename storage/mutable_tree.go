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

package storage

import (
	"fmt"

	"github.com/tendermint/iavl"
	dbm "github.com/tendermint/tendermint/libs/db"
)

type MutableTree struct {
	*iavl.MutableTree
}

func NewMutableTree(db dbm.DB, cacheSize int) *MutableTree {
	tree := iavl.NewMutableTree(db, cacheSize)
	return &MutableTree{
		MutableTree: tree,
	}
}

func (mut *MutableTree) Load(version int64, overwriting bool) error {
	if version <= 0 {
		return fmt.Errorf("trying to load MutableTree from non-positive version: version %d", version)
	}
	var err error
	var treeVersion int64
	if overwriting {
		// Deletes all version above version!
		treeVersion, err = mut.MutableTree.LoadVersionForOverwriting(version)
	} else {
		treeVersion, err = mut.MutableTree.LoadVersion(version)
	}
	if err != nil {
		return fmt.Errorf("could not load current version of MutableTree (version %d): %v", version, err)
	}
	if treeVersion != version {
		return fmt.Errorf("tried to load version %d of MutableTree, but got version %d", version, treeVersion)
	}
	return nil
}

func (mut *MutableTree) Get(key []byte) []byte {
	_, bs := mut.MutableTree.Get(key)
	return bs
}

func (mut *MutableTree) GetImmutable(version int64) (*ImmutableTree, error) {
	tree, err := mut.MutableTree.GetImmutable(version)
	if err != nil {
		return nil, err
	}
	return &ImmutableTree{tree}, nil
}

// Get the current working tree as an ImmutableTree (for the methods - not immutable!)
func (mut *MutableTree) asImmutable() *ImmutableTree {
	return &ImmutableTree{mut.MutableTree.ImmutableTree}
}

func (mut *MutableTree) Iterate(start, end []byte, ascending bool, fn func(key []byte, value []byte) error) error {
	return mut.asImmutable().Iterate(start, end, ascending, fn)
}

func (mut *MutableTree) IterateWriteTree(start, end []byte, ascending bool, fn func(key []byte, value []byte) error) error {
	var err error
	mut.MutableTree.IterateRange(start, end, ascending, func(key, value []byte) bool {
		err = fn(key, value)
		if err != nil {
			// stop
			return true
		}
		return false
	})
	return err
}
