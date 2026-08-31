package bridgehub

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/zakir-web3/bridge/internal/contract"
	"github.com/zakir-web3/bridge/internal/evm"
)

type BridgeContract interface {
	Withdraw(ctx context.Context, chainId *big.Int, withdraw ...*contract.BridgeHubWithdraw) error
	BridgeSignatureSubmitted(ctx context.Context, chainId *big.Int, sings ...*contract.MessageSignature) error
}

type BridgeHub struct {
	*evm.AccountManager
	logger          zerolog.Logger
	cfg             Config
	contract        *contract.BridgeHub
	validatorSet    *contract.ValidatorSet
	bc              BridgeContract
	domainSeparator common.Hash
}

func NewBridgeHub(ctx context.Context, cfg Config) (*BridgeHub, error) {
	cli, err := evm.NewClient(ctx, cfg.NodeURL, &cfg.RetryConfig)
	if err != nil {
		return nil, err
	}

	c, err := contract.NewBridgeHub(cfg.BridgeHubAddress, cli)
	if err != nil {
		return nil, errors.Wrap(err, "new BridgeHub contract")
	}

	acc, err := evm.NewAccountManager(ctx, cfg.AccConfig, cli)
	if err != nil {
		return nil, err
	}
	logger := log.With().
		Str("module", "bridge_hub").
		Uint64("chainId", acc.GetChainID().Uint64()).
		Logger()
	logger.Info().Str("bridgeHub", cfg.BridgeHubAddress.Hex()).Str("validator", acc.GetFrom().Hex()).Msg("initialized BridgeHub")

	domainSeparator, err := c.GetDomainSeparator(ctx)
	if err != nil {
		return nil, err
	}
	logger.Info().Hex("domainSeparator", domainSeparator.Bytes()).Msg("get domain separator from contract")

	validatorSet, err := c.GetValidatorSet(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get validator set")
	}
	logger.Info().Strs("validators", validatorSet.ValidatorsStr()).
		Uint64("epoch", validatorSet.Epoch).
		Uints64("power", validatorSet.Powers).
		Msg("get validator set from contract")
	if !validatorSet.IsValidator(acc.GetFrom()) {
		return nil, errors.New("not a validator")
	}

	return &BridgeHub{
		AccountManager:  acc,
		logger:          logger,
		cfg:             cfg,
		contract:        c,
		validatorSet:    validatorSet,
		domainSeparator: domainSeparator,
	}, nil
}

func (b *BridgeHub) GetValidatorSet() *contract.ValidatorSet {
	return b.validatorSet
}

func (b *BridgeHub) SetBridgeContract(bc BridgeContract) {
	b.bc = bc
}

func (b *BridgeHub) DepositConfirmSolana(ctx context.Context, srcChainID *big.Int, deposits ...*contract.SolanaDeposit) error {
	depositConfirms := make([]contract.DepositConfirm, 0, len(deposits))
	for _, d := range deposits {
		typedData := d.ToTypedData(srcChainID, b.GetChainID(), b.cfg.BridgeHubAddress)
		processed, err := b.hasSignature(ctx, typedData)
		if err != nil {
			return err
		}
		if processed {
			continue
		}
		signature, err := b.SignTypedData(typedData)
		if err != nil {
			return err
		}
		depositConfirms = append(depositConfirms, d.ToDepositConfirm(signature))
	}
	if len(depositConfirms) == 0 {
		return nil
	}
	transaction, err := b.contract.DepositConfirm(b.NewTransactOpts(ctx), depositConfirms)
	if err != nil {
		return errors.Wrap(err, "create deposit confirm transaction")
	}
	b.logger.Info().
		Str("txHash", transaction.Hash().Hex()).
		Uint64("nonce", transaction.Nonce()).
		Msg("Solana deposit confirm transaction sent")
	_, err = b.WaitForTransactionReceipt(ctx, transaction)
	return err
}

func (b *BridgeHub) DepositConfirm(ctx context.Context, chainId *big.Int, transfer ...*contract.BridgeERC20Transfer) error {
	depositConfirms := make([]contract.DepositConfirm, 0, len(transfer))
	for _, t := range transfer {
		typedData := t.ToTypedData(chainId, b.GetChainID(), b.cfg.BridgeHubAddress)
		processed, err := b.hasSignature(ctx, typedData)
		if err != nil {
			return err
		}
		if processed {
			continue
		}
		signature, err := b.SignTypedData(typedData)
		if err != nil {
			return err
		}
		depositConfirm := t.ToDepositConfirm(chainId, signature)
		depositConfirms = append(depositConfirms, depositConfirm)
	}
	if len(depositConfirms) == 0 {
		return nil
	}
	transaction, err := b.contract.DepositConfirm(b.NewTransactOpts(ctx), depositConfirms)
	if err != nil {
		return errors.Wrap(err, "create deposit confirm transaction")
	}
	b.logger.Info().
		Str("txHash", transaction.Hash().Hex()).
		Uint64("nonce", transaction.Nonce()).
		Msg("Deposit transaction sent to the network")
	_, err = b.WaitForTransactionReceipt(ctx, transaction)
	return err
}

