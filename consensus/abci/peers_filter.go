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

package abci

import (
	"fmt"
	"strings"

	"github.com/hyperledger/burrow/consensus/tendermint/codes"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

const (
	peersFilterQueryPath = "/p2p/filter/"
)

func isPeersFilterQuery(query *abciTypes.RequestQuery) bool {
	return strings.HasPrefix(query.Path, peersFilterQueryPath)
}

func (app *App) peersFilter(reqQuery *abciTypes.RequestQuery, respQuery *abciTypes.ResponseQuery) {
	app.logger.TraceMsg("abci.App/Query peers filter query", "query_path", reqQuery.Path)
	path := strings.Split(reqQuery.Path, "/")
	if len(path) != 5 {
		panic(fmt.Errorf("invalid peers filter query path %v", reqQuery.Path))
	}

	filterType := path[3]
	peer := path[4]

	authorizedPeersID, authorizedPeersAddress := app.authorizedPeersProvider()
	var authorizedPeers []string
	switch filterType {
	case "id":
		authorizedPeers = authorizedPeersID
	case "addr":
		authorizedPeers = authorizedPeersAddress
	default:
		panic(fmt.Errorf("invalid peers filter query type %v", reqQuery.Path))
	}

	peerAuthorized := len(authorizedPeers) == 0
	for _, authorizedPeer := range authorizedPeers {
		if authorizedPeer == peer {
			peerAuthorized = true
			break
		}
	}

	if peerAuthorized {
		app.logger.TraceMsg("Peer sync authorized", "peer", peer)
		respQuery.Code = codes.PeerFilterAuthorizedCode
	} else {
		app.logger.InfoMsg("Peer sync forbidden", "peer", peer)
		respQuery.Code = codes.PeerFilterForbiddenCode
	}
}
