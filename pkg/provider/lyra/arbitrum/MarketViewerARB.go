// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum

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

// MarketViewerARBMetaData contains all meta data concerning the MarketViewerARB contract.
var MarketViewerARBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"}],\"name\":\"OnlyNominatedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"RemovingInvalidMarket\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"market\",\"type\":\"tuple\"}],\"name\":\"MarketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"newMarketAddresses\",\"type\":\"tuple\"}],\"name\":\"addMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeAdapter\",\"outputs\":[{\"internalType\":\"contractBaseExchangeAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoardForBase\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"getBoardForStrikeId\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLiquidityBalances\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteBalance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"quoteSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quoteDepositAllowance\",\"type\":\"uint256\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidityBalance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.LiquidityBalance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"getLiveBoards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"marketBoards\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"quoteSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quoteDecimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"baseDecimals\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"marketAddresses\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minDepositWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guardianMultisig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"guardianDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adjustmentNetScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callCollatScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putCollatScalingFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.LiquidityPoolParameters\",\"name\":\"lpParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"liquidityCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardSettlementCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractAdjustmentCBTimeout\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.CircuitBreakerParameters\",\"name\":\"cbParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxStrikesPerBoard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptableSpotPricePercentMove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staleUpdateDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewCap\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.GreekCacheParameters\",\"name\":\"greekCacheParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ivGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatePostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSpotMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateSpotMin\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.ForceCloseParameters\",\"name\":\"forceCloseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minStaticQuoteCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStaticBaseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callSpotPriceShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putSpotPriceShock\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.MinCollateralParameters\",\"name\":\"minCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.PricingParameters\",\"name\":\"pricingParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minForceCloseDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradingCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMinSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMaxSkew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"capSkewsToAbs\",\"type\":\"bool\"}],\"internalType\":\"structOptionMarketPricer.TradeLimitParameters\",\"name\":\"tradeLimitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"defaultVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referenceSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticSkewAdjustment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticIvVariance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeParameters\",\"name\":\"varianceFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"penaltyRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidationFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionToken.PartialCollateralParameters\",\"name\":\"partialCollatParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketParameters\",\"name\":\"marketParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"globalNetGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"liveBoards\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketView\",\"name\":\"marketView\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"}],\"name\":\"getMarketForBase\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"quoteSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quoteDecimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"baseDecimals\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"marketAddresses\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minDepositWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guardianMultisig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"guardianDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adjustmentNetScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callCollatScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putCollatScalingFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.LiquidityPoolParameters\",\"name\":\"lpParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"liquidityCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardSettlementCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractAdjustmentCBTimeout\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.CircuitBreakerParameters\",\"name\":\"cbParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxStrikesPerBoard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptableSpotPricePercentMove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staleUpdateDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewCap\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.GreekCacheParameters\",\"name\":\"greekCacheParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ivGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatePostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSpotMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateSpotMin\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.ForceCloseParameters\",\"name\":\"forceCloseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minStaticQuoteCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStaticBaseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callSpotPriceShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putSpotPriceShock\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.MinCollateralParameters\",\"name\":\"minCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.PricingParameters\",\"name\":\"pricingParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minForceCloseDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradingCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMinSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMaxSkew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"capSkewsToAbs\",\"type\":\"bool\"}],\"internalType\":\"structOptionMarketPricer.TradeLimitParameters\",\"name\":\"tradeLimitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"defaultVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referenceSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticSkewAdjustment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticIvVariance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeParameters\",\"name\":\"varianceFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"penaltyRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidationFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionToken.PartialCollateralParameters\",\"name\":\"partialCollatParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketParameters\",\"name\":\"marketParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"globalNetGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"liveBoards\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketView\",\"name\":\"marketView\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"name\":\"getMarkets\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"quoteSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quoteDecimals\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"baseDecimals\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"marketAddresses\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minDepositWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guardianMultisig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"guardianDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adjustmentNetScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callCollatScalingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putCollatScalingFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.LiquidityPoolParameters\",\"name\":\"lpParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"liquidityCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardSettlementCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractAdjustmentCBTimeout\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.CircuitBreakerParameters\",\"name\":\"cbParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxStrikesPerBoard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptableSpotPricePercentMove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staleUpdateDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewCap\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.GreekCacheParameters\",\"name\":\"greekCacheParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ivGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatePostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSpotMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateSpotMin\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.ForceCloseParameters\",\"name\":\"forceCloseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minStaticQuoteCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStaticBaseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callSpotPriceShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putSpotPriceShock\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.MinCollateralParameters\",\"name\":\"minCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.PricingParameters\",\"name\":\"pricingParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minForceCloseDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradingCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMinSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMaxSkew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"capSkewsToAbs\",\"type\":\"bool\"}],\"internalType\":\"structOptionMarketPricer.TradeLimitParameters\",\"name\":\"tradeLimitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"defaultVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referenceSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticSkewAdjustment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticIvVariance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeParameters\",\"name\":\"varianceFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"penaltyRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidationFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionToken.PartialCollateralParameters\",\"name\":\"partialCollatParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketParameters\",\"name\":\"marketParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"globalNetGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"varianceGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longScaleFactor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"liveBoards\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketView[]\",\"name\":\"markets\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketsView\",\"name\":\"marketsView\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getOwnerPositions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionToken.PositionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structOptionToken.OptionPosition[]\",\"name\":\"positions\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketOptionPositions[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseExchangeAdapter\",\"name\":\"_exchangeAdapter\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketAddresses\",\"outputs\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Decimals\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"optionMarkets\",\"outputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"removeMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketViewerARBABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketViewerARBMetaData.ABI instead.
var MarketViewerARBABI = MarketViewerARBMetaData.ABI

