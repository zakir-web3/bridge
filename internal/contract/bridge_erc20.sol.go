// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BridgeERC20MetaData contains all meta data concerning the BridgeERC20 contract.
var BridgeERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"UnBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"setMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101606040523480156200001257600080fd5b5060405162001fe038038062001fe0833981016040819052620000359162000372565b6040805180820190915260018152603160f81b602082015282908190818460036200006183826200046d565b5060046200007082826200046d565b5062000082915083905060056200016e565b61012052620000938160066200016e565b61014052815160208084019190912060e052815190820120610100524660a0526200012160e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525062000138600033620001a7565b50620001657f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633620001a7565b50505062000593565b60006020835110156200018e57620001868362000259565b9050620001a1565b816200019b84826200046d565b5060ff90505b92915050565b60008281526008602090815260408083206001600160a01b038516845290915281205460ff16620002505760008381526008602090815260408083206001600160a01b03861684529091529020805460ff19166001179055620002073390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001620001a1565b506000620001a1565b600080829050601f8151111562000290578260405163305a27a960e01b815260040162000287919062000539565b60405180910390fd5b80516200029d826200056e565b179392505050565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620002d8578181015183820152602001620002be565b50506000910152565b600082601f830112620002f357600080fd5b81516001600160401b0380821115620003105762000310620002a5565b604051601f8301601f19908116603f011681019082821181831017156200033b576200033b620002a5565b816040528381528660208588010111156200035557600080fd5b62000368846020830160208901620002bb565b9695505050505050565b600080604083850312156200038657600080fd5b82516001600160401b03808211156200039e57600080fd5b620003ac86838701620002e1565b93506020850151915080821115620003c357600080fd5b50620003d285828601620002e1565b9150509250929050565b600181811c90821680620003f157607f821691505b6020821081036200041257634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000468576000816000526020600020601f850160051c81016020861015620004435750805b601f850160051c820191505b8181101562000464578281556001016200044f565b5050505b505050565b81516001600160401b03811115620004895762000489620002a5565b620004a1816200049a8454620003dc565b8462000418565b602080601f831160018114620004d95760008415620004c05750858301515b600019600386901b1c1916600185901b17855562000464565b600085815260208120601f198616915b828110156200050a57888601518255948401946001909101908401620004e9565b5085821015620005295787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60208152600082518060208401526200055a816040850160208701620002bb565b601f01601f19169190910160400192915050565b80516020808301519190811015620004125760001960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516119f2620005ee6000396000610f9701526000610f6a01526000610da601526000610d7e01526000610cd901526000610d0301526000610d2d01526119f26000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80635c975abb1161010f57806395d89b41116100a2578063d539139311610071578063d5391393146103db578063d547741f146103f0578063dd62ed3e14610403578063fe575a871461043c57600080fd5b806395d89b41146103ad578063a217fddf14610349578063a9059cbb146103b5578063d505accf146103c857600080fd5b80637ecebe00116100de5780637ecebe00146103645780638456cb591461037757806384b0196e1461037f57806391d148541461039a57600080fd5b80635c975abb1461031557806370a082311461032057806375b238fc1461034957806378b0bac41461035157600080fd5b8063313ce5671161018757806340c10f191161015657806340c10f19146102c957806342966c68146102dc57806344337ea1146102ef578063537df3b61461030257600080fd5b8063313ce567146102975780633644e515146102a657806336568abe146102ae5780633f4ba83a146102c157600080fd5b806318160ddd116101c357806318160ddd1461023a57806323b872dd1461024c578063248a9ca31461025f5780632f2ff15d1461028257600080fd5b806301ffc9a7146101ea57806306fdde0314610212578063095ea7b314610227575b600080fd5b6101fd6101f836600461164d565b610468565b60405190151581526020015b60405180910390f35b61021a61049f565b60405161020991906116c4565b6101fd6102353660046116f3565b610531565b6002545b604051908152602001610209565b6101fd61025a36600461171d565b610549565b61023e61026d366004611759565b60009081526008602052604090206001015490565b610295610290366004611772565b61056d565b005b60405160128152602001610209565b61023e610598565b6102956102bc366004611772565b6105a7565b6102956105df565b6102956102d73660046116f3565b6105f5565b6102956102ea366004611759565b61065f565b6102956102fd36600461179e565b6106a1565b61029561031036600461179e565b6107a5565b60095460ff166101fd565b61023e61032e36600461179e565b6001600160a01b031660009081526020819052604090205490565b61023e600081565b61029561035f3660046117b9565b61089c565b61023e61037236600461179e565b61092b565b610295610949565b61038761095c565b60405161020997969594939291906117f5565b6101fd6103a8366004611772565b6109a2565b61021a6109cd565b6101fd6103c33660046116f3565b6109dc565b6102956103d636600461188e565b6109ea565b61023e60008051602061199d83398151915281565b6102956103fe366004611772565b610b24565b61023e610411366004611901565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6101fd61044a36600461179e565b6001600160a01b03166000908152600a602052604090205460ff1690565b60006001600160e01b03198216637965db0b60e01b148061049957506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600380546104ae9061192b565b80601f01602080910402602001604051908101604052809291908181526020018280546104da9061192b565b80156105275780601f106104fc57610100808354040283529160200191610527565b820191906000526020600020905b81548152906001019060200180831161050a57829003601f168201915b5050505050905090565b60003361053f818585610b49565b5060019392505050565b600033610557858285610b56565b610562858585610bcf565b506001949350505050565b60008281526008602052604090206001015461058881610c2e565b6105928383610c38565b50505050565b60006105a2610ccc565b905090565b6001600160a01b03811633146105d05760405163334bd91960e11b815260040160405180910390fd5b6105da8282610df7565b505050565b60006105ea81610c2e565b6105f2610e64565b50565b60008051602061199d83398151915261060d81610c2e565b6106178383610eb6565b826001600160a01b03167f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe8360405161065291815260200190565b60405180910390a2505050565b6106693382610ef0565b60405181815233907f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df79060200160405180910390a250565b60006106ac81610c2e565b6001600160a01b0382166106f95760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b60448201526064015b60405180910390fd5b6001600160a01b0382166000908152600a602052604090205460ff16156107585760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e48189b1858dadb1a5cdd1959606a1b60448201526064016106f0565b6001600160a01b0382166000818152600a6020526040808220805460ff19166001179055517fffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b8559190a25050565b60006107b081610c2e565b6001600160a01b0382166107f85760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b60448201526064016106f0565b6001600160a01b0382166000908152600a602052604090205460ff166108525760405162461bcd60e51b815260206004820152600f60248201526e139bdd08189b1858dadb1a5cdd1959608a1b60448201526064016106f0565b6001600160a01b0382166000818152600a6020526040808220805460ff19169055517f117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e9190a25050565b60006108a781610c2e565b6001600160a01b0383166108f55760405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206d696e6572206164647265737360581b60448201526064016106f0565b81156109135761059260008051602061199d83398151915284610c38565b61059260008051602061199d83398151915284610df7565b6001600160a01b038116600090815260076020526040812054610499565b600061095481610c2e565b6105f2610f26565b600060608060008060006060610970610f63565b610978610f90565b60408051600080825260208201909252600f60f81b9b939a50919850469750309650945092509050565b60009182526008602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6060600480546104ae9061192b565b60003361053f818585610bcf565b83421115610a0e5760405163313c898160e11b8152600481018590526024016106f0565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610a5b8c6001600160a01b0316600090815260076020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610ab682610fbd565b90506000610ac682878787610fea565b9050896001600160a01b0316816001600160a01b031614610b0d576040516325c0072360e11b81526001600160a01b0380831660048301528b1660248201526044016106f0565b610b188a8a8a610b49565b50505050505050505050565b600082815260086020526040902060010154610b3f81610c2e565b6105928383610df7565b6105da8383836001611018565b6001600160a01b038381166000908152600160209081526040808320938616835292905220546000198110156105925781811015610bc057604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016106f0565b61059284848484036000611018565b6001600160a01b038316610bf957604051634b637e8f60e11b8152600060048201526024016106f0565b6001600160a01b038216610c235760405163ec442f0560e01b8152600060048201526024016106f0565b6105da8383836110ed565b6105f28133611207565b6000610c4483836109a2565b610cc45760008381526008602090815260408083206001600160a01b03861684529091529020805460ff19166001179055610c7c3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610499565b506000610499565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015610d2557507f000000000000000000000000000000000000000000000000000000000000000046145b15610d4f57507f000000000000000000000000000000000000000000000000000000000000000090565b6105a2604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b6000610e0383836109a2565b15610cc45760008381526008602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610499565b610e6c611240565b6009805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b038216610ee05760405163ec442f0560e01b8152600060048201526024016106f0565b610eec600083836110ed565b5050565b6001600160a01b038216610f1a57604051634b637e8f60e11b8152600060048201526024016106f0565b610eec826000836110ed565b610f2e611265565b6009805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610e993390565b60606105a27f00000000000000000000000000000000000000000000000000000000000000006005611289565b60606105a27f00000000000000000000000000000000000000000000000000000000000000006006611289565b6000610499610fca610ccc565b8360405161190160f01b8152600281019290925260228201526042902090565b600080600080610ffc88888888611334565b92509250925061100c8282611403565b50909695505050505050565b6001600160a01b0384166110425760405163e602df0560e01b8152600060048201526024016106f0565b6001600160a01b03831661106c57604051634a1406b160e11b8152600060048201526024016106f0565b6001600160a01b038085166000908152600160209081526040808320938716835292905220829055801561059257826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516110df91815260200190565b60405180910390a350505050565b60095460ff16156111325760405162461bcd60e51b815260206004820152600f60248201526e151bdad95b881a5cc81c185d5cd959608a1b60448201526064016106f0565b6001600160a01b0383166000908152600a602052604090205460ff16156111935760405162461bcd60e51b815260206004820152601560248201527414d95b99195c881a5cc8189b1858dadb1a5cdd1959605a1b60448201526064016106f0565b6001600160a01b0382166000908152600a602052604090205460ff16156111fc5760405162461bcd60e51b815260206004820152601860248201527f526563697069656e7420697320626c61636b6c6973746564000000000000000060448201526064016106f0565b6105da8383836114bc565b61121182826109a2565b610eec5760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016106f0565b60095460ff1661126357604051638dfc202b60e01b815260040160405180910390fd5b565b60095460ff16156112635760405163d93c066560e01b815260040160405180910390fd5b606060ff83146112a35761129c836115e6565b9050610499565b8180546112af9061192b565b80601f01602080910402602001604051908101604052809291908181526020018280546112db9061192b565b80156113285780601f106112fd57610100808354040283529160200191611328565b820191906000526020600020905b81548152906001019060200180831161130b57829003601f168201915b50505050509050610499565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561136f57506000915060039050826113f9565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156113c3573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166113ef575060009250600191508290506113f9565b9250600091508190505b9450945094915050565b600082600381111561141757611417611965565b03611420575050565b600182600381111561143457611434611965565b036114525760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561146657611466611965565b036114875760405163fce698f760e01b8152600481018290526024016106f0565b600382600381111561149b5761149b611965565b03610eec576040516335e2f38360e21b8152600481018290526024016106f0565b6001600160a01b0383166114e75780600260008282546114dc919061197b565b909155506115599050565b6001600160a01b0383166000908152602081905260409020548181101561153a5760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016106f0565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b03821661157557600280548290039055611594565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516115d991815260200190565b60405180910390a3505050565b606060006115f383611625565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b600060ff8216601f81111561049957604051632cd44ac360e21b815260040160405180910390fd5b60006020828403121561165f57600080fd5b81356001600160e01b03198116811461167757600080fd5b9392505050565b6000815180845260005b818110156116a457602081850181015186830182015201611688565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000611677602083018461167e565b80356001600160a01b03811681146116ee57600080fd5b919050565b6000806040838503121561170657600080fd5b61170f836116d7565b946020939093013593505050565b60008060006060848603121561173257600080fd5b61173b846116d7565b9250611749602085016116d7565b9150604084013590509250925092565b60006020828403121561176b57600080fd5b5035919050565b6000806040838503121561178557600080fd5b82359150611795602084016116d7565b90509250929050565b6000602082840312156117b057600080fd5b611677826116d7565b600080604083850312156117cc57600080fd5b6117d5836116d7565b9150602083013580151581146117ea57600080fd5b809150509250929050565b60ff60f81b881681526000602060e0602084015261181660e084018a61167e565b8381036040850152611828818a61167e565b606085018990526001600160a01b038816608086015260a0850187905284810360c08601528551808252602080880193509091019060005b8181101561187c57835183529284019291840191600101611860565b50909c9b505050505050505050505050565b600080600080600080600060e0888a0312156118a957600080fd5b6118b2886116d7565b96506118c0602089016116d7565b95506040880135945060608801359350608088013560ff811681146118e457600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561191457600080fd5b61191d836116d7565b9150611795602084016116d7565b600181811c9082168061193f57607f821691505b60208210810361195f57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052602160045260246000fd5b8082018082111561049957634e487b7160e01b600052601160045260246000fdfe9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6a2646970667358221220a57ed93536b7c55a6c7989b5ef97055678094875e4fdd19e559852a43181241864736f6c63430008160033",
}

// BridgeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeERC20MetaData.ABI instead.
var BridgeERC20ABI = BridgeERC20MetaData.ABI

// BridgeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeERC20MetaData.Bin instead.
var BridgeERC20Bin = BridgeERC20MetaData.Bin

// DeployBridgeERC20 deploys a new Ethereum contract, binding an instance of BridgeERC20 to it.
func DeployBridgeERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *BridgeERC20, error) {
	parsed, err := BridgeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeERC20Bin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeERC20{BridgeERC20Caller: BridgeERC20Caller{contract: contract}, BridgeERC20Transactor: BridgeERC20Transactor{contract: contract}, BridgeERC20Filterer: BridgeERC20Filterer{contract: contract}}, nil
}

// BridgeERC20 is an auto generated Go binding around an Ethereum contract.
type BridgeERC20 struct {
	BridgeERC20Caller     // Read-only binding to the contract
	BridgeERC20Transactor // Write-only binding to the contract
	BridgeERC20Filterer   // Log filterer for contract events
}

// BridgeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeERC20Session struct {
	Contract     *BridgeERC20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeERC20CallerSession struct {
	Contract *BridgeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeERC20TransactorSession struct {
	Contract     *BridgeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeERC20Raw struct {
	Contract *BridgeERC20 // Generic contract binding to access the raw methods on
}

// BridgeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeERC20CallerRaw struct {
	Contract *BridgeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// BridgeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeERC20TransactorRaw struct {
	Contract *BridgeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeERC20 creates a new instance of BridgeERC20, bound to a specific deployed contract.
func NewBridgeERC20(address common.Address, backend bind.ContractBackend) (*BridgeERC20, error) {
	contract, err := bindBridgeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20{BridgeERC20Caller: BridgeERC20Caller{contract: contract}, BridgeERC20Transactor: BridgeERC20Transactor{contract: contract}, BridgeERC20Filterer: BridgeERC20Filterer{contract: contract}}, nil
}

// NewBridgeERC20Caller creates a new read-only instance of BridgeERC20, bound to a specific deployed contract.
func NewBridgeERC20Caller(address common.Address, caller bind.ContractCaller) (*BridgeERC20Caller, error) {
	contract, err := bindBridgeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20Caller{contract: contract}, nil
}

// NewBridgeERC20Transactor creates a new write-only instance of BridgeERC20, bound to a specific deployed contract.
func NewBridgeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*BridgeERC20Transactor, error) {
	contract, err := bindBridgeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20Transactor{contract: contract}, nil
}

// NewBridgeERC20Filterer creates a new log filterer instance of BridgeERC20, bound to a specific deployed contract.
func NewBridgeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*BridgeERC20Filterer, error) {
	contract, err := bindBridgeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20Filterer{contract: contract}, nil
}

