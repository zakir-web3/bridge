package bridge

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventHashes(t *testing.T) {
	assert.Equal(t, "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef", ERC20TransferEventHash)
	assert.Equal(t, "0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96", DepositEventHash)
	assert.Equal(t, "0x04cafa25a7826c4415eac718e45fc84f69b5539748e0206adbae70f919887548", FinalizedWithdrawalEventHash)
}

func TestIsBridgeToken(t *testing.T) {
	usdc := common.HexToAddress("0x1111111111111111111111111111111111111111")
	other := common.HexToAddress("0x2222222222222222222222222222222222222222")
	cfg := Config{BridgeTokens: []common.Address{usdc}}

	require.True(t, cfg.IsBridgeToken(usdc))
	require.False(t, cfg.IsBridgeToken(other))
}

func TestGetFilterQuery_WatchesWhitelistTransfersIntoBridge(t *testing.T) {
	bridgeAddr := common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	tokenAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	b := &Bridge{cfg: Config{
		BridgeAddress: bridgeAddr,
		BridgeTokens:  []common.Address{tokenAddr},
	}}

	queries := b.GetFilterQuery(10, 20)
	require.Len(t, queries, 2)

	transfer := queries[0]
	require.Equal(t, big.NewInt(10), transfer.FromBlock)
	require.Equal(t, big.NewInt(20), transfer.ToBlock)
	require.Equal(t, []common.Address{tokenAddr}, transfer.Addresses)
	require.Equal(t, [][]common.Hash{
		{common.HexToHash(ERC20TransferEventHash)},
		nil,
		{common.BytesToHash(bridgeAddr.Bytes())},
	}, transfer.Topics)

	finalized := queries[1]
	require.Equal(t, []common.Address{bridgeAddr}, finalized.Addresses)
	require.Equal(t, [][]common.Hash{
		{common.HexToHash(FinalizedWithdrawalEventHash)},
	}, finalized.Topics)
}