// MarketViewerARB is an auto generated Go binding around an Ethereum contract.
type MarketViewerARB struct {
	MarketViewerARBCaller     // Read-only binding to the contract
	MarketViewerARBTransactor // Write-only binding to the contract
	MarketViewerARBFilterer   // Log filterer for contract events
}

// MarketViewerARBCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketViewerARBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketViewerARBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketViewerARBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketViewerARBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketViewerARBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketViewerARBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketViewerARBSession struct {
	Contract     *MarketViewerARB  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketViewerARBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketViewerARBCallerSession struct {
	Contract *MarketViewerARBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MarketViewerARBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketViewerARBTransactorSession struct {
	Contract     *MarketViewerARBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MarketViewerARBRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketViewerARBRaw struct {
	Contract *MarketViewerARB // Generic contract binding to access the raw methods on
}

// MarketViewerARBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketViewerARBCallerRaw struct {
	Contract *MarketViewerARBCaller // Generic read-only contract binding to access the raw methods on
}

// MarketViewerARBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketViewerARBTransactorRaw struct {
	Contract *MarketViewerARBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketViewerARB creates a new instance of MarketViewerARB, bound to a specific deployed contract.
func NewMarketViewerARB(address common.Address, backend bind.ContractBackend) (*MarketViewerARB, error) {
	contract, err := bindMarketViewerARB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketViewerARB{MarketViewerARBCaller: MarketViewerARBCaller{contract: contract}, MarketViewerARBTransactor: MarketViewerARBTransactor{contract: contract}, MarketViewerARBFilterer: MarketViewerARBFilterer{contract: contract}}, nil
}

// NewMarketViewerARBCaller creates a new read-only instance of MarketViewerARB, bound to a specific deployed contract.
func NewMarketViewerARBCaller(address common.Address, caller bind.ContractCaller) (*MarketViewerARBCaller, error) {
	contract, err := bindMarketViewerARB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketViewerARBCaller{contract: contract}, nil
}

// NewMarketViewerARBTransactor creates a new write-only instance of MarketViewerARB, bound to a specific deployed contract.
func NewMarketViewerARBTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketViewerARBTransactor, error) {
	contract, err := bindMarketViewerARB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketViewerARBTransactor{contract: contract}, nil
}

// NewMarketViewerARBFilterer creates a new log filterer instance of MarketViewerARB, bound to a specific deployed contract.
func NewMarketViewerARBFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketViewerARBFilterer, error) {
	contract, err := bindMarketViewerARB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketViewerARBFilterer{contract: contract}, nil
}

// bindMarketViewerARB binds a generic wrapper to an already deployed contract.
func bindMarketViewerARB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketViewerARBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketViewerARB *MarketViewerARBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketViewerARB.Contract.MarketViewerARBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketViewerARB *MarketViewerARBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.MarketViewerARBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketViewerARB *MarketViewerARBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.MarketViewerARBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketViewerARB *MarketViewerARBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketViewerARB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketViewerARB *MarketViewerARBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketViewerARB *MarketViewerARBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.contract.Transact(opts, method, params...)
}

