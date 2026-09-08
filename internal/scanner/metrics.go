package scanner

import (
	"sync/atomic"
	"time"
)

// ScannerMetrics tracks EVM indexer observability counters.
type ScannerMetrics struct {
	LogsFetched      atomic.Uint64
	LogsProcessed    atomic.Uint64
	LogsFailed       atomic.Uint64
	LastScanDuration atomic.Int64 // nanoseconds
}

func (m *ScannerMetrics) RecordScanDuration(d time.Duration) {
	m.LastScanDuration.Store(int64(d))
}

func (m *ScannerMetrics) Snapshot() map[string]uint64 {
	return map[string]uint64{
		"logs_fetched":   m.LogsFetched.Load(),
		"logs_processed": m.LogsProcessed.Load(),
		"logs_failed":    m.LogsFailed.Load(),
	}
}
