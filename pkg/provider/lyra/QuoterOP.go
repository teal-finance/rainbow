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

// QuoterOPMetaData contains all meta data concerning the QuoterOP contract.
var QuoterOPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lyraRegister\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardAlreadySettled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"}],\"name\":\"BoardExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"enumIOptionMarket.NonZeroValues\",\"name\":\"valueType\",\"type\":\"uint8\"}],\"name\":\"ExpectedNonZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"InvalidStrikeId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIOptionMarket\",\"name\":\"_optionMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fullQuotes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOptionMarket\",\"name\":\"_optionMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalPremium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[{\"internalType\":\"contractILyraRegister\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// QuoterOPABI is the input ABI used to generate the binding from.
// Deprecated: Use QuoterOPMetaData.ABI instead.
var QuoterOPABI = QuoterOPMetaData.ABI

// QuoterOP is an auto generated Go binding around an Ethereum contract.
type QuoterOP struct {
	QuoterOPCaller     // Read-only binding to the contract
	QuoterOPTransactor // Write-only binding to the contract
	QuoterOPFilterer   // Log filterer for contract events
}

// QuoterOPCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuoterOPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterOPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuoterOPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterOPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuoterOPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterOPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuoterOPSession struct {
	Contract     *QuoterOP         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterOPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuoterOPCallerSession struct {
	Contract *QuoterOPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QuoterOPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuoterOPTransactorSession struct {
	Contract     *QuoterOPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QuoterOPRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuoterOPRaw struct {
	Contract *QuoterOP // Generic contract binding to access the raw methods on
}

// QuoterOPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuoterOPCallerRaw struct {
	Contract *QuoterOPCaller // Generic read-only contract binding to access the raw methods on
}

// QuoterOPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuoterOPTransactorRaw struct {
	Contract *QuoterOPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuoterOP creates a new instance of QuoterOP, bound to a specific deployed contract.
func NewQuoterOP(address common.Address, backend bind.ContractBackend) (*QuoterOP, error) {
	contract, err := bindQuoterOP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuoterOP{QuoterOPCaller: QuoterOPCaller{contract: contract}, QuoterOPTransactor: QuoterOPTransactor{contract: contract}, QuoterOPFilterer: QuoterOPFilterer{contract: contract}}, nil
}

// NewQuoterOPCaller creates a new read-only instance of QuoterOP, bound to a specific deployed contract.
func NewQuoterOPCaller(address common.Address, caller bind.ContractCaller) (*QuoterOPCaller, error) {
	contract, err := bindQuoterOP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterOPCaller{contract: contract}, nil
}

// NewQuoterOPTransactor creates a new write-only instance of QuoterOP, bound to a specific deployed contract.
func NewQuoterOPTransactor(address common.Address, transactor bind.ContractTransactor) (*QuoterOPTransactor, error) {
	contract, err := bindQuoterOP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterOPTransactor{contract: contract}, nil
}

// NewQuoterOPFilterer creates a new log filterer instance of QuoterOP, bound to a specific deployed contract.
func NewQuoterOPFilterer(address common.Address, filterer bind.ContractFilterer) (*QuoterOPFilterer, error) {
	contract, err := bindQuoterOP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuoterOPFilterer{contract: contract}, nil
}

// bindQuoterOP binds a generic wrapper to an already deployed contract.
func bindQuoterOP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuoterOPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuoterOP *QuoterOPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuoterOP.Contract.QuoterOPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuoterOP *QuoterOPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuoterOP.Contract.QuoterOPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuoterOP *QuoterOPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuoterOP.Contract.QuoterOPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuoterOP *QuoterOPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuoterOP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuoterOP *QuoterOPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuoterOP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuoterOP *QuoterOPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuoterOP.Contract.contract.Transact(opts, method, params...)
}

// FullQuotes is a free data retrieval call binding the contract method 0xae2275b3.
//
// Solidity: function fullQuotes(address _optionMarket, uint256 strikeId, uint256 iterations, uint256 amount) view returns(uint256[], uint256[])
func (_QuoterOP *QuoterOPCaller) FullQuotes(opts *bind.CallOpts, _optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _QuoterOP.contract.Call(opts, &out, "fullQuotes", _optionMarket, strikeId, iterations, amount)

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
func (_QuoterOP *QuoterOPSession) FullQuotes(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	return _QuoterOP.Contract.FullQuotes(&_QuoterOP.CallOpts, _optionMarket, strikeId, iterations, amount)
}

// FullQuotes is a free data retrieval call binding the contract method 0xae2275b3.
//
// Solidity: function fullQuotes(address _optionMarket, uint256 strikeId, uint256 iterations, uint256 amount) view returns(uint256[], uint256[])
func (_QuoterOP *QuoterOPCallerSession) FullQuotes(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	return _QuoterOP.Contract.FullQuotes(&_QuoterOP.CallOpts, _optionMarket, strikeId, iterations, amount)
}

// Quote is a free data retrieval call binding the contract method 0x4c438ca7.
//
// Solidity: function quote(address _optionMarket, uint256 strikeId, uint256 iterations, uint8 optionType, uint256 amount) view returns(uint256 totalPremium, uint256 totalFee)
func (_QuoterOP *QuoterOPCaller) Quote(opts *bind.CallOpts, _optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	var out []interface{}
	err := _QuoterOP.contract.Call(opts, &out, "quote", _optionMarket, strikeId, iterations, optionType, amount)

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
func (_QuoterOP *QuoterOPSession) Quote(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	return _QuoterOP.Contract.Quote(&_QuoterOP.CallOpts, _optionMarket, strikeId, iterations, optionType, amount)
}

// Quote is a free data retrieval call binding the contract method 0x4c438ca7.
//
// Solidity: function quote(address _optionMarket, uint256 strikeId, uint256 iterations, uint8 optionType, uint256 amount) view returns(uint256 totalPremium, uint256 totalFee)
func (_QuoterOP *QuoterOPCallerSession) Quote(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	return _QuoterOP.Contract.Quote(&_QuoterOP.CallOpts, _optionMarket, strikeId, iterations, optionType, amount)
}

// Register is a free data retrieval call binding the contract method 0x1aa3a008.
//
// Solidity: function register() view returns(address)
func (_QuoterOP *QuoterOPCaller) Register(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuoterOP.contract.Call(opts, &out, "register")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Register is a free data retrieval call binding the contract method 0x1aa3a008.
//
// Solidity: function register() view returns(address)
func (_QuoterOP *QuoterOPSession) Register() (common.Address, error) {
	return _QuoterOP.Contract.Register(&_QuoterOP.CallOpts)
}

// Register is a free data retrieval call binding the contract method 0x1aa3a008.
//
// Solidity: function register() view returns(address)
func (_QuoterOP *QuoterOPCallerSession) Register() (common.Address, error) {
	return _QuoterOP.Contract.Register(&_QuoterOP.CallOpts)
}
