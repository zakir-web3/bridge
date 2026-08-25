package bridge

import (
	"bytes"
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/zakir-web3/bridge/internal/contract"
	"github.com/zakir-web3/bridge/internal/evm"
)

const WithdrawEIP712Type = "Withdraw(address user,address destination,address token,uint256 amount,uint256 chainId,uint64 nonce)"

type HubContract interface {
	DepositConfirm(ctx context.Context, chainId *big.Int, transfer ...*contract.BridgeERC20Transfer) error
	SubmitSignatures(ctx context.Context, chainId *big.Int, signatures ...contract.CrossChainMessage) error
	FinalizedWithdrawal(ctx context.Context, chainId *big.Int, withdrawals ...*contract.BridgeFinalizedWithdrawal) error
	GetPendingMsgs(ctx context.Context) ([][32]byte, error)
	GetValidatorSet() *contract.ValidatorSet
}

type Bridge struct {
	*evm.AccountManager
	logger           zerolog.Logger
	cfg              Config
	contract         *contract.Bridge
	erc20            *contract.BridgeERC20
	validatorSet     contract.ValidatorSet
	bridgeHub        HubContract
	domainSeparator  common.Hash
	withdrawTypeHash []byte
}

func NewBridge(ctx context.Context, cfg Config, bridgeHub HubContract) (*Bridge, error) {
	cli, err := evm.NewClient(ctx, cfg.NodeURL, &cfg.RetryConfig)
	if err != nil {
		return nil, err
	}

	erc20, err := contract.NewBridgeERC20(common.Address{}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new BridgeERC20 contract")
	}

	c, err := contract.NewBridge(cfg.BridgeAddress, cli)
	if err != nil {
		return nil, errors.Wrap(err, "new Bridge contract")
	}

	acc, err := evm.NewAccountManager(ctx, cfg.AccConfig, cli)
	if err != nil {
		return nil, err
	}
	logger := log.With().
		Uint64("chainId", acc.GetChainID().Uint64()).
		Str("module", "bridge").Logger()
	logger.Info().Str("bridge", cfg.BridgeAddress.Hex()).Str("validator", acc.GetFrom().Hex()).Msg("initialized Bridge")

	domainSeparator, err := c.GetDomainSeparator(ctx)
	if err != nil {
		return nil, err
	}
	logger.Info().Hex("domainSeparator", domainSeparator.Bytes()).Msg("get domain separator from contract")

	validatorSet := bridgeHub.GetValidatorSet()
	if !validatorSet.IsValidator(acc.GetFrom()) {
		return nil, errors.New("not a validator")
	}
	logger.Info().Strs("validators", validatorSet.ValidatorsStr()).
		Uint64("epoch", validatorSet.Epoch).
		Uints64("power", validatorSet.Powers).
		Msg("get validator set from contract")

	return &Bridge{
		AccountManager:   acc,
		logger:           logger,
		cfg:              cfg,
		contract:         c,
		erc20:            erc20,
		validatorSet:     *validatorSet,
		domainSeparator:  domainSeparator,
		bridgeHub:        bridgeHub,
		withdrawTypeHash: crypto.Keccak256([]byte(WithdrawEIP712Type)),
	}, nil
}

func (b *Bridge) Withdraw(ctx context.Context, chainId *big.Int, withdraws ...*contract.BridgeHubWithdraw) error {
	signatures := make([]contract.CrossChainMessage, 0, len(withdraws))
	for _, w := range withdraws {
		typedData := w.ToTypedData(chainId, b.GetChainID(), b.cfg.BridgeAddress)
		domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
		if err != nil {
			return errors.Wrap(err, "hash domain separator for withdraw")
		}
		if !bytes.Equal(domainSeparator, b.domainSeparator.Bytes()) {
			return errors.Errorf("domain separator mismatch: expected %s, got %s",
				b.domainSeparator.Hex(), domainSeparator.String())
		}
		messageRawData, err := typedData.EncodeData(typedData.PrimaryType, typedData.Message, 1)
		if err != nil {
			return errors.Wrap(err, "encode data for withdraw")
		}
		signature, err := b.SignTypedData(typedData)
		if err != nil {
			return err
		}
		signatures = append(signatures, contract.CrossChainMessage{
			DomainSeparator: common.BytesToHash(domainSeparator),
			Signature:       signature,
			MessageRawData:  messageRawData,
		})
		b.logger.Debug().Hex("domainSeparator", domainSeparator).
			Hex("messageHash", crypto.Keccak256(messageRawData)).
			Any("message", typedData.Message).
			Msg("withdraw message prepared")
		b.logger.Info().Str("user", w.User.Hex()).
			Uint64("nonce", w.Nonce).
			Msg("withdraw signature created")
	}
	if len(signatures) == 0 {
		return nil
	}
	return b.bridgeHub.SubmitSignatures(ctx, b.GetChainID(), signatures...)
}

