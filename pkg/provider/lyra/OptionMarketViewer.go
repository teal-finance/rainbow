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

// ILiquidityPoolLiquidity is an auto generated low-level Go binding around an user-defined struct.
type ILiquidityPoolLiquidity struct {
	FreeCollatLiquidity *big.Int
	UsedCollatLiquidity *big.Int
	FreeDeltaLiquidity  *big.Int
	UsedDeltaLiquidity  *big.Int
}

// ILyraGlobalsPricingGlobals is an auto generated low-level Go binding around an user-defined struct.
type ILyraGlobalsPricingGlobals struct {
	OptionPriceFeeCoefficient *big.Int
	SpotPriceFeeCoefficient   *big.Int
	VegaFeeCoefficient        *big.Int
	VegaNormFactor            *big.Int
	StandardSize              *big.Int
	SkewAdjustmentFactor      *big.Int
	RateAndCarry              *big.Int
	MinDelta                  *big.Int
	VolatilityCutoff          *big.Int
	SpotPrice                 *big.Int
}

// IOptionMarketOptionBoard is an auto generated low-level Go binding around an user-defined struct.
type IOptionMarketOptionBoard struct {
	Id         *big.Int
	Expiry     *big.Int
	Iv         *big.Int
	Frozen     bool
	ListingIds []*big.Int
}

// IOptionMarketOptionListing is an auto generated low-level Go binding around an user-defined struct.
type IOptionMarketOptionListing struct {
	Id        *big.Int
	Strike    *big.Int
	Skew      *big.Int
	LongCall  *big.Int
	ShortCall *big.Int
	LongPut   *big.Int
	ShortPut  *big.Int
	BoardId   *big.Int
}

// IOptionMarketTrade is an auto generated low-level Go binding around an user-defined struct.
type IOptionMarketTrade struct {
	IsBuy     bool
	Amount    *big.Int
	Vol       *big.Int
	Expiry    *big.Int
	Liquidity ILiquidityPoolLiquidity
}

// OptionMarketViewerBoardView is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerBoardView struct {
	BoardId *big.Int
	Expiry  *big.Int
}

// OptionMarketViewerListingView is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerListingView struct {
	ListingId *big.Int
	BoardId   *big.Int
	Strike    *big.Int
	Expiry    *big.Int
	Iv        *big.Int
	Skew      *big.Int
	CallPrice *big.Int
	PutPrice  *big.Int
	CallDelta *big.Int
	PutDelta  *big.Int
	LongCall  *big.Int
	ShortCall *big.Int
	LongPut   *big.Int
	ShortPut  *big.Int
}

// OptionMarketViewerTradePremiumView is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerTradePremiumView struct {
	ListingId      *big.Int
	Premium        *big.Int
	BasePrice      *big.Int
	VegaUtilFee    *big.Int
	OptionPriceFee *big.Int
	SpotPriceFee   *big.Int
	NewIv          *big.Int
}

