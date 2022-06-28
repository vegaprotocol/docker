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

// BridgeLogicMetaData contains all meta data concerning the BridgeLogic contract.
var BridgeLogicMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"erc20_asset_pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vega_public_key\",\"type\":\"bytes32\"}],\"name\":\"Asset_Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lifetime_limit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdraw_threshold\",\"type\":\"uint256\"}],\"name\":\"Asset_Limits_Updated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"vega_asset_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Asset_Listed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Asset_Removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Asset_Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Bridge_Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Bridge_Stopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdraw_delay\",\"type\":\"uint256\"}],\"name\":\"Bridge_Withdraw_Delay_Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"Depositor_Exempted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"Depositor_Exemption_Revoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"default_withdraw_delay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"vega_public_key\",\"type\":\"bytes32\"}],\"name\":\"deposit_asset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20_asset_pool_address\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exempt_depositor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"}],\"name\":\"get_asset_deposit_lifetime_limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vega_asset_id\",\"type\":\"bytes32\"}],\"name\":\"get_asset_source\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_multisig_control_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"}],\"name\":\"get_vega_asset_id\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"}],\"name\":\"get_withdraw_threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"global_resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"global_stop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"}],\"name\":\"is_asset_listed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"is_exempt_depositor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"is_stopped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"vega_asset_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lifetime_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdraw_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"list_asset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"remove_asset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revoke_exempt_depositor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lifetime_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"set_asset_limits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"set_withdraw_delay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"withdraw_asset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeLogicABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeLogicMetaData.ABI instead.
var BridgeLogicABI = BridgeLogicMetaData.ABI

// BridgeLogic is an auto generated Go binding around an Ethereum contract.
type BridgeLogic struct {
	BridgeLogicCaller     // Read-only binding to the contract
	BridgeLogicTransactor // Write-only binding to the contract
	BridgeLogicFilterer   // Log filterer for contract events
}

// BridgeLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeLogicSession struct {
	Contract     *BridgeLogic      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeLogicCallerSession struct {
	Contract *BridgeLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeLogicTransactorSession struct {
	Contract     *BridgeLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeLogicRaw struct {
	Contract *BridgeLogic // Generic contract binding to access the raw methods on
}

// BridgeLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeLogicCallerRaw struct {
	Contract *BridgeLogicCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeLogicTransactorRaw struct {
	Contract *BridgeLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeLogic creates a new instance of BridgeLogic, bound to a specific deployed contract.
func NewBridgeLogic(address common.Address, backend bind.ContractBackend) (*BridgeLogic, error) {
	contract, err := bindBridgeLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeLogic{BridgeLogicCaller: BridgeLogicCaller{contract: contract}, BridgeLogicTransactor: BridgeLogicTransactor{contract: contract}, BridgeLogicFilterer: BridgeLogicFilterer{contract: contract}}, nil
}

// NewBridgeLogicCaller creates a new read-only instance of BridgeLogic, bound to a specific deployed contract.
func NewBridgeLogicCaller(address common.Address, caller bind.ContractCaller) (*BridgeLogicCaller, error) {
	contract, err := bindBridgeLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicCaller{contract: contract}, nil
}

// NewBridgeLogicTransactor creates a new write-only instance of BridgeLogic, bound to a specific deployed contract.
func NewBridgeLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeLogicTransactor, error) {
	contract, err := bindBridgeLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicTransactor{contract: contract}, nil
}

// NewBridgeLogicFilterer creates a new log filterer instance of BridgeLogic, bound to a specific deployed contract.
func NewBridgeLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeLogicFilterer, error) {
	contract, err := bindBridgeLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicFilterer{contract: contract}, nil
}

// bindBridgeLogic binds a generic wrapper to an already deployed contract.
func bindBridgeLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeLogic *BridgeLogicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeLogic.Contract.BridgeLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeLogic *BridgeLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLogic.Contract.BridgeLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeLogic *BridgeLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeLogic.Contract.BridgeLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeLogic *BridgeLogicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeLogic *BridgeLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeLogic *BridgeLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeLogic.Contract.contract.Transact(opts, method, params...)
}

// DefaultWithdrawDelay is a free data retrieval call binding the contract method 0x3f4f199d.
//
// Solidity: function default_withdraw_delay() view returns(uint256)
func (_BridgeLogic *BridgeLogicCaller) DefaultWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "default_withdraw_delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultWithdrawDelay is a free data retrieval call binding the contract method 0x3f4f199d.
//
// Solidity: function default_withdraw_delay() view returns(uint256)
func (_BridgeLogic *BridgeLogicSession) DefaultWithdrawDelay() (*big.Int, error) {
	return _BridgeLogic.Contract.DefaultWithdrawDelay(&_BridgeLogic.CallOpts)
}

// DefaultWithdrawDelay is a free data retrieval call binding the contract method 0x3f4f199d.
//
// Solidity: function default_withdraw_delay() view returns(uint256)
func (_BridgeLogic *BridgeLogicCallerSession) DefaultWithdrawDelay() (*big.Int, error) {
	return _BridgeLogic.Contract.DefaultWithdrawDelay(&_BridgeLogic.CallOpts)
}

// Erc20AssetPoolAddress is a free data retrieval call binding the contract method 0x9356aab8.
//
// Solidity: function erc20_asset_pool_address() view returns(address)
func (_BridgeLogic *BridgeLogicCaller) Erc20AssetPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "erc20_asset_pool_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20AssetPoolAddress is a free data retrieval call binding the contract method 0x9356aab8.
//
// Solidity: function erc20_asset_pool_address() view returns(address)
func (_BridgeLogic *BridgeLogicSession) Erc20AssetPoolAddress() (common.Address, error) {
	return _BridgeLogic.Contract.Erc20AssetPoolAddress(&_BridgeLogic.CallOpts)
}

// Erc20AssetPoolAddress is a free data retrieval call binding the contract method 0x9356aab8.
//
// Solidity: function erc20_asset_pool_address() view returns(address)
func (_BridgeLogic *BridgeLogicCallerSession) Erc20AssetPoolAddress() (common.Address, error) {
	return _BridgeLogic.Contract.Erc20AssetPoolAddress(&_BridgeLogic.CallOpts)
}

// GetAssetDepositLifetimeLimit is a free data retrieval call binding the contract method 0x354a897a.
//
// Solidity: function get_asset_deposit_lifetime_limit(address asset_source) view returns(uint256)
func (_BridgeLogic *BridgeLogicCaller) GetAssetDepositLifetimeLimit(opts *bind.CallOpts, asset_source common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "get_asset_deposit_lifetime_limit", asset_source)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetDepositLifetimeLimit is a free data retrieval call binding the contract method 0x354a897a.
//
// Solidity: function get_asset_deposit_lifetime_limit(address asset_source) view returns(uint256)
func (_BridgeLogic *BridgeLogicSession) GetAssetDepositLifetimeLimit(asset_source common.Address) (*big.Int, error) {
	return _BridgeLogic.Contract.GetAssetDepositLifetimeLimit(&_BridgeLogic.CallOpts, asset_source)
}

