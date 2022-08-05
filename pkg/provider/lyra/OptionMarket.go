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

// OptionMarketLiquidationEventData is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketLiquidationEventData struct {
	RewardBeneficiary common.Address
	Caller            common.Address
	ReturnCollateral  *big.Int
	LpPremiums        *big.Int
	LpFee             *big.Int
	LiquidatorFee     *big.Int
	SmFee             *big.Int
	InsolventAmount   *big.Int
}

// OptionMarketOptionBoard is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketOptionBoard struct {
	Id        *big.Int
	Expiry    *big.Int
	Iv        *big.Int
	Frozen    bool
	StrikeIds []*big.Int
}

// OptionMarketOptionMarketParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketOptionMarketParameters struct {
	MaxBoardExpiry          *big.Int
	SecurityModule          common.Address
	FeePortionReserved      *big.Int
	StaticBaseSettlementFee *big.Int
}

// OptionMarketPricerTradeResult is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketPricerTradeResult struct {
	Amount         *big.Int
	Premium        *big.Int
	OptionPriceFee *big.Int
	SpotPriceFee   *big.Int
	VegaUtilFee    OptionMarketPricerVegaUtilFeeComponents
	VarianceFee    OptionMarketPricerVarianceFeeComponents
	TotalFee       *big.Int
	TotalCost      *big.Int
	VolTraded      *big.Int
	NewBaseIv      *big.Int
	NewSkew        *big.Int
}

// OptionMarketPricerVarianceFeeComponents is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketPricerVarianceFeeComponents struct {
	VarianceFeeCoefficient *big.Int
	Vega                   *big.Int
	VegaCoefficient        *big.Int
	Skew                   *big.Int
	SkewCoefficient        *big.Int
	IvVariance             *big.Int
	IvVarianceCoefficient  *big.Int
	VarianceFee            *big.Int
}

// OptionMarketPricerVegaUtilFeeComponents is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketPricerVegaUtilFeeComponents struct {
	PreTradeAmmNetStdVega  *big.Int
	PostTradeAmmNetStdVega *big.Int
	VegaUtil               *big.Int
	VolTraded              *big.Int
	NAV                    *big.Int
	VegaUtilFee            *big.Int
}

// OptionMarketResult is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketResult struct {
	PositionId *big.Int
	TotalCost  *big.Int
	TotalFee   *big.Int
}

// OptionMarketStrike is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketStrike struct {
	Id             *big.Int
	StrikePrice    *big.Int
	Skew           *big.Int
	LongCall       *big.Int
	ShortCallBase  *big.Int
	ShortCallQuote *big.Int
	LongPut        *big.Int
	ShortPut       *big.Int
	BoardId        *big.Int
}

// OptionMarketTradeEventData is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketTradeEventData struct {
	Expiry          *big.Int
	StrikePrice     *big.Int
	OptionType      uint8
	TradeDirection  uint8
	Amount          *big.Int
	SetCollateralTo *big.Int
	IsForceClose    bool
	SpotPrice       *big.Int
	ReservedFee     *big.Int
	TotalCost       *big.Int
}

// OptionMarketTradeInputParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketTradeInputParameters struct {
	StrikeId        *big.Int
	PositionId      *big.Int
	Iterations      *big.Int
	OptionType      uint8
	Amount          *big.Int
	SetCollateralTo *big.Int
	MinTotalCost    *big.Int
	MaxTotalCost    *big.Int
}

