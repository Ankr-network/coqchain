// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staker

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/Ankr-network/coqchain"
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
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakerMetaData contains all meta data concerning the Staker contract.
var StakerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fineRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vote\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"authorize\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"signerContains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610b3c806100206000396000f3fe6080604052600436106100865760003560e01c806385be77781161005957806385be777814610126578063900cf0cf146101515780639aa445241461017c578063bea23d46146101b9578063d24df893146101f757610086565b806327e235e31461008b5780632e1a7d4d146100c85780633a4b66f1146100f157806342cde4e8146100fb575b600080fd5b34801561009757600080fd5b506100b260048036038101906100ad919061065b565b610222565b6040516100bf91906106a1565b60405180910390f35b3480156100d457600080fd5b506100ef60048036038101906100ea91906106e8565b61023a565b005b6100f96103a5565b005b34801561010757600080fd5b50610110610440565b60405161011d91906106a1565b60405180910390f35b34801561013257600080fd5b5061013b610446565b60405161014891906106a1565b60405180910390f35b34801561015d57600080fd5b5061016661044c565b60405161017391906106a1565b60405180910390f35b34801561018857600080fd5b506101a3600480360381019061019e919061065b565b610452565b6040516101b09190610730565b60405180910390f35b3480156101c557600080fd5b506101e060048036038101906101db919061074b565b610500565b6040516101ee92919061079a565b60405180910390f35b34801561020357600080fd5b5061020c61056a565b6040516102199190610881565b60405180910390f35b60046020528060005260406000206000915090505481565b61024333610452565b15610283576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027a90610900565b60405180910390fd5b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610305576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fc9061096c565b60405180910390fd5b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461035491906109bb565b925050819055503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156103a1573d6000803e3d6000fd5b5050565b600034116103e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103df90610a3b565b60405180910390fd5b34600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104379190610a5b565b92505081905550565b60015481565b60025481565b60005481565b600080600090505b6003805490508110156104f5578273ffffffffffffffffffffffffffffffffffffffff166003828154811061049257610491610a8f565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036104e25760019150506104fb565b80806104ed90610abe565b91505061045a565b50600090505b919050565b6005602052816000526040600020818154811061051c57600080fd5b90600052602060002001600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16905082565b606060038054806020026020016040519081016040528092919081815260200182805480156105ee57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116105a4575b5050505050905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610628826105fd565b9050919050565b6106388161061d565b811461064357600080fd5b50565b6000813590506106558161062f565b92915050565b600060208284031215610671576106706105f8565b5b600061067f84828501610646565b91505092915050565b6000819050919050565b61069b81610688565b82525050565b60006020820190506106b66000830184610692565b92915050565b6106c581610688565b81146106d057600080fd5b50565b6000813590506106e2816106bc565b92915050565b6000602082840312156106fe576106fd6105f8565b5b600061070c848285016106d3565b91505092915050565b60008115159050919050565b61072a81610715565b82525050565b60006020820190506107456000830184610721565b92915050565b60008060408385031215610762576107616105f8565b5b600061077085828601610646565b9250506020610781858286016106d3565b9150509250929050565b6107948161061d565b82525050565b60006040820190506107af600083018561078b565b6107bc6020830184610721565b9392505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6107f88161061d565b82525050565b600061080a83836107ef565b60208301905092915050565b6000602082019050919050565b600061082e826107c3565b61083881856107ce565b9350610843836107df565b8060005b8381101561087457815161085b88826107fe565b975061086683610816565b925050600181019050610847565b5085935050505092915050565b6000602082019050818103600083015261089b8184610823565b905092915050565b600082825260208201905092915050565b7f7374616b696e672c20756e61626c6520746f2077697468647261770000000000600082015250565b60006108ea601b836108a3565b91506108f5826108b4565b602082019050919050565b60006020820190508181036000830152610919816108dd565b9050919050565b7f696e73756666696369656e7420616d6f756e7400000000000000000000000000600082015250565b60006109566013836108a3565b915061096182610920565b602082019050919050565b6000602082019050818103600083015261098581610949565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006109c682610688565b91506109d183610688565b92508282039050818111156109e9576109e861098c565b5b92915050565b7f696e76616c696420616d6f756e74000000000000000000000000000000000000600082015250565b6000610a25600e836108a3565b9150610a30826109ef565b602082019050919050565b60006020820190508181036000830152610a5481610a18565b9050919050565b6000610a6682610688565b9150610a7183610688565b9250828201905080821115610a8957610a8861098c565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610ac982610688565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610afb57610afa61098c565b5b60018201905091905056fea264697066735822122021e6490f7b3fdd7aa89dbc407c0477d18e288fc86f5f62f309dde341d1b7786864736f6c63430008110033",
}

// StakerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakerMetaData.ABI instead.
var StakerABI = StakerMetaData.ABI

// StakerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakerMetaData.Bin instead.
var StakerBin = StakerMetaData.Bin

// DeployStaker deploys a new Ethereum contract, binding an instance of Staker to it.
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

// Staker is an auto generated Go binding around an Ethereum contract.
type Staker struct {
	StakerCaller     // Read-only binding to the contract
	StakerTransactor // Write-only binding to the contract
	StakerFilterer   // Log filterer for contract events
}

// StakerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakerSession struct {
	Contract     *Staker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakerCallerSession struct {
	Contract *StakerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakerTransactorSession struct {
	Contract     *StakerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakerRaw struct {
	Contract *Staker // Generic contract binding to access the raw methods on
}

// StakerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakerCallerRaw struct {
	Contract *StakerCaller // Generic read-only contract binding to access the raw methods on
}

// StakerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
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

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Staker *StakerCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Staker *StakerSession) Epoch() (*big.Int, error) {
	return _Staker.Contract.Epoch(&_Staker.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Staker *StakerCallerSession) Epoch() (*big.Int, error) {
	return _Staker.Contract.Epoch(&_Staker.CallOpts)
}

// FineRatio is a free data retrieval call binding the contract method 0x85be7778.
//
// Solidity: function fineRatio() view returns(uint256)
func (_Staker *StakerCaller) FineRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "fineRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FineRatio is a free data retrieval call binding the contract method 0x85be7778.
//
// Solidity: function fineRatio() view returns(uint256)
func (_Staker *StakerSession) FineRatio() (*big.Int, error) {
	return _Staker.Contract.FineRatio(&_Staker.CallOpts)
}

// FineRatio is a free data retrieval call binding the contract method 0x85be7778.
//
// Solidity: function fineRatio() view returns(uint256)
func (_Staker *StakerCallerSession) FineRatio() (*big.Int, error) {
	return _Staker.Contract.FineRatio(&_Staker.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0xbea23d46.
//
// Solidity: function proposalVotes(address , uint256 ) view returns(address vote, bool authorize)
func (_Staker *StakerCaller) ProposalVotes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Vote      common.Address
	Authorize bool
}, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "proposalVotes", arg0, arg1)

	outstruct := new(struct {
		Vote      common.Address
		Authorize bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Vote = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Authorize = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0xbea23d46.
//
// Solidity: function proposalVotes(address , uint256 ) view returns(address vote, bool authorize)
func (_Staker *StakerSession) ProposalVotes(arg0 common.Address, arg1 *big.Int) (struct {
	Vote      common.Address
	Authorize bool
}, error) {
	return _Staker.Contract.ProposalVotes(&_Staker.CallOpts, arg0, arg1)
}

// ProposalVotes is a free data retrieval call binding the contract method 0xbea23d46.
//
// Solidity: function proposalVotes(address , uint256 ) view returns(address vote, bool authorize)
func (_Staker *StakerCallerSession) ProposalVotes(arg0 common.Address, arg1 *big.Int) (struct {
	Vote      common.Address
	Authorize bool
}, error) {
	return _Staker.Contract.ProposalVotes(&_Staker.CallOpts, arg0, arg1)
}

// SignerContains is a free data retrieval call binding the contract method 0x9aa44524.
//
// Solidity: function signerContains(address _signer) view returns(bool)
func (_Staker *StakerCaller) SignerContains(opts *bind.CallOpts, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "signerContains", _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SignerContains is a free data retrieval call binding the contract method 0x9aa44524.
//
// Solidity: function signerContains(address _signer) view returns(bool)
func (_Staker *StakerSession) SignerContains(_signer common.Address) (bool, error) {
	return _Staker.Contract.SignerContains(&_Staker.CallOpts, _signer)
}

// SignerContains is a free data retrieval call binding the contract method 0x9aa44524.
//
// Solidity: function signerContains(address _signer) view returns(bool)
func (_Staker *StakerCallerSession) SignerContains(_signer common.Address) (bool, error) {
	return _Staker.Contract.SignerContains(&_Staker.CallOpts, _signer)
}

// SignerList is a free data retrieval call binding the contract method 0xd24df893.
//
// Solidity: function signerList() view returns(address[])
func (_Staker *StakerCaller) SignerList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "signerList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SignerList is a free data retrieval call binding the contract method 0xd24df893.
//
// Solidity: function signerList() view returns(address[])
func (_Staker *StakerSession) SignerList() ([]common.Address, error) {
	return _Staker.Contract.SignerList(&_Staker.CallOpts)
}

// SignerList is a free data retrieval call binding the contract method 0xd24df893.
//
// Solidity: function signerList() view returns(address[])
func (_Staker *StakerCallerSession) SignerList() ([]common.Address, error) {
	return _Staker.Contract.SignerList(&_Staker.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Staker *StakerCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Staker *StakerSession) Threshold() (*big.Int, error) {
	return _Staker.Contract.Threshold(&_Staker.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Staker *StakerCallerSession) Threshold() (*big.Int, error) {
	return _Staker.Contract.Threshold(&_Staker.CallOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staker *StakerTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staker *StakerSession) Stake() (*types.Transaction, error) {
	return _Staker.Contract.Stake(&_Staker.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staker *StakerTransactorSession) Stake() (*types.Transaction, error) {
	return _Staker.Contract.Stake(&_Staker.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Staker *StakerTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Staker *StakerSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.Withdraw(&_Staker.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Staker *StakerTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.Withdraw(&_Staker.TransactOpts, _amount)
}
