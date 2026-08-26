package cache

import (
	"os"
	"testing"
)

func TestNewBadgerCache(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "badger_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test creating cache instance
	cache, err := NewBadgerCache(tempDir)
	if err != nil {
		t.Fatalf("Failed to create BadgerCache: %v", err)
	}
	defer cache.db.Close()

	if cache.db == nil {
		t.Error("Expected db to be initialized, got nil")
	}
}

func TestGetLastScannedBlock_Empty(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "badger_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cache, err := NewBadgerCache(tempDir)
	if err != nil {
		t.Fatalf("Failed to create BadgerCache: %v", err)
	}
	defer cache.db.Close()

	// Test getting last scanned block number from empty cache
	chainID := uint64(1)
	blockNumber, err := cache.GetLastScannedBlock(chainID)
	if err != nil {
		t.Fatalf("Failed to get last scanned block: %v", err)
	}

	if blockNumber != 0 {
		t.Errorf("Expected block number to be 0 for empty cache, got %d", blockNumber)
	}
}

func TestSetAndGetLastScannedBlock(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "badger_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cache, err := NewBadgerCache(tempDir)
	if err != nil {
		t.Fatalf("Failed to create BadgerCache: %v", err)
	}
	defer cache.db.Close()

	chainID := uint64(1)
	expectedBlockNumber := uint64(12345)

	// Test setting last scanned block number
	err = cache.SetLastScannedBlock(chainID, expectedBlockNumber)
	if err != nil {
		t.Fatalf("Failed to set last scanned block: %v", err)
	}

	// Test getting last scanned block number
	blockNumber, err := cache.GetLastScannedBlock(chainID)
	if err != nil {
		t.Fatalf("Failed to get last scanned block: %v", err)
	}

	if blockNumber != expectedBlockNumber {
		t.Errorf("Expected block number %d, got %d", expectedBlockNumber, blockNumber)
	}
}

func TestMultipleChains(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "badger_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cache, err := NewBadgerCache(tempDir)
	if err != nil {
		t.Fatalf("Failed to create BadgerCache: %v", err)
	}
	defer cache.db.Close()

	// Test setting and getting block numbers for multiple chains
	testCases := []struct {
		chainID     uint64
		blockNumber uint64
	}{
		{1, 1000},
		{2, 2000},
		{3, 3000},
	}

	// Set block numbers for multiple chains
	for _, tc := range testCases {
		err := cache.SetLastScannedBlock(tc.chainID, tc.blockNumber)
		if err != nil {
			t.Fatalf("Failed to set block number for chain %d: %v", tc.chainID, err)
		}
	}

	// Verify block numbers for each chain
	for _, tc := range testCases {
		blockNumber, err := cache.GetLastScannedBlock(tc.chainID)
		if err != nil {
			t.Fatalf("Failed to get block number for chain %d: %v", tc.chainID, err)
		}

		if blockNumber != tc.blockNumber {
			t.Errorf(
				"Chain %d: expected block number %d, got %d",
				tc.chainID,
				tc.blockNumber,
				blockNumber,
			)
		}
	}
}

func TestUpdateExistingBlock(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "badger_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cache, err := NewBadgerCache(tempDir)
	if err != nil {
		t.Fatalf("Failed to create BadgerCache: %v", err)
	}
	defer cache.db.Close()

	chainID := uint64(1)
	initialBlockNumber := uint64(1000)
	updatedBlockNumber := uint64(2000)

	// Set initial block number
	err = cache.SetLastScannedBlock(chainID, initialBlockNumber)
	if err != nil {
		t.Fatalf("Failed to set initial block number: %v", err)
	}

	// Verify initial value
	blockNumber, err := cache.GetLastScannedBlock(chainID)
	if err != nil {
		t.Fatalf("Failed to get initial block number: %v", err)
	}
	if blockNumber != initialBlockNumber {
		t.Errorf("Expected initial block number %d, got %d", initialBlockNumber, blockNumber)
	}

	// Update block number
	err = cache.SetLastScannedBlock(chainID, updatedBlockNumber)
	if err != nil {
		t.Fatalf("Failed to update block number: %v", err)
	}

	// Verify updated value
	blockNumber, err = cache.GetLastScannedBlock(chainID)
	if err != nil {
		t.Fatalf("Failed to get updated block number: %v", err)
	}
	if blockNumber != updatedBlockNumber {
		t.Errorf("Expected updated block number %d, got %d", updatedBlockNumber, blockNumber)
	}
}
