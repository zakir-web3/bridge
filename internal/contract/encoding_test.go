package contract

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestAddressToBytes32(t *testing.T) {
	addr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	got := AddressToBytes32(addr)
	require.True(t, IsEVMBytes32(got))
	require.Equal(t, addr, Bytes32ToAddress(got))
}

func TestIsEVMBytes32(t *testing.T) {
	var solanaMint [32]byte
	solanaMint[0] = 1
	require.False(t, IsEVMBytes32(solanaMint))
	require.True(t, IsEVMBytes32(AddressToBytes32(common.HexToAddress("0x1"))))
}
