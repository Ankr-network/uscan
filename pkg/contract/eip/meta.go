// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eip

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
)

// MetaMetaData contains all meta data concerning the Meta contract.
var MetaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MetaABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaMetaData.ABI instead.
var MetaABI = MetaMetaData.ABI

// Meta is an auto generated Go binding around an Ethereum contract.
type Meta struct {
	MetaCaller     // Read-only binding to the contract
	MetaTransactor // Write-only binding to the contract
	MetaFilterer   // Log filterer for contract events
}

// MetaCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaSession struct {
	Contract     *Meta             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaCallerSession struct {
	Contract *MetaCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MetaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaTransactorSession struct {
	Contract     *MetaTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaRaw struct {
	Contract *Meta // Generic contract binding to access the raw methods on
}

// MetaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaCallerRaw struct {
	Contract *MetaCaller // Generic read-only contract binding to access the raw methods on
}

// MetaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaTransactorRaw struct {
	Contract *MetaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMeta creates a new instance of Meta, bound to a specific deployed contract.
func NewMeta(address common.Address, backend bind.ContractBackend) (*Meta, error) {
	contract, err := bindMeta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Meta{MetaCaller: MetaCaller{contract: contract}, MetaTransactor: MetaTransactor{contract: contract}, MetaFilterer: MetaFilterer{contract: contract}}, nil
}

// NewMetaCaller creates a new read-only instance of Meta, bound to a specific deployed contract.
func NewMetaCaller(address common.Address, caller bind.ContractCaller) (*MetaCaller, error) {
	contract, err := bindMeta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaCaller{contract: contract}, nil
}

// NewMetaTransactor creates a new write-only instance of Meta, bound to a specific deployed contract.
func NewMetaTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaTransactor, error) {
	contract, err := bindMeta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaTransactor{contract: contract}, nil
}

// NewMetaFilterer creates a new log filterer instance of Meta, bound to a specific deployed contract.
func NewMetaFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaFilterer, error) {
	contract, err := bindMeta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaFilterer{contract: contract}, nil
}

// bindMeta binds a generic wrapper to an already deployed contract.
func bindMeta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meta *MetaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meta.Contract.MetaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meta *MetaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meta.Contract.MetaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meta *MetaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meta.Contract.MetaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meta *MetaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meta *MetaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meta *MetaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meta.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Meta *MetaCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Meta *MetaSession) Decimals() (uint8, error) {
	return _Meta.Contract.Decimals(&_Meta.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Meta *MetaCallerSession) Decimals() (uint8, error) {
	return _Meta.Contract.Decimals(&_Meta.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Meta *MetaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Meta *MetaSession) Name() (string, error) {
	return _Meta.Contract.Name(&_Meta.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Meta *MetaCallerSession) Name() (string, error) {
	return _Meta.Contract.Name(&_Meta.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Meta *MetaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Meta *MetaSession) Symbol() (string, error) {
	return _Meta.Contract.Symbol(&_Meta.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Meta *MetaCallerSession) Symbol() (string, error) {
	return _Meta.Contract.Symbol(&_Meta.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Meta *MetaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Meta *MetaSession) TotalSupply() (*big.Int, error) {
	return _Meta.Contract.TotalSupply(&_Meta.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Meta *MetaCallerSession) TotalSupply() (*big.Int, error) {
	return _Meta.Contract.TotalSupply(&_Meta.CallOpts)
}
