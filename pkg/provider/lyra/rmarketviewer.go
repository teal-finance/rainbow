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

// LiquidityPoolCircuitBreakerParameters is an auto generated low-level Go binding around an user-defined struct.
type LiquidityPoolCircuitBreakerParameters struct {
	LiquidityCBThreshold        *big.Int
	LiquidityCBTimeout          *big.Int
	IvVarianceCBThreshold       *big.Int
	SkewVarianceCBThreshold     *big.Int
	IvVarianceCBTimeout         *big.Int
	SkewVarianceCBTimeout       *big.Int
	BoardSettlementCBTimeout    *big.Int
	ContractAdjustmentCBTimeout *big.Int
}

// LiquidityPoolLiquidity is an auto generated low-level Go binding around an user-defined struct.
type LiquidityPoolLiquidity struct {
	FreeLiquidity           *big.Int
	BurnableLiquidity       *big.Int
	ReservedCollatLiquidity *big.Int
	PendingDeltaLiquidity   *big.Int
	UsedDeltaLiquidity      *big.Int
	NAV                     *big.Int
	LongScaleFactor         *big.Int
}

// LiquidityPoolLiquidityPoolParameters is an auto generated low-level Go binding around an user-defined struct.
type LiquidityPoolLiquidityPoolParameters struct {
	MinDepositWithdraw         *big.Int
	DepositDelay               *big.Int
	WithdrawalDelay            *big.Int
	WithdrawalFee              *big.Int
	GuardianMultisig           common.Address
	GuardianDelay              *big.Int
	AdjustmentNetScalingFactor *big.Int
	CallCollatScalingFactor    *big.Int
	PutCollatScalingFactor     *big.Int
}

// OptionGreekCacheForceCloseParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionGreekCacheForceCloseParameters struct {
	IvGWAVPeriod                *big.Int
	SkewGWAVPeriod              *big.Int
	ShortVolShock               *big.Int
	ShortPostCutoffVolShock     *big.Int
	LongVolShock                *big.Int
	LongPostCutoffVolShock      *big.Int
	LiquidateVolShock           *big.Int
	LiquidatePostCutoffVolShock *big.Int
	ShortSpotMin                *big.Int
	LiquidateSpotMin            *big.Int
}

// OptionGreekCacheGreekCacheParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionGreekCacheGreekCacheParameters struct {
	MaxStrikesPerBoard             *big.Int
	AcceptableSpotPricePercentMove *big.Int
	StaleUpdateDuration            *big.Int
	VarianceIvGWAVPeriod           *big.Int
	VarianceSkewGWAVPeriod         *big.Int
	OptionValueIvGWAVPeriod        *big.Int
	OptionValueSkewGWAVPeriod      *big.Int
	GwavSkewFloor                  *big.Int
	GwavSkewCap                    *big.Int
}

// OptionGreekCacheMinCollateralParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionGreekCacheMinCollateralParameters struct {
	MinStaticQuoteCollateral *big.Int
	MinStaticBaseCollateral  *big.Int
	ShockVolA                *big.Int
	ShockVolPointA           *big.Int
	ShockVolB                *big.Int
	ShockVolPointB           *big.Int
	CallSpotPriceShock       *big.Int
	PutSpotPriceShock        *big.Int
}

// OptionGreekCacheNetGreeks is an auto generated low-level Go binding around an user-defined struct.
type OptionGreekCacheNetGreeks struct {
	NetDelta       *big.Int
	NetStdVega     *big.Int
	NetOptionValue *big.Int
}

// OptionGreekCacheStrikeGreeks is an auto generated low-level Go binding around an user-defined struct.
type OptionGreekCacheStrikeGreeks struct {
	CallDelta *big.Int
	PutDelta  *big.Int
	StdVega   *big.Int
	CallPrice *big.Int
	PutPrice  *big.Int
}

// OptionMarketOptionMarketParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketOptionMarketParameters struct {
	MaxBoardExpiry          *big.Int
	SecurityModule          common.Address
	FeePortionReserved      *big.Int
	StaticBaseSettlementFee *big.Int
}

// OptionMarketPricerPricingParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketPricerPricingParameters struct {
	OptionPriceFeeCoefficient *big.Int
	OptionPriceFee1xPoint     *big.Int
	OptionPriceFee2xPoint     *big.Int
	SpotPriceFeeCoefficient   *big.Int
	SpotPriceFee1xPoint       *big.Int
	SpotPriceFee2xPoint       *big.Int
	VegaFeeCoefficient        *big.Int
	StandardSize              *big.Int
	SkewAdjustmentFactor      *big.Int
}

// OptionMarketPricerTradeLimitParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketPricerTradeLimitParameters struct {
	MinDelta           *big.Int
	MinForceCloseDelta *big.Int
	TradingCutoff      *big.Int
	MinBaseIV          *big.Int
	MaxBaseIV          *big.Int
	MinSkew            *big.Int
	MaxSkew            *big.Int
	MinVol             *big.Int
	MaxVol             *big.Int
	AbsMinSkew         *big.Int
	AbsMaxSkew         *big.Int
	CapSkewsToAbs      bool
}

// OptionMarketPricerVarianceFeeParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketPricerVarianceFeeParameters struct {
	DefaultVarianceFeeCoefficient    *big.Int
	ForceCloseVarianceFeeCoefficient *big.Int
	SkewAdjustmentCoefficient        *big.Int
	ReferenceSkew                    *big.Int
	MinimumStaticSkewAdjustment      *big.Int
	VegaCoefficient                  *big.Int
	MinimumStaticVega                *big.Int
	IvVarianceCoefficient            *big.Int
	MinimumStaticIvVariance          *big.Int
}

