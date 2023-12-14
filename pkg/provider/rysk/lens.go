// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rysk

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

// DHVLensMK1OptionChain is an auto generated low-level Go binding around an user-defined struct.
type DHVLensMK1OptionChain struct {
	Expirations            []uint64
	OptionExpirationDrills []DHVLensMK1OptionExpirationDrill
}

// DHVLensMK1OptionExpirationDrill is an auto generated low-level Go binding around an user-defined struct.
type DHVLensMK1OptionExpirationDrill struct {
	Expiration      uint64
	CallStrikes     []*big.Int
	CallOptionDrill []DHVLensMK1OptionStrikeDrill
	PutStrikes      []*big.Int
	PutOptionDrill  []DHVLensMK1OptionStrikeDrill
	UnderlyingPrice *big.Int
}

// DHVLensMK1OptionStrikeDrill is an auto generated low-level Go binding around an user-defined struct.
type DHVLensMK1OptionStrikeDrill struct {
	Strike   *big.Int
	Sell     DHVLensMK1TradingSpec
	Buy      DHVLensMK1TradingSpec
	Delta    *big.Int
	Exposure *big.Int
}

// DHVLensMK1TradingSpec is an auto generated low-level Go binding around an user-defined struct.
type DHVLensMK1TradingSpec struct {
	Iv              *big.Int
	Quote           *big.Int
	Fee             *big.Int
	Disabled        bool
	PremiumTooSmall bool
}

// RyskMetaData contains all meta data concerning the Rysk contract.
var RyskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_protocol\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_catalogue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pricer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strikeAsset\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"catalogue\",\"outputs\":[{\"internalType\":\"contractOptionCatalogue\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpirations\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptionChain\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"expirations\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expiration\",\"type\":\"uint64\"},{\"internalType\":\"uint128[]\",\"name\":\"callStrikes\",\"type\":\"uint128[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"strike\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"premiumTooSmall\",\"type\":\"bool\"}],\"internalType\":\"structDHVLensMK1.TradingSpec\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"premiumTooSmall\",\"type\":\"bool\"}],\"internalType\":\"structDHVLensMK1.TradingSpec\",\"name\":\"buy\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"exposure\",\"type\":\"int256\"}],\"internalType\":\"structDHVLensMK1.OptionStrikeDrill[]\",\"name\":\"callOptionDrill\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128[]\",\"name\":\"putStrikes\",\"type\":\"uint128[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"strike\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"premiumTooSmall\",\"type\":\"bool\"}],\"internalType\":\"structDHVLensMK1.TradingSpec\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"premiumTooSmall\",\"type\":\"bool\"}],\"internalType\":\"structDHVLensMK1.TradingSpec\",\"name\":\"buy\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"exposure\",\"type\":\"int256\"}],\"internalType\":\"structDHVLensMK1.OptionStrikeDrill[]\",\"name\":\"putOptionDrill\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"}],\"internalType\":\"structDHVLensMK1.OptionExpirationDrill[]\",\"name\":\"optionExpirationDrills\",\"type\":\"tuple[]\"}],\"internalType\":\"structDHVLensMK1.OptionChain\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expiration\",\"type\":\"uint64\"}],\"name\":\"getOptionExpirationDrill\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expiration\",\"type\":\"uint64\"},{\"internalType\":\"uint128[]\",\"name\":\"callStrikes\",\"type\":\"uint128[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"strike\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"premiumTooSmall\",\"type\":\"bool\"}],\"internalType\":\"structDHVLensMK1.TradingSpec\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"premiumTooSmall\",\"type\":\"bool\"}],\"internalType\":\"structDHVLensMK1.TradingSpec\",\"name\":\"buy\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"exposure\",\"type\":\"int256\"}],\"internalType\":\"structDHVLensMK1.OptionStrikeDrill[]\",\"name\":\"callOptionDrill\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128[]\",\"name\":\"putStrikes\",\"type\":\"uint128[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"strike\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"premiumTooSmall\",\"type\":\"bool\"}],\"internalType\":\"structDHVLensMK1.TradingSpec\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"premiumTooSmall\",\"type\":\"bool\"}],\"internalType\":\"structDHVLensMK1.TradingSpec\",\"name\":\"buy\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"exposure\",\"type\":\"int256\"}],\"internalType\":\"structDHVLensMK1.OptionStrikeDrill[]\",\"name\":\"putOptionDrill\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"}],\"internalType\":\"structDHVLensMK1.OptionExpirationDrill\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricer\",\"outputs\":[{\"internalType\":\"contractBeyondPricer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocol\",\"outputs\":[{\"internalType\":\"contractProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strikeAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RyskABI is the input ABI used to generate the binding from.
// Deprecated: Use RyskMetaData.ABI instead.
var RyskABI = RyskMetaData.ABI

// Rysk is an auto generated Go binding around an Ethereum contract.
type Rysk struct {
	RyskCaller     // Read-only binding to the contract
	RyskTransactor // Write-only binding to the contract
	RyskFilterer   // Log filterer for contract events
}

