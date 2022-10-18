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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_votee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"checkVoteStatus\",\"outputs\":[{\"internalType\":\"enumStaker.VoteRes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochProposalVotees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochProposals\",\"outputs\":[{\"internalType\":\"enumStaker.VoteType\",\"name\":\"voteType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"epochVotedByBlockNumber\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fineRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"signerContains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611013806100206000396000f3fe6080604052600436106100c25760003560e01c806342cde4e81161007f578063900cf0cf11610059578063900cf0cf146102815780639aa44524146102ac578063a5a3ce57146102e9578063d24df89314610326576100c2565b806342cde4e8146101ee578063760cdd2f1461021957806385be777814610256576100c2565b80632026f638146100c757806325ee57191461010457806327e235e3146101415780632e1a7d4d1461017e5780633a4b66f1146101a75780633bf81c0c146101b1575b600080fd5b3480156100d357600080fd5b506100ee60048036038101906100e99190610901565b610351565b6040516100fb919061093d565b60405180910390f35b34801561011057600080fd5b5061012b600480360381019061012691906109b6565b610367565b6040516101389190610a80565b60405180910390f35b34801561014d57600080fd5b5061016860048036038101906101639190610a9b565b610418565b604051610175919061093d565b60405180910390f35b34801561018a57600080fd5b506101a560048036038101906101a09190610901565b610430565b005b6101af6105b8565b005b3480156101bd57600080fd5b506101d860048036038101906101d39190610ac8565b610653565b6040516101e59190610b50565b60405180910390f35b3480156101fa57600080fd5b5061020361068b565b604051610210919061093d565b60405180910390f35b34801561022557600080fd5b50610240600480360381019061023b9190610901565b610691565b60405161024d9190610b86565b60405180910390f35b34801561026257600080fd5b5061026b6106c3565b604051610278919061093d565b60405180910390f35b34801561028d57600080fd5b506102966106c9565b6040516102a3919061093d565b60405180910390f35b3480156102b857600080fd5b506102d360048036038101906102ce9190610a9b565b6106cf565b6040516102e09190610b86565b60405180910390f35b3480156102f557600080fd5b50610310600480360381019061030b9190610ba1565b6106fe565b60405161031d9190610bf0565b60405180910390f35b34801561033257600080fd5b5061033b61074c565b6040516103489190610cc9565b60405180910390f35b60008054826103609190610d49565b9050919050565b60006006600061037686610351565b815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690509392505050565b60056020528060005260406000206000915090505481565b6104563360601b6bffffffffffffffffffffffff1916600361083d90919063ffffffff16565b15610496576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048d90610dd7565b60405180910390fd5b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610518576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050f90610e43565b60405180910390fd5b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105679190610e63565b925050819055503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156105b4573d6000803e3d6000fd5b5050565b600034116105fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f290610ee3565b60405180910390fd5b34600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461064a9190610f03565b92505081905550565b6006602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900460ff16905081565b60015481565b6000600860006106a084610351565b815260200190815260200160002060009054906101000a900460ff169050919050565b60025481565b60005481565b60006106f78260601b6bffffffffffffffffffffffff1916600361083d90919063ffffffff16565b9050919050565b6007602052816000526040600020818154811061071a57600080fd5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600061075a600361086a565b90506000815167ffffffffffffffff81111561077957610778610f37565b5b6040519080825280602002602001820160405280156107a75781602001602082028036833780820191505090505b50905060005b8251811015610834578281815181106107c9576107c8610f66565b5b602002602001015160601c8282815181106107e7576107e6610f66565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808061082c90610f95565b9150506107ad565b50809250505090565b600082600101600083815260200190815260200160002060009054906101000a900460ff16905092915050565b6060816000018054806020026020016040519081016040528092919081815260200182805480156108ba57602002820191906000526020600020905b8154815260200190600101908083116108a6575b50505050509050919050565b600080fd5b6000819050919050565b6108de816108cb565b81146108e957600080fd5b50565b6000813590506108fb816108d5565b92915050565b600060208284031215610917576109166108c6565b5b6000610925848285016108ec565b91505092915050565b610937816108cb565b82525050565b6000602082019050610952600083018461092e565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061098382610958565b9050919050565b61099381610978565b811461099e57600080fd5b50565b6000813590506109b08161098a565b92915050565b6000806000606084860312156109cf576109ce6108c6565b5b60006109dd868287016108ec565b93505060206109ee868287016109a1565b92505060406109ff868287016109a1565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110610a4957610a48610a09565b5b50565b6000819050610a5a82610a38565b919050565b6000610a6a82610a4c565b9050919050565b610a7a81610a5f565b82525050565b6000602082019050610a956000830184610a71565b92915050565b600060208284031215610ab157610ab06108c6565b5b6000610abf848285016109a1565b91505092915050565b60008060408385031215610adf57610ade6108c6565b5b6000610aed858286016108ec565b9250506020610afe858286016109a1565b9150509250929050565b60038110610b1957610b18610a09565b5b50565b6000819050610b2a82610b08565b919050565b6000610b3a82610b1c565b9050919050565b610b4a81610b2f565b82525050565b6000602082019050610b656000830184610b41565b92915050565b60008115159050919050565b610b8081610b6b565b82525050565b6000602082019050610b9b6000830184610b77565b92915050565b60008060408385031215610bb857610bb76108c6565b5b6000610bc6858286016108ec565b9250506020610bd7858286016108ec565b9150509250929050565b610bea81610978565b82525050565b6000602082019050610c056000830184610be1565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610c4081610978565b82525050565b6000610c528383610c37565b60208301905092915050565b6000602082019050919050565b6000610c7682610c0b565b610c808185610c16565b9350610c8b83610c27565b8060005b83811015610cbc578151610ca38882610c46565b9750610cae83610c5e565b925050600181019050610c8f565b5085935050505092915050565b60006020820190508181036000830152610ce38184610c6b565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d54826108cb565b9150610d5f836108cb565b925082610d6f57610d6e610ceb565b5b828204905092915050565b600082825260208201905092915050565b7f7374616b696e672c20756e61626c6520746f2077697468647261770000000000600082015250565b6000610dc1601b83610d7a565b9150610dcc82610d8b565b602082019050919050565b60006020820190508181036000830152610df081610db4565b9050919050565b7f696e73756666696369656e7420616d6f756e7400000000000000000000000000600082015250565b6000610e2d601383610d7a565b9150610e3882610df7565b602082019050919050565b60006020820190508181036000830152610e5c81610e20565b9050919050565b6000610e6e826108cb565b9150610e79836108cb565b9250828203905081811115610e9157610e90610d1a565b5b92915050565b7f696e76616c696420616d6f756e74000000000000000000000000000000000000600082015250565b6000610ecd600e83610d7a565b9150610ed882610e97565b602082019050919050565b60006020820190508181036000830152610efc81610ec0565b9050919050565b6000610f0e826108cb565b9150610f19836108cb565b9250828201905080821115610f3157610f30610d1a565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610fa0826108cb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fd257610fd1610d1a565b5b60018201905091905056fea2646970667358221220920b7392adb9863ac98612f64c6bbb3fe88696d8ba8cc1f8bd7758ce02ecaf0f64736f6c63430008110033",
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

