// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opyn

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

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shortPowerPerp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wPowerPerp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quoteCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethQuoteCurrencyPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wPowerPerpPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniPositionManager\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_feeTier\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"BurnShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DepositUniPositionToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralPaid\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"MintShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldNormFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNormFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastModificationTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NormalizationFactorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"OpenVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pausesLeft\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payoutAmount\",\"type\":\"uint256\"}],\"name\":\"RedeemLong\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vauldId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"RedeemShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpExcess\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"}],\"name\":\"ReduceDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexForSettlement\",\"type\":\"uint256\"}],\"name\":\"Shutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"unpauser\",\"type\":\"address\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"UpdateOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawUniPositionToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FUNDING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TWAP_PERIOD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"applyFunding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"burnPowerPerpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wPowerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"burnWPowerPerpAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniTokenId\",\"type\":\"uint256\"}],\"name\":\"depositUniPositionToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethQuoteCurrencyPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTier\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getDenormalizedMark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getDenormalizedMarkForFunding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNormalizationFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getUnscaledIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexForSettlement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutDown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSystemPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"isVaultSafe\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFundingUpdateTimestamp\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPauseTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDebtAmount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniTokenId\",\"type\":\"uint256\"}],\"name\":\"mintPowerPerpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wPowerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniTokenId\",\"type\":\"uint256\"}],\"name\":\"mintWPowerPerpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"normalizationFactor\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausesLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wPerpAmount\",\"type\":\"uint256\"}],\"name\":\"redeemLong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"redeemShort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"reduceDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"reduceDebtShutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shortPowerPerp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutDown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPauseAnyone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPauseOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"NftCollateralId\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"collateralAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint128\",\"name\":\"shortAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wPowerPerp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wPowerPerpPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"withdrawUniPositionToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]any, method string, params ...any) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...any) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]any, method string, params ...any) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...any) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_Controller *ControllerCaller) FUNDINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "FUNDING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_Controller *ControllerSession) FUNDINGPERIOD() (*big.Int, error) {
	return _Controller.Contract.FUNDINGPERIOD(&_Controller.CallOpts)
}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_Controller *ControllerCallerSession) FUNDINGPERIOD() (*big.Int, error) {
	return _Controller.Contract.FUNDINGPERIOD(&_Controller.CallOpts)
}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_Controller *ControllerCaller) TWAPPERIOD(opts *bind.CallOpts) (uint32, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "TWAP_PERIOD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_Controller *ControllerSession) TWAPPERIOD() (uint32, error) {
	return _Controller.Contract.TWAPPERIOD(&_Controller.CallOpts)
}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_Controller *ControllerCallerSession) TWAPPERIOD() (uint32, error) {
	return _Controller.Contract.TWAPPERIOD(&_Controller.CallOpts)
}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Controller *ControllerCaller) EthQuoteCurrencyPool(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "ethQuoteCurrencyPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Controller *ControllerSession) EthQuoteCurrencyPool() (common.Address, error) {
	return _Controller.Contract.EthQuoteCurrencyPool(&_Controller.CallOpts)
}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Controller *ControllerCallerSession) EthQuoteCurrencyPool() (common.Address, error) {
	return _Controller.Contract.EthQuoteCurrencyPool(&_Controller.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Controller *ControllerCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Controller *ControllerSession) FeeRate() (*big.Int, error) {
	return _Controller.Contract.FeeRate(&_Controller.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Controller *ControllerCallerSession) FeeRate() (*big.Int, error) {
	return _Controller.Contract.FeeRate(&_Controller.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Controller *ControllerCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Controller *ControllerSession) FeeRecipient() (common.Address, error) {
	return _Controller.Contract.FeeRecipient(&_Controller.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Controller *ControllerCallerSession) FeeRecipient() (common.Address, error) {
	return _Controller.Contract.FeeRecipient(&_Controller.CallOpts)
}

// FeeTier is a free data retrieval call binding the contract method 0x72f5d98a.
//
// Solidity: function feeTier() view returns(uint24)
func (_Controller *ControllerCaller) FeeTier(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "feeTier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeTier is a free data retrieval call binding the contract method 0x72f5d98a.
//
// Solidity: function feeTier() view returns(uint24)
func (_Controller *ControllerSession) FeeTier() (*big.Int, error) {
	return _Controller.Contract.FeeTier(&_Controller.CallOpts)
}

// FeeTier is a free data retrieval call binding the contract method 0x72f5d98a.
//
// Solidity: function feeTier() view returns(uint24)
func (_Controller *ControllerCallerSession) FeeTier() (*big.Int, error) {
	return _Controller.Contract.FeeTier(&_Controller.CallOpts)
}

// GetDenormalizedMark is a free data retrieval call binding the contract method 0xee3189ff.
//
// Solidity: function getDenormalizedMark(uint32 _period) view returns(uint256)
func (_Controller *ControllerCaller) GetDenormalizedMark(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "getDenormalizedMark", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedMark is a free data retrieval call binding the contract method 0xee3189ff.
//
// Solidity: function getDenormalizedMark(uint32 _period) view returns(uint256)
func (_Controller *ControllerSession) GetDenormalizedMark(_period uint32) (*big.Int, error) {
	return _Controller.Contract.GetDenormalizedMark(&_Controller.CallOpts, _period)
}

// GetDenormalizedMark is a free data retrieval call binding the contract method 0xee3189ff.
//
// Solidity: function getDenormalizedMark(uint32 _period) view returns(uint256)
func (_Controller *ControllerCallerSession) GetDenormalizedMark(_period uint32) (*big.Int, error) {
	return _Controller.Contract.GetDenormalizedMark(&_Controller.CallOpts, _period)
}

// GetDenormalizedMarkForFunding is a free data retrieval call binding the contract method 0x377a1936.
//
// Solidity: function getDenormalizedMarkForFunding(uint32 _period) view returns(uint256)
func (_Controller *ControllerCaller) GetDenormalizedMarkForFunding(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "getDenormalizedMarkForFunding", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedMarkForFunding is a free data retrieval call binding the contract method 0x377a1936.
//
// Solidity: function getDenormalizedMarkForFunding(uint32 _period) view returns(uint256)
func (_Controller *ControllerSession) GetDenormalizedMarkForFunding(_period uint32) (*big.Int, error) {
	return _Controller.Contract.GetDenormalizedMarkForFunding(&_Controller.CallOpts, _period)
}

// GetDenormalizedMarkForFunding is a free data retrieval call binding the contract method 0x377a1936.
//
// Solidity: function getDenormalizedMarkForFunding(uint32 _period) view returns(uint256)
func (_Controller *ControllerCallerSession) GetDenormalizedMarkForFunding(_period uint32) (*big.Int, error) {
	return _Controller.Contract.GetDenormalizedMarkForFunding(&_Controller.CallOpts, _period)
}

// GetExpectedNormalizationFactor is a free data retrieval call binding the contract method 0x24f5f531.
//
// Solidity: function getExpectedNormalizationFactor() view returns(uint256)
func (_Controller *ControllerCaller) GetExpectedNormalizationFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "getExpectedNormalizationFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedNormalizationFactor is a free data retrieval call binding the contract method 0x24f5f531.
//
// Solidity: function getExpectedNormalizationFactor() view returns(uint256)
func (_Controller *ControllerSession) GetExpectedNormalizationFactor() (*big.Int, error) {
	return _Controller.Contract.GetExpectedNormalizationFactor(&_Controller.CallOpts)
}

// GetExpectedNormalizationFactor is a free data retrieval call binding the contract method 0x24f5f531.
//
// Solidity: function getExpectedNormalizationFactor() view returns(uint256)
func (_Controller *ControllerCallerSession) GetExpectedNormalizationFactor() (*big.Int, error) {
	return _Controller.Contract.GetExpectedNormalizationFactor(&_Controller.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0x8cd21d7c.
//
// Solidity: function getIndex(uint32 _period) view returns(uint256)
func (_Controller *ControllerCaller) GetIndex(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "getIndex", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndex is a free data retrieval call binding the contract method 0x8cd21d7c.
//
// Solidity: function getIndex(uint32 _period) view returns(uint256)
func (_Controller *ControllerSession) GetIndex(_period uint32) (*big.Int, error) {
	return _Controller.Contract.GetIndex(&_Controller.CallOpts, _period)
}

// GetIndex is a free data retrieval call binding the contract method 0x8cd21d7c.
//
// Solidity: function getIndex(uint32 _period) view returns(uint256)
func (_Controller *ControllerCallerSession) GetIndex(_period uint32) (*big.Int, error) {
	return _Controller.Contract.GetIndex(&_Controller.CallOpts, _period)
}

// GetUnscaledIndex is a free data retrieval call binding the contract method 0x15aded83.
//
// Solidity: function getUnscaledIndex(uint32 _period) view returns(uint256)
func (_Controller *ControllerCaller) GetUnscaledIndex(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "getUnscaledIndex", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnscaledIndex is a free data retrieval call binding the contract method 0x15aded83.
//
// Solidity: function getUnscaledIndex(uint32 _period) view returns(uint256)
func (_Controller *ControllerSession) GetUnscaledIndex(_period uint32) (*big.Int, error) {
	return _Controller.Contract.GetUnscaledIndex(&_Controller.CallOpts, _period)
}

// GetUnscaledIndex is a free data retrieval call binding the contract method 0x15aded83.
//
// Solidity: function getUnscaledIndex(uint32 _period) view returns(uint256)
func (_Controller *ControllerCallerSession) GetUnscaledIndex(_period uint32) (*big.Int, error) {
	return _Controller.Contract.GetUnscaledIndex(&_Controller.CallOpts, _period)
}

// IndexForSettlement is a free data retrieval call binding the contract method 0xde4a427a.
//
// Solidity: function indexForSettlement() view returns(uint256)
func (_Controller *ControllerCaller) IndexForSettlement(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "indexForSettlement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexForSettlement is a free data retrieval call binding the contract method 0xde4a427a.
//
// Solidity: function indexForSettlement() view returns(uint256)
func (_Controller *ControllerSession) IndexForSettlement() (*big.Int, error) {
	return _Controller.Contract.IndexForSettlement(&_Controller.CallOpts)
}

// IndexForSettlement is a free data retrieval call binding the contract method 0xde4a427a.
//
// Solidity: function indexForSettlement() view returns(uint256)
func (_Controller *ControllerCallerSession) IndexForSettlement() (*big.Int, error) {
	return _Controller.Contract.IndexForSettlement(&_Controller.CallOpts)
}

// IsShutDown is a free data retrieval call binding the contract method 0xff947525.
//
// Solidity: function isShutDown() view returns(bool)
func (_Controller *ControllerCaller) IsShutDown(opts *bind.CallOpts) (bool, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "isShutDown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutDown is a free data retrieval call binding the contract method 0xff947525.
//
// Solidity: function isShutDown() view returns(bool)
func (_Controller *ControllerSession) IsShutDown() (bool, error) {
	return _Controller.Contract.IsShutDown(&_Controller.CallOpts)
}

// IsShutDown is a free data retrieval call binding the contract method 0xff947525.
//
// Solidity: function isShutDown() view returns(bool)
func (_Controller *ControllerCallerSession) IsShutDown() (bool, error) {
	return _Controller.Contract.IsShutDown(&_Controller.CallOpts)
}

// IsSystemPaused is a free data retrieval call binding the contract method 0x7691c4ac.
//
// Solidity: function isSystemPaused() view returns(bool)
func (_Controller *ControllerCaller) IsSystemPaused(opts *bind.CallOpts) (bool, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "isSystemPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSystemPaused is a free data retrieval call binding the contract method 0x7691c4ac.
//
// Solidity: function isSystemPaused() view returns(bool)
func (_Controller *ControllerSession) IsSystemPaused() (bool, error) {
	return _Controller.Contract.IsSystemPaused(&_Controller.CallOpts)
}

// IsSystemPaused is a free data retrieval call binding the contract method 0x7691c4ac.
//
// Solidity: function isSystemPaused() view returns(bool)
func (_Controller *ControllerCallerSession) IsSystemPaused() (bool, error) {
	return _Controller.Contract.IsSystemPaused(&_Controller.CallOpts)
}

// IsVaultSafe is a free data retrieval call binding the contract method 0xa847e674.
//
// Solidity: function isVaultSafe(uint256 _vaultId) view returns(bool)
func (_Controller *ControllerCaller) IsVaultSafe(opts *bind.CallOpts, _vaultId *big.Int) (bool, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "isVaultSafe", _vaultId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVaultSafe is a free data retrieval call binding the contract method 0xa847e674.
//
// Solidity: function isVaultSafe(uint256 _vaultId) view returns(bool)
func (_Controller *ControllerSession) IsVaultSafe(_vaultId *big.Int) (bool, error) {
	return _Controller.Contract.IsVaultSafe(&_Controller.CallOpts, _vaultId)
}

// IsVaultSafe is a free data retrieval call binding the contract method 0xa847e674.
//
// Solidity: function isVaultSafe(uint256 _vaultId) view returns(bool)
func (_Controller *ControllerCallerSession) IsVaultSafe(_vaultId *big.Int) (bool, error) {
	return _Controller.Contract.IsVaultSafe(&_Controller.CallOpts, _vaultId)
}

// LastFundingUpdateTimestamp is a free data retrieval call binding the contract method 0x4be2822c.
//
// Solidity: function lastFundingUpdateTimestamp() view returns(uint128)
func (_Controller *ControllerCaller) LastFundingUpdateTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "lastFundingUpdateTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingUpdateTimestamp is a free data retrieval call binding the contract method 0x4be2822c.
//
// Solidity: function lastFundingUpdateTimestamp() view returns(uint128)
func (_Controller *ControllerSession) LastFundingUpdateTimestamp() (*big.Int, error) {
	return _Controller.Contract.LastFundingUpdateTimestamp(&_Controller.CallOpts)
}

// LastFundingUpdateTimestamp is a free data retrieval call binding the contract method 0x4be2822c.
//
// Solidity: function lastFundingUpdateTimestamp() view returns(uint128)
func (_Controller *ControllerCallerSession) LastFundingUpdateTimestamp() (*big.Int, error) {
	return _Controller.Contract.LastFundingUpdateTimestamp(&_Controller.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Controller *ControllerCaller) LastPauseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "lastPauseTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Controller *ControllerSession) LastPauseTime() (*big.Int, error) {
	return _Controller.Contract.LastPauseTime(&_Controller.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Controller *ControllerCallerSession) LastPauseTime() (*big.Int, error) {
	return _Controller.Contract.LastPauseTime(&_Controller.CallOpts)
}

// NormalizationFactor is a free data retrieval call binding the contract method 0xd5272584.
//
// Solidity: function normalizationFactor() view returns(uint128)
func (_Controller *ControllerCaller) NormalizationFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "normalizationFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NormalizationFactor is a free data retrieval call binding the contract method 0xd5272584.
//
// Solidity: function normalizationFactor() view returns(uint128)
func (_Controller *ControllerSession) NormalizationFactor() (*big.Int, error) {
	return _Controller.Contract.NormalizationFactor(&_Controller.CallOpts)
}

// NormalizationFactor is a free data retrieval call binding the contract method 0xd5272584.
//
// Solidity: function normalizationFactor() view returns(uint128)
func (_Controller *ControllerCallerSession) NormalizationFactor() (*big.Int, error) {
	return _Controller.Contract.NormalizationFactor(&_Controller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Controller *ControllerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Controller *ControllerSession) Oracle() (common.Address, error) {
	return _Controller.Contract.Oracle(&_Controller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Controller *ControllerCallerSession) Oracle() (common.Address, error) {
	return _Controller.Contract.Oracle(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// PausesLeft is a free data retrieval call binding the contract method 0x63b38ae4.
//
// Solidity: function pausesLeft() view returns(uint256)
func (_Controller *ControllerCaller) PausesLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "pausesLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PausesLeft is a free data retrieval call binding the contract method 0x63b38ae4.
//
// Solidity: function pausesLeft() view returns(uint256)
func (_Controller *ControllerSession) PausesLeft() (*big.Int, error) {
	return _Controller.Contract.PausesLeft(&_Controller.CallOpts)
}

// PausesLeft is a free data retrieval call binding the contract method 0x63b38ae4.
//
// Solidity: function pausesLeft() view returns(uint256)
func (_Controller *ControllerCallerSession) PausesLeft() (*big.Int, error) {
	return _Controller.Contract.PausesLeft(&_Controller.CallOpts)
}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Controller *ControllerCaller) QuoteCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "quoteCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Controller *ControllerSession) QuoteCurrency() (common.Address, error) {
	return _Controller.Contract.QuoteCurrency(&_Controller.CallOpts)
}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Controller *ControllerCallerSession) QuoteCurrency() (common.Address, error) {
	return _Controller.Contract.QuoteCurrency(&_Controller.CallOpts)
}

// ShortPowerPerp is a free data retrieval call binding the contract method 0x9d4c9442.
//
// Solidity: function shortPowerPerp() view returns(address)
func (_Controller *ControllerCaller) ShortPowerPerp(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "shortPowerPerp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShortPowerPerp is a free data retrieval call binding the contract method 0x9d4c9442.
//
// Solidity: function shortPowerPerp() view returns(address)
func (_Controller *ControllerSession) ShortPowerPerp() (common.Address, error) {
	return _Controller.Contract.ShortPowerPerp(&_Controller.CallOpts)
}

// ShortPowerPerp is a free data retrieval call binding the contract method 0x9d4c9442.
//
// Solidity: function shortPowerPerp() view returns(address)
func (_Controller *ControllerCallerSession) ShortPowerPerp() (common.Address, error) {
	return _Controller.Contract.ShortPowerPerp(&_Controller.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address operator, uint32 NftCollateralId, uint96 collateralAmount, uint128 shortAmount)
func (_Controller *ControllerCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "vaults", arg0)

	outstruct := new(struct {
		Operator         common.Address
		NftCollateralId  uint32
		CollateralAmount *big.Int
		ShortAmount      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftCollateralId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.CollateralAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ShortAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address operator, uint32 NftCollateralId, uint96 collateralAmount, uint128 shortAmount)
func (_Controller *ControllerSession) Vaults(arg0 *big.Int) (struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}, error) {
	return _Controller.Contract.Vaults(&_Controller.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address operator, uint32 NftCollateralId, uint96 collateralAmount, uint128 shortAmount)
func (_Controller *ControllerCallerSession) Vaults(arg0 *big.Int) (struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}, error) {
	return _Controller.Contract.Vaults(&_Controller.CallOpts, arg0)
}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Controller *ControllerCaller) WPowerPerp(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "wPowerPerp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Controller *ControllerSession) WPowerPerp() (common.Address, error) {
	return _Controller.Contract.WPowerPerp(&_Controller.CallOpts)
}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Controller *ControllerCallerSession) WPowerPerp() (common.Address, error) {
	return _Controller.Contract.WPowerPerp(&_Controller.CallOpts)
}

// WPowerPerpPool is a free data retrieval call binding the contract method 0xb707ab99.
//
// Solidity: function wPowerPerpPool() view returns(address)
func (_Controller *ControllerCaller) WPowerPerpPool(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "wPowerPerpPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPowerPerpPool is a free data retrieval call binding the contract method 0xb707ab99.
//
// Solidity: function wPowerPerpPool() view returns(address)
func (_Controller *ControllerSession) WPowerPerpPool() (common.Address, error) {
	return _Controller.Contract.WPowerPerpPool(&_Controller.CallOpts)
}

// WPowerPerpPool is a free data retrieval call binding the contract method 0xb707ab99.
//
// Solidity: function wPowerPerpPool() view returns(address)
func (_Controller *ControllerCallerSession) WPowerPerpPool() (common.Address, error) {
	return _Controller.Contract.WPowerPerpPool(&_Controller.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Controller *ControllerCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Controller.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Controller *ControllerSession) Weth() (common.Address, error) {
	return _Controller.Contract.Weth(&_Controller.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Controller *ControllerCallerSession) Weth() (common.Address, error) {
	return _Controller.Contract.Weth(&_Controller.CallOpts)
}

// ApplyFunding is a paid mutator transaction binding the contract method 0x200f4b8d.
//
// Solidity: function applyFunding() returns()
func (_Controller *ControllerTransactor) ApplyFunding(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "applyFunding")
}

// ApplyFunding is a paid mutator transaction binding the contract method 0x200f4b8d.
//
// Solidity: function applyFunding() returns()
func (_Controller *ControllerSession) ApplyFunding() (*types.Transaction, error) {
	return _Controller.Contract.ApplyFunding(&_Controller.TransactOpts)
}

// ApplyFunding is a paid mutator transaction binding the contract method 0x200f4b8d.
//
// Solidity: function applyFunding() returns()
func (_Controller *ControllerTransactorSession) ApplyFunding() (*types.Transaction, error) {
	return _Controller.Contract.ApplyFunding(&_Controller.TransactOpts)
}

// BurnPowerPerpAmount is a paid mutator transaction binding the contract method 0x4394318d.
//
// Solidity: function burnPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _withdrawAmount) returns(uint256)
func (_Controller *ControllerTransactor) BurnPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _powerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "burnPowerPerpAmount", _vaultId, _powerPerpAmount, _withdrawAmount)
}

// BurnPowerPerpAmount is a paid mutator transaction binding the contract method 0x4394318d.
//
// Solidity: function burnPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _withdrawAmount) returns(uint256)
func (_Controller *ControllerSession) BurnPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BurnPowerPerpAmount(&_Controller.TransactOpts, _vaultId, _powerPerpAmount, _withdrawAmount)
}

// BurnPowerPerpAmount is a paid mutator transaction binding the contract method 0x4394318d.
//
// Solidity: function burnPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _withdrawAmount) returns(uint256)
func (_Controller *ControllerTransactorSession) BurnPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BurnPowerPerpAmount(&_Controller.TransactOpts, _vaultId, _powerPerpAmount, _withdrawAmount)
}

// BurnWPowerPerpAmount is a paid mutator transaction binding the contract method 0x8632cb03.
//
// Solidity: function burnWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _withdrawAmount) returns()
func (_Controller *ControllerTransactor) BurnWPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _wPowerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "burnWPowerPerpAmount", _vaultId, _wPowerPerpAmount, _withdrawAmount)
}

// BurnWPowerPerpAmount is a paid mutator transaction binding the contract method 0x8632cb03.
//
// Solidity: function burnWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _withdrawAmount) returns()
func (_Controller *ControllerSession) BurnWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BurnWPowerPerpAmount(&_Controller.TransactOpts, _vaultId, _wPowerPerpAmount, _withdrawAmount)
}

// BurnWPowerPerpAmount is a paid mutator transaction binding the contract method 0x8632cb03.
//
// Solidity: function burnWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _withdrawAmount) returns()
func (_Controller *ControllerTransactorSession) BurnWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BurnWPowerPerpAmount(&_Controller.TransactOpts, _vaultId, _wPowerPerpAmount, _withdrawAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _vaultId) payable returns()
func (_Controller *ControllerTransactor) Deposit(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "deposit", _vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _vaultId) payable returns()
func (_Controller *ControllerSession) Deposit(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Deposit(&_Controller.TransactOpts, _vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _vaultId) payable returns()
func (_Controller *ControllerTransactorSession) Deposit(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Deposit(&_Controller.TransactOpts, _vaultId)
}

// DepositUniPositionToken is a paid mutator transaction binding the contract method 0x91b8d34a.
//
// Solidity: function depositUniPositionToken(uint256 _vaultId, uint256 _uniTokenId) returns()
func (_Controller *ControllerTransactor) DepositUniPositionToken(opts *bind.TransactOpts, _vaultId *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "depositUniPositionToken", _vaultId, _uniTokenId)
}

// DepositUniPositionToken is a paid mutator transaction binding the contract method 0x91b8d34a.
//
// Solidity: function depositUniPositionToken(uint256 _vaultId, uint256 _uniTokenId) returns()
func (_Controller *ControllerSession) DepositUniPositionToken(_vaultId *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.DepositUniPositionToken(&_Controller.TransactOpts, _vaultId, _uniTokenId)
}

// DepositUniPositionToken is a paid mutator transaction binding the contract method 0x91b8d34a.
//
// Solidity: function depositUniPositionToken(uint256 _vaultId, uint256 _uniTokenId) returns()
func (_Controller *ControllerTransactorSession) DepositUniPositionToken(_vaultId *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.DepositUniPositionToken(&_Controller.TransactOpts, _vaultId, _uniTokenId)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Controller *ControllerTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Controller *ControllerSession) Donate() (*types.Transaction, error) {
	return _Controller.Contract.Donate(&_Controller.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Controller *ControllerTransactorSession) Donate() (*types.Transaction, error) {
	return _Controller.Contract.Donate(&_Controller.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd296d1f1.
//
// Solidity: function liquidate(uint256 _vaultId, uint256 _maxDebtAmount) returns(uint256)
func (_Controller *ControllerTransactor) Liquidate(opts *bind.TransactOpts, _vaultId *big.Int, _maxDebtAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "liquidate", _vaultId, _maxDebtAmount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd296d1f1.
//
// Solidity: function liquidate(uint256 _vaultId, uint256 _maxDebtAmount) returns(uint256)
func (_Controller *ControllerSession) Liquidate(_vaultId *big.Int, _maxDebtAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Liquidate(&_Controller.TransactOpts, _vaultId, _maxDebtAmount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd296d1f1.
//
// Solidity: function liquidate(uint256 _vaultId, uint256 _maxDebtAmount) returns(uint256)
func (_Controller *ControllerTransactorSession) Liquidate(_vaultId *big.Int, _maxDebtAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Liquidate(&_Controller.TransactOpts, _vaultId, _maxDebtAmount)
}

// MintPowerPerpAmount is a paid mutator transaction binding the contract method 0x1bf7bf6c.
//
// Solidity: function mintPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _uniTokenId) payable returns(uint256, uint256)
func (_Controller *ControllerTransactor) MintPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _powerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "mintPowerPerpAmount", _vaultId, _powerPerpAmount, _uniTokenId)
}

// MintPowerPerpAmount is a paid mutator transaction binding the contract method 0x1bf7bf6c.
//
// Solidity: function mintPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _uniTokenId) payable returns(uint256, uint256)
func (_Controller *ControllerSession) MintPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.MintPowerPerpAmount(&_Controller.TransactOpts, _vaultId, _powerPerpAmount, _uniTokenId)
}

// MintPowerPerpAmount is a paid mutator transaction binding the contract method 0x1bf7bf6c.
//
// Solidity: function mintPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _uniTokenId) payable returns(uint256, uint256)
func (_Controller *ControllerTransactorSession) MintPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.MintPowerPerpAmount(&_Controller.TransactOpts, _vaultId, _powerPerpAmount, _uniTokenId)
}

// MintWPowerPerpAmount is a paid mutator transaction binding the contract method 0x39467918.
//
// Solidity: function mintWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _uniTokenId) payable returns(uint256)
func (_Controller *ControllerTransactor) MintWPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _wPowerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "mintWPowerPerpAmount", _vaultId, _wPowerPerpAmount, _uniTokenId)
}

// MintWPowerPerpAmount is a paid mutator transaction binding the contract method 0x39467918.
//
// Solidity: function mintWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _uniTokenId) payable returns(uint256)
func (_Controller *ControllerSession) MintWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.MintWPowerPerpAmount(&_Controller.TransactOpts, _vaultId, _wPowerPerpAmount, _uniTokenId)
}

// MintWPowerPerpAmount is a paid mutator transaction binding the contract method 0x39467918.
//
// Solidity: function mintWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _uniTokenId) payable returns(uint256)
func (_Controller *ControllerTransactorSession) MintWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.MintWPowerPerpAmount(&_Controller.TransactOpts, _vaultId, _wPowerPerpAmount, _uniTokenId)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Controller *ControllerTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Controller *ControllerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Controller.Contract.OnERC721Received(&_Controller.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Controller *ControllerTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Controller.Contract.OnERC721Received(&_Controller.TransactOpts, arg0, arg1, arg2, arg3)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Controller *ControllerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Controller *ControllerSession) Pause() (*types.Transaction, error) {
	return _Controller.Contract.Pause(&_Controller.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Controller *ControllerTransactorSession) Pause() (*types.Transaction, error) {
	return _Controller.Contract.Pause(&_Controller.TransactOpts)
}

// RedeemLong is a paid mutator transaction binding the contract method 0xac6cd5ef.
//
// Solidity: function redeemLong(uint256 _wPerpAmount) returns()
func (_Controller *ControllerTransactor) RedeemLong(opts *bind.TransactOpts, _wPerpAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "redeemLong", _wPerpAmount)
}

// RedeemLong is a paid mutator transaction binding the contract method 0xac6cd5ef.
//
// Solidity: function redeemLong(uint256 _wPerpAmount) returns()
func (_Controller *ControllerSession) RedeemLong(_wPerpAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.RedeemLong(&_Controller.TransactOpts, _wPerpAmount)
}

// RedeemLong is a paid mutator transaction binding the contract method 0xac6cd5ef.
//
// Solidity: function redeemLong(uint256 _wPerpAmount) returns()
func (_Controller *ControllerTransactorSession) RedeemLong(_wPerpAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.RedeemLong(&_Controller.TransactOpts, _wPerpAmount)
}

// RedeemShort is a paid mutator transaction binding the contract method 0x97efa942.
//
// Solidity: function redeemShort(uint256 _vaultId) returns()
func (_Controller *ControllerTransactor) RedeemShort(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "redeemShort", _vaultId)
}

// RedeemShort is a paid mutator transaction binding the contract method 0x97efa942.
//
// Solidity: function redeemShort(uint256 _vaultId) returns()
func (_Controller *ControllerSession) RedeemShort(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.RedeemShort(&_Controller.TransactOpts, _vaultId)
}

// RedeemShort is a paid mutator transaction binding the contract method 0x97efa942.
//
// Solidity: function redeemShort(uint256 _vaultId) returns()
func (_Controller *ControllerTransactorSession) RedeemShort(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.RedeemShort(&_Controller.TransactOpts, _vaultId)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _vaultId) returns()
func (_Controller *ControllerTransactor) ReduceDebt(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "reduceDebt", _vaultId)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _vaultId) returns()
func (_Controller *ControllerSession) ReduceDebt(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.ReduceDebt(&_Controller.TransactOpts, _vaultId)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _vaultId) returns()
func (_Controller *ControllerTransactorSession) ReduceDebt(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.ReduceDebt(&_Controller.TransactOpts, _vaultId)
}

// ReduceDebtShutdown is a paid mutator transaction binding the contract method 0xfbfc6bc0.
//
// Solidity: function reduceDebtShutdown(uint256 _vaultId) returns()
func (_Controller *ControllerTransactor) ReduceDebtShutdown(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "reduceDebtShutdown", _vaultId)
}

// ReduceDebtShutdown is a paid mutator transaction binding the contract method 0xfbfc6bc0.
//
// Solidity: function reduceDebtShutdown(uint256 _vaultId) returns()
func (_Controller *ControllerSession) ReduceDebtShutdown(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.ReduceDebtShutdown(&_Controller.TransactOpts, _vaultId)
}

// ReduceDebtShutdown is a paid mutator transaction binding the contract method 0xfbfc6bc0.
//
// Solidity: function reduceDebtShutdown(uint256 _vaultId) returns()
func (_Controller *ControllerTransactorSession) ReduceDebtShutdown(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.ReduceDebtShutdown(&_Controller.TransactOpts, _vaultId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newFeeRate) returns()
func (_Controller *ControllerTransactor) SetFeeRate(opts *bind.TransactOpts, _newFeeRate *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setFeeRate", _newFeeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newFeeRate) returns()
func (_Controller *ControllerSession) SetFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetFeeRate(&_Controller.TransactOpts, _newFeeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newFeeRate) returns()
func (_Controller *ControllerTransactorSession) SetFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetFeeRate(&_Controller.TransactOpts, _newFeeRate)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_Controller *ControllerTransactor) SetFeeRecipient(opts *bind.TransactOpts, _newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setFeeRecipient", _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_Controller *ControllerSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetFeeRecipient(&_Controller.TransactOpts, _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_Controller *ControllerTransactorSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetFeeRecipient(&_Controller.TransactOpts, _newFeeRecipient)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_Controller *ControllerTransactor) ShutDown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "shutDown")
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_Controller *ControllerSession) ShutDown() (*types.Transaction, error) {
	return _Controller.Contract.ShutDown(&_Controller.TransactOpts)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_Controller *ControllerTransactorSession) ShutDown() (*types.Transaction, error) {
	return _Controller.Contract.ShutDown(&_Controller.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, newOwner)
}

// UnPauseAnyone is a paid mutator transaction binding the contract method 0x8146b09f.
//
// Solidity: function unPauseAnyone() returns()
func (_Controller *ControllerTransactor) UnPauseAnyone(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "unPauseAnyone")
}

// UnPauseAnyone is a paid mutator transaction binding the contract method 0x8146b09f.
//
// Solidity: function unPauseAnyone() returns()
func (_Controller *ControllerSession) UnPauseAnyone() (*types.Transaction, error) {
	return _Controller.Contract.UnPauseAnyone(&_Controller.TransactOpts)
}

// UnPauseAnyone is a paid mutator transaction binding the contract method 0x8146b09f.
//
// Solidity: function unPauseAnyone() returns()
func (_Controller *ControllerTransactorSession) UnPauseAnyone() (*types.Transaction, error) {
	return _Controller.Contract.UnPauseAnyone(&_Controller.TransactOpts)
}

// UnPauseOwner is a paid mutator transaction binding the contract method 0x07633669.
//
// Solidity: function unPauseOwner() returns()
func (_Controller *ControllerTransactor) UnPauseOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "unPauseOwner")
}

// UnPauseOwner is a paid mutator transaction binding the contract method 0x07633669.
//
// Solidity: function unPauseOwner() returns()
func (_Controller *ControllerSession) UnPauseOwner() (*types.Transaction, error) {
	return _Controller.Contract.UnPauseOwner(&_Controller.TransactOpts)
}

// UnPauseOwner is a paid mutator transaction binding the contract method 0x07633669.
//
// Solidity: function unPauseOwner() returns()
func (_Controller *ControllerTransactorSession) UnPauseOwner() (*types.Transaction, error) {
	return _Controller.Contract.UnPauseOwner(&_Controller.TransactOpts)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xc65a391d.
//
// Solidity: function updateOperator(uint256 _vaultId, address _operator) returns()
func (_Controller *ControllerTransactor) UpdateOperator(opts *bind.TransactOpts, _vaultId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "updateOperator", _vaultId, _operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xc65a391d.
//
// Solidity: function updateOperator(uint256 _vaultId, address _operator) returns()
func (_Controller *ControllerSession) UpdateOperator(_vaultId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Controller.Contract.UpdateOperator(&_Controller.TransactOpts, _vaultId, _operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xc65a391d.
//
// Solidity: function updateOperator(uint256 _vaultId, address _operator) returns()
func (_Controller *ControllerTransactorSession) UpdateOperator(_vaultId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Controller.Contract.UpdateOperator(&_Controller.TransactOpts, _vaultId, _operator)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _vaultId, uint256 _amount) returns()
func (_Controller *ControllerTransactor) Withdraw(opts *bind.TransactOpts, _vaultId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "withdraw", _vaultId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _vaultId, uint256 _amount) returns()
func (_Controller *ControllerSession) Withdraw(_vaultId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Withdraw(&_Controller.TransactOpts, _vaultId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _vaultId, uint256 _amount) returns()
func (_Controller *ControllerTransactorSession) Withdraw(_vaultId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Withdraw(&_Controller.TransactOpts, _vaultId, _amount)
}

// WithdrawUniPositionToken is a paid mutator transaction binding the contract method 0x713d517f.
//
// Solidity: function withdrawUniPositionToken(uint256 _vaultId) returns()
func (_Controller *ControllerTransactor) WithdrawUniPositionToken(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "withdrawUniPositionToken", _vaultId)
}

// WithdrawUniPositionToken is a paid mutator transaction binding the contract method 0x713d517f.
//
// Solidity: function withdrawUniPositionToken(uint256 _vaultId) returns()
func (_Controller *ControllerSession) WithdrawUniPositionToken(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.WithdrawUniPositionToken(&_Controller.TransactOpts, _vaultId)
}

// WithdrawUniPositionToken is a paid mutator transaction binding the contract method 0x713d517f.
//
// Solidity: function withdrawUniPositionToken(uint256 _vaultId) returns()
func (_Controller *ControllerTransactorSession) WithdrawUniPositionToken(_vaultId *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.WithdrawUniPositionToken(&_Controller.TransactOpts, _vaultId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Controller *ControllerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Controller *ControllerSession) Receive() (*types.Transaction, error) {
	return _Controller.Contract.Receive(&_Controller.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Controller *ControllerTransactorSession) Receive() (*types.Transaction, error) {
	return _Controller.Contract.Receive(&_Controller.TransactOpts)
}

// ControllerBurnShortIterator is returned from FilterBurnShort and is used to iterate over the raw logs and unpacked data for BurnShort events raised by the Controller contract.
type ControllerBurnShortIterator struct {
	Event *ControllerBurnShort // Event containing the contract specifics and raw log

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
func (it *ControllerBurnShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerBurnShort)
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
		it.Event = new(ControllerBurnShort)
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
func (it *ControllerBurnShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerBurnShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerBurnShort represents a BurnShort event raised by the Controller contract.
type ControllerBurnShort struct {
	Sender  common.Address
	Amount  *big.Int
	VaultId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnShort is a free log retrieval operation binding the contract event 0xea19ffc45b48de6d95594aacff7214dd24595fdb0c60e98c8f0b269058c2f792.
//
// Solidity: event BurnShort(address sender, uint256 amount, uint256 vaultId)
func (_Controller *ControllerFilterer) FilterBurnShort(opts *bind.FilterOpts) (*ControllerBurnShortIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "BurnShort")
	if err != nil {
		return nil, err
	}
	return &ControllerBurnShortIterator{contract: _Controller.contract, event: "BurnShort", logs: logs, sub: sub}, nil
}

// WatchBurnShort is a free log subscription operation binding the contract event 0xea19ffc45b48de6d95594aacff7214dd24595fdb0c60e98c8f0b269058c2f792.
//
// Solidity: event BurnShort(address sender, uint256 amount, uint256 vaultId)
func (_Controller *ControllerFilterer) WatchBurnShort(opts *bind.WatchOpts, sink chan<- *ControllerBurnShort) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "BurnShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerBurnShort)
				if err := _Controller.contract.UnpackLog(event, "BurnShort", log); err != nil {
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

// ParseBurnShort is a log parse operation binding the contract event 0xea19ffc45b48de6d95594aacff7214dd24595fdb0c60e98c8f0b269058c2f792.
//
// Solidity: event BurnShort(address sender, uint256 amount, uint256 vaultId)
func (_Controller *ControllerFilterer) ParseBurnShort(log types.Log) (*ControllerBurnShort, error) {
	event := new(ControllerBurnShort)
	if err := _Controller.contract.UnpackLog(event, "BurnShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerDepositCollateralIterator is returned from FilterDepositCollateral and is used to iterate over the raw logs and unpacked data for DepositCollateral events raised by the Controller contract.
type ControllerDepositCollateralIterator struct {
	Event *ControllerDepositCollateral // Event containing the contract specifics and raw log

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
func (it *ControllerDepositCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerDepositCollateral)
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
		it.Event = new(ControllerDepositCollateral)
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
func (it *ControllerDepositCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerDepositCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerDepositCollateral represents a DepositCollateral event raised by the Controller contract.
type ControllerDepositCollateral struct {
	Sender  common.Address
	VaultId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositCollateral is a free log retrieval operation binding the contract event 0x3ca13b7aab12bad7472691fe558faa6b25e99099824a0070a88bd5aa84be610f.
//
// Solidity: event DepositCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Controller *ControllerFilterer) FilterDepositCollateral(opts *bind.FilterOpts) (*ControllerDepositCollateralIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "DepositCollateral")
	if err != nil {
		return nil, err
	}
	return &ControllerDepositCollateralIterator{contract: _Controller.contract, event: "DepositCollateral", logs: logs, sub: sub}, nil
}

// WatchDepositCollateral is a free log subscription operation binding the contract event 0x3ca13b7aab12bad7472691fe558faa6b25e99099824a0070a88bd5aa84be610f.
//
// Solidity: event DepositCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Controller *ControllerFilterer) WatchDepositCollateral(opts *bind.WatchOpts, sink chan<- *ControllerDepositCollateral) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "DepositCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerDepositCollateral)
				if err := _Controller.contract.UnpackLog(event, "DepositCollateral", log); err != nil {
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

// ParseDepositCollateral is a log parse operation binding the contract event 0x3ca13b7aab12bad7472691fe558faa6b25e99099824a0070a88bd5aa84be610f.
//
// Solidity: event DepositCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Controller *ControllerFilterer) ParseDepositCollateral(log types.Log) (*ControllerDepositCollateral, error) {
	event := new(ControllerDepositCollateral)
	if err := _Controller.contract.UnpackLog(event, "DepositCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerDepositUniPositionTokenIterator is returned from FilterDepositUniPositionToken and is used to iterate over the raw logs and unpacked data for DepositUniPositionToken events raised by the Controller contract.
type ControllerDepositUniPositionTokenIterator struct {
	Event *ControllerDepositUniPositionToken // Event containing the contract specifics and raw log

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
func (it *ControllerDepositUniPositionTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerDepositUniPositionToken)
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
		it.Event = new(ControllerDepositUniPositionToken)
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
func (it *ControllerDepositUniPositionTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerDepositUniPositionTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerDepositUniPositionToken represents a DepositUniPositionToken event raised by the Controller contract.
type ControllerDepositUniPositionToken struct {
	Sender  common.Address
	VaultId *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositUniPositionToken is a free log retrieval operation binding the contract event 0x3917c2f26ce18614e3aedd1289da672ef6563c5c295f49e9b1697ae0ad315562.
//
// Solidity: event DepositUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Controller *ControllerFilterer) FilterDepositUniPositionToken(opts *bind.FilterOpts) (*ControllerDepositUniPositionTokenIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "DepositUniPositionToken")
	if err != nil {
		return nil, err
	}
	return &ControllerDepositUniPositionTokenIterator{contract: _Controller.contract, event: "DepositUniPositionToken", logs: logs, sub: sub}, nil
}

// WatchDepositUniPositionToken is a free log subscription operation binding the contract event 0x3917c2f26ce18614e3aedd1289da672ef6563c5c295f49e9b1697ae0ad315562.
//
// Solidity: event DepositUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Controller *ControllerFilterer) WatchDepositUniPositionToken(opts *bind.WatchOpts, sink chan<- *ControllerDepositUniPositionToken) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "DepositUniPositionToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerDepositUniPositionToken)
				if err := _Controller.contract.UnpackLog(event, "DepositUniPositionToken", log); err != nil {
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

// ParseDepositUniPositionToken is a log parse operation binding the contract event 0x3917c2f26ce18614e3aedd1289da672ef6563c5c295f49e9b1697ae0ad315562.
//
// Solidity: event DepositUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Controller *ControllerFilterer) ParseDepositUniPositionToken(log types.Log) (*ControllerDepositUniPositionToken, error) {
	event := new(ControllerDepositUniPositionToken)
	if err := _Controller.contract.UnpackLog(event, "DepositUniPositionToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerFeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the Controller contract.
type ControllerFeeRateUpdatedIterator struct {
	Event *ControllerFeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *ControllerFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerFeeRateUpdated)
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
		it.Event = new(ControllerFeeRateUpdated)
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
func (it *ControllerFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerFeeRateUpdated represents a FeeRateUpdated event raised by the Controller contract.
type ControllerFeeRateUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFee, uint256 newFee)
func (_Controller *ControllerFilterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*ControllerFeeRateUpdatedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &ControllerFeeRateUpdatedIterator{contract: _Controller.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFee, uint256 newFee)
func (_Controller *ControllerFilterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *ControllerFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerFeeRateUpdated)
				if err := _Controller.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
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

// ParseFeeRateUpdated is a log parse operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFee, uint256 newFee)
func (_Controller *ControllerFilterer) ParseFeeRateUpdated(log types.Log) (*ControllerFeeRateUpdated, error) {
	event := new(ControllerFeeRateUpdated)
	if err := _Controller.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the Controller contract.
type ControllerFeeRecipientUpdatedIterator struct {
	Event *ControllerFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *ControllerFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerFeeRecipientUpdated)
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
		it.Event = new(ControllerFeeRecipientUpdated)
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
func (it *ControllerFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the Controller contract.
type ControllerFeeRecipientUpdated struct {
	OldFeeRecipient common.Address
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address oldFeeRecipient, address newFeeRecipient)
func (_Controller *ControllerFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts) (*ControllerFeeRecipientUpdatedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return &ControllerFeeRecipientUpdatedIterator{contract: _Controller.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address oldFeeRecipient, address newFeeRecipient)
func (_Controller *ControllerFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *ControllerFeeRecipientUpdated) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerFeeRecipientUpdated)
				if err := _Controller.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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

// ParseFeeRecipientUpdated is a log parse operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address oldFeeRecipient, address newFeeRecipient)
func (_Controller *ControllerFilterer) ParseFeeRecipientUpdated(log types.Log) (*ControllerFeeRecipientUpdated, error) {
	event := new(ControllerFeeRecipientUpdated)
	if err := _Controller.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Controller contract.
type ControllerLiquidateIterator struct {
	Event *ControllerLiquidate // Event containing the contract specifics and raw log

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
func (it *ControllerLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerLiquidate)
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
		it.Event = new(ControllerLiquidate)
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
func (it *ControllerLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerLiquidate represents a Liquidate event raised by the Controller contract.
type ControllerLiquidate struct {
	Liquidator     common.Address
	VaultId        *big.Int
	DebtAmount     *big.Int
	CollateralPaid *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x158ba9ab7bbbd08eeffa4753bad41f4d450e24831d293427308badf3eadd8c76.
//
// Solidity: event Liquidate(address liquidator, uint256 vaultId, uint256 debtAmount, uint256 collateralPaid)
func (_Controller *ControllerFilterer) FilterLiquidate(opts *bind.FilterOpts) (*ControllerLiquidateIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "Liquidate")
	if err != nil {
		return nil, err
	}
	return &ControllerLiquidateIterator{contract: _Controller.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x158ba9ab7bbbd08eeffa4753bad41f4d450e24831d293427308badf3eadd8c76.
//
// Solidity: event Liquidate(address liquidator, uint256 vaultId, uint256 debtAmount, uint256 collateralPaid)
func (_Controller *ControllerFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *ControllerLiquidate) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "Liquidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerLiquidate)
				if err := _Controller.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0x158ba9ab7bbbd08eeffa4753bad41f4d450e24831d293427308badf3eadd8c76.
//
// Solidity: event Liquidate(address liquidator, uint256 vaultId, uint256 debtAmount, uint256 collateralPaid)
func (_Controller *ControllerFilterer) ParseLiquidate(log types.Log) (*ControllerLiquidate, error) {
	event := new(ControllerLiquidate)
	if err := _Controller.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerMintShortIterator is returned from FilterMintShort and is used to iterate over the raw logs and unpacked data for MintShort events raised by the Controller contract.
type ControllerMintShortIterator struct {
	Event *ControllerMintShort // Event containing the contract specifics and raw log

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
func (it *ControllerMintShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerMintShort)
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
		it.Event = new(ControllerMintShort)
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
func (it *ControllerMintShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerMintShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerMintShort represents a MintShort event raised by the Controller contract.
type ControllerMintShort struct {
	Sender  common.Address
	Amount  *big.Int
	VaultId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintShort is a free log retrieval operation binding the contract event 0xb19fa182730a088464dad0e9e0badeb470d0d8d937d854f5caf15c6ad1992c36.
//
// Solidity: event MintShort(address sender, uint256 amount, uint256 vaultId)
func (_Controller *ControllerFilterer) FilterMintShort(opts *bind.FilterOpts) (*ControllerMintShortIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "MintShort")
	if err != nil {
		return nil, err
	}
	return &ControllerMintShortIterator{contract: _Controller.contract, event: "MintShort", logs: logs, sub: sub}, nil
}

// WatchMintShort is a free log subscription operation binding the contract event 0xb19fa182730a088464dad0e9e0badeb470d0d8d937d854f5caf15c6ad1992c36.
//
// Solidity: event MintShort(address sender, uint256 amount, uint256 vaultId)
func (_Controller *ControllerFilterer) WatchMintShort(opts *bind.WatchOpts, sink chan<- *ControllerMintShort) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "MintShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerMintShort)
				if err := _Controller.contract.UnpackLog(event, "MintShort", log); err != nil {
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

// ParseMintShort is a log parse operation binding the contract event 0xb19fa182730a088464dad0e9e0badeb470d0d8d937d854f5caf15c6ad1992c36.
//
// Solidity: event MintShort(address sender, uint256 amount, uint256 vaultId)
func (_Controller *ControllerFilterer) ParseMintShort(log types.Log) (*ControllerMintShort, error) {
	event := new(ControllerMintShort)
	if err := _Controller.contract.UnpackLog(event, "MintShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNormalizationFactorUpdatedIterator is returned from FilterNormalizationFactorUpdated and is used to iterate over the raw logs and unpacked data for NormalizationFactorUpdated events raised by the Controller contract.
type ControllerNormalizationFactorUpdatedIterator struct {
	Event *ControllerNormalizationFactorUpdated // Event containing the contract specifics and raw log

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
func (it *ControllerNormalizationFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNormalizationFactorUpdated)
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
		it.Event = new(ControllerNormalizationFactorUpdated)
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
func (it *ControllerNormalizationFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNormalizationFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNormalizationFactorUpdated represents a NormalizationFactorUpdated event raised by the Controller contract.
type ControllerNormalizationFactorUpdated struct {
	OldNormFactor             *big.Int
	NewNormFactor             *big.Int
	LastModificationTimestamp *big.Int
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterNormalizationFactorUpdated is a free log retrieval operation binding the contract event 0x339e53729b0447795ff69e70a74fed98fc7fef6fe94b7521099b32f0f8de4822.
//
// Solidity: event NormalizationFactorUpdated(uint256 oldNormFactor, uint256 newNormFactor, uint256 lastModificationTimestamp, uint256 timestamp)
func (_Controller *ControllerFilterer) FilterNormalizationFactorUpdated(opts *bind.FilterOpts) (*ControllerNormalizationFactorUpdatedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NormalizationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &ControllerNormalizationFactorUpdatedIterator{contract: _Controller.contract, event: "NormalizationFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchNormalizationFactorUpdated is a free log subscription operation binding the contract event 0x339e53729b0447795ff69e70a74fed98fc7fef6fe94b7521099b32f0f8de4822.
//
// Solidity: event NormalizationFactorUpdated(uint256 oldNormFactor, uint256 newNormFactor, uint256 lastModificationTimestamp, uint256 timestamp)
func (_Controller *ControllerFilterer) WatchNormalizationFactorUpdated(opts *bind.WatchOpts, sink chan<- *ControllerNormalizationFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NormalizationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNormalizationFactorUpdated)
				if err := _Controller.contract.UnpackLog(event, "NormalizationFactorUpdated", log); err != nil {
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

// ParseNormalizationFactorUpdated is a log parse operation binding the contract event 0x339e53729b0447795ff69e70a74fed98fc7fef6fe94b7521099b32f0f8de4822.
//
// Solidity: event NormalizationFactorUpdated(uint256 oldNormFactor, uint256 newNormFactor, uint256 lastModificationTimestamp, uint256 timestamp)
func (_Controller *ControllerFilterer) ParseNormalizationFactorUpdated(log types.Log) (*ControllerNormalizationFactorUpdated, error) {
	event := new(ControllerNormalizationFactorUpdated)
	if err := _Controller.contract.UnpackLog(event, "NormalizationFactorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerOpenVaultIterator is returned from FilterOpenVault and is used to iterate over the raw logs and unpacked data for OpenVault events raised by the Controller contract.
type ControllerOpenVaultIterator struct {
	Event *ControllerOpenVault // Event containing the contract specifics and raw log

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
func (it *ControllerOpenVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerOpenVault)
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
		it.Event = new(ControllerOpenVault)
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
func (it *ControllerOpenVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerOpenVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerOpenVault represents a OpenVault event raised by the Controller contract.
type ControllerOpenVault struct {
	Sender  common.Address
	VaultId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenVault is a free log retrieval operation binding the contract event 0x25ff1e0131762176a9084e92880f880f39d6d0e62134607f37e631efe1bad871.
//
// Solidity: event OpenVault(address sender, uint256 vaultId)
func (_Controller *ControllerFilterer) FilterOpenVault(opts *bind.FilterOpts) (*ControllerOpenVaultIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "OpenVault")
	if err != nil {
		return nil, err
	}
	return &ControllerOpenVaultIterator{contract: _Controller.contract, event: "OpenVault", logs: logs, sub: sub}, nil
}

// WatchOpenVault is a free log subscription operation binding the contract event 0x25ff1e0131762176a9084e92880f880f39d6d0e62134607f37e631efe1bad871.
//
// Solidity: event OpenVault(address sender, uint256 vaultId)
func (_Controller *ControllerFilterer) WatchOpenVault(opts *bind.WatchOpts, sink chan<- *ControllerOpenVault) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "OpenVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerOpenVault)
				if err := _Controller.contract.UnpackLog(event, "OpenVault", log); err != nil {
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

// ParseOpenVault is a log parse operation binding the contract event 0x25ff1e0131762176a9084e92880f880f39d6d0e62134607f37e631efe1bad871.
//
// Solidity: event OpenVault(address sender, uint256 vaultId)
func (_Controller *ControllerFilterer) ParseOpenVault(log types.Log) (*ControllerOpenVault, error) {
	event := new(ControllerOpenVault)
	if err := _Controller.contract.UnpackLog(event, "OpenVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Controller contract.
type ControllerOwnershipTransferredIterator struct {
	Event *ControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerOwnershipTransferred)
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
		it.Event = new(ControllerOwnershipTransferred)
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
func (it *ControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerOwnershipTransferred represents a OwnershipTransferred event raised by the Controller contract.
type ControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []any
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []any
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerOwnershipTransferredIterator{contract: _Controller.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []any
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []any
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerOwnershipTransferred)
				if err := _Controller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Controller *ControllerFilterer) ParseOwnershipTransferred(log types.Log) (*ControllerOwnershipTransferred, error) {
	event := new(ControllerOwnershipTransferred)
	if err := _Controller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Controller contract.
type ControllerPausedIterator struct {
	Event *ControllerPaused // Event containing the contract specifics and raw log

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
func (it *ControllerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerPaused)
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
		it.Event = new(ControllerPaused)
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
func (it *ControllerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerPaused represents a Paused event raised by the Controller contract.
type ControllerPaused struct {
	PausesLeft *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pausesLeft)
func (_Controller *ControllerFilterer) FilterPaused(opts *bind.FilterOpts) (*ControllerPausedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ControllerPausedIterator{contract: _Controller.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pausesLeft)
func (_Controller *ControllerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ControllerPaused) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerPaused)
				if err := _Controller.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pausesLeft)
func (_Controller *ControllerFilterer) ParsePaused(log types.Log) (*ControllerPaused, error) {
	event := new(ControllerPaused)
	if err := _Controller.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerRedeemLongIterator is returned from FilterRedeemLong and is used to iterate over the raw logs and unpacked data for RedeemLong events raised by the Controller contract.
type ControllerRedeemLongIterator struct {
	Event *ControllerRedeemLong // Event containing the contract specifics and raw log

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
func (it *ControllerRedeemLongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRedeemLong)
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
		it.Event = new(ControllerRedeemLong)
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
func (it *ControllerRedeemLongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRedeemLongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRedeemLong represents a RedeemLong event raised by the Controller contract.
type ControllerRedeemLong struct {
	Sender           common.Address
	WPowerPerpAmount *big.Int
	PayoutAmount     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRedeemLong is a free log retrieval operation binding the contract event 0x2131ef4f2f82ca75fe7d2e646ebfa45b6be25e53510c829629c76b641500ec67.
//
// Solidity: event RedeemLong(address sender, uint256 wPowerPerpAmount, uint256 payoutAmount)
func (_Controller *ControllerFilterer) FilterRedeemLong(opts *bind.FilterOpts) (*ControllerRedeemLongIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "RedeemLong")
	if err != nil {
		return nil, err
	}
	return &ControllerRedeemLongIterator{contract: _Controller.contract, event: "RedeemLong", logs: logs, sub: sub}, nil
}

// WatchRedeemLong is a free log subscription operation binding the contract event 0x2131ef4f2f82ca75fe7d2e646ebfa45b6be25e53510c829629c76b641500ec67.
//
// Solidity: event RedeemLong(address sender, uint256 wPowerPerpAmount, uint256 payoutAmount)
func (_Controller *ControllerFilterer) WatchRedeemLong(opts *bind.WatchOpts, sink chan<- *ControllerRedeemLong) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "RedeemLong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRedeemLong)
				if err := _Controller.contract.UnpackLog(event, "RedeemLong", log); err != nil {
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

// ParseRedeemLong is a log parse operation binding the contract event 0x2131ef4f2f82ca75fe7d2e646ebfa45b6be25e53510c829629c76b641500ec67.
//
// Solidity: event RedeemLong(address sender, uint256 wPowerPerpAmount, uint256 payoutAmount)
func (_Controller *ControllerFilterer) ParseRedeemLong(log types.Log) (*ControllerRedeemLong, error) {
	event := new(ControllerRedeemLong)
	if err := _Controller.contract.UnpackLog(event, "RedeemLong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerRedeemShortIterator is returned from FilterRedeemShort and is used to iterate over the raw logs and unpacked data for RedeemShort events raised by the Controller contract.
type ControllerRedeemShortIterator struct {
	Event *ControllerRedeemShort // Event containing the contract specifics and raw log

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
func (it *ControllerRedeemShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRedeemShort)
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
		it.Event = new(ControllerRedeemShort)
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
func (it *ControllerRedeemShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRedeemShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRedeemShort represents a RedeemShort event raised by the Controller contract.
type ControllerRedeemShort struct {
	Sender           common.Address
	VauldId          *big.Int
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRedeemShort is a free log retrieval operation binding the contract event 0x7dff8cdaec6a8d4d1ad32d3c947ed0f0281c3d6456621ef928defae96ec6cddb.
//
// Solidity: event RedeemShort(address sender, uint256 vauldId, uint256 collateralAmount)
func (_Controller *ControllerFilterer) FilterRedeemShort(opts *bind.FilterOpts) (*ControllerRedeemShortIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "RedeemShort")
	if err != nil {
		return nil, err
	}
	return &ControllerRedeemShortIterator{contract: _Controller.contract, event: "RedeemShort", logs: logs, sub: sub}, nil
}

// WatchRedeemShort is a free log subscription operation binding the contract event 0x7dff8cdaec6a8d4d1ad32d3c947ed0f0281c3d6456621ef928defae96ec6cddb.
//
// Solidity: event RedeemShort(address sender, uint256 vauldId, uint256 collateralAmount)
func (_Controller *ControllerFilterer) WatchRedeemShort(opts *bind.WatchOpts, sink chan<- *ControllerRedeemShort) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "RedeemShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRedeemShort)
				if err := _Controller.contract.UnpackLog(event, "RedeemShort", log); err != nil {
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

// ParseRedeemShort is a log parse operation binding the contract event 0x7dff8cdaec6a8d4d1ad32d3c947ed0f0281c3d6456621ef928defae96ec6cddb.
//
// Solidity: event RedeemShort(address sender, uint256 vauldId, uint256 collateralAmount)
func (_Controller *ControllerFilterer) ParseRedeemShort(log types.Log) (*ControllerRedeemShort, error) {
	event := new(ControllerRedeemShort)
	if err := _Controller.contract.UnpackLog(event, "RedeemShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerReduceDebtIterator is returned from FilterReduceDebt and is used to iterate over the raw logs and unpacked data for ReduceDebt events raised by the Controller contract.
type ControllerReduceDebtIterator struct {
	Event *ControllerReduceDebt // Event containing the contract specifics and raw log

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
func (it *ControllerReduceDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerReduceDebt)
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
		it.Event = new(ControllerReduceDebt)
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
func (it *ControllerReduceDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerReduceDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerReduceDebt represents a ReduceDebt event raised by the Controller contract.
type ControllerReduceDebt struct {
	Sender             common.Address
	VaultId            *big.Int
	EthRedeemed        *big.Int
	WPowerPerpRedeemed *big.Int
	WPowerPerpBurned   *big.Int
	WPowerPerpExcess   *big.Int
	Bounty             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReduceDebt is a free log retrieval operation binding the contract event 0xfd0ae2fd36bd955810ae42615bc5ff277c0d0dfcb930f06c9f1777c0ef0752e3.
//
// Solidity: event ReduceDebt(address sender, uint256 vaultId, uint256 ethRedeemed, uint256 wPowerPerpRedeemed, uint256 wPowerPerpBurned, uint256 wPowerPerpExcess, uint256 bounty)
func (_Controller *ControllerFilterer) FilterReduceDebt(opts *bind.FilterOpts) (*ControllerReduceDebtIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "ReduceDebt")
	if err != nil {
		return nil, err
	}
	return &ControllerReduceDebtIterator{contract: _Controller.contract, event: "ReduceDebt", logs: logs, sub: sub}, nil
}

// WatchReduceDebt is a free log subscription operation binding the contract event 0xfd0ae2fd36bd955810ae42615bc5ff277c0d0dfcb930f06c9f1777c0ef0752e3.
//
// Solidity: event ReduceDebt(address sender, uint256 vaultId, uint256 ethRedeemed, uint256 wPowerPerpRedeemed, uint256 wPowerPerpBurned, uint256 wPowerPerpExcess, uint256 bounty)
func (_Controller *ControllerFilterer) WatchReduceDebt(opts *bind.WatchOpts, sink chan<- *ControllerReduceDebt) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "ReduceDebt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerReduceDebt)
				if err := _Controller.contract.UnpackLog(event, "ReduceDebt", log); err != nil {
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

// ParseReduceDebt is a log parse operation binding the contract event 0xfd0ae2fd36bd955810ae42615bc5ff277c0d0dfcb930f06c9f1777c0ef0752e3.
//
// Solidity: event ReduceDebt(address sender, uint256 vaultId, uint256 ethRedeemed, uint256 wPowerPerpRedeemed, uint256 wPowerPerpBurned, uint256 wPowerPerpExcess, uint256 bounty)
func (_Controller *ControllerFilterer) ParseReduceDebt(log types.Log) (*ControllerReduceDebt, error) {
	event := new(ControllerReduceDebt)
	if err := _Controller.contract.UnpackLog(event, "ReduceDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the Controller contract.
type ControllerShutdownIterator struct {
	Event *ControllerShutdown // Event containing the contract specifics and raw log

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
func (it *ControllerShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerShutdown)
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
		it.Event = new(ControllerShutdown)
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
func (it *ControllerShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerShutdown represents a Shutdown event raised by the Controller contract.
type ControllerShutdown struct {
	IndexForSettlement *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 indexForSettlement)
func (_Controller *ControllerFilterer) FilterShutdown(opts *bind.FilterOpts) (*ControllerShutdownIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &ControllerShutdownIterator{contract: _Controller.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 indexForSettlement)
func (_Controller *ControllerFilterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *ControllerShutdown) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerShutdown)
				if err := _Controller.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 indexForSettlement)
func (_Controller *ControllerFilterer) ParseShutdown(log types.Log) (*ControllerShutdown, error) {
	event := new(ControllerShutdown)
	if err := _Controller.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerUnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the Controller contract.
type ControllerUnPausedIterator struct {
	Event *ControllerUnPaused // Event containing the contract specifics and raw log

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
func (it *ControllerUnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerUnPaused)
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
		it.Event = new(ControllerUnPaused)
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
func (it *ControllerUnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerUnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerUnPaused represents a UnPaused event raised by the Controller contract.
type ControllerUnPaused struct {
	Unpauser common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address unpauser)
func (_Controller *ControllerFilterer) FilterUnPaused(opts *bind.FilterOpts) (*ControllerUnPausedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return &ControllerUnPausedIterator{contract: _Controller.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address unpauser)
func (_Controller *ControllerFilterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *ControllerUnPaused) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerUnPaused)
				if err := _Controller.contract.UnpackLog(event, "UnPaused", log); err != nil {
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

// ParseUnPaused is a log parse operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address unpauser)
func (_Controller *ControllerFilterer) ParseUnPaused(log types.Log) (*ControllerUnPaused, error) {
	event := new(ControllerUnPaused)
	if err := _Controller.contract.UnpackLog(event, "UnPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerUpdateOperatorIterator is returned from FilterUpdateOperator and is used to iterate over the raw logs and unpacked data for UpdateOperator events raised by the Controller contract.
type ControllerUpdateOperatorIterator struct {
	Event *ControllerUpdateOperator // Event containing the contract specifics and raw log

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
func (it *ControllerUpdateOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerUpdateOperator)
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
		it.Event = new(ControllerUpdateOperator)
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
func (it *ControllerUpdateOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerUpdateOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerUpdateOperator represents a UpdateOperator event raised by the Controller contract.
type ControllerUpdateOperator struct {
	Sender   common.Address
	VaultId  *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateOperator is a free log retrieval operation binding the contract event 0x3137fc9cd2e33c34f86e29c24d81f3c75b0bce639d3c4ed0d31eeff1160a7ff5.
//
// Solidity: event UpdateOperator(address sender, uint256 vaultId, address operator)
func (_Controller *ControllerFilterer) FilterUpdateOperator(opts *bind.FilterOpts) (*ControllerUpdateOperatorIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "UpdateOperator")
	if err != nil {
		return nil, err
	}
	return &ControllerUpdateOperatorIterator{contract: _Controller.contract, event: "UpdateOperator", logs: logs, sub: sub}, nil
}

// WatchUpdateOperator is a free log subscription operation binding the contract event 0x3137fc9cd2e33c34f86e29c24d81f3c75b0bce639d3c4ed0d31eeff1160a7ff5.
//
// Solidity: event UpdateOperator(address sender, uint256 vaultId, address operator)
func (_Controller *ControllerFilterer) WatchUpdateOperator(opts *bind.WatchOpts, sink chan<- *ControllerUpdateOperator) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "UpdateOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerUpdateOperator)
				if err := _Controller.contract.UnpackLog(event, "UpdateOperator", log); err != nil {
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

// ParseUpdateOperator is a log parse operation binding the contract event 0x3137fc9cd2e33c34f86e29c24d81f3c75b0bce639d3c4ed0d31eeff1160a7ff5.
//
// Solidity: event UpdateOperator(address sender, uint256 vaultId, address operator)
func (_Controller *ControllerFilterer) ParseUpdateOperator(log types.Log) (*ControllerUpdateOperator, error) {
	event := new(ControllerUpdateOperator)
	if err := _Controller.contract.UnpackLog(event, "UpdateOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerWithdrawCollateralIterator is returned from FilterWithdrawCollateral and is used to iterate over the raw logs and unpacked data for WithdrawCollateral events raised by the Controller contract.
type ControllerWithdrawCollateralIterator struct {
	Event *ControllerWithdrawCollateral // Event containing the contract specifics and raw log

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
func (it *ControllerWithdrawCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerWithdrawCollateral)
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
		it.Event = new(ControllerWithdrawCollateral)
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
func (it *ControllerWithdrawCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerWithdrawCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerWithdrawCollateral represents a WithdrawCollateral event raised by the Controller contract.
type ControllerWithdrawCollateral struct {
	Sender  common.Address
	VaultId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCollateral is a free log retrieval operation binding the contract event 0x627a692d5a03ab34732c0d2aa319f3ecdebdc4528f383eabcb25441dc0a70cfb.
//
// Solidity: event WithdrawCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Controller *ControllerFilterer) FilterWithdrawCollateral(opts *bind.FilterOpts) (*ControllerWithdrawCollateralIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "WithdrawCollateral")
	if err != nil {
		return nil, err
	}
	return &ControllerWithdrawCollateralIterator{contract: _Controller.contract, event: "WithdrawCollateral", logs: logs, sub: sub}, nil
}

// WatchWithdrawCollateral is a free log subscription operation binding the contract event 0x627a692d5a03ab34732c0d2aa319f3ecdebdc4528f383eabcb25441dc0a70cfb.
//
// Solidity: event WithdrawCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Controller *ControllerFilterer) WatchWithdrawCollateral(opts *bind.WatchOpts, sink chan<- *ControllerWithdrawCollateral) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "WithdrawCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerWithdrawCollateral)
				if err := _Controller.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
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

// ParseWithdrawCollateral is a log parse operation binding the contract event 0x627a692d5a03ab34732c0d2aa319f3ecdebdc4528f383eabcb25441dc0a70cfb.
//
// Solidity: event WithdrawCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Controller *ControllerFilterer) ParseWithdrawCollateral(log types.Log) (*ControllerWithdrawCollateral, error) {
	event := new(ControllerWithdrawCollateral)
	if err := _Controller.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerWithdrawUniPositionTokenIterator is returned from FilterWithdrawUniPositionToken and is used to iterate over the raw logs and unpacked data for WithdrawUniPositionToken events raised by the Controller contract.
type ControllerWithdrawUniPositionTokenIterator struct {
	Event *ControllerWithdrawUniPositionToken // Event containing the contract specifics and raw log

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
func (it *ControllerWithdrawUniPositionTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerWithdrawUniPositionToken)
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
		it.Event = new(ControllerWithdrawUniPositionToken)
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
func (it *ControllerWithdrawUniPositionTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerWithdrawUniPositionTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerWithdrawUniPositionToken represents a WithdrawUniPositionToken event raised by the Controller contract.
type ControllerWithdrawUniPositionToken struct {
	Sender  common.Address
	VaultId *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawUniPositionToken is a free log retrieval operation binding the contract event 0xe59f38fa1264fc25c9f0185eee136eaf810d90b8e7293b342e4037c68720177a.
//
// Solidity: event WithdrawUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Controller *ControllerFilterer) FilterWithdrawUniPositionToken(opts *bind.FilterOpts) (*ControllerWithdrawUniPositionTokenIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "WithdrawUniPositionToken")
	if err != nil {
		return nil, err
	}
	return &ControllerWithdrawUniPositionTokenIterator{contract: _Controller.contract, event: "WithdrawUniPositionToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawUniPositionToken is a free log subscription operation binding the contract event 0xe59f38fa1264fc25c9f0185eee136eaf810d90b8e7293b342e4037c68720177a.
//
// Solidity: event WithdrawUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Controller *ControllerFilterer) WatchWithdrawUniPositionToken(opts *bind.WatchOpts, sink chan<- *ControllerWithdrawUniPositionToken) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "WithdrawUniPositionToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerWithdrawUniPositionToken)
				if err := _Controller.contract.UnpackLog(event, "WithdrawUniPositionToken", log); err != nil {
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

// ParseWithdrawUniPositionToken is a log parse operation binding the contract event 0xe59f38fa1264fc25c9f0185eee136eaf810d90b8e7293b342e4037c68720177a.
//
// Solidity: event WithdrawUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Controller *ControllerFilterer) ParseWithdrawUniPositionToken(log types.Log) (*ControllerWithdrawUniPositionToken, error) {
	event := new(ControllerWithdrawUniPositionToken)
	if err := _Controller.contract.UnpackLog(event, "WithdrawUniPositionToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
