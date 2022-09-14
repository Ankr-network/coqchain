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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sigs\",\"type\":\"address[]\"}],\"name\":\"commitSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526064600055606460008190555061103d806100206000396000f3fe60806040526004361061007b5760003560e01c8063757991a81161004e578063757991a81461014c57806394cf795e14610177578063d0e30db0146101a2578063f3fef3a3146101ac5761007b565b806302fb4d851461008057806327e235e3146100a95780633cf4ac6d146100e657806370a082311461010f575b600080fd5b34801561008c57600080fd5b506100a760048036038101906100a291906108a9565b6101d5565b005b3480156100b557600080fd5b506100d060048036038101906100cb9190610927565b6102c6565b6040516100dd9190610963565b60405180910390f35b3480156100f257600080fd5b5061010d60048036038101906101089190610ad7565b6102de565b005b34801561011b57600080fd5b5061013660048036038101906101319190610927565b610423565b6040516101439190610963565b60405180910390f35b34801561015857600080fd5b5061016161046c565b60405161016e9190610963565b60405180910390f35b34801561018357600080fd5b5061018c610475565b6040516101999190610bde565b60405180910390f35b6101aa610503565b005b3480156101b857600080fd5b506101d360048036038101906101ce91906108a9565b61059e565b005b60008054436101e49190610c2f565b03610224576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021b90610cbd565b60405180910390fd5b61022d33610719565b61026c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026390610d29565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102bb9190610d78565b925050819055505050565b60026020528060005260406000206000915090505481565b60008054436102ed9190610c2f565b0361032d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032490610cbd565b60405180910390fd5b61033633610719565b610375576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036c90610d29565b60405180910390fd5b6001600061038391906107c3565b60005b815181101561041f5760018282815181106103a4576103a3610dac565b5b60200260200101519080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061041790610ddb565b915050610386565b5050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008054905090565b606060018054806020026020016040519081016040528092919081815260200182805480156104f957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116104af575b5050505050905090565b60003403610546576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053d90610e6f565b60405180910390fd5b34600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105959190610e8f565b92505081905550565b6105a733610719565b6105e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105dd90610d29565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610668576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065f90610f35565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405161068e90610f86565b60006040518083038185875af1925050503d80600081146106cb576040519150601f19603f3d011682016040523d82523d6000602084013e6106d0565b606091505b5050905080610714576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070b90610fe7565b60405180910390fd5b505050565b6000806000905060005b6001805490508110156107b9578373ffffffffffffffffffffffffffffffffffffffff166001828154811061075b5761075a610dac565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036107a657600191505b80806107b190610ddb565b915050610723565b5080915050919050565b50805460008255906000526020600020908101906107e191906107e4565b50565b5b808211156107fd5760008160009055506001016107e5565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061084082610815565b9050919050565b61085081610835565b811461085b57600080fd5b50565b60008135905061086d81610847565b92915050565b6000819050919050565b61088681610873565b811461089157600080fd5b50565b6000813590506108a38161087d565b92915050565b600080604083850312156108c0576108bf61080b565b5b60006108ce8582860161085e565b92505060206108df85828601610894565b9150509250929050565b60006108f482610815565b9050919050565b610904816108e9565b811461090f57600080fd5b50565b600081359050610921816108fb565b92915050565b60006020828403121561093d5761093c61080b565b5b600061094b84828501610912565b91505092915050565b61095d81610873565b82525050565b60006020820190506109786000830184610954565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6109cc82610983565b810181811067ffffffffffffffff821117156109eb576109ea610994565b5b80604052505050565b60006109fe610801565b9050610a0a82826109c3565b919050565b600067ffffffffffffffff821115610a2a57610a29610994565b5b602082029050602081019050919050565b600080fd5b6000610a53610a4e84610a0f565b6109f4565b90508083825260208201905060208402830185811115610a7657610a75610a3b565b5b835b81811015610a9f5780610a8b8882610912565b845260208401935050602081019050610a78565b5050509392505050565b600082601f830112610abe57610abd61097e565b5b8135610ace848260208601610a40565b91505092915050565b600060208284031215610aed57610aec61080b565b5b600082013567ffffffffffffffff811115610b0b57610b0a610810565b5b610b1784828501610aa9565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610b55816108e9565b82525050565b6000610b678383610b4c565b60208301905092915050565b6000602082019050919050565b6000610b8b82610b20565b610b958185610b2b565b9350610ba083610b3c565b8060005b83811015610bd1578151610bb88882610b5b565b9750610bc383610b73565b925050600181019050610ba4565b5085935050505092915050565b60006020820190508181036000830152610bf88184610b80565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610c3a82610873565b9150610c4583610873565b925082610c5557610c54610c00565b5b828206905092915050565b600082825260208201905092915050565b7f6e6f7420636865636b706f696e74000000000000000000000000000000000000600082015250565b6000610ca7600e83610c60565b9150610cb282610c71565b602082019050919050565b60006020820190508181036000830152610cd681610c9a565b9050919050565b7f6e6f74207369676e657200000000000000000000000000000000000000000000600082015250565b6000610d13600a83610c60565b9150610d1e82610cdd565b602082019050919050565b60006020820190508181036000830152610d4281610d06565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d8382610873565b9150610d8e83610873565b9250828203905081811115610da657610da5610d49565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610de682610873565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e1857610e17610d49565b5b600182019050919050565b7f77726f6e672076616c7565000000000000000000000000000000000000000000600082015250565b6000610e59600b83610c60565b9150610e6482610e23565b602082019050919050565b60006020820190508181036000830152610e8881610e4c565b9050919050565b6000610e9a82610873565b9150610ea583610873565b9250828201905080821115610ebd57610ebc610d49565b5b92915050565b7f776974686472617720616d6f756e742073686f756c64206265206c657373206f60008201527f7220657175616c20796f75722062616c616e6365000000000000000000000000602082015250565b6000610f1f603483610c60565b9150610f2a82610ec3565b604082019050919050565b60006020820190508181036000830152610f4e81610f12565b9050919050565b600081905092915050565b50565b6000610f70600083610f55565b9150610f7b82610f60565b600082019050919050565b6000610f9182610f63565b9150819050919050565b7f4661696c656420746f2073656e64000000000000000000000000000000000000600082015250565b6000610fd1600e83610c60565b9150610fdc82610f9b565b602082019050919050565b6000602082019050818103600083015261100081610fc4565b905091905056fea2646970667358221220b5f8ba5ead37b21ae17ef061cbfd5dbe3125b5d8c92da362eac5a4e1ff30b16464736f6c63430008100033",
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

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_Staker *StakerCaller) GetEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "getEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_Staker *StakerSession) GetEpoch() (*big.Int, error) {
	return _Staker.Contract.GetEpoch(&_Staker.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_Staker *StakerCallerSession) GetEpoch() (*big.Int, error) {
	return _Staker.Contract.GetEpoch(&_Staker.CallOpts)
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

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address _to, uint256 _amount) returns()
func (_Staker *StakerTransactor) Slash(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "slash", _to, _amount)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address _to, uint256 _amount) returns()
func (_Staker *StakerSession) Slash(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.Slash(&_Staker.TransactOpts, _to, _amount)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address _to, uint256 _amount) returns()
func (_Staker *StakerTransactorSession) Slash(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.Slash(&_Staker.TransactOpts, _to, _amount)
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
