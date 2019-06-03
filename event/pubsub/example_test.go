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

package pubsub_test

import (
	"context"
	"testing"

	"github.com/hyperledger/burrow/event/pubsub"
	"github.com/hyperledger/burrow/event/query"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	s := pubsub.NewServer()
	s.Start()
	defer s.Stop()

	ctx := context.Background()
	ch, err := s.Subscribe(ctx, "example-client", query.MustParse("abci.account.name='John'"), 1)
	require.NoError(t, err)
	err = s.PublishWithTags(ctx, "Tombstone", query.TagMap(map[string]interface{}{"abci.account.name": "John"}))
	require.NoError(t, err)
	assertReceive(t, "Tombstone", ch)
}