// LyrapMetaData contains all meta data concerning the Lyrap contract.
var LyrapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"internalType\":\"structIOptionMarket.OptionListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"listingIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIOptionMarket.OptionBoard\",\"name\":\"board\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"}],\"internalType\":\"structILiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"}],\"internalType\":\"structIOptionMarket.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaNormFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rateAndCarry\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"volatilityCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"internalType\":\"structILyraGlobals.PricingGlobals\",\"name\":\"pricingGlobals\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"}],\"name\":\"_getPremiumForTrade\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaUtilFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newIv\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.TradePremiumView\",\"name\":\"premium\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blackScholes\",\"outputs\":[{\"internalType\":\"contractIBlackScholes\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"listingIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIOptionMarket.OptionBoard\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_boardId\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getClosePremiumsForBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaUtilFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newIv\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.TradePremiumView[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"internalType\":\"structIOptionMarket.OptionListing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getListingView\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"longCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPut\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.ListingView\",\"name\":\"listingView\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getListingViewAndBalance\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"longCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPut\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.ListingView\",\"name\":\"listingView\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"longCallAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutAmt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getListingsForBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"longCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPut\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.ListingView[]\",\"name\":\"boardListings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiveBoards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"boards\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_boardId\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getOpenPremiumsForBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaUtilFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newIv\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.TradePremiumView[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPremiumForClose\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaUtilFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newIv\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.TradePremiumView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPremiumForOpen\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaUtilFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newIv\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.TradePremiumView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPremiumForTrade\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaUtilFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newIv\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.TradePremiumView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_boardId\",\"type\":\"uint256\"},{\"internalType\":\"enumIOptionMarket.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPremiumsForBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaUtilFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newIv\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.TradePremiumView[]\",\"name\":\"tradePremiums\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globals\",\"outputs\":[{\"internalType\":\"contractILyraGlobals\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"greekCache\",\"outputs\":[{\"internalType\":\"contractIOptionGreekCache\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILyraGlobals\",\"name\":\"_globals\",\"type\":\"address\"},{\"internalType\":\"contractIOptionMarket\",\"name\":\"_optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractIOptionMarketPricer\",\"name\":\"_optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractIOptionGreekCache\",\"name\":\"_greekCache\",\"type\":\"address\"},{\"internalType\":\"contractIOptionToken\",\"name\":\"_optionToken\",\"type\":\"address\"},{\"internalType\":\"contractILiquidityPool\",\"name\":\"_liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractIBlackScholes\",\"name\":\"_blackScholes\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityPool\",\"outputs\":[{\"internalType\":\"contractILiquidityPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optionMarket\",\"outputs\":[{\"internalType\":\"contractIOptionMarket\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optionMarketPricer\",\"outputs\":[{\"internalType\":\"contractIOptionMarketPricer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optionToken\",\"outputs\":[{\"internalType\":\"contractIOptionToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LyrapABI is the input ABI used to generate the binding from.
// Deprecated: Use LyrapMetaData.ABI instead.
var LyrapABI = LyrapMetaData.ABI

// Lyrap is an auto generated Go binding around an Ethereum contract.
type Lyrap struct {
	LyrapCaller     // Read-only binding to the contract
	LyrapTransactor // Write-only binding to the contract
	LyrapFilterer   // Log filterer for contract events
}

// LyrapCaller is an auto generated read-only Go binding around an Ethereum contract.
type LyrapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LyrapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LyrapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LyrapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LyrapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LyrapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LyrapSession struct {
	Contract     *Lyrap            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LyrapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LyrapCallerSession struct {
	Contract *LyrapCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LyrapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LyrapTransactorSession struct {
	Contract     *LyrapTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LyrapRaw is an auto generated low-level Go binding around an Ethereum contract.
type LyrapRaw struct {
	Contract *Lyrap // Generic contract binding to access the raw methods on
}

// LyrapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LyrapCallerRaw struct {
	Contract *LyrapCaller // Generic read-only contract binding to access the raw methods on
}

// LyrapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LyrapTransactorRaw struct {
	Contract *LyrapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLyrap creates a new instance of Lyrap, bound to a specific deployed contract.
func NewLyrap(address common.Address, backend bind.ContractBackend) (*Lyrap, error) {
	contract, err := bindLyrap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lyrap{LyrapCaller: LyrapCaller{contract: contract}, LyrapTransactor: LyrapTransactor{contract: contract}, LyrapFilterer: LyrapFilterer{contract: contract}}, nil
}

// NewLyrapCaller creates a new read-only instance of Lyrap, bound to a specific deployed contract.
func NewLyrapCaller(address common.Address, caller bind.ContractCaller) (*LyrapCaller, error) {
	contract, err := bindLyrap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LyrapCaller{contract: contract}, nil
}

// NewLyrapTransactor creates a new write-only instance of Lyrap, bound to a specific deployed contract.
func NewLyrapTransactor(address common.Address, transactor bind.ContractTransactor) (*LyrapTransactor, error) {
	contract, err := bindLyrap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LyrapTransactor{contract: contract}, nil
}

// NewLyrapFilterer creates a new log filterer instance of Lyrap, bound to a specific deployed contract.
func NewLyrapFilterer(address common.Address, filterer bind.ContractFilterer) (*LyrapFilterer, error) {
	contract, err := bindLyrap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LyrapFilterer{contract: contract}, nil
}

// bindLyrap binds a generic wrapper to an already deployed contract.
func bindLyrap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LyrapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lyrap *LyrapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lyrap.Contract.LyrapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lyrap *LyrapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lyrap.Contract.LyrapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lyrap *LyrapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lyrap.Contract.LyrapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lyrap *LyrapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lyrap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lyrap *LyrapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lyrap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lyrap *LyrapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lyrap.Contract.contract.Transact(opts, method, params...)
}

// InternalGetPremiumForTrade is a free data retrieval call binding the contract method 0x5dd8c74d.
//
// Solidity: function _getPremiumForTrade((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) listing, (uint256,uint256,uint256,bool,uint256[]) board, (bool,uint256,uint256,uint256,(uint256,uint256,uint256,uint256)) trade, (uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256) pricingGlobals, bool isCall) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) premium)
func (_Lyrap *LyrapCaller) InternalGetPremiumForTrade(opts *bind.CallOpts, listing IOptionMarketOptionListing, board IOptionMarketOptionBoard, trade IOptionMarketTrade, pricingGlobals ILyraGlobalsPricingGlobals, isCall bool) (OptionMarketViewerTradePremiumView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "_getPremiumForTrade", listing, board, trade, pricingGlobals, isCall)

	if err != nil {
		return *new(OptionMarketViewerTradePremiumView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerTradePremiumView)).(*OptionMarketViewerTradePremiumView)

	return out0, err

}

// InternalGetPremiumForTrade is a free data retrieval call binding the contract method 0x5dd8c74d.
//
// Solidity: function _getPremiumForTrade((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) listing, (uint256,uint256,uint256,bool,uint256[]) board, (bool,uint256,uint256,uint256,(uint256,uint256,uint256,uint256)) trade, (uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256) pricingGlobals, bool isCall) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) premium)
func (_Lyrap *LyrapSession) InternalGetPremiumForTrade(listing IOptionMarketOptionListing, board IOptionMarketOptionBoard, trade IOptionMarketTrade, pricingGlobals ILyraGlobalsPricingGlobals, isCall bool) (OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.InternalGetPremiumForTrade(&_Lyrap.CallOpts, listing, board, trade, pricingGlobals, isCall)
}

// InternalGetPremiumForTrade is a free data retrieval call binding the contract method 0x5dd8c74d.
//
// Solidity: function _getPremiumForTrade((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) listing, (uint256,uint256,uint256,bool,uint256[]) board, (bool,uint256,uint256,uint256,(uint256,uint256,uint256,uint256)) trade, (uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256) pricingGlobals, bool isCall) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) premium)
func (_Lyrap *LyrapCallerSession) InternalGetPremiumForTrade(listing IOptionMarketOptionListing, board IOptionMarketOptionBoard, trade IOptionMarketTrade, pricingGlobals ILyraGlobalsPricingGlobals, isCall bool) (OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.InternalGetPremiumForTrade(&_Lyrap.CallOpts, listing, board, trade, pricingGlobals, isCall)
}

// BlackScholes is a free data retrieval call binding the contract method 0x3c195aa7.
//
// Solidity: function blackScholes() view returns(address)
func (_Lyrap *LyrapCaller) BlackScholes(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "blackScholes")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlackScholes is a free data retrieval call binding the contract method 0x3c195aa7.
//
// Solidity: function blackScholes() view returns(address)
func (_Lyrap *LyrapSession) BlackScholes() (common.Address, error) {
	return _Lyrap.Contract.BlackScholes(&_Lyrap.CallOpts)
}

// BlackScholes is a free data retrieval call binding the contract method 0x3c195aa7.
//
// Solidity: function blackScholes() view returns(address)
func (_Lyrap *LyrapCallerSession) BlackScholes() (common.Address, error) {
	return _Lyrap.Contract.BlackScholes(&_Lyrap.CallOpts)
}

// GetBoard is a free data retrieval call binding the contract method 0x45e09e54.
//
// Solidity: function getBoard(uint256 boardId) view returns((uint256,uint256,uint256,bool,uint256[]))
func (_Lyrap *LyrapCaller) GetBoard(opts *bind.CallOpts, boardId *big.Int) (IOptionMarketOptionBoard, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getBoard", boardId)

	if err != nil {
		return *new(IOptionMarketOptionBoard), err
	}

	out0 := *abi.ConvertType(out[0], new(IOptionMarketOptionBoard)).(*IOptionMarketOptionBoard)

	return out0, err

}

// GetBoard is a free data retrieval call binding the contract method 0x45e09e54.
//
// Solidity: function getBoard(uint256 boardId) view returns((uint256,uint256,uint256,bool,uint256[]))
func (_Lyrap *LyrapSession) GetBoard(boardId *big.Int) (IOptionMarketOptionBoard, error) {
	return _Lyrap.Contract.GetBoard(&_Lyrap.CallOpts, boardId)
}

// GetBoard is a free data retrieval call binding the contract method 0x45e09e54.
//
// Solidity: function getBoard(uint256 boardId) view returns((uint256,uint256,uint256,bool,uint256[]))
func (_Lyrap *LyrapCallerSession) GetBoard(boardId *big.Int) (IOptionMarketOptionBoard, error) {
	return _Lyrap.Contract.GetBoard(&_Lyrap.CallOpts, boardId)
}

// GetClosePremiumsForBoard is a free data retrieval call binding the contract method 0x1740aac5.
//
// Solidity: function getClosePremiumsForBoard(uint256 _boardId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Lyrap *LyrapCaller) GetClosePremiumsForBoard(opts *bind.CallOpts, _boardId *big.Int, tradeType uint8, amount *big.Int) ([]OptionMarketViewerTradePremiumView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getClosePremiumsForBoard", _boardId, tradeType, amount)

	if err != nil {
		return *new([]OptionMarketViewerTradePremiumView), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerTradePremiumView)).(*[]OptionMarketViewerTradePremiumView)

	return out0, err

}