// OptionMarketViewerBoardView is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerBoardView struct {
	Market           common.Address
	BoardId          *big.Int
	Expiry           *big.Int
	BaseIv           *big.Int
	PriceAtExpiry    *big.Int
	IsPaused         bool
	VarianceGwavIv   *big.Int
	ForceCloseGwavIv *big.Int
	LongScaleFactor  *big.Int
	NetGreeks        OptionGreekCacheNetGreeks
	Strikes          []OptionMarketViewerStrikeView
}

// OptionMarketViewerLiquidityBalance is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerLiquidityBalance struct {
	QuoteAsset            common.Address
	QuoteBalance          *big.Int
	QuoteSymbol           string
	QuoteDepositAllowance *big.Int
	LiquidityToken        common.Address
	LiquidityBalance      *big.Int
}

// OptionMarketViewerMarketOptionPositions is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerMarketOptionPositions struct {
	Market    common.Address
	Positions []OptionTokenOptionPosition
}

// OptionMarketViewerMarketParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerMarketParameters struct {
	OptionMarketParams  OptionMarketOptionMarketParameters
	LpParams            LiquidityPoolLiquidityPoolParameters
	CbParams            LiquidityPoolCircuitBreakerParameters
	GreekCacheParams    OptionGreekCacheGreekCacheParameters
	ForceCloseParams    OptionGreekCacheForceCloseParameters
	MinCollatParams     OptionGreekCacheMinCollateralParameters
	PricingParams       OptionMarketPricerPricingParameters
	TradeLimitParams    OptionMarketPricerTradeLimitParameters
	VarianceFeeParams   OptionMarketPricerVarianceFeeParameters
	PartialCollatParams OptionTokenPartialCollateralParameters
}

// OptionMarketViewerMarketView is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerMarketView struct {
	IsPaused         bool
	SpotPrice        *big.Int
	QuoteSymbol      string
	QuoteDecimals    *big.Int
	BaseSymbol       string
	BaseDecimals     *big.Int
	Liquidity        LiquidityPoolLiquidity
	MarketAddresses  OptionMarketViewerOptionMarketAddresses
	MarketParameters OptionMarketViewerMarketParameters
	GlobalNetGreeks  OptionGreekCacheNetGreeks
	LiveBoards       []OptionMarketViewerBoardView
}

// OptionMarketViewerMarketsView is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerMarketsView struct {
	IsPaused bool
	Markets  []OptionMarketViewerMarketView
}

// OptionMarketViewerOptionMarketAddresses is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerOptionMarketAddresses struct {
	LiquidityPool      common.Address
	LiquidityToken     common.Address
	GreekCache         common.Address
	OptionMarket       common.Address
	OptionMarketPricer common.Address
	OptionToken        common.Address
	ShortCollateral    common.Address
	PoolHedger         common.Address
	QuoteAsset         common.Address
	BaseAsset          common.Address
}

// OptionMarketViewerStrikeView is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerStrikeView struct {
	StrikeId                   *big.Int
	BoardId                    *big.Int
	StrikePrice                *big.Int
	Skew                       *big.Int
	ForceCloseSkew             *big.Int
	CachedGreeks               OptionGreekCacheStrikeGreeks
	BaseReturnedRatio          *big.Int
	LongCallOpenInterest       *big.Int
	LongPutOpenInterest        *big.Int
	ShortCallBaseOpenInterest  *big.Int
	ShortCallQuoteOpenInterest *big.Int
	ShortPutOpenInterest       *big.Int
}

// OptionTokenOptionPosition is an auto generated low-level Go binding around an user-defined struct.
type OptionTokenOptionPosition struct {
	PositionId *big.Int
	StrikeId   *big.Int
	OptionType uint8
	Amount     *big.Int
	Collateral *big.Int
	State      uint8
}

// OptionTokenPartialCollateralParameters is an auto generated low-level Go binding around an user-defined struct.
type OptionTokenPartialCollateralParameters struct {
	PenaltyRatio       *big.Int
	LiquidatorFeeRatio *big.Int
	SmFeeRatio         *big.Int
	MinLiquidationFee  *big.Int
}