func (b *BridgeHub) SubmitSignatures(ctx context.Context, _ *big.Int, signatures ...contract.CrossChainMessage) error {
	validSignatures := make([]contract.CrossChainMessage, 0, len(signatures))
	for i := 0; i < len(signatures); i++ {
		signature := signatures[i]
		messageHash := common.BytesToHash(crypto.Keccak256(signature.MessageRawData))
		address, err := contract.RecoverSignerAddress(signature.DomainSeparator, messageHash, signature.Signature)
		if err != nil {
			return err
		}
		if !b.validatorSet.IsValidator(address) {
			return errors.Errorf("signer %s is not a validator", address.Hex())
		}
		messageSignature, err := b.contract.GetBridgeMsgSign(ctx, messageHash, address)
		if err != nil {
			return err
		}
		if !messageSignature.IsEmpty() {
			b.logger.Warn().
				Str("messageHash", messageHash.Hex()).
				Str("signer", address.Hex()).
				Msg("Signature already submitted, skipping")
			continue
		}
		validSignatures = append(validSignatures, signature)
	}
	if len(validSignatures) == 0 {
		return nil
	}
	transaction, err := b.contract.SubmitBridgeSignatures(b.NewTransactOpts(ctx), validSignatures)
	if err != nil {
		return errors.Wrap(err, "create submit signatures transaction")
	}
	b.logger.Info().
		Str("txHash", transaction.Hash().Hex()).
		Uint64("nonce", transaction.Nonce()).
		Msg("Submit signatures transaction sent to the network")
	_, err = b.WaitForTransactionReceipt(ctx, transaction)
	return err
}

func (b *BridgeHub) FinalizedWithdrawal(ctx context.Context, chainId *big.Int, withdrawals ...*contract.BridgeFinalizedWithdrawal) error {
	withdrawConfirms := make([]contract.WithdrawConfirm, 0, len(withdrawals))
	for _, w := range withdrawals {
		typedData := w.ToTypedData(chainId, b.GetChainID(), b.cfg.BridgeHubAddress)
		processed, err := b.hasSignature(ctx, typedData)
		if err != nil {
			return err
		}
		if processed {
			continue
		}
		signature, err := b.SignTypedData(typedData)
		if err != nil {
			return err
		}
		withdrawConfirms = append(withdrawConfirms, w.ToWithdrawConfirm(chainId, signature))
	}
	if len(withdrawConfirms) == 0 {
		return nil
	}
	transaction, err := b.contract.WithdrawConfirm(b.NewTransactOpts(ctx), withdrawConfirms)
	if err != nil {
		return errors.Wrap(err, "create withdraw confirm transaction")
	}
	b.logger.Info().
		Str("txHash", transaction.Hash().Hex()).
		Uint64("nonce", transaction.Nonce()).
		Msg("Withdraw transaction sent to the network")
	_, err = b.WaitForTransactionReceipt(ctx, transaction)
	return err
}

func (b *BridgeHub) hasSignature(ctx context.Context, typedData apitypes.TypedData) (bool, error) {
	messageHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return false, errors.Wrap(err, "hash struct for deposit confirm")
	}
	processed, err := b.contract.IsMsgProcessed(ctx, common.BytesToHash(messageHash))
	if err != nil {
		return false, err
	}
	if processed {
		b.logger.Info().
			Hex("messageHash", messageHash).
			Str("primaryType", typedData.PrimaryType).
			Uint64("chainId", b.GetChainID().Uint64()).
			Msg("Message already processed, skipping")
		return true, nil
	}
	sig, err := b.contract.GetMsgSign(ctx, common.BytesToHash(messageHash), b.GetFrom())
	if err != nil {
		return false, err
	}
	logger := b.logger.Info().
		Hex("messageHash", messageHash).
		Str("primaryType", typedData.PrimaryType).
		Uint64("chainId", b.GetChainID().Uint64())
	if !sig.IsEmpty() {
		logger.Msg("Signature already submitted, skipping")
		return true, nil
	}
	logger.Msg("Signature not found, will submit")
	return false, nil
}

func (b *BridgeHub) HaveEnoughPower(cumulativePower uint64) bool {
	totalPower := uint64(0)
	for _, power := range b.validatorSet.Powers {
		totalPower += power
	}
	return 3*cumulativePower > 2*totalPower
	// return cumulativePower == totalPower
}

func (b *BridgeHub) GetPendingMsgs(ctx context.Context) ([][32]byte, error) {
	pendingMessages, err := b.contract.GetPendingMessages(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, errors.Wrap(err, "get pending messages")
	}
	return pendingMessages, nil
}