// GetClosePremiumsForBoard is a free data retrieval call binding the contract method 0x1740aac5.
//
// Solidity: function getClosePremiumsForBoard(uint256 _boardId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Lyrap *LyrapSession) GetClosePremiumsForBoard(_boardId *big.Int, tradeType uint8, amount *big.Int) ([]OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetClosePremiumsForBoard(&_Lyrap.CallOpts, _boardId, tradeType, amount)
}

// GetClosePremiumsForBoard is a free data retrieval call binding the contract method 0x1740aac5.
//
// Solidity: function getClosePremiumsForBoard(uint256 _boardId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Lyrap *LyrapCallerSession) GetClosePremiumsForBoard(_boardId *big.Int, tradeType uint8, amount *big.Int) ([]OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetClosePremiumsForBoard(&_Lyrap.CallOpts, _boardId, tradeType, amount)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapCaller) GetListing(opts *bind.CallOpts, listingId *big.Int) (IOptionMarketOptionListing, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getListing", listingId)

	if err != nil {
		return *new(IOptionMarketOptionListing), err
	}

	out0 := *abi.ConvertType(out[0], new(IOptionMarketOptionListing)).(*IOptionMarketOptionListing)

	return out0, err

}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapSession) GetListing(listingId *big.Int) (IOptionMarketOptionListing, error) {
	return _Lyrap.Contract.GetListing(&_Lyrap.CallOpts, listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapCallerSession) GetListing(listingId *big.Int) (IOptionMarketOptionListing, error) {
	return _Lyrap.Contract.GetListing(&_Lyrap.CallOpts, listingId)
}

// GetListingView is a free data retrieval call binding the contract method 0x236c70d1.
//
// Solidity: function getListingView(uint256 listingId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256,uint256,uint256) listingView)
func (_Lyrap *LyrapCaller) GetListingView(opts *bind.CallOpts, listingId *big.Int) (OptionMarketViewerListingView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getListingView", listingId)

	if err != nil {
		return *new(OptionMarketViewerListingView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerListingView)).(*OptionMarketViewerListingView)

	return out0, err

}

// GetListingView is a free data retrieval call binding the contract method 0x236c70d1.
//
// Solidity: function getListingView(uint256 listingId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256,uint256,uint256) listingView)
func (_Lyrap *LyrapSession) GetListingView(listingId *big.Int) (OptionMarketViewerListingView, error) {
	return _Lyrap.Contract.GetListingView(&_Lyrap.CallOpts, listingId)
}

// GetListingView is a free data retrieval call binding the contract method 0x236c70d1.
//
// Solidity: function getListingView(uint256 listingId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256,uint256,uint256) listingView)
func (_Lyrap *LyrapCallerSession) GetListingView(listingId *big.Int) (OptionMarketViewerListingView, error) {
	return _Lyrap.Contract.GetListingView(&_Lyrap.CallOpts, listingId)
}

// GetListingViewAndBalance is a free data retrieval call binding the contract method 0x8dd915d1.
//
// Solidity: function getListingViewAndBalance(uint256 listingId, address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256,uint256,uint256) listingView, uint256 longCallAmt, uint256 longPutAmt, uint256 shortCallAmt, uint256 shortPutAmt)
func (_Lyrap *LyrapCaller) GetListingViewAndBalance(opts *bind.CallOpts, listingId *big.Int, user common.Address) (struct {
	ListingView  OptionMarketViewerListingView
	LongCallAmt  *big.Int
	LongPutAmt   *big.Int
	ShortCallAmt *big.Int
	ShortPutAmt  *big.Int
}, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getListingViewAndBalance", listingId, user)

	outstruct := new(struct {
		ListingView  OptionMarketViewerListingView
		LongCallAmt  *big.Int
		LongPutAmt   *big.Int
		ShortCallAmt *big.Int
		ShortPutAmt  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListingView = *abi.ConvertType(out[0], new(OptionMarketViewerListingView)).(*OptionMarketViewerListingView)
	outstruct.LongCallAmt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LongPutAmt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ShortCallAmt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ShortPutAmt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetListingViewAndBalance is a free data retrieval call binding the contract method 0x8dd915d1.
//
// Solidity: function getListingViewAndBalance(uint256 listingId, address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256,uint256,uint256) listingView, uint256 longCallAmt, uint256 longPutAmt, uint256 shortCallAmt, uint256 shortPutAmt)
func (_Lyrap *LyrapSession) GetListingViewAndBalance(listingId *big.Int, user common.Address) (struct {
	ListingView  OptionMarketViewerListingView
	LongCallAmt  *big.Int
	LongPutAmt   *big.Int
	ShortCallAmt *big.Int
	ShortPutAmt  *big.Int
}, error) {
	return _Lyrap.Contract.GetListingViewAndBalance(&_Lyrap.CallOpts, listingId, user)
}

// GetListingViewAndBalance is a free data retrieval call binding the contract method 0x8dd915d1.
//
// Solidity: function getListingViewAndBalance(uint256 listingId, address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256,uint256,uint256) listingView, uint256 longCallAmt, uint256 longPutAmt, uint256 shortCallAmt, uint256 shortPutAmt)
func (_Lyrap *LyrapCallerSession) GetListingViewAndBalance(listingId *big.Int, user common.Address) (struct {
	ListingView  OptionMarketViewerListingView
	LongCallAmt  *big.Int
	LongPutAmt   *big.Int
	ShortCallAmt *big.Int
	ShortPutAmt  *big.Int
}, error) {
	return _Lyrap.Contract.GetListingViewAndBalance(&_Lyrap.CallOpts, listingId, user)
}

// GetListingsForBoard is a free data retrieval call binding the contract method 0x19ca692a.
//
// Solidity: function getListingsForBoard(uint256 boardId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256,uint256,uint256)[] boardListings)
func (_Lyrap *LyrapCaller) GetListingsForBoard(opts *bind.CallOpts, boardId *big.Int) ([]OptionMarketViewerListingView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getListingsForBoard", boardId)

	if err != nil {
		return *new([]OptionMarketViewerListingView), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerListingView)).(*[]OptionMarketViewerListingView)

	return out0, err

}

// GetListingsForBoard is a free data retrieval call binding the contract method 0x19ca692a.
//
// Solidity: function getListingsForBoard(uint256 boardId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256,uint256,uint256)[] boardListings)
func (_Lyrap *LyrapSession) GetListingsForBoard(boardId *big.Int) ([]OptionMarketViewerListingView, error) {
	return _Lyrap.Contract.GetListingsForBoard(&_Lyrap.CallOpts, boardId)
}

// GetListingsForBoard is a free data retrieval call binding the contract method 0x19ca692a.
//
// Solidity: function getListingsForBoard(uint256 boardId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256,int256,uint256,uint256,uint256,uint256)[] boardListings)
func (_Lyrap *LyrapCallerSession) GetListingsForBoard(boardId *big.Int) ([]OptionMarketViewerListingView, error) {
	return _Lyrap.Contract.GetListingsForBoard(&_Lyrap.CallOpts, boardId)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0x7e7088eb.
//
// Solidity: function getLiveBoards() view returns((uint256,uint256)[] boards)
func (_Lyrap *LyrapCaller) GetLiveBoards(opts *bind.CallOpts) ([]OptionMarketViewerBoardView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getLiveBoards")

	if err != nil {
		return *new([]OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerBoardView)).(*[]OptionMarketViewerBoardView)

	return out0, err

}

// GetLiveBoards is a free data retrieval call binding the contract method 0x7e7088eb.
//
// Solidity: function getLiveBoards() view returns((uint256,uint256)[] boards)
func (_Lyrap *LyrapSession) GetLiveBoards() ([]OptionMarketViewerBoardView, error) {
	return _Lyrap.Contract.GetLiveBoards(&_Lyrap.CallOpts)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0x7e7088eb.
//
// Solidity: function getLiveBoards() view returns((uint256,uint256)[] boards)
func (_Lyrap *LyrapCallerSession) GetLiveBoards() ([]OptionMarketViewerBoardView, error) {
	return _Lyrap.Contract.GetLiveBoards(&_Lyrap.CallOpts)
}

// GetOpenPremiumsForBoard is a free data retrieval call binding the contract method 0xf8f6f7ce.
//
// Solidity: function getOpenPremiumsForBoard(uint256 _boardId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Lyrap *LyrapCaller) GetOpenPremiumsForBoard(opts *bind.CallOpts, _boardId *big.Int, tradeType uint8, amount *big.Int) ([]OptionMarketViewerTradePremiumView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getOpenPremiumsForBoard", _boardId, tradeType, amount)

	if err != nil {
		return *new([]OptionMarketViewerTradePremiumView), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerTradePremiumView)).(*[]OptionMarketViewerTradePremiumView)

	return out0, err

}

// GetOpenPremiumsForBoard is a free data retrieval call binding the contract method 0xf8f6f7ce.
//
// Solidity: function getOpenPremiumsForBoard(uint256 _boardId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Lyrap *LyrapSession) GetOpenPremiumsForBoard(_boardId *big.Int, tradeType uint8, amount *big.Int) ([]OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetOpenPremiumsForBoard(&_Lyrap.CallOpts, _boardId, tradeType, amount)
}

// GetOpenPremiumsForBoard is a free data retrieval call binding the contract method 0xf8f6f7ce.
//
// Solidity: function getOpenPremiumsForBoard(uint256 _boardId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Lyrap *LyrapCallerSession) GetOpenPremiumsForBoard(_boardId *big.Int, tradeType uint8, amount *big.Int) ([]OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetOpenPremiumsForBoard(&_Lyrap.CallOpts, _boardId, tradeType, amount)
}

// GetPremiumForClose is a free data retrieval call binding the contract method 0xf43acb68.
//
// Solidity: function getPremiumForClose(uint256 _listingId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapCaller) GetPremiumForClose(opts *bind.CallOpts, _listingId *big.Int, tradeType uint8, amount *big.Int) (OptionMarketViewerTradePremiumView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getPremiumForClose", _listingId, tradeType, amount)

	if err != nil {
		return *new(OptionMarketViewerTradePremiumView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerTradePremiumView)).(*OptionMarketViewerTradePremiumView)

	return out0, err

}

