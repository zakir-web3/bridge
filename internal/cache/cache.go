package cache

import (
	"encoding/binary"
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// Default cache key prefix for scanned blocks
	scannedBlockKeyPrefix = "scanned_block"
	scannedSlotKeyPrefix  = "scanned_slot"
)

// BadgerCache implements cache interface using Badger database
type BadgerCache struct {
	db *badger.DB
}

// NewBadgerCache creates a new Badger cache instance
func NewBadgerCache(source string) (*BadgerCache, error) {
	options := badger.DefaultOptions(source)
	options.Logger = &logger{log: log.Logger.Level(zerolog.InfoLevel)}
	db, err := badger.Open(options)
	if err != nil {
		return nil, err
	}
	return &BadgerCache{db: db}, nil
}

// getChainKey generates a chain-specific cache key
func getChainKey(chainID uint64) string {
	return fmt.Sprintf("%s_%d", scannedBlockKeyPrefix, chainID)
}

func getSlotKey(chainID uint64) string {
	return fmt.Sprintf("%s_%d", scannedSlotKeyPrefix, chainID)
}

// GetLastScannedSlot retrieves the last scanned slot for a Solana chain id.
func (c *BadgerCache) GetLastScannedSlot(chainID uint64) (uint64, error) {
	var slot uint64
	key := getSlotKey(chainID)

	err := c.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				slot = 0
				return nil
			}
			return errors.Wrap(err, "get last scanned slot")
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return errors.Wrap(err, "copy value")
		}
		if len(val) != 8 {
			return errors.New("invalid slot value length")
		}
		slot = binary.BigEndian.Uint64(val)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return slot, nil
}

// SetLastScannedSlot updates the last scanned slot for a Solana chain id.
func (c *BadgerCache) SetLastScannedSlot(chainID, slot uint64) error {
	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val, slot)
	key := getSlotKey(chainID)

	err := c.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), val)
	})
	return err
}

// GetLastScannedBlock retrieves the last scanned block number for a specific chain
func (c *BadgerCache) GetLastScannedBlock(chainID uint64) (uint64, error) {
	var blockNumber uint64
	key := getChainKey(chainID)

	err := c.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				// If key not found, return 0 (indicating never scanned)
				blockNumber = 0
				return nil
			}
			return errors.Wrap(err, "get last scanned block")
		}

		val, err := item.ValueCopy(nil)
		if err != nil {
			return errors.Wrap(err, "copy value")
		}

		if len(val) != 8 {
			return errors.New("invalid block number value length")
		}

		blockNumber = binary.BigEndian.Uint64(val)
		return nil
	})
	if err != nil {
		return 0, err
	}

	return blockNumber, nil
}

// SetLastScannedBlock updates the last scanned block number for a specific chain
func (c *BadgerCache) SetLastScannedBlock(chainID, blockNumber uint64) error {
	// Convert uint64 to 8-byte big-endian format
	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val, blockNumber)
	key := getChainKey(chainID)

	err := c.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), val)
		if err != nil {
			return errors.Wrap(err, "set last scanned block")
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
