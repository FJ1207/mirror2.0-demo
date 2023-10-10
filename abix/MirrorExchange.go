// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abix

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

// Fee is an auto generated low-level Go binding around an user-defined struct.
type Fee struct {
	Rate      uint16
	Recipient common.Address
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader         common.Address
	Side           uint8
	Collection     common.Address
	AssetType      uint8
	TokenId        *big.Int
	Amount         *big.Int
	PaymentToken   common.Address
	Price          *big.Int
	ListingTime    *big.Int
	ExpirationTime *big.Int
	Fees           []Fee
	Salt           *big.Int
}

// SignedOrder is an auto generated low-level Go binding around an user-defined struct.
type SignedOrder struct {
	Order Order
	V     uint8
	R     [32]byte
	S     [32]byte
}

// MirrorExchangeMetaData contains all meta data concerning the MirrorExchange contract.
var MirrorExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"name\":\"NewFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"NewFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"NewGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractITransferDelegate\",\"name\":\"transferDelegate\",\"type\":\"address\"}],\"name\":\"NewTransferDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Opened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"sell\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"buy\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"buyHash\",\"type\":\"bytes32\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INVERSE_BASIS_POINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MFIBO\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WFIBO\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelledOrFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedOrder\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedOrder\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransferDelegate\",\"name\":\"_transferDelegate\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransferDelegate\",\"name\":\"_transferDelegate\",\"type\":\"address\"}],\"name\":\"setTransferDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferDelegate\",\"outputs\":[{\"internalType\":\"contractITransferDelegate\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MirrorExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use MirrorExchangeMetaData.ABI instead.
var MirrorExchangeABI = MirrorExchangeMetaData.ABI

// MirrorExchange is an auto generated Go binding around an Ethereum contract.
type MirrorExchange struct {
	MirrorExchangeCaller     // Read-only binding to the contract
	MirrorExchangeTransactor // Write-only binding to the contract
	MirrorExchangeFilterer   // Log filterer for contract events
}

// MirrorExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MirrorExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MirrorExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MirrorExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MirrorExchangeSession struct {
	Contract     *MirrorExchange   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MirrorExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MirrorExchangeCallerSession struct {
	Contract *MirrorExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MirrorExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MirrorExchangeTransactorSession struct {
	Contract     *MirrorExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MirrorExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MirrorExchangeRaw struct {
	Contract *MirrorExchange // Generic contract binding to access the raw methods on
}

// MirrorExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MirrorExchangeCallerRaw struct {
	Contract *MirrorExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// MirrorExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MirrorExchangeTransactorRaw struct {
	Contract *MirrorExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMirrorExchange creates a new instance of MirrorExchange, bound to a specific deployed contract.
func NewMirrorExchange(address common.Address, backend bind.ContractBackend) (*MirrorExchange, error) {
	contract, err := bindMirrorExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MirrorExchange{MirrorExchangeCaller: MirrorExchangeCaller{contract: contract}, MirrorExchangeTransactor: MirrorExchangeTransactor{contract: contract}, MirrorExchangeFilterer: MirrorExchangeFilterer{contract: contract}}, nil
}

// NewMirrorExchangeCaller creates a new read-only instance of MirrorExchange, bound to a specific deployed contract.
func NewMirrorExchangeCaller(address common.Address, caller bind.ContractCaller) (*MirrorExchangeCaller, error) {
	contract, err := bindMirrorExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeCaller{contract: contract}, nil
}

// NewMirrorExchangeTransactor creates a new write-only instance of MirrorExchange, bound to a specific deployed contract.
func NewMirrorExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*MirrorExchangeTransactor, error) {
	contract, err := bindMirrorExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeTransactor{contract: contract}, nil
}

// NewMirrorExchangeFilterer creates a new log filterer instance of MirrorExchange, bound to a specific deployed contract.
func NewMirrorExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*MirrorExchangeFilterer, error) {
	contract, err := bindMirrorExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeFilterer{contract: contract}, nil
}

// bindMirrorExchange binds a generic wrapper to an already deployed contract.
func bindMirrorExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MirrorExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorExchange *MirrorExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MirrorExchange.Contract.MirrorExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorExchange *MirrorExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorExchange.Contract.MirrorExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorExchange *MirrorExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorExchange.Contract.MirrorExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorExchange *MirrorExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MirrorExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorExchange *MirrorExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorExchange *MirrorExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorExchange.Contract.contract.Transact(opts, method, params...)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_MirrorExchange *MirrorExchangeCaller) FEETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "FEE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_MirrorExchange *MirrorExchangeSession) FEETYPEHASH() ([32]byte, error) {
	return _MirrorExchange.Contract.FEETYPEHASH(&_MirrorExchange.CallOpts)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_MirrorExchange *MirrorExchangeCallerSession) FEETYPEHASH() ([32]byte, error) {
	return _MirrorExchange.Contract.FEETYPEHASH(&_MirrorExchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_MirrorExchange *MirrorExchangeCaller) INVERSEBASISPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "INVERSE_BASIS_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_MirrorExchange *MirrorExchangeSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _MirrorExchange.Contract.INVERSEBASISPOINT(&_MirrorExchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_MirrorExchange *MirrorExchangeCallerSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _MirrorExchange.Contract.INVERSEBASISPOINT(&_MirrorExchange.CallOpts)
}