// GetPremiumForClose is a free data retrieval call binding the contract method 0xf43acb68.
//
// Solidity: function getPremiumForClose(uint256 _listingId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapSession) GetPremiumForClose(_listingId *big.Int, tradeType uint8, amount *big.Int) (OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetPremiumForClose(&_Lyrap.CallOpts, _listingId, tradeType, amount)
}

// GetPremiumForClose is a free data retrieval call binding the contract method 0xf43acb68.
//
// Solidity: function getPremiumForClose(uint256 _listingId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapCallerSession) GetPremiumForClose(_listingId *big.Int, tradeType uint8, amount *big.Int) (OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetPremiumForClose(&_Lyrap.CallOpts, _listingId, tradeType, amount)
}

// GetPremiumForOpen is a free data retrieval call binding the contract method 0x3904b0c8.
//
// Solidity: function getPremiumForOpen(uint256 _listingId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapCaller) GetPremiumForOpen(opts *bind.CallOpts, _listingId *big.Int, tradeType uint8, amount *big.Int) (OptionMarketViewerTradePremiumView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getPremiumForOpen", _listingId, tradeType, amount)

	if err != nil {
		return *new(OptionMarketViewerTradePremiumView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerTradePremiumView)).(*OptionMarketViewerTradePremiumView)

	return out0, err

}

