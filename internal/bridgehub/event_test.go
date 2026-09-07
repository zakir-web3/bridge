package bridgehub

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/zakir-web3/bridge/internal/contract"
)

func TestEventHashes(t *testing.T) {
	assert.Equal(t, "0xddf7473863baaba91098e18c11ea6e972fb72f38f768f5e5d3e984207414ff13", WithdrawEventHash)
	assert.Equal(t, "0xc36d3c025c0f565ea8451b940f66ac343971b6b6d515ea860aed580c6408af2b", BridgeSignatureSubmittedEventHash)
}

func TestGetFilterQuery_WatchesWithdrawAndBridgeSignatures(t *testing.T) {
	hubAddr := common.HexToAddress("0xcccccccccccccccccccccccccccccccccccccccc")
	h := &BridgeHub{cfg: Config{BridgeHubAddress: hubAddr}}

	queries := h.GetFilterQuery(5, 15)
	require.Len(t, queries, 1)
	require.Equal(t, big.NewInt(5), queries[0].FromBlock)
	require.Equal(t, big.NewInt(15), queries[0].ToBlock)
	require.Equal(t, []common.Address{hubAddr}, queries[0].Addresses)
	require.Equal(t, [][]common.Hash{{
		common.HexToHash(WithdrawEventHash),
		common.HexToHash(BridgeSignatureSubmittedEventHash),
	}}, queries[0].Topics)
}

func TestHaveEnoughPower_RequiresStrictTwoThirds(t *testing.T) {
	h := &BridgeHub{validatorSet: &contract.ValidatorSet{Powers: []uint64{70, 30}}}

	// total=100, threshold is 3*power > 200, so power >= 67
	require.False(t, h.HaveEnoughPower(66))
	require.True(t, h.HaveEnoughPower(67))
	require.True(t, h.HaveEnoughPower(100))
}