// ExchangeAdapter is a free data retrieval call binding the contract method 0x2cbcda25.
//
// Solidity: function exchangeAdapter() view returns(address)
func (_MarketViewerARB *MarketViewerARBCaller) ExchangeAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "exchangeAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeAdapter is a free data retrieval call binding the contract method 0x2cbcda25.
//
// Solidity: function exchangeAdapter() view returns(address)
func (_MarketViewerARB *MarketViewerARBSession) ExchangeAdapter() (common.Address, error) {
	return _MarketViewerARB.Contract.ExchangeAdapter(&_MarketViewerARB.CallOpts)
}

// ExchangeAdapter is a free data retrieval call binding the contract method 0x2cbcda25.
//
// Solidity: function exchangeAdapter() view returns(address)
func (_MarketViewerARB *MarketViewerARBCallerSession) ExchangeAdapter() (common.Address, error) {
	return _MarketViewerARB.Contract.ExchangeAdapter(&_MarketViewerARB.CallOpts)
}

// GetBoard is a free data retrieval call binding the contract method 0x5d033f7f.
//
// Solidity: function getBoard(address market, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerARB *MarketViewerARBCaller) GetBoard(opts *bind.CallOpts, market common.Address, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getBoard", market, boardId)

	if err != nil {
		return *new(OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerBoardView)).(*OptionMarketViewerBoardView)

	return out0, err

}

// GetBoard is a free data retrieval call binding the contract method 0x5d033f7f.
//
// Solidity: function getBoard(address market, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerARB *MarketViewerARBSession) GetBoard(market common.Address, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerARB.Contract.GetBoard(&_MarketViewerARB.CallOpts, market, boardId)
}

// GetBoard is a free data retrieval call binding the contract method 0x5d033f7f.
//
// Solidity: function getBoard(address market, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerARB *MarketViewerARBCallerSession) GetBoard(market common.Address, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerARB.Contract.GetBoard(&_MarketViewerARB.CallOpts, market, boardId)
}

// GetBoardForBase is a free data retrieval call binding the contract method 0x173b7f77.
//
// Solidity: function getBoardForBase(string baseSymbol, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerARB *MarketViewerARBCaller) GetBoardForBase(opts *bind.CallOpts, baseSymbol string, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getBoardForBase", baseSymbol, boardId)

	if err != nil {
		return *new(OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerBoardView)).(*OptionMarketViewerBoardView)

	return out0, err

}

// GetBoardForBase is a free data retrieval call binding the contract method 0x173b7f77.
//
// Solidity: function getBoardForBase(string baseSymbol, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerARB *MarketViewerARBSession) GetBoardForBase(baseSymbol string, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerARB.Contract.GetBoardForBase(&_MarketViewerARB.CallOpts, baseSymbol, boardId)
}

// GetBoardForBase is a free data retrieval call binding the contract method 0x173b7f77.
//
// Solidity: function getBoardForBase(string baseSymbol, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerARB *MarketViewerARBCallerSession) GetBoardForBase(baseSymbol string, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerARB.Contract.GetBoardForBase(&_MarketViewerARB.CallOpts, baseSymbol, boardId)
}

// GetBoardForStrikeId is a free data retrieval call binding the contract method 0xac6e227b.
//
// Solidity: function getBoardForStrikeId(address market, uint256 strikeId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerARB *MarketViewerARBCaller) GetBoardForStrikeId(opts *bind.CallOpts, market common.Address, strikeId *big.Int) (OptionMarketViewerBoardView, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getBoardForStrikeId", market, strikeId)

	if err != nil {
		return *new(OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerBoardView)).(*OptionMarketViewerBoardView)

	return out0, err

}

// GetBoardForStrikeId is a free data retrieval call binding the contract method 0xac6e227b.
//
// Solidity: function getBoardForStrikeId(address market, uint256 strikeId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerARB *MarketViewerARBSession) GetBoardForStrikeId(market common.Address, strikeId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerARB.Contract.GetBoardForStrikeId(&_MarketViewerARB.CallOpts, market, strikeId)
}

// GetBoardForStrikeId is a free data retrieval call binding the contract method 0xac6e227b.
//
// Solidity: function getBoardForStrikeId(address market, uint256 strikeId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerARB *MarketViewerARBCallerSession) GetBoardForStrikeId(market common.Address, strikeId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerARB.Contract.GetBoardForStrikeId(&_MarketViewerARB.CallOpts, market, strikeId)
}

