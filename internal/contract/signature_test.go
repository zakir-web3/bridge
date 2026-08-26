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
		Destination: userAddr,
		Token:       tokenAddr,
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
	require.Equal(t, "0x2d6d11d3b045962a4c69066bfc5511cdb6f608d12fde188b54ef657ab7672573", message.String())

	digest := common.HexToHash("0xd77998b767c9c3e7c538f8fae2cd60480c9965bd8205cf3749bcb1b075e5ec9a")
	hash, _, err := apitypes.TypedDataAndHash(typedData)
	require.NoError(t, err)
	require.Equal(t, digest, common.BytesToHash(hash))
	require.Equal(t, digest, crypto.Keccak256Hash([]byte(fmt.Sprintf(
		"\x19\x01%s%s",
		string(domainSeparator.Bytes()),
		string(message.Bytes()),
	))))

	sig := common.Hex2Bytes("7573f2ea9d55409b8911b3f1826cae722cc201bf8778d4c3c4829db9cef3807b34368a0e9757596b772f5757afd5ab4d13666b2c45d0682796ac9ad0eb49085b01")
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
