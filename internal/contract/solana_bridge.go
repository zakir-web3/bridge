package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// VerifyingContractFromProgramID derives the EIP-712 verifyingContract pseudo-address
// used by the Solana bridge program: keccak256(program_id)[12..32].
func VerifyingContractFromProgramID(programID [32]byte) common.Address {
	hash := crypto.Keccak256Hash(programID[:])
	return common.BytesToAddress(hash[12:32])
}

// SolanaBridgeDomainSeparator returns the Bridge-domain EIP-712 separator for a Solana
// source chain (chainId + program-derived verifyingContract).
func SolanaBridgeDomainSeparator(chainID *big.Int, programID [32]byte) (common.Hash, error) {
	verifyingContract := VerifyingContractFromProgramID(programID)
	domain := NewBridgeDomain(chainID, verifyingContract)
	typedData := apitypes.TypedData{
		Domain: domain,
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "EIP712Domain",
		Message:     domain.Map(),
	}
	separator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(separator), nil
}

// BridgeHubWithdrawSolanaTypedData builds Bridge-domain typed data for a Hub withdraw
// destined for Solana (verifyingContract derived from program ID, not EVM Bridge address).
func BridgeHubWithdrawSolanaTypedData(
	w *BridgeHubWithdraw,
	srcChainID *big.Int,
	programID [32]byte,
) apitypes.TypedData {
	verifyingContract := VerifyingContractFromProgramID(programID)
	return w.ToTypedData(nil, srcChainID, verifyingContract)
}
