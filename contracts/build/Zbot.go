// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zbot

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ZbotABI is the input ABI used to generate the binding from.
const ZbotABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Tip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"tip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ZbotBin is the compiled bytecode used for deploying new contracts.
var ZbotBin = "0x60806040523480156200001157600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060400160405280600581526020017f312e322e30000000000000000000000000000000000000000000000000000000815250600490805190602001906200009f929190620000a6565b506200015c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620000de57600085556200012a565b82601f10620000f957805160ff19168380011785556200012a565b828001600101855582156200012a579182015b82811115620001295782518255916020019190600101906200010c565b5b5090506200013991906200013d565b5090565b5b80821115620001585760008160009055506001016200013e565b5090565b61176f806200016c6000396000f3fe6080604052600436106100745760003560e01c80633d80b5321161004e5780633d80b532146105ce57806354fd4d50146106aa578063a26e11861461073a578063bf40fac1146107f5576102b0565b806320cebbf6146102b557806331fb67c21461041e57806332434a2e146104e6576102b0565b366102b0576060600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561014e5780601f106101235761010080835404028352916020019161014e565b820191906000526020600020905b81548152906001019060200180831161013157829003601f168201915b505050505090506000815114156101cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f4e6f207573657220494420666f756e6420666f7220616464726573732e00000081525060200191505060405180910390fd5b6102ad600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102a35780601f10610278576101008083540402835291602001916102a3565b820191906000526020600020905b81548152906001019060200180831161028657829003601f168201915b50505050506108e7565b50005b600080fd5b3480156102c157600080fd5b5061041c600480360360608110156102d857600080fd5b81019080803590602001906401000000008111156102f557600080fd5b82018360208201111561030757600080fd5b8035906020019184600183028401116401000000008311171561032957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561038c57600080fd5b82018360208201111561039e57600080fd5b803590602001918460018302840111640100000000831117156103c057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610b14565b005b34801561042a57600080fd5b506104e46004803603602081101561044157600080fd5b810190808035906020019064010000000081111561045e57600080fd5b82018360208201111561047057600080fd5b8035906020019184600183028401116401000000008311171561049257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610eb8565b005b3480156104f257600080fd5b506105cc6004803603604081101561050957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561054657600080fd5b82018360208201111561055857600080fd5b8035906020019184600183028401116401000000008311171561057a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061124c565b005b3480156105da57600080fd5b50610694600480360360208110156105f157600080fd5b810190808035906020019064010000000081111561060e57600080fd5b82018360208201111561062057600080fd5b8035906020019184600183028401116401000000008311171561064257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506114c7565b6040518082815260200191505060405180910390f35b3480156106b657600080fd5b506106bf61153a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106ff5780820151818401526020810190506106e4565b50505050905090810190601f16801561072c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6107f36004803603602081101561075057600080fd5b810190808035906020019064010000000081111561076d57600080fd5b82018360208201111561077f57600080fd5b803590602001918460018302840111640100000000831117156107a157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506108e7565b005b34801561080157600080fd5b506108bb6004803603602081101561081857600080fd5b810190808035906020019064010000000081111561083557600080fd5b82018360208201111561084757600080fd5b8035906020019184600183028401116401000000008311171561086957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506115d8565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600073ffffffffffffffffffffffffffffffffffffffff166001826040518082805190602001908083835b602083106109355780518252602082019150602081019050602083039250610912565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156109f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806117176023913960400191505060405180910390fd5b346002826040518082805190602001908083835b60208310610a2d5780518252602082019150602081019050602083039250610a0a565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600082825401925050819055507f135290006f1871309577d248cf00619459a79b4a735638108a2bb080319c90d681346040518080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610ad6578082015181840152602081019050610abb565b50505050905090810190601f168015610b035780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bd5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f596f7520617265206e6f7420746865206f776e65722e0000000000000000000081525060200191505060405180910390fd5b806002846040518082805190602001908083835b60208310610c0c5780518252602082019150602081019050602083039250610be9565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020541015610cb4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f4e6f7420656e6f7567682062616c616e63652e0000000000000000000000000081525060200191505060405180910390fd5b806002846040518082805190602001908083835b60208310610ceb5780518252602082019150602081019050602083039250610cc8565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008282540392505081905550806002836040518082805190602001908083835b60208310610d625780518252602082019150602081019050602083039250610d3f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600082825401925050819055507fc58663f31c3549520494f7c83d9787b7bf8bb4878e12f0bf3ce44e272626a0bd838383604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015610e10578082015181840152602081019050610df5565b50505050905090810190601f168015610e3d5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b83811015610e76578082015181840152602081019050610e5b565b50505050905090810190601f168015610ea35780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f79576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f596f7520617265206e6f7420746865206f776e65722e0000000000000000000081525060200191505060405180910390fd5b60006001826040518082805190602001908083835b60208310610fb15780518252602082019150602081019050602083039250610f8e565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561108d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806117176023913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc6002846040518082805190602001908083835b602083106110dd57805182526020820191506020810190506020830392506110ba565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549081150290604051600060405180830381858888f1935050505015801561113c573d6000803e3d6000fd5b5060006002836040518082805190602001908083835b602083106111755780518252602082019150602081019050602083039250611152565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020819055507fcb36594c54e38123fa4ff1ce9db90c80bc294a2101aeea2c7b172e5a9043a69c826040518080602001828103825283818151815260200191508051906020019080838360005b8381101561120e5780820151818401526020810190506111f3565b50505050905090810190601f16801561123b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461130d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f596f7520617265206e6f7420746865206f776e65722e0000000000000000000081525060200191505060405180910390fd5b816001826040518082805190602001908083835b602083106113445780518252602082019150602081019050602083039250611321565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020908051906020019061140892919061166b565b507f6ba0831d2f62ae5cbf7214bcc1d79c5da1d705f12811efda0beaa840006f874e8282604051808373ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561148857808201518184015260208101905061146d565b50505050905090810190601f1680156114b55780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b60006002826040518082805190602001908083835b602083106114ff57805182526020820191506020810190506020830392506114dc565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115d05780601f106115a5576101008083540402835291602001916115d0565b820191906000526020600020905b8154815290600101906020018083116115b357829003601f168201915b505050505081565b60006001826040518082805190602001908083835b6020831061161057805182526020820191506020810190506020830392506115ed565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826116a157600085556116e8565b82601f106116ba57805160ff19168380011785556116e8565b828001600101855582156116e8579182015b828111156116e75782518255916020019190600101906116cc565b5b5090506116f591906116f9565b5090565b5b808211156117125760008160009055506001016116fa565b509056fe4e6f206164647265737320726567697374657265642077697468206163636f756e742ea26469706673582212209874b73ca38386e3454b2c2562d1b9a024a6e6c5b58ac5c29d5b1ffcf56d080e64736f6c63430007050033"

