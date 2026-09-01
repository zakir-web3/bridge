package solana

import (
	"math/big"
	"strings"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pkg/errors"

	"github.com/zakir-web3/bridge/internal/evm"
)

const (
	defaultSignaturesPageLimit = 1000
	defaultTxFetchRetries      = 5
	defaultRescanSlots         = 128
	defaultSlotLookback        = 64
)

type Config struct {
	SlotScannerConfig `mapstructure:",squash"        toml:",squash"`
	NodeURL           string          `mapstructure:"node_url"        toml:"node_url"`
	RetryConfig       evm.RetryConfig `mapstructure:",squash"        toml:",squash"`
	ChainID           *big.Int        `mapstructure:"chain_id"        toml:"chain_id"`
	ProgramID         solana.PublicKey
	ProgramIDStr      string `mapstructure:"program_id"      toml:"program_id"`
	BridgeMints       []solana.PublicKey
	BridgeMintsStr    []string `mapstructure:"bridge_mints"    toml:"bridge_mints"`
	Commitment        string   `mapstructure:"commitment"      toml:"commitment"`
}

func (c *Config) DecodePubkeys() error {
	if c.ProgramIDStr != "" {
		pub, err := solana.PublicKeyFromBase58(c.ProgramIDStr)
		if err != nil {
			return errors.Wrap(err, "program_id")
		}
		c.ProgramID = pub
	}
	c.BridgeMints = make([]solana.PublicKey, 0, len(c.BridgeMintsStr))
	for i, mint := range c.BridgeMintsStr {
		pub, err := solana.PublicKeyFromBase58(mint)
		if err != nil {
			return errors.Wrapf(err, "bridge_mints[%d]", i)
		}
		c.BridgeMints = append(c.BridgeMints, pub)
	}
	return nil
}

func (c Config) Enabled() bool {
	return c.ProgramIDStr != ""
}

func (c Config) CommitmentType() rpc.CommitmentType {
	switch strings.ToLower(strings.TrimSpace(c.Commitment)) {
	case "", "finalized":
		return rpc.CommitmentFinalized
	case "confirmed":
		return rpc.CommitmentConfirmed
	case "processed":
		return rpc.CommitmentProcessed
	default:
		return rpc.CommitmentFinalized
	}
}

func (c *Config) Validate() error {
	if !c.Enabled() {
		return nil
	}
	if c.NodeURL == "" {
		return errors.New("node_url is required")
	}
	if err := c.RetryConfig.Validate(); err != nil {
		return errors.Wrap(err, "retry_config")
	}
	if err := c.SlotScannerConfig.Validate(); err != nil {
		return err
	}
	if c.ChainID == nil || c.ChainID.Sign() <= 0 {
		return errors.New("chain_id is required")
	}
	if c.ProgramIDStr == "" {
		return errors.New("program_id is required")
	}
	if len(c.BridgeMintsStr) == 0 {
		return errors.New("bridge_mints is required")
	}
	return c.DecodePubkeys()
}

func (c Config) IsBridgeMint(mint solana.PublicKey) bool {
	for _, m := range c.BridgeMints {
		if m.Equals(mint) {
			return true
		}
	}
	return false
}

// SlotScannerConfig drives slot-based polling.
type SlotScannerConfig struct {
	Interval            time.Duration `mapstructure:"interval"                toml:"interval"`
	StartSlot           uint64        `mapstructure:"start_slot"                toml:"start_slot"`
	SlotInterval        uint64        `mapstructure:"slot_interval"             toml:"slot_interval"`
	SlotDelay           uint64        `mapstructure:"slot_delay"                toml:"slot_delay"`
	SlotLookback        uint64        `mapstructure:"slot_lookback"             toml:"slot_lookback"`
	RescanSlots         uint64        `mapstructure:"rescan_slots"              toml:"rescan_slots"`
	SignaturesPageLimit int           `mapstructure:"signatures_page_limit"     toml:"signatures_page_limit"`
	TxFetchRetries      int           `mapstructure:"tx_fetch_retries"          toml:"tx_fetch_retries"`
	ClearCache          bool          `mapstructure:"clear_cache"               toml:"clear_cache"`
}

func (c *SlotScannerConfig) applyDefaults() {
	if c.SlotLookback == 0 {
		c.SlotLookback = defaultSlotLookback
	}
	if c.RescanSlots == 0 {
		c.RescanSlots = defaultRescanSlots
	}
	if c.SignaturesPageLimit == 0 {
		c.SignaturesPageLimit = defaultSignaturesPageLimit
	}
	if c.TxFetchRetries == 0 {
		c.TxFetchRetries = defaultTxFetchRetries
	}
}

func (c SlotScannerConfig) Validate() error {
	if c.Interval == 0 {
		return errors.New("interval is required")
	}
	if c.SlotInterval == 0 {
		return errors.New("slot_interval is required")
	}
	return nil
}
