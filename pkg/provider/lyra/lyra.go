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

// LyraMetaData contains all meta data concerning the Lyra contract.
var LyraMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"}],\"name\":\"BoardBaseIvSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"}],\"name\":\"BoardCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"BoardFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUserLongProfitQuote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBoardLongCallCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBoardLongPutCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAMMShortCallProfitBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAMMShortPutProfitQuote\",\"type\":\"uint256\"}],\"name\":\"BoardLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"}],\"name\":\"ListingAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"}],\"name\":\"ListingSkewSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"name\":\"PositionClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"name\":\"PositionOpened\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"}],\"name\":\"addListingToBoard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boardToPriceAtExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"closePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"strikes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"skews\",\"type\":\"uint256[]\"}],\"name\":\"createOptionBoard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoardListings\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiveBoards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_liveBoards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILyraGlobals\",\"name\":\"_globals\",\"type\":\"address\"},{\"internalType\":\"contractILiquidityPool\",\"name\":\"_liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractIOptionMarketPricer\",\"name\":\"_optionPricer\",\"type\":\"address\"},{\"internalType\":\"contractIOptionGreekCache\",\"name\":\"_greekCache\",\"type\":\"address\"},{\"internalType\":\"contractIShortCollateral\",\"name\":\"_shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractIOptionToken\",\"name\":\"_optionToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_baseAsset\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_errorMessages\",\"type\":\"string[]\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"liquidateExpiredBoard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listingToBaseReturnedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExpiryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"openPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"optionBoards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"optionListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"}],\"name\":\"setBoardBaseIv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setBoardFrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"}],\"name\":\"setListingSkew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"}],\"name\":\"settleOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LyraABI is the input ABI used to generate the binding from.
// Deprecated: Use LyraMetaData.ABI instead.
var LyraABI = LyraMetaData.ABI

// Lyra is an auto generated Go binding around an Ethereum contract.
type Lyra struct {
	LyraCaller     // Read-only binding to the contract
	LyraTransactor // Write-only binding to the contract
	LyraFilterer   // Log filterer for contract events
}

