package scanner

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// EventProcessor defines the interface for processing blockchain events
type EventProcessor interface {
	// ProcessLog processes a single log event from the blockchain
	ProcessLog(ctx context.Context, log types.Log) error
	// GetFilterQuery returns an array of filter queries for scanning specific block ranges
	// This allows scanning for multiple topic combinations in a single block range
	GetFilterQuery(startBlock, endBlock uint64) []ethereum.FilterQuery
	// GetClient returns the Ethereum client used for blockchain interactions
	GetClient() *ethclient.Client
	// GetChainID returns the chain ID of the connected blockchain
	GetChainID() *big.Int
}

// Cache defines the interface for storing and retrieving blockchain scanning state
type Cache interface {
	// GetLastScannedBlock retrieves the last scanned block number for a specific chain
	GetLastScannedBlock(chainID uint64) (uint64, error)
	// SetLastScannedBlock updates the last scanned block number for a specific chain
	SetLastScannedBlock(chainID, blockNumber uint64) error
}

type Scanner struct {
	logger    zerolog.Logger
	cfg       Config
	chainID   *big.Int
	cli       *ethclient.Client
	cache     Cache
	processor EventProcessor
}

func NewScanner(cfg Config, cache Cache, processor EventProcessor) *Scanner {
	logger := log.With().
		Str("module", "scanner").
		Uint64("chainID", processor.GetChainID().Uint64()).
		Logger()
	return &Scanner{
		logger:    logger,
		cfg:       cfg,
		chainID:   processor.GetChainID(),
		cli:       processor.GetClient(),
		cache:     cache,
		processor: processor,
	}
}

// ScanBlockRange scans a range of blocks for events
func (s *Scanner) ScanBlockRange(ctx context.Context) error {
	if s.cfg.ClearCache {
		s.logger.Info().Msg("Clearing last scanned block cache")
		if err := s.cache.SetLastScannedBlock(s.chainID.Uint64(), 0); err != nil {
			return errors.Wrap(err, "clear last scanned block cache")
		}
		return errors.New("cache cleared, update config to disable clear_cache and restart scanner")
	}
	// Get current block height
	currentBlock, err := s.cli.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "get current block number")
	}

	// Get the last scanned block height
	lastScannedBlock, err := s.cache.GetLastScannedBlock(s.chainID.Uint64())
	if err != nil {
		return errors.Wrap(err, "get last scanned block")
	}

	// Calculate start block
	startBlock := lastScannedBlock
	if startBlock == 0 {
		startBlock = s.cfg.StartBlock
	}
	if s.cfg.StartBlock > startBlock {
		startBlock = s.cfg.StartBlock
	}

	// Calculate end block (considering block delay for confirmation)
	if currentBlock <= s.cfg.BlockDelay {
		s.logger.Debug().
			Uint64("startBlock", startBlock).
			Uint64("currentBlock", currentBlock).
			Uint64("blockDelay", s.cfg.BlockDelay).
			Msg("Current block is not beyond block delay, waiting for next interval")
		return nil
	}
	endBlock := currentBlock - s.cfg.BlockDelay
	if endBlock <= startBlock {
		// If end block is less than start block, no new blocks to scan
		s.logger.Debug().
			Uint64("startBlock", startBlock).
			Uint64("endBlock", endBlock).
			Uint64("currentBlock", currentBlock).
			Msg("No new blocks to scan, waiting for next interval")
		return nil
	}

	// Limit the number of blocks scanned per iteration
	if endBlock-startBlock+1 > s.cfg.BlockInterval {
		endBlock = startBlock + s.cfg.BlockInterval - 1
	}

	s.logger.Info().
		Uint64("startBlock", startBlock).
		Uint64("endBlock", endBlock).
		Uint64("currentBlock", currentBlock).
		Msg("Scanning block range")

	// Scan the blocks
	if err = s.filterBlock(ctx, startBlock, endBlock); err != nil {
		s.logger.Error().Err(err).
			Uint64("startBlock", startBlock).
			Uint64("endBlock", endBlock).
			Msg("Failed to scan block range")
		return err
	}

	// Update the last scanned block height
	if err = s.cache.SetLastScannedBlock(s.chainID.Uint64(), endBlock); err != nil {
		return errors.Wrap(err, "set last scanned block")
	}

	s.logger.Debug().
		Uint64("lastScannedBlock", endBlock).
		Msg("Block range scanned successfully")

	return nil
}

func (s *Scanner) filterBlock(ctx context.Context, startBlock, endBlock uint64) error {
	// Get filter query parameters
	filterQueries := s.processor.GetFilterQuery(startBlock, endBlock)

	// Fetch logs from the blockchain
	var logs []types.Log
	for _, filterQuery := range filterQueries {
		ls, err := s.cli.FilterLogs(ctx, filterQuery)
		if err != nil {
			return errors.Wrap(err, "filter logs")
		}
		logs = append(logs, ls...)
	}

	// Process each log event
	for _, l := range logs {
		if err := s.processor.ProcessLog(ctx, l); err != nil {
			s.logger.Error().Err(err).
				Uint64("blockNumber", l.BlockNumber).
				Msg("Failed to process log")
			return err
		}
	}
	return nil
}