// MarketviewerMetaData contains all meta data concerning the Marketviewer contract.
var MarketviewerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"}],\"name\":\"OnlyNominatedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"RemovingInvalidMarket\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"market\",\"type\":\"tuple\"}],\"name\":\"MarketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"newMarketAddresses\",\"type\":\"tuple\"}],\"name\":\"addMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeAdapter\",\"outputs\":[{\"internalType\":\"contractBaseExchangeAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoardForBase\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"getBoardForStrikeId\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLiquidityBalances\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteBalance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"quoteSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quoteDepositAllowance\",\"type\":\"uint256\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidityBalance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.LiquidityBalance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"getLiveBoards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"marketBoards\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"quoteSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quoteDecimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"baseDecimals\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"marketAddresses\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minDepositWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guardianMultisig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"guardianDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adjustmentNetScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callCollatScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putCollatScalingFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.LiquidityPoolParameters\",\"name\":\"lpParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"liquidityCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardSettlementCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractAdjustmentCBTimeout\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.CircuitBreakerParameters\",\"name\":\"cbParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxStrikesPerBoard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptableSpotPricePercentMove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staleUpdateDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewCap\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.GreekCacheParameters\",\"name\":\"greekCacheParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ivGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatePostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSpotMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateSpotMin\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.ForceCloseParameters\",\"name\":\"forceCloseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minStaticQuoteCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStaticBaseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callSpotPriceShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putSpotPriceShock\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.MinCollateralParameters\",\"name\":\"minCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.PricingParameters\",\"name\":\"pricingParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minForceCloseDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradingCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMinSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMaxSkew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"capSkewsToAbs\",\"type\":\"bool\"}],\"internalType\":\"structOptionMarketPricer.TradeLimitParameters\",\"name\":\"tradeLimitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"defaultVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referenceSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticSkewAdjustment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticIvVariance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeParameters\",\"name\":\"varianceFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"penaltyRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidationFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionToken.PartialCollateralParameters\",\"name\":\"partialCollatParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketParameters\",\"name\":\"marketParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"globalNetGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"liveBoards\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketView\",\"name\":\"marketView\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"}],\"name\":\"getMarketForBase\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"quoteSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quoteDecimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"baseDecimals\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"marketAddresses\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minDepositWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guardianMultisig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"guardianDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adjustmentNetScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callCollatScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putCollatScalingFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.LiquidityPoolParameters\",\"name\":\"lpParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"liquidityCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardSettlementCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractAdjustmentCBTimeout\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.CircuitBreakerParameters\",\"name\":\"cbParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxStrikesPerBoard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptableSpotPricePercentMove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staleUpdateDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewCap\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.GreekCacheParameters\",\"name\":\"greekCacheParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ivGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatePostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSpotMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateSpotMin\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.ForceCloseParameters\",\"name\":\"forceCloseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minStaticQuoteCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStaticBaseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callSpotPriceShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putSpotPriceShock\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.MinCollateralParameters\",\"name\":\"minCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.PricingParameters\",\"name\":\"pricingParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minForceCloseDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradingCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMinSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMaxSkew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"capSkewsToAbs\",\"type\":\"bool\"}],\"internalType\":\"structOptionMarketPricer.TradeLimitParameters\",\"name\":\"tradeLimitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"defaultVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referenceSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticSkewAdjustment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticIvVariance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeParameters\",\"name\":\"varianceFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"penaltyRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidationFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionToken.PartialCollateralParameters\",\"name\":\"partialCollatParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketParameters\",\"name\":\"marketParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"globalNetGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"liveBoards\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketView\",\"name\":\"marketView\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"name\":\"getMarkets\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"quoteSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quoteDecimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"baseDecimals\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"marketAddresses\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minDepositWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guardianMultisig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"guardianDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adjustmentNetScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callCollatScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putCollatScalingFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.LiquidityPoolParameters\",\"name\":\"lpParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"liquidityCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardSettlementCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractAdjustmentCBTimeout\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.CircuitBreakerParameters\",\"name\":\"cbParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxStrikesPerBoard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptableSpotPricePercentMove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staleUpdateDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewCap\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.GreekCacheParameters\",\"name\":\"greekCacheParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ivGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatePostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSpotMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateSpotMin\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.ForceCloseParameters\",\"name\":\"forceCloseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minStaticQuoteCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStaticBaseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callSpotPriceShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putSpotPriceShock\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.MinCollateralParameters\",\"name\":\"minCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.PricingParameters\",\"name\":\"pricingParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minForceCloseDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradingCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMinSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMaxSkew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"capSkewsToAbs\",\"type\":\"bool\"}],\"internalType\":\"structOptionMarketPricer.TradeLimitParameters\",\"name\":\"tradeLimitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"defaultVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referenceSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticSkewAdjustment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticIvVariance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeParameters\",\"name\":\"varianceFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"penaltyRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidationFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionToken.PartialCollateralParameters\",\"name\":\"partialCollatParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketParameters\",\"name\":\"marketParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"globalNetGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"liveBoards\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketView[]\",\"name\":\"markets\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketsView\",\"name\":\"marketsView\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getOwnerPositions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionToken.PositionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structOptionToken.OptionPosition[]\",\"name\":\"positions\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketOptionPositions[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseExchangeAdapter\",\"name\":\"_exchangeAdapter\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketAddresses\",\"outputs\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"optionMarkets\",\"outputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"removeMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketviewerABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketviewerMetaData.ABI instead.
var MarketviewerABI = MarketviewerMetaData.ABI

// Marketviewer is an auto generated Go binding around an Ethereum contract.
type Marketviewer struct {
	MarketviewerCaller     // Read-only binding to the contract
	MarketviewerTransactor // Write-only binding to the contract
	MarketviewerFilterer   // Log filterer for contract events
}

// MarketviewerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketviewerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketviewerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketviewerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketviewerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketviewerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketviewerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketviewerSession struct {
	Contract     *Marketviewer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketviewerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketviewerCallerSession struct {
	Contract *MarketviewerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MarketviewerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketviewerTransactorSession struct {
	Contract     *MarketviewerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MarketviewerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketviewerRaw struct {
	Contract *Marketviewer // Generic contract binding to access the raw methods on
}

// MarketviewerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketviewerCallerRaw struct {
	Contract *MarketviewerCaller // Generic read-only contract binding to access the raw methods on
}

// MarketviewerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketviewerTransactorRaw struct {
	Contract *MarketviewerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketviewer creates a new instance of Marketviewer, bound to a specific deployed contract.
func NewMarketviewer(address common.Address, backend bind.ContractBackend) (*Marketviewer, error) {
	contract, err := bindMarketviewer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketviewer{MarketviewerCaller: MarketviewerCaller{contract: contract}, MarketviewerTransactor: MarketviewerTransactor{contract: contract}, MarketviewerFilterer: MarketviewerFilterer{contract: contract}}, nil
}

// NewMarketviewerCaller creates a new read-only instance of Marketviewer, bound to a specific deployed contract.
func NewMarketviewerCaller(address common.Address, caller bind.ContractCaller) (*MarketviewerCaller, error) {
	contract, err := bindMarketviewer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketviewerCaller{contract: contract}, nil
}

// NewMarketviewerTransactor creates a new write-only instance of Marketviewer, bound to a specific deployed contract.
func NewMarketviewerTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketviewerTransactor, error) {
	contract, err := bindMarketviewer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketviewerTransactor{contract: contract}, nil
}

