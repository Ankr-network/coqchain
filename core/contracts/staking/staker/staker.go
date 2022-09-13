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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"in_epoch\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sigs\",\"type\":\"address[]\"}],\"name\":\"commitSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526064600055604051610ed1380380610ed1833981810160405281019061002a9190610088565b6000810361003f576064600081905550610047565b806000819055505b506100b5565b600080fd5b6000819050919050565b61006581610052565b811461007057600080fd5b50565b6000815190506100828161005c565b92915050565b60006020828403121561009e5761009d61004d565b5b60006100ac84828501610073565b91505092915050565b610e0d806100c46000396000f3fe6080604052600436106100555760003560e01c806302fb4d851461005a57806327e235e3146100835780633cf4ac6d146100c0578063722713f7146100e9578063d0e30db014610114578063f3fef3a31461011e575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c919061078d565b610147565b005b34801561008f57600080fd5b506100aa60048036038101906100a5919061080b565b610242565b6040516100b79190610847565b60405180910390f35b3480156100cc57600080fd5b506100e760048036038101906100e291906109bb565b61025a565b005b3480156100f557600080fd5b506100fe6103a9565b60405161010b9190610847565b60405180910390f35b61011c6103f0565b005b34801561012a57600080fd5b506101456004803603810190610140919061078d565b610479565b005b600080600054436101589190610a33565b141590508061019c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019390610ac1565b60405180910390fd5b6101a5336105fd565b9050806101e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101de90610b2d565b60405180910390fd5b81600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102369190610b7c565b92505081905550505050565b60036020528060005260406000206000915090505481565b6000806000544361026b9190610a33565b14159050806102af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a690610ac1565b60405180910390fd5b6102b8336105fd565b9050806102fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f190610b2d565b60405180910390fd5b6001600061030891906106a7565b60005b82518110156103a457600183828151811061032957610328610bb0565b5b60200260200101519080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061039c90610bdf565b91505061030b565b505050565b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003403610433576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042a90610c73565b60405180910390fd5b34600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550565b6000610484336105fd565b9050806104c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bd90610b2d565b60405180910390fd5b81600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101590508061054b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054290610d05565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff168360405161057190610d56565b60006040518083038185875af1925050503d80600081146105ae576040519150601f19603f3d011682016040523d82523d6000602084013e6105b3565b606091505b50509050806105f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ee90610db7565b60405180910390fd5b50505050565b6000806000905060005b60018054905081101561069d578373ffffffffffffffffffffffffffffffffffffffff166001828154811061063f5761063e610bb0565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361068a57600191505b808061069590610bdf565b915050610607565b5080915050919050565b50805460008255906000526020600020908101906106c591906106c8565b50565b5b808211156106e15760008160009055506001016106c9565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610724826106f9565b9050919050565b61073481610719565b811461073f57600080fd5b50565b6000813590506107518161072b565b92915050565b6000819050919050565b61076a81610757565b811461077557600080fd5b50565b60008135905061078781610761565b92915050565b600080604083850312156107a4576107a36106ef565b5b60006107b285828601610742565b92505060206107c385828601610778565b9150509250929050565b60006107d8826106f9565b9050919050565b6107e8816107cd565b81146107f357600080fd5b50565b600081359050610805816107df565b92915050565b600060208284031215610821576108206106ef565b5b600061082f848285016107f6565b91505092915050565b61084181610757565b82525050565b600060208201905061085c6000830184610838565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6108b082610867565b810181811067ffffffffffffffff821117156108cf576108ce610878565b5b80604052505050565b60006108e26106e5565b90506108ee82826108a7565b919050565b600067ffffffffffffffff82111561090e5761090d610878565b5b602082029050602081019050919050565b600080fd5b6000610937610932846108f3565b6108d8565b9050808382526020820190506020840283018581111561095a5761095961091f565b5b835b81811015610983578061096f88826107f6565b84526020840193505060208101905061095c565b5050509392505050565b600082601f8301126109a2576109a1610862565b5b81356109b2848260208601610924565b91505092915050565b6000602082840312156109d1576109d06106ef565b5b600082013567ffffffffffffffff8111156109ef576109ee6106f4565b5b6109fb8482850161098d565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610a3e82610757565b9150610a4983610757565b925082610a5957610a58610a04565b5b828206905092915050565b600082825260208201905092915050565b7f6e6f7420636865636b706f696e74000000000000000000000000000000000000600082015250565b6000610aab600e83610a64565b9150610ab682610a75565b602082019050919050565b60006020820190508181036000830152610ada81610a9e565b9050919050565b7f6e6f74207369676e657200000000000000000000000000000000000000000000600082015250565b6000610b17600a83610a64565b9150610b2282610ae1565b602082019050919050565b60006020820190508181036000830152610b4681610b0a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b8782610757565b9150610b9283610757565b9250828203905081811115610baa57610ba9610b4d565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610bea82610757565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c1c57610c1b610b4d565b5b600182019050919050565b7f77726f6e672076616c7565000000000000000000000000000000000000000000600082015250565b6000610c5d600b83610a64565b9150610c6882610c27565b602082019050919050565b60006020820190508181036000830152610c8c81610c50565b9050919050565b7f776974686472617720616d6f756e742073686f756c64206265206c657373206f60008201527f7220657175616c20796f75722062616c616e6365000000000000000000000000602082015250565b6000610cef603483610a64565b9150610cfa82610c93565b604082019050919050565b60006020820190508181036000830152610d1e81610ce2565b9050919050565b600081905092915050565b50565b6000610d40600083610d25565b9150610d4b82610d30565b600082019050919050565b6000610d6182610d33565b9150819050919050565b7f4661696c656420746f2073656e64000000000000000000000000000000000000600082015250565b6000610da1600e83610a64565b9150610dac82610d6b565b602082019050919050565b60006020820190508181036000830152610dd081610d94565b905091905056fea26469706673582212200f4ba89e84fbd7155bd9dd7ef689f1d4d2879c515a438fdfd8baf7148827f28464736f6c63430008100033",
}

// StakerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakerMetaData.ABI instead.
var StakerABI = StakerMetaData.ABI

// StakerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakerMetaData.Bin instead.
var StakerBin = StakerMetaData.Bin

// DeployStaker deploys a new coqchain contract, binding an instance of Staker to it.
func DeployStaker(auth *bind.TransactOpts, backend bind.ContractBackend, in_epoch *big.Int) (common.Address, *types.Transaction, *Staker, error) {
	parsed, err := StakerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakerBin), backend, in_epoch)
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
