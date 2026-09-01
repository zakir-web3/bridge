package solana

import (
	"context"
	"math"
	"math/big"
	"sort"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/zakir-web3/bridge/internal/cache"
)

// ScanCache stores Solana scan progress and processed transaction markers.
type ScanCache interface {
	GetLastScannedSlot(chainID uint64) (uint64, error)
	SetLastScannedSlot(chainID, slot uint64) error
	GetSolanaCheckpoint(chainID uint64) (cache.SolanaCheckpoint, error)
	SetSolanaCheckpoint(chainID uint64, cp cache.SolanaCheckpoint) error
	MarkSolanaTransactionProcessed(chainID uint64, sig solana.Signature, slot uint64) (bool, error)
	IsSolanaTransactionProcessed(chainID uint64, sig solana.Signature) (bool, error)
}

// SlotProcessor handles Solana program transactions.
type SlotProcessor interface {
	GetChainID() *big.Int
	GetClient() *Client
	GetProgramID() solana.PublicKey
	ProcessTransaction(ctx context.Context, sig solana.Signature, tx *rpc.GetTransactionResult) error
}

// Scanner polls Solana slots and processes program transactions.
type Scanner struct {
	logger    zerolog.Logger
	cfg       SlotScannerConfig
	chainID   *big.Int
	client    *Client
	programID solana.PublicKey
	cache     ScanCache
	processor SlotProcessor
	metrics   *ScannerMetrics
}

func NewScanner(cfg SlotScannerConfig, scanCache ScanCache, processor SlotProcessor) *Scanner {
	cfg.applyDefaults()
	logger := log.With().
		Str("module", "solana_scanner").
		Uint64("chainID", processor.GetChainID().Uint64()).
		Logger()
	metrics := &ScannerMetrics{}
	client := processor.GetClient()
	if client != nil {
		client.metrics = metrics
	}
	return &Scanner{
		logger:    logger,
		cfg:       cfg,
		chainID:   processor.GetChainID(),
		client:    client,
		processor: processor,
		programID: processor.GetProgramID(),
		cache:     scanCache,
		metrics:   metrics,
	}
}

// Metrics returns current scanner counters.
func (s *Scanner) Metrics() *ScannerMetrics {
	return s.metrics
}