// GetAssetDepositLifetimeLimit is a free data retrieval call binding the contract method 0x354a897a.
//
// Solidity: function get_asset_deposit_lifetime_limit(address asset_source) view returns(uint256)
func (_BridgeLogic *BridgeLogicCallerSession) GetAssetDepositLifetimeLimit(asset_source common.Address) (*big.Int, error) {
	return _BridgeLogic.Contract.GetAssetDepositLifetimeLimit(&_BridgeLogic.CallOpts, asset_source)
}

// GetAssetSource is a free data retrieval call binding the contract method 0x786b0bc0.
//
// Solidity: function get_asset_source(bytes32 vega_asset_id) view returns(address)
func (_BridgeLogic *BridgeLogicCaller) GetAssetSource(opts *bind.CallOpts, vega_asset_id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "get_asset_source", vega_asset_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAssetSource is a free data retrieval call binding the contract method 0x786b0bc0.
//
// Solidity: function get_asset_source(bytes32 vega_asset_id) view returns(address)
func (_BridgeLogic *BridgeLogicSession) GetAssetSource(vega_asset_id [32]byte) (common.Address, error) {
	return _BridgeLogic.Contract.GetAssetSource(&_BridgeLogic.CallOpts, vega_asset_id)
}

// GetAssetSource is a free data retrieval call binding the contract method 0x786b0bc0.
//
// Solidity: function get_asset_source(bytes32 vega_asset_id) view returns(address)
func (_BridgeLogic *BridgeLogicCallerSession) GetAssetSource(vega_asset_id [32]byte) (common.Address, error) {
	return _BridgeLogic.Contract.GetAssetSource(&_BridgeLogic.CallOpts, vega_asset_id)
}

// GetMultisigControlAddress is a free data retrieval call binding the contract method 0xc58dc3b9.
//
// Solidity: function get_multisig_control_address() view returns(address)
func (_BridgeLogic *BridgeLogicCaller) GetMultisigControlAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "get_multisig_control_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMultisigControlAddress is a free data retrieval call binding the contract method 0xc58dc3b9.
//
// Solidity: function get_multisig_control_address() view returns(address)
func (_BridgeLogic *BridgeLogicSession) GetMultisigControlAddress() (common.Address, error) {
	return _BridgeLogic.Contract.GetMultisigControlAddress(&_BridgeLogic.CallOpts)
}

// GetMultisigControlAddress is a free data retrieval call binding the contract method 0xc58dc3b9.
//
// Solidity: function get_multisig_control_address() view returns(address)
func (_BridgeLogic *BridgeLogicCallerSession) GetMultisigControlAddress() (common.Address, error) {
	return _BridgeLogic.Contract.GetMultisigControlAddress(&_BridgeLogic.CallOpts)
}

// GetVegaAssetId is a free data retrieval call binding the contract method 0xa06b5d39.
//
// Solidity: function get_vega_asset_id(address asset_source) view returns(bytes32)
func (_BridgeLogic *BridgeLogicCaller) GetVegaAssetId(opts *bind.CallOpts, asset_source common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "get_vega_asset_id", asset_source)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVegaAssetId is a free data retrieval call binding the contract method 0xa06b5d39.
//
// Solidity: function get_vega_asset_id(address asset_source) view returns(bytes32)
func (_BridgeLogic *BridgeLogicSession) GetVegaAssetId(asset_source common.Address) ([32]byte, error) {
	return _BridgeLogic.Contract.GetVegaAssetId(&_BridgeLogic.CallOpts, asset_source)
}

// GetVegaAssetId is a free data retrieval call binding the contract method 0xa06b5d39.
//
// Solidity: function get_vega_asset_id(address asset_source) view returns(bytes32)
func (_BridgeLogic *BridgeLogicCallerSession) GetVegaAssetId(asset_source common.Address) ([32]byte, error) {
	return _BridgeLogic.Contract.GetVegaAssetId(&_BridgeLogic.CallOpts, asset_source)
}

// GetWithdrawThreshold is a free data retrieval call binding the contract method 0xe8a7bce0.
//
// Solidity: function get_withdraw_threshold(address asset_source) view returns(uint256)
func (_BridgeLogic *BridgeLogicCaller) GetWithdrawThreshold(opts *bind.CallOpts, asset_source common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "get_withdraw_threshold", asset_source)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawThreshold is a free data retrieval call binding the contract method 0xe8a7bce0.
//
// Solidity: function get_withdraw_threshold(address asset_source) view returns(uint256)
func (_BridgeLogic *BridgeLogicSession) GetWithdrawThreshold(asset_source common.Address) (*big.Int, error) {
	return _BridgeLogic.Contract.GetWithdrawThreshold(&_BridgeLogic.CallOpts, asset_source)
}

// GetWithdrawThreshold is a free data retrieval call binding the contract method 0xe8a7bce0.
//
// Solidity: function get_withdraw_threshold(address asset_source) view returns(uint256)
func (_BridgeLogic *BridgeLogicCallerSession) GetWithdrawThreshold(asset_source common.Address) (*big.Int, error) {
	return _BridgeLogic.Contract.GetWithdrawThreshold(&_BridgeLogic.CallOpts, asset_source)
}

// IsAssetListed is a free data retrieval call binding the contract method 0x7fd27b7f.
//
// Solidity: function is_asset_listed(address asset_source) view returns(bool)
func (_BridgeLogic *BridgeLogicCaller) IsAssetListed(opts *bind.CallOpts, asset_source common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "is_asset_listed", asset_source)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetListed is a free data retrieval call binding the contract method 0x7fd27b7f.
//
// Solidity: function is_asset_listed(address asset_source) view returns(bool)
func (_BridgeLogic *BridgeLogicSession) IsAssetListed(asset_source common.Address) (bool, error) {
	return _BridgeLogic.Contract.IsAssetListed(&_BridgeLogic.CallOpts, asset_source)
}

// IsAssetListed is a free data retrieval call binding the contract method 0x7fd27b7f.
//
// Solidity: function is_asset_listed(address asset_source) view returns(bool)
func (_BridgeLogic *BridgeLogicCallerSession) IsAssetListed(asset_source common.Address) (bool, error) {
	return _BridgeLogic.Contract.IsAssetListed(&_BridgeLogic.CallOpts, asset_source)
}

// IsExemptDepositor is a free data retrieval call binding the contract method 0x15c0df9d.
//
// Solidity: function is_exempt_depositor(address depositor) view returns(bool)
func (_BridgeLogic *BridgeLogicCaller) IsExemptDepositor(opts *bind.CallOpts, depositor common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "is_exempt_depositor", depositor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExemptDepositor is a free data retrieval call binding the contract method 0x15c0df9d.
//
// Solidity: function is_exempt_depositor(address depositor) view returns(bool)
func (_BridgeLogic *BridgeLogicSession) IsExemptDepositor(depositor common.Address) (bool, error) {
	return _BridgeLogic.Contract.IsExemptDepositor(&_BridgeLogic.CallOpts, depositor)
}

// IsExemptDepositor is a free data retrieval call binding the contract method 0x15c0df9d.
//
// Solidity: function is_exempt_depositor(address depositor) view returns(bool)
func (_BridgeLogic *BridgeLogicCallerSession) IsExemptDepositor(depositor common.Address) (bool, error) {
	return _BridgeLogic.Contract.IsExemptDepositor(&_BridgeLogic.CallOpts, depositor)
}

// IsStopped is a free data retrieval call binding the contract method 0xe272e9d0.
//
// Solidity: function is_stopped() view returns(bool)
func (_BridgeLogic *BridgeLogicCaller) IsStopped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeLogic.contract.Call(opts, &out, "is_stopped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStopped is a free data retrieval call binding the contract method 0xe272e9d0.
//
// Solidity: function is_stopped() view returns(bool)
func (_BridgeLogic *BridgeLogicSession) IsStopped() (bool, error) {
	return _BridgeLogic.Contract.IsStopped(&_BridgeLogic.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0xe272e9d0.
//
// Solidity: function is_stopped() view returns(bool)
func (_BridgeLogic *BridgeLogicCallerSession) IsStopped() (bool, error) {
	return _BridgeLogic.Contract.IsStopped(&_BridgeLogic.CallOpts)
}

// DepositAsset is a paid mutator transaction binding the contract method 0xf7683932.
//
// Solidity: function deposit_asset(address asset_source, uint256 amount, bytes32 vega_public_key) returns()
func (_BridgeLogic *BridgeLogicTransactor) DepositAsset(opts *bind.TransactOpts, asset_source common.Address, amount *big.Int, vega_public_key [32]byte) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "deposit_asset", asset_source, amount, vega_public_key)
}

// DepositAsset is a paid mutator transaction binding the contract method 0xf7683932.
//
// Solidity: function deposit_asset(address asset_source, uint256 amount, bytes32 vega_public_key) returns()
func (_BridgeLogic *BridgeLogicSession) DepositAsset(asset_source common.Address, amount *big.Int, vega_public_key [32]byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.DepositAsset(&_BridgeLogic.TransactOpts, asset_source, amount, vega_public_key)
}

// DepositAsset is a paid mutator transaction binding the contract method 0xf7683932.
//
// Solidity: function deposit_asset(address asset_source, uint256 amount, bytes32 vega_public_key) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) DepositAsset(asset_source common.Address, amount *big.Int, vega_public_key [32]byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.DepositAsset(&_BridgeLogic.TransactOpts, asset_source, amount, vega_public_key)
}

// ExemptDepositor is a paid mutator transaction binding the contract method 0xb76fbb75.
//
// Solidity: function exempt_depositor() returns()
func (_BridgeLogic *BridgeLogicTransactor) ExemptDepositor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "exempt_depositor")
}

