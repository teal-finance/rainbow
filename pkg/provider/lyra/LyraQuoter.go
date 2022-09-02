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

// LyraqMetaData contains all meta data concerning the Lyraq contract.
var LyraqMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lyraRegister\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardAlreadySettled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"}],\"name\":\"BoardExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"enumIOptionMarket.NonZeroValues\",\"name\":\"valueType\",\"type\":\"uint8\"}],\"name\":\"ExpectedNonZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"InvalidStrikeId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIOptionMarket\",\"name\":\"_optionMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fullQuotes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOptionMarket\",\"name\":\"_optionMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalPremium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[{\"internalType\":\"contractILyraRegister\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LyraqABI is the input ABI used to generate the binding from.
// Deprecated: Use LyraqMetaData.ABI instead.
var LyraqABI = LyraqMetaData.ABI

// Lyraq is an auto generated Go binding around an Ethereum contract.
type Lyraq struct {
	LyraqCaller     // Read-only binding to the contract
	LyraqTransactor // Write-only binding to the contract
	LyraqFilterer   // Log filterer for contract events
}

// LyraqCaller is an auto generated read-only Go binding around an Ethereum contract.
type LyraqCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LyraqTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LyraqTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LyraqFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LyraqFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LyraqSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LyraqSession struct {
	Contract     *Lyraq            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LyraqCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LyraqCallerSession struct {
	Contract *LyraqCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LyraqTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LyraqTransactorSession struct {
	Contract     *LyraqTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LyraqRaw is an auto generated low-level Go binding around an Ethereum contract.
type LyraqRaw struct {
	Contract *Lyraq // Generic contract binding to access the raw methods on
}

// LyraqCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LyraqCallerRaw struct {
	Contract *LyraqCaller // Generic read-only contract binding to access the raw methods on
}

// LyraqTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LyraqTransactorRaw struct {
	Contract *LyraqTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLyraq creates a new instance of Lyraq, bound to a specific deployed contract.
func NewLyraq(address common.Address, backend bind.ContractBackend) (*Lyraq, error) {
	contract, err := bindLyraq(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lyraq{LyraqCaller: LyraqCaller{contract: contract}, LyraqTransactor: LyraqTransactor{contract: contract}, LyraqFilterer: LyraqFilterer{contract: contract}}, nil
}

// NewLyraqCaller creates a new read-only instance of Lyraq, bound to a specific deployed contract.
func NewLyraqCaller(address common.Address, caller bind.ContractCaller) (*LyraqCaller, error) {
	contract, err := bindLyraq(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LyraqCaller{contract: contract}, nil
}

// NewLyraqTransactor creates a new write-only instance of Lyraq, bound to a specific deployed contract.
func NewLyraqTransactor(address common.Address, transactor bind.ContractTransactor) (*LyraqTransactor, error) {
	contract, err := bindLyraq(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LyraqTransactor{contract: contract}, nil
}

// NewLyraqFilterer creates a new log filterer instance of Lyraq, bound to a specific deployed contract.
func NewLyraqFilterer(address common.Address, filterer bind.ContractFilterer) (*LyraqFilterer, error) {
	contract, err := bindLyraq(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LyraqFilterer{contract: contract}, nil
}

// bindLyraq binds a generic wrapper to an already deployed contract.
func bindLyraq(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LyraqABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lyraq *LyraqRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lyraq.Contract.LyraqCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lyraq *LyraqRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lyraq.Contract.LyraqTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lyraq *LyraqRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lyraq.Contract.LyraqTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lyraq *LyraqCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lyraq.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lyraq *LyraqTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lyraq.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lyraq *LyraqTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lyraq.Contract.contract.Transact(opts, method, params...)
}

// FullQuotes is a free data retrieval call binding the contract method 0xae2275b3.
//
// Solidity: function fullQuotes(address _optionMarket, uint256 strikeId, uint256 iterations, uint256 amount) view returns(uint256[], uint256[])
func (_Lyraq *LyraqCaller) FullQuotes(opts *bind.CallOpts, _optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Lyraq.contract.Call(opts, &out, "fullQuotes", _optionMarket, strikeId, iterations, amount)

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
func (_Lyraq *LyraqSession) FullQuotes(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	return _Lyraq.Contract.FullQuotes(&_Lyraq.CallOpts, _optionMarket, strikeId, iterations, amount)
}

// FullQuotes is a free data retrieval call binding the contract method 0xae2275b3.
//
// Solidity: function fullQuotes(address _optionMarket, uint256 strikeId, uint256 iterations, uint256 amount) view returns(uint256[], uint256[])
func (_Lyraq *LyraqCallerSession) FullQuotes(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	return _Lyraq.Contract.FullQuotes(&_Lyraq.CallOpts, _optionMarket, strikeId, iterations, amount)
}

// Quote is a free data retrieval call binding the contract method 0x4c438ca7.
//
// Solidity: function quote(address _optionMarket, uint256 strikeId, uint256 iterations, uint8 optionType, uint256 amount) view returns(uint256 totalPremium, uint256 totalFee)
func (_Lyraq *LyraqCaller) Quote(opts *bind.CallOpts, _optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	var out []interface{}
	err := _Lyraq.contract.Call(opts, &out, "quote", _optionMarket, strikeId, iterations, optionType, amount)

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
func (_Lyraq *LyraqSession) Quote(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	return _Lyraq.Contract.Quote(&_Lyraq.CallOpts, _optionMarket, strikeId, iterations, optionType, amount)
}

// Quote is a free data retrieval call binding the contract method 0x4c438ca7.
//
// Solidity: function quote(address _optionMarket, uint256 strikeId, uint256 iterations, uint8 optionType, uint256 amount) view returns(uint256 totalPremium, uint256 totalFee)
func (_Lyraq *LyraqCallerSession) Quote(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	return _Lyraq.Contract.Quote(&_Lyraq.CallOpts, _optionMarket, strikeId, iterations, optionType, amount)
}

// Register is a free data retrieval call binding the contract method 0x1aa3a008.
//
// Solidity: function register() view returns(address)
func (_Lyraq *LyraqCaller) Register(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyraq.contract.Call(opts, &out, "register")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Register is a free data retrieval call binding the contract method 0x1aa3a008.
//
// Solidity: function register() view returns(address)
func (_Lyraq *LyraqSession) Register() (common.Address, error) {
	return _Lyraq.Contract.Register(&_Lyraq.CallOpts)
}

// Register is a free data retrieval call binding the contract method 0x1aa3a008.
//
// Solidity: function register() view returns(address)
func (_Lyraq *LyraqCallerSession) Register() (common.Address, error) {
	return _Lyraq.Contract.Register(&_Lyraq.CallOpts)
}
