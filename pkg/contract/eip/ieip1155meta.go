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

// Ieip1155metaMetaData contains all meta data concerning the Ieip1155meta contract.
var Ieip1155metaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Ieip1155metaABI is the input ABI used to generate the binding from.
// Deprecated: Use Ieip1155metaMetaData.ABI instead.
var Ieip1155metaABI = Ieip1155metaMetaData.ABI

// Ieip1155meta is an auto generated Go binding around an Ethereum contract.
type Ieip1155meta struct {
	Ieip1155metaCaller     // Read-only binding to the contract
	Ieip1155metaTransactor // Write-only binding to the contract
	Ieip1155metaFilterer   // Log filterer for contract events
}

// Ieip1155metaCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ieip1155metaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ieip1155metaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ieip1155metaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ieip1155metaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ieip1155metaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ieip1155metaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ieip1155metaSession struct {
	Contract     *Ieip1155meta     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ieip1155metaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ieip1155metaCallerSession struct {
	Contract *Ieip1155metaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Ieip1155metaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ieip1155metaTransactorSession struct {
	Contract     *Ieip1155metaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Ieip1155metaRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ieip1155metaRaw struct {
	Contract *Ieip1155meta // Generic contract binding to access the raw methods on
}

// Ieip1155metaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ieip1155metaCallerRaw struct {
	Contract *Ieip1155metaCaller // Generic read-only contract binding to access the raw methods on
}

// Ieip1155metaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ieip1155metaTransactorRaw struct {
	Contract *Ieip1155metaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIeip1155meta creates a new instance of Ieip1155meta, bound to a specific deployed contract.
func NewIeip1155meta(address common.Address, backend bind.ContractBackend) (*Ieip1155meta, error) {
	contract, err := bindIeip1155meta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ieip1155meta{Ieip1155metaCaller: Ieip1155metaCaller{contract: contract}, Ieip1155metaTransactor: Ieip1155metaTransactor{contract: contract}, Ieip1155metaFilterer: Ieip1155metaFilterer{contract: contract}}, nil
}

// NewIeip1155metaCaller creates a new read-only instance of Ieip1155meta, bound to a specific deployed contract.
func NewIeip1155metaCaller(address common.Address, caller bind.ContractCaller) (*Ieip1155metaCaller, error) {
	contract, err := bindIeip1155meta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ieip1155metaCaller{contract: contract}, nil
}

// NewIeip1155metaTransactor creates a new write-only instance of Ieip1155meta, bound to a specific deployed contract.
func NewIeip1155metaTransactor(address common.Address, transactor bind.ContractTransactor) (*Ieip1155metaTransactor, error) {
	contract, err := bindIeip1155meta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ieip1155metaTransactor{contract: contract}, nil
}

// NewIeip1155metaFilterer creates a new log filterer instance of Ieip1155meta, bound to a specific deployed contract.
func NewIeip1155metaFilterer(address common.Address, filterer bind.ContractFilterer) (*Ieip1155metaFilterer, error) {
	contract, err := bindIeip1155meta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ieip1155metaFilterer{contract: contract}, nil
}

// bindIeip1155meta binds a generic wrapper to an already deployed contract.
func bindIeip1155meta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ieip1155metaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ieip1155meta *Ieip1155metaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ieip1155meta.Contract.Ieip1155metaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ieip1155meta *Ieip1155metaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ieip1155meta.Contract.Ieip1155metaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ieip1155meta *Ieip1155metaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ieip1155meta.Contract.Ieip1155metaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ieip1155meta *Ieip1155metaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ieip1155meta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ieip1155meta *Ieip1155metaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ieip1155meta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ieip1155meta *Ieip1155metaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ieip1155meta.Contract.contract.Transact(opts, method, params...)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Ieip1155meta *Ieip1155metaCaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Ieip1155meta.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Ieip1155meta *Ieip1155metaSession) Uri(id *big.Int) (string, error) {
	return _Ieip1155meta.Contract.Uri(&_Ieip1155meta.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Ieip1155meta *Ieip1155metaCallerSession) Uri(id *big.Int) (string, error) {
	return _Ieip1155meta.Contract.Uri(&_Ieip1155meta.CallOpts, id)
}
