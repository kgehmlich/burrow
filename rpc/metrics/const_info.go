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

package metrics

import (
	"github.com/hyperledger/burrow/acm/acmstate"
	"github.com/hyperledger/burrow/rpc"
	core_types "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

// For mocking purposes
type constInfo struct {
	acmstate.AccountStats
	*rpc.ResultUnconfirmedTxs
	*rpc.ResultStatus
	NodePeers  []core_types.Peer
	BlockMetas []*types.BlockMeta
}

func (is *constInfo) Status() (*rpc.ResultStatus, error) {
	return is.ResultStatus, nil
}

func (is *constInfo) UnconfirmedTxs(maxTxs int64) (*rpc.ResultUnconfirmedTxs, error) {
	return is.ResultUnconfirmedTxs, nil
}

func (is *constInfo) Blocks(minHeight, maxHeight int64) (*rpc.ResultBlocks, error) {
	var lastHeight uint64
	var lo, hi int
	for i, bm := range is.BlockMetas {
		height := bm.Header.Height
		if height < minHeight {
			lo = i + 1
		} else if height <= maxHeight {
			hi = i + 1
			lastHeight = uint64(height)
		}
	}
	return &rpc.ResultBlocks{
		LastHeight: lastHeight,
		BlockMetas: is.BlockMetas[lo:hi],
	}, nil
}

func (is *constInfo) Peers() []core_types.Peer {
	return is.NodePeers
}

func (is *constInfo) Stats() acmstate.AccountStatsGetter {
	return is
}

func (is *constInfo) GetAccountStats() acmstate.AccountStats {
	return is.AccountStats
}