// LyraMetaData contains all meta data concerning the Lyra contract.
var LyraMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"}],\"name\":\"AlreadyInitialised\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BaseTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardAlreadySettled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"}],\"name\":\"BoardExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardNotExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"BoardNotFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"enumOptionMarket.NonZeroValues\",\"name\":\"valueType\",\"type\":\"uint8\"}],\"name\":\"ExpectedNonZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"InvalidBoardId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"}],\"name\":\"InvalidExpiryTimestamp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"}],\"name\":\"InvalidOptionMarketParams\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"InvalidStrikeId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"}],\"name\":\"OnlyNominatedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"}],\"name\":\"OnlySecurityModule\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"QuoteTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikesLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewsLength\",\"type\":\"uint256\"}],\"name\":\"StrikeSkewLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"TotalCostOutsideOfSpecifiedBounds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tradeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"TradeIterationsHasRemainder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"}],\"name\":\"BoardBaseIvSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"BoardCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"BoardFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spotPriceAtExpiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUserLongProfitQuote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBoardLongCallCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBoardLongPutCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAMMShortCallProfitBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAMMShortCallProfitQuote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAMMShortPutProfitQuote\",\"type\":\"uint256\"}],\"name\":\"BoardSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"}],\"name\":\"OptionMarketParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"SMClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"}],\"name\":\"StrikeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"}],\"name\":\"StrikeSkewSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"enumOptionMarket.TradeDirection\",\"name\":\"tradeDirection\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"setCollateralTo\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForceClose\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOptionMarket.TradeEventData\",\"name\":\"trade\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"preTradeAmmNetStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"postTradeAmmNetStdVega\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"vegaUtil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"volTraded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaUtilFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VegaUtilFeeComponents\",\"name\":\"vegaUtilFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"varianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVariance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeComponents\",\"name\":\"varianceFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"volTraded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newBaseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newSkew\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOptionMarketPricer.TradeResult[]\",\"name\":\"tradeResults\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"rewardBeneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"returnCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpPremiums\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"insolventAmount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOptionMarket.LiquidationEventData\",\"name\":\"liquidation\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountCollateral\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"}],\"name\":\"addStrikeToBoard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boardToPriceAtExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"setCollateralTo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.TradeInputParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"closePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.Result\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"strikePrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"skews\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"createOptionBoard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"setCollateralTo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.TradeInputParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"forceClosePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.Result\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"forceSettleBoard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoardAndStrikeDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"strikeIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structOptionMarket.OptionBoard\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.Strike[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoardStrikes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"strikeIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiveBoards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_liveBoards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumLiveBoards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numLiveBoards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getOptionBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"strikeIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structOptionMarket.OptionBoard\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptionMarketParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"getSettlementParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikeToBaseReturned\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"getStrike\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.Strike\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"getStrikeAndBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.Strike\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iv\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"strikeIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structOptionMarket.OptionBoard\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"getStrikeAndExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractSynthetixAdapter\",\"name\":\"_synthetixAdapter\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityPool\",\"name\":\"_liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"_optionPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"_greekCache\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"_shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"_optionToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_baseAsset\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardBeneficiary\",\"type\":\"address\"}],\"name\":\"liquidatePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"setCollateralTo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.TradeInputParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"openPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.Result\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"}],\"name\":\"setBoardBaseIv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setBoardFrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"_optionMarketParams\",\"type\":\"tuple\"}],\"name\":\"setOptionMarketParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"}],\"name\":\"setStrikeSkew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"settleExpiredBoard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetBoardAndStrikeDetails is a free data retrieval call binding the contract method 0x1f18a342.
//
// Solidity: function getBoardAndStrikeDetails(uint256 boardId) view returns((uint256,uint256,uint256,bool,uint256[]), (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[], uint256[], uint256)
func (_Lyra *LyraCaller) GetBoardAndStrikeDetails(opts *bind.CallOpts, boardId *big.Int) (OptionMarketOptionBoard, []OptionMarketStrike, []*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getBoardAndStrikeDetails", boardId)

	if err != nil {
		return *new(OptionMarketOptionBoard), *new([]OptionMarketStrike), *new([]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketOptionBoard)).(*OptionMarketOptionBoard)
	out1 := *abi.ConvertType(out[1], new([]OptionMarketStrike)).(*[]OptionMarketStrike)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetBoardAndStrikeDetails is a free data retrieval call binding the contract method 0x1f18a342.
//
// Solidity: function getBoardAndStrikeDetails(uint256 boardId) view returns((uint256,uint256,uint256,bool,uint256[]), (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[], uint256[], uint256)
func (_Lyra *LyraSession) GetBoardAndStrikeDetails(boardId *big.Int) (OptionMarketOptionBoard, []OptionMarketStrike, []*big.Int, *big.Int, error) {
	return _Lyra.Contract.GetBoardAndStrikeDetails(&_Lyra.CallOpts, boardId)
}

// GetBoardAndStrikeDetails is a free data retrieval call binding the contract method 0x1f18a342.
//
// Solidity: function getBoardAndStrikeDetails(uint256 boardId) view returns((uint256,uint256,uint256,bool,uint256[]), (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[], uint256[], uint256)
func (_Lyra *LyraCallerSession) GetBoardAndStrikeDetails(boardId *big.Int) (OptionMarketOptionBoard, []OptionMarketStrike, []*big.Int, *big.Int, error) {
	return _Lyra.Contract.GetBoardAndStrikeDetails(&_Lyra.CallOpts, boardId)
}

// GetBoardStrikes is a free data retrieval call binding the contract method 0x5d9d310c.
//
// Solidity: function getBoardStrikes(uint256 boardId) view returns(uint256[] strikeIds)
func (_Lyra *LyraCaller) GetBoardStrikes(opts *bind.CallOpts, boardId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getBoardStrikes", boardId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBoardStrikes is a free data retrieval call binding the contract method 0x5d9d310c.
//
// Solidity: function getBoardStrikes(uint256 boardId) view returns(uint256[] strikeIds)
func (_Lyra *LyraSession) GetBoardStrikes(boardId *big.Int) ([]*big.Int, error) {
	return _Lyra.Contract.GetBoardStrikes(&_Lyra.CallOpts, boardId)
}

// GetBoardStrikes is a free data retrieval call binding the contract method 0x5d9d310c.
//
// Solidity: function getBoardStrikes(uint256 boardId) view returns(uint256[] strikeIds)
func (_Lyra *LyraCallerSession) GetBoardStrikes(boardId *big.Int) ([]*big.Int, error) {
	return _Lyra.Contract.GetBoardStrikes(&_Lyra.CallOpts, boardId)
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

// GetNumLiveBoards is a free data retrieval call binding the contract method 0xeed6601a.
//
// Solidity: function getNumLiveBoards() view returns(uint256 numLiveBoards)
func (_Lyra *LyraCaller) GetNumLiveBoards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getNumLiveBoards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumLiveBoards is a free data retrieval call binding the contract method 0xeed6601a.
//
// Solidity: function getNumLiveBoards() view returns(uint256 numLiveBoards)
func (_Lyra *LyraSession) GetNumLiveBoards() (*big.Int, error) {
	return _Lyra.Contract.GetNumLiveBoards(&_Lyra.CallOpts)
}

// GetNumLiveBoards is a free data retrieval call binding the contract method 0xeed6601a.
//
// Solidity: function getNumLiveBoards() view returns(uint256 numLiveBoards)
func (_Lyra *LyraCallerSession) GetNumLiveBoards() (*big.Int, error) {
	return _Lyra.Contract.GetNumLiveBoards(&_Lyra.CallOpts)
}

// GetOptionBoard is a free data retrieval call binding the contract method 0x16a54f50.
//
// Solidity: function getOptionBoard(uint256 boardId) view returns((uint256,uint256,uint256,bool,uint256[]))
func (_Lyra *LyraCaller) GetOptionBoard(opts *bind.CallOpts, boardId *big.Int) (OptionMarketOptionBoard, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getOptionBoard", boardId)

	if err != nil {
		return *new(OptionMarketOptionBoard), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketOptionBoard)).(*OptionMarketOptionBoard)

	return out0, err

}

// GetOptionBoard is a free data retrieval call binding the contract method 0x16a54f50.
//
// Solidity: function getOptionBoard(uint256 boardId) view returns((uint256,uint256,uint256,bool,uint256[]))
func (_Lyra *LyraSession) GetOptionBoard(boardId *big.Int) (OptionMarketOptionBoard, error) {
	return _Lyra.Contract.GetOptionBoard(&_Lyra.CallOpts, boardId)
}

// GetOptionBoard is a free data retrieval call binding the contract method 0x16a54f50.
//
// Solidity: function getOptionBoard(uint256 boardId) view returns((uint256,uint256,uint256,bool,uint256[]))
func (_Lyra *LyraCallerSession) GetOptionBoard(boardId *big.Int) (OptionMarketOptionBoard, error) {
	return _Lyra.Contract.GetOptionBoard(&_Lyra.CallOpts, boardId)
}

// GetOptionMarketParams is a free data retrieval call binding the contract method 0x3105dd9c.
//
// Solidity: function getOptionMarketParams() view returns((uint256,address,uint256,uint256))
func (_Lyra *LyraCaller) GetOptionMarketParams(opts *bind.CallOpts) (OptionMarketOptionMarketParameters, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getOptionMarketParams")

	if err != nil {
		return *new(OptionMarketOptionMarketParameters), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketOptionMarketParameters)).(*OptionMarketOptionMarketParameters)

	return out0, err

}

// GetOptionMarketParams is a free data retrieval call binding the contract method 0x3105dd9c.
//
// Solidity: function getOptionMarketParams() view returns((uint256,address,uint256,uint256))
func (_Lyra *LyraSession) GetOptionMarketParams() (OptionMarketOptionMarketParameters, error) {
	return _Lyra.Contract.GetOptionMarketParams(&_Lyra.CallOpts)
}

// GetOptionMarketParams is a free data retrieval call binding the contract method 0x3105dd9c.
//
// Solidity: function getOptionMarketParams() view returns((uint256,address,uint256,uint256))
func (_Lyra *LyraCallerSession) GetOptionMarketParams() (OptionMarketOptionMarketParameters, error) {
	return _Lyra.Contract.GetOptionMarketParams(&_Lyra.CallOpts)
}

// GetSettlementParameters is a free data retrieval call binding the contract method 0x1fdb6cbd.
//
// Solidity: function getSettlementParameters(uint256 strikeId) view returns(uint256 strikePrice, uint256 priceAtExpiry, uint256 strikeToBaseReturned)
func (_Lyra *LyraCaller) GetSettlementParameters(opts *bind.CallOpts, strikeId *big.Int) (struct {
	StrikePrice          *big.Int
	PriceAtExpiry        *big.Int
	StrikeToBaseReturned *big.Int
}, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getSettlementParameters", strikeId)

	outstruct := new(struct {
		StrikePrice          *big.Int
		PriceAtExpiry        *big.Int
		StrikeToBaseReturned *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StrikePrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PriceAtExpiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StrikeToBaseReturned = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSettlementParameters is a free data retrieval call binding the contract method 0x1fdb6cbd.
//
// Solidity: function getSettlementParameters(uint256 strikeId) view returns(uint256 strikePrice, uint256 priceAtExpiry, uint256 strikeToBaseReturned)
func (_Lyra *LyraSession) GetSettlementParameters(strikeId *big.Int) (struct {
	StrikePrice          *big.Int
	PriceAtExpiry        *big.Int
	StrikeToBaseReturned *big.Int
}, error) {
	return _Lyra.Contract.GetSettlementParameters(&_Lyra.CallOpts, strikeId)
}

// GetSettlementParameters is a free data retrieval call binding the contract method 0x1fdb6cbd.
//
// Solidity: function getSettlementParameters(uint256 strikeId) view returns(uint256 strikePrice, uint256 priceAtExpiry, uint256 strikeToBaseReturned)
func (_Lyra *LyraCallerSession) GetSettlementParameters(strikeId *big.Int) (struct {
	StrikePrice          *big.Int
	PriceAtExpiry        *big.Int
	StrikeToBaseReturned *big.Int
}, error) {
	return _Lyra.Contract.GetSettlementParameters(&_Lyra.CallOpts, strikeId)
}

// GetStrike is a free data retrieval call binding the contract method 0xa6063c05.
//
// Solidity: function getStrike(uint256 strikeId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyra *LyraCaller) GetStrike(opts *bind.CallOpts, strikeId *big.Int) (OptionMarketStrike, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getStrike", strikeId)

	if err != nil {
		return *new(OptionMarketStrike), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketStrike)).(*OptionMarketStrike)

	return out0, err

}

// GetStrike is a free data retrieval call binding the contract method 0xa6063c05.
//
// Solidity: function getStrike(uint256 strikeId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyra *LyraSession) GetStrike(strikeId *big.Int) (OptionMarketStrike, error) {
	return _Lyra.Contract.GetStrike(&_Lyra.CallOpts, strikeId)
}

// GetStrike is a free data retrieval call binding the contract method 0xa6063c05.
//
// Solidity: function getStrike(uint256 strikeId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Lyra *LyraCallerSession) GetStrike(strikeId *big.Int) (OptionMarketStrike, error) {
	return _Lyra.Contract.GetStrike(&_Lyra.CallOpts, strikeId)
}