// NewMarketviewerFilterer creates a new log filterer instance of Marketviewer, bound to a specific deployed contract.
func NewMarketviewerFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketviewerFilterer, error) {
	contract, err := bindMarketviewer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketviewerFilterer{contract: contract}, nil
}

// bindMarketviewer binds a generic wrapper to an already deployed contract.
func bindMarketviewer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketviewerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketviewer *MarketviewerRaw) Call(opts *bind.CallOpts, result *[]any, method string, params ...any) error {
	return _Marketviewer.Contract.MarketviewerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketviewer *MarketviewerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketviewer.Contract.MarketviewerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketviewer *MarketviewerRaw) Transact(opts *bind.TransactOpts, method string, params ...any) (*types.Transaction, error) {
	return _Marketviewer.Contract.MarketviewerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketviewer *MarketviewerCallerRaw) Call(opts *bind.CallOpts, result *[]any, method string, params ...any) error {
	return _Marketviewer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketviewer *MarketviewerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketviewer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketviewer *MarketviewerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...any) (*types.Transaction, error) {
	return _Marketviewer.Contract.contract.Transact(opts, method, params...)
}

// ExchangeAdapter is a free data retrieval call binding the contract method 0x2cbcda25.
//
// Solidity: function exchangeAdapter() view returns(address)
func (_Marketviewer *MarketviewerCaller) ExchangeAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "exchangeAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeAdapter is a free data retrieval call binding the contract method 0x2cbcda25.
//
// Solidity: function exchangeAdapter() view returns(address)
func (_Marketviewer *MarketviewerSession) ExchangeAdapter() (common.Address, error) {
	return _Marketviewer.Contract.ExchangeAdapter(&_Marketviewer.CallOpts)
}

// ExchangeAdapter is a free data retrieval call binding the contract method 0x2cbcda25.
//
// Solidity: function exchangeAdapter() view returns(address)
func (_Marketviewer *MarketviewerCallerSession) ExchangeAdapter() (common.Address, error) {
	return _Marketviewer.Contract.ExchangeAdapter(&_Marketviewer.CallOpts)
}

// GetBoard is a free data retrieval call binding the contract method 0x5d033f7f.
//
// Solidity: function getBoard(address market, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_Marketviewer *MarketviewerCaller) GetBoard(opts *bind.CallOpts, market common.Address, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getBoard", market, boardId)

	if err != nil {
		return *new(OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerBoardView)).(*OptionMarketViewerBoardView)

	return out0, err

}

// GetBoard is a free data retrieval call binding the contract method 0x5d033f7f.
//
// Solidity: function getBoard(address market, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_Marketviewer *MarketviewerSession) GetBoard(market common.Address, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _Marketviewer.Contract.GetBoard(&_Marketviewer.CallOpts, market, boardId)
}

// GetBoard is a free data retrieval call binding the contract method 0x5d033f7f.
//
// Solidity: function getBoard(address market, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_Marketviewer *MarketviewerCallerSession) GetBoard(market common.Address, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _Marketviewer.Contract.GetBoard(&_Marketviewer.CallOpts, market, boardId)
}

// GetBoardForBase is a free data retrieval call binding the contract method 0x173b7f77.
//
// Solidity: function getBoardForBase(string baseSymbol, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_Marketviewer *MarketviewerCaller) GetBoardForBase(opts *bind.CallOpts, baseSymbol string, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getBoardForBase", baseSymbol, boardId)

	if err != nil {
		return *new(OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerBoardView)).(*OptionMarketViewerBoardView)

	return out0, err

}

// GetBoardForBase is a free data retrieval call binding the contract method 0x173b7f77.
//
// Solidity: function getBoardForBase(string baseSymbol, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_Marketviewer *MarketviewerSession) GetBoardForBase(baseSymbol string, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _Marketviewer.Contract.GetBoardForBase(&_Marketviewer.CallOpts, baseSymbol, boardId)
}

// GetBoardForBase is a free data retrieval call binding the contract method 0x173b7f77.
//
// Solidity: function getBoardForBase(string baseSymbol, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_Marketviewer *MarketviewerCallerSession) GetBoardForBase(baseSymbol string, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _Marketviewer.Contract.GetBoardForBase(&_Marketviewer.CallOpts, baseSymbol, boardId)
}

// GetBoardForStrikeId is a free data retrieval call binding the contract method 0xac6e227b.
//
// Solidity: function getBoardForStrikeId(address market, uint256 strikeId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_Marketviewer *MarketviewerCaller) GetBoardForStrikeId(opts *bind.CallOpts, market common.Address, strikeId *big.Int) (OptionMarketViewerBoardView, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getBoardForStrikeId", market, strikeId)

	if err != nil {
		return *new(OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerBoardView)).(*OptionMarketViewerBoardView)

	return out0, err

}

// GetBoardForStrikeId is a free data retrieval call binding the contract method 0xac6e227b.
//
// Solidity: function getBoardForStrikeId(address market, uint256 strikeId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_Marketviewer *MarketviewerSession) GetBoardForStrikeId(market common.Address, strikeId *big.Int) (OptionMarketViewerBoardView, error) {
	return _Marketviewer.Contract.GetBoardForStrikeId(&_Marketviewer.CallOpts, market, strikeId)
}

// GetBoardForStrikeId is a free data retrieval call binding the contract method 0xac6e227b.
//
// Solidity: function getBoardForStrikeId(address market, uint256 strikeId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_Marketviewer *MarketviewerCallerSession) GetBoardForStrikeId(market common.Address, strikeId *big.Int) (OptionMarketViewerBoardView, error) {
	return _Marketviewer.Contract.GetBoardForStrikeId(&_Marketviewer.CallOpts, market, strikeId)
}

