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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Consume\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"consume\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setadmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405233600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610b4a806100546000396000f3fe60806040526004361061004a5760003560e01c80631dedc6f71461004f5780633ccfd60b1461005957806370a0823114610070578063a9059cbb146100ad578063dded49cb146100c9575b600080fd5b6100576100f2565b005b34801561006557600080fd5b5061006e6101c5565b005b34801561007c57600080fd5b50610097600480360381019061009291906106a8565b6103a9565b6040516100a491906106ee565b60405180910390f35b6100c760048036038101906100c29190610735565b6103f1565b005b3480156100d557600080fd5b506100f060048036038101906100eb91906106a8565b61055b565b005b60003403610135576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012c906107d2565b60405180910390fd5b346000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546101839190610821565b925050819055507fb3762e93ec66871dd27c421b64edc79636345ff0a949cd04f7f8efce5bd4240e33346040516101bb929190610864565b60405180910390a1565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610255576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024c906108d9565b60405180910390fd5b60004790506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826040516102a29061092a565b60006040518083038185875af1925050503d80600081146102df576040519150601f19603f3d011682016040523d82523d6000602084013e6102e4565b606091505b5050905080610328576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031f9061098b565b60405180910390fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364338360405161039d929190610864565b60405180910390a15050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610472576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046990610a1d565b60405180910390fd5b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104c09190610a3d565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105159190610821565b925050819055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef33838360405161054f93929190610a71565b60405180910390a15050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e290610af4565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106758261064a565b9050919050565b6106858161066a565b811461069057600080fd5b50565b6000813590506106a28161067c565b92915050565b6000602082840312156106be576106bd610645565b5b60006106cc84828501610693565b91505092915050565b6000819050919050565b6106e8816106d5565b82525050565b600060208201905061070360008301846106df565b92915050565b610712816106d5565b811461071d57600080fd5b50565b60008135905061072f81610709565b92915050565b6000806040838503121561074c5761074b610645565b5b600061075a85828601610693565b925050602061076b85828601610720565b9150509250929050565b600082825260208201905092915050565b7f77726f6e672076616c7565000000000000000000000000000000000000000000600082015250565b60006107bc600b83610775565b91506107c782610786565b602082019050919050565b600060208201905081810360008301526107eb816107af565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061082c826106d5565b9150610837836106d5565b925082820190508082111561084f5761084e6107f2565b5b92915050565b61085e8161066a565b82525050565b60006040820190506108796000830185610855565b61088660208301846106df565b9392505050565b7f6f6e6c79206f776e657200000000000000000000000000000000000000000000600082015250565b60006108c3600a83610775565b91506108ce8261088d565b602082019050919050565b600060208201905081810360008301526108f2816108b6565b9050919050565b600081905092915050565b50565b60006109146000836108f9565b915061091f82610904565b600082019050919050565b600061093582610907565b9150819050919050565b7f4661696c656420746f2073656e64000000000000000000000000000000000000600082015250565b6000610975600e83610775565b91506109808261093f565b602082019050919050565b600060208201905081810360008301526109a481610968565b9050919050565b7f7472616e7366657220616d6f756e742073686f756c64206265206c657373206f60008201527f7220657175616c20796f75722062616c616e6365000000000000000000000000602082015250565b6000610a07603483610775565b9150610a12826109ab565b604082019050919050565b60006020820190508181036000830152610a36816109fa565b9050919050565b6000610a48826106d5565b9150610a53836106d5565b9250828203905081811115610a6b57610a6a6107f2565b5b92915050565b6000606082019050610a866000830186610855565b610a936020830185610855565b610aa060408301846106df565b949350505050565b7f6f6e6c79206f776e65722063616e20696e766f6b650000000000000000000000600082015250565b6000610ade601583610775565b9150610ae982610aa8565b602082019050919050565b60006020820190508181036000830152610b0d81610ad1565b905091905056fea2646970667358221220c6760702efa5acecbe2939b68bb8a10b06b2ef2f7de427bd0d41b17d6c113de364736f6c63430008100033",
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

// Consume is a paid mutator transaction binding the contract method 0x1dedc6f7.
//
// Solidity: function consume() payable returns()
func (_Staker *StakerTransactor) Consume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "consume")
}

// Consume is a paid mutator transaction binding the contract method 0x1dedc6f7.
//
// Solidity: function consume() payable returns()
func (_Staker *StakerSession) Consume() (*types.Transaction, error) {
	return _Staker.Contract.Consume(&_Staker.TransactOpts)
}

// Consume is a paid mutator transaction binding the contract method 0x1dedc6f7.
//
// Solidity: function consume() payable returns()
func (_Staker *StakerTransactorSession) Consume() (*types.Transaction, error) {
	return _Staker.Contract.Consume(&_Staker.TransactOpts)
}

// Setadmin is a paid mutator transaction binding the contract method 0xdded49cb.
//
// Solidity: function setadmin(address _admin) returns()
func (_Staker *StakerTransactor) Setadmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "setadmin", _admin)
}