// CheckVoteStatus is a free data retrieval call binding the contract method 0x25ee5719.
//
// Solidity: function checkVoteStatus(uint256 _blockNumber, address _votee, address _voter) view returns(uint8)
func (_Staker *StakerCaller) CheckVoteStatus(opts *bind.CallOpts, _blockNumber *big.Int, _votee common.Address, _voter common.Address) (uint8, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "checkVoteStatus", _blockNumber, _votee, _voter)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CheckVoteStatus is a free data retrieval call binding the contract method 0x25ee5719.
//
// Solidity: function checkVoteStatus(uint256 _blockNumber, address _votee, address _voter) view returns(uint8)
func (_Staker *StakerSession) CheckVoteStatus(_blockNumber *big.Int, _votee common.Address, _voter common.Address) (uint8, error) {
	return _Staker.Contract.CheckVoteStatus(&_Staker.CallOpts, _blockNumber, _votee, _voter)
}

// CheckVoteStatus is a free data retrieval call binding the contract method 0x25ee5719.
//
// Solidity: function checkVoteStatus(uint256 _blockNumber, address _votee, address _voter) view returns(uint8)
func (_Staker *StakerCallerSession) CheckVoteStatus(_blockNumber *big.Int, _votee common.Address, _voter common.Address) (uint8, error) {
	return _Staker.Contract.CheckVoteStatus(&_Staker.CallOpts, _blockNumber, _votee, _voter)
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

// EpochProposalVotees is a free data retrieval call binding the contract method 0xa5a3ce57.
//
// Solidity: function epochProposalVotees(uint256 , uint256 ) view returns(address)
func (_Staker *StakerCaller) EpochProposalVotees(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "epochProposalVotees", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EpochProposalVotees is a free data retrieval call binding the contract method 0xa5a3ce57.
//
// Solidity: function epochProposalVotees(uint256 , uint256 ) view returns(address)
func (_Staker *StakerSession) EpochProposalVotees(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Staker.Contract.EpochProposalVotees(&_Staker.CallOpts, arg0, arg1)
}

// EpochProposalVotees is a free data retrieval call binding the contract method 0xa5a3ce57.
//
// Solidity: function epochProposalVotees(uint256 , uint256 ) view returns(address)
func (_Staker *StakerCallerSession) EpochProposalVotees(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Staker.Contract.EpochProposalVotees(&_Staker.CallOpts, arg0, arg1)
}

// EpochProposals is a free data retrieval call binding the contract method 0x3bf81c0c.
//
// Solidity: function epochProposals(uint256 , address ) view returns(uint8 voteType)
func (_Staker *StakerCaller) EpochProposals(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (uint8, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "epochProposals", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// EpochProposals is a free data retrieval call binding the contract method 0x3bf81c0c.
//
// Solidity: function epochProposals(uint256 , address ) view returns(uint8 voteType)
func (_Staker *StakerSession) EpochProposals(arg0 *big.Int, arg1 common.Address) (uint8, error) {
	return _Staker.Contract.EpochProposals(&_Staker.CallOpts, arg0, arg1)
}

// EpochProposals is a free data retrieval call binding the contract method 0x3bf81c0c.
//
// Solidity: function epochProposals(uint256 , address ) view returns(uint8 voteType)
func (_Staker *StakerCallerSession) EpochProposals(arg0 *big.Int, arg1 common.Address) (uint8, error) {
	return _Staker.Contract.EpochProposals(&_Staker.CallOpts, arg0, arg1)
}

// EpochVotedByBlockNumber is a free data retrieval call binding the contract method 0x760cdd2f.
//
// Solidity: function epochVotedByBlockNumber(uint256 _blockNumber) view returns(bool)
func (_Staker *StakerCaller) EpochVotedByBlockNumber(opts *bind.CallOpts, _blockNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "epochVotedByBlockNumber", _blockNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochVotedByBlockNumber is a free data retrieval call binding the contract method 0x760cdd2f.
//
// Solidity: function epochVotedByBlockNumber(uint256 _blockNumber) view returns(bool)
func (_Staker *StakerSession) EpochVotedByBlockNumber(_blockNumber *big.Int) (bool, error) {
	return _Staker.Contract.EpochVotedByBlockNumber(&_Staker.CallOpts, _blockNumber)
}

// EpochVotedByBlockNumber is a free data retrieval call binding the contract method 0x760cdd2f.
//
// Solidity: function epochVotedByBlockNumber(uint256 _blockNumber) view returns(bool)
func (_Staker *StakerCallerSession) EpochVotedByBlockNumber(_blockNumber *big.Int) (bool, error) {
	return _Staker.Contract.EpochVotedByBlockNumber(&_Staker.CallOpts, _blockNumber)
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

// GetCycle is a free data retrieval call binding the contract method 0x2026f638.
//
// Solidity: function getCycle(uint256 _blockNumber) view returns(uint256)
func (_Staker *StakerCaller) GetCycle(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "getCycle", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCycle is a free data retrieval call binding the contract method 0x2026f638.
//
// Solidity: function getCycle(uint256 _blockNumber) view returns(uint256)
func (_Staker *StakerSession) GetCycle(_blockNumber *big.Int) (*big.Int, error) {
	return _Staker.Contract.GetCycle(&_Staker.CallOpts, _blockNumber)
}

// GetCycle is a free data retrieval call binding the contract method 0x2026f638.
//
// Solidity: function getCycle(uint256 _blockNumber) view returns(uint256)
func (_Staker *StakerCallerSession) GetCycle(_blockNumber *big.Int) (*big.Int, error) {
	return _Staker.Contract.GetCycle(&_Staker.CallOpts, _blockNumber)
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
