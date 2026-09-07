package contract

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)

// AddressToBytes32 right-pads a 20-byte EVM address into bytes32.
func AddressToBytes32(addr common.Address) [32]byte {
	var out [32]byte
	copy(out[12:], addr.Bytes())
	return out
}

// Bytes32ToAddress extracts a 20-byte EVM address from bytes32.
func Bytes32ToAddress(value [32]byte) common.Address {
	return common.BytesToAddress(value[12:])
}

// IsEVMBytes32 reports whether value encodes an EVM address in the low 20 bytes.
func IsEVMBytes32(value [32]byte) bool {
	for i := 0; i < 12; i++ {
		if value[i] != 0 {
			return false
		}
	}
	return true
}

// Bytes32Hex returns the canonical 0x-prefixed hex encoding for EIP-712 bytes32 fields.
func Bytes32Hex(value [32]byte) string {
	return "0x" + hex.EncodeToString(value[:])
}