// GetStrikeAndBoard is a free data retrieval call binding the contract method 0xc4c4a0d0.
//
// Solidity: function getStrikeAndBoard(uint256 strikeId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256), (uint256,uint256,uint256,bool,uint256[]))
func (_Lyra *LyraCaller) GetStrikeAndBoard(opts *bind.CallOpts, strikeId *big.Int) (OptionMarketStrike, OptionMarketOptionBoard, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getStrikeAndBoard", strikeId)

	if err != nil {
		return *new(OptionMarketStrike), *new(OptionMarketOptionBoard), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketStrike)).(*OptionMarketStrike)
	out1 := *abi.ConvertType(out[1], new(OptionMarketOptionBoard)).(*OptionMarketOptionBoard)

	return out0, out1, err

}

// GetStrikeAndBoard is a free data retrieval call binding the contract method 0xc4c4a0d0.
//
// Solidity: function getStrikeAndBoard(uint256 strikeId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256), (uint256,uint256,uint256,bool,uint256[]))
func (_Lyra *LyraSession) GetStrikeAndBoard(strikeId *big.Int) (OptionMarketStrike, OptionMarketOptionBoard, error) {
	return _Lyra.Contract.GetStrikeAndBoard(&_Lyra.CallOpts, strikeId)
}

// GetStrikeAndBoard is a free data retrieval call binding the contract method 0xc4c4a0d0.
//
// Solidity: function getStrikeAndBoard(uint256 strikeId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256), (uint256,uint256,uint256,bool,uint256[]))
func (_Lyra *LyraCallerSession) GetStrikeAndBoard(strikeId *big.Int) (OptionMarketStrike, OptionMarketOptionBoard, error) {
	return _Lyra.Contract.GetStrikeAndBoard(&_Lyra.CallOpts, strikeId)
}

