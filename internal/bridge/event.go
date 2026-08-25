package bridge

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/zakir-web3/bridge/internal/contract"
)

// ERC20TransferEventHash cast keccak "Transfer(address,address,uint256)"
const ERC20TransferEventHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

// DepositEventHash cast keccak "Deposit(address,address,address,uint256)"
const DepositEventHash = "0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96"

// FinalizedWithdrawalEventHash cast keccak "FinalizedWithdrawal(bytes32,address,address,address,uint256,uint64)"
const FinalizedWithdrawalEventHash = "0x04cafa25a7826c4415eac718e45fc84f69b5539748e0206adbae70f919887548"

func (b *Bridge) ProcessLog(ctx context.Context, log types.Log) error {
	if len(log.Topics) == 0 {
		return nil
	}
	switch log.Topics[0] {
	case common.HexToHash(ERC20TransferEventHash):
		if !b.cfg.IsBridgeToken(log.Address) {
			b.logger.Warn().
				Str("address", log.Address.Hex()).
				Msg("skipping bridge token transfer event")
			return nil
		}
		transferEvent, err := b.erc20.ParseTransfer(log)
		if err != nil {
			return errors.Wrap(err, "parse transfer event")
		}

		if transferEvent.To != b.cfg.BridgeAddress {
			return nil
		}
		depositEvent, err := b.QueryDepositEvent(ctx, log.TxHash)
		if err != nil {
			return errors.Wrap(err, "query deposit event")
		}
		if depositEvent != nil {
			transferEvent.To = depositEvent.Destination
		} else {
			transferEvent.To = transferEvent.From
		}
		b.logger.Info().Str("user", transferEvent.From.Hex()).
			Str("destination", transferEvent.To.Hex()).
			Str("amount", transferEvent.Value.String()).
			Msg("bridge token transfer to bridge address")
		if err = b.bridgeHub.DepositConfirm(ctx, b.GetChainID(), transferEvent); err != nil {
			return errors.Wrap(err, "handle bridge hub transfer")
		}
		return nil
	case common.HexToHash(FinalizedWithdrawalEventHash):
		if log.Address != b.cfg.BridgeAddress {
			b.logger.Warn().
				Str("address", log.Address.Hex()).
				Str("expected", b.cfg.BridgeAddress.Hex()).
				Msg("skipping finalized withdrawal event")
			return nil
		}
		withdrawEvent, err := b.contract.ParseFinalizedWithdrawal(log)
		if err != nil {
			return errors.Wrap(err, "parse finalized withdrawal event")
		}
		if err = b.bridgeHub.FinalizedWithdrawal(ctx, b.GetChainID(), withdrawEvent); err != nil {
			return errors.Wrap(err, "handle bridge hub finalized withdraw")
		}
		return nil
	default:
		return nil
	}
}

func (b *Bridge) GetFilterQuery(startBlock, endBlock uint64) []ethereum.FilterQuery {
	// Return multiple filter queries for different event types
	return []ethereum.FilterQuery{
		// Query for ERC20 transfer events from bridge tokens
		{
			FromBlock: big.NewInt(int64(startBlock)),
			ToBlock:   big.NewInt(int64(endBlock)),
			Addresses: b.cfg.BridgeTokens,
			Topics: [][]common.Hash{
				{
					common.HexToHash(ERC20TransferEventHash),
				},
				{},
				{
					common.BytesToHash(b.cfg.BridgeAddress.Bytes()),
				},
			},
		},
		// Query for finalized withdrawal events from bridge address
		{
			FromBlock: big.NewInt(int64(startBlock)),
			ToBlock:   big.NewInt(int64(endBlock)),
			Addresses: []common.Address{b.cfg.BridgeAddress},
			Topics: [][]common.Hash{
				{
					common.HexToHash(FinalizedWithdrawalEventHash),
				},
			},
		},
	}
}

func (b *Bridge) QueryDepositEvent(ctx context.Context, txHash common.Hash) (*contract.BridgeDeposit, error) {
	receipt, err := b.GetClient().TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get transaction receipt for %s", txHash.Hex())
	}

	for _, log := range receipt.Logs {
		if log.Address != b.cfg.BridgeAddress {
			continue
		}

		if len(log.Topics) == 0 || log.Topics[0] != common.HexToHash(DepositEventHash) {
			continue
		}

		depositEvent, err := b.contract.ParseDeposit(*log)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse deposit event from log %d", log.Index)
		}

		b.logger.Debug().
			Str("txHash", txHash.Hex()).
			Str("user", depositEvent.User.Hex()).
			Str("destination", depositEvent.Destination.Hex()).
			Str("token", depositEvent.Token.Hex()).
			Str("amount", depositEvent.Amount.String()).
			Uint("logIndex", log.Index).
			Msg("found deposit event in transaction receipt")

		return depositEvent, nil
	}
	return nil, nil
}