// GetLiquidityBalances is a free data retrieval call binding the contract method 0x4af60f1a.
//
// Solidity: function getLiquidityBalances(address account) view returns((address,uint256,string,uint256,address,uint256)[])
func (_MarketViewerARB *MarketViewerARBCaller) GetLiquidityBalances(opts *bind.CallOpts, account common.Address) ([]OptionMarketViewerLiquidityBalance, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getLiquidityBalances", account)

	if err != nil {
		return *new([]OptionMarketViewerLiquidityBalance), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerLiquidityBalance)).(*[]OptionMarketViewerLiquidityBalance)

	return out0, err

}

// GetLiquidityBalances is a free data retrieval call binding the contract method 0x4af60f1a.
//
// Solidity: function getLiquidityBalances(address account) view returns((address,uint256,string,uint256,address,uint256)[])
func (_MarketViewerARB *MarketViewerARBSession) GetLiquidityBalances(account common.Address) ([]OptionMarketViewerLiquidityBalance, error) {
	return _MarketViewerARB.Contract.GetLiquidityBalances(&_MarketViewerARB.CallOpts, account)
}

// GetLiquidityBalances is a free data retrieval call binding the contract method 0x4af60f1a.
//
// Solidity: function getLiquidityBalances(address account) view returns((address,uint256,string,uint256,address,uint256)[])
func (_MarketViewerARB *MarketViewerARBCallerSession) GetLiquidityBalances(account common.Address) ([]OptionMarketViewerLiquidityBalance, error) {
	return _MarketViewerARB.Contract.GetLiquidityBalances(&_MarketViewerARB.CallOpts, account)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0xb0862c0e.
//
// Solidity: function getLiveBoards(address market) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[] marketBoards)
func (_MarketViewerARB *MarketViewerARBCaller) GetLiveBoards(opts *bind.CallOpts, market common.Address) ([]OptionMarketViewerBoardView, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getLiveBoards", market)

	if err != nil {
		return *new([]OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerBoardView)).(*[]OptionMarketViewerBoardView)

	return out0, err

}

// GetLiveBoards is a free data retrieval call binding the contract method 0xb0862c0e.
//
// Solidity: function getLiveBoards(address market) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[] marketBoards)
func (_MarketViewerARB *MarketViewerARBSession) GetLiveBoards(market common.Address) ([]OptionMarketViewerBoardView, error) {
	return _MarketViewerARB.Contract.GetLiveBoards(&_MarketViewerARB.CallOpts, market)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0xb0862c0e.
//
// Solidity: function getLiveBoards(address market) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[] marketBoards)
func (_MarketViewerARB *MarketViewerARBCallerSession) GetLiveBoards(market common.Address) ([]OptionMarketViewerBoardView, error) {
	return _MarketViewerARB.Contract.GetLiveBoards(&_MarketViewerARB.CallOpts, market)
}

// GetMarket is a free data retrieval call binding the contract method 0xd4dfadbf.
//
// Solidity: function getMarket(address market) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_MarketViewerARB *MarketViewerARBCaller) GetMarket(opts *bind.CallOpts, market common.Address) (OptionMarketViewerMarketView, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getMarket", market)

	if err != nil {
		return *new(OptionMarketViewerMarketView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerMarketView)).(*OptionMarketViewerMarketView)

	return out0, err

}

// GetMarket is a free data retrieval call binding the contract method 0xd4dfadbf.
//
// Solidity: function getMarket(address market) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_MarketViewerARB *MarketViewerARBSession) GetMarket(market common.Address) (OptionMarketViewerMarketView, error) {
	return _MarketViewerARB.Contract.GetMarket(&_MarketViewerARB.CallOpts, market)
}

// GetMarket is a free data retrieval call binding the contract method 0xd4dfadbf.
//
// Solidity: function getMarket(address market) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_MarketViewerARB *MarketViewerARBCallerSession) GetMarket(market common.Address) (OptionMarketViewerMarketView, error) {
	return _MarketViewerARB.Contract.GetMarket(&_MarketViewerARB.CallOpts, market)
}

// GetMarketAddresses is a free data retrieval call binding the contract method 0x97ce0d31.
//
// Solidity: function getMarketAddresses() view returns((address,address,address,address,address,address,address,address,address,address)[])
func (_MarketViewerARB *MarketViewerARBCaller) GetMarketAddresses(opts *bind.CallOpts) ([]OptionMarketViewerOptionMarketAddresses, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getMarketAddresses")

	if err != nil {
		return *new([]OptionMarketViewerOptionMarketAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerOptionMarketAddresses)).(*[]OptionMarketViewerOptionMarketAddresses)

	return out0, err

}

