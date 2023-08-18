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
	_ = abi.ConvertType
)

// QuoterMetaData contains all meta data concerning the Quoter contract.
var QuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lyraRegister\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardAlreadySettled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"}],\"name\":\"BoardExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"enumIOptionMarket.NonZeroValues\",\"name\":\"valueType\",\"type\":\"uint8\"}],\"name\":\"ExpectedNonZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"InvalidStrikeId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIOptionMarket\",\"name\":\"_optionMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fullQuotes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOptionMarket\",\"name\":\"_optionMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalPremium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// QuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use QuoterMetaData.ABI instead.
var QuoterABI = QuoterMetaData.ABI

// Quoter is an auto generated Go binding around an Ethereum contract.
type Quoter struct {
	QuoterCaller     // Read-only binding to the contract
	QuoterTransactor // Write-only binding to the contract
	QuoterFilterer   // Log filterer for contract events
}

// QuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuoterSession struct {
	Contract     *Quoter           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuoterCallerSession struct {
	Contract *QuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuoterTransactorSession struct {
	Contract     *QuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuoterRaw struct {
	Contract *Quoter // Generic contract binding to access the raw methods on
}

// QuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuoterCallerRaw struct {
	Contract *QuoterCaller // Generic read-only contract binding to access the raw methods on
}

// QuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuoterTransactorRaw struct {
	Contract *QuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuoter creates a new instance of Quoter, bound to a specific deployed contract.
func NewQuoter(address common.Address, backend bind.ContractBackend) (*Quoter, error) {
	contract, err := bindQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quoter{QuoterCaller: QuoterCaller{contract: contract}, QuoterTransactor: QuoterTransactor{contract: contract}, QuoterFilterer: QuoterFilterer{contract: contract}}, nil
}

// NewQuoterCaller creates a new read-only instance of Quoter, bound to a specific deployed contract.
func NewQuoterCaller(address common.Address, caller bind.ContractCaller) (*QuoterCaller, error) {
	contract, err := bindQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterCaller{contract: contract}, nil
}

// NewQuoterTransactor creates a new write-only instance of Quoter, bound to a specific deployed contract.
func NewQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*QuoterTransactor, error) {
	contract, err := bindQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterTransactor{contract: contract}, nil
}

// NewQuoterFilterer creates a new log filterer instance of Quoter, bound to a specific deployed contract.
func NewQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*QuoterFilterer, error) {
	contract, err := bindQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuoterFilterer{contract: contract}, nil
}

// bindQuoter binds a generic wrapper to an already deployed contract.
func bindQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quoter *QuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quoter.Contract.QuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quoter *QuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.Contract.QuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quoter *QuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quoter.Contract.QuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quoter *QuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quoter *QuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quoter *QuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quoter.Contract.contract.Transact(opts, method, params...)
}

// FullQuotes is a free data retrieval call binding the contract method 0xae2275b3.
//
// Solidity: function fullQuotes(address _optionMarket, uint256 strikeId, uint256 iterations, uint256 amount) view returns(uint256[], uint256[])
func (_Quoter *QuoterCaller) FullQuotes(opts *bind.CallOpts, _optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "fullQuotes", _optionMarket, strikeId, iterations, amount)

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
func (_Quoter *QuoterSession) FullQuotes(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	return _Quoter.Contract.FullQuotes(&_Quoter.CallOpts, _optionMarket, strikeId, iterations, amount)
}

// FullQuotes is a free data retrieval call binding the contract method 0xae2275b3.
//
// Solidity: function fullQuotes(address _optionMarket, uint256 strikeId, uint256 iterations, uint256 amount) view returns(uint256[], uint256[])
func (_Quoter *QuoterCallerSession) FullQuotes(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error) {
	return _Quoter.Contract.FullQuotes(&_Quoter.CallOpts, _optionMarket, strikeId, iterations, amount)
}

// Quote is a free data retrieval call binding the contract method 0x4c438ca7.
//
// Solidity: function quote(address _optionMarket, uint256 strikeId, uint256 iterations, uint8 optionType, uint256 amount) view returns(uint256 totalPremium, uint256 totalFee)
func (_Quoter *QuoterCaller) Quote(opts *bind.CallOpts, _optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "quote", _optionMarket, strikeId, iterations, optionType, amount)

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
func (_Quoter *QuoterSession) Quote(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	return _Quoter.Contract.Quote(&_Quoter.CallOpts, _optionMarket, strikeId, iterations, optionType, amount)
}

// Quote is a free data retrieval call binding the contract method 0x4c438ca7.
//
// Solidity: function quote(address _optionMarket, uint256 strikeId, uint256 iterations, uint8 optionType, uint256 amount) view returns(uint256 totalPremium, uint256 totalFee)
func (_Quoter *QuoterCallerSession) Quote(_optionMarket common.Address, strikeId *big.Int, iterations *big.Int, optionType uint8, amount *big.Int) (struct {
	TotalPremium *big.Int
	TotalFee     *big.Int
}, error) {
	return _Quoter.Contract.Quote(&_Quoter.CallOpts, _optionMarket, strikeId, iterations, optionType, amount)
}
