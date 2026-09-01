package solana

import (
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/zakir-web3/bridge/internal/cache"
	"github.com/zakir-web3/bridge/internal/evm"
)

func TestShouldSkipSignature(t *testing.T) {
	var sigA, sigB solana.Signature
	sigA[0] = 1
	sigB[0] = 2

	checkpoint := cache.SolanaCheckpoint{Slot: 100, Signature: sigA}

	require.True(t, shouldSkipSignature(&rpc.TransactionSignature{Slot: 99, Signature: sigB}, checkpoint))
	require.True(t, shouldSkipSignature(&rpc.TransactionSignature{Slot: 100, Signature: sigA}, checkpoint))
	require.False(t, shouldSkipSignature(&rpc.TransactionSignature{Slot: 100, Signature: sigB}, checkpoint))
	require.False(t, shouldSkipSignature(&rpc.TransactionSignature{Slot: 101, Signature: sigA}, checkpoint))
}

func TestSlotScannerConfigApplyDefaults(t *testing.T) {
	cfg := SlotScannerConfig{Interval: 1, SlotInterval: 1}
	cfg.applyDefaults()
	require.Equal(t, uint64(defaultSlotLookback), cfg.SlotLookback)
	require.Equal(t, uint64(defaultRescanSlots), cfg.RescanSlots)
	require.Equal(t, defaultSignaturesPageLimit, cfg.SignaturesPageLimit)
	require.Equal(t, defaultTxFetchRetries, cfg.TxFetchRetries)
}

type testScanCache struct {
	lastScanned uint64
	checkpoint  cache.SolanaCheckpoint
	processed   map[solana.Signature]bool
	setCalled   bool
}

func newTestScanCache(lastScanned uint64) *testScanCache {
	return &testScanCache{
		lastScanned: lastScanned,
		processed:   make(map[solana.Signature]bool),
	}
}

func (c *testScanCache) GetLastScannedSlot(_ uint64) (uint64, error) {
	return c.lastScanned, nil
}

func (c *testScanCache) SetLastScannedSlot(_ uint64, slot uint64) error {
	c.setCalled = true
	c.lastScanned = slot
	return nil
}

func (c *testScanCache) GetSolanaCheckpoint(_ uint64) (cache.SolanaCheckpoint, error) {
	return c.checkpoint, nil
}

func (c *testScanCache) SetSolanaCheckpoint(_ uint64, cp cache.SolanaCheckpoint) error {
	c.checkpoint = cp
	return nil
}

func (c *testScanCache) MarkSolanaTransactionProcessed(_ uint64, sig solana.Signature, _ uint64) (bool, error) {
	if c.processed[sig] {
		return false, nil
	}
	c.processed[sig] = true
	return true, nil
}

func (c *testScanCache) IsSolanaTransactionProcessed(_ uint64, sig solana.Signature) (bool, error) {
	return c.processed[sig], nil
}

type testSlotProcessor struct {
	client    *Client
	chainID   *big.Int
	programID solana.PublicKey
	processed []solana.Signature
}

func (p *testSlotProcessor) GetChainID() *big.Int {
	return p.chainID
}

func (p *testSlotProcessor) GetClient() *Client {
	return p.client
}

func (p *testSlotProcessor) GetProgramID() solana.PublicKey {
	return p.programID
}

func (p *testSlotProcessor) ProcessTransaction(_ context.Context, sig solana.Signature, _ *rpc.GetTransactionResult) error {
	p.processed = append(p.processed, sig)
	return nil
}

type mockSolanaRPC struct {
	slot        uint64
	signatures  []*rpc.TransactionSignature
	transaction *rpc.GetTransactionResult
}

func (m *mockSolanaRPC) serveHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req struct {
		Method string          `json:"method"`
		ID     json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result any
	switch req.Method {
	case "getSlot":
		result = m.slot
	case "getSignaturesForAddress":
		result = m.signatures
	case "getTransaction":
		result = m.transaction
	default:
		http.Error(w, "unexpected method: "+req.Method, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"jsonrpc": "2.0",
		"id":      req.ID,
		"result":  result,
	})
}

