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

package state

import (
	"fmt"

	"github.com/hyperledger/burrow/execution/names"
)

var _ names.IterableReader = &State{}

func (s *ReadState) GetName(name string) (*names.Entry, error) {
	tree, err := s.Forest.Reader(keys.Name.Prefix())
	if err != nil {
		return nil, err
	}
	entryBytes := tree.Get(keys.Name.KeyNoPrefix(name))
	if entryBytes == nil {
		return nil, nil
	}

	return names.DecodeEntry(entryBytes)
}

func (ws *writeState) UpdateName(entry *names.Entry) error {
	tree, err := ws.forest.Writer(keys.Name.Prefix())
	if err != nil {
		return err
	}
	bs, err := entry.Encode()
	if err != nil {
		return err
	}
	tree.Set(keys.Name.KeyNoPrefix(entry.Name), bs)
	return nil
}

func (ws *writeState) RemoveName(name string) error {
	tree, err := ws.forest.Writer(keys.Name.Prefix())
	if err != nil {
		return err
	}
	tree.Delete(keys.Name.KeyNoPrefix(name))
	return nil
}

func (s *ReadState) IterateNames(consumer func(*names.Entry) error) error {
	tree, err := s.Forest.Reader(keys.Name.Prefix())
	if err != nil {
		return err
	}
	return tree.Iterate(nil, nil, true, func(key []byte, value []byte) error {
		entry, err := names.DecodeEntry(value)
		if err != nil {
			return fmt.Errorf("State.IterateNames() could not iterate over names: %v", err)
		}
		return consumer(entry)
	})
}