// ScanSlotRange fetches signatures for the program within the slot window.
func (s *Scanner) ScanSlotRange(ctx context.Context) error {
	started := time.Now()
	defer func() {
		s.metrics.RecordScanDuration(time.Since(started))
	}()

	if s.cfg.ClearCache {
		s.logger.Info().Msg("clearing Solana scan cache")
		if err := s.cache.SetLastScannedSlot(s.chainID.Uint64(), 0); err != nil {
			return errors.Wrap(err, "clear last scanned slot cache")
		}
		if err := s.cache.SetSolanaCheckpoint(s.chainID.Uint64(), cache.SolanaCheckpoint{}); err != nil {
			return errors.Wrap(err, "clear solana checkpoint")
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

	fetchStartSlot := startSlot
	if fetchStartSlot > s.cfg.SlotLookback {
		fetchStartSlot -= s.cfg.SlotLookback
	}
	if s.cfg.RescanSlots > 0 && endSlot > s.cfg.RescanSlots {
		rescanStart := endSlot - s.cfg.RescanSlots
		if rescanStart < fetchStartSlot {
			fetchStartSlot = rescanStart
		}
	}

	checkpoint, err := s.cache.GetSolanaCheckpoint(s.chainID.Uint64())
	if err != nil {
		return errors.Wrap(err, "get solana checkpoint")
	}

	s.logger.Info().
		Uint64("startSlot", startSlot).
		Uint64("endSlot", endSlot).
		Uint64("fetchStartSlot", fetchStartSlot).
		Uint64("currentSlot", currentSlot).
		Uint64("checkpointSlot", checkpoint.Slot).
		Str("checkpointSignature", checkpoint.Signature.String()).
		Uint64("backlogSlots", endSlot-startSlot).
		Msg("scanning slot range")

	signatures, err := s.paginateSignatures(ctx, fetchStartSlot, endSlot)
	if err != nil {
		return err
	}
	s.metrics.SignaturesFetched.Add(uint64(len(signatures)))

	if len(signatures) == 0 {
		s.logger.Debug().Msg("no signatures in slot window")
	} else if err := s.processSignatures(ctx, signatures, checkpoint); err != nil {
		return err
	}

	if err := s.cache.SetLastScannedSlot(s.chainID.Uint64(), endSlot); err != nil {
		return errors.Wrap(err, "set last scanned slot")
	}

	s.logger.Info().
		Uint64("endSlot", endSlot).
		Interface("metrics", s.metrics.Snapshot()).
		Msg("slot range scan completed")
	return nil
}

func (s *Scanner) paginateSignatures(ctx context.Context, minSlot, maxSlot uint64) ([]*rpc.TransactionSignature, error) {
	limit := s.cfg.SignaturesPageLimit
	var (
		all    []*rpc.TransactionSignature
		before *solana.Signature
	)

	for {
		opts := &rpc.GetSignaturesForAddressOpts{
			Limit: &limit,
		}
		if before != nil {
			opts.Before = *before
		}
		if minSlot > 0 {
			opts.MinContextSlot = &minSlot
		}

		page, err := s.client.GetSignaturesForAddress(ctx, s.programID, opts)
		if err != nil {
			return nil, errors.Wrap(err, "get transaction signatures for program")
		}
		s.metrics.PagesFetched.Add(1)

		if len(page) == 0 {
			break
		}

		oldestSlot := uint64(math.MaxUint64)
		for _, txSig := range page {
			if txSig.Slot < oldestSlot {
				oldestSlot = txSig.Slot
			}
			if txSig.Err != nil {
				s.metrics.TransactionsSkipped.Add(1)
				continue
			}
			if txSig.Slot < minSlot || txSig.Slot > maxSlot {
				continue
			}
			all = append(all, txSig)
		}

		last := page[len(page)-1].Signature
		before = &last

		if oldestSlot < minSlot {
			break
		}
		if len(page) < limit {
			break
		}
	}

	sort.SliceStable(all, func(i, j int) bool {
		if all[i].Slot != all[j].Slot {
			return all[i].Slot < all[j].Slot
		}
		return all[i].Signature.String() < all[j].Signature.String()
	})
	return all, nil
}

func (s *Scanner) processSignatures(ctx context.Context, signatures []*rpc.TransactionSignature, checkpoint cache.SolanaCheckpoint) error {
	for _, txSig := range signatures {
		if shouldSkipSignature(txSig, checkpoint) {
			s.metrics.TransactionsSkipped.Add(1)
			continue
		}
		processed, err := s.cache.IsSolanaTransactionProcessed(s.chainID.Uint64(), txSig.Signature)
		if err != nil {
			return errors.Wrap(err, "check processed transaction")
		}
		if processed {
			s.metrics.TransactionsSkipped.Add(1)
			continue
		}

		tx, err := s.client.GetTransactionWithRetry(ctx, txSig.Signature, s.cfg.TxFetchRetries)
		if err != nil {
			s.metrics.TransactionsFailed.Add(1)
			return errors.Wrapf(err, "get transaction %s", txSig.Signature.String())
		}
		s.metrics.TransactionsFetched.Add(1)

		if err := s.processor.ProcessTransaction(ctx, txSig.Signature, tx); err != nil {
			s.metrics.TransactionsFailed.Add(1)
			s.logger.Error().
				Err(err).
				Str("signature", txSig.Signature.String()).
				Uint64("slot", txSig.Slot).
				Msg("process transaction failed")
			return err
		}

		if _, err := s.cache.MarkSolanaTransactionProcessed(s.chainID.Uint64(), txSig.Signature, txSig.Slot); err != nil {
			return errors.Wrap(err, "mark transaction processed")
		}
		if err := s.cache.SetSolanaCheckpoint(s.chainID.Uint64(), cache.SolanaCheckpoint{
			Slot:      txSig.Slot,
			Signature: txSig.Signature,
		}); err != nil {
			return errors.Wrap(err, "set solana checkpoint")
		}
		s.metrics.TransactionsOK.Add(1)
	}
	return nil
}

func shouldSkipSignature(txSig *rpc.TransactionSignature, checkpoint cache.SolanaCheckpoint) bool {
	if checkpoint.Signature == (solana.Signature{}) {
		return false
	}
	if txSig.Slot < checkpoint.Slot {
		return true
	}
	if txSig.Slot > checkpoint.Slot {
		return false
	}
	return txSig.Signature.String() <= checkpoint.Signature.String()
}
