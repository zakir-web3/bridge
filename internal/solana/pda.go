package solana

import (
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
)

const noncesPerPage = 8192

func findConfigPDA(programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, programID)
}

func findValidatorSetPDA(programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("validator_set")}, programID)
}

func findVaultStatePDA(mint, programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("vault_state"), mint.Bytes()}, programID)
}

func findVaultAuthorityPDA(mint, programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("vault"), mint.Bytes()}, programID)
}

func findNoncePagePDA(nonce uint64, programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	page := nonce / noncesPerPage
	pageBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(pageBytes, page)
	return solana.FindProgramAddress([][]byte{[]byte("nonce_page"), pageBytes}, programID)
}

func findAssociatedTokenAddress(owner, mint solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindAssociatedTokenAddress(owner, mint)
	return addr, err
}
