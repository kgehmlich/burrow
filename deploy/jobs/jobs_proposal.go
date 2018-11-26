package jobs

import (
	"fmt"

	"github.com/hyperledger/burrow/binary"
	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/deploy/def"
	"github.com/hyperledger/burrow/deploy/loader"
	"github.com/hyperledger/burrow/deploy/proposals"
	"github.com/hyperledger/burrow/deploy/util"
	"github.com/hyperledger/burrow/txs"
	"github.com/hyperledger/burrow/txs/payload"

	log "github.com/sirupsen/logrus"
)

func recurseJobs(proposeBatch *payload.BatchTx, jobs []*def.Job, prop *def.Proposal, do *def.DeployArgs, parentScript *def.Playbook, client *def.Client) error {
	script := def.Playbook{Jobs: jobs, Account: useDefault(prop.Source, parentScript.Account), Parent: parentScript}

	for _, job := range script.Jobs {
		load, err := job.Payload()
		if err != nil {
			return fmt.Errorf("could not get Job payload: %v", load)
		}

		err = util.PreProcessFields(load, do, &script, client)
		if err != nil {
			return err
		}
		// Revalidate with possible replacements
		err = load.Validate()
		if err != nil {
			return fmt.Errorf("error validating job %s after pre-processing variables: %v", job.Name, err)
		}

		switch load.(type) {
		case *def.Meta:
			announceProposalJob(job.Name, "UpdateAccount")
			// load the package
			log.WithField("=>", job.Meta.File).Info("Loading sub YAML")
			metaScript, err := loader.LoadPackage(job.Meta.File)
			if err != nil {
				return err
			}
			err = recurseJobs(proposeBatch, metaScript.Jobs, prop, do, &script, client)
			if err != nil {
				return err
			}

		case *def.UpdateAccount:
			announceProposalJob(job.Name, "UpdateAccount")
			tx, _, err := FormulateUpdateAccountJob(job.UpdateAccount, script.Account, client)
			if err != nil {
				return err
			}
			proposeBatch.Txs = append(proposeBatch.Txs, &payload.Any{GovTx: tx})

		case *def.RegisterName:
			announceProposalJob(job.Name, "RegisterName")
			txs, err := FormulateRegisterNameJob(job.RegisterName, do, script.Account, client)
			if err != nil {
				return err
			}
			for _, tx := range txs {
				proposeBatch.Txs = append(proposeBatch.Txs, &payload.Any{NameTx: tx})
			}
		case *def.Call:
			announceProposalJob(job.Name, "Call")
			tx, err := FormulateCallJob(job.Call, do, &script, client)
			if err != nil {
				return err
			}
			proposeBatch.Txs = append(proposeBatch.Txs, &payload.Any{CallTx: tx})
		case *def.Deploy:
			announceProposalJob(job.Name, "Deploy")
			deployTxs, _, err := FormulateDeployJob(job.Deploy, do, &script, client, job.Intermediate)
			if err != nil {
				return err
			}
			var deployAddress crypto.Address
			// Predict address
			callee, err := crypto.AddressFromHexString(job.Deploy.Source)
			if err != nil {
				return err
			}
			for _, tx := range deployTxs {
				proposeBatch.Txs = append(proposeBatch.Txs, &payload.Any{CallTx: tx})
				txEnv := txs.NewTx(tx)

				deployAddress = crypto.NewContractAddress(callee, txEnv.Hash())
			}
			job.Result = deployAddress.String()
		case *def.Permission:
			announceProposalJob(job.Name, "Permission")
			tx, err := FormulatePermissionJob(job.Permission, script.Account, client)
			if err != nil {
				return err
			}
			proposeBatch.Txs = append(proposeBatch.Txs, &payload.Any{PermsTx: tx})
		case *def.Send:
			announceProposalJob(job.Name, "Send")
			tx, err := FormulateSendJob(job.Send, script.Account, client)
			if err != nil {
				return err
			}
			proposeBatch.Txs = append(proposeBatch.Txs, &payload.Any{SendTx: tx})
		default:
			return fmt.Errorf("jobs %s illegal job type for proposal", job.Name)
		}
	}

	return nil
}

func ProposalJob(prop *def.Proposal, do *def.DeployArgs, parentScript *def.Playbook, client *def.Client) (string, error) {
	var proposeBatch payload.BatchTx

	err := recurseJobs(&proposeBatch, prop.Jobs, prop, do, parentScript, client)
	if err != nil {
		return "", err
	}

	proposal := payload.Proposal{Name: prop.Name, Description: prop.Description, BatchTx: &proposeBatch}

	proposalInput, err := client.TxInput(prop.ProposalAddress, "", prop.ProposalSequence, false)
	if err != nil {
		return "", err
	}
	proposal.BatchTx.Inputs = []*payload.TxInput{proposalInput}
	proposalHash := proposal.Hash()

	var proposalTx *payload.ProposalTx
	if do.ProposeVerify {
		ballot, err := client.GetProposal(proposalHash)
		if err != nil {
			log.Warnf("Proposal could NOT be verified, error %v", err)
			return "", err
		}

		err = proposals.ProposalExpired(ballot.Proposal, client)
		if err != nil {
			log.Warnf("Proposal verify FAILED: %v", err)
			return "", err
		}

		log.Warnf("Proposal VERIFY SUCCESSFUL")
		log.Warnf("Proposal has %d votes:", len(ballot.Votes))
		for _, v := range ballot.Votes {
			log.Warnf("\t%s\n", v.Address)
		}

		return "", err
	} else if do.ProposeVote {
		ballot, err := client.GetProposal(proposalHash)
		if err != nil {
			log.Warnf("Proposal could not be found: %v", err)
			return "", err
		}

		err = proposals.ProposalExpired(ballot.Proposal, client)
		if err != nil {
			log.Warnf("Proposal error: %v", err)
			return "", err
		}

		// proposal is there and current, let's vote for it
		input, err := client.TxInput(parentScript.Account, "", prop.Sequence, true)
		if err != nil {
			return "", err
		}

		log.Warnf("Voting for proposal with hash: %x\n", proposalHash)

		h := binary.HexBytes(proposalHash)
		proposalTx = &payload.ProposalTx{ProposalHash: &h, VotingWeight: 1, Input: input}
	} else if do.ProposeCreate {
		input, err := client.TxInput(useDefault(prop.Source, parentScript.Account), "", prop.Sequence, true)
		if err != nil {
			return "", err
		}
		log.Warnf("Creating Proposal with hash: %x\n", proposalHash)

		proposalTx = &payload.ProposalTx{VotingWeight: 1, Input: input, Proposal: &proposal}
	} else {
		log.Errorf("please specify one of --proposal-create, --proposal-vote, --proposal-verify")
		return "", nil
	}

	txe, err := client.SignAndBroadcast(proposalTx)
	if err != nil {
		var err = util.ChainErrorHandler(proposalTx.Input.Address.String(), err)
		return "", err
	}

	result := fmt.Sprintf("%X", txe.Receipt.TxHash)

	return result, nil
}
