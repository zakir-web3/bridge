package contract

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	solanaProgramIDC4Yxx = "C4YxxrnCKnE4hVdTPcmTZN6yuHp5U9xVXRs3VanEeYfq"
	// CI deploy ID after anchor keys sync; cross-checked in withdraw.rs Go conformance tests.
	solanaProgramIDAlt = "5SPkxd2xhiG3aZbnW9ro7gdfhg3AZ44iR1k5rt9DJJ2"
	solanaCanonicalChainID = 900001
)

func TestVerifyingContractFromProgramID(t *testing.T) {
	programID := solana.MustPublicKeyFromBase58(solanaProgramIDC4Yxx)
	got := VerifyingContractFromProgramID(programID)
	assert.Equal(t, common.HexToAddress("0xdb94ec3d773ea0a5b9b89b4bf28ed21dec3f0f8f"), got)
}

func TestSolanaBridgeDomainSeparator(t *testing.T) {
	programID := solana.MustPublicKeyFromBase58(solanaProgramIDC4Yxx)
	separator, err := SolanaBridgeDomainSeparator(big.NewInt(solanaCanonicalChainID), programID)
	require.NoError(t, err)
	assert.Equal(t, "0xe038ca293e650b49e9781d6f45d165e6ac0f202e3e6d4e00c07072d8088c3633", separator.Hex())
}

// Conformance: solana/programs/bridge/src/withdraw.rs test_domain_separator_go_vector_alt_program_id.
func TestSolanaBridgeDomainSeparator_AltProgramIDMatchesRustVector(t *testing.T) {
	programID := solana.MustPublicKeyFromBase58(solanaProgramIDAlt)
	separator, err := SolanaBridgeDomainSeparator(big.NewInt(solanaCanonicalChainID), programID)
	require.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0xE723c846F25a42d2782D5bb51a5485A6bf46D395"), VerifyingContractFromProgramID(programID))
	assert.Equal(t, "0x3347d4f0929752e8b301efa4e933d2f41e8d0c4849ad2e39f3f75482422536ef", separator.Hex())
}

func TestSolanaBridgeDomainSeparator_DifferentProgramIDsDiffer(t *testing.T) {
	c4yxx := solana.MustPublicKeyFromBase58(solanaProgramIDC4Yxx)
	alt := solana.MustPublicKeyFromBase58(solanaProgramIDAlt)
	chainID := big.NewInt(solanaCanonicalChainID)

	dsC4, err := SolanaBridgeDomainSeparator(chainID, c4yxx)
	require.NoError(t, err)
	dsAlt, err := SolanaBridgeDomainSeparator(chainID, alt)
	require.NoError(t, err)
	require.NotEqual(t, dsC4, dsAlt)
}

// Conformance: solana/programs/bridge/src/withdraw.rs test_full_digest_solana_withdraw_go_vector_alt_program_id.
func TestBridgeHubWithdrawSolanaTypedData_DigestMatchesRustVector(t *testing.T) {
	chainID := big.NewInt(solanaCanonicalChainID)
	programID := solana.MustPublicKeyFromBase58(solanaProgramIDAlt)

	user := common.HexToAddress("0x1000000000000000000000000000000000000001")
	withdraw := &BridgeHubWithdraw{
		User:        user,
		Destination: AddressToBytes32(user),
		Token:       AddressToBytes32(common.HexToAddress("0x2000000000000000000000000000000000000002")),
		Amount:      big.NewInt(1e18),
		Nonce:       12345,
	}

	typedData := BridgeHubWithdrawSolanaTypedData(withdraw, chainID, programID)

	domainSep, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	require.NoError(t, err)
	assert.Equal(t, "0x3347d4f0929752e8b301efa4e933d2f41e8d0c4849ad2e39f3f75482422536ef", common.BytesToHash(domainSep).Hex())

	structHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	require.NoError(t, err)
	assert.Equal(t, "0x0750425c9837f1a20d7ad334d5677fc7b06aa5c9029113dc03ee79df64f092fe", common.BytesToHash(structHash).Hex())

	digest, _, err := apitypes.TypedDataAndHash(typedData)
	require.NoError(t, err)
	assert.Equal(t, "0x98bc4e91da645a2c1d0aa949fe25f35836513332ae72a5d4d21127edb8408f6d", common.BytesToHash(digest).Hex())
}