// GetMarketAddresses is a free data retrieval call binding the contract method 0x97ce0d31.
//
// Solidity: function getMarketAddresses() view returns((address,address,address,address,address,address,address,address,address,address)[])
func (_MarketViewerARB *MarketViewerARBSession) GetMarketAddresses() ([]OptionMarketViewerOptionMarketAddresses, error) {
	return _MarketViewerARB.Contract.GetMarketAddresses(&_MarketViewerARB.CallOpts)
}

// GetMarketAddresses is a free data retrieval call binding the contract method 0x97ce0d31.
//
// Solidity: function getMarketAddresses() view returns((address,address,address,address,address,address,address,address,address,address)[])
func (_MarketViewerARB *MarketViewerARBCallerSession) GetMarketAddresses() ([]OptionMarketViewerOptionMarketAddresses, error) {
	return _MarketViewerARB.Contract.GetMarketAddresses(&_MarketViewerARB.CallOpts)
}

// GetMarketForBase is a free data retrieval call binding the contract method 0x96c121b5.
//
// Solidity: function getMarketForBase(string baseSymbol) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_MarketViewerARB *MarketViewerARBCaller) GetMarketForBase(opts *bind.CallOpts, baseSymbol string) (OptionMarketViewerMarketView, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getMarketForBase", baseSymbol)

	if err != nil {
		return *new(OptionMarketViewerMarketView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerMarketView)).(*OptionMarketViewerMarketView)

	return out0, err

}

// GetMarketForBase is a free data retrieval call binding the contract method 0x96c121b5.
//
// Solidity: function getMarketForBase(string baseSymbol) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_MarketViewerARB *MarketViewerARBSession) GetMarketForBase(baseSymbol string) (OptionMarketViewerMarketView, error) {
	return _MarketViewerARB.Contract.GetMarketForBase(&_MarketViewerARB.CallOpts, baseSymbol)
}

// GetMarketForBase is a free data retrieval call binding the contract method 0x96c121b5.
//
// Solidity: function getMarketForBase(string baseSymbol) view returns((bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[]) marketView)
func (_MarketViewerARB *MarketViewerARBCallerSession) GetMarketForBase(baseSymbol string) (OptionMarketViewerMarketView, error) {
	return _MarketViewerARB.Contract.GetMarketForBase(&_MarketViewerARB.CallOpts, baseSymbol)
}

// GetMarkets is a free data retrieval call binding the contract method 0x1139e3f1.
//
// Solidity: function getMarkets(address[] markets) view returns((bool,(bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[])[]) marketsView)
func (_MarketViewerARB *MarketViewerARBCaller) GetMarkets(opts *bind.CallOpts, markets []common.Address) (OptionMarketViewerMarketsView, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getMarkets", markets)

	if err != nil {
		return *new(OptionMarketViewerMarketsView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerMarketsView)).(*OptionMarketViewerMarketsView)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0x1139e3f1.
//
// Solidity: function getMarkets(address[] markets) view returns((bool,(bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[])[]) marketsView)
func (_MarketViewerARB *MarketViewerARBSession) GetMarkets(markets []common.Address) (OptionMarketViewerMarketsView, error) {
	return _MarketViewerARB.Contract.GetMarkets(&_MarketViewerARB.CallOpts, markets)
}

// GetMarkets is a free data retrieval call binding the contract method 0x1139e3f1.
//
// Solidity: function getMarkets(address[] markets) view returns((bool,(bool,uint256,string,uint256,string,uint256,(uint256,uint256,uint256,uint256,uint256,uint256,uint256),(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[])[]) marketsView)
func (_MarketViewerARB *MarketViewerARBCallerSession) GetMarkets(markets []common.Address) (OptionMarketViewerMarketsView, error) {
	return _MarketViewerARB.Contract.GetMarkets(&_MarketViewerARB.CallOpts, markets)
}

// GetOwnerPositions is a free data retrieval call binding the contract method 0x845e1259.
//
// Solidity: function getOwnerPositions(address account) view returns((address,(uint256,uint256,uint8,uint256,uint256,uint8)[])[])
func (_MarketViewerARB *MarketViewerARBCaller) GetOwnerPositions(opts *bind.CallOpts, account common.Address) ([]OptionMarketViewerMarketOptionPositions, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "getOwnerPositions", account)

	if err != nil {
		return *new([]OptionMarketViewerMarketOptionPositions), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerMarketOptionPositions)).(*[]OptionMarketViewerMarketOptionPositions)

	return out0, err

}