// RyskCaller is an auto generated read-only Go binding around an Ethereum contract.
type RyskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RyskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RyskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RyskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RyskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RyskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RyskSession struct {
	Contract     *Rysk             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RyskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RyskCallerSession struct {
	Contract *RyskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RyskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RyskTransactorSession struct {
	Contract     *RyskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RyskRaw is an auto generated low-level Go binding around an Ethereum contract.
type RyskRaw struct {
	Contract *Rysk // Generic contract binding to access the raw methods on
}

// RyskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RyskCallerRaw struct {
	Contract *RyskCaller // Generic read-only contract binding to access the raw methods on
}

// RyskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RyskTransactorRaw struct {
	Contract *RyskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRysk creates a new instance of Rysk, bound to a specific deployed contract.
func NewRysk(address common.Address, backend bind.ContractBackend) (*Rysk, error) {
	contract, err := bindRysk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rysk{RyskCaller: RyskCaller{contract: contract}, RyskTransactor: RyskTransactor{contract: contract}, RyskFilterer: RyskFilterer{contract: contract}}, nil
}

// NewRyskCaller creates a new read-only instance of Rysk, bound to a specific deployed contract.
func NewRyskCaller(address common.Address, caller bind.ContractCaller) (*RyskCaller, error) {
	contract, err := bindRysk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RyskCaller{contract: contract}, nil
}

// NewRyskTransactor creates a new write-only instance of Rysk, bound to a specific deployed contract.
func NewRyskTransactor(address common.Address, transactor bind.ContractTransactor) (*RyskTransactor, error) {
	contract, err := bindRysk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RyskTransactor{contract: contract}, nil
}

// NewRyskFilterer creates a new log filterer instance of Rysk, bound to a specific deployed contract.
func NewRyskFilterer(address common.Address, filterer bind.ContractFilterer) (*RyskFilterer, error) {
	contract, err := bindRysk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RyskFilterer{contract: contract}, nil
}

// bindRysk binds a generic wrapper to an already deployed contract.
func bindRysk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RyskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rysk *RyskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rysk.Contract.RyskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rysk *RyskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rysk.Contract.RyskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rysk *RyskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rysk.Contract.RyskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rysk *RyskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rysk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rysk *RyskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rysk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rysk *RyskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rysk.Contract.contract.Transact(opts, method, params...)
}

// Catalogue is a free data retrieval call binding the contract method 0x24e4013a.
//
// Solidity: function catalogue() view returns(address)
func (_Rysk *RyskCaller) Catalogue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rysk.contract.Call(opts, &out, "catalogue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Catalogue is a free data retrieval call binding the contract method 0x24e4013a.
//
// Solidity: function catalogue() view returns(address)
func (_Rysk *RyskSession) Catalogue() (common.Address, error) {
	return _Rysk.Contract.Catalogue(&_Rysk.CallOpts)
}

// Catalogue is a free data retrieval call binding the contract method 0x24e4013a.
//
// Solidity: function catalogue() view returns(address)
func (_Rysk *RyskCallerSession) Catalogue() (common.Address, error) {
	return _Rysk.Contract.Catalogue(&_Rysk.CallOpts)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_Rysk *RyskCaller) CollateralAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rysk.contract.Call(opts, &out, "collateralAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_Rysk *RyskSession) CollateralAsset() (common.Address, error) {
	return _Rysk.Contract.CollateralAsset(&_Rysk.CallOpts)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_Rysk *RyskCallerSession) CollateralAsset() (common.Address, error) {
	return _Rysk.Contract.CollateralAsset(&_Rysk.CallOpts)
}

// GetExpirations is a free data retrieval call binding the contract method 0x56e9d1e4.
//
// Solidity: function getExpirations() view returns(uint64[])
func (_Rysk *RyskCaller) GetExpirations(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _Rysk.contract.Call(opts, &out, "getExpirations")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetExpirations is a free data retrieval call binding the contract method 0x56e9d1e4.
//
// Solidity: function getExpirations() view returns(uint64[])
func (_Rysk *RyskSession) GetExpirations() ([]uint64, error) {
	return _Rysk.Contract.GetExpirations(&_Rysk.CallOpts)
}

// GetExpirations is a free data retrieval call binding the contract method 0x56e9d1e4.
//
// Solidity: function getExpirations() view returns(uint64[])
func (_Rysk *RyskCallerSession) GetExpirations() ([]uint64, error) {
	return _Rysk.Contract.GetExpirations(&_Rysk.CallOpts)
}

// GetOptionChain is a free data retrieval call binding the contract method 0x756fb8c9.
//
// Solidity: function getOptionChain() view returns((uint64[],(uint64,uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint256)[]))
func (_Rysk *RyskCaller) GetOptionChain(opts *bind.CallOpts) (DHVLensMK1OptionChain, error) {
	var out []interface{}
	err := _Rysk.contract.Call(opts, &out, "getOptionChain")

	if err != nil {
		return *new(DHVLensMK1OptionChain), err
	}

	out0 := *abi.ConvertType(out[0], new(DHVLensMK1OptionChain)).(*DHVLensMK1OptionChain)

	return out0, err

}

// GetOptionChain is a free data retrieval call binding the contract method 0x756fb8c9.
//
// Solidity: function getOptionChain() view returns((uint64[],(uint64,uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint256)[]))
func (_Rysk *RyskSession) GetOptionChain() (DHVLensMK1OptionChain, error) {
	return _Rysk.Contract.GetOptionChain(&_Rysk.CallOpts)
}

// GetOptionChain is a free data retrieval call binding the contract method 0x756fb8c9.
//
// Solidity: function getOptionChain() view returns((uint64[],(uint64,uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint256)[]))
func (_Rysk *RyskCallerSession) GetOptionChain() (DHVLensMK1OptionChain, error) {
	return _Rysk.Contract.GetOptionChain(&_Rysk.CallOpts)
}

// GetOptionExpirationDrill is a free data retrieval call binding the contract method 0xe8581f4b.
//
// Solidity: function getOptionExpirationDrill(uint64 expiration) view returns((uint64,uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint256))
func (_Rysk *RyskCaller) GetOptionExpirationDrill(opts *bind.CallOpts, expiration uint64) (DHVLensMK1OptionExpirationDrill, error) {
	var out []interface{}
	err := _Rysk.contract.Call(opts, &out, "getOptionExpirationDrill", expiration)

	if err != nil {
		return *new(DHVLensMK1OptionExpirationDrill), err
	}

	out0 := *abi.ConvertType(out[0], new(DHVLensMK1OptionExpirationDrill)).(*DHVLensMK1OptionExpirationDrill)

	return out0, err

}

// GetOptionExpirationDrill is a free data retrieval call binding the contract method 0xe8581f4b.
//
// Solidity: function getOptionExpirationDrill(uint64 expiration) view returns((uint64,uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint256))
func (_Rysk *RyskSession) GetOptionExpirationDrill(expiration uint64) (DHVLensMK1OptionExpirationDrill, error) {
	return _Rysk.Contract.GetOptionExpirationDrill(&_Rysk.CallOpts, expiration)
}

// GetOptionExpirationDrill is a free data retrieval call binding the contract method 0xe8581f4b.
//
// Solidity: function getOptionExpirationDrill(uint64 expiration) view returns((uint64,uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint128[],(uint128,(uint256,uint256,uint256,bool,bool),(uint256,uint256,uint256,bool,bool),int256,int256)[],uint256))
func (_Rysk *RyskCallerSession) GetOptionExpirationDrill(expiration uint64) (DHVLensMK1OptionExpirationDrill, error) {
	return _Rysk.Contract.GetOptionExpirationDrill(&_Rysk.CallOpts, expiration)
}

// Pricer is a free data retrieval call binding the contract method 0xa6138ed9.
//
// Solidity: function pricer() view returns(address)
func (_Rysk *RyskCaller) Pricer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rysk.contract.Call(opts, &out, "pricer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pricer is a free data retrieval call binding the contract method 0xa6138ed9.
//
// Solidity: function pricer() view returns(address)
func (_Rysk *RyskSession) Pricer() (common.Address, error) {
	return _Rysk.Contract.Pricer(&_Rysk.CallOpts)
}

// Pricer is a free data retrieval call binding the contract method 0xa6138ed9.
//
// Solidity: function pricer() view returns(address)
func (_Rysk *RyskCallerSession) Pricer() (common.Address, error) {
	return _Rysk.Contract.Pricer(&_Rysk.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(address)
func (_Rysk *RyskCaller) Protocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rysk.contract.Call(opts, &out, "protocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(address)
func (_Rysk *RyskSession) Protocol() (common.Address, error) {
	return _Rysk.Contract.Protocol(&_Rysk.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(address)
func (_Rysk *RyskCallerSession) Protocol() (common.Address, error) {
	return _Rysk.Contract.Protocol(&_Rysk.CallOpts)
}

// StrikeAsset is a free data retrieval call binding the contract method 0x17d69bc8.
//
// Solidity: function strikeAsset() view returns(address)
func (_Rysk *RyskCaller) StrikeAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rysk.contract.Call(opts, &out, "strikeAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrikeAsset is a free data retrieval call binding the contract method 0x17d69bc8.
//
// Solidity: function strikeAsset() view returns(address)
func (_Rysk *RyskSession) StrikeAsset() (common.Address, error) {
	return _Rysk.Contract.StrikeAsset(&_Rysk.CallOpts)
}

// StrikeAsset is a free data retrieval call binding the contract method 0x17d69bc8.
//
// Solidity: function strikeAsset() view returns(address)
func (_Rysk *RyskCallerSession) StrikeAsset() (common.Address, error) {
	return _Rysk.Contract.StrikeAsset(&_Rysk.CallOpts)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_Rysk *RyskCaller) UnderlyingAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rysk.contract.Call(opts, &out, "underlyingAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_Rysk *RyskSession) UnderlyingAsset() (common.Address, error) {
	return _Rysk.Contract.UnderlyingAsset(&_Rysk.CallOpts)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_Rysk *RyskCallerSession) UnderlyingAsset() (common.Address, error) {
	return _Rysk.Contract.UnderlyingAsset(&_Rysk.CallOpts)
}