// GetStrikeAndExpiry is a free data retrieval call binding the contract method 0xcf6bcba0.
//
// Solidity: function getStrikeAndExpiry(uint256 strikeId) view returns(uint256 strikePrice, uint256 expiry)
func (_Lyra *LyraCaller) GetStrikeAndExpiry(opts *bind.CallOpts, strikeId *big.Int) (struct {
	StrikePrice *big.Int
	Expiry      *big.Int
}, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "getStrikeAndExpiry", strikeId)

	outstruct := new(struct {
		StrikePrice *big.Int
		Expiry      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StrikePrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStrikeAndExpiry is a free data retrieval call binding the contract method 0xcf6bcba0.
//
// Solidity: function getStrikeAndExpiry(uint256 strikeId) view returns(uint256 strikePrice, uint256 expiry)
func (_Lyra *LyraSession) GetStrikeAndExpiry(strikeId *big.Int) (struct {
	StrikePrice *big.Int
	Expiry      *big.Int
}, error) {
	return _Lyra.Contract.GetStrikeAndExpiry(&_Lyra.CallOpts, strikeId)
}

// GetStrikeAndExpiry is a free data retrieval call binding the contract method 0xcf6bcba0.
//
// Solidity: function getStrikeAndExpiry(uint256 strikeId) view returns(uint256 strikePrice, uint256 expiry)
func (_Lyra *LyraCallerSession) GetStrikeAndExpiry(strikeId *big.Int) (struct {
	StrikePrice *big.Int
	Expiry      *big.Int
}, error) {
	return _Lyra.Contract.GetStrikeAndExpiry(&_Lyra.CallOpts, strikeId)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Lyra *LyraCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Lyra *LyraSession) NominatedOwner() (common.Address, error) {
	return _Lyra.Contract.NominatedOwner(&_Lyra.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Lyra *LyraCallerSession) NominatedOwner() (common.Address, error) {
	return _Lyra.Contract.NominatedOwner(&_Lyra.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lyra *LyraCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lyra.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lyra *LyraSession) Owner() (common.Address, error) {
	return _Lyra.Contract.Owner(&_Lyra.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lyra *LyraCallerSession) Owner() (common.Address, error) {
	return _Lyra.Contract.Owner(&_Lyra.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Lyra *LyraTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Lyra *LyraSession) AcceptOwnership() (*types.Transaction, error) {
	return _Lyra.Contract.AcceptOwnership(&_Lyra.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Lyra *LyraTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Lyra.Contract.AcceptOwnership(&_Lyra.TransactOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xa8f35adf.
//
// Solidity: function addCollateral(uint256 positionId, uint256 amountCollateral) returns()
func (_Lyra *LyraTransactor) AddCollateral(opts *bind.TransactOpts, positionId *big.Int, amountCollateral *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "addCollateral", positionId, amountCollateral)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xa8f35adf.
//
// Solidity: function addCollateral(uint256 positionId, uint256 amountCollateral) returns()
func (_Lyra *LyraSession) AddCollateral(positionId *big.Int, amountCollateral *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.AddCollateral(&_Lyra.TransactOpts, positionId, amountCollateral)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xa8f35adf.
//
// Solidity: function addCollateral(uint256 positionId, uint256 amountCollateral) returns()
func (_Lyra *LyraTransactorSession) AddCollateral(positionId *big.Int, amountCollateral *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.AddCollateral(&_Lyra.TransactOpts, positionId, amountCollateral)
}

// AddStrikeToBoard is a paid mutator transaction binding the contract method 0x05c8954a.
//
// Solidity: function addStrikeToBoard(uint256 boardId, uint256 strikePrice, uint256 skew) returns()
func (_Lyra *LyraTransactor) AddStrikeToBoard(opts *bind.TransactOpts, boardId *big.Int, strikePrice *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "addStrikeToBoard", boardId, strikePrice, skew)
}

// AddStrikeToBoard is a paid mutator transaction binding the contract method 0x05c8954a.
//
// Solidity: function addStrikeToBoard(uint256 boardId, uint256 strikePrice, uint256 skew) returns()
func (_Lyra *LyraSession) AddStrikeToBoard(boardId *big.Int, strikePrice *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.AddStrikeToBoard(&_Lyra.TransactOpts, boardId, strikePrice, skew)
}

// AddStrikeToBoard is a paid mutator transaction binding the contract method 0x05c8954a.
//
// Solidity: function addStrikeToBoard(uint256 boardId, uint256 strikePrice, uint256 skew) returns()
func (_Lyra *LyraTransactorSession) AddStrikeToBoard(boardId *big.Int, strikePrice *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.AddStrikeToBoard(&_Lyra.TransactOpts, boardId, strikePrice, skew)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x92b4632c.
//
// Solidity: function closePosition((uint256,uint256,uint256,uint8,uint256,uint256,uint256,uint256) params) returns((uint256,uint256,uint256) result)
func (_Lyra *LyraTransactor) ClosePosition(opts *bind.TransactOpts, params OptionMarketTradeInputParameters) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "closePosition", params)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x92b4632c.
//
// Solidity: function closePosition((uint256,uint256,uint256,uint8,uint256,uint256,uint256,uint256) params) returns((uint256,uint256,uint256) result)
func (_Lyra *LyraSession) ClosePosition(params OptionMarketTradeInputParameters) (*types.Transaction, error) {
	return _Lyra.Contract.ClosePosition(&_Lyra.TransactOpts, params)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x92b4632c.
//
// Solidity: function closePosition((uint256,uint256,uint256,uint8,uint256,uint256,uint256,uint256) params) returns((uint256,uint256,uint256) result)
func (_Lyra *LyraTransactorSession) ClosePosition(params OptionMarketTradeInputParameters) (*types.Transaction, error) {
	return _Lyra.Contract.ClosePosition(&_Lyra.TransactOpts, params)
}

// CreateOptionBoard is a paid mutator transaction binding the contract method 0x2dd0776b.
//
// Solidity: function createOptionBoard(uint256 expiry, uint256 baseIV, uint256[] strikePrices, uint256[] skews, bool frozen) returns(uint256 boardId)
func (_Lyra *LyraTransactor) CreateOptionBoard(opts *bind.TransactOpts, expiry *big.Int, baseIV *big.Int, strikePrices []*big.Int, skews []*big.Int, frozen bool) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "createOptionBoard", expiry, baseIV, strikePrices, skews, frozen)
}

// CreateOptionBoard is a paid mutator transaction binding the contract method 0x2dd0776b.
//
// Solidity: function createOptionBoard(uint256 expiry, uint256 baseIV, uint256[] strikePrices, uint256[] skews, bool frozen) returns(uint256 boardId)
func (_Lyra *LyraSession) CreateOptionBoard(expiry *big.Int, baseIV *big.Int, strikePrices []*big.Int, skews []*big.Int, frozen bool) (*types.Transaction, error) {
	return _Lyra.Contract.CreateOptionBoard(&_Lyra.TransactOpts, expiry, baseIV, strikePrices, skews, frozen)
}

// CreateOptionBoard is a paid mutator transaction binding the contract method 0x2dd0776b.
//
// Solidity: function createOptionBoard(uint256 expiry, uint256 baseIV, uint256[] strikePrices, uint256[] skews, bool frozen) returns(uint256 boardId)
func (_Lyra *LyraTransactorSession) CreateOptionBoard(expiry *big.Int, baseIV *big.Int, strikePrices []*big.Int, skews []*big.Int, frozen bool) (*types.Transaction, error) {
	return _Lyra.Contract.CreateOptionBoard(&_Lyra.TransactOpts, expiry, baseIV, strikePrices, skews, frozen)
}

// ForceClosePosition is a paid mutator transaction binding the contract method 0xe4e83e3d.
//
// Solidity: function forceClosePosition((uint256,uint256,uint256,uint8,uint256,uint256,uint256,uint256) params) returns((uint256,uint256,uint256) result)
func (_Lyra *LyraTransactor) ForceClosePosition(opts *bind.TransactOpts, params OptionMarketTradeInputParameters) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "forceClosePosition", params)
}

// ForceClosePosition is a paid mutator transaction binding the contract method 0xe4e83e3d.
//
// Solidity: function forceClosePosition((uint256,uint256,uint256,uint8,uint256,uint256,uint256,uint256) params) returns((uint256,uint256,uint256) result)
func (_Lyra *LyraSession) ForceClosePosition(params OptionMarketTradeInputParameters) (*types.Transaction, error) {
	return _Lyra.Contract.ForceClosePosition(&_Lyra.TransactOpts, params)
}

// ForceClosePosition is a paid mutator transaction binding the contract method 0xe4e83e3d.
//
// Solidity: function forceClosePosition((uint256,uint256,uint256,uint8,uint256,uint256,uint256,uint256) params) returns((uint256,uint256,uint256) result)
func (_Lyra *LyraTransactorSession) ForceClosePosition(params OptionMarketTradeInputParameters) (*types.Transaction, error) {
	return _Lyra.Contract.ForceClosePosition(&_Lyra.TransactOpts, params)
}

// ForceSettleBoard is a paid mutator transaction binding the contract method 0x90e32fba.
//
// Solidity: function forceSettleBoard(uint256 boardId) returns()
func (_Lyra *LyraTransactor) ForceSettleBoard(opts *bind.TransactOpts, boardId *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "forceSettleBoard", boardId)
}

// ForceSettleBoard is a paid mutator transaction binding the contract method 0x90e32fba.
//
// Solidity: function forceSettleBoard(uint256 boardId) returns()
func (_Lyra *LyraSession) ForceSettleBoard(boardId *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.ForceSettleBoard(&_Lyra.TransactOpts, boardId)
}

// ForceSettleBoard is a paid mutator transaction binding the contract method 0x90e32fba.
//
// Solidity: function forceSettleBoard(uint256 boardId) returns()
func (_Lyra *LyraTransactorSession) ForceSettleBoard(boardId *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.ForceSettleBoard(&_Lyra.TransactOpts, boardId)
}

// Init is a paid mutator transaction binding the contract method 0x525240c0.
//
// Solidity: function init(address _synthetixAdapter, address _liquidityPool, address _optionPricer, address _greekCache, address _shortCollateral, address _optionToken, address _quoteAsset, address _baseAsset) returns()
func (_Lyra *LyraTransactor) Init(opts *bind.TransactOpts, _synthetixAdapter common.Address, _liquidityPool common.Address, _optionPricer common.Address, _greekCache common.Address, _shortCollateral common.Address, _optionToken common.Address, _quoteAsset common.Address, _baseAsset common.Address) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "init", _synthetixAdapter, _liquidityPool, _optionPricer, _greekCache, _shortCollateral, _optionToken, _quoteAsset, _baseAsset)
}

// Init is a paid mutator transaction binding the contract method 0x525240c0.
//
// Solidity: function init(address _synthetixAdapter, address _liquidityPool, address _optionPricer, address _greekCache, address _shortCollateral, address _optionToken, address _quoteAsset, address _baseAsset) returns()
func (_Lyra *LyraSession) Init(_synthetixAdapter common.Address, _liquidityPool common.Address, _optionPricer common.Address, _greekCache common.Address, _shortCollateral common.Address, _optionToken common.Address, _quoteAsset common.Address, _baseAsset common.Address) (*types.Transaction, error) {
	return _Lyra.Contract.Init(&_Lyra.TransactOpts, _synthetixAdapter, _liquidityPool, _optionPricer, _greekCache, _shortCollateral, _optionToken, _quoteAsset, _baseAsset)
}

// Init is a paid mutator transaction binding the contract method 0x525240c0.
//
// Solidity: function init(address _synthetixAdapter, address _liquidityPool, address _optionPricer, address _greekCache, address _shortCollateral, address _optionToken, address _quoteAsset, address _baseAsset) returns()
func (_Lyra *LyraTransactorSession) Init(_synthetixAdapter common.Address, _liquidityPool common.Address, _optionPricer common.Address, _greekCache common.Address, _shortCollateral common.Address, _optionToken common.Address, _quoteAsset common.Address, _baseAsset common.Address) (*types.Transaction, error) {
	return _Lyra.Contract.Init(&_Lyra.TransactOpts, _synthetixAdapter, _liquidityPool, _optionPricer, _greekCache, _shortCollateral, _optionToken, _quoteAsset, _baseAsset)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0x5f036de1.
//
// Solidity: function liquidatePosition(uint256 positionId, address rewardBeneficiary) returns()
func (_Lyra *LyraTransactor) LiquidatePosition(opts *bind.TransactOpts, positionId *big.Int, rewardBeneficiary common.Address) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "liquidatePosition", positionId, rewardBeneficiary)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0x5f036de1.
//
// Solidity: function liquidatePosition(uint256 positionId, address rewardBeneficiary) returns()
func (_Lyra *LyraSession) LiquidatePosition(positionId *big.Int, rewardBeneficiary common.Address) (*types.Transaction, error) {
	return _Lyra.Contract.LiquidatePosition(&_Lyra.TransactOpts, positionId, rewardBeneficiary)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0x5f036de1.
//
// Solidity: function liquidatePosition(uint256 positionId, address rewardBeneficiary) returns()
func (_Lyra *LyraTransactorSession) LiquidatePosition(positionId *big.Int, rewardBeneficiary common.Address) (*types.Transaction, error) {
	return _Lyra.Contract.LiquidatePosition(&_Lyra.TransactOpts, positionId, rewardBeneficiary)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Lyra *LyraTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Lyra *LyraSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Lyra.Contract.NominateNewOwner(&_Lyra.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Lyra *LyraTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Lyra.Contract.NominateNewOwner(&_Lyra.TransactOpts, _owner)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x0037f2bc.
//
// Solidity: function openPosition((uint256,uint256,uint256,uint8,uint256,uint256,uint256,uint256) params) returns((uint256,uint256,uint256) result)
func (_Lyra *LyraTransactor) OpenPosition(opts *bind.TransactOpts, params OptionMarketTradeInputParameters) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "openPosition", params)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x0037f2bc.
//
// Solidity: function openPosition((uint256,uint256,uint256,uint8,uint256,uint256,uint256,uint256) params) returns((uint256,uint256,uint256) result)
func (_Lyra *LyraSession) OpenPosition(params OptionMarketTradeInputParameters) (*types.Transaction, error) {
	return _Lyra.Contract.OpenPosition(&_Lyra.TransactOpts, params)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x0037f2bc.
//
// Solidity: function openPosition((uint256,uint256,uint256,uint8,uint256,uint256,uint256,uint256) params) returns((uint256,uint256,uint256) result)
func (_Lyra *LyraTransactorSession) OpenPosition(params OptionMarketTradeInputParameters) (*types.Transaction, error) {
	return _Lyra.Contract.OpenPosition(&_Lyra.TransactOpts, params)
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

// SetOptionMarketParams is a paid mutator transaction binding the contract method 0x1227e500.
//
// Solidity: function setOptionMarketParams((uint256,address,uint256,uint256) _optionMarketParams) returns()
func (_Lyra *LyraTransactor) SetOptionMarketParams(opts *bind.TransactOpts, _optionMarketParams OptionMarketOptionMarketParameters) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "setOptionMarketParams", _optionMarketParams)
}

// SetOptionMarketParams is a paid mutator transaction binding the contract method 0x1227e500.
//
// Solidity: function setOptionMarketParams((uint256,address,uint256,uint256) _optionMarketParams) returns()
func (_Lyra *LyraSession) SetOptionMarketParams(_optionMarketParams OptionMarketOptionMarketParameters) (*types.Transaction, error) {
	return _Lyra.Contract.SetOptionMarketParams(&_Lyra.TransactOpts, _optionMarketParams)
}

// SetOptionMarketParams is a paid mutator transaction binding the contract method 0x1227e500.
//
// Solidity: function setOptionMarketParams((uint256,address,uint256,uint256) _optionMarketParams) returns()
func (_Lyra *LyraTransactorSession) SetOptionMarketParams(_optionMarketParams OptionMarketOptionMarketParameters) (*types.Transaction, error) {
	return _Lyra.Contract.SetOptionMarketParams(&_Lyra.TransactOpts, _optionMarketParams)
}

// SetStrikeSkew is a paid mutator transaction binding the contract method 0xe7a5897b.
//
// Solidity: function setStrikeSkew(uint256 strikeId, uint256 skew) returns()
func (_Lyra *LyraTransactor) SetStrikeSkew(opts *bind.TransactOpts, strikeId *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "setStrikeSkew", strikeId, skew)
}

// SetStrikeSkew is a paid mutator transaction binding the contract method 0xe7a5897b.
//
// Solidity: function setStrikeSkew(uint256 strikeId, uint256 skew) returns()
func (_Lyra *LyraSession) SetStrikeSkew(strikeId *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.SetStrikeSkew(&_Lyra.TransactOpts, strikeId, skew)
}

// SetStrikeSkew is a paid mutator transaction binding the contract method 0xe7a5897b.
//
// Solidity: function setStrikeSkew(uint256 strikeId, uint256 skew) returns()
func (_Lyra *LyraTransactorSession) SetStrikeSkew(strikeId *big.Int, skew *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.SetStrikeSkew(&_Lyra.TransactOpts, strikeId, skew)
}

// SettleExpiredBoard is a paid mutator transaction binding the contract method 0x7c1de425.
//
// Solidity: function settleExpiredBoard(uint256 boardId) returns()
func (_Lyra *LyraTransactor) SettleExpiredBoard(opts *bind.TransactOpts, boardId *big.Int) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "settleExpiredBoard", boardId)
}

// SettleExpiredBoard is a paid mutator transaction binding the contract method 0x7c1de425.
//
// Solidity: function settleExpiredBoard(uint256 boardId) returns()
func (_Lyra *LyraSession) SettleExpiredBoard(boardId *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.SettleExpiredBoard(&_Lyra.TransactOpts, boardId)
}

// SettleExpiredBoard is a paid mutator transaction binding the contract method 0x7c1de425.
//
// Solidity: function settleExpiredBoard(uint256 boardId) returns()
func (_Lyra *LyraTransactorSession) SettleExpiredBoard(boardId *big.Int) (*types.Transaction, error) {
	return _Lyra.Contract.SettleExpiredBoard(&_Lyra.TransactOpts, boardId)
}

// SmClaim is a paid mutator transaction binding the contract method 0xebc20866.
//
// Solidity: function smClaim() returns()
func (_Lyra *LyraTransactor) SmClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lyra.contract.Transact(opts, "smClaim")
}

// SmClaim is a paid mutator transaction binding the contract method 0xebc20866.
//
// Solidity: function smClaim() returns()
func (_Lyra *LyraSession) SmClaim() (*types.Transaction, error) {
	return _Lyra.Contract.SmClaim(&_Lyra.TransactOpts)
}

// SmClaim is a paid mutator transaction binding the contract method 0xebc20866.
//
// Solidity: function smClaim() returns()
func (_Lyra *LyraTransactorSession) SmClaim() (*types.Transaction, error) {
	return _Lyra.Contract.SmClaim(&_Lyra.TransactOpts)
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
	Frozen  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBoardCreated is a free log retrieval operation binding the contract event 0xcb94f87a9b05d8957b230ed8ef82ef8ef24fb02e374b4b1300c402ccb3b8868e.
//
// Solidity: event BoardCreated(uint256 indexed boardId, uint256 expiry, uint256 baseIv, bool frozen)
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

// WatchBoardCreated is a free log subscription operation binding the contract event 0xcb94f87a9b05d8957b230ed8ef82ef8ef24fb02e374b4b1300c402ccb3b8868e.
//
// Solidity: event BoardCreated(uint256 indexed boardId, uint256 expiry, uint256 baseIv, bool frozen)
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

// ParseBoardCreated is a log parse operation binding the contract event 0xcb94f87a9b05d8957b230ed8ef82ef8ef24fb02e374b4b1300c402ccb3b8868e.
//
// Solidity: event BoardCreated(uint256 indexed boardId, uint256 expiry, uint256 baseIv, bool frozen)
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

// LyraBoardSettledIterator is returned from FilterBoardSettled and is used to iterate over the raw logs and unpacked data for BoardSettled events raised by the Lyra contract.
type LyraBoardSettledIterator struct {
	Event *LyraBoardSettled // Event containing the contract specifics and raw log

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
func (it *LyraBoardSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraBoardSettled)
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
		it.Event = new(LyraBoardSettled)
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
func (it *LyraBoardSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraBoardSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraBoardSettled represents a BoardSettled event raised by the Lyra contract.
type LyraBoardSettled struct {
	BoardId                      *big.Int
	SpotPriceAtExpiry            *big.Int
	TotalUserLongProfitQuote     *big.Int
	TotalBoardLongCallCollateral *big.Int
	TotalBoardLongPutCollateral  *big.Int
	TotalAMMShortCallProfitBase  *big.Int
	TotalAMMShortCallProfitQuote *big.Int
	TotalAMMShortPutProfitQuote  *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterBoardSettled is a free log retrieval operation binding the contract event 0xdb2372d9a91045fa42195170e972ab572d182278c275fbb277fed93ac5ddda82.
//
// Solidity: event BoardSettled(uint256 indexed boardId, uint256 spotPriceAtExpiry, uint256 totalUserLongProfitQuote, uint256 totalBoardLongCallCollateral, uint256 totalBoardLongPutCollateral, uint256 totalAMMShortCallProfitBase, uint256 totalAMMShortCallProfitQuote, uint256 totalAMMShortPutProfitQuote)
func (_Lyra *LyraFilterer) FilterBoardSettled(opts *bind.FilterOpts, boardId []*big.Int) (*LyraBoardSettledIterator, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "BoardSettled", boardIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraBoardSettledIterator{contract: _Lyra.contract, event: "BoardSettled", logs: logs, sub: sub}, nil
}

// WatchBoardSettled is a free log subscription operation binding the contract event 0xdb2372d9a91045fa42195170e972ab572d182278c275fbb277fed93ac5ddda82.
//
// Solidity: event BoardSettled(uint256 indexed boardId, uint256 spotPriceAtExpiry, uint256 totalUserLongProfitQuote, uint256 totalBoardLongCallCollateral, uint256 totalBoardLongPutCollateral, uint256 totalAMMShortCallProfitBase, uint256 totalAMMShortCallProfitQuote, uint256 totalAMMShortPutProfitQuote)
func (_Lyra *LyraFilterer) WatchBoardSettled(opts *bind.WatchOpts, sink chan<- *LyraBoardSettled, boardId []*big.Int) (event.Subscription, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "BoardSettled", boardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraBoardSettled)
				if err := _Lyra.contract.UnpackLog(event, "BoardSettled", log); err != nil {
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

// ParseBoardSettled is a log parse operation binding the contract event 0xdb2372d9a91045fa42195170e972ab572d182278c275fbb277fed93ac5ddda82.
//
// Solidity: event BoardSettled(uint256 indexed boardId, uint256 spotPriceAtExpiry, uint256 totalUserLongProfitQuote, uint256 totalBoardLongCallCollateral, uint256 totalBoardLongPutCollateral, uint256 totalAMMShortCallProfitBase, uint256 totalAMMShortCallProfitQuote, uint256 totalAMMShortPutProfitQuote)
func (_Lyra *LyraFilterer) ParseBoardSettled(log types.Log) (*LyraBoardSettled, error) {
	event := new(LyraBoardSettled)
	if err := _Lyra.contract.UnpackLog(event, "BoardSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraOptionMarketParamsSetIterator is returned from FilterOptionMarketParamsSet and is used to iterate over the raw logs and unpacked data for OptionMarketParamsSet events raised by the Lyra contract.
type LyraOptionMarketParamsSetIterator struct {
	Event *LyraOptionMarketParamsSet // Event containing the contract specifics and raw log

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
func (it *LyraOptionMarketParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraOptionMarketParamsSet)
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
		it.Event = new(LyraOptionMarketParamsSet)
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
func (it *LyraOptionMarketParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraOptionMarketParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraOptionMarketParamsSet represents a OptionMarketParamsSet event raised by the Lyra contract.
type LyraOptionMarketParamsSet struct {
	OptionMarketParams OptionMarketOptionMarketParameters
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOptionMarketParamsSet is a free log retrieval operation binding the contract event 0x19854f162a92eaba887a94b15691fb68f39b5ffc2a48f1eccd9ebac965432d80.
//
// Solidity: event OptionMarketParamsSet((uint256,address,uint256,uint256) optionMarketParams)
func (_Lyra *LyraFilterer) FilterOptionMarketParamsSet(opts *bind.FilterOpts) (*LyraOptionMarketParamsSetIterator, error) {

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "OptionMarketParamsSet")
	if err != nil {
		return nil, err
	}
	return &LyraOptionMarketParamsSetIterator{contract: _Lyra.contract, event: "OptionMarketParamsSet", logs: logs, sub: sub}, nil
}

// WatchOptionMarketParamsSet is a free log subscription operation binding the contract event 0x19854f162a92eaba887a94b15691fb68f39b5ffc2a48f1eccd9ebac965432d80.
//
// Solidity: event OptionMarketParamsSet((uint256,address,uint256,uint256) optionMarketParams)
func (_Lyra *LyraFilterer) WatchOptionMarketParamsSet(opts *bind.WatchOpts, sink chan<- *LyraOptionMarketParamsSet) (event.Subscription, error) {

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "OptionMarketParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraOptionMarketParamsSet)
				if err := _Lyra.contract.UnpackLog(event, "OptionMarketParamsSet", log); err != nil {
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

// ParseOptionMarketParamsSet is a log parse operation binding the contract event 0x19854f162a92eaba887a94b15691fb68f39b5ffc2a48f1eccd9ebac965432d80.
//
// Solidity: event OptionMarketParamsSet((uint256,address,uint256,uint256) optionMarketParams)
func (_Lyra *LyraFilterer) ParseOptionMarketParamsSet(log types.Log) (*LyraOptionMarketParamsSet, error) {
	event := new(LyraOptionMarketParamsSet)
	if err := _Lyra.contract.UnpackLog(event, "OptionMarketParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Lyra contract.
type LyraOwnerChangedIterator struct {
	Event *LyraOwnerChanged // Event containing the contract specifics and raw log

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
func (it *LyraOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraOwnerChanged)
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
		it.Event = new(LyraOwnerChanged)
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
func (it *LyraOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraOwnerChanged represents a OwnerChanged event raised by the Lyra contract.
type LyraOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Lyra *LyraFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*LyraOwnerChangedIterator, error) {

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &LyraOwnerChangedIterator{contract: _Lyra.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Lyra *LyraFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *LyraOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraOwnerChanged)
				if err := _Lyra.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Lyra *LyraFilterer) ParseOwnerChanged(log types.Log) (*LyraOwnerChanged, error) {
	event := new(LyraOwnerChanged)
	if err := _Lyra.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Lyra contract.
type LyraOwnerNominatedIterator struct {
	Event *LyraOwnerNominated // Event containing the contract specifics and raw log

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
func (it *LyraOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraOwnerNominated)
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
		it.Event = new(LyraOwnerNominated)
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
func (it *LyraOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraOwnerNominated represents a OwnerNominated event raised by the Lyra contract.
type LyraOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Lyra *LyraFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*LyraOwnerNominatedIterator, error) {

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &LyraOwnerNominatedIterator{contract: _Lyra.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Lyra *LyraFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *LyraOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraOwnerNominated)
				if err := _Lyra.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Lyra *LyraFilterer) ParseOwnerNominated(log types.Log) (*LyraOwnerNominated, error) {
	event := new(LyraOwnerNominated)
	if err := _Lyra.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraSMClaimedIterator is returned from FilterSMClaimed and is used to iterate over the raw logs and unpacked data for SMClaimed events raised by the Lyra contract.
type LyraSMClaimedIterator struct {
	Event *LyraSMClaimed // Event containing the contract specifics and raw log

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
func (it *LyraSMClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraSMClaimed)
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
		it.Event = new(LyraSMClaimed)
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
func (it *LyraSMClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraSMClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraSMClaimed represents a SMClaimed event raised by the Lyra contract.
type LyraSMClaimed struct {
	SecurityModule common.Address
	QuoteAmount    *big.Int
	BaseAmount     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSMClaimed is a free log retrieval operation binding the contract event 0x37b3d08ef7ea83f3a646520b4d8d248a92ffba0f65943f3bc3a8d58dd149e94b.
//
// Solidity: event SMClaimed(address securityModule, uint256 quoteAmount, uint256 baseAmount)
func (_Lyra *LyraFilterer) FilterSMClaimed(opts *bind.FilterOpts) (*LyraSMClaimedIterator, error) {

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "SMClaimed")
	if err != nil {
		return nil, err
	}
	return &LyraSMClaimedIterator{contract: _Lyra.contract, event: "SMClaimed", logs: logs, sub: sub}, nil
}

// WatchSMClaimed is a free log subscription operation binding the contract event 0x37b3d08ef7ea83f3a646520b4d8d248a92ffba0f65943f3bc3a8d58dd149e94b.
//
// Solidity: event SMClaimed(address securityModule, uint256 quoteAmount, uint256 baseAmount)
func (_Lyra *LyraFilterer) WatchSMClaimed(opts *bind.WatchOpts, sink chan<- *LyraSMClaimed) (event.Subscription, error) {

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "SMClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraSMClaimed)
				if err := _Lyra.contract.UnpackLog(event, "SMClaimed", log); err != nil {
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

// ParseSMClaimed is a log parse operation binding the contract event 0x37b3d08ef7ea83f3a646520b4d8d248a92ffba0f65943f3bc3a8d58dd149e94b.
//
// Solidity: event SMClaimed(address securityModule, uint256 quoteAmount, uint256 baseAmount)
func (_Lyra *LyraFilterer) ParseSMClaimed(log types.Log) (*LyraSMClaimed, error) {
	event := new(LyraSMClaimed)
	if err := _Lyra.contract.UnpackLog(event, "SMClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraStrikeAddedIterator is returned from FilterStrikeAdded and is used to iterate over the raw logs and unpacked data for StrikeAdded events raised by the Lyra contract.
type LyraStrikeAddedIterator struct {
	Event *LyraStrikeAdded // Event containing the contract specifics and raw log

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
func (it *LyraStrikeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraStrikeAdded)
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
		it.Event = new(LyraStrikeAdded)
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
func (it *LyraStrikeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraStrikeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraStrikeAdded represents a StrikeAdded event raised by the Lyra contract.
type LyraStrikeAdded struct {
	BoardId     *big.Int
	StrikeId    *big.Int
	StrikePrice *big.Int
	Skew        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrikeAdded is a free log retrieval operation binding the contract event 0x8dba15d5569538a96f1404ee74da67080a64e02c5cfbadf4b8908f90320178a1.
//
// Solidity: event StrikeAdded(uint256 indexed boardId, uint256 indexed strikeId, uint256 strikePrice, uint256 skew)
func (_Lyra *LyraFilterer) FilterStrikeAdded(opts *bind.FilterOpts, boardId []*big.Int, strikeId []*big.Int) (*LyraStrikeAddedIterator, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}
	var strikeIdRule []interface{}
	for _, strikeIdItem := range strikeId {
		strikeIdRule = append(strikeIdRule, strikeIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "StrikeAdded", boardIdRule, strikeIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraStrikeAddedIterator{contract: _Lyra.contract, event: "StrikeAdded", logs: logs, sub: sub}, nil
}

// WatchStrikeAdded is a free log subscription operation binding the contract event 0x8dba15d5569538a96f1404ee74da67080a64e02c5cfbadf4b8908f90320178a1.
//
// Solidity: event StrikeAdded(uint256 indexed boardId, uint256 indexed strikeId, uint256 strikePrice, uint256 skew)
func (_Lyra *LyraFilterer) WatchStrikeAdded(opts *bind.WatchOpts, sink chan<- *LyraStrikeAdded, boardId []*big.Int, strikeId []*big.Int) (event.Subscription, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}
	var strikeIdRule []interface{}
	for _, strikeIdItem := range strikeId {
		strikeIdRule = append(strikeIdRule, strikeIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "StrikeAdded", boardIdRule, strikeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraStrikeAdded)
				if err := _Lyra.contract.UnpackLog(event, "StrikeAdded", log); err != nil {
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

// ParseStrikeAdded is a log parse operation binding the contract event 0x8dba15d5569538a96f1404ee74da67080a64e02c5cfbadf4b8908f90320178a1.
//
// Solidity: event StrikeAdded(uint256 indexed boardId, uint256 indexed strikeId, uint256 strikePrice, uint256 skew)
func (_Lyra *LyraFilterer) ParseStrikeAdded(log types.Log) (*LyraStrikeAdded, error) {
	event := new(LyraStrikeAdded)
	if err := _Lyra.contract.UnpackLog(event, "StrikeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraStrikeSkewSetIterator is returned from FilterStrikeSkewSet and is used to iterate over the raw logs and unpacked data for StrikeSkewSet events raised by the Lyra contract.
type LyraStrikeSkewSetIterator struct {
	Event *LyraStrikeSkewSet // Event containing the contract specifics and raw log

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
func (it *LyraStrikeSkewSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraStrikeSkewSet)
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
		it.Event = new(LyraStrikeSkewSet)
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
func (it *LyraStrikeSkewSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraStrikeSkewSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraStrikeSkewSet represents a StrikeSkewSet event raised by the Lyra contract.
type LyraStrikeSkewSet struct {
	StrikeId *big.Int
	Skew     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrikeSkewSet is a free log retrieval operation binding the contract event 0xe494e8ea1592ce5eff63e093b9aa6d5208883099ea0a8734e64a518226fd443e.
//
// Solidity: event StrikeSkewSet(uint256 indexed strikeId, uint256 skew)
func (_Lyra *LyraFilterer) FilterStrikeSkewSet(opts *bind.FilterOpts, strikeId []*big.Int) (*LyraStrikeSkewSetIterator, error) {

	var strikeIdRule []interface{}
	for _, strikeIdItem := range strikeId {
		strikeIdRule = append(strikeIdRule, strikeIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "StrikeSkewSet", strikeIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraStrikeSkewSetIterator{contract: _Lyra.contract, event: "StrikeSkewSet", logs: logs, sub: sub}, nil
}

// WatchStrikeSkewSet is a free log subscription operation binding the contract event 0xe494e8ea1592ce5eff63e093b9aa6d5208883099ea0a8734e64a518226fd443e.
//
// Solidity: event StrikeSkewSet(uint256 indexed strikeId, uint256 skew)
func (_Lyra *LyraFilterer) WatchStrikeSkewSet(opts *bind.WatchOpts, sink chan<- *LyraStrikeSkewSet, strikeId []*big.Int) (event.Subscription, error) {

	var strikeIdRule []interface{}
	for _, strikeIdItem := range strikeId {
		strikeIdRule = append(strikeIdRule, strikeIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "StrikeSkewSet", strikeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraStrikeSkewSet)
				if err := _Lyra.contract.UnpackLog(event, "StrikeSkewSet", log); err != nil {
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

// ParseStrikeSkewSet is a log parse operation binding the contract event 0xe494e8ea1592ce5eff63e093b9aa6d5208883099ea0a8734e64a518226fd443e.
//
// Solidity: event StrikeSkewSet(uint256 indexed strikeId, uint256 skew)
func (_Lyra *LyraFilterer) ParseStrikeSkewSet(log types.Log) (*LyraStrikeSkewSet, error) {
	event := new(LyraStrikeSkewSet)
	if err := _Lyra.contract.UnpackLog(event, "StrikeSkewSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LyraTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Lyra contract.
type LyraTradeIterator struct {
	Event *LyraTrade // Event containing the contract specifics and raw log

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
func (it *LyraTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LyraTrade)
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
		it.Event = new(LyraTrade)
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
func (it *LyraTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LyraTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LyraTrade represents a Trade event raised by the Lyra contract.
type LyraTrade struct {
	Trader       common.Address
	StrikeId     *big.Int
	PositionId   *big.Int
	Trade        OptionMarketTradeEventData
	TradeResults []OptionMarketPricerTradeResult
	Liquidation  OptionMarketLiquidationEventData
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0xb587f4a964cfbcea150e428bfd686fc9305b9c66a45f0daa0ff6f3e99e612d4d.
//
// Solidity: event Trade(address indexed trader, uint256 indexed strikeId, uint256 indexed positionId, (uint256,uint256,uint8,uint8,uint256,uint256,bool,uint256,uint256,uint256) trade, (uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256)[] tradeResults, (address,address,uint256,uint256,uint256,uint256,uint256,uint256) liquidation, uint256 timestamp)
func (_Lyra *LyraFilterer) FilterTrade(opts *bind.FilterOpts, trader []common.Address, strikeId []*big.Int, positionId []*big.Int) (*LyraTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var strikeIdRule []interface{}
	for _, strikeIdItem := range strikeId {
		strikeIdRule = append(strikeIdRule, strikeIdItem)
	}
	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _Lyra.contract.FilterLogs(opts, "Trade", traderRule, strikeIdRule, positionIdRule)
	if err != nil {
		return nil, err
	}
	return &LyraTradeIterator{contract: _Lyra.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xb587f4a964cfbcea150e428bfd686fc9305b9c66a45f0daa0ff6f3e99e612d4d.
//
// Solidity: event Trade(address indexed trader, uint256 indexed strikeId, uint256 indexed positionId, (uint256,uint256,uint8,uint8,uint256,uint256,bool,uint256,uint256,uint256) trade, (uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256)[] tradeResults, (address,address,uint256,uint256,uint256,uint256,uint256,uint256) liquidation, uint256 timestamp)
func (_Lyra *LyraFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *LyraTrade, trader []common.Address, strikeId []*big.Int, positionId []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var strikeIdRule []interface{}
	for _, strikeIdItem := range strikeId {
		strikeIdRule = append(strikeIdRule, strikeIdItem)
	}
	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _Lyra.contract.WatchLogs(opts, "Trade", traderRule, strikeIdRule, positionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LyraTrade)
				if err := _Lyra.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0xb587f4a964cfbcea150e428bfd686fc9305b9c66a45f0daa0ff6f3e99e612d4d.
//
// Solidity: event Trade(address indexed trader, uint256 indexed strikeId, uint256 indexed positionId, (uint256,uint256,uint8,uint8,uint256,uint256,bool,uint256,uint256,uint256) trade, (uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256)[] tradeResults, (address,address,uint256,uint256,uint256,uint256,uint256,uint256) liquidation, uint256 timestamp)
func (_Lyra *LyraFilterer) ParseTrade(log types.Log) (*LyraTrade, error) {
	event := new(LyraTrade)
	if err := _Lyra.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
