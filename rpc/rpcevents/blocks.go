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

package rpcevents

import (
	"github.com/hyperledger/burrow/execution/exec"
)

// Get bounds suitable for events.Provider
func (br *BlockRange) Bounds(latestBlockHeight uint64) (startHeight, endHeight uint64, streaming bool) {
	// End bound is exclusive in state.GetEvents so we increment the height
	return br.GetStart().Bound(latestBlockHeight), br.GetEnd().Bound(latestBlockHeight) + 1,
		br.GetEnd().GetType() == Bound_STREAM
}

func (b *Bound) Bound(latestBlockHeight uint64) uint64 {
	if b == nil {
		return latestBlockHeight
	}
	switch b.Type {
	case Bound_ABSOLUTE:
		return b.GetIndex()
	case Bound_RELATIVE:
		if b.Index < latestBlockHeight {
			return latestBlockHeight - b.Index
		}
		return 0
	case Bound_FIRST:
		return 0
	case Bound_LATEST, Bound_STREAM:
		return latestBlockHeight
	default:
		return latestBlockHeight
	}
}

func AbsoluteBound(index uint64) *Bound {
	return &Bound{
		Index: index,
		Type:  Bound_ABSOLUTE,
	}
}

func RelativeBound(index uint64) *Bound {
	return &Bound{
		Index: index,
		Type:  Bound_RELATIVE,
	}
}

func LatestBound() *Bound {
	return &Bound{
		Type: Bound_LATEST,
	}
}

func StreamBound() *Bound {
	return &Bound{
		Type: Bound_STREAM,
	}
}

func NewBlockRange(start, end *Bound) *BlockRange {
	return &BlockRange{
		Start: start,
		End:   end,
	}
}

func AbsoluteRange(start, end uint64) *BlockRange {
	return NewBlockRange(AbsoluteBound(start), AbsoluteBound(end))
}

func SingleBlock(height uint64) *BlockRange {
	return AbsoluteRange(height, height+1)
}

func ConsumeBlockExecutions(stream ExecutionEvents_StreamClient, consumer func(*exec.BlockExecution) error) error {
	var be *exec.BlockExecution
	var err error
	for be, err = exec.ConsumeBlockExecution(stream); err == nil; be, err = exec.ConsumeBlockExecution(stream) {
		err = consumer(be)
		if err != nil {
			return err
		}
	}
	return err
}