// ExemptDepositor is a paid mutator transaction binding the contract method 0xb76fbb75.
//
// Solidity: function exempt_depositor() returns()
func (_BridgeLogic *BridgeLogicSession) ExemptDepositor() (*types.Transaction, error) {
	return _BridgeLogic.Contract.ExemptDepositor(&_BridgeLogic.TransactOpts)
}

// ExemptDepositor is a paid mutator transaction binding the contract method 0xb76fbb75.
//
// Solidity: function exempt_depositor() returns()
func (_BridgeLogic *BridgeLogicTransactorSession) ExemptDepositor() (*types.Transaction, error) {
	return _BridgeLogic.Contract.ExemptDepositor(&_BridgeLogic.TransactOpts)
}

// GlobalResume is a paid mutator transaction binding the contract method 0xd72ed529.
//
// Solidity: function global_resume(uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactor) GlobalResume(opts *bind.TransactOpts, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "global_resume", nonce, signatures)
}

// GlobalResume is a paid mutator transaction binding the contract method 0xd72ed529.
//
// Solidity: function global_resume(uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicSession) GlobalResume(nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.GlobalResume(&_BridgeLogic.TransactOpts, nonce, signatures)
}

// GlobalResume is a paid mutator transaction binding the contract method 0xd72ed529.
//
// Solidity: function global_resume(uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) GlobalResume(nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.GlobalResume(&_BridgeLogic.TransactOpts, nonce, signatures)
}

// GlobalStop is a paid mutator transaction binding the contract method 0x9dfd3c88.
//
// Solidity: function global_stop(uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactor) GlobalStop(opts *bind.TransactOpts, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "global_stop", nonce, signatures)
}

// GlobalStop is a paid mutator transaction binding the contract method 0x9dfd3c88.
//
// Solidity: function global_stop(uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicSession) GlobalStop(nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.GlobalStop(&_BridgeLogic.TransactOpts, nonce, signatures)
}

// GlobalStop is a paid mutator transaction binding the contract method 0x9dfd3c88.
//
// Solidity: function global_stop(uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) GlobalStop(nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.GlobalStop(&_BridgeLogic.TransactOpts, nonce, signatures)
}

// ListAsset is a paid mutator transaction binding the contract method 0x0ff3562c.
//
// Solidity: function list_asset(address asset_source, bytes32 vega_asset_id, uint256 lifetime_limit, uint256 withdraw_threshold, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactor) ListAsset(opts *bind.TransactOpts, asset_source common.Address, vega_asset_id [32]byte, lifetime_limit *big.Int, withdraw_threshold *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "list_asset", asset_source, vega_asset_id, lifetime_limit, withdraw_threshold, nonce, signatures)
}

// ListAsset is a paid mutator transaction binding the contract method 0x0ff3562c.
//
// Solidity: function list_asset(address asset_source, bytes32 vega_asset_id, uint256 lifetime_limit, uint256 withdraw_threshold, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicSession) ListAsset(asset_source common.Address, vega_asset_id [32]byte, lifetime_limit *big.Int, withdraw_threshold *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.ListAsset(&_BridgeLogic.TransactOpts, asset_source, vega_asset_id, lifetime_limit, withdraw_threshold, nonce, signatures)
}

