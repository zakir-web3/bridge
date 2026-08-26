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
	assert.Equal(t, "0x7ba97cc0a03c9de2b8d972e5a42d8867d1c3483652b679f5b8b8b5b62a4ef576", WithdrawEventHash)
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