// GetLiquidityBalances is a free data retrieval call binding the contract method 0x4af60f1a.
//
// Solidity: function getLiquidityBalances(address account) view returns((address,uint256,string,uint256,address,uint256)[])
func (_Marketviewer *MarketviewerCaller) GetLiquidityBalances(opts *bind.CallOpts, account common.Address) ([]OptionMarketViewerLiquidityBalance, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getLiquidityBalances", account)

	if err != nil {
		return *new([]OptionMarketViewerLiquidityBalance), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerLiquidityBalance)).(*[]OptionMarketViewerLiquidityBalance)

	return out0, err

}

// GetLiquidityBalances is a free data retrieval call binding the contract method 0x4af60f1a.
//
// Solidity: function getLiquidityBalances(address account) view returns((address,uint256,string,uint256,address,uint256)[])
func (_Marketviewer *MarketviewerSession) GetLiquidityBalances(account common.Address) ([]OptionMarketViewerLiquidityBalance, error) {
	return _Marketviewer.Contract.GetLiquidityBalances(&_Marketviewer.CallOpts, account)
}

// GetLiquidityBalances is a free data retrieval call binding the contract method 0x4af60f1a.
//
// Solidity: function getLiquidityBalances(address account) view returns((address,uint256,string,uint256,address,uint256)[])
func (_Marketviewer *MarketviewerCallerSession) GetLiquidityBalances(account common.Address) ([]OptionMarketViewerLiquidityBalance, error) {
	return _Marketviewer.Contract.GetLiquidityBalances(&_Marketviewer.CallOpts, account)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0xb0862c0e.
//
// Solidity: function getLiveBoards(address market) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[] marketBoards)
func (_Marketviewer *MarketviewerCaller) GetLiveBoards(opts *bind.CallOpts, market common.Address) ([]OptionMarketViewerBoardView, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getLiveBoards", market)

	if err != nil {
		return *new([]OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerBoardView)).(*[]OptionMarketViewerBoardView)

	return out0, err

}

// GetLiveBoards is a free data retrieval call binding the contract method 0xb0862c0e.
//
// Solidity: function getLiveBoards(address market) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[] marketBoards)
func (_Marketviewer *MarketviewerSession) GetLiveBoards(market common.Address) ([]OptionMarketViewerBoardView, error) {
	return _Marketviewer.Contract.GetLiveBoards(&_Marketviewer.CallOpts, market)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0xb0862c0e.
//
// Solidity: function getLiveBoards(address market) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[] marketBoards)
func (_Marketviewer *MarketviewerCallerSession) GetLiveBoards(market common.Address) ([]OptionMarketViewerBoardView, error) {
	return _Marketviewer.Contract.GetLiveBoards(&_Marketviewer.CallOpts, market)
}

// GetMarket is a free data retrieval call binding the contract method 0xd4dfadbf.
//
// Solidity: function getMarket(address market) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_Marketviewer *MarketviewerCaller) GetMarket(opts *bind.CallOpts, market common.Address) (OptionMarketViewerMarketView, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getMarket", market)

	if err != nil {
		return *new(OptionMarketViewerMarketView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerMarketView)).(*OptionMarketViewerMarketView)

	return out0, err

}

// GetMarket is a free data retrieval call binding the contract method 0xd4dfadbf.
//
// Solidity: function getMarket(address market) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_Marketviewer *MarketviewerSession) GetMarket(market common.Address) (OptionMarketViewerMarketView, error) {
	return _Marketviewer.Contract.GetMarket(&_Marketviewer.CallOpts, market)
}

// GetMarket is a free data retrieval call binding the contract method 0xd4dfadbf.
//
// Solidity: function getMarket(address market) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_Marketviewer *MarketviewerCallerSession) GetMarket(market common.Address) (OptionMarketViewerMarketView, error) {
	return _Marketviewer.Contract.GetMarket(&_Marketviewer.CallOpts, market)
}

// GetMarketAddresses is a free data retrieval call binding the contract method 0x97ce0d31.
//
// Solidity: function getMarketAddresses() view returns((address,address,address,address,address,address,address,address,address,address)[])
func (_Marketviewer *MarketviewerCaller) GetMarketAddresses(opts *bind.CallOpts) ([]OptionMarketViewerOptionMarketAddresses, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getMarketAddresses")

	if err != nil {
		return *new([]OptionMarketViewerOptionMarketAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerOptionMarketAddresses)).(*[]OptionMarketViewerOptionMarketAddresses)

	return out0, err

}

// GetMarketAddresses is a free data retrieval call binding the contract method 0x97ce0d31.
//
// Solidity: function getMarketAddresses() view returns((address,address,address,address,address,address,address,address,address,address)[])
func (_Marketviewer *MarketviewerSession) GetMarketAddresses() ([]OptionMarketViewerOptionMarketAddresses, error) {
	return _Marketviewer.Contract.GetMarketAddresses(&_Marketviewer.CallOpts)
}

// GetMarketAddresses is a free data retrieval call binding the contract method 0x97ce0d31.
//
// Solidity: function getMarketAddresses() view returns((address,address,address,address,address,address,address,address,address,address)[])
func (_Marketviewer *MarketviewerCallerSession) GetMarketAddresses() ([]OptionMarketViewerOptionMarketAddresses, error) {
	return _Marketviewer.Contract.GetMarketAddresses(&_Marketviewer.CallOpts)
}

// GetMarketForBase is a free data retrieval call binding the contract method 0x96c121b5.
//
// Solidity: function getMarketForBase(string baseSymbol) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_Marketviewer *MarketviewerCaller) GetMarketForBase(opts *bind.CallOpts, baseSymbol string) (OptionMarketViewerMarketView, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getMarketForBase", baseSymbol)

	if err != nil {
		return *new(OptionMarketViewerMarketView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerMarketView)).(*OptionMarketViewerMarketView)

	return out0, err

}