// GetPremiumForOpen is a free data retrieval call binding the contract method 0x3904b0c8.
//
// Solidity: function getPremiumForOpen(uint256 _listingId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapSession) GetPremiumForOpen(_listingId *big.Int, tradeType uint8, amount *big.Int) (OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetPremiumForOpen(&_Lyrap.CallOpts, _listingId, tradeType, amount)
}

// GetPremiumForOpen is a free data retrieval call binding the contract method 0x3904b0c8.
//
// Solidity: function getPremiumForOpen(uint256 _listingId, uint8 tradeType, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapCallerSession) GetPremiumForOpen(_listingId *big.Int, tradeType uint8, amount *big.Int) (OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetPremiumForOpen(&_Lyrap.CallOpts, _listingId, tradeType, amount)
}

// GetPremiumForTrade is a free data retrieval call binding the contract method 0x467e24c1.
//
// Solidity: function getPremiumForTrade(uint256 _listingId, uint8 tradeType, bool isBuy, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapCaller) GetPremiumForTrade(opts *bind.CallOpts, _listingId *big.Int, tradeType uint8, isBuy bool, amount *big.Int) (OptionMarketViewerTradePremiumView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getPremiumForTrade", _listingId, tradeType, isBuy, amount)

	if err != nil {
		return *new(OptionMarketViewerTradePremiumView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerTradePremiumView)).(*OptionMarketViewerTradePremiumView)

	return out0, err

}