func testScannerConfig() SlotScannerConfig {
	return SlotScannerConfig{
		Interval:            1,
		StartSlot:           1,
		SlotInterval:        100,
		SlotDelay:           10,
		SlotLookback:        0,
		RescanSlots:         0,
		SignaturesPageLimit: 10,
		TxFetchRetries:      0,
	}
}

func newTestScanner(t *testing.T, cfg SlotScannerConfig, scanCache ScanCache, mockRPC *mockSolanaRPC) (*Scanner, *testSlotProcessor) {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(mockRPC.serveHTTP))
	t.Cleanup(srv.Close)

	metrics := &ScannerMetrics{}
	client, err := NewClient(srv.URL, rpc.CommitmentFinalized, evm.DefaultRetryConfig(), metrics)
	require.NoError(t, err)

	programID := solana.MustPublicKeyFromBase58("11111111111111111111111111111111")
	processor := &testSlotProcessor{
		client:    client,
		chainID:   big.NewInt(901),
		programID: programID,
	}
	return NewScanner(cfg, scanCache, processor), processor
}

func TestScanSlotRange_NoScanWhenBehindDelay(t *testing.T) {
	t.Parallel()

	cache := newTestScanCache(5)
	scanner, _ := newTestScanner(t, testScannerConfig(), cache, &mockSolanaRPC{slot: 10})

	err := scanner.ScanSlotRange(context.Background())
	require.NoError(t, err)
	require.False(t, cache.setCalled)
}

func TestScanSlotRange_NoScanWhenCaughtUp(t *testing.T) {
	t.Parallel()

	cache := newTestScanCache(90)
	scanner, _ := newTestScanner(t, testScannerConfig(), cache, &mockSolanaRPC{slot: 100})

	err := scanner.ScanSlotRange(context.Background())
	require.NoError(t, err)
	require.False(t, cache.setCalled)
}

func TestScanSlotRange_ClearCache(t *testing.T) {
	t.Parallel()

	scanCache := newTestScanCache(500)
	scanCache.checkpoint = cache.SolanaCheckpoint{Slot: 400}
	cfg := testScannerConfig()
	cfg.ClearCache = true

	scanner, _ := newTestScanner(t, cfg, scanCache, &mockSolanaRPC{slot: 1000})

	err := scanner.ScanSlotRange(context.Background())
	require.Error(t, err)
	require.Contains(t, err.Error(), "cache cleared")
	require.Equal(t, uint64(0), scanCache.lastScanned)
	require.Equal(t, cache.SolanaCheckpoint{}, scanCache.checkpoint)
}

func TestScanSlotRange_UpdatesCacheWithNoSignatures(t *testing.T) {
	t.Parallel()

	cache := newTestScanCache(80)
	scanner, _ := newTestScanner(t, testScannerConfig(), cache, &mockSolanaRPC{
		slot:       100,
		signatures: nil,
	})

	err := scanner.ScanSlotRange(context.Background())
	require.NoError(t, err)
	require.True(t, cache.setCalled)
	require.Equal(t, uint64(90), cache.lastScanned)
}

func TestScanSlotRange_ProcessesSignatures(t *testing.T) {
	t.Parallel()

	var sig solana.Signature
	sig[0] = 42

	cache := newTestScanCache(80)
	scanner, processor := newTestScanner(t, testScannerConfig(), cache, &mockSolanaRPC{
		slot: 100,
		signatures: []*rpc.TransactionSignature{
			{Signature: sig, Slot: 85},
		},
		transaction: &rpc.GetTransactionResult{},
	})

	err := scanner.ScanSlotRange(context.Background())
	require.NoError(t, err)
	require.True(t, cache.setCalled)
	require.Equal(t, uint64(90), cache.lastScanned)
	require.Equal(t, []solana.Signature{sig}, processor.processed)
	require.True(t, cache.processed[sig])
	require.Equal(t, uint64(85), cache.checkpoint.Slot)
	require.Equal(t, sig, cache.checkpoint.Signature)
	require.Equal(t, uint64(1), scanner.Metrics().SignaturesFetched.Load())
	require.Equal(t, uint64(1), scanner.Metrics().TransactionsOK.Load())
}