// GetMarketForBase is a free data retrieval call binding the contract method 0x96c121b5.
//
// Solidity: function getMarketForBase(string baseSymbol) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_Marketviewer *MarketviewerSession) GetMarketForBase(baseSymbol string) (OptionMarketViewerMarketView, error) {
	return _Marketviewer.Contract.GetMarketForBase(&_Marketviewer.CallOpts, baseSymbol)
}

// GetMarketForBase is a free data retrieval call binding the contract method 0x96c121b5.
//
// Solidity: function getMarketForBase(string baseSymbol) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_Marketviewer *MarketviewerCallerSession) GetMarketForBase(baseSymbol string) (OptionMarketViewerMarketView, error) {
	return _Marketviewer.Contract.GetMarketForBase(&_Marketviewer.CallOpts, baseSymbol)
}

// GetMarkets is a free data retrieval call binding the contract method 0x1139e3f1.
//
// Solidity: function getMarkets(address[] markets) view returns((bool,(bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[])[]) marketsView)
func (_Marketviewer *MarketviewerCaller) GetMarkets(opts *bind.CallOpts, markets []common.Address) (OptionMarketViewerMarketsView, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getMarkets", markets)

	if err != nil {
		return *new(OptionMarketViewerMarketsView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerMarketsView)).(*OptionMarketViewerMarketsView)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0x1139e3f1.
//
// Solidity: function getMarkets(address[] markets) view returns((bool,(bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[])[]) marketsView)
func (_Marketviewer *MarketviewerSession) GetMarkets(markets []common.Address) (OptionMarketViewerMarketsView, error) {
	return _Marketviewer.Contract.GetMarkets(&_Marketviewer.CallOpts, markets)
}

// GetMarkets is a free data retrieval call binding the contract method 0x1139e3f1.
//
// Solidity: function getMarkets(address[] markets) view returns((bool,(bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[])[]) marketsView)
func (_Marketviewer *MarketviewerCallerSession) GetMarkets(markets []common.Address) (OptionMarketViewerMarketsView, error) {
	return _Marketviewer.Contract.GetMarkets(&_Marketviewer.CallOpts, markets)
}

// GetOwnerPositions is a free data retrieval call binding the contract method 0x845e1259.
//
// Solidity: function getOwnerPositions(address account) view returns((address,(uint256,uint256,uint8,uint256,uint256,uint8)[])[])
func (_Marketviewer *MarketviewerCaller) GetOwnerPositions(opts *bind.CallOpts, account common.Address) ([]OptionMarketViewerMarketOptionPositions, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "getOwnerPositions", account)

	if err != nil {
		return *new([]OptionMarketViewerMarketOptionPositions), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerMarketOptionPositions)).(*[]OptionMarketViewerMarketOptionPositions)

	return out0, err

}

// GetOwnerPositions is a free data retrieval call binding the contract method 0x845e1259.
//
// Solidity: function getOwnerPositions(address account) view returns((address,(uint256,uint256,uint8,uint256,uint256,uint8)[])[])
func (_Marketviewer *MarketviewerSession) GetOwnerPositions(account common.Address) ([]OptionMarketViewerMarketOptionPositions, error) {
	return _Marketviewer.Contract.GetOwnerPositions(&_Marketviewer.CallOpts, account)
}