// GetPremiumForTrade is a free data retrieval call binding the contract method 0x467e24c1.
//
// Solidity: function getPremiumForTrade(uint256 _listingId, uint8 tradeType, bool isBuy, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapSession) GetPremiumForTrade(_listingId *big.Int, tradeType uint8, isBuy bool, amount *big.Int) (OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetPremiumForTrade(&_Lyrap.CallOpts, _listingId, tradeType, isBuy, amount)
}

// GetPremiumForTrade is a free data retrieval call binding the contract method 0x467e24c1.
//
// Solidity: function getPremiumForTrade(uint256 _listingId, uint8 tradeType, bool isBuy, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyrap *LyrapCallerSession) GetPremiumForTrade(_listingId *big.Int, tradeType uint8, isBuy bool, amount *big.Int) (OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetPremiumForTrade(&_Lyrap.CallOpts, _listingId, tradeType, isBuy, amount)
}

// GetPremiumsForBoard is a free data retrieval call binding the contract method 0x552a2590.
//
// Solidity: function getPremiumsForBoard(uint256 _boardId, uint8 tradeType, bool isBuy, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] tradePremiums)
func (_Lyrap *LyrapCaller) GetPremiumsForBoard(opts *bind.CallOpts, _boardId *big.Int, tradeType uint8, isBuy bool, amount *big.Int) ([]OptionMarketViewerTradePremiumView, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "getPremiumsForBoard", _boardId, tradeType, isBuy, amount)

	if err != nil {
		return *new([]OptionMarketViewerTradePremiumView), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerTradePremiumView)).(*[]OptionMarketViewerTradePremiumView)

	return out0, err

}

