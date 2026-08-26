package contract

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBridgeHubWithdraw_ToTypedData(t *testing.T) {
	userAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	destinationAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	tokenAddr := common.HexToAddress("0x2000000000000000000000000000000000000002")
	verifyingContract := common.HexToAddress("0x3000000000000000000000000000000000000003")

	srcChainId := big.NewInt(97)
	chainId := big.NewInt(1337)

	withdraw := &BridgeHubWithdraw{
		User:        userAddr,
		Destination: destinationAddr,
		Token:       tokenAddr,
		Amount:      big.NewInt(1e18),
		ChainId:     srcChainId,
		Nonce:       12345,
		Raw: types.Log{
			Address:     verifyingContract,
			BlockNumber: uint64(1),
			TxHash: common.HexToHash(
				"0x0000000000000000000000000000000000000000000000000000000000000000",
			),
			Index: uint(2),
		},
	}

	typedData := withdraw.ToTypedData(nil, chainId, verifyingContract)
	assert.NotNil(t, typedData)

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	assert.NoError(t, err)
	assert.Equal(t, "0xc41231ad1b7e9b5eed07b994b8348bb7fc45680241556e46317c680e8c588550", domainSeparator.String())

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	assert.NoError(t, err)
	assert.Equal(t, "0x2d6d11d3b045962a4c69066bfc5511cdb6f608d12fde188b54ef657ab7672573", typedDataHash.String())

	hash, _, err := apitypes.TypedDataAndHash(typedData)
	assert.NoError(t, err)
	assert.Equal(t, "0xd77998b767c9c3e7c538f8fae2cd60480c9965bd8205cf3749bcb1b075e5ec9a", common.BytesToHash(hash).String())
}

func TestBridgeHubRequestedValidatorSetUpdate_ToTypedData(t *testing.T) {
	verifyingContract := common.HexToAddress("0x3000000000000000000000000000000000000003")
	chainId := big.NewInt(1337)

	validatorSetUpdate := &BridgeHubRequestedValidatorSetUpdate{
		NewEpoch: 1,
		HotAddresses: []common.Address{
			common.HexToAddress("0x1000000000000000000000000000000000000001"),
		},
		ColdAddresses: []common.Address{
			common.HexToAddress("0x1000000000000000000000000000000000000002"),
		},
		Powers: []uint64{2},
	}

	typedData := validatorSetUpdate.ToTypedData(nil, chainId, verifyingContract)
	assert.NotNil(t, typedData)

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	assert.NoError(t, err)
	assert.NotEmpty(t, domainSeparator.String())

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	assert.NoError(t, err)
	assert.NotEmpty(t, typedDataHash.String())

	hash, _, err := apitypes.TypedDataAndHash(typedData)
	assert.NoError(t, err)
	assert.NotEmpty(t, common.BytesToHash(hash).String())
}

func TestValidatorSet_IsValidator(t *testing.T) {
	v1 := common.HexToAddress("0x1000000000000000000000000000000000000001")
	v2 := common.HexToAddress("0x1000000000000000000000000000000000000002")
	set := &ValidatorSet{Validators: []common.Address{v1}}

	assert.True(t, set.IsValidator(v1))
	assert.False(t, set.IsValidator(v2))
}

func TestMessageSignature_ToWithdrawalRequest(t *testing.T) {
	user := common.HexToAddress("0x1000000000000000000000000000000000000001")
	token := common.HexToAddress("0x2000000000000000000000000000000000000002")
	verifying := common.HexToAddress("0x3000000000000000000000000000000000000003")
	v0 := common.HexToAddress("0x4000000000000000000000000000000000000004")
	v1 := common.HexToAddress("0x5000000000000000000000000000000000000005")
	chainId := big.NewInt(1337)
	sig := Signature{R: big.NewInt(11), S: big.NewInt(22), V: 27}

	withdraw := &BridgeHubWithdraw{
		User:        user,
		Destination: user,
		Token:       token,
		Amount:      big.NewInt(1e18),
		Nonce:       12345,
	}
	typedData := withdraw.ToTypedData(nil, chainId, verifying)
	raw, err := typedData.EncodeData(typedData.PrimaryType, typedData.Message, 1)
	require.NoError(t, err)

	msg := &MessageSignature{
		RawData:    raw,
		Signers:    []common.Address{v0, v1},
		Signatures: []Signature{{}, sig},
	}
	got, err := msg.ToWithdrawalRequest(ValidatorSet{Validators: []common.Address{v0, v1}})
	require.NoError(t, err)
	assert.Equal(t, user, got.User)
	assert.Equal(t, user, got.Destination)
	assert.Equal(t, token, got.Token)
	assert.Equal(t, big.NewInt(1e18), got.Amount)
	assert.Equal(t, chainId, got.ChainId)
	assert.Equal(t, uint64(12345), got.Nonce)
	assert.Equal(t, []Signature{sig, sig}, got.Signatures)
}

func TestMessageSignature_ToWithdrawalRequest_RejectsShortData(t *testing.T) {
	_, err := (&MessageSignature{RawData: []byte{1, 2, 3}}).ToWithdrawalRequest(ValidatorSet{})
	require.EqualError(t, err, "raw data too short")
}

func TestMessageSignature_ToWithdrawalRequest_RejectsSignerMismatch(t *testing.T) {
	user := common.HexToAddress("0x1000000000000000000000000000000000000001")
	token := common.HexToAddress("0x2000000000000000000000000000000000000002")
	verifying := common.HexToAddress("0x3000000000000000000000000000000000000003")
	typedData := (&BridgeHubWithdraw{
		User:        user,
		Destination: user,
		Token:       token,
		Amount:      big.NewInt(1),
		Nonce:       1,
	}).ToTypedData(nil, big.NewInt(1337), verifying)
	raw, err := typedData.EncodeData(typedData.PrimaryType, typedData.Message, 1)
	require.NoError(t, err)

	_, err = (&MessageSignature{
		RawData: raw,
		Signers: []common.Address{common.HexToAddress("0x1")},
	}).ToWithdrawalRequest(ValidatorSet{Validators: []common.Address{common.HexToAddress("0x2")}})
	require.EqualError(t, err, "signers do not match validator set")
}
