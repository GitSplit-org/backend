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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"name\":\"assignAddressToUsername\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"usernames\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsInEther\",\"type\":\"uint256[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"getAddressForUsername\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"temporaryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"usernameToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amountInWei\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x05d01f50.
//
// Solidity: function balances(address , string ) view returns(uint256)
func (_Contract *ContractCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x05d01f50.
//
// Solidity: function balances(address , string ) view returns(uint256)
func (_Contract *ContractSession) Balances(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0x05d01f50.
//
// Solidity: function balances(address , string ) view returns(uint256)
func (_Contract *ContractCallerSession) Balances(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0, arg1)
}

// GetAddressForUsername is a free data retrieval call binding the contract method 0x095a0535.
//
// Solidity: function getAddressForUsername(string username) view returns(address)
func (_Contract *ContractCaller) GetAddressForUsername(opts *bind.CallOpts, username string) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAddressForUsername", username)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressForUsername is a free data retrieval call binding the contract method 0x095a0535.
//
// Solidity: function getAddressForUsername(string username) view returns(address)
func (_Contract *ContractSession) GetAddressForUsername(username string) (common.Address, error) {
	return _Contract.Contract.GetAddressForUsername(&_Contract.CallOpts, username)
}

// GetAddressForUsername is a free data retrieval call binding the contract method 0x095a0535.
//
// Solidity: function getAddressForUsername(string username) view returns(address)
func (_Contract *ContractCallerSession) GetAddressForUsername(username string) (common.Address, error) {
	return _Contract.Contract.GetAddressForUsername(&_Contract.CallOpts, username)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string username) view returns(uint256)
func (_Contract *ContractCaller) GetBalance(opts *bind.CallOpts, username string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBalance", username)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string username) view returns(uint256)
func (_Contract *ContractSession) GetBalance(username string) (*big.Int, error) {
	return _Contract.Contract.GetBalance(&_Contract.CallOpts, username)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string username) view returns(uint256)
func (_Contract *ContractCallerSession) GetBalance(username string) (*big.Int, error) {
	return _Contract.Contract.GetBalance(&_Contract.CallOpts, username)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// TemporaryAddress is a free data retrieval call binding the contract method 0x284c3b3b.
//
// Solidity: function temporaryAddress() view returns(address)
func (_Contract *ContractCaller) TemporaryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "temporaryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TemporaryAddress is a free data retrieval call binding the contract method 0x284c3b3b.
//
// Solidity: function temporaryAddress() view returns(address)
func (_Contract *ContractSession) TemporaryAddress() (common.Address, error) {
	return _Contract.Contract.TemporaryAddress(&_Contract.CallOpts)
}

// TemporaryAddress is a free data retrieval call binding the contract method 0x284c3b3b.
//
// Solidity: function temporaryAddress() view returns(address)
func (_Contract *ContractCallerSession) TemporaryAddress() (common.Address, error) {
	return _Contract.Contract.TemporaryAddress(&_Contract.CallOpts)
}

// UsernameToAddress is a free data retrieval call binding the contract method 0xf825f143.
//
// Solidity: function usernameToAddress(string ) view returns(address)
func (_Contract *ContractCaller) UsernameToAddress(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "usernameToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsernameToAddress is a free data retrieval call binding the contract method 0xf825f143.
//
// Solidity: function usernameToAddress(string ) view returns(address)
func (_Contract *ContractSession) UsernameToAddress(arg0 string) (common.Address, error) {
	return _Contract.Contract.UsernameToAddress(&_Contract.CallOpts, arg0)
}

// UsernameToAddress is a free data retrieval call binding the contract method 0xf825f143.
//
// Solidity: function usernameToAddress(string ) view returns(address)
func (_Contract *ContractCallerSession) UsernameToAddress(arg0 string) (common.Address, error) {
	return _Contract.Contract.UsernameToAddress(&_Contract.CallOpts, arg0)
}

// AssignAddressToUsername is a paid mutator transaction binding the contract method 0xbacd359a.
//
// Solidity: function assignAddressToUsername(string username, address walletAddress) returns()
func (_Contract *ContractTransactor) AssignAddressToUsername(opts *bind.TransactOpts, username string, walletAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "assignAddressToUsername", username, walletAddress)
}

// AssignAddressToUsername is a paid mutator transaction binding the contract method 0xbacd359a.
//
// Solidity: function assignAddressToUsername(string username, address walletAddress) returns()
func (_Contract *ContractSession) AssignAddressToUsername(username string, walletAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AssignAddressToUsername(&_Contract.TransactOpts, username, walletAddress)
}

// AssignAddressToUsername is a paid mutator transaction binding the contract method 0xbacd359a.
//
// Solidity: function assignAddressToUsername(string username, address walletAddress) returns()
func (_Contract *ContractTransactorSession) AssignAddressToUsername(username string, walletAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AssignAddressToUsername(&_Contract.TransactOpts, username, walletAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x877e531e.
//
// Solidity: function deposit(string[] usernames, uint256[] amountsInEther) payable returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, usernames []string, amountsInEther []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", usernames, amountsInEther)
}

// Deposit is a paid mutator transaction binding the contract method 0x877e531e.
//
// Solidity: function deposit(string[] usernames, uint256[] amountsInEther) payable returns()
func (_Contract *ContractSession) Deposit(usernames []string, amountsInEther []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, usernames, amountsInEther)
}

// Deposit is a paid mutator transaction binding the contract method 0x877e531e.
//
// Solidity: function deposit(string[] usernames, uint256[] amountsInEther) payable returns()
func (_Contract *ContractTransactorSession) Deposit(usernames []string, amountsInEther []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, usernames, amountsInEther)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string username, uint256 amountInWei) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, username string, amountInWei *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", username, amountInWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string username, uint256 amountInWei) returns()
func (_Contract *ContractSession) Withdraw(username string, amountInWei *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, username, amountInWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string username, uint256 amountInWei) returns()
func (_Contract *ContractTransactorSession) Withdraw(username string, amountInWei *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, username, amountInWei)
}