// GetOwnerPositions is a free data retrieval call binding the contract method 0x845e1259.
//
// Solidity: function getOwnerPositions(address account) view returns((address,(uint256,uint256,uint8,uint256,uint256,uint8)[])[])
func (_MarketViewerARB *MarketViewerARBSession) GetOwnerPositions(account common.Address) ([]OptionMarketViewerMarketOptionPositions, error) {
	return _MarketViewerARB.Contract.GetOwnerPositions(&_MarketViewerARB.CallOpts, account)
}

// GetOwnerPositions is a free data retrieval call binding the contract method 0x845e1259.
//
// Solidity: function getOwnerPositions(address account) view returns((address,(uint256,uint256,uint8,uint256,uint256,uint8)[])[])
func (_MarketViewerARB *MarketViewerARBCallerSession) GetOwnerPositions(account common.Address) ([]OptionMarketViewerMarketOptionPositions, error) {
	return _MarketViewerARB.Contract.GetOwnerPositions(&_MarketViewerARB.CallOpts, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_MarketViewerARB *MarketViewerARBCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_MarketViewerARB *MarketViewerARBSession) Initialized() (bool, error) {
	return _MarketViewerARB.Contract.Initialized(&_MarketViewerARB.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_MarketViewerARB *MarketViewerARBCallerSession) Initialized() (bool, error) {
	return _MarketViewerARB.Contract.Initialized(&_MarketViewerARB.CallOpts)
}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address shortCollateral, address poolHedger, address quoteAsset, address baseAsset)
func (_MarketViewerARB *MarketViewerARBCaller) MarketAddresses(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "marketAddresses", arg0)

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
func (_MarketViewerARB *MarketViewerARBSession) MarketAddresses(arg0 common.Address) (struct {
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
	return _MarketViewerARB.Contract.MarketAddresses(&_MarketViewerARB.CallOpts, arg0)
}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address shortCollateral, address poolHedger, address quoteAsset, address baseAsset)
func (_MarketViewerARB *MarketViewerARBCallerSession) MarketAddresses(arg0 common.Address) (struct {
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
	return _MarketViewerARB.Contract.MarketAddresses(&_MarketViewerARB.CallOpts, arg0)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_MarketViewerARB *MarketViewerARBCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_MarketViewerARB *MarketViewerARBSession) NominatedOwner() (common.Address, error) {
	return _MarketViewerARB.Contract.NominatedOwner(&_MarketViewerARB.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_MarketViewerARB *MarketViewerARBCallerSession) NominatedOwner() (common.Address, error) {
	return _MarketViewerARB.Contract.NominatedOwner(&_MarketViewerARB.CallOpts)
}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_MarketViewerARB *MarketViewerARBCaller) OptionMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "optionMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_MarketViewerARB *MarketViewerARBSession) OptionMarkets(arg0 *big.Int) (common.Address, error) {
	return _MarketViewerARB.Contract.OptionMarkets(&_MarketViewerARB.CallOpts, arg0)
}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_MarketViewerARB *MarketViewerARBCallerSession) OptionMarkets(arg0 *big.Int) (common.Address, error) {
	return _MarketViewerARB.Contract.OptionMarkets(&_MarketViewerARB.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketViewerARB *MarketViewerARBCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketViewerARB.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketViewerARB *MarketViewerARBSession) Owner() (common.Address, error) {
	return _MarketViewerARB.Contract.Owner(&_MarketViewerARB.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketViewerARB *MarketViewerARBCallerSession) Owner() (common.Address, error) {
	return _MarketViewerARB.Contract.Owner(&_MarketViewerARB.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MarketViewerARB *MarketViewerARBTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketViewerARB.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MarketViewerARB *MarketViewerARBSession) AcceptOwnership() (*types.Transaction, error) {
	return _MarketViewerARB.Contract.AcceptOwnership(&_MarketViewerARB.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MarketViewerARB *MarketViewerARBTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MarketViewerARB.Contract.AcceptOwnership(&_MarketViewerARB.TransactOpts)
}

// AddMarket is a paid mutator transaction binding the contract method 0x7f68720f.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_MarketViewerARB *MarketViewerARBTransactor) AddMarket(opts *bind.TransactOpts, newMarketAddresses OptionMarketViewerOptionMarketAddresses) (*types.Transaction, error) {
	return _MarketViewerARB.contract.Transact(opts, "addMarket", newMarketAddresses)
}

// AddMarket is a paid mutator transaction binding the contract method 0x7f68720f.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_MarketViewerARB *MarketViewerARBSession) AddMarket(newMarketAddresses OptionMarketViewerOptionMarketAddresses) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.AddMarket(&_MarketViewerARB.TransactOpts, newMarketAddresses)
}

// AddMarket is a paid mutator transaction binding the contract method 0x7f68720f.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_MarketViewerARB *MarketViewerARBTransactorSession) AddMarket(newMarketAddresses OptionMarketViewerOptionMarketAddresses) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.AddMarket(&_MarketViewerARB.TransactOpts, newMarketAddresses)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _exchangeAdapter) returns()
func (_MarketViewerARB *MarketViewerARBTransactor) Init(opts *bind.TransactOpts, _exchangeAdapter common.Address) (*types.Transaction, error) {
	return _MarketViewerARB.contract.Transact(opts, "init", _exchangeAdapter)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _exchangeAdapter) returns()
func (_MarketViewerARB *MarketViewerARBSession) Init(_exchangeAdapter common.Address) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.Init(&_MarketViewerARB.TransactOpts, _exchangeAdapter)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _exchangeAdapter) returns()
func (_MarketViewerARB *MarketViewerARBTransactorSession) Init(_exchangeAdapter common.Address) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.Init(&_MarketViewerARB.TransactOpts, _exchangeAdapter)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_MarketViewerARB *MarketViewerARBTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _MarketViewerARB.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_MarketViewerARB *MarketViewerARBSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.NominateNewOwner(&_MarketViewerARB.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_MarketViewerARB *MarketViewerARBTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.NominateNewOwner(&_MarketViewerARB.TransactOpts, _owner)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_MarketViewerARB *MarketViewerARBTransactor) RemoveMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _MarketViewerARB.contract.Transact(opts, "removeMarket", market)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_MarketViewerARB *MarketViewerARBSession) RemoveMarket(market common.Address) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.RemoveMarket(&_MarketViewerARB.TransactOpts, market)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_MarketViewerARB *MarketViewerARBTransactorSession) RemoveMarket(market common.Address) (*types.Transaction, error) {
	return _MarketViewerARB.Contract.RemoveMarket(&_MarketViewerARB.TransactOpts, market)
}