// ListAsset is a paid mutator transaction binding the contract method 0x0ff3562c.
//
// Solidity: function list_asset(address asset_source, bytes32 vega_asset_id, uint256 lifetime_limit, uint256 withdraw_threshold, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) ListAsset(asset_source common.Address, vega_asset_id [32]byte, lifetime_limit *big.Int, withdraw_threshold *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.ListAsset(&_BridgeLogic.TransactOpts, asset_source, vega_asset_id, lifetime_limit, withdraw_threshold, nonce, signatures)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0xc76de358.
//
// Solidity: function remove_asset(address asset_source, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactor) RemoveAsset(opts *bind.TransactOpts, asset_source common.Address, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "remove_asset", asset_source, nonce, signatures)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0xc76de358.
//
// Solidity: function remove_asset(address asset_source, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicSession) RemoveAsset(asset_source common.Address, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.RemoveAsset(&_BridgeLogic.TransactOpts, asset_source, nonce, signatures)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0xc76de358.
//
// Solidity: function remove_asset(address asset_source, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) RemoveAsset(asset_source common.Address, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.RemoveAsset(&_BridgeLogic.TransactOpts, asset_source, nonce, signatures)
}

// RevokeExemptDepositor is a paid mutator transaction binding the contract method 0x6a1c6fa4.
//
// Solidity: function revoke_exempt_depositor() returns()
func (_BridgeLogic *BridgeLogicTransactor) RevokeExemptDepositor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "revoke_exempt_depositor")
}

// RevokeExemptDepositor is a paid mutator transaction binding the contract method 0x6a1c6fa4.
//
// Solidity: function revoke_exempt_depositor() returns()
func (_BridgeLogic *BridgeLogicSession) RevokeExemptDepositor() (*types.Transaction, error) {
	return _BridgeLogic.Contract.RevokeExemptDepositor(&_BridgeLogic.TransactOpts)
}

// RevokeExemptDepositor is a paid mutator transaction binding the contract method 0x6a1c6fa4.
//
// Solidity: function revoke_exempt_depositor() returns()
func (_BridgeLogic *BridgeLogicTransactorSession) RevokeExemptDepositor() (*types.Transaction, error) {
	return _BridgeLogic.Contract.RevokeExemptDepositor(&_BridgeLogic.TransactOpts)
}

// SetAssetLimits is a paid mutator transaction binding the contract method 0x41fb776d.
//
// Solidity: function set_asset_limits(address asset_source, uint256 lifetime_limit, uint256 threshold, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactor) SetAssetLimits(opts *bind.TransactOpts, asset_source common.Address, lifetime_limit *big.Int, threshold *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "set_asset_limits", asset_source, lifetime_limit, threshold, nonce, signatures)
}

// SetAssetLimits is a paid mutator transaction binding the contract method 0x41fb776d.
//
// Solidity: function set_asset_limits(address asset_source, uint256 lifetime_limit, uint256 threshold, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicSession) SetAssetLimits(asset_source common.Address, lifetime_limit *big.Int, threshold *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetAssetLimits(&_BridgeLogic.TransactOpts, asset_source, lifetime_limit, threshold, nonce, signatures)
}

// SetAssetLimits is a paid mutator transaction binding the contract method 0x41fb776d.
//
// Solidity: function set_asset_limits(address asset_source, uint256 lifetime_limit, uint256 threshold, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) SetAssetLimits(asset_source common.Address, lifetime_limit *big.Int, threshold *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetAssetLimits(&_BridgeLogic.TransactOpts, asset_source, lifetime_limit, threshold, nonce, signatures)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0x5a246728.
//
// Solidity: function set_withdraw_delay(uint256 delay, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactor) SetWithdrawDelay(opts *bind.TransactOpts, delay *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "set_withdraw_delay", delay, nonce, signatures)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0x5a246728.
//
// Solidity: function set_withdraw_delay(uint256 delay, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicSession) SetWithdrawDelay(delay *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetWithdrawDelay(&_BridgeLogic.TransactOpts, delay, nonce, signatures)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0x5a246728.
//
// Solidity: function set_withdraw_delay(uint256 delay, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) SetWithdrawDelay(delay *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetWithdrawDelay(&_BridgeLogic.TransactOpts, delay, nonce, signatures)
}

// WithdrawAsset is a paid mutator transaction binding the contract method 0x3ad90635.
//
// Solidity: function withdraw_asset(address asset_source, uint256 amount, address target, uint256 creation, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactor) WithdrawAsset(opts *bind.TransactOpts, asset_source common.Address, amount *big.Int, target common.Address, creation *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "withdraw_asset", asset_source, amount, target, creation, nonce, signatures)
}

// WithdrawAsset is a paid mutator transaction binding the contract method 0x3ad90635.
//
// Solidity: function withdraw_asset(address asset_source, uint256 amount, address target, uint256 creation, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicSession) WithdrawAsset(asset_source common.Address, amount *big.Int, target common.Address, creation *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.WithdrawAsset(&_BridgeLogic.TransactOpts, asset_source, amount, target, creation, nonce, signatures)
}

// WithdrawAsset is a paid mutator transaction binding the contract method 0x3ad90635.
//
// Solidity: function withdraw_asset(address asset_source, uint256 amount, address target, uint256 creation, uint256 nonce, bytes signatures) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) WithdrawAsset(asset_source common.Address, amount *big.Int, target common.Address, creation *big.Int, nonce *big.Int, signatures []byte) (*types.Transaction, error) {
	return _BridgeLogic.Contract.WithdrawAsset(&_BridgeLogic.TransactOpts, asset_source, amount, target, creation, nonce, signatures)
}

