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

package rpcquery

import (
	"context"
	"fmt"

	"github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/acm/acmstate"
	"github.com/hyperledger/burrow/acm/validator"
	"github.com/hyperledger/burrow/bcm"
	"github.com/hyperledger/burrow/consensus/tendermint"
	"github.com/hyperledger/burrow/event/query"
	"github.com/hyperledger/burrow/execution/names"
	"github.com/hyperledger/burrow/execution/proposal"
	"github.com/hyperledger/burrow/execution/state"
	"github.com/hyperledger/burrow/logging"
	"github.com/hyperledger/burrow/rpc"
	"github.com/hyperledger/burrow/txs/payload"
	"github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

type queryServer struct {
	accounts    acmstate.IterableStatsReader
	nameReg     names.IterableReader
	proposalReg proposal.IterableReader
	blockchain  bcm.BlockchainInfo
	validators  validator.History
	nodeView    *tendermint.NodeView
	logger      *logging.Logger
}

var _ QueryServer = &queryServer{}

func NewQueryServer(state acmstate.IterableStatsReader, nameReg names.IterableReader, proposalReg proposal.IterableReader,
	blockchain bcm.BlockchainInfo, validators validator.History, nodeView *tendermint.NodeView, logger *logging.Logger) *queryServer {
	return &queryServer{
		accounts:    state,
		nameReg:     nameReg,
		proposalReg: proposalReg,
		blockchain:  blockchain,
		validators:  validators,
		nodeView:    nodeView,
		logger:      logger,
	}
}

func (qs *queryServer) Status(ctx context.Context, param *StatusParam) (*rpc.ResultStatus, error) {
	return rpc.Status(qs.blockchain, qs.validators, qs.nodeView, param.BlockTimeWithin, param.BlockSeenTimeWithin)
}

// Account state

func (qs *queryServer) GetAccount(ctx context.Context, param *GetAccountParam) (*acm.Account, error) {
	acc, err := qs.accounts.GetAccount(param.Address)
	if acc == nil {
		acc = &acm.Account{}
	}
	return acc, err
}

func (qs *queryServer) GetStorage(ctx context.Context, param *GetStorageParam) (*StorageValue, error) {
	val, err := qs.accounts.GetStorage(param.Address, param.Key)
	return &StorageValue{Value: val}, err
}

func (qs *queryServer) ListAccounts(param *ListAccountsParam, stream Query_ListAccountsServer) error {
	qry, err := query.NewOrEmpty(param.Query)
	var streamErr error
	err = qs.accounts.IterateAccounts(func(acc *acm.Account) error {
		if qry.Matches(acc.Tagged()) {
			return stream.Send(acc)
		} else {
			return nil
		}
	})
	if err != nil {
		return err
	}
	return streamErr
}

// Names

func (qs *queryServer) GetName(ctx context.Context, param *GetNameParam) (entry *names.Entry, err error) {
	entry, err = qs.nameReg.GetName(param.Name)
	if entry == nil && err == nil {
		err = fmt.Errorf("name %s not found", param.Name)
	}
	return
}

func (qs *queryServer) ListNames(param *ListNamesParam, stream Query_ListNamesServer) error {
	qry, err := query.NewOrEmpty(param.Query)
	if err != nil {
		return err
	}
	var streamErr error
	err = qs.nameReg.IterateNames(func(entry *names.Entry) error {
		if qry.Matches(entry.Tagged()) {
			return stream.Send(entry)
		} else {
			return nil
		}
	})
	if err != nil {
		return err
	}
	return streamErr
}

// Validators

func (qs *queryServer) GetValidatorSet(ctx context.Context, param *GetValidatorSetParam) (*ValidatorSet, error) {
	set := validator.Copy(qs.validators.Validators(0))
	return &ValidatorSet{
		Set: set.Validators(),
	}, nil
}

func (qs *queryServer) GetValidatorSetHistory(ctx context.Context, param *GetValidatorSetHistoryParam) (*ValidatorSetHistory, error) {
	lookback := int(param.IncludePrevious)
	switch {
	case lookback == 0:
		lookback = 1
	case lookback < 0 || lookback > state.DefaultValidatorsWindowSize:
		lookback = state.DefaultValidatorsWindowSize
	}
	height := qs.blockchain.LastBlockHeight()
	if height < uint64(lookback) {
		lookback = int(height)
	}
	history := &ValidatorSetHistory{}
	for i := 0; i < lookback; i++ {
		set := validator.Copy(qs.validators.Validators(i))
		vs := &ValidatorSet{
			Height: height - uint64(i),
			Set:    set.Validators(),
		}
		history.History = append(history.History, vs)
	}
	return history, nil
}

// proposals

func (qs *queryServer) GetProposal(ctx context.Context, param *GetProposalParam) (proposal *payload.Ballot, err error) {
	proposal, err = qs.proposalReg.GetProposal(param.Hash)
	if proposal == nil && err == nil {
		err = fmt.Errorf("proposal %x not found", param.Hash)
	}
	return
}

func (qs *queryServer) ListProposals(param *ListProposalsParam, stream Query_ListProposalsServer) error {
	var streamErr error
	err := qs.proposalReg.IterateProposals(func(hash []byte, ballot *payload.Ballot) error {
		if param.GetProposed() == false || ballot.ProposalState == payload.Ballot_PROPOSED {
			return stream.Send(&ProposalResult{Hash: hash, Ballot: ballot})
		} else {
			return nil
		}
	})
	if err != nil {
		return err
	}
	return streamErr
}

func (qs *queryServer) GetStats(ctx context.Context, param *GetStatsParam) (*Stats, error) {
	stats := qs.accounts.GetAccountStats()

	return &Stats{
		AccountsWithCode:    stats.AccountsWithCode,
		AccountsWithoutCode: stats.AccountsWithoutCode,
	}, nil
}

// Tendermint and blocks

func (qs *queryServer) GetBlockHeader(ctx context.Context, param *GetBlockParam) (*types.Header, error) {
	header, err := qs.blockchain.GetBlockHeader(param.Height)
	if err != nil {
		return nil, err
	}
	abciHeader := tmtypes.TM2PB.Header(header)
	return &abciHeader, nil
}