// MarketViewerARBMarketAddedIterator is returned from FilterMarketAdded and is used to iterate over the raw logs and unpacked data for MarketAdded events raised by the MarketViewerARB contract.
type MarketViewerARBMarketAddedIterator struct {
	Event *MarketViewerARBMarketAdded // Event containing the contract specifics and raw log

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
func (it *MarketViewerARBMarketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketViewerARBMarketAdded)
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
		it.Event = new(MarketViewerARBMarketAdded)
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
func (it *MarketViewerARBMarketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketViewerARBMarketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketViewerARBMarketAdded represents a MarketAdded event raised by the MarketViewerARB contract.
type MarketViewerARBMarketAdded struct {
	Market OptionMarketViewerOptionMarketAddresses
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketAdded is a free log retrieval operation binding the contract event 0xd3777d9870fa8e7682a422021d48549c180f7b06a647478f5211caef6b6a4ee8.
//
// Solidity: event MarketAdded((address,address,address,address,address,address,address,address,address,address) market)
func (_MarketViewerARB *MarketViewerARBFilterer) FilterMarketAdded(opts *bind.FilterOpts) (*MarketViewerARBMarketAddedIterator, error) {

	logs, sub, err := _MarketViewerARB.contract.FilterLogs(opts, "MarketAdded")
	if err != nil {
		return nil, err
	}
	return &MarketViewerARBMarketAddedIterator{contract: _MarketViewerARB.contract, event: "MarketAdded", logs: logs, sub: sub}, nil
}

// WatchMarketAdded is a free log subscription operation binding the contract event 0xd3777d9870fa8e7682a422021d48549c180f7b06a647478f5211caef6b6a4ee8.
//
// Solidity: event MarketAdded((address,address,address,address,address,address,address,address,address,address) market)
func (_MarketViewerARB *MarketViewerARBFilterer) WatchMarketAdded(opts *bind.WatchOpts, sink chan<- *MarketViewerARBMarketAdded) (event.Subscription, error) {

	logs, sub, err := _MarketViewerARB.contract.WatchLogs(opts, "MarketAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketViewerARBMarketAdded)
				if err := _MarketViewerARB.contract.UnpackLog(event, "MarketAdded", log); err != nil {
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
func (_MarketViewerARB *MarketViewerARBFilterer) ParseMarketAdded(log types.Log) (*MarketViewerARBMarketAdded, error) {
	event := new(MarketViewerARBMarketAdded)
	if err := _MarketViewerARB.contract.UnpackLog(event, "MarketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketViewerARBMarketRemovedIterator is returned from FilterMarketRemoved and is used to iterate over the raw logs and unpacked data for MarketRemoved events raised by the MarketViewerARB contract.
type MarketViewerARBMarketRemovedIterator struct {
	Event *MarketViewerARBMarketRemoved // Event containing the contract specifics and raw log

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
func (it *MarketViewerARBMarketRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketViewerARBMarketRemoved)
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
		it.Event = new(MarketViewerARBMarketRemoved)
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
func (it *MarketViewerARBMarketRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketViewerARBMarketRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketViewerARBMarketRemoved represents a MarketRemoved event raised by the MarketViewerARB contract.
type MarketViewerARBMarketRemoved struct {
	Market common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketRemoved is a free log retrieval operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address market)
func (_MarketViewerARB *MarketViewerARBFilterer) FilterMarketRemoved(opts *bind.FilterOpts) (*MarketViewerARBMarketRemovedIterator, error) {

	logs, sub, err := _MarketViewerARB.contract.FilterLogs(opts, "MarketRemoved")
	if err != nil {
		return nil, err
	}
	return &MarketViewerARBMarketRemovedIterator{contract: _MarketViewerARB.contract, event: "MarketRemoved", logs: logs, sub: sub}, nil
}

// WatchMarketRemoved is a free log subscription operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address market)
func (_MarketViewerARB *MarketViewerARBFilterer) WatchMarketRemoved(opts *bind.WatchOpts, sink chan<- *MarketViewerARBMarketRemoved) (event.Subscription, error) {

	logs, sub, err := _MarketViewerARB.contract.WatchLogs(opts, "MarketRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketViewerARBMarketRemoved)
				if err := _MarketViewerARB.contract.UnpackLog(event, "MarketRemoved", log); err != nil {
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
func (_MarketViewerARB *MarketViewerARBFilterer) ParseMarketRemoved(log types.Log) (*MarketViewerARBMarketRemoved, error) {
	event := new(MarketViewerARBMarketRemoved)
	if err := _MarketViewerARB.contract.UnpackLog(event, "MarketRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketViewerARBOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the MarketViewerARB contract.
type MarketViewerARBOwnerChangedIterator struct {
	Event *MarketViewerARBOwnerChanged // Event containing the contract specifics and raw log

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
func (it *MarketViewerARBOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketViewerARBOwnerChanged)
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
		it.Event = new(MarketViewerARBOwnerChanged)
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
func (it *MarketViewerARBOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketViewerARBOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketViewerARBOwnerChanged represents a OwnerChanged event raised by the MarketViewerARB contract.
type MarketViewerARBOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_MarketViewerARB *MarketViewerARBFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*MarketViewerARBOwnerChangedIterator, error) {

	logs, sub, err := _MarketViewerARB.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &MarketViewerARBOwnerChangedIterator{contract: _MarketViewerARB.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_MarketViewerARB *MarketViewerARBFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *MarketViewerARBOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _MarketViewerARB.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketViewerARBOwnerChanged)
				if err := _MarketViewerARB.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_MarketViewerARB *MarketViewerARBFilterer) ParseOwnerChanged(log types.Log) (*MarketViewerARBOwnerChanged, error) {
	event := new(MarketViewerARBOwnerChanged)
	if err := _MarketViewerARB.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketViewerARBOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the MarketViewerARB contract.
type MarketViewerARBOwnerNominatedIterator struct {
	Event *MarketViewerARBOwnerNominated // Event containing the contract specifics and raw log

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
func (it *MarketViewerARBOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketViewerARBOwnerNominated)
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
		it.Event = new(MarketViewerARBOwnerNominated)
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
func (it *MarketViewerARBOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketViewerARBOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketViewerARBOwnerNominated represents a OwnerNominated event raised by the MarketViewerARB contract.
type MarketViewerARBOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_MarketViewerARB *MarketViewerARBFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*MarketViewerARBOwnerNominatedIterator, error) {

	logs, sub, err := _MarketViewerARB.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &MarketViewerARBOwnerNominatedIterator{contract: _MarketViewerARB.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_MarketViewerARB *MarketViewerARBFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *MarketViewerARBOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _MarketViewerARB.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketViewerARBOwnerNominated)
				if err := _MarketViewerARB.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_MarketViewerARB *MarketViewerARBFilterer) ParseOwnerNominated(log types.Log) (*MarketViewerARBOwnerNominated, error) {
	event := new(MarketViewerARBOwnerNominated)
	if err := _MarketViewerARB.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