// BridgeLogicAssetDepositedIterator is returned from FilterAssetDeposited and is used to iterate over the raw logs and unpacked data for AssetDeposited events raised by the BridgeLogic contract.
type BridgeLogicAssetDepositedIterator struct {
	Event *BridgeLogicAssetDeposited // Event containing the contract specifics and raw log

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
func (it *BridgeLogicAssetDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicAssetDeposited)
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
		it.Event = new(BridgeLogicAssetDeposited)
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
func (it *BridgeLogicAssetDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicAssetDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicAssetDeposited represents a AssetDeposited event raised by the BridgeLogic contract.
type BridgeLogicAssetDeposited struct {
	UserAddress   common.Address
	AssetSource   common.Address
	Amount        *big.Int
	VegaPublicKey [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAssetDeposited is a free log retrieval operation binding the contract event 0x3724ff5e82ddc640a08d68b0b782a5991aea0de51a8dd10a59cdbe5b3ec4e6bf.
//
// Solidity: event Asset_Deposited(address indexed user_address, address indexed asset_source, uint256 amount, bytes32 vega_public_key)
func (_BridgeLogic *BridgeLogicFilterer) FilterAssetDeposited(opts *bind.FilterOpts, user_address []common.Address, asset_source []common.Address) (*BridgeLogicAssetDepositedIterator, error) {

	var user_addressRule []interface{}
	for _, user_addressItem := range user_address {
		user_addressRule = append(user_addressRule, user_addressItem)
	}
	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Asset_Deposited", user_addressRule, asset_sourceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicAssetDepositedIterator{contract: _BridgeLogic.contract, event: "Asset_Deposited", logs: logs, sub: sub}, nil
}

// WatchAssetDeposited is a free log subscription operation binding the contract event 0x3724ff5e82ddc640a08d68b0b782a5991aea0de51a8dd10a59cdbe5b3ec4e6bf.
//
// Solidity: event Asset_Deposited(address indexed user_address, address indexed asset_source, uint256 amount, bytes32 vega_public_key)
func (_BridgeLogic *BridgeLogicFilterer) WatchAssetDeposited(opts *bind.WatchOpts, sink chan<- *BridgeLogicAssetDeposited, user_address []common.Address, asset_source []common.Address) (event.Subscription, error) {

	var user_addressRule []interface{}
	for _, user_addressItem := range user_address {
		user_addressRule = append(user_addressRule, user_addressItem)
	}
	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Asset_Deposited", user_addressRule, asset_sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicAssetDeposited)
				if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Deposited", log); err != nil {
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

// ParseAssetDeposited is a log parse operation binding the contract event 0x3724ff5e82ddc640a08d68b0b782a5991aea0de51a8dd10a59cdbe5b3ec4e6bf.
//
// Solidity: event Asset_Deposited(address indexed user_address, address indexed asset_source, uint256 amount, bytes32 vega_public_key)
func (_BridgeLogic *BridgeLogicFilterer) ParseAssetDeposited(log types.Log) (*BridgeLogicAssetDeposited, error) {
	event := new(BridgeLogicAssetDeposited)
	if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogicAssetLimitsUpdatedIterator is returned from FilterAssetLimitsUpdated and is used to iterate over the raw logs and unpacked data for AssetLimitsUpdated events raised by the BridgeLogic contract.
type BridgeLogicAssetLimitsUpdatedIterator struct {
	Event *BridgeLogicAssetLimitsUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeLogicAssetLimitsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicAssetLimitsUpdated)
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
		it.Event = new(BridgeLogicAssetLimitsUpdated)
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
func (it *BridgeLogicAssetLimitsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicAssetLimitsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicAssetLimitsUpdated represents a AssetLimitsUpdated event raised by the BridgeLogic contract.
type BridgeLogicAssetLimitsUpdated struct {
	AssetSource       common.Address
	LifetimeLimit     *big.Int
	WithdrawThreshold *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAssetLimitsUpdated is a free log retrieval operation binding the contract event 0xfc7eab762b8751ad85c101fd1025c763b4e8d48f2093f506629b606618e884fe.
//
// Solidity: event Asset_Limits_Updated(address indexed asset_source, uint256 lifetime_limit, uint256 withdraw_threshold)
func (_BridgeLogic *BridgeLogicFilterer) FilterAssetLimitsUpdated(opts *bind.FilterOpts, asset_source []common.Address) (*BridgeLogicAssetLimitsUpdatedIterator, error) {

	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Asset_Limits_Updated", asset_sourceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicAssetLimitsUpdatedIterator{contract: _BridgeLogic.contract, event: "Asset_Limits_Updated", logs: logs, sub: sub}, nil
}

// WatchAssetLimitsUpdated is a free log subscription operation binding the contract event 0xfc7eab762b8751ad85c101fd1025c763b4e8d48f2093f506629b606618e884fe.
//
// Solidity: event Asset_Limits_Updated(address indexed asset_source, uint256 lifetime_limit, uint256 withdraw_threshold)
func (_BridgeLogic *BridgeLogicFilterer) WatchAssetLimitsUpdated(opts *bind.WatchOpts, sink chan<- *BridgeLogicAssetLimitsUpdated, asset_source []common.Address) (event.Subscription, error) {

	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Asset_Limits_Updated", asset_sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicAssetLimitsUpdated)
				if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Limits_Updated", log); err != nil {
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

// ParseAssetLimitsUpdated is a log parse operation binding the contract event 0xfc7eab762b8751ad85c101fd1025c763b4e8d48f2093f506629b606618e884fe.
//
// Solidity: event Asset_Limits_Updated(address indexed asset_source, uint256 lifetime_limit, uint256 withdraw_threshold)
func (_BridgeLogic *BridgeLogicFilterer) ParseAssetLimitsUpdated(log types.Log) (*BridgeLogicAssetLimitsUpdated, error) {
	event := new(BridgeLogicAssetLimitsUpdated)
	if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Limits_Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogicAssetListedIterator is returned from FilterAssetListed and is used to iterate over the raw logs and unpacked data for AssetListed events raised by the BridgeLogic contract.
type BridgeLogicAssetListedIterator struct {
	Event *BridgeLogicAssetListed // Event containing the contract specifics and raw log

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
func (it *BridgeLogicAssetListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicAssetListed)
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
		it.Event = new(BridgeLogicAssetListed)
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
func (it *BridgeLogicAssetListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicAssetListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicAssetListed represents a AssetListed event raised by the BridgeLogic contract.
type BridgeLogicAssetListed struct {
	AssetSource common.Address
	VegaAssetId [32]byte
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetListed is a free log retrieval operation binding the contract event 0x4180d77d05ff0d31650c548c23f2de07a3da3ad42e3dd6edd817b438a150452e.
//
// Solidity: event Asset_Listed(address indexed asset_source, bytes32 indexed vega_asset_id, uint256 nonce)
func (_BridgeLogic *BridgeLogicFilterer) FilterAssetListed(opts *bind.FilterOpts, asset_source []common.Address, vega_asset_id [][32]byte) (*BridgeLogicAssetListedIterator, error) {

	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}
	var vega_asset_idRule []interface{}
	for _, vega_asset_idItem := range vega_asset_id {
		vega_asset_idRule = append(vega_asset_idRule, vega_asset_idItem)
	}

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Asset_Listed", asset_sourceRule, vega_asset_idRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicAssetListedIterator{contract: _BridgeLogic.contract, event: "Asset_Listed", logs: logs, sub: sub}, nil
}

// WatchAssetListed is a free log subscription operation binding the contract event 0x4180d77d05ff0d31650c548c23f2de07a3da3ad42e3dd6edd817b438a150452e.
//
// Solidity: event Asset_Listed(address indexed asset_source, bytes32 indexed vega_asset_id, uint256 nonce)
func (_BridgeLogic *BridgeLogicFilterer) WatchAssetListed(opts *bind.WatchOpts, sink chan<- *BridgeLogicAssetListed, asset_source []common.Address, vega_asset_id [][32]byte) (event.Subscription, error) {

	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}
	var vega_asset_idRule []interface{}
	for _, vega_asset_idItem := range vega_asset_id {
		vega_asset_idRule = append(vega_asset_idRule, vega_asset_idItem)
	}

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Asset_Listed", asset_sourceRule, vega_asset_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicAssetListed)
				if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Listed", log); err != nil {
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

// ParseAssetListed is a log parse operation binding the contract event 0x4180d77d05ff0d31650c548c23f2de07a3da3ad42e3dd6edd817b438a150452e.
//
// Solidity: event Asset_Listed(address indexed asset_source, bytes32 indexed vega_asset_id, uint256 nonce)
func (_BridgeLogic *BridgeLogicFilterer) ParseAssetListed(log types.Log) (*BridgeLogicAssetListed, error) {
	event := new(BridgeLogicAssetListed)
	if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Listed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogicAssetRemovedIterator is returned from FilterAssetRemoved and is used to iterate over the raw logs and unpacked data for AssetRemoved events raised by the BridgeLogic contract.
type BridgeLogicAssetRemovedIterator struct {
	Event *BridgeLogicAssetRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeLogicAssetRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicAssetRemoved)
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
		it.Event = new(BridgeLogicAssetRemoved)
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
func (it *BridgeLogicAssetRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicAssetRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicAssetRemoved represents a AssetRemoved event raised by the BridgeLogic contract.
type BridgeLogicAssetRemoved struct {
	AssetSource common.Address
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetRemoved is a free log retrieval operation binding the contract event 0x58ad5e799e2df93ab408be0e5c1870d44c80b5bca99dfaf7ddf0dab5e6b155c9.
//
// Solidity: event Asset_Removed(address indexed asset_source, uint256 nonce)
func (_BridgeLogic *BridgeLogicFilterer) FilterAssetRemoved(opts *bind.FilterOpts, asset_source []common.Address) (*BridgeLogicAssetRemovedIterator, error) {

	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Asset_Removed", asset_sourceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicAssetRemovedIterator{contract: _BridgeLogic.contract, event: "Asset_Removed", logs: logs, sub: sub}, nil
}

// WatchAssetRemoved is a free log subscription operation binding the contract event 0x58ad5e799e2df93ab408be0e5c1870d44c80b5bca99dfaf7ddf0dab5e6b155c9.
//
// Solidity: event Asset_Removed(address indexed asset_source, uint256 nonce)
func (_BridgeLogic *BridgeLogicFilterer) WatchAssetRemoved(opts *bind.WatchOpts, sink chan<- *BridgeLogicAssetRemoved, asset_source []common.Address) (event.Subscription, error) {

	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Asset_Removed", asset_sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicAssetRemoved)
				if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Removed", log); err != nil {
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

// ParseAssetRemoved is a log parse operation binding the contract event 0x58ad5e799e2df93ab408be0e5c1870d44c80b5bca99dfaf7ddf0dab5e6b155c9.
//
// Solidity: event Asset_Removed(address indexed asset_source, uint256 nonce)
func (_BridgeLogic *BridgeLogicFilterer) ParseAssetRemoved(log types.Log) (*BridgeLogicAssetRemoved, error) {
	event := new(BridgeLogicAssetRemoved)
	if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Removed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogicAssetWithdrawnIterator is returned from FilterAssetWithdrawn and is used to iterate over the raw logs and unpacked data for AssetWithdrawn events raised by the BridgeLogic contract.
type BridgeLogicAssetWithdrawnIterator struct {
	Event *BridgeLogicAssetWithdrawn // Event containing the contract specifics and raw log

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
func (it *BridgeLogicAssetWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicAssetWithdrawn)
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
		it.Event = new(BridgeLogicAssetWithdrawn)
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
func (it *BridgeLogicAssetWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicAssetWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicAssetWithdrawn represents a AssetWithdrawn event raised by the BridgeLogic contract.
type BridgeLogicAssetWithdrawn struct {
	UserAddress common.Address
	AssetSource common.Address
	Amount      *big.Int
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetWithdrawn is a free log retrieval operation binding the contract event 0xa79be4f3361e32d396d64c478ecef73732cb40b2a75702c3b3b3226a2c83b5df.
//
// Solidity: event Asset_Withdrawn(address indexed user_address, address indexed asset_source, uint256 amount, uint256 nonce)
func (_BridgeLogic *BridgeLogicFilterer) FilterAssetWithdrawn(opts *bind.FilterOpts, user_address []common.Address, asset_source []common.Address) (*BridgeLogicAssetWithdrawnIterator, error) {

	var user_addressRule []interface{}
	for _, user_addressItem := range user_address {
		user_addressRule = append(user_addressRule, user_addressItem)
	}
	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Asset_Withdrawn", user_addressRule, asset_sourceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicAssetWithdrawnIterator{contract: _BridgeLogic.contract, event: "Asset_Withdrawn", logs: logs, sub: sub}, nil
}

// WatchAssetWithdrawn is a free log subscription operation binding the contract event 0xa79be4f3361e32d396d64c478ecef73732cb40b2a75702c3b3b3226a2c83b5df.
//
// Solidity: event Asset_Withdrawn(address indexed user_address, address indexed asset_source, uint256 amount, uint256 nonce)
func (_BridgeLogic *BridgeLogicFilterer) WatchAssetWithdrawn(opts *bind.WatchOpts, sink chan<- *BridgeLogicAssetWithdrawn, user_address []common.Address, asset_source []common.Address) (event.Subscription, error) {

	var user_addressRule []interface{}
	for _, user_addressItem := range user_address {
		user_addressRule = append(user_addressRule, user_addressItem)
	}
	var asset_sourceRule []interface{}
	for _, asset_sourceItem := range asset_source {
		asset_sourceRule = append(asset_sourceRule, asset_sourceItem)
	}

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Asset_Withdrawn", user_addressRule, asset_sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicAssetWithdrawn)
				if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Withdrawn", log); err != nil {
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

// ParseAssetWithdrawn is a log parse operation binding the contract event 0xa79be4f3361e32d396d64c478ecef73732cb40b2a75702c3b3b3226a2c83b5df.
//
// Solidity: event Asset_Withdrawn(address indexed user_address, address indexed asset_source, uint256 amount, uint256 nonce)
func (_BridgeLogic *BridgeLogicFilterer) ParseAssetWithdrawn(log types.Log) (*BridgeLogicAssetWithdrawn, error) {
	event := new(BridgeLogicAssetWithdrawn)
	if err := _BridgeLogic.contract.UnpackLog(event, "Asset_Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogicBridgeResumedIterator is returned from FilterBridgeResumed and is used to iterate over the raw logs and unpacked data for BridgeResumed events raised by the BridgeLogic contract.
type BridgeLogicBridgeResumedIterator struct {
	Event *BridgeLogicBridgeResumed // Event containing the contract specifics and raw log

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
func (it *BridgeLogicBridgeResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicBridgeResumed)
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
		it.Event = new(BridgeLogicBridgeResumed)
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
func (it *BridgeLogicBridgeResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicBridgeResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicBridgeResumed represents a BridgeResumed event raised by the BridgeLogic contract.
type BridgeLogicBridgeResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBridgeResumed is a free log retrieval operation binding the contract event 0x79c02b0e60e0f00fe0370791204f2f175fe3f06f4816f3506ad4fa1b8e8cde0f.
//
// Solidity: event Bridge_Resumed()
func (_BridgeLogic *BridgeLogicFilterer) FilterBridgeResumed(opts *bind.FilterOpts) (*BridgeLogicBridgeResumedIterator, error) {

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Bridge_Resumed")
	if err != nil {
		return nil, err
	}
	return &BridgeLogicBridgeResumedIterator{contract: _BridgeLogic.contract, event: "Bridge_Resumed", logs: logs, sub: sub}, nil
}

// WatchBridgeResumed is a free log subscription operation binding the contract event 0x79c02b0e60e0f00fe0370791204f2f175fe3f06f4816f3506ad4fa1b8e8cde0f.
//
// Solidity: event Bridge_Resumed()
func (_BridgeLogic *BridgeLogicFilterer) WatchBridgeResumed(opts *bind.WatchOpts, sink chan<- *BridgeLogicBridgeResumed) (event.Subscription, error) {

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Bridge_Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicBridgeResumed)
				if err := _BridgeLogic.contract.UnpackLog(event, "Bridge_Resumed", log); err != nil {
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

// ParseBridgeResumed is a log parse operation binding the contract event 0x79c02b0e60e0f00fe0370791204f2f175fe3f06f4816f3506ad4fa1b8e8cde0f.
//
// Solidity: event Bridge_Resumed()
func (_BridgeLogic *BridgeLogicFilterer) ParseBridgeResumed(log types.Log) (*BridgeLogicBridgeResumed, error) {
	event := new(BridgeLogicBridgeResumed)
	if err := _BridgeLogic.contract.UnpackLog(event, "Bridge_Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogicBridgeStoppedIterator is returned from FilterBridgeStopped and is used to iterate over the raw logs and unpacked data for BridgeStopped events raised by the BridgeLogic contract.
type BridgeLogicBridgeStoppedIterator struct {
	Event *BridgeLogicBridgeStopped // Event containing the contract specifics and raw log

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
func (it *BridgeLogicBridgeStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicBridgeStopped)
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
		it.Event = new(BridgeLogicBridgeStopped)
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
func (it *BridgeLogicBridgeStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicBridgeStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicBridgeStopped represents a BridgeStopped event raised by the BridgeLogic contract.
type BridgeLogicBridgeStopped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBridgeStopped is a free log retrieval operation binding the contract event 0x129d99581c8e70519df1f0733d3212f33d0ed3ea6144adacc336c647f1d36382.
//
// Solidity: event Bridge_Stopped()
func (_BridgeLogic *BridgeLogicFilterer) FilterBridgeStopped(opts *bind.FilterOpts) (*BridgeLogicBridgeStoppedIterator, error) {

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Bridge_Stopped")
	if err != nil {
		return nil, err
	}
	return &BridgeLogicBridgeStoppedIterator{contract: _BridgeLogic.contract, event: "Bridge_Stopped", logs: logs, sub: sub}, nil
}

// WatchBridgeStopped is a free log subscription operation binding the contract event 0x129d99581c8e70519df1f0733d3212f33d0ed3ea6144adacc336c647f1d36382.
//
// Solidity: event Bridge_Stopped()
func (_BridgeLogic *BridgeLogicFilterer) WatchBridgeStopped(opts *bind.WatchOpts, sink chan<- *BridgeLogicBridgeStopped) (event.Subscription, error) {

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Bridge_Stopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicBridgeStopped)
				if err := _BridgeLogic.contract.UnpackLog(event, "Bridge_Stopped", log); err != nil {
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

// ParseBridgeStopped is a log parse operation binding the contract event 0x129d99581c8e70519df1f0733d3212f33d0ed3ea6144adacc336c647f1d36382.
//
// Solidity: event Bridge_Stopped()
func (_BridgeLogic *BridgeLogicFilterer) ParseBridgeStopped(log types.Log) (*BridgeLogicBridgeStopped, error) {
	event := new(BridgeLogicBridgeStopped)
	if err := _BridgeLogic.contract.UnpackLog(event, "Bridge_Stopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogicBridgeWithdrawDelaySetIterator is returned from FilterBridgeWithdrawDelaySet and is used to iterate over the raw logs and unpacked data for BridgeWithdrawDelaySet events raised by the BridgeLogic contract.
type BridgeLogicBridgeWithdrawDelaySetIterator struct {
	Event *BridgeLogicBridgeWithdrawDelaySet // Event containing the contract specifics and raw log

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
func (it *BridgeLogicBridgeWithdrawDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicBridgeWithdrawDelaySet)
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
		it.Event = new(BridgeLogicBridgeWithdrawDelaySet)
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
func (it *BridgeLogicBridgeWithdrawDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicBridgeWithdrawDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicBridgeWithdrawDelaySet represents a BridgeWithdrawDelaySet event raised by the BridgeLogic contract.
type BridgeLogicBridgeWithdrawDelaySet struct {
	WithdrawDelay *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeWithdrawDelaySet is a free log retrieval operation binding the contract event 0x1c7e8f73a01b8af4e18dd34455a42a45ad742bdb79cfda77bbdf50db2391fc88.
//
// Solidity: event Bridge_Withdraw_Delay_Set(uint256 withdraw_delay)
func (_BridgeLogic *BridgeLogicFilterer) FilterBridgeWithdrawDelaySet(opts *bind.FilterOpts) (*BridgeLogicBridgeWithdrawDelaySetIterator, error) {

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Bridge_Withdraw_Delay_Set")
	if err != nil {
		return nil, err
	}
	return &BridgeLogicBridgeWithdrawDelaySetIterator{contract: _BridgeLogic.contract, event: "Bridge_Withdraw_Delay_Set", logs: logs, sub: sub}, nil
}

// WatchBridgeWithdrawDelaySet is a free log subscription operation binding the contract event 0x1c7e8f73a01b8af4e18dd34455a42a45ad742bdb79cfda77bbdf50db2391fc88.
//
// Solidity: event Bridge_Withdraw_Delay_Set(uint256 withdraw_delay)
func (_BridgeLogic *BridgeLogicFilterer) WatchBridgeWithdrawDelaySet(opts *bind.WatchOpts, sink chan<- *BridgeLogicBridgeWithdrawDelaySet) (event.Subscription, error) {

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Bridge_Withdraw_Delay_Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicBridgeWithdrawDelaySet)
				if err := _BridgeLogic.contract.UnpackLog(event, "Bridge_Withdraw_Delay_Set", log); err != nil {
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

// ParseBridgeWithdrawDelaySet is a log parse operation binding the contract event 0x1c7e8f73a01b8af4e18dd34455a42a45ad742bdb79cfda77bbdf50db2391fc88.
//
// Solidity: event Bridge_Withdraw_Delay_Set(uint256 withdraw_delay)
func (_BridgeLogic *BridgeLogicFilterer) ParseBridgeWithdrawDelaySet(log types.Log) (*BridgeLogicBridgeWithdrawDelaySet, error) {
	event := new(BridgeLogicBridgeWithdrawDelaySet)
	if err := _BridgeLogic.contract.UnpackLog(event, "Bridge_Withdraw_Delay_Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogicDepositorExemptedIterator is returned from FilterDepositorExempted and is used to iterate over the raw logs and unpacked data for DepositorExempted events raised by the BridgeLogic contract.
type BridgeLogicDepositorExemptedIterator struct {
	Event *BridgeLogicDepositorExempted // Event containing the contract specifics and raw log

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
func (it *BridgeLogicDepositorExemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicDepositorExempted)
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
		it.Event = new(BridgeLogicDepositorExempted)
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
func (it *BridgeLogicDepositorExemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicDepositorExemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicDepositorExempted represents a DepositorExempted event raised by the BridgeLogic contract.
type BridgeLogicDepositorExempted struct {
	Depositor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositorExempted is a free log retrieval operation binding the contract event 0xf56e0868b913034a60dbca9c89ee79f8b0fa18dadbc5f6665f2f9a2cf3f51cdb.
//
// Solidity: event Depositor_Exempted(address indexed depositor)
func (_BridgeLogic *BridgeLogicFilterer) FilterDepositorExempted(opts *bind.FilterOpts, depositor []common.Address) (*BridgeLogicDepositorExemptedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Depositor_Exempted", depositorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicDepositorExemptedIterator{contract: _BridgeLogic.contract, event: "Depositor_Exempted", logs: logs, sub: sub}, nil
}

// WatchDepositorExempted is a free log subscription operation binding the contract event 0xf56e0868b913034a60dbca9c89ee79f8b0fa18dadbc5f6665f2f9a2cf3f51cdb.
//
// Solidity: event Depositor_Exempted(address indexed depositor)
func (_BridgeLogic *BridgeLogicFilterer) WatchDepositorExempted(opts *bind.WatchOpts, sink chan<- *BridgeLogicDepositorExempted, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Depositor_Exempted", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicDepositorExempted)
				if err := _BridgeLogic.contract.UnpackLog(event, "Depositor_Exempted", log); err != nil {
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

// ParseDepositorExempted is a log parse operation binding the contract event 0xf56e0868b913034a60dbca9c89ee79f8b0fa18dadbc5f6665f2f9a2cf3f51cdb.
//
// Solidity: event Depositor_Exempted(address indexed depositor)
func (_BridgeLogic *BridgeLogicFilterer) ParseDepositorExempted(log types.Log) (*BridgeLogicDepositorExempted, error) {
	event := new(BridgeLogicDepositorExempted)
	if err := _BridgeLogic.contract.UnpackLog(event, "Depositor_Exempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogicDepositorExemptionRevokedIterator is returned from FilterDepositorExemptionRevoked and is used to iterate over the raw logs and unpacked data for DepositorExemptionRevoked events raised by the BridgeLogic contract.
type BridgeLogicDepositorExemptionRevokedIterator struct {
	Event *BridgeLogicDepositorExemptionRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeLogicDepositorExemptionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogicDepositorExemptionRevoked)
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
		it.Event = new(BridgeLogicDepositorExemptionRevoked)
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
func (it *BridgeLogicDepositorExemptionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogicDepositorExemptionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogicDepositorExemptionRevoked represents a DepositorExemptionRevoked event raised by the BridgeLogic contract.
type BridgeLogicDepositorExemptionRevoked struct {
	Depositor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositorExemptionRevoked is a free log retrieval operation binding the contract event 0xe74b113dca87276d976f476a9b4b9da3c780a3262eaabad051ee4e98912936a4.
//
// Solidity: event Depositor_Exemption_Revoked(address indexed depositor)
func (_BridgeLogic *BridgeLogicFilterer) FilterDepositorExemptionRevoked(opts *bind.FilterOpts, depositor []common.Address) (*BridgeLogicDepositorExemptionRevokedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _BridgeLogic.contract.FilterLogs(opts, "Depositor_Exemption_Revoked", depositorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicDepositorExemptionRevokedIterator{contract: _BridgeLogic.contract, event: "Depositor_Exemption_Revoked", logs: logs, sub: sub}, nil
}

// WatchDepositorExemptionRevoked is a free log subscription operation binding the contract event 0xe74b113dca87276d976f476a9b4b9da3c780a3262eaabad051ee4e98912936a4.
//
// Solidity: event Depositor_Exemption_Revoked(address indexed depositor)
func (_BridgeLogic *BridgeLogicFilterer) WatchDepositorExemptionRevoked(opts *bind.WatchOpts, sink chan<- *BridgeLogicDepositorExemptionRevoked, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _BridgeLogic.contract.WatchLogs(opts, "Depositor_Exemption_Revoked", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogicDepositorExemptionRevoked)
				if err := _BridgeLogic.contract.UnpackLog(event, "Depositor_Exemption_Revoked", log); err != nil {
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

// ParseDepositorExemptionRevoked is a log parse operation binding the contract event 0xe74b113dca87276d976f476a9b4b9da3c780a3262eaabad051ee4e98912936a4.
//
// Solidity: event Depositor_Exemption_Revoked(address indexed depositor)
func (_BridgeLogic *BridgeLogicFilterer) ParseDepositorExemptionRevoked(log types.Log) (*BridgeLogicDepositorExemptionRevoked, error) {
	event := new(BridgeLogicDepositorExemptionRevoked)
	if err := _BridgeLogic.contract.UnpackLog(event, "Depositor_Exemption_Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
