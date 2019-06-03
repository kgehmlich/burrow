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

package payload

import (
	"crypto/sha256"
	"fmt"

	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func NewProposalTx(propsal *Proposal) *ProposalTx {
	return &ProposalTx{
		Proposal: propsal,
	}
}

func (tx *ProposalTx) Type() Type {
	return TypeProposal
}

func (tx *ProposalTx) GetInputs() []*TxInput {
	return []*TxInput{tx.Input}
}

func (tx *ProposalTx) String() string {
	return fmt.Sprintf("ProposalTx{%v}", tx.Proposal)
}

func (tx *ProposalTx) Any() *Any {
	return &Any{
		ProposalTx: tx,
	}
}

func DecodeProposal(proposalBytes []byte) (*Proposal, error) {
	proposal := new(Proposal)
	err := cdc.UnmarshalBinaryBare(proposalBytes, proposal)
	if err != nil {
		return nil, err
	}
	return proposal, nil
}

func (p *Proposal) Encode() ([]byte, error) {
	return cdc.MarshalBinaryBare(p)
}

func (p *Proposal) Hash() []byte {
	bs, err := p.Encode()
	if err != nil {
		panic("failed to encode Proposal")
	}

	hash := sha256.Sum256(bs)

	return hash[:]
}

func (p *Proposal) String() string {
	return ""
}

func (v *Vote) String() string {
	return v.Address.String()
}

func DecodeBallot(ballotBytes []byte) (*Ballot, error) {
	ballot := new(Ballot)
	err := cdc.UnmarshalBinaryBare(ballotBytes, ballot)
	if err != nil {
		return nil, err
	}
	return ballot, nil
}

func (p *Ballot) Encode() ([]byte, error) {
	return cdc.MarshalBinaryBare(p)
}
