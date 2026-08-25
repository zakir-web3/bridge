package bridgehub

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

// WithdrawEventHash cast keccak "Withdraw(bytes32,address,address,address,uint256,uint256,uint64)"
const WithdrawEventHash = "0x7ba97cc0a03c9de2b8d972e5a42d8867d1c3483652b679f5b8b8b5b62a4ef576"

// BridgeSignatureSubmittedEventHash cast keccak "BridgeSignatureSubmitted(bytes32,address,uint64)"
const BridgeSignatureSubmittedEventHash = "0xc36d3c025c0f565ea8451b940f66ac343971b6b6d515ea860aed580c6408af2b"

func (b *BridgeHub) ProcessLog(ctx context.Context, log types.Log) error {
	if len(log.Topics) == 0 || log.Address != b.cfg.BridgeHubAddress {
		b.logger.Warn().
			Str("address", log.Address.Hex()).
			Str("expected", b.cfg.BridgeHubAddress.Hex()).
			Msg("skipping event")
		return nil
	}
	switch log.Topics[0] {
	case common.HexToHash(WithdrawEventHash):
		withdrawEvent, err := b.contract.ParseWithdraw(log)
		if err != nil {
			return errors.Wrap(err, "parse withdraw event")
		}
		if err = b.bc.Withdraw(ctx, b.GetChainID(), withdrawEvent); err != nil {
			return errors.Wrap(err, "handle withdraw confirm")
		}
		return nil
	case common.HexToHash(BridgeSignatureSubmittedEventHash):
		signatureSubmittedEvent, err := b.contract.ParseBridgeSignatureSubmitted(log)
		if err != nil {
			return errors.Wrap(err, "parse bridge signature submitted event")
		}
		if !b.HaveEnoughPower(signatureSubmittedEvent.TotalPower) {
			b.logger.Debug().Hex("messageHash", signatureSubmittedEvent.Message[:]).
				Uint64("totalPower", signatureSubmittedEvent.TotalPower).
				Msg("not enough power to handle bridge signature submitted")
			return nil
		}
		if signatureSubmittedEvent.Signer != b.GetFrom() {
			b.logger.Info().Hex("messageHash", signatureSubmittedEvent.Message[:]).
				Str("signer", signatureSubmittedEvent.Signer.Hex()).
				Msg("signature submitted by another signer, skipping")
			return nil
		}
		signs, err := b.contract.GetBridgeMsgSigns(ctx, signatureSubmittedEvent.Message)
		if err != nil {
			return err
		}
		if err = b.bc.BridgeSignatureSubmitted(ctx, b.GetChainID(), signs); err != nil {
			return errors.Wrap(err, "handle bridge signature submitted")
		}
		return nil
	default:
		return nil
	}
}

func (b *BridgeHub) GetFilterQuery(startBlock, endBlock uint64) []ethereum.FilterQuery {
	return []ethereum.FilterQuery{
		{
			FromBlock: big.NewInt(int64(startBlock)),
			ToBlock:   big.NewInt(int64(endBlock)),
			Addresses: []common.Address{b.cfg.BridgeHubAddress},
			Topics: [][]common.Hash{
				{
					common.HexToHash(WithdrawEventHash),
					common.HexToHash(BridgeSignatureSubmittedEventHash),
				},
			},
		},
	}
}