// bindBridgeERC20 binds a generic wrapper to an already deployed contract.
func bindBridgeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeERC20 *BridgeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeERC20.Contract.BridgeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeERC20 *BridgeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeERC20.Contract.BridgeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeERC20 *BridgeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeERC20.Contract.BridgeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeERC20 *BridgeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeERC20 *BridgeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeERC20 *BridgeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeERC20.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Caller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Session) ADMINROLE() ([32]byte, error) {
	return _BridgeERC20.Contract.ADMINROLE(&_BridgeERC20.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20CallerSession) ADMINROLE() ([32]byte, error) {
	return _BridgeERC20.Contract.ADMINROLE(&_BridgeERC20.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeERC20.Contract.DEFAULTADMINROLE(&_BridgeERC20.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeERC20.Contract.DEFAULTADMINROLE(&_BridgeERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _BridgeERC20.Contract.DOMAINSEPARATOR(&_BridgeERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BridgeERC20.Contract.DOMAINSEPARATOR(&_BridgeERC20.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Caller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Session) MINTERROLE() ([32]byte, error) {
	return _BridgeERC20.Contract.MINTERROLE(&_BridgeERC20.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_BridgeERC20 *BridgeERC20CallerSession) MINTERROLE() ([32]byte, error) {
	return _BridgeERC20.Contract.MINTERROLE(&_BridgeERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BridgeERC20 *BridgeERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BridgeERC20 *BridgeERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeERC20.Contract.Allowance(&_BridgeERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BridgeERC20 *BridgeERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeERC20.Contract.Allowance(&_BridgeERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BridgeERC20 *BridgeERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BridgeERC20 *BridgeERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BridgeERC20.Contract.BalanceOf(&_BridgeERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BridgeERC20 *BridgeERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BridgeERC20.Contract.BalanceOf(&_BridgeERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_BridgeERC20 *BridgeERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_BridgeERC20 *BridgeERC20Session) Decimals() (uint8, error) {
	return _BridgeERC20.Contract.Decimals(&_BridgeERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_BridgeERC20 *BridgeERC20CallerSession) Decimals() (uint8, error) {
	return _BridgeERC20.Contract.Decimals(&_BridgeERC20.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgeERC20 *BridgeERC20Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgeERC20 *BridgeERC20Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BridgeERC20.Contract.Eip712Domain(&_BridgeERC20.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgeERC20 *BridgeERC20CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BridgeERC20.Contract.Eip712Domain(&_BridgeERC20.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeERC20 *BridgeERC20Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeERC20.Contract.GetRoleAdmin(&_BridgeERC20.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeERC20 *BridgeERC20CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeERC20.Contract.GetRoleAdmin(&_BridgeERC20.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeERC20 *BridgeERC20Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeERC20 *BridgeERC20Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeERC20.Contract.HasRole(&_BridgeERC20.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeERC20 *BridgeERC20CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeERC20.Contract.HasRole(&_BridgeERC20.CallOpts, role, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_BridgeERC20 *BridgeERC20Caller) IsBlacklisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "isBlacklisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_BridgeERC20 *BridgeERC20Session) IsBlacklisted(account common.Address) (bool, error) {
	return _BridgeERC20.Contract.IsBlacklisted(&_BridgeERC20.CallOpts, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_BridgeERC20 *BridgeERC20CallerSession) IsBlacklisted(account common.Address) (bool, error) {
	return _BridgeERC20.Contract.IsBlacklisted(&_BridgeERC20.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeERC20 *BridgeERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeERC20 *BridgeERC20Session) Name() (string, error) {
	return _BridgeERC20.Contract.Name(&_BridgeERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeERC20 *BridgeERC20CallerSession) Name() (string, error) {
	return _BridgeERC20.Contract.Name(&_BridgeERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BridgeERC20 *BridgeERC20Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BridgeERC20 *BridgeERC20Session) Nonces(owner common.Address) (*big.Int, error) {
	return _BridgeERC20.Contract.Nonces(&_BridgeERC20.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BridgeERC20 *BridgeERC20CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BridgeERC20.Contract.Nonces(&_BridgeERC20.CallOpts, owner)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeERC20 *BridgeERC20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeERC20 *BridgeERC20Session) Paused() (bool, error) {
	return _BridgeERC20.Contract.Paused(&_BridgeERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeERC20 *BridgeERC20CallerSession) Paused() (bool, error) {
	return _BridgeERC20.Contract.Paused(&_BridgeERC20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeERC20 *BridgeERC20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeERC20 *BridgeERC20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeERC20.Contract.SupportsInterface(&_BridgeERC20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeERC20 *BridgeERC20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeERC20.Contract.SupportsInterface(&_BridgeERC20.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeERC20 *BridgeERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeERC20 *BridgeERC20Session) Symbol() (string, error) {
	return _BridgeERC20.Contract.Symbol(&_BridgeERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeERC20 *BridgeERC20CallerSession) Symbol() (string, error) {
	return _BridgeERC20.Contract.Symbol(&_BridgeERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BridgeERC20 *BridgeERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BridgeERC20 *BridgeERC20Session) TotalSupply() (*big.Int, error) {
	return _BridgeERC20.Contract.TotalSupply(&_BridgeERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BridgeERC20 *BridgeERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _BridgeERC20.Contract.TotalSupply(&_BridgeERC20.CallOpts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(address account) returns()
func (_BridgeERC20 *BridgeERC20Transactor) AddToBlacklist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "addToBlacklist", account)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(address account) returns()
func (_BridgeERC20 *BridgeERC20Session) AddToBlacklist(account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.AddToBlacklist(&_BridgeERC20.TransactOpts, account)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(address account) returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) AddToBlacklist(account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.AddToBlacklist(&_BridgeERC20.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BridgeERC20 *BridgeERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BridgeERC20 *BridgeERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Approve(&_BridgeERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BridgeERC20 *BridgeERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Approve(&_BridgeERC20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_BridgeERC20 *BridgeERC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_BridgeERC20 *BridgeERC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Burn(&_BridgeERC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Burn(&_BridgeERC20.TransactOpts, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeERC20 *BridgeERC20Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeERC20 *BridgeERC20Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.GrantRole(&_BridgeERC20.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.GrantRole(&_BridgeERC20.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BridgeERC20 *BridgeERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BridgeERC20 *BridgeERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Mint(&_BridgeERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Mint(&_BridgeERC20.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeERC20 *BridgeERC20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeERC20 *BridgeERC20Session) Pause() (*types.Transaction, error) {
	return _BridgeERC20.Contract.Pause(&_BridgeERC20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeERC20.Contract.Pause(&_BridgeERC20.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BridgeERC20 *BridgeERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BridgeERC20 *BridgeERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Permit(&_BridgeERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Permit(&_BridgeERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x537df3b6.
//
// Solidity: function removeFromBlacklist(address account) returns()
func (_BridgeERC20 *BridgeERC20Transactor) RemoveFromBlacklist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "removeFromBlacklist", account)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x537df3b6.
//
// Solidity: function removeFromBlacklist(address account) returns()
func (_BridgeERC20 *BridgeERC20Session) RemoveFromBlacklist(account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.RemoveFromBlacklist(&_BridgeERC20.TransactOpts, account)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x537df3b6.
//
// Solidity: function removeFromBlacklist(address account) returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) RemoveFromBlacklist(account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.RemoveFromBlacklist(&_BridgeERC20.TransactOpts, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeERC20 *BridgeERC20Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeERC20 *BridgeERC20Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.RenounceRole(&_BridgeERC20.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.RenounceRole(&_BridgeERC20.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeERC20 *BridgeERC20Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeERC20 *BridgeERC20Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.RevokeRole(&_BridgeERC20.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeERC20.Contract.RevokeRole(&_BridgeERC20.TransactOpts, role, account)
}

// SetMiner is a paid mutator transaction binding the contract method 0x78b0bac4.
//
// Solidity: function setMiner(address miner, bool isActive) returns()
func (_BridgeERC20 *BridgeERC20Transactor) SetMiner(opts *bind.TransactOpts, miner common.Address, isActive bool) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "setMiner", miner, isActive)
}

// SetMiner is a paid mutator transaction binding the contract method 0x78b0bac4.
//
// Solidity: function setMiner(address miner, bool isActive) returns()
func (_BridgeERC20 *BridgeERC20Session) SetMiner(miner common.Address, isActive bool) (*types.Transaction, error) {
	return _BridgeERC20.Contract.SetMiner(&_BridgeERC20.TransactOpts, miner, isActive)
}

// SetMiner is a paid mutator transaction binding the contract method 0x78b0bac4.
//
// Solidity: function setMiner(address miner, bool isActive) returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) SetMiner(miner common.Address, isActive bool) (*types.Transaction, error) {
	return _BridgeERC20.Contract.SetMiner(&_BridgeERC20.TransactOpts, miner, isActive)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BridgeERC20 *BridgeERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BridgeERC20 *BridgeERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Transfer(&_BridgeERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BridgeERC20 *BridgeERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.Transfer(&_BridgeERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BridgeERC20 *BridgeERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BridgeERC20 *BridgeERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.TransferFrom(&_BridgeERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BridgeERC20 *BridgeERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeERC20.Contract.TransferFrom(&_BridgeERC20.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeERC20 *BridgeERC20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeERC20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeERC20 *BridgeERC20Session) Unpause() (*types.Transaction, error) {
	return _BridgeERC20.Contract.Unpause(&_BridgeERC20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeERC20 *BridgeERC20TransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeERC20.Contract.Unpause(&_BridgeERC20.TransactOpts)
}

// BridgeERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BridgeERC20 contract.
type BridgeERC20ApprovalIterator struct {
	Event *BridgeERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20Approval represents a Approval event raised by the BridgeERC20 contract.
type BridgeERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BridgeERC20 *BridgeERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BridgeERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20ApprovalIterator{contract: _BridgeERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BridgeERC20 *BridgeERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BridgeERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20Approval)
				if err := _BridgeERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BridgeERC20 *BridgeERC20Filterer) ParseApproval(log types.Log) (*BridgeERC20Approval, error) {
	event := new(BridgeERC20Approval)
	if err := _BridgeERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20BlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the BridgeERC20 contract.
type BridgeERC20BlacklistedIterator struct {
	Event *BridgeERC20Blacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20BlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20Blacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20Blacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20BlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20BlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20Blacklisted represents a Blacklisted event raised by the BridgeERC20 contract.
type BridgeERC20Blacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed account)
func (_BridgeERC20 *BridgeERC20Filterer) FilterBlacklisted(opts *bind.FilterOpts, account []common.Address) (*BridgeERC20BlacklistedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "Blacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20BlacklistedIterator{contract: _BridgeERC20.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed account)
func (_BridgeERC20 *BridgeERC20Filterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *BridgeERC20Blacklisted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "Blacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20Blacklisted)
				if err := _BridgeERC20.contract.UnpackLog(event, "Blacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklisted is a log parse operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed account)
func (_BridgeERC20 *BridgeERC20Filterer) ParseBlacklisted(log types.Log) (*BridgeERC20Blacklisted, error) {
	event := new(BridgeERC20Blacklisted)
	if err := _BridgeERC20.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20BurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the BridgeERC20 contract.
type BridgeERC20BurnedIterator struct {
	Event *BridgeERC20Burned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20BurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20Burned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20Burned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20BurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20BurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20Burned represents a Burned event raised by the BridgeERC20 contract.
type BridgeERC20Burned struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_BridgeERC20 *BridgeERC20Filterer) FilterBurned(opts *bind.FilterOpts, from []common.Address) (*BridgeERC20BurnedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20BurnedIterator{contract: _BridgeERC20.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_BridgeERC20 *BridgeERC20Filterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *BridgeERC20Burned, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20Burned)
				if err := _BridgeERC20.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurned is a log parse operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_BridgeERC20 *BridgeERC20Filterer) ParseBurned(log types.Log) (*BridgeERC20Burned, error) {
	event := new(BridgeERC20Burned)
	if err := _BridgeERC20.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BridgeERC20 contract.
type BridgeERC20EIP712DomainChangedIterator struct {
	Event *BridgeERC20EIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20EIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20EIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20EIP712DomainChanged represents a EIP712DomainChanged event raised by the BridgeERC20 contract.
type BridgeERC20EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgeERC20 *BridgeERC20Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BridgeERC20EIP712DomainChangedIterator, error) {

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeERC20EIP712DomainChangedIterator{contract: _BridgeERC20.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgeERC20 *BridgeERC20Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BridgeERC20EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20EIP712DomainChanged)
				if err := _BridgeERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgeERC20 *BridgeERC20Filterer) ParseEIP712DomainChanged(log types.Log) (*BridgeERC20EIP712DomainChanged, error) {
	event := new(BridgeERC20EIP712DomainChanged)
	if err := _BridgeERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the BridgeERC20 contract.
type BridgeERC20MintedIterator struct {
	Event *BridgeERC20Minted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20Minted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20Minted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20Minted represents a Minted event raised by the BridgeERC20 contract.
type BridgeERC20Minted struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_BridgeERC20 *BridgeERC20Filterer) FilterMinted(opts *bind.FilterOpts, to []common.Address) (*BridgeERC20MintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20MintedIterator{contract: _BridgeERC20.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_BridgeERC20 *BridgeERC20Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *BridgeERC20Minted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20Minted)
				if err := _BridgeERC20.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinted is a log parse operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_BridgeERC20 *BridgeERC20Filterer) ParseMinted(log types.Log) (*BridgeERC20Minted, error) {
	event := new(BridgeERC20Minted)
	if err := _BridgeERC20.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeERC20 contract.
type BridgeERC20PausedIterator struct {
	Event *BridgeERC20Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20Paused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20Paused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20Paused represents a Paused event raised by the BridgeERC20 contract.
type BridgeERC20Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeERC20 *BridgeERC20Filterer) FilterPaused(opts *bind.FilterOpts) (*BridgeERC20PausedIterator, error) {

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeERC20PausedIterator{contract: _BridgeERC20.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeERC20 *BridgeERC20Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeERC20Paused) (event.Subscription, error) {

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20Paused)
				if err := _BridgeERC20.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeERC20 *BridgeERC20Filterer) ParsePaused(log types.Log) (*BridgeERC20Paused, error) {
	event := new(BridgeERC20Paused)
	if err := _BridgeERC20.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeERC20 contract.
type BridgeERC20RoleAdminChangedIterator struct {
	Event *BridgeERC20RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20RoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20RoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20RoleAdminChanged represents a RoleAdminChanged event raised by the BridgeERC20 contract.
type BridgeERC20RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeERC20 *BridgeERC20Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeERC20RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20RoleAdminChangedIterator{contract: _BridgeERC20.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeERC20 *BridgeERC20Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeERC20RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20RoleAdminChanged)
				if err := _BridgeERC20.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeERC20 *BridgeERC20Filterer) ParseRoleAdminChanged(log types.Log) (*BridgeERC20RoleAdminChanged, error) {
	event := new(BridgeERC20RoleAdminChanged)
	if err := _BridgeERC20.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeERC20 contract.
type BridgeERC20RoleGrantedIterator struct {
	Event *BridgeERC20RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20RoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20RoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20RoleGranted represents a RoleGranted event raised by the BridgeERC20 contract.
type BridgeERC20RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeERC20 *BridgeERC20Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeERC20RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20RoleGrantedIterator{contract: _BridgeERC20.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeERC20 *BridgeERC20Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeERC20RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20RoleGranted)
				if err := _BridgeERC20.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeERC20 *BridgeERC20Filterer) ParseRoleGranted(log types.Log) (*BridgeERC20RoleGranted, error) {
	event := new(BridgeERC20RoleGranted)
	if err := _BridgeERC20.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeERC20 contract.
type BridgeERC20RoleRevokedIterator struct {
	Event *BridgeERC20RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20RoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20RoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20RoleRevoked represents a RoleRevoked event raised by the BridgeERC20 contract.
type BridgeERC20RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeERC20 *BridgeERC20Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeERC20RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20RoleRevokedIterator{contract: _BridgeERC20.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeERC20 *BridgeERC20Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeERC20RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20RoleRevoked)
				if err := _BridgeERC20.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeERC20 *BridgeERC20Filterer) ParseRoleRevoked(log types.Log) (*BridgeERC20RoleRevoked, error) {
	event := new(BridgeERC20RoleRevoked)
	if err := _BridgeERC20.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BridgeERC20 contract.
type BridgeERC20TransferIterator struct {
	Event *BridgeERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20Transfer represents a Transfer event raised by the BridgeERC20 contract.
type BridgeERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BridgeERC20 *BridgeERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BridgeERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20TransferIterator{contract: _BridgeERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BridgeERC20 *BridgeERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BridgeERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20Transfer)
				if err := _BridgeERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BridgeERC20 *BridgeERC20Filterer) ParseTransfer(log types.Log) (*BridgeERC20Transfer, error) {
	event := new(BridgeERC20Transfer)
	if err := _BridgeERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20UnBlacklistedIterator is returned from FilterUnBlacklisted and is used to iterate over the raw logs and unpacked data for UnBlacklisted events raised by the BridgeERC20 contract.
type BridgeERC20UnBlacklistedIterator struct {
	Event *BridgeERC20UnBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20UnBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20UnBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20UnBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20UnBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20UnBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20UnBlacklisted represents a UnBlacklisted event raised by the BridgeERC20 contract.
type BridgeERC20UnBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnBlacklisted is a free log retrieval operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed account)
func (_BridgeERC20 *BridgeERC20Filterer) FilterUnBlacklisted(opts *bind.FilterOpts, account []common.Address) (*BridgeERC20UnBlacklistedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "UnBlacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20UnBlacklistedIterator{contract: _BridgeERC20.contract, event: "UnBlacklisted", logs: logs, sub: sub}, nil
}

// WatchUnBlacklisted is a free log subscription operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed account)
func (_BridgeERC20 *BridgeERC20Filterer) WatchUnBlacklisted(opts *bind.WatchOpts, sink chan<- *BridgeERC20UnBlacklisted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "UnBlacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20UnBlacklisted)
				if err := _BridgeERC20.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnBlacklisted is a log parse operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed account)
func (_BridgeERC20 *BridgeERC20Filterer) ParseUnBlacklisted(log types.Log) (*BridgeERC20UnBlacklisted, error) {
	event := new(BridgeERC20UnBlacklisted)
	if err := _BridgeERC20.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeERC20UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeERC20 contract.
type BridgeERC20UnpausedIterator struct {
	Event *BridgeERC20Unpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20Unpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20Unpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20Unpaused represents a Unpaused event raised by the BridgeERC20 contract.
type BridgeERC20Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeERC20 *BridgeERC20Filterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeERC20UnpausedIterator, error) {

	logs, sub, err := _BridgeERC20.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeERC20UnpausedIterator{contract: _BridgeERC20.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeERC20 *BridgeERC20Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeERC20Unpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeERC20.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20Unpaused)
				if err := _BridgeERC20.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeERC20 *BridgeERC20Filterer) ParseUnpaused(log types.Log) (*BridgeERC20Unpaused, error) {
	event := new(BridgeERC20Unpaused)
	if err := _BridgeERC20.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
