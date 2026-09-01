package solana

import (
	"context"
	"math"
	"net/http"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/jsonrpc"
	"github.com/pkg/errors"

	"github.com/zakir-web3/bridge/internal/evm"
)

const defaultRPCTimeout = 30 * time.Second

// Client wraps Solana RPC with retry.
type Client struct {
	rpc        *rpc.Client
	retry      *evm.RetryConfig
	commitment rpc.CommitmentType
	timeout    time.Duration
	metrics    *ScannerMetrics
}

func NewClient(nodeURL string, commitment rpc.CommitmentType, config *evm.RetryConfig, metrics *ScannerMetrics) (*Client, error) {
	if nodeURL == "" {
		return nil, errors.New("node_url is required")
	}
	if commitment == "" {
		commitment = rpc.CommitmentFinalized
	}
	if config == nil {
		config = evm.DefaultRetryConfig()
	}

	httpClient := &http.Client{
		Transport: evm.NewRetryTransport(http.DefaultTransport, config),
		Timeout:   defaultRPCTimeout,
	}
	opts := &jsonrpc.RPCClientOpts{HTTPClient: httpClient}
	rpcClient := jsonrpc.NewClientWithOpts(nodeURL, opts)

	return &Client{
		rpc:        rpc.NewWithCustomRPCClient(rpcClient),
		retry:      config,
		commitment: commitment,
		timeout:    defaultRPCTimeout,
		metrics:    metrics,
	}, nil
}

func (c *Client) withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if _, ok := ctx.Deadline(); ok {
		return ctx, func() {}
	}
	return context.WithTimeout(ctx, c.timeout)
}

func (c *Client) GetSlot(ctx context.Context) (uint64, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()

	slot, err := c.rpc.GetSlot(ctx, c.commitment)
	if err != nil {
		return 0, errors.Wrap(err, "get slot")
	}
	return slot, nil
}

func (c *Client) GetTransaction(ctx context.Context, sig solana.Signature) (*rpc.GetTransactionResult, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()

	version := uint64(0)
	tx, err := c.rpc.GetTransaction(ctx, sig, &rpc.GetTransactionOpts{
		Encoding:                       solana.EncodingBase64,
		Commitment:                     c.commitment,
		MaxSupportedTransactionVersion: &version,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get transaction")
	}
	return tx, nil
}

// GetTransactionWithRetry retries only when RPC succeeds but returns a null transaction
// (not yet indexed). Transport-level failures are already handled by RetryTransport.
func (c *Client) GetTransactionWithRetry(ctx context.Context, sig solana.Signature, maxRetries int) (*rpc.GetTransactionResult, error) {
	if maxRetries < 0 {
		maxRetries = 0
	}

	var lastErr error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		tx, err := c.GetTransaction(ctx, sig)
		if err != nil {
			return nil, err
		}
		if tx != nil {
			return tx, nil
		}

		lastErr = errors.Errorf("transaction %s not available (null)", sig.String())
		if c.metrics != nil {
			c.metrics.RPCRetries.Add(1)
		}
		if attempt >= maxRetries {
			break
		}
		delay := calculateBackoff(attempt+1, c.retry)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(delay):
		}
	}
	return nil, lastErr
}

func (c *Client) GetSignaturesForAddress(
	ctx context.Context,
	address solana.PublicKey,
	opts *rpc.GetSignaturesForAddressOpts,
) ([]*rpc.TransactionSignature, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()

	if opts != nil && opts.Commitment == "" {
		opts.Commitment = c.commitment
	}

	txSigs, err := c.rpc.GetSignaturesForAddressWithOpts(ctx, address, opts)
	if err != nil {
		return nil, errors.Wrap(err, "get signatures for address")
	}
	return txSigs, nil
}

func calculateBackoff(attempt int, config *evm.RetryConfig) time.Duration {
	if config == nil {
		config = evm.DefaultRetryConfig()
	}
	delay := time.Duration(float64(config.BaseDelay) * math.Pow(config.BackoffRate, float64(attempt-1)))
	if config.MaxDelay > 0 && delay > config.MaxDelay {
		delay = config.MaxDelay
	}
	return delay
}
