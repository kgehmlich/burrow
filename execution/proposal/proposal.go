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

package proposal

import (
	"github.com/hyperledger/burrow/txs/payload"
)

type Reader interface {
	GetProposal(proposalHash []byte) (*payload.Ballot, error)
}

type Writer interface {
	// Updates the name entry creating it if it does not exist
	UpdateProposal(proposalHash []byte, proposal *payload.Ballot) error
	// Remove the name entry
	RemoveProposal(proposalHash []byte) error
}

type ReaderWriter interface {
	Reader
	Writer
}

type Iterable interface {
	IterateProposals(consumer func(proposalHash []byte, proposal *payload.Ballot) error) (err error)
}

type IterableReader interface {
	Iterable
	Reader
}

type IterableReaderWriter interface {
	Iterable
	ReaderWriter
}
