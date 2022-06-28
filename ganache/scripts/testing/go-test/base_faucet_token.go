// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// BaseFaucetTokenMetaData contains all meta data concerning the BaseFaucetToken contract.
var BaseFaucetTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"total_supply_whole_tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"faucet_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bridge_address\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"vega_public_keys\",\"type\":\"bytes32[]\"}],\"name\":\"admin_deposit_bulk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bridge_address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"vega_public_key\",\"type\":\"bytes32\"}],\"name\":\"admin_deposit_single\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staking_bridge_address\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"vega_public_keys\",\"type\":\"bytes32[]\"}],\"name\":\"admin_stake_bulk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"faucet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BaseFaucetTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseFaucetTokenMetaData.ABI instead.
var BaseFaucetTokenABI = BaseFaucetTokenMetaData.ABI

// BaseFaucetToken is an auto generated Go binding around an Ethereum contract.
type BaseFaucetToken struct {
	BaseFaucetTokenCaller     // Read-only binding to the contract
	BaseFaucetTokenTransactor // Write-only binding to the contract
	BaseFaucetTokenFilterer   // Log filterer for contract events
}

// BaseFaucetTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseFaucetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFaucetTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseFaucetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFaucetTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseFaucetTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFaucetTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseFaucetTokenSession struct {
	Contract     *BaseFaucetToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseFaucetTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseFaucetTokenCallerSession struct {
	Contract *BaseFaucetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BaseFaucetTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseFaucetTokenTransactorSession struct {
	Contract     *BaseFaucetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BaseFaucetTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseFaucetTokenRaw struct {
	Contract *BaseFaucetToken // Generic contract binding to access the raw methods on
}

// BaseFaucetTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseFaucetTokenCallerRaw struct {
	Contract *BaseFaucetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BaseFaucetTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseFaucetTokenTransactorRaw struct {
	Contract *BaseFaucetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseFaucetToken creates a new instance of BaseFaucetToken, bound to a specific deployed contract.
func NewBaseFaucetToken(address common.Address, backend bind.ContractBackend) (*BaseFaucetToken, error) {
	contract, err := bindBaseFaucetToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseFaucetToken{BaseFaucetTokenCaller: BaseFaucetTokenCaller{contract: contract}, BaseFaucetTokenTransactor: BaseFaucetTokenTransactor{contract: contract}, BaseFaucetTokenFilterer: BaseFaucetTokenFilterer{contract: contract}}, nil
}

// NewBaseFaucetTokenCaller creates a new read-only instance of BaseFaucetToken, bound to a specific deployed contract.
func NewBaseFaucetTokenCaller(address common.Address, caller bind.ContractCaller) (*BaseFaucetTokenCaller, error) {
	contract, err := bindBaseFaucetToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseFaucetTokenCaller{contract: contract}, nil
}

// NewBaseFaucetTokenTransactor creates a new write-only instance of BaseFaucetToken, bound to a specific deployed contract.
func NewBaseFaucetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseFaucetTokenTransactor, error) {
	contract, err := bindBaseFaucetToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseFaucetTokenTransactor{contract: contract}, nil
}

// NewBaseFaucetTokenFilterer creates a new log filterer instance of BaseFaucetToken, bound to a specific deployed contract.
func NewBaseFaucetTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseFaucetTokenFilterer, error) {
	contract, err := bindBaseFaucetToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseFaucetTokenFilterer{contract: contract}, nil
}

// bindBaseFaucetToken binds a generic wrapper to an already deployed contract.
func bindBaseFaucetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseFaucetTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseFaucetToken *BaseFaucetTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseFaucetToken.Contract.BaseFaucetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseFaucetToken *BaseFaucetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.BaseFaucetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseFaucetToken *BaseFaucetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.BaseFaucetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseFaucetToken *BaseFaucetTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseFaucetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseFaucetToken *BaseFaucetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseFaucetToken *BaseFaucetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BaseFaucetToken *BaseFaucetTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseFaucetToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BaseFaucetToken *BaseFaucetTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BaseFaucetToken.Contract.Allowance(&_BaseFaucetToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BaseFaucetToken *BaseFaucetTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BaseFaucetToken.Contract.Allowance(&_BaseFaucetToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BaseFaucetToken *BaseFaucetTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseFaucetToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BaseFaucetToken *BaseFaucetTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BaseFaucetToken.Contract.BalanceOf(&_BaseFaucetToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BaseFaucetToken *BaseFaucetTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BaseFaucetToken.Contract.BalanceOf(&_BaseFaucetToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BaseFaucetToken *BaseFaucetTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BaseFaucetToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BaseFaucetToken *BaseFaucetTokenSession) Decimals() (uint8, error) {
	return _BaseFaucetToken.Contract.Decimals(&_BaseFaucetToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BaseFaucetToken *BaseFaucetTokenCallerSession) Decimals() (uint8, error) {
	return _BaseFaucetToken.Contract.Decimals(&_BaseFaucetToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseFaucetToken.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenSession) IsOwner() (bool, error) {
	return _BaseFaucetToken.Contract.IsOwner(&_BaseFaucetToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenCallerSession) IsOwner() (bool, error) {
	return _BaseFaucetToken.Contract.IsOwner(&_BaseFaucetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BaseFaucetToken *BaseFaucetTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseFaucetToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BaseFaucetToken *BaseFaucetTokenSession) Name() (string, error) {
	return _BaseFaucetToken.Contract.Name(&_BaseFaucetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BaseFaucetToken *BaseFaucetTokenCallerSession) Name() (string, error) {
	return _BaseFaucetToken.Contract.Name(&_BaseFaucetToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseFaucetToken *BaseFaucetTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseFaucetToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseFaucetToken *BaseFaucetTokenSession) Owner() (common.Address, error) {
	return _BaseFaucetToken.Contract.Owner(&_BaseFaucetToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseFaucetToken *BaseFaucetTokenCallerSession) Owner() (common.Address, error) {
	return _BaseFaucetToken.Contract.Owner(&_BaseFaucetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BaseFaucetToken *BaseFaucetTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseFaucetToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BaseFaucetToken *BaseFaucetTokenSession) Symbol() (string, error) {
	return _BaseFaucetToken.Contract.Symbol(&_BaseFaucetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BaseFaucetToken *BaseFaucetTokenCallerSession) Symbol() (string, error) {
	return _BaseFaucetToken.Contract.Symbol(&_BaseFaucetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BaseFaucetToken *BaseFaucetTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseFaucetToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BaseFaucetToken *BaseFaucetTokenSession) TotalSupply() (*big.Int, error) {
	return _BaseFaucetToken.Contract.TotalSupply(&_BaseFaucetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BaseFaucetToken *BaseFaucetTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BaseFaucetToken.Contract.TotalSupply(&_BaseFaucetToken.CallOpts)
}

// AdminDepositBulk is a paid mutator transaction binding the contract method 0xbc36878e.
//
// Solidity: function admin_deposit_bulk(uint256 amount, address bridge_address, bytes32[] vega_public_keys) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactor) AdminDepositBulk(opts *bind.TransactOpts, amount *big.Int, bridge_address common.Address, vega_public_keys [][32]byte) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "admin_deposit_bulk", amount, bridge_address, vega_public_keys)
}

// AdminDepositBulk is a paid mutator transaction binding the contract method 0xbc36878e.
//
// Solidity: function admin_deposit_bulk(uint256 amount, address bridge_address, bytes32[] vega_public_keys) returns()
func (_BaseFaucetToken *BaseFaucetTokenSession) AdminDepositBulk(amount *big.Int, bridge_address common.Address, vega_public_keys [][32]byte) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.AdminDepositBulk(&_BaseFaucetToken.TransactOpts, amount, bridge_address, vega_public_keys)
}

// AdminDepositBulk is a paid mutator transaction binding the contract method 0xbc36878e.
//
// Solidity: function admin_deposit_bulk(uint256 amount, address bridge_address, bytes32[] vega_public_keys) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) AdminDepositBulk(amount *big.Int, bridge_address common.Address, vega_public_keys [][32]byte) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.AdminDepositBulk(&_BaseFaucetToken.TransactOpts, amount, bridge_address, vega_public_keys)
}

// AdminDepositSingle is a paid mutator transaction binding the contract method 0xb777374c.
//
// Solidity: function admin_deposit_single(uint256 amount, address bridge_address, bytes32 vega_public_key) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactor) AdminDepositSingle(opts *bind.TransactOpts, amount *big.Int, bridge_address common.Address, vega_public_key [32]byte) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "admin_deposit_single", amount, bridge_address, vega_public_key)
}

// AdminDepositSingle is a paid mutator transaction binding the contract method 0xb777374c.
//
// Solidity: function admin_deposit_single(uint256 amount, address bridge_address, bytes32 vega_public_key) returns()
func (_BaseFaucetToken *BaseFaucetTokenSession) AdminDepositSingle(amount *big.Int, bridge_address common.Address, vega_public_key [32]byte) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.AdminDepositSingle(&_BaseFaucetToken.TransactOpts, amount, bridge_address, vega_public_key)
}

// AdminDepositSingle is a paid mutator transaction binding the contract method 0xb777374c.
//
// Solidity: function admin_deposit_single(uint256 amount, address bridge_address, bytes32 vega_public_key) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) AdminDepositSingle(amount *big.Int, bridge_address common.Address, vega_public_key [32]byte) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.AdminDepositSingle(&_BaseFaucetToken.TransactOpts, amount, bridge_address, vega_public_key)
}

// AdminStakeBulk is a paid mutator transaction binding the contract method 0xd779cae8.
//
// Solidity: function admin_stake_bulk(uint256 amount, address staking_bridge_address, bytes32[] vega_public_keys) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactor) AdminStakeBulk(opts *bind.TransactOpts, amount *big.Int, staking_bridge_address common.Address, vega_public_keys [][32]byte) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "admin_stake_bulk", amount, staking_bridge_address, vega_public_keys)
}

// AdminStakeBulk is a paid mutator transaction binding the contract method 0xd779cae8.
//
// Solidity: function admin_stake_bulk(uint256 amount, address staking_bridge_address, bytes32[] vega_public_keys) returns()
func (_BaseFaucetToken *BaseFaucetTokenSession) AdminStakeBulk(amount *big.Int, staking_bridge_address common.Address, vega_public_keys [][32]byte) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.AdminStakeBulk(&_BaseFaucetToken.TransactOpts, amount, staking_bridge_address, vega_public_keys)
}

// AdminStakeBulk is a paid mutator transaction binding the contract method 0xd779cae8.
//
// Solidity: function admin_stake_bulk(uint256 amount, address staking_bridge_address, bytes32[] vega_public_keys) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) AdminStakeBulk(amount *big.Int, staking_bridge_address common.Address, vega_public_keys [][32]byte) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.AdminStakeBulk(&_BaseFaucetToken.TransactOpts, amount, staking_bridge_address, vega_public_keys)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Approve(&_BaseFaucetToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Approve(&_BaseFaucetToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.DecreaseAllowance(&_BaseFaucetToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.DecreaseAllowance(&_BaseFaucetToken.TransactOpts, spender, subtractedValue)
}

// Faucet is a paid mutator transaction binding the contract method 0xde5f72fd.
//
// Solidity: function faucet() returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactor) Faucet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "faucet")
}

// Faucet is a paid mutator transaction binding the contract method 0xde5f72fd.
//
// Solidity: function faucet() returns()
func (_BaseFaucetToken *BaseFaucetTokenSession) Faucet() (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Faucet(&_BaseFaucetToken.TransactOpts)
}

// Faucet is a paid mutator transaction binding the contract method 0xde5f72fd.
//
// Solidity: function faucet() returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) Faucet() (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Faucet(&_BaseFaucetToken.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.IncreaseAllowance(&_BaseFaucetToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.IncreaseAllowance(&_BaseFaucetToken.TransactOpts, spender, addedValue)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address account, uint256 value) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactor) Issue(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "issue", account, value)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address account, uint256 value) returns()
func (_BaseFaucetToken *BaseFaucetTokenSession) Issue(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Issue(&_BaseFaucetToken.TransactOpts, account, value)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address account, uint256 value) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) Issue(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Issue(&_BaseFaucetToken.TransactOpts, account, value)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseFaucetToken *BaseFaucetTokenSession) Kill() (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Kill(&_BaseFaucetToken.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) Kill() (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Kill(&_BaseFaucetToken.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BaseFaucetToken *BaseFaucetTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Mint(&_BaseFaucetToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Mint(&_BaseFaucetToken.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseFaucetToken *BaseFaucetTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.RenounceOwnership(&_BaseFaucetToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.RenounceOwnership(&_BaseFaucetToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Transfer(&_BaseFaucetToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.Transfer(&_BaseFaucetToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.TransferFrom(&_BaseFaucetToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.TransferFrom(&_BaseFaucetToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseFaucetToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseFaucetToken *BaseFaucetTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.TransferOwnership(&_BaseFaucetToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseFaucetToken *BaseFaucetTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseFaucetToken.Contract.TransferOwnership(&_BaseFaucetToken.TransactOpts, newOwner)
}

// BaseFaucetTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BaseFaucetToken contract.
type BaseFaucetTokenApprovalIterator struct {
	Event *BaseFaucetTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseFaucetTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseFaucetTokenApproval)
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
		it.Event = new(BaseFaucetTokenApproval)
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
func (it *BaseFaucetTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseFaucetTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseFaucetTokenApproval represents a Approval event raised by the BaseFaucetToken contract.
type BaseFaucetTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BaseFaucetToken *BaseFaucetTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BaseFaucetTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BaseFaucetToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BaseFaucetTokenApprovalIterator{contract: _BaseFaucetToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BaseFaucetToken *BaseFaucetTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BaseFaucetTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BaseFaucetToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseFaucetTokenApproval)
				if err := _BaseFaucetToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BaseFaucetToken *BaseFaucetTokenFilterer) ParseApproval(log types.Log) (*BaseFaucetTokenApproval, error) {
	event := new(BaseFaucetTokenApproval)
	if err := _BaseFaucetToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseFaucetTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaseFaucetToken contract.
type BaseFaucetTokenOwnershipTransferredIterator struct {
	Event *BaseFaucetTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseFaucetTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseFaucetTokenOwnershipTransferred)
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
		it.Event = new(BaseFaucetTokenOwnershipTransferred)
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
func (it *BaseFaucetTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseFaucetTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseFaucetTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BaseFaucetToken contract.
type BaseFaucetTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BaseFaucetToken *BaseFaucetTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaseFaucetTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaseFaucetToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaseFaucetTokenOwnershipTransferredIterator{contract: _BaseFaucetToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BaseFaucetToken *BaseFaucetTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaseFaucetTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaseFaucetToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseFaucetTokenOwnershipTransferred)
				if err := _BaseFaucetToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BaseFaucetToken *BaseFaucetTokenFilterer) ParseOwnershipTransferred(log types.Log) (*BaseFaucetTokenOwnershipTransferred, error) {
	event := new(BaseFaucetTokenOwnershipTransferred)
	if err := _BaseFaucetToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseFaucetTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BaseFaucetToken contract.
type BaseFaucetTokenTransferIterator struct {
	Event *BaseFaucetTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseFaucetTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseFaucetTokenTransfer)
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
		it.Event = new(BaseFaucetTokenTransfer)
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
func (it *BaseFaucetTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseFaucetTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseFaucetTokenTransfer represents a Transfer event raised by the BaseFaucetToken contract.
type BaseFaucetTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BaseFaucetToken *BaseFaucetTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BaseFaucetTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseFaucetToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaseFaucetTokenTransferIterator{contract: _BaseFaucetToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BaseFaucetToken *BaseFaucetTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BaseFaucetTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseFaucetToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseFaucetTokenTransfer)
				if err := _BaseFaucetToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BaseFaucetToken *BaseFaucetTokenFilterer) ParseTransfer(log types.Log) (*BaseFaucetTokenTransfer, error) {
	event := new(BaseFaucetTokenTransfer)
	if err := _BaseFaucetToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
