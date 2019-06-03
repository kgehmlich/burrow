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

package bcm

import (
	"fmt"
	"runtime/debug"

	"github.com/hyperledger/burrow/txs"
	"github.com/tendermint/tendermint/blockchain"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/state"
	"github.com/tendermint/tendermint/types"
)

type BlockStore struct {
	txDecoder txs.Decoder
	state.BlockStoreRPC
}

func NewBlockStore(blockStore state.BlockStoreRPC) *BlockStore {
	return &BlockStore{
		txDecoder:     txs.NewAminoCodec(),
		BlockStoreRPC: blockStore,
	}
}

func NewBlockExplorer(dbBackendType db.DBBackendType, dbDir string) *BlockStore {
	return NewBlockStore(blockchain.NewBlockStore(db.NewDB("blockstore", dbBackendType, dbDir)))
}

func (bs *BlockStore) Block(height int64) (_ *Block, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("BlockStore.Block(): could not get block at height %v: %v", height, r)
		}
	}()

	tmBlock := bs.LoadBlock(height)
	if tmBlock == nil {
		return nil, fmt.Errorf("could not pull block at height: %v", height)
	}
	return NewBlock(bs.txDecoder, tmBlock), nil
}

func (bs *BlockStore) BlockMeta(height int64) (_ *types.BlockMeta, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("BlockStore.BlockMeta() could not get BlockMeta at height %d: %v\n%s",
				height, r, debug.Stack())
		}
	}()
	return bs.LoadBlockMeta(height), nil
}

// Iterate over blocks between start (inclusive) and end (exclusive)
func (bs *BlockStore) Blocks(start, end int64, iter func(*Block) (stop bool)) (stopped bool, err error) {
	if end > 0 && start >= end {
		return false, fmt.Errorf("end height must be strictly greater than start height")
	}
	if start <= 0 {
		// From first block
		start = 1
	}
	if end < 0 {
		// -1 means include the very last block so + 1 for offset
		end = bs.Height() + end + 1
	}

	for height := start; height <= end; height++ {
		block, err := bs.Block(height)
		if err != nil {
			return false, err
		}
		if iter(block) {
			return true, nil
		}
	}

	return false, nil
}
