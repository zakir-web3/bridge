package evm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/zakir-web3/bridge/internal/contract"
)

type AccountManager struct {
	cfg        Config
	chainId    *big.Int
	cli        *ethclient.Client
	transactor *bind.TransactOpts
	logger     zerolog.Logger
}

func NewAccountManager(ctx context.Context, cfg Config, cli *ethclient.Client) (*AccountManager, error) {
	chainID, err := cli.ChainID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get chain id")
	}
	if cfg.ChainId.Cmp(chainID) != 0 {
		return nil, errors.Errorf("chain ID mismatch: config %s, client %s", cfg.ChainId.String(), chainID.String())
	}

	transactor := NewKeyedTransactor(cfg.PrivKey, chainID, cfg.MaxGasLimit)
	transactor.Context = ctx
	transactor.NoSend = cfg.NoSend

	logger := log.With().
		Uint64("chainId", chainID.Uint64()).
		Str("address", transactor.From.Hex()).
		Str("module", "acc").Logger()

	return &AccountManager{
		cfg:        cfg,
		chainId:    chainID,
		cli:        cli,
		transactor: transactor,
		logger:     logger,
	}, nil
}

func (am *AccountManager) GetChainID() *big.Int {
	return am.chainId
}

func (am *AccountManager) GetFrom() common.Address {
	return am.transactor.From
}

func (am *AccountManager) GetClient() *ethclient.Client {
	return am.cli
}

func (am *AccountManager) SignTypedData(typedData apitypes.TypedData) (contract.Signature, error) {
	hash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return contract.Signature{}, errors.Wrap(err, "get typed data hash")
	}
	sig, err := crypto.Sign(hash, am.cfg.PrivKey)
	if err != nil {
		return contract.Signature{}, err
	}
	am.logger.Debug().Hex("hash", hash).Hex("sig", sig).Msg("sign typed data")
	r := big.NewInt(0).SetBytes(sig[:32])
	s := big.NewInt(0).SetBytes(sig[32:64])
	v := sig[64] + 27 // EIP-155 compatibility
	return contract.Signature{R: r, S: s, V: v}, nil
}

func (am *AccountManager) NewTransactOpts(ctx context.Context) *bind.TransactOpts {
	transactor := *am.transactor
	transactor.Context = ctx
	transactor.GasTipCap = am.cfg.GasTipCap
	transactor.GasFeeCap = am.cfg.GasFeeCap
	return &transactor
}

func (am *AccountManager) WaitForTransactionReceipt(ctx context.Context, transaction *types.Transaction) (*types.Receipt, error) {
	if am.cfg.NoSend {
		return nil, nil
	}
	receipt, err := bind.WaitMined(ctx, am.cli, transaction.Hash())
	if err != nil {
		return nil, errors.Wrap(err, "wait for transaction to be mined")
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, errors.Errorf("transaction failed with status %d", receipt.Status)
	}
	am.logger.Info().
		Str("txHash", transaction.Hash().Hex()).
		Msg("transaction mined successfully")
	return receipt, nil
}