// GetPremiumsForBoard is a free data retrieval call binding the contract method 0x552a2590.
//
// Solidity: function getPremiumsForBoard(uint256 _boardId, uint8 tradeType, bool isBuy, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] tradePremiums)
func (_Lyrap *LyrapSession) GetPremiumsForBoard(_boardId *big.Int, tradeType uint8, isBuy bool, amount *big.Int) ([]OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetPremiumsForBoard(&_Lyrap.CallOpts, _boardId, tradeType, isBuy, amount)
}

// GetPremiumsForBoard is a free data retrieval call binding the contract method 0x552a2590.
//
// Solidity: function getPremiumsForBoard(uint256 _boardId, uint8 tradeType, bool isBuy, uint256 amount) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] tradePremiums)
func (_Lyrap *LyrapCallerSession) GetPremiumsForBoard(_boardId *big.Int, tradeType uint8, isBuy bool, amount *big.Int) ([]OptionMarketViewerTradePremiumView, error) {
	return _Lyrap.Contract.GetPremiumsForBoard(&_Lyrap.CallOpts, _boardId, tradeType, isBuy, amount)
}

// Globals is a free data retrieval call binding the contract method 0xc3124525.
//
// Solidity: function globals() view returns(address)
func (_Lyrap *LyrapCaller) Globals(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "globals")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Globals is a free data retrieval call binding the contract method 0xc3124525.
//
// Solidity: function globals() view returns(address)
func (_Lyrap *LyrapSession) Globals() (common.Address, error) {
	return _Lyrap.Contract.Globals(&_Lyrap.CallOpts)
}

// Globals is a free data retrieval call binding the contract method 0xc3124525.
//
// Solidity: function globals() view returns(address)
func (_Lyrap *LyrapCallerSession) Globals() (common.Address, error) {
	return _Lyrap.Contract.Globals(&_Lyrap.CallOpts)
}

// GreekCache is a free data retrieval call binding the contract method 0x38b74054.
//
// Solidity: function greekCache() view returns(address)
func (_Lyrap *LyrapCaller) GreekCache(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "greekCache")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GreekCache is a free data retrieval call binding the contract method 0x38b74054.
//
// Solidity: function greekCache() view returns(address)
func (_Lyrap *LyrapSession) GreekCache() (common.Address, error) {
	return _Lyrap.Contract.GreekCache(&_Lyrap.CallOpts)
}

// GreekCache is a free data retrieval call binding the contract method 0x38b74054.
//
// Solidity: function greekCache() view returns(address)
func (_Lyrap *LyrapCallerSession) GreekCache() (common.Address, error) {
	return _Lyrap.Contract.GreekCache(&_Lyrap.CallOpts)
}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_Lyrap *LyrapCaller) LiquidityPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "liquidityPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_Lyrap *LyrapSession) LiquidityPool() (common.Address, error) {
	return _Lyrap.Contract.LiquidityPool(&_Lyrap.CallOpts)
}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_Lyrap *LyrapCallerSession) LiquidityPool() (common.Address, error) {
	return _Lyrap.Contract.LiquidityPool(&_Lyrap.CallOpts)
}