// DeployZbot deploys a new Ethereum contract, binding an instance of Zbot to it.
func DeployZbot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Zbot, error) {
	parsed, err := abi.JSON(strings.NewReader(ZbotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZbotBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zbot{ZbotCaller: ZbotCaller{contract: contract}, ZbotTransactor: ZbotTransactor{contract: contract}, ZbotFilterer: ZbotFilterer{contract: contract}}, nil
}

// Zbot is an auto generated Go binding around an Ethereum contract.
type Zbot struct {
	ZbotCaller     // Read-only binding to the contract
	ZbotTransactor // Write-only binding to the contract
	ZbotFilterer   // Log filterer for contract events
}

// ZbotCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZbotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZbotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZbotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZbotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZbotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZbotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZbotSession struct {
	Contract     *Zbot             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZbotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZbotCallerSession struct {
	Contract *ZbotCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZbotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZbotTransactorSession struct {
	Contract     *ZbotTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZbotRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZbotRaw struct {
	Contract *Zbot // Generic contract binding to access the raw methods on
}

// ZbotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZbotCallerRaw struct {
	Contract *ZbotCaller // Generic read-only contract binding to access the raw methods on
}

// ZbotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZbotTransactorRaw struct {
	Contract *ZbotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZbot creates a new instance of Zbot, bound to a specific deployed contract.
func NewZbot(address common.Address, backend bind.ContractBackend) (*Zbot, error) {
	contract, err := bindZbot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zbot{ZbotCaller: ZbotCaller{contract: contract}, ZbotTransactor: ZbotTransactor{contract: contract}, ZbotFilterer: ZbotFilterer{contract: contract}}, nil
}

// NewZbotCaller creates a new read-only instance of Zbot, bound to a specific deployed contract.
func NewZbotCaller(address common.Address, caller bind.ContractCaller) (*ZbotCaller, error) {
	contract, err := bindZbot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZbotCaller{contract: contract}, nil
}

// NewZbotTransactor creates a new write-only instance of Zbot, bound to a specific deployed contract.
func NewZbotTransactor(address common.Address, transactor bind.ContractTransactor) (*ZbotTransactor, error) {
	contract, err := bindZbot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZbotTransactor{contract: contract}, nil
}

// NewZbotFilterer creates a new log filterer instance of Zbot, bound to a specific deployed contract.
func NewZbotFilterer(address common.Address, filterer bind.ContractFilterer) (*ZbotFilterer, error) {
	contract, err := bindZbot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZbotFilterer{contract: contract}, nil
}

// bindZbot binds a generic wrapper to an already deployed contract.
func bindZbot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZbotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zbot *ZbotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zbot.Contract.ZbotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zbot *ZbotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zbot.Contract.ZbotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zbot *ZbotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zbot.Contract.ZbotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zbot *ZbotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zbot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zbot *ZbotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zbot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zbot *ZbotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zbot.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x3d80b532.
//
// Solidity: function balance(string id) view returns(uint256)
func (_Zbot *ZbotCaller) Balance(opts *bind.CallOpts, id string) (*big.Int, error) {
	var out []interface{}
	err := _Zbot.contract.Call(opts, &out, "balance", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x3d80b532.
//
// Solidity: function balance(string id) view returns(uint256)
func (_Zbot *ZbotSession) Balance(id string) (*big.Int, error) {
	return _Zbot.Contract.Balance(&_Zbot.CallOpts, id)
}

// Balance is a free data retrieval call binding the contract method 0x3d80b532.
//
// Solidity: function balance(string id) view returns(uint256)
func (_Zbot *ZbotCallerSession) Balance(id string) (*big.Int, error) {
	return _Zbot.Contract.Balance(&_Zbot.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string id) view returns(address)
func (_Zbot *ZbotCaller) GetAddress(opts *bind.CallOpts, id string) (common.Address, error) {
	var out []interface{}
	err := _Zbot.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string id) view returns(address)
func (_Zbot *ZbotSession) GetAddress(id string) (common.Address, error) {
	return _Zbot.Contract.GetAddress(&_Zbot.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string id) view returns(address)
func (_Zbot *ZbotCallerSession) GetAddress(id string) (common.Address, error) {
	return _Zbot.Contract.GetAddress(&_Zbot.CallOpts, id)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Zbot *ZbotCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zbot.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Zbot *ZbotSession) Version() (string, error) {
	return _Zbot.Contract.Version(&_Zbot.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Zbot *ZbotCallerSession) Version() (string, error) {
	return _Zbot.Contract.Version(&_Zbot.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string id) payable returns()
func (_Zbot *ZbotTransactor) Deposit(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _Zbot.contract.Transact(opts, "deposit", id)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string id) payable returns()
func (_Zbot *ZbotSession) Deposit(id string) (*types.Transaction, error) {
	return _Zbot.Contract.Deposit(&_Zbot.TransactOpts, id)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string id) payable returns()
func (_Zbot *ZbotTransactorSession) Deposit(id string) (*types.Transaction, error) {
	return _Zbot.Contract.Deposit(&_Zbot.TransactOpts, id)
}

// Register is a paid mutator transaction binding the contract method 0x32434a2e.
//
// Solidity: function register(address addr, string id) returns()
func (_Zbot *ZbotTransactor) Register(opts *bind.TransactOpts, addr common.Address, id string) (*types.Transaction, error) {
	return _Zbot.contract.Transact(opts, "register", addr, id)
}

// Register is a paid mutator transaction binding the contract method 0x32434a2e.
//
// Solidity: function register(address addr, string id) returns()
func (_Zbot *ZbotSession) Register(addr common.Address, id string) (*types.Transaction, error) {
	return _Zbot.Contract.Register(&_Zbot.TransactOpts, addr, id)
}

// Register is a paid mutator transaction binding the contract method 0x32434a2e.
//
// Solidity: function register(address addr, string id) returns()
func (_Zbot *ZbotTransactorSession) Register(addr common.Address, id string) (*types.Transaction, error) {
	return _Zbot.Contract.Register(&_Zbot.TransactOpts, addr, id)
}

// Tip is a paid mutator transaction binding the contract method 0x20cebbf6.
//
// Solidity: function tip(string from, string to, uint256 amount) returns()
func (_Zbot *ZbotTransactor) Tip(opts *bind.TransactOpts, from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Zbot.contract.Transact(opts, "tip", from, to, amount)
}

// Tip is a paid mutator transaction binding the contract method 0x20cebbf6.
//
// Solidity: function tip(string from, string to, uint256 amount) returns()
func (_Zbot *ZbotSession) Tip(from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Zbot.Contract.Tip(&_Zbot.TransactOpts, from, to, amount)
}

// Tip is a paid mutator transaction binding the contract method 0x20cebbf6.
//
// Solidity: function tip(string from, string to, uint256 amount) returns()
func (_Zbot *ZbotTransactorSession) Tip(from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Zbot.Contract.Tip(&_Zbot.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string id) returns()
func (_Zbot *ZbotTransactor) Withdraw(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _Zbot.contract.Transact(opts, "withdraw", id)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string id) returns()
func (_Zbot *ZbotSession) Withdraw(id string) (*types.Transaction, error) {
	return _Zbot.Contract.Withdraw(&_Zbot.TransactOpts, id)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string id) returns()
func (_Zbot *ZbotTransactorSession) Withdraw(id string) (*types.Transaction, error) {
	return _Zbot.Contract.Withdraw(&_Zbot.TransactOpts, id)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Zbot *ZbotTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zbot.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Zbot *ZbotSession) Receive() (*types.Transaction, error) {
	return _Zbot.Contract.Receive(&_Zbot.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Zbot *ZbotTransactorSession) Receive() (*types.Transaction, error) {
	return _Zbot.Contract.Receive(&_Zbot.TransactOpts)
}

// ZbotDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Zbot contract.
type ZbotDepositIterator struct {
	Event *ZbotDeposit // Event containing the contract specifics and raw log

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
func (it *ZbotDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZbotDeposit)
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
		it.Event = new(ZbotDeposit)
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
func (it *ZbotDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZbotDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZbotDeposit represents a Deposit event raised by the Zbot contract.
type ZbotDeposit struct {
	Id     string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x135290006f1871309577d248cf00619459a79b4a735638108a2bb080319c90d6.
//
// Solidity: event Deposit(string id, uint256 amount)
func (_Zbot *ZbotFilterer) FilterDeposit(opts *bind.FilterOpts) (*ZbotDepositIterator, error) {

	logs, sub, err := _Zbot.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ZbotDepositIterator{contract: _Zbot.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x135290006f1871309577d248cf00619459a79b4a735638108a2bb080319c90d6.
//
// Solidity: event Deposit(string id, uint256 amount)
func (_Zbot *ZbotFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZbotDeposit) (event.Subscription, error) {

	logs, sub, err := _Zbot.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZbotDeposit)
				if err := _Zbot.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x135290006f1871309577d248cf00619459a79b4a735638108a2bb080319c90d6.
//
// Solidity: event Deposit(string id, uint256 amount)
func (_Zbot *ZbotFilterer) ParseDeposit(log types.Log) (*ZbotDeposit, error) {
	event := new(ZbotDeposit)
	if err := _Zbot.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZbotRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Zbot contract.
type ZbotRegisterIterator struct {
	Event *ZbotRegister // Event containing the contract specifics and raw log

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
func (it *ZbotRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZbotRegister)
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
		it.Event = new(ZbotRegister)
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
func (it *ZbotRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZbotRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZbotRegister represents a Register event raised by the Zbot contract.
type ZbotRegister struct {
	Addr common.Address
	Id   string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x6ba0831d2f62ae5cbf7214bcc1d79c5da1d705f12811efda0beaa840006f874e.
//
// Solidity: event Register(address addr, string id)
func (_Zbot *ZbotFilterer) FilterRegister(opts *bind.FilterOpts) (*ZbotRegisterIterator, error) {

	logs, sub, err := _Zbot.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &ZbotRegisterIterator{contract: _Zbot.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x6ba0831d2f62ae5cbf7214bcc1d79c5da1d705f12811efda0beaa840006f874e.
//
// Solidity: event Register(address addr, string id)
func (_Zbot *ZbotFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *ZbotRegister) (event.Subscription, error) {

	logs, sub, err := _Zbot.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZbotRegister)
				if err := _Zbot.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0x6ba0831d2f62ae5cbf7214bcc1d79c5da1d705f12811efda0beaa840006f874e.
//
// Solidity: event Register(address addr, string id)
func (_Zbot *ZbotFilterer) ParseRegister(log types.Log) (*ZbotRegister, error) {
	event := new(ZbotRegister)
	if err := _Zbot.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZbotTipIterator is returned from FilterTip and is used to iterate over the raw logs and unpacked data for Tip events raised by the Zbot contract.
type ZbotTipIterator struct {
	Event *ZbotTip // Event containing the contract specifics and raw log

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
func (it *ZbotTipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZbotTip)
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
		it.Event = new(ZbotTip)
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
func (it *ZbotTipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZbotTipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZbotTip represents a Tip event raised by the Zbot contract.
type ZbotTip struct {
	From   string
	To     string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTip is a free log retrieval operation binding the contract event 0xc58663f31c3549520494f7c83d9787b7bf8bb4878e12f0bf3ce44e272626a0bd.
//
// Solidity: event Tip(string from, string to, uint256 amount)
func (_Zbot *ZbotFilterer) FilterTip(opts *bind.FilterOpts) (*ZbotTipIterator, error) {

	logs, sub, err := _Zbot.contract.FilterLogs(opts, "Tip")
	if err != nil {
		return nil, err
	}
	return &ZbotTipIterator{contract: _Zbot.contract, event: "Tip", logs: logs, sub: sub}, nil
}

// WatchTip is a free log subscription operation binding the contract event 0xc58663f31c3549520494f7c83d9787b7bf8bb4878e12f0bf3ce44e272626a0bd.
//
// Solidity: event Tip(string from, string to, uint256 amount)
func (_Zbot *ZbotFilterer) WatchTip(opts *bind.WatchOpts, sink chan<- *ZbotTip) (event.Subscription, error) {

	logs, sub, err := _Zbot.contract.WatchLogs(opts, "Tip")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZbotTip)
				if err := _Zbot.contract.UnpackLog(event, "Tip", log); err != nil {
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

// ParseTip is a log parse operation binding the contract event 0xc58663f31c3549520494f7c83d9787b7bf8bb4878e12f0bf3ce44e272626a0bd.
//
// Solidity: event Tip(string from, string to, uint256 amount)
func (_Zbot *ZbotFilterer) ParseTip(log types.Log) (*ZbotTip, error) {
	event := new(ZbotTip)
	if err := _Zbot.contract.UnpackLog(event, "Tip", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZbotWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Zbot contract.
type ZbotWithdrawIterator struct {
	Event *ZbotWithdraw // Event containing the contract specifics and raw log

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
func (it *ZbotWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZbotWithdraw)
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
		it.Event = new(ZbotWithdraw)
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
func (it *ZbotWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZbotWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZbotWithdraw represents a Withdraw event raised by the Zbot contract.
type ZbotWithdraw struct {
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xcb36594c54e38123fa4ff1ce9db90c80bc294a2101aeea2c7b172e5a9043a69c.
//
// Solidity: event Withdraw(string id)
func (_Zbot *ZbotFilterer) FilterWithdraw(opts *bind.FilterOpts) (*ZbotWithdrawIterator, error) {

	logs, sub, err := _Zbot.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &ZbotWithdrawIterator{contract: _Zbot.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xcb36594c54e38123fa4ff1ce9db90c80bc294a2101aeea2c7b172e5a9043a69c.
//
// Solidity: event Withdraw(string id)
func (_Zbot *ZbotFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZbotWithdraw) (event.Subscription, error) {

	logs, sub, err := _Zbot.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZbotWithdraw)
				if err := _Zbot.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xcb36594c54e38123fa4ff1ce9db90c80bc294a2101aeea2c7b172e5a9043a69c.
//
// Solidity: event Withdraw(string id)
func (_Zbot *ZbotFilterer) ParseWithdraw(log types.Log) (*ZbotWithdraw, error) {
	event := new(ZbotWithdraw)
	if err := _Zbot.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
