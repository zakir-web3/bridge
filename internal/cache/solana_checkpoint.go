package cache

import (
	"encoding/binary"
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/gagliardetto/solana-go"
	"github.com/pkg/errors"
)

const (
	solanaCheckpointKeyPrefix = "solana_checkpoint"
	solanaProcessedKeyPrefix  = "solana_processed"
)

// SolanaCheckpoint stores the last successfully processed slot and signature.
type SolanaCheckpoint struct {
	Slot      uint64
	Signature solana.Signature
}

func solanaCheckpointKey(chainID uint64) string {
	return fmt.Sprintf("%s_%d", solanaCheckpointKeyPrefix, chainID)
}

func solanaProcessedKey(chainID uint64, sig solana.Signature) string {
	return fmt.Sprintf("%s_%d_%s", solanaProcessedKeyPrefix, chainID, sig.String())
}

// GetSolanaCheckpoint returns the persisted scan checkpoint for a Solana chain.
// A zero signature means only the slot field is meaningful (legacy or slot-only).
func (c *BadgerCache) GetSolanaCheckpoint(chainID uint64) (SolanaCheckpoint, error) {
	var cp SolanaCheckpoint
	key := solanaCheckpointKey(chainID)

	err := c.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return nil
			}
			return errors.Wrap(err, "get solana checkpoint")
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return errors.Wrap(err, "copy checkpoint value")
		}
		if len(val) < 8 {
			return errors.New("invalid solana checkpoint value length")
		}
		cp.Slot = binary.BigEndian.Uint64(val[:8])
		if len(val) >= 8+solana.SignatureLength {
			copy(cp.Signature[:], val[8:8+solana.SignatureLength])
		}
		return nil
	})
	if err != nil {
		return SolanaCheckpoint{}, err
	}
	return cp, nil
}

// SetSolanaCheckpoint persists slot+signature after a transaction is processed.
func (c *BadgerCache) SetSolanaCheckpoint(chainID uint64, cp SolanaCheckpoint) error {
	val := make([]byte, 8+solana.SignatureLength)
	binary.BigEndian.PutUint64(val[:8], cp.Slot)
	copy(val[8:], cp.Signature[:])
	key := solanaCheckpointKey(chainID)

	return c.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), val)
	})
}

// MarkSolanaTransactionProcessed records a processed signature. Returns true when
// the signature was newly recorded and false when it already existed.
func (c *BadgerCache) MarkSolanaTransactionProcessed(chainID uint64, sig solana.Signature, slot uint64) (bool, error) {
	key := solanaProcessedKey(chainID, sig)
	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val, slot)
	newlyMarked := false

	err := c.db.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		if err == nil {
			return nil
		}
		if !errors.Is(err, badger.ErrKeyNotFound) {
			return errors.Wrap(err, "check processed transaction")
		}
		if err := txn.Set([]byte(key), val); err != nil {
			return errors.Wrap(err, "set processed transaction")
		}
		newlyMarked = true
		return nil
	})
	return newlyMarked, err
}

// IsSolanaTransactionProcessed reports whether a signature was already processed.
func (c *BadgerCache) IsSolanaTransactionProcessed(chainID uint64, sig solana.Signature) (bool, error) {
	key := solanaProcessedKey(chainID, sig)
	found := false

	err := c.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "get processed transaction")
		}
		found = true
		return nil
	})
	return found, err
}
