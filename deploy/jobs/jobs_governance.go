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

package jobs

import (
	"fmt"

	"github.com/hyperledger/burrow/logging"
	"github.com/hyperledger/burrow/txs/payload"

	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/deploy/def"
	"github.com/hyperledger/burrow/deploy/util"
	"github.com/hyperledger/burrow/execution/evm/abi"
)

func FormulateUpdateAccountJob(gov *def.UpdateAccount, account string, client *def.Client, logger *logging.Logger) (*payload.GovTx, []*abi.Variable, error) {
	gov.Source = FirstOf(gov.Source, account)
	perms := make([]string, len(gov.Permissions))

	for i, p := range gov.Permissions {
		perms[i] = string(p)
	}
	arg := &def.GovArg{
		Input:       gov.Source,
		Sequence:    gov.Sequence,
		Power:       gov.Power,
		Native:      gov.Native,
		Roles:       gov.Roles,
		Permissions: perms,
	}
	newAccountMatch := def.NewKeyRegex.FindStringSubmatch(gov.Target)
	if len(newAccountMatch) > 0 {
		keyName, curveType := def.KeyNameCurveType(newAccountMatch)
		publicKey, err := client.CreateKey(keyName, curveType, logger)
		if err != nil {
			return nil, nil, fmt.Errorf("could not create key for new account: %v", err)
		}
		arg.Address = publicKey.GetAddress().String()
		arg.PublicKey = publicKey.String()
	} else if len(gov.Target) == crypto.AddressHexLength {
		arg.Address = gov.Target
	} else {
		arg.PublicKey = gov.Target
	}

	tx, err := client.UpdateAccount(arg, logger)
	if err != nil {
		return nil, nil, err
	}

	return tx, util.Variables(arg), nil
}

func UpdateAccountJob(gov *def.UpdateAccount, account string, tx *payload.GovTx, client *def.Client, logger *logging.Logger) error {
	txe, err := client.SignAndBroadcast(tx, logger)
	if err != nil {
		return util.ChainErrorHandler(account, err, logger)
	}

	util.ReadTxSignAndBroadcast(txe, err, logger)
	if err != nil {
		return err
	}

	return nil
}
