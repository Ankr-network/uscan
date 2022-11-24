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

// Ieip165MetaData contains all meta data concerning the Ieip165 contract.
var Ieip165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Ieip165ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ieip165MetaData.ABI instead.
var Ieip165ABI = Ieip165MetaData.ABI

// Ieip165 is an auto generated Go binding around an Ethereum contract.
type Ieip165 struct {
	Ieip165Caller     // Read-only binding to the contract
	Ieip165Transactor // Write-only binding to the contract
	Ieip165Filterer   // Log filterer for contract events
}

// Ieip165Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ieip165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ieip165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ieip165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ieip165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ieip165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ieip165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ieip165Session struct {
	Contract     *Ieip165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ieip165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ieip165CallerSession struct {
	Contract *Ieip165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Ieip165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ieip165TransactorSession struct {
	Contract     *Ieip165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Ieip165Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ieip165Raw struct {
	Contract *Ieip165 // Generic contract binding to access the raw methods on
}

// Ieip165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ieip165CallerRaw struct {
	Contract *Ieip165Caller // Generic read-only contract binding to access the raw methods on
}

// Ieip165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ieip165TransactorRaw struct {
	Contract *Ieip165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIeip165 creates a new instance of Ieip165, bound to a specific deployed contract.
func NewIeip165(address common.Address, backend bind.ContractBackend) (*Ieip165, error) {
	contract, err := bindIeip165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ieip165{Ieip165Caller: Ieip165Caller{contract: contract}, Ieip165Transactor: Ieip165Transactor{contract: contract}, Ieip165Filterer: Ieip165Filterer{contract: contract}}, nil
}

// NewIeip165Caller creates a new read-only instance of Ieip165, bound to a specific deployed contract.
func NewIeip165Caller(address common.Address, caller bind.ContractCaller) (*Ieip165Caller, error) {
	contract, err := bindIeip165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ieip165Caller{contract: contract}, nil
}

// NewIeip165Transactor creates a new write-only instance of Ieip165, bound to a specific deployed contract.
func NewIeip165Transactor(address common.Address, transactor bind.ContractTransactor) (*Ieip165Transactor, error) {
	contract, err := bindIeip165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ieip165Transactor{contract: contract}, nil
}

// NewIeip165Filterer creates a new log filterer instance of Ieip165, bound to a specific deployed contract.
func NewIeip165Filterer(address common.Address, filterer bind.ContractFilterer) (*Ieip165Filterer, error) {
	contract, err := bindIeip165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ieip165Filterer{contract: contract}, nil
}

// bindIeip165 binds a generic wrapper to an already deployed contract.
func bindIeip165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ieip165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ieip165 *Ieip165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ieip165.Contract.Ieip165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ieip165 *Ieip165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ieip165.Contract.Ieip165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ieip165 *Ieip165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ieip165.Contract.Ieip165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ieip165 *Ieip165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ieip165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ieip165 *Ieip165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ieip165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ieip165 *Ieip165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ieip165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ieip165 *Ieip165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ieip165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ieip165 *Ieip165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ieip165.Contract.SupportsInterface(&_Ieip165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ieip165 *Ieip165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ieip165.Contract.SupportsInterface(&_Ieip165.CallOpts, interfaceId)
}
