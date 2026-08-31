package solana

import (
	"bytes"
	"encoding/binary"
)

// EncodeDepositInstructionData returns serialized deposit instruction data for tests.
func EncodeDepositInstructionData(destination [20]byte, amount uint64) []byte {
	data := make([]byte, 8+20+8)
	copy(data[:8], depositInstructionDiscriminator)
	copy(data[8:28], destination[:])
	binary.LittleEndian.PutUint64(data[28:], amount)
	return data
}

// IsDepositInstruction reports whether instruction data targets the deposit handler.
func IsDepositInstruction(data []byte) bool {
	return len(data) >= 8 && bytes.Equal(data[:8], depositInstructionDiscriminator)
}
