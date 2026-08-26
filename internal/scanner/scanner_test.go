package scanner

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
)

type testProcessor struct {
	client  *ethclient.Client
	chainID *big.Int
}

func (p *testProcessor) ProcessLog(_ context.Context, _ types.Log) error   { return nil }
func (p *testProcessor) GetFilterQuery(_, _ uint64) []ethereum.FilterQuery { return nil }
func (p *testProcessor) GetClient() *ethclient.Client                      { return p.client }
func (p *testProcessor) GetChainID() *big.Int                              { return p.chainID }

type testCache struct {
	lastBlock uint64
	setCalled bool
}

func (c *testCache) GetLastScannedBlock(_ uint64) (uint64, error) { return c.lastBlock, nil }
func (c *testCache) SetLastScannedBlock(_, blockNumber uint64) error {
	c.setCalled = true
	c.lastBlock = blockNumber
	return nil
}

func TestScanBlockRange_NoScanWhenCurrentBlockNotBeyondDelay(t *testing.T) {
	t.Parallel()

	// Return block number 10 for eth_blockNumber.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xa"}`))
	}))
	defer srv.Close()

	rpcClient, err := rpc.DialOptions(context.Background(), srv.URL)
	require.NoError(t, err)
	defer rpcClient.Close()

	processor := &testProcessor{
		client:  ethclient.NewClient(rpcClient),
		chainID: big.NewInt(1),
	}
	cache := &testCache{lastBlock: 5}
	s := NewScanner(Config{
		StartBlock:    1,
		BlockInterval: 100,
		BlockDelay:    10,
	}, cache, processor)

	err = s.ScanBlockRange(context.Background())
	require.NoError(t, err)
	require.False(t, cache.setCalled, "scanner should not update cache when no blocks are available")
}

func TestScanBlockRange_UpdateCacheAfterSuccessfulScan(t *testing.T) {
	t.Parallel()

	call := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		call++
		if call == 1 {
			_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x64"}`)) // 100
			return
		}
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":[]}`)) // eth_getLogs
	}))
	defer srv.Close()

	rpcClient, err := rpc.DialOptions(context.Background(), srv.URL)
	require.NoError(t, err)
	defer rpcClient.Close()

	processor := &testProcessor{
		client:  ethclient.NewClient(rpcClient),
		chainID: big.NewInt(1),
	}
	cache := &testCache{lastBlock: 90}
	s := NewScanner(Config{
		StartBlock:    1,
		BlockInterval: 10,
		BlockDelay:    1,
	}, cache, processor)

	// Provide at least one query to force FilterLogs path.
	processorWithQuery := &testProcessorWithQuery{testProcessor: processor}
	s.processor = processorWithQuery

	err = s.ScanBlockRange(context.Background())
	require.NoError(t, err)
	require.True(t, cache.setCalled)
	require.Equal(t, uint64(99), cache.lastBlock)
}

type testProcessorWithQuery struct {
	*testProcessor
}

func (p *testProcessorWithQuery) GetFilterQuery(startBlock, endBlock uint64) []ethereum.FilterQuery {
	return []ethereum.FilterQuery{{
		FromBlock: new(big.Int).SetUint64(startBlock),
		ToBlock:   new(big.Int).SetUint64(endBlock),
	}}
}
