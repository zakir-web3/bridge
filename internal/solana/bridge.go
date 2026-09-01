package solana

import (
	"context"
	"math/big"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/zakir-web3/bridge/internal/bridgehub"
)

type Bridge struct {
	logger  zerolog.Logger
	cfg     Config
	client  *Client
	chainID *big.Int
	hub     *bridgehub.BridgeHub
}

func NewBridge(ctx context.Context, cfg Config, hub *bridgehub.BridgeHub) (*Bridge, error) {
	client, err := NewClient(cfg.NodeURL, cfg.CommitmentType(), &cfg.RetryConfig, nil)
	if err != nil {
		return nil, err
	}
	logger := log.With().
		Str("module", "solana").
		Uint64("chainId", cfg.ChainID.Uint64()).
		Logger()
	logger.Info().
		Str("programId", cfg.ProgramID.String()).
		Msg("initialized Solana bridge scanner")
	return &Bridge{
		logger:  logger,
		cfg:     cfg,
		client:  client,
		chainID: cfg.ChainID,
		hub:     hub,
	}, nil
}

func (b *Bridge) GetChainID() *big.Int {
	return b.chainID
}

func (b *Bridge) GetClient() *Client {
	return b.client
}

func (b *Bridge) GetProgramID() solana.PublicKey {
	return b.cfg.ProgramID
}

func (b *Bridge) ProcessTransaction(ctx context.Context, sig solana.Signature, tx *rpc.GetTransactionResult) error {
	parsed, err := ParseDepositFromTransaction(b.cfg.ProgramID, tx, sig)
	if err != nil {
		b.logger.Debug().
			Err(err).
			Str("signature", sig.String()).
			Msg("skip non-deposit Solana transaction")
		return nil
	}
	mintPub := solana.PublicKeyFromBytes(parsed.Token[:])
	if !b.cfg.IsBridgeMint(mintPub) {
		b.logger.Warn().Str("mint", mintPub.String()).Msg("skipping non-whitelisted mint")
		return nil
	}
	parsed.ChainID = b.chainID
	b.logger.Info().
		Str("signature", sig.String()).
		Str("mint", mintPub.String()).
		Str("destination", parsed.Destination.Hex()).
		Str("amount", parsed.Amount.String()).
		Uint64("blockNumber", parsed.BlockNumber).
		Msg("solana deposit detected")
	return b.hub.DepositConfirmSolana(ctx, b.chainID, &parsed.SolanaDeposit)
}