// LyraCaller is an auto generated read-only Go binding around an Ethereum contract.
type LyraCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LyraTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LyraTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LyraFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LyraFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LyraSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LyraSession struct {
	Contract     *Lyra             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LyraCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LyraCallerSession struct {
	Contract *LyraCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LyraTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LyraTransactorSession struct {
	Contract     *LyraTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LyraRaw is an auto generated low-level Go binding around an Ethereum contract.
type LyraRaw struct {
	Contract *Lyra // Generic contract binding to access the raw methods on
}

// LyraCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LyraCallerRaw struct {
	Contract *LyraCaller // Generic read-only contract binding to access the raw methods on
}

// LyraTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LyraTransactorRaw struct {
	Contract *LyraTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLyra creates a new instance of Lyra, bound to a specific deployed contract.
func NewLyra(address common.Address, backend bind.ContractBackend) (*Lyra, error) {
	contract, err := bindLyra(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lyra{LyraCaller: LyraCaller{contract: contract}, LyraTransactor: LyraTransactor{contract: contract}, LyraFilterer: LyraFilterer{contract: contract}}, nil
}

// NewLyraCaller creates a new read-only instance of Lyra, bound to a specific deployed contract.
func NewLyraCaller(address common.Address, caller bind.ContractCaller) (*LyraCaller, error) {
	contract, err := bindLyra(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LyraCaller{contract: contract}, nil
}

// NewLyraTransactor creates a new write-only instance of Lyra, bound to a specific deployed contract.
func NewLyraTransactor(address common.Address, transactor bind.ContractTransactor) (*LyraTransactor, error) {
	contract, err := bindLyra(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LyraTransactor{contract: contract}, nil
}

// NewLyraFilterer creates a new log filterer instance of Lyra, bound to a specific deployed contract.
func NewLyraFilterer(address common.Address, filterer bind.ContractFilterer) (*LyraFilterer, error) {
	contract, err := bindLyra(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LyraFilterer{contract: contract}, nil
}

// bindLyra binds a generic wrapper to an already deployed contract.
func bindLyra(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LyraABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lyra *LyraRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lyra.Contract.LyraCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lyra *LyraRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lyra.Contract.LyraTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lyra *LyraRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lyra.Contract.LyraTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lyra *LyraCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lyra.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lyra *LyraTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lyra.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lyra *LyraTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lyra.Contract.contract.Transact(opts, method, params...)
}

// BoardToPriceAtExpiry is a free data retrieval call binding the contract method 0xd1e9e811.
//
// Solidity: function boardToPriceAtExpiry(uint256 ) view returns(uint256)
func (_Lyra *LyraCaller) BoardToPriceAtExpiry(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "boardToPriceAtExpiry", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoardToPriceAtExpiry is a free data retrieval call binding the contract method 0xd1e9e811.
//
// Solidity: function boardToPriceAtExpiry(uint256 ) view returns(uint256)
func (_Lyra *LyraSession) BoardToPriceAtExpiry(arg0 *big.Int) (*big.Int, error) {
	return _Lyra.Contract.BoardToPriceAtExpiry(&_Lyra.CallOpts, arg0)
}

// BoardToPriceAtExpiry is a free data retrieval call binding the contract method 0xd1e9e811.
//
// Solidity: function boardToPriceAtExpiry(uint256 ) view returns(uint256)
func (_Lyra *LyraCallerSession) BoardToPriceAtExpiry(arg0 *big.Int) (*big.Int, error) {
	return _Lyra.Contract.BoardToPriceAtExpiry(&_Lyra.CallOpts, arg0)
}

// GetBoardListings is a free data retrieval call binding the contract method 0xc39eeefb.
//
// Solidity: function getBoardListings(uint256 boardId) view returns(uint256[])
func (_Lyra *LyraCaller) GetBoardListings(opts *bind.CallOpts, boardId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getBoardListings", boardId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBoardListings is a free data retrieval call binding the contract method 0xc39eeefb.
//
// Solidity: function getBoardListings(uint256 boardId) view returns(uint256[])
func (_Lyra *LyraSession) GetBoardListings(boardId *big.Int) ([]*big.Int, error) {
	return _Lyra.Contract.GetBoardListings(&_Lyra.CallOpts, boardId)
}

// GetBoardListings is a free data retrieval call binding the contract method 0xc39eeefb.
//
// Solidity: function getBoardListings(uint256 boardId) view returns(uint256[])
func (_Lyra *LyraCallerSession) GetBoardListings(boardId *big.Int) ([]*big.Int, error) {
	return _Lyra.Contract.GetBoardListings(&_Lyra.CallOpts, boardId)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0x7e7088eb.
//
// Solidity: function getLiveBoards() view returns(uint256[] _liveBoards)
func (_Lyra *LyraCaller) GetLiveBoards(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getLiveBoards")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLiveBoards is a free data retrieval call binding the contract method 0x7e7088eb.
//
// Solidity: function getLiveBoards() view returns(uint256[] _liveBoards)
func (_Lyra *LyraSession) GetLiveBoards() ([]*big.Int, error) {
	return _Lyra.Contract.GetLiveBoards(&_Lyra.CallOpts)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0x7e7088eb.
//
// Solidity: function getLiveBoards() view returns(uint256[] _liveBoards)
func (_Lyra *LyraCallerSession) GetLiveBoards() ([]*big.Int, error) {
	return _Lyra.Contract.GetLiveBoards(&_Lyra.CallOpts)
}

// ListingToBaseReturnedRatio is a free data retrieval call binding the contract method 0x9b805da8.
//
// Solidity: function listingToBaseReturnedRatio(uint256 ) view returns(uint256)
func (_Lyra *LyraCaller) ListingToBaseReturnedRatio(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "listingToBaseReturnedRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListingToBaseReturnedRatio is a free data retrieval call binding the contract method 0x9b805da8.
//
// Solidity: function listingToBaseReturnedRatio(uint256 ) view returns(uint256)
func (_Lyra *LyraSession) ListingToBaseReturnedRatio(arg0 *big.Int) (*big.Int, error) {
	return _Lyra.Contract.ListingToBaseReturnedRatio(&_Lyra.CallOpts, arg0)
}

// ListingToBaseReturnedRatio is a free data retrieval call binding the contract method 0x9b805da8.
//
// Solidity: function listingToBaseReturnedRatio(uint256 ) view returns(uint256)
func (_Lyra *LyraCallerSession) ListingToBaseReturnedRatio(arg0 *big.Int) (*big.Int, error) {
	return _Lyra.Contract.ListingToBaseReturnedRatio(&_Lyra.CallOpts, arg0)
}

// MaxExpiryTimestamp is a free data retrieval call binding the contract method 0x2eb6534f.
//
// Solidity: function maxExpiryTimestamp() view returns(uint256)
func (_Lyra *LyraCaller) MaxExpiryTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "maxExpiryTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExpiryTimestamp is a free data retrieval call binding the contract method 0x2eb6534f.
//
// Solidity: function maxExpiryTimestamp() view returns(uint256)
func (_Lyra *LyraSession) MaxExpiryTimestamp() (*big.Int, error) {
	return _Lyra.Contract.MaxExpiryTimestamp(&_Lyra.CallOpts)
}

// MaxExpiryTimestamp is a free data retrieval call binding the contract method 0x2eb6534f.
//
// Solidity: function maxExpiryTimestamp() view returns(uint256)
func (_Lyra *LyraCallerSession) MaxExpiryTimestamp() (*big.Int, error) {
	return _Lyra.Contract.MaxExpiryTimestamp(&_Lyra.CallOpts)
}

// OptionBoards is a free data retrieval call binding the contract method 0x35359675.
//
// Solidity: function optionBoards(uint256 ) view returns(uint256 id, uint256 expiry, uint256 iv, bool frozen)
func (_Lyra *LyraCaller) OptionBoards(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Expiry *big.Int
	Iv     *big.Int
	Frozen bool
}, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "optionBoards", arg0)

	outstruct := new(struct {
		Id     *big.Int
		Expiry *big.Int
		Iv     *big.Int
		Frozen bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Iv = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Frozen = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// OptionBoards is a free data retrieval call binding the contract method 0x35359675.
//
// Solidity: function optionBoards(uint256 ) view returns(uint256 id, uint256 expiry, uint256 iv, bool frozen)
func (_Lyra *LyraSession) OptionBoards(arg0 *big.Int) (struct {
	Id     *big.Int
	Expiry *big.Int
	Iv     *big.Int
	Frozen bool
}, error) {
	return _Lyra.Contract.OptionBoards(&_Lyra.CallOpts, arg0)
}

// OptionBoards is a free data retrieval call binding the contract method 0x35359675.
//
// Solidity: function optionBoards(uint256 ) view returns(uint256 id, uint256 expiry, uint256 iv, bool frozen)
func (_Lyra *LyraCallerSession) OptionBoards(arg0 *big.Int) (struct {
	Id     *big.Int
	Expiry *big.Int
	Iv     *big.Int
	Frozen bool
}, error) {
	return _Lyra.Contract.OptionBoards(&_Lyra.CallOpts, arg0)
}

// OptionListings is a free data retrieval call binding the contract method 0x3289cf8f.
//
// Solidity: function optionListings(uint256 ) view returns(uint256 id, uint256 strike, uint256 skew, uint256 longCall, uint256 shortCall, uint256 longPut, uint256 shortPut, uint256 boardId)
func (_Lyra *LyraCaller) OptionListings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Strike    *big.Int
	Skew      *big.Int
	LongCall  *big.Int
	ShortCall *big.Int
	LongPut   *big.Int
	ShortPut  *big.Int
	BoardId   *big.Int
}, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "optionListings", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Strike    *big.Int
		Skew      *big.Int
		LongCall  *big.Int
		ShortCall *big.Int
		LongPut   *big.Int
		ShortPut  *big.Int
		BoardId   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Strike = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Skew = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LongCall = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ShortCall = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LongPut = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ShortPut = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.BoardId = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OptionListings is a free data retrieval call binding the contract method 0x3289cf8f.
//
// Solidity: function optionListings(uint256 ) view returns(uint256 id, uint256 strike, uint256 skew, uint256 longCall, uint256 shortCall, uint256 longPut, uint256 shortPut, uint256 boardId)
func (_Lyra *LyraSession) OptionListings(arg0 *big.Int) (struct {
	Id        *big.Int
	Strike    *big.Int
	Skew      *big.Int
	LongCall  *big.Int
	ShortCall *big.Int
	LongPut   *big.Int
	ShortPut  *big.Int
	BoardId   *big.Int
}, error) {
	return _Lyra.Contract.OptionListings(&_Lyra.CallOpts, arg0)
}

// OptionListings is a free data retrieval call binding the contract method 0x3289cf8f.
//
// Solidity: function optionListings(uint256 ) view returns(uint256 id, uint256 strike, uint256 skew, uint256 longCall, uint256 shortCall, uint256 longPut, uint256 shortPut, uint256 boardId)
func (_Lyra *LyraCallerSession) OptionListings(arg0 *big.Int) (struct {
	Id        *big.Int
	Strike    *big.Int
	Skew      *big.Int
	LongCall  *big.Int
	ShortCall *big.Int
	LongPut   *big.Int
	ShortPut  *big.Int
	BoardId   *big.Int
}, error) {
	return _Lyra.Contract.OptionListings(&_Lyra.CallOpts, arg0)
}

// AddListingToBoard is a paid mutator transaction binding the contract method 0x2f4aff0e.
//
// Solidity: function addListingToBoard(uint256 boardId, uint256 strike, uint256 skew) returns()
func (_Lyra *LyraTransactor) AddListingToBoard(opts *bind.TransactOpts, boardId *big.Int, strike *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "addListingToBoard", boardId, strike, skew)
}

// AddListingToBoard is a paid mutator transaction binding the contract method 0x2f4aff0e.
//
// Solidity: function addListingToBoard(uint256 boardId, uint256 strike, uint256 skew) returns()
func (_Lyra *LyraSession) AddListingToBoard(boardId *big.Int, strike *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.AddListingToBoard(&_Lyra.TransactOpts, boardId, strike, skew)
}

// AddListingToBoard is a paid mutator transaction binding the contract method 0x2f4aff0e.
//
// Solidity: function addListingToBoard(uint256 boardId, uint256 strike, uint256 skew) returns()
func (_Lyra *LyraTransactorSession) AddListingToBoard(boardId *big.Int, strike *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.AddListingToBoard(&_Lyra.TransactOpts, boardId, strike, skew)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x97dd2524.
//
// Solidity: function closePosition(uint256 _listingId, uint8 tradeType, uint256 amount) returns(uint256 totalCost)
func (_Lyra *LyraTransactor) ClosePosition(opts *bind.TransactOpts, _listingId *big.Int, tradeType uint8, amount *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "closePosition", _listingId, tradeType, amount)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x97dd2524.
//
// Solidity: function closePosition(uint256 _listingId, uint8 tradeType, uint256 amount) returns(uint256 totalCost)
func (_Lyra *LyraSession) ClosePosition(_listingId *big.Int, tradeType uint8, amount *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.ClosePosition(&_Lyra.TransactOpts, _listingId, tradeType, amount)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x97dd2524.
//
// Solidity: function closePosition(uint256 _listingId, uint8 tradeType, uint256 amount) returns(uint256 totalCost)
func (_Lyra *LyraTransactorSession) ClosePosition(_listingId *big.Int, tradeType uint8, amount *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.ClosePosition(&_Lyra.TransactOpts, _listingId, tradeType, amount)
}

// CreateOptionBoard is a paid mutator transaction binding the contract method 0xdd2f0d36.
//
// Solidity: function createOptionBoard(uint256 expiry, uint256 baseIV, uint256[] strikes, uint256[] skews) returns(uint256)
func (_Lyra *LyraTransactor) CreateOptionBoard(opts *bind.TransactOpts, expiry *big.Int, baseIV *big.Int, strikes []*big.Int, skews []*big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "createOptionBoard", expiry, baseIV, strikes, skews)
}

// CreateOptionBoard is a paid mutator transaction binding the contract method 0xdd2f0d36.
//
// Solidity: function createOptionBoard(uint256 expiry, uint256 baseIV, uint256[] strikes, uint256[] skews) returns(uint256)
func (_Lyra *LyraSession) CreateOptionBoard(expiry *big.Int, baseIV *big.Int, strikes []*big.Int, skews []*big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.CreateOptionBoard(&_Lyra.TransactOpts, expiry, baseIV, strikes, skews)
}

// CreateOptionBoard is a paid mutator transaction binding the contract method 0xdd2f0d36.
//
// Solidity: function createOptionBoard(uint256 expiry, uint256 baseIV, uint256[] strikes, uint256[] skews) returns(uint256)
func (_Lyra *LyraTransactorSession) CreateOptionBoard(expiry *big.Int, baseIV *big.Int, strikes []*big.Int, skews []*big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.CreateOptionBoard(&_Lyra.TransactOpts, expiry, baseIV, strikes, skews)
}

// Init is a paid mutator transaction binding the contract method 0xf2333a58.
//
// Solidity: function init(address _globals, address _liquidityPool, address _optionPricer, address _greekCache, address _shortCollateral, address _optionToken, address _quoteAsset, address _baseAsset, string[] _errorMessages) returns()
func (_Lyra *LyraTransactor) Init(opts *bind.TransactOpts, _globals common.Address, _liquidityPool common.Address, _optionPricer common.Address, _greekCache common.Address, _shortCollateral common.Address, _optionToken common.Address, _quoteAsset common.Address, _baseAsset common.Address, _errorMessages []string) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "init", _globals, _liquidityPool, _optionPricer, _greekCache, _shortCollateral, _optionToken, _quoteAsset, _baseAsset, _errorMessages)
}

// Init is a paid mutator transaction binding the contract method 0xf2333a58.
//
// Solidity: function init(address _globals, address _liquidityPool, address _optionPricer, address _greekCache, address _shortCollateral, address _optionToken, address _quoteAsset, address _baseAsset, string[] _errorMessages) returns()
func (_Lyra *LyraSession) Init(_globals common.Address, _liquidityPool common.Address, _optionPricer common.Address, _greekCache common.Address, _shortCollateral common.Address, _optionToken common.Address, _quoteAsset common.Address, _baseAsset common.Address, _errorMessages []string) (*types.Transaction, error) {
	return _Lyra.Contract.Init(&_Lyra.TransactOpts, _globals, _liquidityPool, _optionPricer, _greekCache, _shortCollateral, _optionToken, _quoteAsset, _baseAsset, _errorMessages)
}

// Init is a paid mutator transaction binding the contract method 0xf2333a58.
//
// Solidity: function init(address _globals, address _liquidityPool, address _optionPricer, address _greekCache, address _shortCollateral, address _optionToken, address _quoteAsset, address _baseAsset, string[] _errorMessages) returns()
func (_Lyra *LyraTransactorSession) Init(_globals common.Address, _liquidityPool common.Address, _optionPricer common.Address, _greekCache common.Address, _shortCollateral common.Address, _optionToken common.Address, _quoteAsset common.Address, _baseAsset common.Address, _errorMessages []string) (*types.Transaction, error) {
	return _Lyra.Contract.Init(&_Lyra.TransactOpts, _globals, _liquidityPool, _optionPricer, _greekCache, _shortCollateral, _optionToken, _quoteAsset, _baseAsset, _errorMessages)
}

// LiquidateExpiredBoard is a paid mutator transaction binding the contract method 0x42a47e13.
//
// Solidity: function liquidateExpiredBoard(uint256 boardId) returns()
func (_Lyra *LyraTransactor) LiquidateExpiredBoard(opts *bind.TransactOpts, boardId *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "liquidateExpiredBoard", boardId)
}

// LiquidateExpiredBoard is a paid mutator transaction binding the contract method 0x42a47e13.
//
// Solidity: function liquidateExpiredBoard(uint256 boardId) returns()
func (_Lyra *LyraSession) LiquidateExpiredBoard(boardId *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.LiquidateExpiredBoard(&_Lyra.TransactOpts, boardId)
}

// LiquidateExpiredBoard is a paid mutator transaction binding the contract method 0x42a47e13.
//
// Solidity: function liquidateExpiredBoard(uint256 boardId) returns()
func (_Lyra *LyraTransactorSession) LiquidateExpiredBoard(boardId *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.LiquidateExpiredBoard(&_Lyra.TransactOpts, boardId)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x862a3394.
//
// Solidity: function openPosition(uint256 _listingId, uint8 tradeType, uint256 amount) returns(uint256 totalCost)
func (_Lyra *LyraTransactor) OpenPosition(opts *bind.TransactOpts, _listingId *big.Int, tradeType uint8, amount *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "openPosition", _listingId, tradeType, amount)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x862a3394.
//
// Solidity: function openPosition(uint256 _listingId, uint8 tradeType, uint256 amount) returns(uint256 totalCost)
func (_Lyra *LyraSession) OpenPosition(_listingId *big.Int, tradeType uint8, amount *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.OpenPosition(&_Lyra.TransactOpts, _listingId, tradeType, amount)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x862a3394.
//
// Solidity: function openPosition(uint256 _listingId, uint8 tradeType, uint256 amount) returns(uint256 totalCost)
func (_Lyra *LyraTransactorSession) OpenPosition(_listingId *big.Int, tradeType uint8, amount *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.OpenPosition(&_Lyra.TransactOpts, _listingId, tradeType, amount)
}

// SetBoardBaseIv is a paid mutator transaction binding the contract method 0x18cc7e86.
//
// Solidity: function setBoardBaseIv(uint256 boardId, uint256 baseIv) returns()
func (_Lyra *LyraTransactor) SetBoardBaseIv(opts *bind.TransactOpts, boardId *big.Int, baseIv *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "setBoardBaseIv", boardId, baseIv)
}

// SetBoardBaseIv is a paid mutator transaction binding the contract method 0x18cc7e86.
//
// Solidity: function setBoardBaseIv(uint256 boardId, uint256 baseIv) returns()
func (_Lyra *LyraSession) SetBoardBaseIv(boardId *big.Int, baseIv *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.SetBoardBaseIv(&_Lyra.TransactOpts, boardId, baseIv)
}

// SetBoardBaseIv is a paid mutator transaction binding the contract method 0x18cc7e86.
//
// Solidity: function setBoardBaseIv(uint256 boardId, uint256 baseIv) returns()
func (_Lyra *LyraTransactorSession) SetBoardBaseIv(boardId *big.Int, baseIv *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.SetBoardBaseIv(&_Lyra.TransactOpts, boardId, baseIv)
}

// SetBoardFrozen is a paid mutator transaction binding the contract method 0xa9c9d125.
//
// Solidity: function setBoardFrozen(uint256 boardId, bool frozen) returns()
func (_Lyra *LyraTransactor) SetBoardFrozen(opts *bind.TransactOpts, boardId *big.Int, frozen bool) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "setBoardFrozen", boardId, frozen)
}

// SetBoardFrozen is a paid mutator transaction binding the contract method 0xa9c9d125.
//
// Solidity: function setBoardFrozen(uint256 boardId, bool frozen) returns()
func (_Lyra *LyraSession) SetBoardFrozen(boardId *big.Int, frozen bool) (*types.Transaction, error) {
	return _Lyra.Contract.SetBoardFrozen(&_Lyra.TransactOpts, boardId, frozen)
}

// SetBoardFrozen is a paid mutator transaction binding the contract method 0xa9c9d125.
//
// Solidity: function setBoardFrozen(uint256 boardId, bool frozen) returns()
func (_Lyra *LyraTransactorSession) SetBoardFrozen(boardId *big.Int, frozen bool) (*types.Transaction, error) {
	return _Lyra.Contract.SetBoardFrozen(&_Lyra.TransactOpts, boardId, frozen)
}

// SetListingSkew is a paid mutator transaction binding the contract method 0xd6ef9be3.
//
// Solidity: function setListingSkew(uint256 listingId, uint256 skew) returns()
func (_Lyra *LyraTransactor) SetListingSkew(opts *bind.TransactOpts, listingId *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "setListingSkew", listingId, skew)
}

// SetListingSkew is a paid mutator transaction binding the contract method 0xd6ef9be3.
//
// Solidity: function setListingSkew(uint256 listingId, uint256 skew) returns()
func (_Lyra *LyraSession) SetListingSkew(listingId *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.SetListingSkew(&_Lyra.TransactOpts, listingId, skew)
}

// SetListingSkew is a paid mutator transaction binding the contract method 0xd6ef9be3.
//
// Solidity: function setListingSkew(uint256 listingId, uint256 skew) returns()
func (_Lyra *LyraTransactorSession) SetListingSkew(listingId *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.SetListingSkew(&_Lyra.TransactOpts, listingId, skew)
}

// SettleOptions is a paid mutator transaction binding the contract method 0xa92e61eb.
//
// Solidity: function settleOptions(uint256 listingId, uint8 tradeType) returns()
func (_Lyra *LyraTransactor) SettleOptions(opts *bind.TransactOpts, listingId *big.Int, tradeType uint8) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "settleOptions", listingId, tradeType)
}

// SettleOptions is a paid mutator transaction binding the contract method 0xa92e61eb.
//
// Solidity: function settleOptions(uint256 listingId, uint8 tradeType) returns()
func (_Lyra *LyraSession) SettleOptions(listingId *big.Int, tradeType uint8) (*types.Transaction, error) {
	return _Lyra.Contract.SettleOptions(&_Lyra.TransactOpts, listingId, tradeType)
}

// SettleOptions is a paid mutator transaction binding the contract method 0xa92e61eb.
//
// Solidity: function settleOptions(uint256 listingId, uint8 tradeType) returns()
func (_Lyra *LyraTransactorSession) SettleOptions(listingId *big.Int, tradeType uint8) (*types.Transaction, error) {
	return _Lyra.Contract.SettleOptions(&_Lyra.TransactOpts, listingId, tradeType)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lyra *LyraTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lyra *LyraSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lyra.Contract.TransferOwnership(&_Lyra.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lyra *LyraTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lyra.Contract.TransferOwnership(&_Lyra.TransactOpts, newOwner)
}

// LyraBoardBaseIvSetIterator is returned from FilterBoardBaseIvSet and is used to iterate over the raw logs and unpacked data for BoardBaseIvSet events raised by the Lyra contract.
type LyraBoardBaseIvSetIterator struct {
	Event *LyraBoardBaseIvSet // Event containing the contract specifics and raw log

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
func (it *LyraBoardBaseIvSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraBoardBaseIvSet)
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
		it.Event = new(LyraBoardBaseIvSet)
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
func (it *LyraBoardBaseIvSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraBoardBaseIvSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraBoardBaseIvSet represents a BoardBaseIvSet event raised by the Lyra contract.
type LyraBoardBaseIvSet struct {
	BoardId *big.Int
	BaseIv  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBoardBaseIvSet is a free log retrieval operation binding the contract event 0x27dc10bc12529bac536af6dbf5d4b270673ac7aeb848c334e4038ef55ecce881.
//
// Solidity: event BoardBaseIvSet(uint256 indexed boardId, uint256 baseIv)
func (_Lyra *LyraFilterer) FilterBoardBaseIvSet(opts *bind.FilterOpts, boardId []*big.Int) (*LyraBoardBaseIvSetIterator, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "BoardBaseIvSet", boardIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraBoardBaseIvSetIterator{contract: _Lyra.contract, event: "BoardBaseIvSet", logs: logs, sub: sub}, nil
}

// WatchBoardBaseIvSet is a free log subscription operation binding the contract event 0x27dc10bc12529bac536af6dbf5d4b270673ac7aeb848c334e4038ef55ecce881.
//
// Solidity: event BoardBaseIvSet(uint256 indexed boardId, uint256 baseIv)
func (_Lyra *LyraFilterer) WatchBoardBaseIvSet(opts *bind.WatchOpts, sink chan<- *LyraBoardBaseIvSet, boardId []*big.Int) (event.Subscription, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "BoardBaseIvSet", boardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraBoardBaseIvSet)
				if err := _Lyra.contract.UnpackLog(event, "BoardBaseIvSet", log); err != nil {
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

// ParseBoardBaseIvSet is a log parse operation binding the contract event 0x27dc10bc12529bac536af6dbf5d4b270673ac7aeb848c334e4038ef55ecce881.
//
// Solidity: event BoardBaseIvSet(uint256 indexed boardId, uint256 baseIv)
func (_Lyra *LyraFilterer) ParseBoardBaseIvSet(log types.Log) (*LyraBoardBaseIvSet, error) {
	event := new(LyraBoardBaseIvSet)
	if err := _Lyra.contract.UnpackLog(event, "BoardBaseIvSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraBoardCreatedIterator is returned from FilterBoardCreated and is used to iterate over the raw logs and unpacked data for BoardCreated events raised by the Lyra contract.
type LyraBoardCreatedIterator struct {
	Event *LyraBoardCreated // Event containing the contract specifics and raw log

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
func (it *LyraBoardCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraBoardCreated)
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
		it.Event = new(LyraBoardCreated)
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
func (it *LyraBoardCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraBoardCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraBoardCreated represents a BoardCreated event raised by the Lyra contract.
type LyraBoardCreated struct {
	BoardId *big.Int
	Expiry  *big.Int
	BaseIv  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBoardCreated is a free log retrieval operation binding the contract event 0x2eb8cef7f145b51f7d30c1ffc842e46161a2b1624415ce69192c99ec2106ad87.
//
// Solidity: event BoardCreated(uint256 indexed boardId, uint256 expiry, uint256 baseIv)
func (_Lyra *LyraFilterer) FilterBoardCreated(opts *bind.FilterOpts, boardId []*big.Int) (*LyraBoardCreatedIterator, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "BoardCreated", boardIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraBoardCreatedIterator{contract: _Lyra.contract, event: "BoardCreated", logs: logs, sub: sub}, nil
}

// WatchBoardCreated is a free log subscription operation binding the contract event 0x2eb8cef7f145b51f7d30c1ffc842e46161a2b1624415ce69192c99ec2106ad87.
//
// Solidity: event BoardCreated(uint256 indexed boardId, uint256 expiry, uint256 baseIv)
func (_Lyra *LyraFilterer) WatchBoardCreated(opts *bind.WatchOpts, sink chan<- *LyraBoardCreated, boardId []*big.Int) (event.Subscription, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "BoardCreated", boardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraBoardCreated)
				if err := _Lyra.contract.UnpackLog(event, "BoardCreated", log); err != nil {
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

// ParseBoardCreated is a log parse operation binding the contract event 0x2eb8cef7f145b51f7d30c1ffc842e46161a2b1624415ce69192c99ec2106ad87.
//
// Solidity: event BoardCreated(uint256 indexed boardId, uint256 expiry, uint256 baseIv)
func (_Lyra *LyraFilterer) ParseBoardCreated(log types.Log) (*LyraBoardCreated, error) {
	event := new(LyraBoardCreated)
	if err := _Lyra.contract.UnpackLog(event, "BoardCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraBoardFrozenIterator is returned from FilterBoardFrozen and is used to iterate over the raw logs and unpacked data for BoardFrozen events raised by the Lyra contract.
type LyraBoardFrozenIterator struct {
	Event *LyraBoardFrozen // Event containing the contract specifics and raw log

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
func (it *LyraBoardFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraBoardFrozen)
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
		it.Event = new(LyraBoardFrozen)
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
func (it *LyraBoardFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraBoardFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraBoardFrozen represents a BoardFrozen event raised by the Lyra contract.
type LyraBoardFrozen struct {
	BoardId *big.Int
	Frozen  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBoardFrozen is a free log retrieval operation binding the contract event 0xab7e756517bb425436c10403644a884802e0b2d5105f9f5386823b7c42ca5d5f.
//
// Solidity: event BoardFrozen(uint256 indexed boardId, bool frozen)
func (_Lyra *LyraFilterer) FilterBoardFrozen(opts *bind.FilterOpts, boardId []*big.Int) (*LyraBoardFrozenIterator, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "BoardFrozen", boardIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraBoardFrozenIterator{contract: _Lyra.contract, event: "BoardFrozen", logs: logs, sub: sub}, nil
}

// WatchBoardFrozen is a free log subscription operation binding the contract event 0xab7e756517bb425436c10403644a884802e0b2d5105f9f5386823b7c42ca5d5f.
//
// Solidity: event BoardFrozen(uint256 indexed boardId, bool frozen)
func (_Lyra *LyraFilterer) WatchBoardFrozen(opts *bind.WatchOpts, sink chan<- *LyraBoardFrozen, boardId []*big.Int) (event.Subscription, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "BoardFrozen", boardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraBoardFrozen)
				if err := _Lyra.contract.UnpackLog(event, "BoardFrozen", log); err != nil {
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

// ParseBoardFrozen is a log parse operation binding the contract event 0xab7e756517bb425436c10403644a884802e0b2d5105f9f5386823b7c42ca5d5f.
//
// Solidity: event BoardFrozen(uint256 indexed boardId, bool frozen)
func (_Lyra *LyraFilterer) ParseBoardFrozen(log types.Log) (*LyraBoardFrozen, error) {
	event := new(LyraBoardFrozen)
	if err := _Lyra.contract.UnpackLog(event, "BoardFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraBoardLiquidatedIterator is returned from FilterBoardLiquidated and is used to iterate over the raw logs and unpacked data for BoardLiquidated events raised by the Lyra contract.
type LyraBoardLiquidatedIterator struct {
	Event *LyraBoardLiquidated // Event containing the contract specifics and raw log

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
func (it *LyraBoardLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraBoardLiquidated)
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
		it.Event = new(LyraBoardLiquidated)
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
func (it *LyraBoardLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraBoardLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraBoardLiquidated represents a BoardLiquidated event raised by the Lyra contract.
type LyraBoardLiquidated struct {
	BoardId                      *big.Int
	TotalUserLongProfitQuote     *big.Int
	TotalBoardLongCallCollateral *big.Int
	TotalBoardLongPutCollateral  *big.Int
	TotalAMMShortCallProfitBase  *big.Int
	TotalAMMShortPutProfitQuote  *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterBoardLiquidated is a free log retrieval operation binding the contract event 0x3f24fabf249b0bbc2acff84ed6a1b001961a11a3d32fc78ec28407708d3c238b.
//
// Solidity: event BoardLiquidated(uint256 indexed boardId, uint256 totalUserLongProfitQuote, uint256 totalBoardLongCallCollateral, uint256 totalBoardLongPutCollateral, uint256 totalAMMShortCallProfitBase, uint256 totalAMMShortPutProfitQuote)
func (_Lyra *LyraFilterer) FilterBoardLiquidated(opts *bind.FilterOpts, boardId []*big.Int) (*LyraBoardLiquidatedIterator, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "BoardLiquidated", boardIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraBoardLiquidatedIterator{contract: _Lyra.contract, event: "BoardLiquidated", logs: logs, sub: sub}, nil
}

// WatchBoardLiquidated is a free log subscription operation binding the contract event 0x3f24fabf249b0bbc2acff84ed6a1b001961a11a3d32fc78ec28407708d3c238b.
//
// Solidity: event BoardLiquidated(uint256 indexed boardId, uint256 totalUserLongProfitQuote, uint256 totalBoardLongCallCollateral, uint256 totalBoardLongPutCollateral, uint256 totalAMMShortCallProfitBase, uint256 totalAMMShortPutProfitQuote)
func (_Lyra *LyraFilterer) WatchBoardLiquidated(opts *bind.WatchOpts, sink chan<- *LyraBoardLiquidated, boardId []*big.Int) (event.Subscription, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "BoardLiquidated", boardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraBoardLiquidated)
				if err := _Lyra.contract.UnpackLog(event, "BoardLiquidated", log); err != nil {
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

// ParseBoardLiquidated is a log parse operation binding the contract event 0x3f24fabf249b0bbc2acff84ed6a1b001961a11a3d32fc78ec28407708d3c238b.
//
// Solidity: event BoardLiquidated(uint256 indexed boardId, uint256 totalUserLongProfitQuote, uint256 totalBoardLongCallCollateral, uint256 totalBoardLongPutCollateral, uint256 totalAMMShortCallProfitBase, uint256 totalAMMShortPutProfitQuote)
func (_Lyra *LyraFilterer) ParseBoardLiquidated(log types.Log) (*LyraBoardLiquidated, error) {
	event := new(LyraBoardLiquidated)
	if err := _Lyra.contract.UnpackLog(event, "BoardLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraListingAddedIterator is returned from FilterListingAdded and is used to iterate over the raw logs and unpacked data for ListingAdded events raised by the Lyra contract.
type LyraListingAddedIterator struct {
	Event *LyraListingAdded // Event containing the contract specifics and raw log

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
func (it *LyraListingAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraListingAdded)
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
		it.Event = new(LyraListingAdded)
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
func (it *LyraListingAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraListingAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraListingAdded represents a ListingAdded event raised by the Lyra contract.
type LyraListingAdded struct {
	BoardId   *big.Int
	ListingId *big.Int
	Strike    *big.Int
	Skew      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingAdded is a free log retrieval operation binding the contract event 0xb022fbd99a76d7d6d20bd5b22c69b0e8bd1b756764c76b117306d29bbc16e09d.
//
// Solidity: event ListingAdded(uint256 indexed boardId, uint256 indexed listingId, uint256 strike, uint256 skew)
func (_Lyra *LyraFilterer) FilterListingAdded(opts *bind.FilterOpts, boardId []*big.Int, listingId []*big.Int) (*LyraListingAddedIterator, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "ListingAdded", boardIdRule, listingIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraListingAddedIterator{contract: _Lyra.contract, event: "ListingAdded", logs: logs, sub: sub}, nil
}

// WatchListingAdded is a free log subscription operation binding the contract event 0xb022fbd99a76d7d6d20bd5b22c69b0e8bd1b756764c76b117306d29bbc16e09d.
//
// Solidity: event ListingAdded(uint256 indexed boardId, uint256 indexed listingId, uint256 strike, uint256 skew)
func (_Lyra *LyraFilterer) WatchListingAdded(opts *bind.WatchOpts, sink chan<- *LyraListingAdded, boardId []*big.Int, listingId []*big.Int) (event.Subscription, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "ListingAdded", boardIdRule, listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraListingAdded)
				if err := _Lyra.contract.UnpackLog(event, "ListingAdded", log); err != nil {
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

// ParseListingAdded is a log parse operation binding the contract event 0xb022fbd99a76d7d6d20bd5b22c69b0e8bd1b756764c76b117306d29bbc16e09d.
//
// Solidity: event ListingAdded(uint256 indexed boardId, uint256 indexed listingId, uint256 strike, uint256 skew)
func (_Lyra *LyraFilterer) ParseListingAdded(log types.Log) (*LyraListingAdded, error) {
	event := new(LyraListingAdded)
	if err := _Lyra.contract.UnpackLog(event, "ListingAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraListingSkewSetIterator is returned from FilterListingSkewSet and is used to iterate over the raw logs and unpacked data for ListingSkewSet events raised by the Lyra contract.
type LyraListingSkewSetIterator struct {
	Event *LyraListingSkewSet // Event containing the contract specifics and raw log

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
func (it *LyraListingSkewSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraListingSkewSet)
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
		it.Event = new(LyraListingSkewSet)
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
func (it *LyraListingSkewSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraListingSkewSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraListingSkewSet represents a ListingSkewSet event raised by the Lyra contract.
type LyraListingSkewSet struct {
	ListingId *big.Int
	Skew      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingSkewSet is a free log retrieval operation binding the contract event 0x807f7e992c23befc534fdba4f8206693b5b6c4dd10b67620b355911654d0a957.
//
// Solidity: event ListingSkewSet(uint256 indexed listingId, uint256 skew)
func (_Lyra *LyraFilterer) FilterListingSkewSet(opts *bind.FilterOpts, listingId []*big.Int) (*LyraListingSkewSetIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "ListingSkewSet", listingIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraListingSkewSetIterator{contract: _Lyra.contract, event: "ListingSkewSet", logs: logs, sub: sub}, nil
}

// WatchListingSkewSet is a free log subscription operation binding the contract event 0x807f7e992c23befc534fdba4f8206693b5b6c4dd10b67620b355911654d0a957.
//
// Solidity: event ListingSkewSet(uint256 indexed listingId, uint256 skew)
func (_Lyra *LyraFilterer) WatchListingSkewSet(opts *bind.WatchOpts, sink chan<- *LyraListingSkewSet, listingId []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "ListingSkewSet", listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraListingSkewSet)
				if err := _Lyra.contract.UnpackLog(event, "ListingSkewSet", log); err != nil {
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

// ParseListingSkewSet is a log parse operation binding the contract event 0x807f7e992c23befc534fdba4f8206693b5b6c4dd10b67620b355911654d0a957.
//
// Solidity: event ListingSkewSet(uint256 indexed listingId, uint256 skew)
func (_Lyra *LyraFilterer) ParseListingSkewSet(log types.Log) (*LyraListingSkewSet, error) {
	event := new(LyraListingSkewSet)
	if err := _Lyra.contract.UnpackLog(event, "ListingSkewSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lyra contract.
type LyraOwnershipTransferredIterator struct {
	Event *LyraOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LyraOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraOwnershipTransferred)
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
		it.Event = new(LyraOwnershipTransferred)
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
func (it *LyraOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraOwnershipTransferred represents a OwnershipTransferred event raised by the Lyra contract.
type LyraOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lyra *LyraFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LyraOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LyraOwnershipTransferredIterator{contract: _Lyra.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lyra *LyraFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LyraOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraOwnershipTransferred)
				if err := _Lyra.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Lyra *LyraFilterer) ParseOwnershipTransferred(log types.Log) (*LyraOwnershipTransferred, error) {
	event := new(LyraOwnershipTransferred)
	if err := _Lyra.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraPositionClosedIterator is returned from FilterPositionClosed and is used to iterate over the raw logs and unpacked data for PositionClosed events raised by the Lyra contract.
type LyraPositionClosedIterator struct {
	Event *LyraPositionClosed // Event containing the contract specifics and raw log

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
func (it *LyraPositionClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraPositionClosed)
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
		it.Event = new(LyraPositionClosed)
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
func (it *LyraPositionClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraPositionClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraPositionClosed represents a PositionClosed event raised by the Lyra contract.
type LyraPositionClosed struct {
	Trader    common.Address
	ListingId *big.Int
	TradeType uint8
	Amount    *big.Int
	TotalCost *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPositionClosed is a free log retrieval operation binding the contract event 0x5a0cd91f269c536b8167c8fc4e2b859e3db49a3fb022e5db3e380a5897540584.
//
// Solidity: event PositionClosed(address indexed trader, uint256 indexed listingId, uint8 indexed tradeType, uint256 amount, uint256 totalCost)
func (_Lyra *LyraFilterer) FilterPositionClosed(opts *bind.FilterOpts, trader []common.Address, listingId []*big.Int, tradeType []uint8) (*LyraPositionClosedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var tradeTypeRule []interface{}
	for _, tradeTypeItem := range tradeType {
		tradeTypeRule = append(tradeTypeRule, tradeTypeItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "PositionClosed", traderRule, listingIdRule, tradeTypeRule)
	if err != nil {
		return nil, err
	}
	return &LyraPositionClosedIterator{contract: _Lyra.contract, event: "PositionClosed", logs: logs, sub: sub}, nil
}

// WatchPositionClosed is a free log subscription operation binding the contract event 0x5a0cd91f269c536b8167c8fc4e2b859e3db49a3fb022e5db3e380a5897540584.
//
// Solidity: event PositionClosed(address indexed trader, uint256 indexed listingId, uint8 indexed tradeType, uint256 amount, uint256 totalCost)
func (_Lyra *LyraFilterer) WatchPositionClosed(opts *bind.WatchOpts, sink chan<- *LyraPositionClosed, trader []common.Address, listingId []*big.Int, tradeType []uint8) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var tradeTypeRule []interface{}
	for _, tradeTypeItem := range tradeType {
		tradeTypeRule = append(tradeTypeRule, tradeTypeItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "PositionClosed", traderRule, listingIdRule, tradeTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraPositionClosed)
				if err := _Lyra.contract.UnpackLog(event, "PositionClosed", log); err != nil {
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

// ParsePositionClosed is a log parse operation binding the contract event 0x5a0cd91f269c536b8167c8fc4e2b859e3db49a3fb022e5db3e380a5897540584.
//
// Solidity: event PositionClosed(address indexed trader, uint256 indexed listingId, uint8 indexed tradeType, uint256 amount, uint256 totalCost)
func (_Lyra *LyraFilterer) ParsePositionClosed(log types.Log) (*LyraPositionClosed, error) {
	event := new(LyraPositionClosed)
	if err := _Lyra.contract.UnpackLog(event, "PositionClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraPositionOpenedIterator is returned from FilterPositionOpened and is used to iterate over the raw logs and unpacked data for PositionOpened events raised by the Lyra contract.
type LyraPositionOpenedIterator struct {
	Event *LyraPositionOpened // Event containing the contract specifics and raw log

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
func (it *LyraPositionOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraPositionOpened)
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
		it.Event = new(LyraPositionOpened)
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
func (it *LyraPositionOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraPositionOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraPositionOpened represents a PositionOpened event raised by the Lyra contract.
type LyraPositionOpened struct {
	Trader    common.Address
	ListingId *big.Int
	TradeType uint8
	Amount    *big.Int
	TotalCost *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPositionOpened is a free log retrieval operation binding the contract event 0x9a99f0d83ea15216987ecd4ed332a7338eafafe12db86adf283528ec63c02d3f.
//
// Solidity: event PositionOpened(address indexed trader, uint256 indexed listingId, uint8 indexed tradeType, uint256 amount, uint256 totalCost)
func (_Lyra *LyraFilterer) FilterPositionOpened(opts *bind.FilterOpts, trader []common.Address, listingId []*big.Int, tradeType []uint8) (*LyraPositionOpenedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var tradeTypeRule []interface{}
	for _, tradeTypeItem := range tradeType {
		tradeTypeRule = append(tradeTypeRule, tradeTypeItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "PositionOpened", traderRule, listingIdRule, tradeTypeRule)
	if err != nil {
		return nil, err
	}
	return &LyraPositionOpenedIterator{contract: _Lyra.contract, event: "PositionOpened", logs: logs, sub: sub}, nil
}

// WatchPositionOpened is a free log subscription operation binding the contract event 0x9a99f0d83ea15216987ecd4ed332a7338eafafe12db86adf283528ec63c02d3f.
//
// Solidity: event PositionOpened(address indexed trader, uint256 indexed listingId, uint8 indexed tradeType, uint256 amount, uint256 totalCost)
func (_Lyra *LyraFilterer) WatchPositionOpened(opts *bind.WatchOpts, sink chan<- *LyraPositionOpened, trader []common.Address, listingId []*big.Int, tradeType []uint8) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var tradeTypeRule []interface{}
	for _, tradeTypeItem := range tradeType {
		tradeTypeRule = append(tradeTypeRule, tradeTypeItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "PositionOpened", traderRule, listingIdRule, tradeTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraPositionOpened)
				if err := _Lyra.contract.UnpackLog(event, "PositionOpened", log); err != nil {
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

// ParsePositionOpened is a log parse operation binding the contract event 0x9a99f0d83ea15216987ecd4ed332a7338eafafe12db86adf283528ec63c02d3f.
//
// Solidity: event PositionOpened(address indexed trader, uint256 indexed listingId, uint8 indexed tradeType, uint256 amount, uint256 totalCost)
func (_Lyra *LyraFilterer) ParsePositionOpened(log types.Log) (*LyraPositionOpened, error) {
	event := new(LyraPositionOpened)
	if err := _Lyra.contract.UnpackLog(event, "PositionOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