// Setadmin is a paid mutator transaction binding the contract method 0xdded49cb.
//
// Solidity: function setadmin(address _admin) returns()
func (_Staker *StakerSession) Setadmin(_admin common.Address) (*types.Transaction, error) {
	return _Staker.Contract.Setadmin(&_Staker.TransactOpts, _admin)
}

// Setadmin is a paid mutator transaction binding the contract method 0xdded49cb.
//
// Solidity: function setadmin(address _admin) returns()
func (_Staker *StakerTransactorSession) Setadmin(_admin common.Address) (*types.Transaction, error) {
	return _Staker.Contract.Setadmin(&_Staker.TransactOpts, _admin)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) payable returns()
func (_Staker *StakerTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) payable returns()
func (_Staker *StakerSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.Transfer(&_Staker.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) payable returns()
func (_Staker *StakerTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.Transfer(&_Staker.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staker *StakerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staker *StakerSession) Withdraw() (*types.Transaction, error) {
	return _Staker.Contract.Withdraw(&_Staker.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staker *StakerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Staker.Contract.Withdraw(&_Staker.TransactOpts)
}

// StakerConsumeIterator is returned from FilterConsume and is used to iterate over the raw logs and unpacked data for Consume events raised by the Staker contract.
type StakerConsumeIterator struct {
	Event *StakerConsume // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  coqchain.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakerConsumeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerConsume)
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
		it.Event = new(StakerConsume)
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
func (it *StakerConsumeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerConsumeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerConsume represents a Consume event raised by the Staker contract.
type StakerConsume struct {
	Sender common.Address
	Amt    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConsume is a free log retrieval operation binding the contract event 0xb3762e93ec66871dd27c421b64edc79636345ff0a949cd04f7f8efce5bd4240e.
//
// Solidity: event Consume(address sender, uint256 amt)
func (_Staker *StakerFilterer) FilterConsume(opts *bind.FilterOpts) (*StakerConsumeIterator, error) {

	logs, sub, err := _Staker.contract.FilterLogs(opts, "Consume")
	if err != nil {
		return nil, err
	}
	return &StakerConsumeIterator{contract: _Staker.contract, event: "Consume", logs: logs, sub: sub}, nil
}

// WatchConsume is a free log subscription operation binding the contract event 0xb3762e93ec66871dd27c421b64edc79636345ff0a949cd04f7f8efce5bd4240e.
//
// Solidity: event Consume(address sender, uint256 amt)
func (_Staker *StakerFilterer) WatchConsume(opts *bind.WatchOpts, sink chan<- *StakerConsume) (event.Subscription, error) {

	logs, sub, err := _Staker.contract.WatchLogs(opts, "Consume")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerConsume)
				if err := _Staker.contract.UnpackLog(event, "Consume", log); err != nil {
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

// ParseConsume is a log parse operation binding the contract event 0xb3762e93ec66871dd27c421b64edc79636345ff0a949cd04f7f8efce5bd4240e.
//
// Solidity: event Consume(address sender, uint256 amt)
func (_Staker *StakerFilterer) ParseConsume(log types.Log) (*StakerConsume, error) {
	event := new(StakerConsume)
	if err := _Staker.contract.UnpackLog(event, "Consume", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Staker contract.
type StakerTransferIterator struct {
	Event *StakerTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  coqchain.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerTransfer)
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
		it.Event = new(StakerTransfer)
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
func (it *StakerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerTransfer represents a Transfer event raised by the Staker contract.
type StakerTransfer struct {
	From common.Address
	To   common.Address
	Amt  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 amt)
func (_Staker *StakerFilterer) FilterTransfer(opts *bind.FilterOpts) (*StakerTransferIterator, error) {

	logs, sub, err := _Staker.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &StakerTransferIterator{contract: _Staker.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 amt)
func (_Staker *StakerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakerTransfer) (event.Subscription, error) {

	logs, sub, err := _Staker.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerTransfer)
				if err := _Staker.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address from, address to, uint256 amt)
func (_Staker *StakerFilterer) ParseTransfer(log types.Log) (*StakerTransfer, error) {
	event := new(StakerTransfer)
	if err := _Staker.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Staker contract.
type StakerWithdrawIterator struct {
	Event *StakerWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  coqchain.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerWithdraw)
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
		it.Event = new(StakerWithdraw)
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
func (it *StakerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerWithdraw represents a Withdraw event raised by the Staker contract.
type StakerWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address to, uint256 amount)
func (_Staker *StakerFilterer) FilterWithdraw(opts *bind.FilterOpts) (*StakerWithdrawIterator, error) {

	logs, sub, err := _Staker.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &StakerWithdrawIterator{contract: _Staker.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address to, uint256 amount)
func (_Staker *StakerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StakerWithdraw) (event.Subscription, error) {

	logs, sub, err := _Staker.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerWithdraw)
				if err := _Staker.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address to, uint256 amount)
func (_Staker *StakerFilterer) ParseWithdraw(log types.Log) (*StakerWithdraw, error) {
	event := new(StakerWithdraw)
	if err := _Staker.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