// OptionMarket is a free data retrieval call binding the contract method 0xbb4a9f88.
//
// Solidity: function optionMarket() view returns(address)
func (_Lyrap *LyrapCaller) OptionMarket(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "optionMarket")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionMarket is a free data retrieval call binding the contract method 0xbb4a9f88.
//
// Solidity: function optionMarket() view returns(address)
func (_Lyrap *LyrapSession) OptionMarket() (common.Address, error) {
	return _Lyrap.Contract.OptionMarket(&_Lyrap.CallOpts)
}

// OptionMarket is a free data retrieval call binding the contract method 0xbb4a9f88.
//
// Solidity: function optionMarket() view returns(address)
func (_Lyrap *LyrapCallerSession) OptionMarket() (common.Address, error) {
	return _Lyrap.Contract.OptionMarket(&_Lyrap.CallOpts)
}

// OptionMarketPricer is a free data retrieval call binding the contract method 0xa5a249d0.
//
// Solidity: function optionMarketPricer() view returns(address)
func (_Lyrap *LyrapCaller) OptionMarketPricer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "optionMarketPricer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionMarketPricer is a free data retrieval call binding the contract method 0xa5a249d0.
//
// Solidity: function optionMarketPricer() view returns(address)
func (_Lyrap *LyrapSession) OptionMarketPricer() (common.Address, error) {
	return _Lyrap.Contract.OptionMarketPricer(&_Lyrap.CallOpts)
}

// OptionMarketPricer is a free data retrieval call binding the contract method 0xa5a249d0.
//
// Solidity: function optionMarketPricer() view returns(address)
func (_Lyrap *LyrapCallerSession) OptionMarketPricer() (common.Address, error) {
	return _Lyrap.Contract.OptionMarketPricer(&_Lyrap.CallOpts)
}

// OptionToken is a free data retrieval call binding the contract method 0x2bab754b.
//
// Solidity: function optionToken() view returns(address)
func (_Lyrap *LyrapCaller) OptionToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyrap.contract.Call(opts, &out, "optionToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionToken is a free data retrieval call binding the contract method 0x2bab754b.
//
// Solidity: function optionToken() view returns(address)
func (_Lyrap *LyrapSession) OptionToken() (common.Address, error) {
	return _Lyrap.Contract.OptionToken(&_Lyrap.CallOpts)
}

// OptionToken is a free data retrieval call binding the contract method 0x2bab754b.
//
// Solidity: function optionToken() view returns(address)
func (_Lyrap *LyrapCallerSession) OptionToken() (common.Address, error) {
	return _Lyrap.Contract.OptionToken(&_Lyrap.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0xf0d4d592.
//
// Solidity: function init(address _globals, address _optionMarket, address _optionMarketPricer, address _greekCache, address _optionToken, address _liquidityPool, address _blackScholes) returns()
func (_Lyrap *LyrapTransactor) Init(opts *bind.TransactOpts, _globals common.Address, _optionMarket common.Address, _optionMarketPricer common.Address, _greekCache common.Address, _optionToken common.Address, _liquidityPool common.Address, _blackScholes common.Address) (*types.Transaction, error) {
	return _Lyrap.contract.Transact(opts, "init", _globals, _optionMarket, _optionMarketPricer, _greekCache, _optionToken, _liquidityPool, _blackScholes)
}

// Init is a paid mutator transaction binding the contract method 0xf0d4d592.
//
// Solidity: function init(address _globals, address _optionMarket, address _optionMarketPricer, address _greekCache, address _optionToken, address _liquidityPool, address _blackScholes) returns()
func (_Lyrap *LyrapSession) Init(_globals common.Address, _optionMarket common.Address, _optionMarketPricer common.Address, _greekCache common.Address, _optionToken common.Address, _liquidityPool common.Address, _blackScholes common.Address) (*types.Transaction, error) {
	return _Lyrap.Contract.Init(&_Lyrap.TransactOpts, _globals, _optionMarket, _optionMarketPricer, _greekCache, _optionToken, _liquidityPool, _blackScholes)
}

// Init is a paid mutator transaction binding the contract method 0xf0d4d592.
//
// Solidity: function init(address _globals, address _optionMarket, address _optionMarketPricer, address _greekCache, address _optionToken, address _liquidityPool, address _blackScholes) returns()
func (_Lyrap *LyrapTransactorSession) Init(_globals common.Address, _optionMarket common.Address, _optionMarketPricer common.Address, _greekCache common.Address, _optionToken common.Address, _liquidityPool common.Address, _blackScholes common.Address) (*types.Transaction, error) {
	return _Lyrap.Contract.Init(&_Lyrap.TransactOpts, _globals, _optionMarket, _optionMarketPricer, _greekCache, _optionToken, _liquidityPool, _blackScholes)
}
