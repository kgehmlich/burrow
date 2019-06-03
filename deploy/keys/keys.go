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

package keys

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/hyperledger/burrow/keys"
	"github.com/hyperledger/burrow/logging"
)

type LocalKeyClient struct {
	keys.KeyClient
}

var keysTimeout = 5 * time.Second

// Returns an initialized key client to a docker container
// running the keys server
// Adding the Ip address is optional and should only be used
// for passing data
func InitKeyClient(keysUrl string) (*LocalKeyClient, error) {
	aliveCh := make(chan struct{})
	localKeyClient, err := keys.NewRemoteKeyClient(keysUrl, logging.NewNoopLogger())
	if err != nil {
		return nil, err
	}

	err = localKeyClient.HealthCheck()

	go func() {
		for err != nil {
			err = localKeyClient.HealthCheck()
		}
		aliveCh <- struct{}{}
	}()
	select {
	case <-time.After(keysTimeout):
		return nil, fmt.Errorf("keys instance did not become responsive after %s: %v", keysTimeout, err)
	case <-aliveCh:
		return &LocalKeyClient{localKeyClient}, nil
	}
}

// Keyclient returns a list of keys that it is aware of.
// params:
// host - search for keys on the host
// container - search for keys on the container
// quiet - don't print output, just return the list you find
func (keys *LocalKeyClient) ListKeys(keysPath string, quiet bool, logger *logging.Logger) ([]string, error) {
	var result []string
	addrs, err := ioutil.ReadDir(keysPath)
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		result = append(result, addr.Name())
	}
	if !quiet {
		if len(addrs) == 0 {
			logger.InfoMsg("No keys found on host")
		} else {
			// First key.
			logger.InfoMsg("The keys on host", result)
		}
	}

	return result, nil
}