// GetOwnerPositions is a free data retrieval call binding the contract method 0x845e1259.
//
// Solidity: function getOwnerPositions(address account) view returns((address,(uint256,uint256,uint8,uint256,uint256,uint8)[])[])
func (_Marketviewer *MarketviewerCallerSession) GetOwnerPositions(account common.Address) ([]OptionMarketViewerMarketOptionPositions, error) {
	return _Marketviewer.Contract.GetOwnerPositions(&_Marketviewer.CallOpts, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Marketviewer *MarketviewerCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Marketviewer *MarketviewerSession) Initialized() (bool, error) {
	return _Marketviewer.Contract.Initialized(&_Marketviewer.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Marketviewer *MarketviewerCallerSession) Initialized() (bool, error) {
	return _Marketviewer.Contract.Initialized(&_Marketviewer.CallOpts)
}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address shortCollateral, address poolHedger, address quoteAsset, address baseAsset)
func (_Marketviewer *MarketviewerCaller) MarketAddresses(opts *bind.CallOpts, arg0 common.Address) (struct {
	LiquidityPool      common.Address
	LiquidityToken     common.Address
	GreekCache         common.Address
	OptionMarket       common.Address
	OptionMarketPricer common.Address
	OptionToken        common.Address
	ShortCollateral    common.Address
	PoolHedger         common.Address
	QuoteAsset         common.Address
	BaseAsset          common.Address
}, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "marketAddresses", arg0)

	outstruct := new(struct {
		LiquidityPool      common.Address
		LiquidityToken     common.Address
		GreekCache         common.Address
		OptionMarket       common.Address
		OptionMarketPricer common.Address
		OptionToken        common.Address
		ShortCollateral    common.Address
		PoolHedger         common.Address
		QuoteAsset         common.Address
		BaseAsset          common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityPool = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LiquidityToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.GreekCache = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OptionMarket = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.OptionMarketPricer = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.OptionToken = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ShortCollateral = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.PoolHedger = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.QuoteAsset = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.BaseAsset = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address shortCollateral, address poolHedger, address quoteAsset, address baseAsset)
func (_Marketviewer *MarketviewerSession) MarketAddresses(arg0 common.Address) (struct {
	LiquidityPool      common.Address
	LiquidityToken     common.Address
	GreekCache         common.Address
	OptionMarket       common.Address
	OptionMarketPricer common.Address
	OptionToken        common.Address
	ShortCollateral    common.Address
	PoolHedger         common.Address
	QuoteAsset         common.Address
	BaseAsset          common.Address
}, error) {
	return _Marketviewer.Contract.MarketAddresses(&_Marketviewer.CallOpts, arg0)
}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address shortCollateral, address poolHedger, address quoteAsset, address baseAsset)
func (_Marketviewer *MarketviewerCallerSession) MarketAddresses(arg0 common.Address) (struct {
	LiquidityPool      common.Address
	LiquidityToken     common.Address
	GreekCache         common.Address
	OptionMarket       common.Address
	OptionMarketPricer common.Address
	OptionToken        common.Address
	ShortCollateral    common.Address
	PoolHedger         common.Address
	QuoteAsset         common.Address
	BaseAsset          common.Address
}, error) {
	return _Marketviewer.Contract.MarketAddresses(&_Marketviewer.CallOpts, arg0)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Marketviewer *MarketviewerCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Marketviewer *MarketviewerSession) NominatedOwner() (common.Address, error) {
	return _Marketviewer.Contract.NominatedOwner(&_Marketviewer.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Marketviewer *MarketviewerCallerSession) NominatedOwner() (common.Address, error) {
	return _Marketviewer.Contract.NominatedOwner(&_Marketviewer.CallOpts)
}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_Marketviewer *MarketviewerCaller) OptionMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "optionMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_Marketviewer *MarketviewerSession) OptionMarkets(arg0 *big.Int) (common.Address, error) {
	return _Marketviewer.Contract.OptionMarkets(&_Marketviewer.CallOpts, arg0)
}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_Marketviewer *MarketviewerCallerSession) OptionMarkets(arg0 *big.Int) (common.Address, error) {
	return _Marketviewer.Contract.OptionMarkets(&_Marketviewer.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketviewer *MarketviewerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Marketviewer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketviewer *MarketviewerSession) Owner() (common.Address, error) {
	return _Marketviewer.Contract.Owner(&_Marketviewer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketviewer *MarketviewerCallerSession) Owner() (common.Address, error) {
	return _Marketviewer.Contract.Owner(&_Marketviewer.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Marketviewer *MarketviewerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketviewer.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Marketviewer *MarketviewerSession) AcceptOwnership() (*types.Transaction, error) {
	return _Marketviewer.Contract.AcceptOwnership(&_Marketviewer.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Marketviewer *MarketviewerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Marketviewer.Contract.AcceptOwnership(&_Marketviewer.TransactOpts)
}

// AddMarket is a paid mutator transaction binding the contract method 0x7f68720f.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_Marketviewer *MarketviewerTransactor) AddMarket(opts *bind.TransactOpts, newMarketAddresses OptionMarketViewerOptionMarketAddresses) (*types.Transaction, error) {
	return _Marketviewer.contract.Transact(opts, "addMarket", newMarketAddresses)
}

// AddMarket is a paid mutator transaction binding the contract method 0x7f68720f.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_Marketviewer *MarketviewerSession) AddMarket(newMarketAddresses OptionMarketViewerOptionMarketAddresses) (*types.Transaction, error) {
	return _Marketviewer.Contract.AddMarket(&_Marketviewer.TransactOpts, newMarketAddresses)
}

// AddMarket is a paid mutator transaction binding the contract method 0x7f68720f.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_Marketviewer *MarketviewerTransactorSession) AddMarket(newMarketAddresses OptionMarketViewerOptionMarketAddresses) (*types.Transaction, error) {
	return _Marketviewer.Contract.AddMarket(&_Marketviewer.TransactOpts, newMarketAddresses)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _exchangeAdapter) returns()
func (_Marketviewer *MarketviewerTransactor) Init(opts *bind.TransactOpts, _exchangeAdapter common.Address) (*types.Transaction, error) {
	return _Marketviewer.contract.Transact(opts, "init", _exchangeAdapter)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _exchangeAdapter) returns()
func (_Marketviewer *MarketviewerSession) Init(_exchangeAdapter common.Address) (*types.Transaction, error) {
	return _Marketviewer.Contract.Init(&_Marketviewer.TransactOpts, _exchangeAdapter)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _exchangeAdapter) returns()
func (_Marketviewer *MarketviewerTransactorSession) Init(_exchangeAdapter common.Address) (*types.Transaction, error) {
	return _Marketviewer.Contract.Init(&_Marketviewer.TransactOpts, _exchangeAdapter)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Marketviewer *MarketviewerTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Marketviewer.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Marketviewer *MarketviewerSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Marketviewer.Contract.NominateNewOwner(&_Marketviewer.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Marketviewer *MarketviewerTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Marketviewer.Contract.NominateNewOwner(&_Marketviewer.TransactOpts, _owner)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_Marketviewer *MarketviewerTransactor) RemoveMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _Marketviewer.contract.Transact(opts, "removeMarket", market)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_Marketviewer *MarketviewerSession) RemoveMarket(market common.Address) (*types.Transaction, error) {
	return _Marketviewer.Contract.RemoveMarket(&_Marketviewer.TransactOpts, market)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_Marketviewer *MarketviewerTransactorSession) RemoveMarket(market common.Address) (*types.Transaction, error) {
	return _Marketviewer.Contract.RemoveMarket(&_Marketviewer.TransactOpts, market)
}

// MarketviewerMarketAddedIterator is returned from FilterMarketAdded and is used to iterate over the raw logs and unpacked data for MarketAdded events raised by the Marketviewer contract.
type MarketviewerMarketAddedIterator struct {
	Event *MarketviewerMarketAdded // Event containing the contract specifics and raw log

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
func (it *MarketviewerMarketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketviewerMarketAdded)
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
		it.Event = new(MarketviewerMarketAdded)
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
func (it *MarketviewerMarketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketviewerMarketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketviewerMarketAdded represents a MarketAdded event raised by the Marketviewer contract.
type MarketviewerMarketAdded struct {
	Market OptionMarketViewerOptionMarketAddresses
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketAdded is a free log retrieval operation binding the contract event 0xd3777d9870fa8e7682a422021d48549c180f7b06a647478f5211caef6b6a4ee8.
//
// Solidity: event MarketAdded((address,address,address,address,address,address,address,address,address,address) market)
func (_Marketviewer *MarketviewerFilterer) FilterMarketAdded(opts *bind.FilterOpts) (*MarketviewerMarketAddedIterator, error) {

	logs, sub, err := _Marketviewer.contract.FilterLogs(opts, "MarketAdded")
	if err != nil {
		return nil, err
	}
	return &MarketviewerMarketAddedIterator{contract: _Marketviewer.contract, event: "MarketAdded", logs: logs, sub: sub}, nil
}

// WatchMarketAdded is a free log subscription operation binding the contract event 0xd3777d9870fa8e7682a422021d48549c180f7b06a647478f5211caef6b6a4ee8.
//
// Solidity: event MarketAdded((address,address,address,address,address,address,address,address,address,address) market)
func (_Marketviewer *MarketviewerFilterer) WatchMarketAdded(opts *bind.WatchOpts, sink chan<- *MarketviewerMarketAdded) (event.Subscription, error) {

	logs, sub, err := _Marketviewer.contract.WatchLogs(opts, "MarketAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketviewerMarketAdded)
				if err := _Marketviewer.contract.UnpackLog(event, "MarketAdded", log); err != nil {
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

// ParseMarketAdded is a log parse operation binding the contract event 0xd3777d9870fa8e7682a422021d48549c180f7b06a647478f5211caef6b6a4ee8.
//
// Solidity: event MarketAdded((address,address,address,address,address,address,address,address,address,address) market)
func (_Marketviewer *MarketviewerFilterer) ParseMarketAdded(log types.Log) (*MarketviewerMarketAdded, error) {
	event := new(MarketviewerMarketAdded)
	if err := _Marketviewer.contract.UnpackLog(event, "MarketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketviewerMarketRemovedIterator is returned from FilterMarketRemoved and is used to iterate over the raw logs and unpacked data for MarketRemoved events raised by the Marketviewer contract.
type MarketviewerMarketRemovedIterator struct {
	Event *MarketviewerMarketRemoved // Event containing the contract specifics and raw log

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
func (it *MarketviewerMarketRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketviewerMarketRemoved)
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
		it.Event = new(MarketviewerMarketRemoved)
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
func (it *MarketviewerMarketRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketviewerMarketRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketviewerMarketRemoved represents a MarketRemoved event raised by the Marketviewer contract.
type MarketviewerMarketRemoved struct {
	Market common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketRemoved is a free log retrieval operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address market)
func (_Marketviewer *MarketviewerFilterer) FilterMarketRemoved(opts *bind.FilterOpts) (*MarketviewerMarketRemovedIterator, error) {

	logs, sub, err := _Marketviewer.contract.FilterLogs(opts, "MarketRemoved")
	if err != nil {
		return nil, err
	}
	return &MarketviewerMarketRemovedIterator{contract: _Marketviewer.contract, event: "MarketRemoved", logs: logs, sub: sub}, nil
}

// WatchMarketRemoved is a free log subscription operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address market)
func (_Marketviewer *MarketviewerFilterer) WatchMarketRemoved(opts *bind.WatchOpts, sink chan<- *MarketviewerMarketRemoved) (event.Subscription, error) {

	logs, sub, err := _Marketviewer.contract.WatchLogs(opts, "MarketRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketviewerMarketRemoved)
				if err := _Marketviewer.contract.UnpackLog(event, "MarketRemoved", log); err != nil {
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

// ParseMarketRemoved is a log parse operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address market)
func (_Marketviewer *MarketviewerFilterer) ParseMarketRemoved(log types.Log) (*MarketviewerMarketRemoved, error) {
	event := new(MarketviewerMarketRemoved)
	if err := _Marketviewer.contract.UnpackLog(event, "MarketRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketviewerOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Marketviewer contract.
type MarketviewerOwnerChangedIterator struct {
	Event *MarketviewerOwnerChanged // Event containing the contract specifics and raw log

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
func (it *MarketviewerOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketviewerOwnerChanged)
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
		it.Event = new(MarketviewerOwnerChanged)
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
func (it *MarketviewerOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketviewerOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketviewerOwnerChanged represents a OwnerChanged event raised by the Marketviewer contract.
type MarketviewerOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Marketviewer *MarketviewerFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*MarketviewerOwnerChangedIterator, error) {

	logs, sub, err := _Marketviewer.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &MarketviewerOwnerChangedIterator{contract: _Marketviewer.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Marketviewer *MarketviewerFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *MarketviewerOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Marketviewer.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketviewerOwnerChanged)
				if err := _Marketviewer.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_Marketviewer *MarketviewerFilterer) ParseOwnerChanged(log types.Log) (*MarketviewerOwnerChanged, error) {
	event := new(MarketviewerOwnerChanged)
	if err := _Marketviewer.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketviewerOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Marketviewer contract.
type MarketviewerOwnerNominatedIterator struct {
	Event *MarketviewerOwnerNominated // Event containing the contract specifics and raw log

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
func (it *MarketviewerOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketviewerOwnerNominated)
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
		it.Event = new(MarketviewerOwnerNominated)
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
func (it *MarketviewerOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketviewerOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketviewerOwnerNominated represents a OwnerNominated event raised by the Marketviewer contract.
type MarketviewerOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Marketviewer *MarketviewerFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*MarketviewerOwnerNominatedIterator, error) {

	logs, sub, err := _Marketviewer.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &MarketviewerOwnerNominatedIterator{contract: _Marketviewer.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Marketviewer *MarketviewerFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *MarketviewerOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Marketviewer.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketviewerOwnerNominated)
				if err := _Marketviewer.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_Marketviewer *MarketviewerFilterer) ParseOwnerNominated(log types.Log) (*MarketviewerOwnerNominated, error) {
	event := new(MarketviewerOwnerNominated)
	if err := _Marketviewer.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