// MFIBO is a free data retrieval call binding the contract method 0xe57da445.
//
// Solidity: function MFIBO() view returns(address)
func (_MirrorExchange *MirrorExchangeCaller) MFIBO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "MFIBO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MFIBO is a free data retrieval call binding the contract method 0xe57da445.
//
// Solidity: function MFIBO() view returns(address)
func (_MirrorExchange *MirrorExchangeSession) MFIBO() (common.Address, error) {
	return _MirrorExchange.Contract.MFIBO(&_MirrorExchange.CallOpts)
}

// MFIBO is a free data retrieval call binding the contract method 0xe57da445.
//
// Solidity: function MFIBO() view returns(address)
func (_MirrorExchange *MirrorExchangeCallerSession) MFIBO() (common.Address, error) {
	return _MirrorExchange.Contract.MFIBO(&_MirrorExchange.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_MirrorExchange *MirrorExchangeCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_MirrorExchange *MirrorExchangeSession) NAME() (string, error) {
	return _MirrorExchange.Contract.NAME(&_MirrorExchange.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_MirrorExchange *MirrorExchangeCallerSession) NAME() (string, error) {
	return _MirrorExchange.Contract.NAME(&_MirrorExchange.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_MirrorExchange *MirrorExchangeCaller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_MirrorExchange *MirrorExchangeSession) ORDERTYPEHASH() ([32]byte, error) {
	return _MirrorExchange.Contract.ORDERTYPEHASH(&_MirrorExchange.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_MirrorExchange *MirrorExchangeCallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _MirrorExchange.Contract.ORDERTYPEHASH(&_MirrorExchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_MirrorExchange *MirrorExchangeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_MirrorExchange *MirrorExchangeSession) VERSION() (string, error) {
	return _MirrorExchange.Contract.VERSION(&_MirrorExchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_MirrorExchange *MirrorExchangeCallerSession) VERSION() (string, error) {
	return _MirrorExchange.Contract.VERSION(&_MirrorExchange.CallOpts)
}

// WFIBO is a free data retrieval call binding the contract method 0x3ab6fe1f.
//
// Solidity: function WFIBO() view returns(address)
func (_MirrorExchange *MirrorExchangeCaller) WFIBO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "WFIBO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WFIBO is a free data retrieval call binding the contract method 0x3ab6fe1f.
//
// Solidity: function WFIBO() view returns(address)
func (_MirrorExchange *MirrorExchangeSession) WFIBO() (common.Address, error) {
	return _MirrorExchange.Contract.WFIBO(&_MirrorExchange.CallOpts)
}

// WFIBO is a free data retrieval call binding the contract method 0x3ab6fe1f.
//
// Solidity: function WFIBO() view returns(address)
func (_MirrorExchange *MirrorExchangeCallerSession) WFIBO() (common.Address, error) {
	return _MirrorExchange.Contract.WFIBO(&_MirrorExchange.CallOpts)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_MirrorExchange *MirrorExchangeCaller) CancelledOrFilled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "cancelledOrFilled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_MirrorExchange *MirrorExchangeSession) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _MirrorExchange.Contract.CancelledOrFilled(&_MirrorExchange.CallOpts, arg0)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_MirrorExchange *MirrorExchangeCallerSession) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _MirrorExchange.Contract.CancelledOrFilled(&_MirrorExchange.CallOpts, arg0)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_MirrorExchange *MirrorExchangeCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_MirrorExchange *MirrorExchangeSession) FeeRate() (*big.Int, error) {
	return _MirrorExchange.Contract.FeeRate(&_MirrorExchange.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_MirrorExchange *MirrorExchangeCallerSession) FeeRate() (*big.Int, error) {
	return _MirrorExchange.Contract.FeeRate(&_MirrorExchange.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MirrorExchange *MirrorExchangeCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MirrorExchange *MirrorExchangeSession) FeeRecipient() (common.Address, error) {
	return _MirrorExchange.Contract.FeeRecipient(&_MirrorExchange.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MirrorExchange *MirrorExchangeCallerSession) FeeRecipient() (common.Address, error) {
	return _MirrorExchange.Contract.FeeRecipient(&_MirrorExchange.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_MirrorExchange *MirrorExchangeCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_MirrorExchange *MirrorExchangeSession) Governor() (common.Address, error) {
	return _MirrorExchange.Contract.Governor(&_MirrorExchange.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_MirrorExchange *MirrorExchangeCallerSession) Governor() (common.Address, error) {
	return _MirrorExchange.Contract.Governor(&_MirrorExchange.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_MirrorExchange *MirrorExchangeCaller) IsOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_MirrorExchange *MirrorExchangeSession) IsOpen() (*big.Int, error) {
	return _MirrorExchange.Contract.IsOpen(&_MirrorExchange.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_MirrorExchange *MirrorExchangeCallerSession) IsOpen() (*big.Int, error) {
	return _MirrorExchange.Contract.IsOpen(&_MirrorExchange.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MirrorExchange *MirrorExchangeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MirrorExchange *MirrorExchangeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MirrorExchange.Contract.Nonces(&_MirrorExchange.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MirrorExchange *MirrorExchangeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MirrorExchange.Contract.Nonces(&_MirrorExchange.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MirrorExchange *MirrorExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MirrorExchange *MirrorExchangeSession) Owner() (common.Address, error) {
	return _MirrorExchange.Contract.Owner(&_MirrorExchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MirrorExchange *MirrorExchangeCallerSession) Owner() (common.Address, error) {
	return _MirrorExchange.Contract.Owner(&_MirrorExchange.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MirrorExchange *MirrorExchangeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MirrorExchange *MirrorExchangeSession) ProxiableUUID() ([32]byte, error) {
	return _MirrorExchange.Contract.ProxiableUUID(&_MirrorExchange.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MirrorExchange *MirrorExchangeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MirrorExchange.Contract.ProxiableUUID(&_MirrorExchange.CallOpts)
}

// TransferDelegate is a free data retrieval call binding the contract method 0xdd164a3b.
//
// Solidity: function transferDelegate() view returns(address)
func (_MirrorExchange *MirrorExchangeCaller) TransferDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MirrorExchange.contract.Call(opts, &out, "transferDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransferDelegate is a free data retrieval call binding the contract method 0xdd164a3b.
//
// Solidity: function transferDelegate() view returns(address)
func (_MirrorExchange *MirrorExchangeSession) TransferDelegate() (common.Address, error) {
	return _MirrorExchange.Contract.TransferDelegate(&_MirrorExchange.CallOpts)
}

// TransferDelegate is a free data retrieval call binding the contract method 0xdd164a3b.
//
// Solidity: function transferDelegate() view returns(address)
func (_MirrorExchange *MirrorExchangeCallerSession) TransferDelegate() (common.Address, error) {
	return _MirrorExchange.Contract.TransferDelegate(&_MirrorExchange.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x842afbd8.
//
// Solidity: function cancelOrder((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256) order) returns()
func (_MirrorExchange *MirrorExchangeTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x842afbd8.
//
// Solidity: function cancelOrder((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256) order) returns()
func (_MirrorExchange *MirrorExchangeSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _MirrorExchange.Contract.CancelOrder(&_MirrorExchange.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x842afbd8.
//
// Solidity: function cancelOrder((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256) order) returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _MirrorExchange.Contract.CancelOrder(&_MirrorExchange.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x98b7b7f5.
//
// Solidity: function cancelOrders((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256)[] orders) returns()
func (_MirrorExchange *MirrorExchangeTransactor) CancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x98b7b7f5.
//
// Solidity: function cancelOrders((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256)[] orders) returns()
func (_MirrorExchange *MirrorExchangeSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _MirrorExchange.Contract.CancelOrders(&_MirrorExchange.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x98b7b7f5.
//
// Solidity: function cancelOrders((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256)[] orders) returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _MirrorExchange.Contract.CancelOrders(&_MirrorExchange.TransactOpts, orders)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_MirrorExchange *MirrorExchangeTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_MirrorExchange *MirrorExchangeSession) Close() (*types.Transaction, error) {
	return _MirrorExchange.Contract.Close(&_MirrorExchange.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) Close() (*types.Transaction, error) {
	return _MirrorExchange.Contract.Close(&_MirrorExchange.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x142f4eb0.
//
// Solidity: function execute(((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256),uint8,bytes32,bytes32) sell, ((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256),uint8,bytes32,bytes32) buy) payable returns()
func (_MirrorExchange *MirrorExchangeTransactor) Execute(opts *bind.TransactOpts, sell SignedOrder, buy SignedOrder) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "execute", sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x142f4eb0.
//
// Solidity: function execute(((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256),uint8,bytes32,bytes32) sell, ((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256),uint8,bytes32,bytes32) buy) payable returns()
func (_MirrorExchange *MirrorExchangeSession) Execute(sell SignedOrder, buy SignedOrder) (*types.Transaction, error) {
	return _MirrorExchange.Contract.Execute(&_MirrorExchange.TransactOpts, sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x142f4eb0.
//
// Solidity: function execute(((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256),uint8,bytes32,bytes32) sell, ((address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256),uint8,bytes32,bytes32) buy) payable returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) Execute(sell SignedOrder, buy SignedOrder) (*types.Transaction, error) {
	return _MirrorExchange.Contract.Execute(&_MirrorExchange.TransactOpts, sell, buy)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_MirrorExchange *MirrorExchangeTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_MirrorExchange *MirrorExchangeSession) IncrementNonce() (*types.Transaction, error) {
	return _MirrorExchange.Contract.IncrementNonce(&_MirrorExchange.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _MirrorExchange.Contract.IncrementNonce(&_MirrorExchange.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _transferDelegate) returns()
func (_MirrorExchange *MirrorExchangeTransactor) Initialize(opts *bind.TransactOpts, _transferDelegate common.Address) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "initialize", _transferDelegate)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _transferDelegate) returns()
func (_MirrorExchange *MirrorExchangeSession) Initialize(_transferDelegate common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.Initialize(&_MirrorExchange.TransactOpts, _transferDelegate)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _transferDelegate) returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) Initialize(_transferDelegate common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.Initialize(&_MirrorExchange.TransactOpts, _transferDelegate)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_MirrorExchange *MirrorExchangeTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_MirrorExchange *MirrorExchangeSession) Open() (*types.Transaction, error) {
	return _MirrorExchange.Contract.Open(&_MirrorExchange.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) Open() (*types.Transaction, error) {
	return _MirrorExchange.Contract.Open(&_MirrorExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MirrorExchange *MirrorExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MirrorExchange *MirrorExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _MirrorExchange.Contract.RenounceOwnership(&_MirrorExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MirrorExchange.Contract.RenounceOwnership(&_MirrorExchange.TransactOpts)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_MirrorExchange *MirrorExchangeTransactor) SetFeeRate(opts *bind.TransactOpts, _feeRate *big.Int) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "setFeeRate", _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_MirrorExchange *MirrorExchangeSession) SetFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _MirrorExchange.Contract.SetFeeRate(&_MirrorExchange.TransactOpts, _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) SetFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _MirrorExchange.Contract.SetFeeRate(&_MirrorExchange.TransactOpts, _feeRate)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_MirrorExchange *MirrorExchangeTransactor) SetFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "setFeeRecipient", _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_MirrorExchange *MirrorExchangeSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.SetFeeRecipient(&_MirrorExchange.TransactOpts, _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.SetFeeRecipient(&_MirrorExchange.TransactOpts, _feeRecipient)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_MirrorExchange *MirrorExchangeTransactor) SetGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "setGovernor", _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_MirrorExchange *MirrorExchangeSession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.SetGovernor(&_MirrorExchange.TransactOpts, _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.SetGovernor(&_MirrorExchange.TransactOpts, _governor)
}

// SetTransferDelegate is a paid mutator transaction binding the contract method 0x1f2430e9.
//
// Solidity: function setTransferDelegate(address _transferDelegate) returns()
func (_MirrorExchange *MirrorExchangeTransactor) SetTransferDelegate(opts *bind.TransactOpts, _transferDelegate common.Address) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "setTransferDelegate", _transferDelegate)
}

// SetTransferDelegate is a paid mutator transaction binding the contract method 0x1f2430e9.
//
// Solidity: function setTransferDelegate(address _transferDelegate) returns()
func (_MirrorExchange *MirrorExchangeSession) SetTransferDelegate(_transferDelegate common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.SetTransferDelegate(&_MirrorExchange.TransactOpts, _transferDelegate)
}

// SetTransferDelegate is a paid mutator transaction binding the contract method 0x1f2430e9.
//
// Solidity: function setTransferDelegate(address _transferDelegate) returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) SetTransferDelegate(_transferDelegate common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.SetTransferDelegate(&_MirrorExchange.TransactOpts, _transferDelegate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MirrorExchange *MirrorExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MirrorExchange *MirrorExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.TransferOwnership(&_MirrorExchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.TransferOwnership(&_MirrorExchange.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MirrorExchange *MirrorExchangeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MirrorExchange *MirrorExchangeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.UpgradeTo(&_MirrorExchange.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MirrorExchange.Contract.UpgradeTo(&_MirrorExchange.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MirrorExchange *MirrorExchangeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MirrorExchange.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MirrorExchange *MirrorExchangeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MirrorExchange.Contract.UpgradeToAndCall(&_MirrorExchange.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MirrorExchange *MirrorExchangeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MirrorExchange.Contract.UpgradeToAndCall(&_MirrorExchange.TransactOpts, newImplementation, data)
}

// MirrorExchangeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the MirrorExchange contract.
type MirrorExchangeAdminChangedIterator struct {
	Event *MirrorExchangeAdminChanged // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeAdminChanged)
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
		it.Event = new(MirrorExchangeAdminChanged)
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
func (it *MirrorExchangeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeAdminChanged represents a AdminChanged event raised by the MirrorExchange contract.
type MirrorExchangeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MirrorExchange *MirrorExchangeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*MirrorExchangeAdminChangedIterator, error) {

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeAdminChangedIterator{contract: _MirrorExchange.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MirrorExchange *MirrorExchangeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *MirrorExchangeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeAdminChanged)
				if err := _MirrorExchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MirrorExchange *MirrorExchangeFilterer) ParseAdminChanged(log types.Log) (*MirrorExchangeAdminChanged, error) {
	event := new(MirrorExchangeAdminChanged)
	if err := _MirrorExchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the MirrorExchange contract.
type MirrorExchangeBeaconUpgradedIterator struct {
	Event *MirrorExchangeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeBeaconUpgraded)
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
		it.Event = new(MirrorExchangeBeaconUpgraded)
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
func (it *MirrorExchangeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeBeaconUpgraded represents a BeaconUpgraded event raised by the MirrorExchange contract.
type MirrorExchangeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MirrorExchange *MirrorExchangeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*MirrorExchangeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeBeaconUpgradedIterator{contract: _MirrorExchange.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MirrorExchange *MirrorExchangeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *MirrorExchangeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeBeaconUpgraded)
				if err := _MirrorExchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MirrorExchange *MirrorExchangeFilterer) ParseBeaconUpgraded(log types.Log) (*MirrorExchangeBeaconUpgraded, error) {
	event := new(MirrorExchangeBeaconUpgraded)
	if err := _MirrorExchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the MirrorExchange contract.
type MirrorExchangeClosedIterator struct {
	Event *MirrorExchangeClosed // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeClosed)
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
		it.Event = new(MirrorExchangeClosed)
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
func (it *MirrorExchangeClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeClosed represents a Closed event raised by the MirrorExchange contract.
type MirrorExchangeClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_MirrorExchange *MirrorExchangeFilterer) FilterClosed(opts *bind.FilterOpts) (*MirrorExchangeClosedIterator, error) {

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeClosedIterator{contract: _MirrorExchange.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_MirrorExchange *MirrorExchangeFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *MirrorExchangeClosed) (event.Subscription, error) {

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeClosed)
				if err := _MirrorExchange.contract.UnpackLog(event, "Closed", log); err != nil {
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

// ParseClosed is a log parse operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_MirrorExchange *MirrorExchangeFilterer) ParseClosed(log types.Log) (*MirrorExchangeClosed, error) {
	event := new(MirrorExchangeClosed)
	if err := _MirrorExchange.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MirrorExchange contract.
type MirrorExchangeInitializedIterator struct {
	Event *MirrorExchangeInitialized // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeInitialized)
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
		it.Event = new(MirrorExchangeInitialized)
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
func (it *MirrorExchangeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeInitialized represents a Initialized event raised by the MirrorExchange contract.
type MirrorExchangeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MirrorExchange *MirrorExchangeFilterer) FilterInitialized(opts *bind.FilterOpts) (*MirrorExchangeInitializedIterator, error) {

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeInitializedIterator{contract: _MirrorExchange.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MirrorExchange *MirrorExchangeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MirrorExchangeInitialized) (event.Subscription, error) {

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeInitialized)
				if err := _MirrorExchange.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MirrorExchange *MirrorExchangeFilterer) ParseInitialized(log types.Log) (*MirrorExchangeInitialized, error) {
	event := new(MirrorExchangeInitialized)
	if err := _MirrorExchange.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeNewFeeRateIterator is returned from FilterNewFeeRate and is used to iterate over the raw logs and unpacked data for NewFeeRate events raised by the MirrorExchange contract.
type MirrorExchangeNewFeeRateIterator struct {
	Event *MirrorExchangeNewFeeRate // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeNewFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeNewFeeRate)
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
		it.Event = new(MirrorExchangeNewFeeRate)
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
func (it *MirrorExchangeNewFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeNewFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeNewFeeRate represents a NewFeeRate event raised by the MirrorExchange contract.
type MirrorExchangeNewFeeRate struct {
	FeeRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFeeRate is a free log retrieval operation binding the contract event 0x788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae.
//
// Solidity: event NewFeeRate(uint256 feeRate)
func (_MirrorExchange *MirrorExchangeFilterer) FilterNewFeeRate(opts *bind.FilterOpts) (*MirrorExchangeNewFeeRateIterator, error) {

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "NewFeeRate")
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeNewFeeRateIterator{contract: _MirrorExchange.contract, event: "NewFeeRate", logs: logs, sub: sub}, nil
}

// WatchNewFeeRate is a free log subscription operation binding the contract event 0x788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae.
//
// Solidity: event NewFeeRate(uint256 feeRate)
func (_MirrorExchange *MirrorExchangeFilterer) WatchNewFeeRate(opts *bind.WatchOpts, sink chan<- *MirrorExchangeNewFeeRate) (event.Subscription, error) {

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "NewFeeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeNewFeeRate)
				if err := _MirrorExchange.contract.UnpackLog(event, "NewFeeRate", log); err != nil {
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

// ParseNewFeeRate is a log parse operation binding the contract event 0x788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae.
//
// Solidity: event NewFeeRate(uint256 feeRate)
func (_MirrorExchange *MirrorExchangeFilterer) ParseNewFeeRate(log types.Log) (*MirrorExchangeNewFeeRate, error) {
	event := new(MirrorExchangeNewFeeRate)
	if err := _MirrorExchange.contract.UnpackLog(event, "NewFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeNewFeeRecipientIterator is returned from FilterNewFeeRecipient and is used to iterate over the raw logs and unpacked data for NewFeeRecipient events raised by the MirrorExchange contract.
type MirrorExchangeNewFeeRecipientIterator struct {
	Event *MirrorExchangeNewFeeRecipient // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeNewFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeNewFeeRecipient)
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
		it.Event = new(MirrorExchangeNewFeeRecipient)
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
func (it *MirrorExchangeNewFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeNewFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeNewFeeRecipient represents a NewFeeRecipient event raised by the MirrorExchange contract.
type MirrorExchangeNewFeeRecipient struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewFeeRecipient is a free log retrieval operation binding the contract event 0x412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23.
//
// Solidity: event NewFeeRecipient(address feeRecipient)
func (_MirrorExchange *MirrorExchangeFilterer) FilterNewFeeRecipient(opts *bind.FilterOpts) (*MirrorExchangeNewFeeRecipientIterator, error) {

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "NewFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeNewFeeRecipientIterator{contract: _MirrorExchange.contract, event: "NewFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewFeeRecipient is a free log subscription operation binding the contract event 0x412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23.
//
// Solidity: event NewFeeRecipient(address feeRecipient)
func (_MirrorExchange *MirrorExchangeFilterer) WatchNewFeeRecipient(opts *bind.WatchOpts, sink chan<- *MirrorExchangeNewFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "NewFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeNewFeeRecipient)
				if err := _MirrorExchange.contract.UnpackLog(event, "NewFeeRecipient", log); err != nil {
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

// ParseNewFeeRecipient is a log parse operation binding the contract event 0x412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23.
//
// Solidity: event NewFeeRecipient(address feeRecipient)
func (_MirrorExchange *MirrorExchangeFilterer) ParseNewFeeRecipient(log types.Log) (*MirrorExchangeNewFeeRecipient, error) {
	event := new(MirrorExchangeNewFeeRecipient)
	if err := _MirrorExchange.contract.UnpackLog(event, "NewFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeNewGovernorIterator is returned from FilterNewGovernor and is used to iterate over the raw logs and unpacked data for NewGovernor events raised by the MirrorExchange contract.
type MirrorExchangeNewGovernorIterator struct {
	Event *MirrorExchangeNewGovernor // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeNewGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeNewGovernor)
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
		it.Event = new(MirrorExchangeNewGovernor)
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
func (it *MirrorExchangeNewGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeNewGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeNewGovernor represents a NewGovernor event raised by the MirrorExchange contract.
type MirrorExchangeNewGovernor struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewGovernor is a free log retrieval operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address governor)
func (_MirrorExchange *MirrorExchangeFilterer) FilterNewGovernor(opts *bind.FilterOpts) (*MirrorExchangeNewGovernorIterator, error) {

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeNewGovernorIterator{contract: _MirrorExchange.contract, event: "NewGovernor", logs: logs, sub: sub}, nil
}

// WatchNewGovernor is a free log subscription operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address governor)
func (_MirrorExchange *MirrorExchangeFilterer) WatchNewGovernor(opts *bind.WatchOpts, sink chan<- *MirrorExchangeNewGovernor) (event.Subscription, error) {

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeNewGovernor)
				if err := _MirrorExchange.contract.UnpackLog(event, "NewGovernor", log); err != nil {
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

// ParseNewGovernor is a log parse operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address governor)
func (_MirrorExchange *MirrorExchangeFilterer) ParseNewGovernor(log types.Log) (*MirrorExchangeNewGovernor, error) {
	event := new(MirrorExchangeNewGovernor)
	if err := _MirrorExchange.contract.UnpackLog(event, "NewGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeNewTransferDelegateIterator is returned from FilterNewTransferDelegate and is used to iterate over the raw logs and unpacked data for NewTransferDelegate events raised by the MirrorExchange contract.
type MirrorExchangeNewTransferDelegateIterator struct {
	Event *MirrorExchangeNewTransferDelegate // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeNewTransferDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeNewTransferDelegate)
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
		it.Event = new(MirrorExchangeNewTransferDelegate)
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
func (it *MirrorExchangeNewTransferDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeNewTransferDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeNewTransferDelegate represents a NewTransferDelegate event raised by the MirrorExchange contract.
type MirrorExchangeNewTransferDelegate struct {
	TransferDelegate common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewTransferDelegate is a free log retrieval operation binding the contract event 0x51cf954b2fbd80f88c0093810ea055bed0d40cecb34bbe8361aa130a290b309e.
//
// Solidity: event NewTransferDelegate(address indexed transferDelegate)
func (_MirrorExchange *MirrorExchangeFilterer) FilterNewTransferDelegate(opts *bind.FilterOpts, transferDelegate []common.Address) (*MirrorExchangeNewTransferDelegateIterator, error) {

	var transferDelegateRule []interface{}
	for _, transferDelegateItem := range transferDelegate {
		transferDelegateRule = append(transferDelegateRule, transferDelegateItem)
	}

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "NewTransferDelegate", transferDelegateRule)
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeNewTransferDelegateIterator{contract: _MirrorExchange.contract, event: "NewTransferDelegate", logs: logs, sub: sub}, nil
}

// WatchNewTransferDelegate is a free log subscription operation binding the contract event 0x51cf954b2fbd80f88c0093810ea055bed0d40cecb34bbe8361aa130a290b309e.
//
// Solidity: event NewTransferDelegate(address indexed transferDelegate)
func (_MirrorExchange *MirrorExchangeFilterer) WatchNewTransferDelegate(opts *bind.WatchOpts, sink chan<- *MirrorExchangeNewTransferDelegate, transferDelegate []common.Address) (event.Subscription, error) {

	var transferDelegateRule []interface{}
	for _, transferDelegateItem := range transferDelegate {
		transferDelegateRule = append(transferDelegateRule, transferDelegateItem)
	}

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "NewTransferDelegate", transferDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeNewTransferDelegate)
				if err := _MirrorExchange.contract.UnpackLog(event, "NewTransferDelegate", log); err != nil {
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

// ParseNewTransferDelegate is a log parse operation binding the contract event 0x51cf954b2fbd80f88c0093810ea055bed0d40cecb34bbe8361aa130a290b309e.
//
// Solidity: event NewTransferDelegate(address indexed transferDelegate)
func (_MirrorExchange *MirrorExchangeFilterer) ParseNewTransferDelegate(log types.Log) (*MirrorExchangeNewTransferDelegate, error) {
	event := new(MirrorExchangeNewTransferDelegate)
	if err := _MirrorExchange.contract.UnpackLog(event, "NewTransferDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeNonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the MirrorExchange contract.
type MirrorExchangeNonceIncrementedIterator struct {
	Event *MirrorExchangeNonceIncremented // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeNonceIncremented)
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
		it.Event = new(MirrorExchangeNonceIncremented)
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
func (it *MirrorExchangeNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeNonceIncremented represents a NonceIncremented event raised by the MirrorExchange contract.
type MirrorExchangeNonceIncremented struct {
	Trader   common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_MirrorExchange *MirrorExchangeFilterer) FilterNonceIncremented(opts *bind.FilterOpts, trader []common.Address) (*MirrorExchangeNonceIncrementedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeNonceIncrementedIterator{contract: _MirrorExchange.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_MirrorExchange *MirrorExchangeFilterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *MirrorExchangeNonceIncremented, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeNonceIncremented)
				if err := _MirrorExchange.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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

// ParseNonceIncremented is a log parse operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_MirrorExchange *MirrorExchangeFilterer) ParseNonceIncremented(log types.Log) (*MirrorExchangeNonceIncremented, error) {
	event := new(MirrorExchangeNonceIncremented)
	if err := _MirrorExchange.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeOpenedIterator is returned from FilterOpened and is used to iterate over the raw logs and unpacked data for Opened events raised by the MirrorExchange contract.
type MirrorExchangeOpenedIterator struct {
	Event *MirrorExchangeOpened // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeOpened)
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
		it.Event = new(MirrorExchangeOpened)
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
func (it *MirrorExchangeOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeOpened represents a Opened event raised by the MirrorExchange contract.
type MirrorExchangeOpened struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOpened is a free log retrieval operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_MirrorExchange *MirrorExchangeFilterer) FilterOpened(opts *bind.FilterOpts) (*MirrorExchangeOpenedIterator, error) {

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeOpenedIterator{contract: _MirrorExchange.contract, event: "Opened", logs: logs, sub: sub}, nil
}

// WatchOpened is a free log subscription operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_MirrorExchange *MirrorExchangeFilterer) WatchOpened(opts *bind.WatchOpts, sink chan<- *MirrorExchangeOpened) (event.Subscription, error) {

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeOpened)
				if err := _MirrorExchange.contract.UnpackLog(event, "Opened", log); err != nil {
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

// ParseOpened is a log parse operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_MirrorExchange *MirrorExchangeFilterer) ParseOpened(log types.Log) (*MirrorExchangeOpened, error) {
	event := new(MirrorExchangeOpened)
	if err := _MirrorExchange.contract.UnpackLog(event, "Opened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the MirrorExchange contract.
type MirrorExchangeOrderCancelledIterator struct {
	Event *MirrorExchangeOrderCancelled // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeOrderCancelled)
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
		it.Event = new(MirrorExchangeOrderCancelled)
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
func (it *MirrorExchangeOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeOrderCancelled represents a OrderCancelled event raised by the MirrorExchange contract.
type MirrorExchangeOrderCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_MirrorExchange *MirrorExchangeFilterer) FilterOrderCancelled(opts *bind.FilterOpts) (*MirrorExchangeOrderCancelledIterator, error) {

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeOrderCancelledIterator{contract: _MirrorExchange.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_MirrorExchange *MirrorExchangeFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *MirrorExchangeOrderCancelled) (event.Subscription, error) {

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeOrderCancelled)
				if err := _MirrorExchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_MirrorExchange *MirrorExchangeFilterer) ParseOrderCancelled(log types.Log) (*MirrorExchangeOrderCancelled, error) {
	event := new(MirrorExchangeOrderCancelled)
	if err := _MirrorExchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the MirrorExchange contract.
type MirrorExchangeOrdersMatchedIterator struct {
	Event *MirrorExchangeOrdersMatched // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeOrdersMatched)
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
		it.Event = new(MirrorExchangeOrdersMatched)
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
func (it *MirrorExchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeOrdersMatched represents a OrdersMatched event raised by the MirrorExchange contract.
type MirrorExchangeOrdersMatched struct {
	Maker    common.Address
	Taker    common.Address
	Sell     Order
	SellHash [32]byte
	Buy      Order
	BuyHash  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x83b9a672ba7c839b8b4284abb20790cdb697804487c36ac03bbc9e3367ef54db.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256) sell, bytes32 sellHash, (address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256) buy, bytes32 buyHash)
func (_MirrorExchange *MirrorExchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address) (*MirrorExchangeOrdersMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeOrdersMatchedIterator{contract: _MirrorExchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x83b9a672ba7c839b8b4284abb20790cdb697804487c36ac03bbc9e3367ef54db.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256) sell, bytes32 sellHash, (address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256) buy, bytes32 buyHash)
func (_MirrorExchange *MirrorExchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *MirrorExchangeOrdersMatched, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeOrdersMatched)
				if err := _MirrorExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x83b9a672ba7c839b8b4284abb20790cdb697804487c36ac03bbc9e3367ef54db.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256) sell, bytes32 sellHash, (address,uint8,address,uint8,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256) buy, bytes32 buyHash)
func (_MirrorExchange *MirrorExchangeFilterer) ParseOrdersMatched(log types.Log) (*MirrorExchangeOrdersMatched, error) {
	event := new(MirrorExchangeOrdersMatched)
	if err := _MirrorExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MirrorExchange contract.
type MirrorExchangeOwnershipTransferredIterator struct {
	Event *MirrorExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeOwnershipTransferred)
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
		it.Event = new(MirrorExchangeOwnershipTransferred)
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
func (it *MirrorExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the MirrorExchange contract.
type MirrorExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MirrorExchange *MirrorExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MirrorExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeOwnershipTransferredIterator{contract: _MirrorExchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MirrorExchange *MirrorExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MirrorExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeOwnershipTransferred)
				if err := _MirrorExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MirrorExchange *MirrorExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*MirrorExchangeOwnershipTransferred, error) {
	event := new(MirrorExchangeOwnershipTransferred)
	if err := _MirrorExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MirrorExchangeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MirrorExchange contract.
type MirrorExchangeUpgradedIterator struct {
	Event *MirrorExchangeUpgraded // Event containing the contract specifics and raw log

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
func (it *MirrorExchangeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MirrorExchangeUpgraded)
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
		it.Event = new(MirrorExchangeUpgraded)
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
func (it *MirrorExchangeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MirrorExchangeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MirrorExchangeUpgraded represents a Upgraded event raised by the MirrorExchange contract.
type MirrorExchangeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MirrorExchange *MirrorExchangeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MirrorExchangeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MirrorExchange.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MirrorExchangeUpgradedIterator{contract: _MirrorExchange.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MirrorExchange *MirrorExchangeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MirrorExchangeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MirrorExchange.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MirrorExchangeUpgraded)
				if err := _MirrorExchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MirrorExchange *MirrorExchangeFilterer) ParseUpgraded(log types.Log) (*MirrorExchangeUpgraded, error) {
	event := new(MirrorExchangeUpgraded)
	if err := _MirrorExchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
