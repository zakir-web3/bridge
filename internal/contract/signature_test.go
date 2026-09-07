package contract

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/require"
)

func TestRecoverSignerAddress(t *testing.T) {
	userAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	tokenAddr := common.HexToAddress("0x2000000000000000000000000000000000000002")
	verifyingContract := common.HexToAddress("0x3000000000000000000000000000000000000003")

	withdraw := &BridgeHubWithdraw{
		User:        userAddr,
		Destination: AddressToBytes32(userAddr),
		Token:       AddressToBytes32(tokenAddr),
		Amount:      big.NewInt(1e18),
		ChainId:     big.NewInt(97),
		Nonce:       12345,
		Raw: types.Log{
			Address:     verifyingContract,
			BlockNumber: 1,
			TxHash:      common.Hash{},
			Index:       2,
		},
	}
	typedData := withdraw.ToTypedData(nil, big.NewInt(1337), verifyingContract)

	domainSeparatorBytes, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	require.NoError(t, err)
	domainSeparator := common.BytesToHash(domainSeparatorBytes)
	require.Equal(t, "0xc41231ad1b7e9b5eed07b994b8348bb7fc45680241556e46317c680e8c588550", domainSeparator.String())

	messageBytes, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	require.NoError(t, err)
	message := common.BytesToHash(messageBytes)
	require.Equal(t, "0x2486f3baf331176f06d4eda1e971483466ca196931705258bc5ef971b81d4ad2", message.String())

	digest := common.HexToHash("0x04f7234d8a5f15bd6cd22589a81bbff694d8b707f81df216635d52a3934a84e8")
	hash, _, err := apitypes.TypedDataAndHash(typedData)
	require.NoError(t, err)
	require.Equal(t, digest, common.BytesToHash(hash))
	require.Equal(t, digest, crypto.Keccak256Hash([]byte(fmt.Sprintf(
		"\x19\x01%s%s",
		string(domainSeparator.Bytes()),
		string(message.Bytes()),
	))))

	sig := common.Hex2Bytes("529011f95cf9db893320030b9b22cd8535f68d4ffabed22a6509b5dfa57669963cf19f6e687d2b3dc808b514debe0d2772ad86853ea82bb94c89db47a4f65eac01")
	want := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	got, err := RecoverSignerAddress(domainSeparator, message, Signature{
		R: new(big.Int).SetBytes(sig[:32]),
		S: new(big.Int).SetBytes(sig[32:64]),
		V: sig[64] + 27,
	})
	require.NoError(t, err)
	require.Equal(t, want, got)

	pub, err := crypto.SigToPub(digest.Bytes(), sig)
	require.NoError(t, err)
	require.Equal(t, want, crypto.PubkeyToAddress(*pub))
}

func TestSignatureIsEmpty(t *testing.T) {
	require.True(t, (&Signature{}).IsEmpty())
	require.True(t, (&Signature{R: big.NewInt(0), S: big.NewInt(1)}).IsEmpty())
	require.True(t, (&Signature{R: big.NewInt(1), S: big.NewInt(0)}).IsEmpty())
	require.False(t, (&Signature{R: big.NewInt(1), S: big.NewInt(2), V: 28}).IsEmpty())
}