func (b *Bridge) BridgeSignatureSubmitted(ctx context.Context, _ *big.Int, signs ...*contract.MessageSignature) error {
	if !b.cfg.EnableRequestWithdraw {
		b.logger.Info().Msg("request withdraw is disabled, skipping withdraw")
		return nil
	}
	withdrawals := make([]contract.WithdrawalRequest, 0, len(signs))
	for _, sign := range signs {
		if !b.HaveEnoughPower(sign.TotalPower) {
			b.logger.Debug().
				Uint64("totalPower", sign.TotalPower).
				Msg("not enough power to handle bridge signature submitted")
			continue
		}
		if len(sign.RawData) < 32 || !bytes.Equal(sign.RawData[:32], b.withdrawTypeHash) {
			b.logger.Warn().Hex("expectedTypeHash", b.withdrawTypeHash).
				Hex("actualTypeHash", sign.RawData).
				Msg("withdraw type hash mismatch, skipping")
			continue
		}
		messageHash := common.BytesToHash(crypto.Keccak256(sign.RawData))
		requestedWithdrawals, err := b.contract.GetRequestedWithdrawals(ctx, messageHash)
		if err != nil {
			return err
		}
		if requestedWithdrawals.RequestedTime != 0 {
			b.logger.Info().Hex("messageHash", messageHash.Bytes()).
				Uint64("nonce", requestedWithdrawals.Nonce).
				Msg("withdrawal already requested, skipping")
			continue
		}
		withdrawalRequest, err := sign.ToWithdrawalRequest(b.validatorSet)
		if err != nil {
			return err
		}
		b.logger.Info().Str("user", withdrawalRequest.User.Hex()).
			Uint64("nonce", withdrawalRequest.Nonce).
			Msg("bridge signature submitted for withdrawal")
		withdrawals = append(withdrawals, withdrawalRequest)
	}
	if len(withdrawals) == 0 {
		return nil
	}
	opts := b.NewTransactOpts(ctx)
	opts.GasLimit = b.cfg.RequestWithdrawGasLimit
	transaction, err := b.contract.BatchedRequestWithdrawals(opts, withdrawals, b.validatorSet)
	if err != nil {
		return errors.Wrap(err, "batched request withdrawals")
	}
	b.logger.Info().
		Str("txHash", transaction.Hash().Hex()).
		Uint64("nonce", transaction.Nonce()).
		Msg("Batched request withdrawals transaction sent to the network")
	_, err = b.WaitForTransactionReceipt(ctx, transaction)
	if err != nil {
		return errors.Wrap(err, "wait for batched request withdrawals transaction to be mined")
	}
	return nil
}

func (b *Bridge) FinalizeWithdrawals(ctx context.Context) error {
	if !b.cfg.SendFinalizeWithdrawals {
		b.logger.Info().Msg("send_finalize_withdrawals is disabled, skipping finalize withdrawals")
		return nil
	}
	pendingMsgs, err := b.bridgeHub.GetPendingMsgs(ctx)
	if err != nil {
		return err
	}
	disputePeriod, err := b.contract.GetDisputePeriod(ctx)
	if err != nil {
		return err
	}
	messageHashes := make([][32]byte, 0, len(pendingMsgs))
	for _, msg := range pendingMsgs {
		finalizeWithdrawals, err := b.contract.IsFinalizeWithdrawals(ctx, msg)
		if err != nil {
			return err
		}
		if finalizeWithdrawals {
			b.logger.Info().Hex("messageHash", msg[:]).
				Msg("message already finalized, skipping")
			continue
		}
		withdrawals, err := b.contract.GetRequestedWithdrawals(ctx, msg)
		if err != nil {
			return err
		}
		if withdrawals.RequestedTime == 0 || uint64(time.Now().Unix()) <= (withdrawals.RequestedTime+disputePeriod+1) {
			b.logger.Info().Hex("messageHash", msg[:]).
				Uint64("nonce", withdrawals.Nonce).
				Msg("message is still in dispute period, skipping")
			continue
		}
		if withdrawals.Amount.Cmp(b.cfg.MaxWithdrawAmount) > 0 {
			b.logger.Warn().Hex("messageHash", msg[:]).
				Uint64("nonce", withdrawals.Nonce).
				Str("amount", withdrawals.Amount.String()).
				Msg("withdrawal amount exceeds max allowed, skipping")
			continue
		}
		messageHashes = append(messageHashes, msg)
	}
	if len(messageHashes) == 0 {
		return nil
	}
	transaction, err := b.contract.BatchedFinalizeWithdrawals(b.NewTransactOpts(ctx), messageHashes)
	if err != nil {
		return errors.Wrap(err, "batched finalize withdrawals")
	}
	b.logger.Info().
		Str("txHash", transaction.Hash().Hex()).
		Uint64("nonce", transaction.Nonce()).
		Msg("Batched finalize withdrawals transaction sent to the network")
	_, err = b.WaitForTransactionReceipt(ctx, transaction)
	if err != nil {
		return errors.Wrap(err, "wait for batched finalize withdrawals transaction to be mined")
	}
	return nil
}

func (b *Bridge) HaveEnoughPower(cumulativePower uint64) bool {
	totalPower := uint64(0)
	for _, power := range b.validatorSet.Powers {
		totalPower += power
	}
	return 3*cumulativePower > 2*totalPower
}
