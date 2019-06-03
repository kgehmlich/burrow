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

// +build forensics

package forensics

import (
	"bytes"
	"fmt"

	"github.com/hyperledger/burrow/storage"

	"github.com/hyperledger/burrow/execution/state"

	"github.com/hyperledger/burrow/bcm"
	"github.com/hyperledger/burrow/binary"
	"github.com/hyperledger/burrow/core"
	"github.com/hyperledger/burrow/event"
	"github.com/hyperledger/burrow/execution"
	"github.com/hyperledger/burrow/execution/exec"
	"github.com/hyperledger/burrow/genesis"
	"github.com/hyperledger/burrow/logging"
	"github.com/hyperledger/burrow/txs"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/types"
)

type Replay struct {
	explorer   *bcm.BlockStore
	burrowDB   dbm.DB
	blockchain *bcm.Blockchain
	genesisDoc *genesis.GenesisDoc
	logger     *logging.Logger
}

type ReplayCapture struct {
	AppHashBefore binary.HexBytes
	AppHashAfter  binary.HexBytes
	TxExecutions  []*exec.TxExecution
}

func (recap *ReplayCapture) String() string {
	return fmt.Sprintf("ReplayCapture[%v -> %v]", recap.AppHashBefore, recap.AppHashAfter)
}

func NewReplay(dbDir string, genesisDoc *genesis.GenesisDoc, logger *logging.Logger) *Replay {
	//burrowDB := core.NewBurrowDB(dbDir)
	// Avoid writing through to underlying DB
	burrowDB := storage.NewCacheDB(dbm.NewDB(core.BurrowDBName, dbm.GoLevelDBBackend, dbDir))
	return &Replay{
		explorer:   bcm.NewBlockExplorer(dbm.LevelDBBackend, dbDir),
		burrowDB:   burrowDB,
		blockchain: bcm.NewBlockchain(burrowDB, genesisDoc),
		genesisDoc: genesisDoc,
		logger:     logger,
	}
}

func (re *Replay) LatestBlockchain() (*bcm.Blockchain, error) {
	_, blockchain, err := bcm.LoadOrNewBlockchain(re.burrowDB, re.genesisDoc, re.logger)
	if err != nil {
		return nil, err
	}
	re.blockchain = blockchain
	return blockchain, nil
}

func (re *Replay) State(height uint64) (*state.State, error) {
	return state.LoadState(re.burrowDB, execution.VersionAtHeight(height))
}

func (re *Replay) Block(height uint64) (*ReplayCapture, error) {
	recap := new(ReplayCapture)
	// Load and commit previous block
	block, err := re.explorer.Block(int64(height - 1))
	if err != nil {
		return nil, err
	}
	err = re.blockchain.CommitBlockAtHeight(block.Time, block.Hash(), block.Header.AppHash, uint64(block.Height))
	if err != nil {
		return nil, err
	}
	// block.AppHash is hash after txs from previous block have been applied - it's the state we want to load on top
	// of which we will reapply this block txs
	st, err := re.State(height - 1)
	if err != nil {
		return nil, err
	}
	// Load block for replay
	block, err = re.explorer.Block(int64(height))
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(st.Hash(), block.AppHash) {
		return nil, fmt.Errorf("state hash (%X) retrieved for block AppHash (%X) do not match",
			st.Hash(), block.AppHash)
	}
	recap.AppHashBefore = binary.HexBytes(block.AppHash)

	// Get our commit machinery
	committer := execution.NewBatchCommitter(st, execution.ParamsFromGenesis(re.genesisDoc), re.blockchain,
		event.NewEmitter(), re.logger)

	var txe *exec.TxExecution
	var execErr error
	_, err = block.Transactions(func(txEnv *txs.Envelope) (stop bool) {
		txe, execErr = committer.Execute(txEnv)
		if execErr != nil {
			return true
		}
		recap.TxExecutions = append(recap.TxExecutions, txe)
		return false
	})
	if err != nil {
		return nil, err
	}
	if execErr != nil {
		return nil, execErr
	}
	abciHeader := types.TM2PB.Header(&block.Header)
	recap.AppHashAfter, err = committer.Commit(&abciHeader)
	if err != nil {
		return nil, err
	}
	block, err = re.explorer.Block(int64(height + 1))
	if err != nil {
		return nil, err
	}
	fmt.Println(block.AppHash)
	return recap, nil
}

func (re *Replay) Blocks(startHeight, endHeight uint64) ([]*ReplayCapture, error) {
	var err error
	var st *state.State
	if startHeight > 1 {
		// Load and commit previous block
		block, err := re.explorer.Block(int64(startHeight - 1))
		if err != nil {
			return nil, err
		}
		err = re.blockchain.CommitBlockAtHeight(block.Time, block.Hash(), block.Header.AppHash, uint64(block.Height))
		if err != nil {
			return nil, err
		}
		// block.AppHash is hash after txs from previous block have been applied - it's the state we want to load on top
		// of which we will reapply this block txs
		st, err = re.State(startHeight - 1)
		if err != nil {
			return nil, err
		}
	} else {
		st, err = state.MakeGenesisState(re.burrowDB, re.genesisDoc)
		if err != nil {
			return nil, err
		}
		err = st.InitialCommit()
		if err != nil {
			return nil, err
		}
	}
	recaps := make([]*ReplayCapture, 0, endHeight-startHeight+1)
	for height := startHeight; height < endHeight; height++ {
		recap := new(ReplayCapture)
		// Load block for replay
		block, err := re.explorer.Block(int64(height))
		if err != nil {
			return nil, err
		}
		if height > 1 && !bytes.Equal(st.Hash(), block.AppHash) {
			return nil, fmt.Errorf("state hash (%X) retrieved for block AppHash (%X) do not match",
				st.Hash(), block.AppHash[:])
		}
		recap.AppHashBefore = binary.HexBytes(block.AppHash)

		// Get our commit machinery
		committer := execution.NewBatchCommitter(st, execution.ParamsFromGenesis(re.genesisDoc), re.blockchain,
			event.NewEmitter(), re.logger)

		var txe *exec.TxExecution
		var execErr error
		_, err = block.Transactions(func(txEnv *txs.Envelope) (stop bool) {
			txe, execErr = committer.Execute(txEnv)
			if execErr != nil {
				return true
			}
			recap.TxExecutions = append(recap.TxExecutions, txe)
			return false
		})
		if err != nil {
			return nil, err
		}
		if execErr != nil {
			return nil, execErr
		}
		abciHeader := types.TM2PB.Header(&block.Header)
		recap.AppHashAfter, err = committer.Commit(&abciHeader)
		if err != nil {
			return nil, err
		}
		recaps = append(recaps, recap)
	}
	return recaps, nil
}
