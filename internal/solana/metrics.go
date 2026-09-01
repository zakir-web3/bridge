package solana

import (
	"sync/atomic"
	"time"
)

// ScannerMetrics tracks indexer observability counters.
type ScannerMetrics struct {
	SignaturesFetched   atomic.Uint64
	TransactionsFetched atomic.Uint64
	TransactionsSkipped atomic.Uint64
	TransactionsFailed  atomic.Uint64
	TransactionsOK      atomic.Uint64
	RPCRetries          atomic.Uint64
	PagesFetched        atomic.Uint64
	LastScanDuration    atomic.Int64 // nanoseconds
}

func (m *ScannerMetrics) RecordScanDuration(d time.Duration) {
	m.LastScanDuration.Store(int64(d))
}

func (m *ScannerMetrics) Snapshot() map[string]uint64 {
	return map[string]uint64{
		"signatures_fetched":   m.SignaturesFetched.Load(),
		"transactions_fetched": m.TransactionsFetched.Load(),
		"transactions_skipped": m.TransactionsSkipped.Load(),
		"transactions_failed":  m.TransactionsFailed.Load(),
		"transactions_ok":      m.TransactionsOK.Load(),
		"rpc_retries":          m.RPCRetries.Load(),
		"pages_fetched":        m.PagesFetched.Load(),
	}
}
