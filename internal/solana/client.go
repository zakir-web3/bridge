package solana

import (
	"context"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pkg/errors"
)

const defaultRPCTimeout = 30 * time.Second

// defaultCommitment balances devnet finality with local-test responsiveness.
var defaultCommitment = rpc.CommitmentConfirmed

// Client wraps Solana RPC with finalized commitment.
type Client struct {
	rpc *rpc.Client
}

func NewClient(nodeURL string) (*Client, error) {
	if nodeURL == "" {
		return nil, errors.New("node_url is required")
	}
	return &Client{rpc: rpc.New(nodeURL)}, nil
}

func (c *Client) GetSlot(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultRPCTimeout)
	defer cancel()
	slot, err := c.rpc.GetSlot(ctx, defaultCommitment)
	if err != nil {
		return 0, errors.Wrap(err, "get slot")
	}
	return slot, nil
}

func (c *Client) GetTransaction(ctx context.Context, sig solana.Signature) (*rpc.GetTransactionResult, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultRPCTimeout)
	defer cancel()
	version := uint64(0)
	result, err := c.rpc.GetTransaction(ctx, sig, &rpc.GetTransactionOpts{
		Encoding:                       solana.EncodingBase64,
		Commitment:                     defaultCommitment,
		MaxSupportedTransactionVersion: &version,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get transaction")
	}
	return result, nil
}

func (c *Client) GetSignaturesForAddress(
	ctx context.Context,
	address solana.PublicKey,
	opts *rpc.GetSignaturesForAddressOpts,
) ([]*rpc.TransactionSignature, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultRPCTimeout)
	defer cancel()
	sigs, err := c.rpc.GetSignaturesForAddressWithOpts(ctx, address, opts)
	if err != nil {
		return nil, errors.Wrap(err, "get signatures for address")
	}
	return sigs, nil
}
