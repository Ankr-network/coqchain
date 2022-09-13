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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sigs\",\"type\":\"address[]\"}],\"name\":\"commitSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526064600055610e36806100186000396000f3fe6080604052600436106100555760003560e01c806302fb4d851461005a57806327e235e3146100835780633cf4ac6d146100c0578063722713f7146100e9578063d0e30db014610114578063f3fef3a31461011e575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c9190610782565b610147565b005b34801561008f57600080fd5b506100aa60048036038101906100a59190610800565b610238565b6040516100b7919061083c565b60405180910390f35b3480156100cc57600080fd5b506100e760048036038101906100e291906109b0565b610250565b005b3480156100f557600080fd5b506100fe610395565b60405161010b919061083c565b60405180910390f35b61011c6103dc565b005b34801561012a57600080fd5b5061014560048036038101906101409190610782565b610477565b005b60008054436101569190610a28565b03610196576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018d90610ab6565b60405180910390fd5b61019f336105f2565b6101de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d590610b22565b60405180910390fd5b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461022d9190610b71565b925050819055505050565b60036020528060005260406000206000915090505481565b600080544361025f9190610a28565b0361029f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029690610ab6565b60405180910390fd5b6102a8336105f2565b6102e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102de90610b22565b60405180910390fd5b600160006102f5919061069c565b60005b815181101561039157600182828151811061031657610315610ba5565b5b60200260200101519080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061038990610bd4565b9150506102f8565b5050565b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b6000340361041f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041690610c68565b60405180910390fd5b34600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461046e9190610c88565b92505081905550565b610480336105f2565b6104bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b690610b22565b60405180910390fd5b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610541576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053890610d2e565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405161056790610d7f565b60006040518083038185875af1925050503d80600081146105a4576040519150601f19603f3d011682016040523d82523d6000602084013e6105a9565b606091505b50509050806105ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e490610de0565b60405180910390fd5b505050565b6000806000905060005b600180549050811015610692578373ffffffffffffffffffffffffffffffffffffffff166001828154811061063457610633610ba5565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361067f57600191505b808061068a90610bd4565b9150506105fc565b5080915050919050565b50805460008255906000526020600020908101906106ba91906106bd565b50565b5b808211156106d65760008160009055506001016106be565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610719826106ee565b9050919050565b6107298161070e565b811461073457600080fd5b50565b60008135905061074681610720565b92915050565b6000819050919050565b61075f8161074c565b811461076a57600080fd5b50565b60008135905061077c81610756565b92915050565b60008060408385031215610799576107986106e4565b5b60006107a785828601610737565b92505060206107b88582860161076d565b9150509250929050565b60006107cd826106ee565b9050919050565b6107dd816107c2565b81146107e857600080fd5b50565b6000813590506107fa816107d4565b92915050565b600060208284031215610816576108156106e4565b5b6000610824848285016107eb565b91505092915050565b6108368161074c565b82525050565b6000602082019050610851600083018461082d565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6108a58261085c565b810181811067ffffffffffffffff821117156108c4576108c361086d565b5b80604052505050565b60006108d76106da565b90506108e3828261089c565b919050565b600067ffffffffffffffff8211156109035761090261086d565b5b602082029050602081019050919050565b600080fd5b600061092c610927846108e8565b6108cd565b9050808382526020820190506020840283018581111561094f5761094e610914565b5b835b81811015610978578061096488826107eb565b845260208401935050602081019050610951565b5050509392505050565b600082601f83011261099757610996610857565b5b81356109a7848260208601610919565b91505092915050565b6000602082840312156109c6576109c56106e4565b5b600082013567ffffffffffffffff8111156109e4576109e36106e9565b5b6109f084828501610982565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610a338261074c565b9150610a3e8361074c565b925082610a4e57610a4d6109f9565b5b828206905092915050565b600082825260208201905092915050565b7f6e6f7420636865636b706f696e74000000000000000000000000000000000000600082015250565b6000610aa0600e83610a59565b9150610aab82610a6a565b602082019050919050565b60006020820190508181036000830152610acf81610a93565b9050919050565b7f6e6f74207369676e657200000000000000000000000000000000000000000000600082015250565b6000610b0c600a83610a59565b9150610b1782610ad6565b602082019050919050565b60006020820190508181036000830152610b3b81610aff565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b7c8261074c565b9150610b878361074c565b9250828203905081811115610b9f57610b9e610b42565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610bdf8261074c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c1157610c10610b42565b5b600182019050919050565b7f77726f6e672076616c7565000000000000000000000000000000000000000000600082015250565b6000610c52600b83610a59565b9150610c5d82610c1c565b602082019050919050565b60006020820190508181036000830152610c8181610c45565b9050919050565b6000610c938261074c565b9150610c9e8361074c565b9250828201905080821115610cb657610cb5610b42565b5b92915050565b7f776974686472617720616d6f756e742073686f756c64206265206c657373206f60008201527f7220657175616c20796f75722062616c616e6365000000000000000000000000602082015250565b6000610d18603483610a59565b9150610d2382610cbc565b604082019050919050565b60006020820190508181036000830152610d4781610d0b565b9050919050565b600081905092915050565b50565b6000610d69600083610d4e565b9150610d7482610d59565b600082019050919050565b6000610d8a82610d5c565b9150819050919050565b7f4661696c656420746f2073656e64000000000000000000000000000000000000600082015250565b6000610dca600e83610a59565b9150610dd582610d94565b602082019050919050565b60006020820190508181036000830152610df981610dbd565b905091905056fea2646970667358221220b630f752a364bdd73787c69f5365d0ba19c2645330549663c1de97993d5ede4464736f6c63430008100033",
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

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Staker *StakerCaller) BalanceOf(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "balanceOf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Staker *StakerSession) BalanceOf() (*big.Int, error) {
	return _Staker.Contract.BalanceOf(&_Staker.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Staker *StakerCallerSession) BalanceOf() (*big.Int, error) {
	return _Staker.Contract.BalanceOf(&_Staker.CallOpts)
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
