// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lyra

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

// QuoterARBMetaData contains all meta data concerning the QuoterARB contract.
var QuoterARBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lyraRegister\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardAlreadySettled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"}],\"name\":\"BoardExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"enumIOptionMarket.NonZeroValues\",\"name\":\"valueType\",\"type\":\"uint8\"}],\"name\":\"ExpectedNonZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"InvalidStrikeId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIOptionMarket\",\"name\":\"_optionMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fullQuotes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOptionMarket\",\"name\":\"_optionMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalPremium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// QuoterARBABI is the input ABI used to generate the binding from.
// Deprecated: Use QuoterARBMetaData.ABI instead.
var QuoterARBABI = QuoterARBMetaData.ABI

// QuoterARB is an auto generated Go binding around an Ethereum contract.
type QuoterARB struct {
	QuoterARBCaller     // Read-only binding to the contract
	QuoterARBTransactor // Write-only binding to the contract
	QuoterARBFilterer   // Log filterer for contract events
}

// QuoterARBCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuoterARBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterARBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuoterARBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterARBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuoterARBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterARBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuoterARBSession struct {
	Contract     *QuoterARB        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterARBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuoterARBCallerSession struct {
	Contract *QuoterARBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// QuoterARBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuoterARBTransactorSession struct {
	Contract     *QuoterARBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// QuoterARBRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuoterARBRaw struct {
	Contract *QuoterARB // Generic contract binding to access the raw methods on
}

// QuoterARBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuoterARBCallerRaw struct {
	Contract *QuoterARBCaller // Generic read-only contract binding to access the raw methods on
}

// QuoterARBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuoterARBTransactorRaw struct {
	Contract *QuoterARBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuoterARB creates a new instance of QuoterARB, bound to a specific deployed contract.
func NewQuoterARB(address common.Address, backend bind.ContractBackend) (*QuoterARB, error) {
	contract, err := bindQuoterARB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuoterARB{QuoterARBCaller: QuoterARBCaller{contract: contract}, QuoterARBTransactor: QuoterARBTransactor{contract: contract}, QuoterARBFilterer: QuoterARBFilterer{contract: contract}}, nil
}

// NewQuoterARBCaller creates a new read-only instance of QuoterARB, bound to a specific deployed contract.
func NewQuoterARBCaller(address common.Address, caller bind.ContractCaller) (*QuoterARBCaller, error) {
	contract, err := bindQuoterARB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterARBCaller{contract: contract}, nil
}

// NewQuoterARBTransactor creates a new write-only instance of QuoterARB, bound to a specific deployed contract.
func NewQuoterARBTransactor(address common.Address, transactor bind.ContractTransactor) (*QuoterARBTransactor, error) {
	contract, err := bindQuoterARB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterARBTransactor{contract: contract}, nil
}

// NewQuoterARBFilterer creates a new log filterer instance of QuoterARB, bound to a specific deployed contract.
func NewQuoterARBFilterer(address common.Address, filterer bind.ContractFilterer) (*QuoterARBFilterer, error) {
	contract, err := bindQuoterARB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuoterARBFilterer{contract: contract}, nil
}

// bindQuoterARB binds a generic wrapper to an already deployed contract.
func bindQuoterARB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuoterARBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuoterARB *QuoterARBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuoterARB.Contract.QuoterARBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuoterARB *QuoterARBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuoterARB.Contract.QuoterARBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuoterARB *QuoterARBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuoterARB.Contract.QuoterARBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuoterARB *QuoterARBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuoterARB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuoterARB *QuoterARBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuoterARB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuoterARB *QuoterARBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuoterARB.Contract.contract.Transact(opts, method, params...)
}

// FullQuotes is a free data retrieval call binding the contract method 0xae2275b3.
//
// Solidity: function fullQuotes(address _optionMarket, uint256 strikeId, uint256 iterations, uint256 amount) view returns(uint256[], uint256[])
func (_QuoterARB *QuoterARBCaller) FullQuotes(opts *bind.CallOpts, _optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _QuoterARB.contract.Call(opts, &out, "fullQuotes", _optionMarket, strikeId, iterations, amount)

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// FullQuotes is a free data retrieval call binding the contract method 0xae2275b3.
//
// Solidity: function fullQuotes(address _optionMarket, uint256 strikeId, uint256 iterations, uint256 amount) view returns(uint256[], uint256[])
func (_QuoterARB *QuoterARBSession) FullQuotes(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	return _QuoterARB.Contract.FullQuotes(&_QuoterARB.CallOpts, _optionMarket, strikeId, iterations, amount)
}

// FullQuotes is a free data retrieval call binding the contract method 0xae2275b3.
//
// Solidity: function fullQuotes(address _optionMarket, uint256 strikeId, uint256 iterations, uint256 amount) view returns(uint256[], uint256[])
func (_QuoterARB *QuoterARBCallerSession) FullQuotes(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	return _QuoterARB.Contract.FullQuotes(&_QuoterARB.CallOpts, _optionMarket, strikeId, iterations, amount)
}

// Quote is a free data retrieval call binding the contract method 0x4c438ca7.
//
// Solidity: function quote(address _optionMarket, uint256 strikeId, uint256 iterations, uint8 optionType, uint256 amount) view returns(uint256 totalPremium, uint256 totalFee)
func (_QuoterARB *QuoterARBCaller) Quote(opts *bind.CallOpts, _optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	var out []interface{}
	err := _QuoterARB.contract.Call(opts, &out, "quote", _optionMarket, strikeId, iterations, optionType, amount)

	outstruct := new(struct {
		TotalPremium *big.Int
		TotalFee     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalPremium = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Quote is a free data retrieval call binding the contract method 0x4c438ca7.
//
// Solidity: function quote(address _optionMarket, uint256 strikeId, uint256 iterations, uint8 optionType, uint256 amount) view returns(uint256 totalPremium, uint256 totalFee)
func (_QuoterARB *QuoterARBSession) Quote(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	return _QuoterARB.Contract.Quote(&_QuoterARB.CallOpts, _optionMarket, strikeId, iterations, optionType, amount)
}

// Quote is a free data retrieval call binding the contract method 0x4c438ca7.
//
// Solidity: function quote(address _optionMarket, uint256 strikeId, uint256 iterations, uint8 optionType, uint256 amount) view returns(uint256 totalPremium, uint256 totalFee)
func (_QuoterARB *QuoterARBCallerSession) Quote(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	return _QuoterARB.Contract.Quote(&_QuoterARB.CallOpts, _optionMarket, strikeId, iterations, optionType, amount)
}
