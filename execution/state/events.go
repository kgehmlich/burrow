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
	"io"

	"github.com/hyperledger/burrow/execution/exec"
)

func (ws *writeState) AddBlock(be *exec.BlockExecution) error {
	tree, err := ws.forest.Writer(keys.Event.Prefix())
	if err != nil {
		return err
	}
	txHashTree, err := ws.forest.Writer(keys.TxHash.Prefix())
	if err != nil {
		return err
	}
	// Index transactions so they can be retrieved by their TxHash
	for index, ev := range be.StreamEvents() {
		// Store with prefix for scanning later
		key := keys.Event.KeyNoPrefix(be.Height, uint64(index))
		bs, err := ev.Encode()
		if err != nil {
			return err
		}
		// Set StreamEvent itself
		tree.Set(key, bs)
		if ev.BeginTx != nil {
			// Set reference to TxExecution
			txHashTree.Set(keys.TxHash.KeyNoPrefix(ev.BeginTx.TxHeader.TxHash), key)
		}
	}
	return nil
}

func (s *ReadState) IterateStreamEvents(start, end *exec.StreamKey, consumer func(*exec.StreamEvent) error) error {
	tree, err := s.Forest.Reader(keys.Event.Prefix())
	if err != nil {
		return err
	}
	var startKey, endKey []byte
	if start != nil {
		startKey = keys.Event.KeyNoPrefix(start.Height, start.Index)
	}
	if end != nil {
		endKey = keys.Event.KeyNoPrefix(end.Height, end.Index)
	}
	return tree.Iterate(startKey, endKey, true, func(_, value []byte) error {
		txe, err := exec.DecodeStreamEvent(value)
		if err != nil {
			return fmt.Errorf("error unmarshalling BlockExecution in GetBlocks: %v", err)
		}
		return consumer(txe)
	})
}

func (s *ReadState) TxsAtHeight(height uint64) ([]*exec.TxExecution, error) {
	var stack exec.TxStack
	var txExecutions []*exec.TxExecution
	err := s.IterateStreamEvents(&exec.StreamKey{Height: height}, &exec.StreamKey{Height: height + 1},
		func(ev *exec.StreamEvent) error {
			// Keep trying to consume TxExecutions at from events at this height
			txe := stack.Consume(ev)
			if txe != nil {
				txExecutions = append(txExecutions, txe)
			}
			return nil
		})
	if err != nil && err != io.EOF {
		return nil, err
	}
	return txExecutions, nil
}

func (s *ReadState) StreamEvent(height, index uint64) (*exec.StreamEvent, error) {
	tree, err := s.Forest.Reader(keys.Event.Prefix())
	if err != nil {
		return nil, err
	}
	// Note: stored with prefix for scanning
	bs := tree.Get(keys.Event.KeyNoPrefix(height, index))
	if len(bs) == 0 {
		return nil, nil
	}
	return exec.DecodeStreamEvent(bs)
}

func (s *ReadState) TxByHash(txHash []byte) (*exec.TxExecution, error) {
	const errHeader = "TxByHash():"
	eventKeyByTxHashTree, err := s.Forest.Reader(keys.TxHash.Prefix())
	if err != nil {
		return nil, err
	}
	eventKey := eventKeyByTxHashTree.Get(keys.TxHash.KeyNoPrefix(txHash))
	if len(eventKey) == 0 {
		return nil, nil
	}
	eventTree, err := s.Forest.Reader(keys.Event.Prefix())
	if err != nil {
		return nil, err
	}
	bs := eventTree.Get(eventKey)
	if len(bs) == 0 {
		return nil, fmt.Errorf("%s could not retieve transaction with TxHash %X despite finding reference",
			errHeader, txHash)
	}
	ev, err := exec.DecodeStreamEvent(bs)
	if err != nil {
		return nil, err
	}
	if ev.BeginTx == nil {
		return nil, fmt.Errorf("%s StreamEvent %v is not a transaction despite being indexed as such",
			errHeader, ev)
	}

	start := new(exec.StreamKey)
	// Scan out position in storage
	err = keys.Event.ScanNoPrefix(eventKey, &start.Height, &start.Index)
	if err != nil {
		return nil, fmt.Errorf("%s could not scan height and index from tx key %X: %v", errHeader, eventKey, err)
	}
	// Iterate to end of block - we will break the iteration once we have scanned the tx so this is an upper bound
	end := &exec.StreamKey{Height: start.Height + 1}

	// Establish iteration state
	var stack exec.TxStack
	var txe *exec.TxExecution
	err = s.IterateStreamEvents(start, end, func(ev *exec.StreamEvent) error {
		txe = stack.Consume(ev)
		if txe != nil {
			return io.EOF
		}
		return nil
	})
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("%s error iterating over stream events %v", errHeader, err)
	}
	// Possibly nil if not found
	return txe, nil
}
