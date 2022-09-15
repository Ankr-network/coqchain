// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staker

import (
	"errors"
	"math/big"
	"strings"

	coqchain "github.com/Ankr-network/coqchain"
	"github.com/Ankr-network/coqchain/accounts/abi"
	"github.com/Ankr-network/coqchain/accounts/abi/bind"
	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/core/types"
	"github.com/Ankr-network/coqchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = coqchain.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakerMetaData contains all meta data concerning the Staker contract.
var StakerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sigs\",\"type\":\"address[]\"}],\"name\":\"commitSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526110ef806100136000396000f3fe6080604052600436106100555760003560e01c806327e235e31461005a5780633cf4ac6d1461009757806370a08231146100c057806394cf795e146100fd578063d0e30db014610128578063f3fef3a314610132575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c9190610991565b61015b565b60405161008e91906109d7565b60405180910390f35b3480156100a357600080fd5b506100be60048036038101906100b99190610b4b565b610173565b005b3480156100cc57600080fd5b506100e760048036038101906100e29190610991565b6104c0565b6040516100f491906109d7565b60405180910390f35b34801561010957600080fd5b50610112610509565b60405161011f9190610c52565b60405180910390f35b610130610597565b005b34801561013e57600080fd5b5061015960048036038101906101549190610cde565b610632565b005b60026020528060005260406000206000915090505481565b61017c336107ad565b6101bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b290610d7b565b60405180910390fd5b600160006101c99190610857565b60005b60008054905081101561028d576001600082815481106101ef576101ee610d9b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061028590610df9565b9150506101cc565b5080600090805190602001906102a4929190610878565b5060005b6001805490508110156104bc576102fc600182815481106102cc576102cb610d9b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166107ad565b15801561038657506000600260006001848154811061031e5761031d610d9b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054115b156104a957600a60026000600184815481106103a5576103a4610d9b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546104159190610e70565b600260006001848154811061042d5761042c610d9b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104a19190610ea1565b925050819055505b80806104b490610df9565b9150506102a8565b5050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600080548060200260200160405190810160405280929190818152602001828054801561058d57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610543575b5050505050905090565b600034036105da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d190610f21565b60405180910390fd5b34600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106299190610f41565b92505081905550565b61063b336107ad565b61067a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067190610d7b565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156106fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f390610fe7565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405161072290611038565b60006040518083038185875af1925050503d806000811461075f576040519150601f19603f3d011682016040523d82523d6000602084013e610764565b606091505b50509050806107a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079f90611099565b60405180910390fd5b505050565b6000806000905060005b60008054905081101561084d578373ffffffffffffffffffffffffffffffffffffffff16600082815481106107ef576107ee610d9b565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361083a57600191505b808061084590610df9565b9150506107b7565b5080915050919050565b50805460008255906000526020600020908101906108759190610902565b50565b8280548282559060005260206000209081019282156108f1579160200282015b828111156108f05782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610898565b5b5090506108fe9190610902565b5090565b5b8082111561091b576000816000905550600101610903565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061095e82610933565b9050919050565b61096e81610953565b811461097957600080fd5b50565b60008135905061098b81610965565b92915050565b6000602082840312156109a7576109a6610929565b5b60006109b58482850161097c565b91505092915050565b6000819050919050565b6109d1816109be565b82525050565b60006020820190506109ec60008301846109c8565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610a40826109f7565b810181811067ffffffffffffffff82111715610a5f57610a5e610a08565b5b80604052505050565b6000610a7261091f565b9050610a7e8282610a37565b919050565b600067ffffffffffffffff821115610a9e57610a9d610a08565b5b602082029050602081019050919050565b600080fd5b6000610ac7610ac284610a83565b610a68565b90508083825260208201905060208402830185811115610aea57610ae9610aaf565b5b835b81811015610b135780610aff888261097c565b845260208401935050602081019050610aec565b5050509392505050565b600082601f830112610b3257610b316109f2565b5b8135610b42848260208601610ab4565b91505092915050565b600060208284031215610b6157610b60610929565b5b600082013567ffffffffffffffff811115610b7f57610b7e61092e565b5b610b8b84828501610b1d565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610bc981610953565b82525050565b6000610bdb8383610bc0565b60208301905092915050565b6000602082019050919050565b6000610bff82610b94565b610c098185610b9f565b9350610c1483610bb0565b8060005b83811015610c45578151610c2c8882610bcf565b9750610c3783610be7565b925050600181019050610c18565b5085935050505092915050565b60006020820190508181036000830152610c6c8184610bf4565b905092915050565b6000610c7f82610933565b9050919050565b610c8f81610c74565b8114610c9a57600080fd5b50565b600081359050610cac81610c86565b92915050565b610cbb816109be565b8114610cc657600080fd5b50565b600081359050610cd881610cb2565b92915050565b60008060408385031215610cf557610cf4610929565b5b6000610d0385828601610c9d565b9250506020610d1485828601610cc9565b9150509250929050565b600082825260208201905092915050565b7f6e6f74207369676e657200000000000000000000000000000000000000000000600082015250565b6000610d65600a83610d1e565b9150610d7082610d2f565b602082019050919050565b60006020820190508181036000830152610d9481610d58565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e04826109be565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e3657610e35610dca565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610e7b826109be565b9150610e86836109be565b925082610e9657610e95610e41565b5b828204905092915050565b6000610eac826109be565b9150610eb7836109be565b9250828203905081811115610ecf57610ece610dca565b5b92915050565b7f77726f6e672076616c7565000000000000000000000000000000000000000000600082015250565b6000610f0b600b83610d1e565b9150610f1682610ed5565b602082019050919050565b60006020820190508181036000830152610f3a81610efe565b9050919050565b6000610f4c826109be565b9150610f57836109be565b9250828201905080821115610f6f57610f6e610dca565b5b92915050565b7f776974686472617720616d6f756e742073686f756c64206265206c657373206f60008201527f7220657175616c20796f75722062616c616e6365000000000000000000000000602082015250565b6000610fd1603483610d1e565b9150610fdc82610f75565b604082019050919050565b6000602082019050818103600083015261100081610fc4565b9050919050565b600081905092915050565b50565b6000611022600083611007565b915061102d82611012565b600082019050919050565b600061104382611015565b9150819050919050565b7f4661696c656420746f2073656e64000000000000000000000000000000000000600082015250565b6000611083600e83610d1e565b915061108e8261104d565b602082019050919050565b600060208201905081810360008301526110b281611076565b905091905056fea264697066735822122008882fcb5278a1659ae3b442498741512b2fca26d5d3ad774ddce21883a680f464736f6c63430008100033",
}

// StakerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakerMetaData.ABI instead.
var StakerABI = StakerMetaData.ABI

// StakerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakerMetaData.Bin instead.
var StakerBin = StakerMetaData.Bin

// DeployStaker deploys a new coqchain contract, binding an instance of Staker to it.
func DeployStaker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Staker, error) {
	parsed, err := StakerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staker{StakerCaller: StakerCaller{contract: contract}, StakerTransactor: StakerTransactor{contract: contract}, StakerFilterer: StakerFilterer{contract: contract}}, nil
}

// Staker is an auto generated Go binding around an coqchain contract.
type Staker struct {
	StakerCaller     // Read-only binding to the contract
	StakerTransactor // Write-only binding to the contract
	StakerFilterer   // Log filterer for contract events
}

// StakerCaller is an auto generated read-only Go binding around an coqchain contract.
type StakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerTransactor is an auto generated write-only Go binding around an coqchain contract.
type StakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerFilterer is an auto generated log filtering Go binding around an coqchain contract events.
type StakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerSession is an auto generated Go binding around an coqchain contract,
// with pre-set call and transact options.
type StakerSession struct {
	Contract     *Staker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakerCallerSession is an auto generated read-only Go binding around an coqchain contract,
// with pre-set call options.
type StakerCallerSession struct {
	Contract *StakerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakerTransactorSession is an auto generated write-only Go binding around an coqchain contract,
// with pre-set transact options.
type StakerTransactorSession struct {
	Contract     *StakerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakerRaw is an auto generated low-level Go binding around an coqchain contract.
type StakerRaw struct {
	Contract *Staker // Generic contract binding to access the raw methods on
}

// StakerCallerRaw is an auto generated low-level read-only Go binding around an coqchain contract.
type StakerCallerRaw struct {
	Contract *StakerCaller // Generic read-only contract binding to access the raw methods on
}

// StakerTransactorRaw is an auto generated low-level write-only Go binding around an coqchain contract.
type StakerTransactorRaw struct {
	Contract *StakerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaker creates a new instance of Staker, bound to a specific deployed contract.
func NewStaker(address common.Address, backend bind.ContractBackend) (*Staker, error) {
	contract, err := bindStaker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staker{StakerCaller: StakerCaller{contract: contract}, StakerTransactor: StakerTransactor{contract: contract}, StakerFilterer: StakerFilterer{contract: contract}}, nil
}

// NewStakerCaller creates a new read-only instance of Staker, bound to a specific deployed contract.
func NewStakerCaller(address common.Address, caller bind.ContractCaller) (*StakerCaller, error) {
	contract, err := bindStaker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakerCaller{contract: contract}, nil
}

// NewStakerTransactor creates a new write-only instance of Staker, bound to a specific deployed contract.
func NewStakerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakerTransactor, error) {
	contract, err := bindStaker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakerTransactor{contract: contract}, nil
}

// NewStakerFilterer creates a new log filterer instance of Staker, bound to a specific deployed contract.
func NewStakerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakerFilterer, error) {
	contract, err := bindStaker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakerFilterer{contract: contract}, nil
}

// bindStaker binds a generic wrapper to an already deployed contract.
func bindStaker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staker *StakerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staker.Contract.StakerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staker *StakerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.Contract.StakerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staker *StakerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staker.Contract.StakerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staker *StakerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staker *StakerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staker *StakerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staker.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Staker *StakerCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Staker *StakerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Staker.Contract.BalanceOf(&_Staker.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Staker *StakerCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Staker.Contract.BalanceOf(&_Staker.CallOpts, addr)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Staker *StakerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Staker *StakerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Staker.Contract.Balances(&_Staker.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Staker *StakerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Staker.Contract.Balances(&_Staker.CallOpts, arg0)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Staker *StakerCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Staker *StakerSession) GetSigners() ([]common.Address, error) {
	return _Staker.Contract.GetSigners(&_Staker.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Staker *StakerCallerSession) GetSigners() ([]common.Address, error) {
	return _Staker.Contract.GetSigners(&_Staker.CallOpts)
}

// CommitSigners is a paid mutator transaction binding the contract method 0x3cf4ac6d.
//
// Solidity: function commitSigners(address[] sigs) returns()
func (_Staker *StakerTransactor) CommitSigners(opts *bind.TransactOpts, sigs []common.Address) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "commitSigners", sigs)
}

// CommitSigners is a paid mutator transaction binding the contract method 0x3cf4ac6d.
//
// Solidity: function commitSigners(address[] sigs) returns()
func (_Staker *StakerSession) CommitSigners(sigs []common.Address) (*types.Transaction, error) {
	return _Staker.Contract.CommitSigners(&_Staker.TransactOpts, sigs)
}

// CommitSigners is a paid mutator transaction binding the contract method 0x3cf4ac6d.
//
// Solidity: function commitSigners(address[] sigs) returns()
func (_Staker *StakerTransactorSession) CommitSigners(sigs []common.Address) (*types.Transaction, error) {
	return _Staker.Contract.CommitSigners(&_Staker.TransactOpts, sigs)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Staker *StakerTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Staker *StakerSession) Deposit() (*types.Transaction, error) {
	return _Staker.Contract.Deposit(&_Staker.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Staker *StakerTransactorSession) Deposit() (*types.Transaction, error) {
	return _Staker.Contract.Deposit(&_Staker.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_Staker *StakerTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "withdraw", _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_Staker *StakerSession) Withdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.Withdraw(&_Staker.TransactOpts, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_Staker *StakerTransactorSession) Withdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.Withdraw(&_Staker.TransactOpts, _to, _amount)
}
