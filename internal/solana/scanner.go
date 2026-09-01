package solana

import (
	"context"
	"math/big"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// slotLookback re-fetches recent signatures so deposits are not missed when RPC
// indexing lags behind the slot cursor.
const slotLookback = 64

// SlotCache stores Solana scan progress by chain id.
type SlotCache interface {
	GetLastScannedSlot(chainID uint64) (uint64, error)
	SetLastScannedSlot(chainID, slot uint64) error
}

// SlotScannerConfig drives slot-based polling.
type SlotScannerConfig struct {
	Interval     time.Duration
	StartSlot    uint64
	SlotInterval uint64
	SlotDelay    uint64
	ClearCache   bool
}

// SlotProcessor handles Solana program signatures.
type SlotProcessor interface {
	GetChainID() *big.Int
	GetClient() *Client
	GetProgramID() solana.PublicKey
	ProcessSignature(ctx context.Context, sig solana.Signature) error
}

// Scanner polls Solana slots and processes program transactions.
type Scanner struct {
	logger    zerolog.Logger
	cfg       SlotScannerConfig
	chainID   *big.Int
	client    *Client
	programID solana.PublicKey
	cache     SlotCache
	processor SlotProcessor
}

func NewScanner(cfg SlotScannerConfig, cache SlotCache, processor SlotProcessor) *Scanner {
	logger := log.With().
		Str("module", "solana_scanner").
		Uint64("chainID", processor.GetChainID().Uint64()).
		Logger()
	return &Scanner{
		logger:    logger,
		cfg:       cfg,
		chainID:   processor.GetChainID(),
		client:    processor.GetClient(),
		programID: processor.GetProgramID(),
		cache:     cache,
		processor: processor,
	}
}

// ScanSlotRange fetches signatures for the program within the slot window.
func (s *Scanner) ScanSlotRange(ctx context.Context) error {
	if s.cfg.ClearCache {
		s.logger.Info().Msg("Clearing last scanned slot cache")
		if err := s.cache.SetLastScannedSlot(s.chainID.Uint64(), 0); err != nil {
			return errors.Wrap(err, "clear last scanned slot cache")
		}
		return errors.New("cache cleared, update config to disable clear_cache and restart scanner")
	}

	currentSlot, err := s.client.GetSlot(ctx)
	if err != nil {
		return errors.Wrap(err, "get current slot")
	}

	lastScanned, err := s.cache.GetLastScannedSlot(s.chainID.Uint64())
	if err != nil {
		return errors.Wrap(err, "get last scanned slot")
	}

	startSlot := lastScanned
	if startSlot == 0 {
		startSlot = s.cfg.StartSlot
	}
	if s.cfg.StartSlot > startSlot {
		startSlot = s.cfg.StartSlot
	}

	if currentSlot <= s.cfg.SlotDelay {
		return nil
	}
	endSlot := currentSlot - s.cfg.SlotDelay
	if endSlot <= startSlot {
		return nil
	}
	if endSlot-startSlot+1 > s.cfg.SlotInterval {
		endSlot = startSlot + s.cfg.SlotInterval - 1
	}

	s.logger.Info().
		Uint64("startSlot", startSlot).
		Uint64("endSlot", endSlot).
		Uint64("currentSlot", currentSlot).
		Msg("Scanning slot range")

	if err := s.fetchSignatures(ctx, startSlot, endSlot); err != nil {
		return err
	}

	if err := s.cache.SetLastScannedSlot(s.chainID.Uint64(), endSlot); err != nil {
		return errors.Wrap(err, "set last scanned slot")
	}
	return nil
}

func (s *Scanner) fetchSignatures(ctx context.Context, startSlot, endSlot uint64) error {
	fetchStartSlot := startSlot
	if fetchStartSlot > slotLookback {
		fetchStartSlot -= slotLookback
	}
	limit := 1000
	opts := &rpc.GetSignaturesForAddressOpts{
		Limit:          &limit,
		MinContextSlot: &fetchStartSlot,
		Commitment:     defaultCommitment,
	}
	txSigs, err := s.client.GetSignaturesForAddress(ctx, s.programID, opts)
	if err != nil {
		return errors.Wrap(err, "get transaction signatures for program")
	}
	for _, txSig := range txSigs {
		if txSig.Err != nil {
			continue
		}
		if txSig.Slot < fetchStartSlot || txSig.Slot > endSlot {
			continue
		}
		if err := s.processor.ProcessSignature(ctx, txSig.Signature); err != nil {
			s.logger.Error().Err(err).Str("signature", txSig.Signature.String()).Msg("process transaction signature failed")
			return err
		}
	}
	return nil
}
