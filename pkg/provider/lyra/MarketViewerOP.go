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

// LiquidityPoolLiquidity is an auto generated low-level Go binding around an user-defined struct.
type LiquidityPoolLiquidity struct {
	FreeLiquidity         *big.Int
	BurnableLiquidity     *big.Int
	UsedCollatLiquidity   *big.Int
	PendingDeltaLiquidity *big.Int
	UsedDeltaLiquidity    *big.Int
	NAV                   *big.Int
}

// LiquidityPoolLiquidityPoolParameters is an auto generated low-level Go binding around an user-defined struct.
type LiquidityPoolLiquidityPoolParameters struct {
	MinDepositWithdraw       *big.Int
	DepositDelay             *big.Int
	WithdrawalDelay          *big.Int
	WithdrawalFee            *big.Int
	LiquidityCBThreshold     *big.Int
	LiquidityCBTimeout       *big.Int
	IvVarianceCBThreshold    *big.Int
	SkewVarianceCBThreshold  *big.Int
	IvVarianceCBTimeout      *big.Int
	SkewVarianceCBTimeout    *big.Int
	GuardianMultisig         common.Address
	GuardianDelay            *big.Int
	BoardSettlementCBTimeout *big.Int
	MaxFeePaid               *big.Int
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
	RateAndCarry                   *big.Int
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
	ForceCloseGwavIV *big.Int
	NetGreeks        OptionGreekCacheNetGreeks
	Strikes          []OptionMarketViewerStrikeView
}

// OptionMarketViewerLiquidityBalanceAndAllowance is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerLiquidityBalanceAndAllowance struct {
	Token     common.Address
	Balance   *big.Int
	Allowance *big.Int
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
	GreekCacheParams    OptionGreekCacheGreekCacheParameters
	ForceCloseParams    OptionGreekCacheForceCloseParameters
	MinCollatParams     OptionGreekCacheMinCollateralParameters
	PricingParams       OptionMarketPricerPricingParameters
	TradeLimitParams    OptionMarketPricerTradeLimitParameters
	VarianceFeeParams   OptionMarketPricerVarianceFeeParameters
	PartialCollatParams OptionTokenPartialCollateralParameters
	PoolHedgerParams    PoolHedgerPoolHedgerParameters
}

// OptionMarketViewerMarketView is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerMarketView struct {
	IsPaused               bool
	TotalQueuedDeposits    *big.Int
	TotalQueuedWithdrawals *big.Int
	TokenPrice             *big.Int
	MarketAddresses        OptionMarketViewerOptionMarketAddresses
	MarketParameters       OptionMarketViewerMarketParameters
	Liquidity              LiquidityPoolLiquidity
	GlobalNetGreeks        OptionGreekCacheNetGreeks
	ExchangeParams         SynthetixAdapterExchangeParams
}

// OptionMarketViewerMarketViewWithBoards is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerMarketViewWithBoards struct {
	IsPaused               bool
	TotalQueuedDeposits    *big.Int
	TotalQueuedWithdrawals *big.Int
	TokenPrice             *big.Int
	MarketAddresses        OptionMarketViewerOptionMarketAddresses
	MarketParameters       OptionMarketViewerMarketParameters
	Liquidity              LiquidityPoolLiquidity
	GlobalNetGreeks        OptionGreekCacheNetGreeks
	LiveBoards             []OptionMarketViewerBoardView
	ExchangeParams         SynthetixAdapterExchangeParams
}

// OptionMarketViewerMarketsView is an auto generated low-level Go binding around an user-defined struct.
type OptionMarketViewerMarketsView struct {
	AddressResolver common.Address
	IsPaused        bool
	Markets         []OptionMarketViewerMarketView
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

// PoolHedgerPoolHedgerParameters is an auto generated low-level Go binding around an user-defined struct.
type PoolHedgerPoolHedgerParameters struct {
	InteractionDelay *big.Int
	HedgeCap         *big.Int
}

// SynthetixAdapterExchangeParams is an auto generated low-level Go binding around an user-defined struct.
type SynthetixAdapterExchangeParams struct {
	SpotPrice        *big.Int
	QuoteKey         [32]byte
	BaseKey          [32]byte
	QuoteBaseFeeRate *big.Int
	BaseQuoteFeeRate *big.Int
}

// MarketViewerOPMetaData contains all meta data concerning the MarketViewerOP contract.
var MarketViewerOPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"}],\"name\":\"OnlyNominatedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"RemovingInvalidMarket\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"market\",\"type\":\"tuple\"}],\"name\":\"MarketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"newMarketAddresses\",\"type\":\"tuple\"}],\"name\":\"addMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoard\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIV\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"baseKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoardForBaseKey\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIV\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"}],\"name\":\"getBoardForStrikeId\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIV\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLiquidityBalancesAndAllowances\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.LiquidityBalanceAndAllowance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"getLiveBoards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIV\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"marketBoards\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalQueuedDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalQueuedWithdrawals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"marketAddresses\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minDepositWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guardianMultisig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"guardianDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardSettlementCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.LiquidityPoolParameters\",\"name\":\"lpParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxStrikesPerBoard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptableSpotPricePercentMove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staleUpdateDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewCap\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rateAndCarry\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.GreekCacheParameters\",\"name\":\"greekCacheParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ivGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatePostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSpotMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateSpotMin\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.ForceCloseParameters\",\"name\":\"forceCloseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minStaticQuoteCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStaticBaseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callSpotPriceShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putSpotPriceShock\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.MinCollateralParameters\",\"name\":\"minCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.PricingParameters\",\"name\":\"pricingParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minForceCloseDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradingCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMinSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMaxSkew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"capSkewsToAbs\",\"type\":\"bool\"}],\"internalType\":\"structOptionMarketPricer.TradeLimitParameters\",\"name\":\"tradeLimitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"defaultVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referenceSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticSkewAdjustment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticIvVariance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeParameters\",\"name\":\"varianceFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"penaltyRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidationFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionToken.PartialCollateralParameters\",\"name\":\"partialCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"interactionDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hedgeCap\",\"type\":\"uint256\"}],\"internalType\":\"structPoolHedger.PoolHedgerParameters\",\"name\":\"poolHedgerParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketParameters\",\"name\":\"marketParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"globalNetGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIV\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"liveBoards\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"quoteKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"baseKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quoteBaseFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseQuoteFeeRate\",\"type\":\"uint256\"}],\"internalType\":\"structSynthetixAdapter.ExchangeParams\",\"name\":\"exchangeParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketViewWithBoards\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"baseKey\",\"type\":\"bytes32\"}],\"name\":\"getMarketForBaseKey\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalQueuedDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalQueuedWithdrawals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"marketAddresses\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minDepositWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guardianMultisig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"guardianDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardSettlementCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.LiquidityPoolParameters\",\"name\":\"lpParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxStrikesPerBoard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptableSpotPricePercentMove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staleUpdateDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewCap\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rateAndCarry\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.GreekCacheParameters\",\"name\":\"greekCacheParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ivGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatePostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSpotMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateSpotMin\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.ForceCloseParameters\",\"name\":\"forceCloseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minStaticQuoteCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStaticBaseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callSpotPriceShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putSpotPriceShock\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.MinCollateralParameters\",\"name\":\"minCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.PricingParameters\",\"name\":\"pricingParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minForceCloseDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradingCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMinSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMaxSkew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"capSkewsToAbs\",\"type\":\"bool\"}],\"internalType\":\"structOptionMarketPricer.TradeLimitParameters\",\"name\":\"tradeLimitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"defaultVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referenceSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticSkewAdjustment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticIvVariance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeParameters\",\"name\":\"varianceFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"penaltyRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidationFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionToken.PartialCollateralParameters\",\"name\":\"partialCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"interactionDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hedgeCap\",\"type\":\"uint256\"}],\"internalType\":\"structPoolHedger.PoolHedgerParameters\",\"name\":\"poolHedgerParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketParameters\",\"name\":\"marketParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"globalNetGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseIv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAtExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseGwavIV\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"netGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseSkew\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"callDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"putDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"stdVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putPrice\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.StrikeGreeks\",\"name\":\"cachedGreeks\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"baseReturnedRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longCallOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPutOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallBaseOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortCallQuoteOpenInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPutOpenInterest\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketViewer.StrikeView[]\",\"name\":\"strikes\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.BoardView[]\",\"name\":\"liveBoards\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"quoteKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"baseKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quoteBaseFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseQuoteFeeRate\",\"type\":\"uint256\"}],\"internalType\":\"structSynthetixAdapter.ExchangeParams\",\"name\":\"exchangeParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketViewWithBoards\",\"name\":\"market\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"name\":\"getMarkets\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIAddressResolver\",\"name\":\"addressResolver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalQueuedDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalQueuedWithdrawals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structOptionMarketViewer.OptionMarketAddresses\",\"name\":\"marketAddresses\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxBoardExpiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"securityModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePortionReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBaseSettlementFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarket.OptionMarketParameters\",\"name\":\"optionMarketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minDepositWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewVarianceCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guardianMultisig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"guardianDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boardSettlementCBTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.LiquidityPoolParameters\",\"name\":\"lpParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxStrikesPerBoard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptableSpotPricePercentMove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staleUpdateDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"varianceSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueIvGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionValueSkewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gwavSkewCap\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rateAndCarry\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.GreekCacheParameters\",\"name\":\"greekCacheParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ivGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewGWAVPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatePostCutoffVolShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSpotMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateSpotMin\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.ForceCloseParameters\",\"name\":\"forceCloseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minStaticQuoteCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStaticBaseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shockVolPointB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callSpotPriceShock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"putSpotPriceShock\",\"type\":\"uint256\"}],\"internalType\":\"structOptionGreekCache.MinCollateralParameters\",\"name\":\"minCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"optionPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee1xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceFee2xPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"standardSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentFactor\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.PricingParameters\",\"name\":\"pricingParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"minDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minForceCloseDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradingCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBaseIV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMinSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absMaxSkew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"capSkewsToAbs\",\"type\":\"bool\"}],\"internalType\":\"structOptionMarketPricer.TradeLimitParameters\",\"name\":\"tradeLimitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"defaultVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forceCloseVarianceFeeCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewAdjustmentCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referenceSkew\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticSkewAdjustment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vegaCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticVega\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ivVarianceCoefficient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStaticIvVariance\",\"type\":\"uint256\"}],\"internalType\":\"structOptionMarketPricer.VarianceFeeParameters\",\"name\":\"varianceFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"penaltyRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidatorFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"smFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidationFee\",\"type\":\"uint256\"}],\"internalType\":\"structOptionToken.PartialCollateralParameters\",\"name\":\"partialCollatParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"interactionDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hedgeCap\",\"type\":\"uint256\"}],\"internalType\":\"structPoolHedger.PoolHedgerParameters\",\"name\":\"poolHedgerParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketParameters\",\"name\":\"marketParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedCollatLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedDeltaLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NAV\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityPool.Liquidity\",\"name\":\"liquidity\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"netDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netStdVega\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netOptionValue\",\"type\":\"int256\"}],\"internalType\":\"structOptionGreekCache.NetGreeks\",\"name\":\"globalNetGreeks\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"quoteKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"baseKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quoteBaseFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseQuoteFeeRate\",\"type\":\"uint256\"}],\"internalType\":\"structSynthetixAdapter.ExchangeParams\",\"name\":\"exchangeParams\",\"type\":\"tuple\"}],\"internalType\":\"structOptionMarketViewer.MarketView[]\",\"name\":\"markets\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketsView\",\"name\":\"marketsView\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getOwnerPositions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionToken.PositionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structOptionToken.OptionPosition[]\",\"name\":\"positions\",\"type\":\"tuple[]\"}],\"internalType\":\"structOptionMarketViewer.MarketOptionPositions[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getOwnerPositionsInRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikeId\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionMarket.OptionType\",\"name\":\"optionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumOptionToken.PositionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structOptionToken.OptionPosition[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractSynthetixAdapter\",\"name\":\"_synthetixAdapter\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketAddresses\",\"outputs\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"optionMarkets\",\"outputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"removeMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synthetixAdapter\",\"outputs\":[{\"internalType\":\"contractSynthetixAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MarketViewerOPABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketViewerOPMetaData.ABI instead.
var MarketViewerOPABI = MarketViewerOPMetaData.ABI

// MarketViewerOP is an auto generated Go binding around an Ethereum contract.
type MarketViewerOP struct {
	MarketViewerOPCaller     // Read-only binding to the contract
	MarketViewerOPTransactor // Write-only binding to the contract
	MarketViewerOPFilterer   // Log filterer for contract events
}

// MarketViewerOPCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketViewerOPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketViewerOPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketViewerOPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketViewerOPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketViewerOPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketViewerOPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketViewerOPSession struct {
	Contract     *MarketViewerOP   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketViewerOPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketViewerOPCallerSession struct {
	Contract *MarketViewerOPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MarketViewerOPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketViewerOPTransactorSession struct {
	Contract     *MarketViewerOPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MarketViewerOPRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketViewerOPRaw struct {
	Contract *MarketViewerOP // Generic contract binding to access the raw methods on
}

// MarketViewerOPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketViewerOPCallerRaw struct {
	Contract *MarketViewerOPCaller // Generic read-only contract binding to access the raw methods on
}

// MarketViewerOPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketViewerOPTransactorRaw struct {
	Contract *MarketViewerOPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketViewerOP creates a new instance of MarketViewerOP, bound to a specific deployed contract.
func NewMarketViewerOP(address common.Address, backend bind.ContractBackend) (*MarketViewerOP, error) {
	contract, err := bindMarketViewerOP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketViewerOP{MarketViewerOPCaller: MarketViewerOPCaller{contract: contract}, MarketViewerOPTransactor: MarketViewerOPTransactor{contract: contract}, MarketViewerOPFilterer: MarketViewerOPFilterer{contract: contract}}, nil
}

// NewMarketViewerOPCaller creates a new read-only instance of MarketViewerOP, bound to a specific deployed contract.
func NewMarketViewerOPCaller(address common.Address, caller bind.ContractCaller) (*MarketViewerOPCaller, error) {
	contract, err := bindMarketViewerOP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketViewerOPCaller{contract: contract}, nil
}

// NewMarketViewerOPTransactor creates a new write-only instance of MarketViewerOP, bound to a specific deployed contract.
func NewMarketViewerOPTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketViewerOPTransactor, error) {
	contract, err := bindMarketViewerOP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketViewerOPTransactor{contract: contract}, nil
}

// NewMarketViewerOPFilterer creates a new log filterer instance of MarketViewerOP, bound to a specific deployed contract.
func NewMarketViewerOPFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketViewerOPFilterer, error) {
	contract, err := bindMarketViewerOP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketViewerOPFilterer{contract: contract}, nil
}

// bindMarketViewerOP binds a generic wrapper to an already deployed contract.
func bindMarketViewerOP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketViewerOPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketViewerOP *MarketViewerOPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketViewerOP.Contract.MarketViewerOPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketViewerOP *MarketViewerOPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.MarketViewerOPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketViewerOP *MarketViewerOPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.MarketViewerOPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketViewerOP *MarketViewerOPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketViewerOP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketViewerOP *MarketViewerOPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketViewerOP *MarketViewerOPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.contract.Transact(opts, method, params...)
}

// GetBoard is a free data retrieval call binding the contract method 0x5d033f7f.
//
// Solidity: function getBoard(address market, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerOP *MarketViewerOPCaller) GetBoard(opts *bind.CallOpts, market common.Address, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getBoard", market, boardId)

	if err != nil {
		return *new(OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerBoardView)).(*OptionMarketViewerBoardView)

	return out0, err

}

// GetBoard is a free data retrieval call binding the contract method 0x5d033f7f.
//
// Solidity: function getBoard(address market, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerOP *MarketViewerOPSession) GetBoard(market common.Address, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerOP.Contract.GetBoard(&_MarketViewerOP.CallOpts, market, boardId)
}

// GetBoard is a free data retrieval call binding the contract method 0x5d033f7f.
//
// Solidity: function getBoard(address market, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerOP *MarketViewerOPCallerSession) GetBoard(market common.Address, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerOP.Contract.GetBoard(&_MarketViewerOP.CallOpts, market, boardId)
}

// GetBoardForBaseKey is a free data retrieval call binding the contract method 0xa9306dba.
//
// Solidity: function getBoardForBaseKey(bytes32 baseKey, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerOP *MarketViewerOPCaller) GetBoardForBaseKey(opts *bind.CallOpts, baseKey [32]byte, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getBoardForBaseKey", baseKey, boardId)

	if err != nil {
		return *new(OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerBoardView)).(*OptionMarketViewerBoardView)

	return out0, err

}

// GetBoardForBaseKey is a free data retrieval call binding the contract method 0xa9306dba.
//
// Solidity: function getBoardForBaseKey(bytes32 baseKey, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerOP *MarketViewerOPSession) GetBoardForBaseKey(baseKey [32]byte, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerOP.Contract.GetBoardForBaseKey(&_MarketViewerOP.CallOpts, baseKey, boardId)
}

// GetBoardForBaseKey is a free data retrieval call binding the contract method 0xa9306dba.
//
// Solidity: function getBoardForBaseKey(bytes32 baseKey, uint256 boardId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerOP *MarketViewerOPCallerSession) GetBoardForBaseKey(baseKey [32]byte, boardId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerOP.Contract.GetBoardForBaseKey(&_MarketViewerOP.CallOpts, baseKey, boardId)
}

// GetBoardForStrikeId is a free data retrieval call binding the contract method 0xac6e227b.
//
// Solidity: function getBoardForStrikeId(address market, uint256 strikeId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerOP *MarketViewerOPCaller) GetBoardForStrikeId(opts *bind.CallOpts, market common.Address, strikeId *big.Int) (OptionMarketViewerBoardView, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getBoardForStrikeId", market, strikeId)

	if err != nil {
		return *new(OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerBoardView)).(*OptionMarketViewerBoardView)

	return out0, err

}

// GetBoardForStrikeId is a free data retrieval call binding the contract method 0xac6e227b.
//
// Solidity: function getBoardForStrikeId(address market, uint256 strikeId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerOP *MarketViewerOPSession) GetBoardForStrikeId(market common.Address, strikeId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerOP.Contract.GetBoardForStrikeId(&_MarketViewerOP.CallOpts, market, strikeId)
}

// GetBoardForStrikeId is a free data retrieval call binding the contract method 0xac6e227b.
//
// Solidity: function getBoardForStrikeId(address market, uint256 strikeId) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[]))
func (_MarketViewerOP *MarketViewerOPCallerSession) GetBoardForStrikeId(market common.Address, strikeId *big.Int) (OptionMarketViewerBoardView, error) {
	return _MarketViewerOP.Contract.GetBoardForStrikeId(&_MarketViewerOP.CallOpts, market, strikeId)
}

// GetLiquidityBalancesAndAllowances is a free data retrieval call binding the contract method 0x25646545.
//
// Solidity: function getLiquidityBalancesAndAllowances(address[] markets, address account) view returns((address,uint256,uint256)[])
func (_MarketViewerOP *MarketViewerOPCaller) GetLiquidityBalancesAndAllowances(opts *bind.CallOpts, markets []common.Address, account common.Address) ([]OptionMarketViewerLiquidityBalanceAndAllowance, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getLiquidityBalancesAndAllowances", markets, account)

	if err != nil {
		return *new([]OptionMarketViewerLiquidityBalanceAndAllowance), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerLiquidityBalanceAndAllowance)).(*[]OptionMarketViewerLiquidityBalanceAndAllowance)

	return out0, err

}

// GetLiquidityBalancesAndAllowances is a free data retrieval call binding the contract method 0x25646545.
//
// Solidity: function getLiquidityBalancesAndAllowances(address[] markets, address account) view returns((address,uint256,uint256)[])
func (_MarketViewerOP *MarketViewerOPSession) GetLiquidityBalancesAndAllowances(markets []common.Address, account common.Address) ([]OptionMarketViewerLiquidityBalanceAndAllowance, error) {
	return _MarketViewerOP.Contract.GetLiquidityBalancesAndAllowances(&_MarketViewerOP.CallOpts, markets, account)
}

// GetLiquidityBalancesAndAllowances is a free data retrieval call binding the contract method 0x25646545.
//
// Solidity: function getLiquidityBalancesAndAllowances(address[] markets, address account) view returns((address,uint256,uint256)[])
func (_MarketViewerOP *MarketViewerOPCallerSession) GetLiquidityBalancesAndAllowances(markets []common.Address, account common.Address) ([]OptionMarketViewerLiquidityBalanceAndAllowance, error) {
	return _MarketViewerOP.Contract.GetLiquidityBalancesAndAllowances(&_MarketViewerOP.CallOpts, markets, account)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0xb0862c0e.
//
// Solidity: function getLiveBoards(address market) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[] marketBoards)
func (_MarketViewerOP *MarketViewerOPCaller) GetLiveBoards(opts *bind.CallOpts, market common.Address) ([]OptionMarketViewerBoardView, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getLiveBoards", market)

	if err != nil {
		return *new([]OptionMarketViewerBoardView), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerBoardView)).(*[]OptionMarketViewerBoardView)

	return out0, err

}

// GetLiveBoards is a free data retrieval call binding the contract method 0xb0862c0e.
//
// Solidity: function getLiveBoards(address market) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[] marketBoards)
func (_MarketViewerOP *MarketViewerOPSession) GetLiveBoards(market common.Address) ([]OptionMarketViewerBoardView, error) {
	return _MarketViewerOP.Contract.GetLiveBoards(&_MarketViewerOP.CallOpts, market)
}

// GetLiveBoards is a free data retrieval call binding the contract method 0xb0862c0e.
//
// Solidity: function getLiveBoards(address market) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[] marketBoards)
func (_MarketViewerOP *MarketViewerOPCallerSession) GetLiveBoards(market common.Address) ([]OptionMarketViewerBoardView, error) {
	return _MarketViewerOP.Contract.GetLiveBoards(&_MarketViewerOP.CallOpts, market)
}

// GetMarket is a free data retrieval call binding the contract method 0xd4dfadbf.
//
// Solidity: function getMarket(address market) view returns((bool,uint256,uint256,uint256,(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[],(uint256,bytes32,bytes32,uint256,uint256)))
func (_MarketViewerOP *MarketViewerOPCaller) GetMarket(opts *bind.CallOpts, market common.Address) (OptionMarketViewerMarketViewWithBoards, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getMarket", market)

	if err != nil {
		return *new(OptionMarketViewerMarketViewWithBoards), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerMarketViewWithBoards)).(*OptionMarketViewerMarketViewWithBoards)

	return out0, err

}

// GetMarket is a free data retrieval call binding the contract method 0xd4dfadbf.
//
// Solidity: function getMarket(address market) view returns((bool,uint256,uint256,uint256,(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[],(uint256,bytes32,bytes32,uint256,uint256)))
func (_MarketViewerOP *MarketViewerOPSession) GetMarket(market common.Address) (OptionMarketViewerMarketViewWithBoards, error) {
	return _MarketViewerOP.Contract.GetMarket(&_MarketViewerOP.CallOpts, market)
}

// GetMarket is a free data retrieval call binding the contract method 0xd4dfadbf.
//
// Solidity: function getMarket(address market) view returns((bool,uint256,uint256,uint256,(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[],(uint256,bytes32,bytes32,uint256,uint256)))
func (_MarketViewerOP *MarketViewerOPCallerSession) GetMarket(market common.Address) (OptionMarketViewerMarketViewWithBoards, error) {
	return _MarketViewerOP.Contract.GetMarket(&_MarketViewerOP.CallOpts, market)
}

// GetMarketAddresses is a free data retrieval call binding the contract method 0x97ce0d31.
//
// Solidity: function getMarketAddresses() view returns((address,address,address,address,address,address,address,address,address,address)[])
func (_MarketViewerOP *MarketViewerOPCaller) GetMarketAddresses(opts *bind.CallOpts) ([]OptionMarketViewerOptionMarketAddresses, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getMarketAddresses")

	if err != nil {
		return *new([]OptionMarketViewerOptionMarketAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerOptionMarketAddresses)).(*[]OptionMarketViewerOptionMarketAddresses)

	return out0, err

}

// GetMarketAddresses is a free data retrieval call binding the contract method 0x97ce0d31.
//
// Solidity: function getMarketAddresses() view returns((address,address,address,address,address,address,address,address,address,address)[])
func (_MarketViewerOP *MarketViewerOPSession) GetMarketAddresses() ([]OptionMarketViewerOptionMarketAddresses, error) {
	return _MarketViewerOP.Contract.GetMarketAddresses(&_MarketViewerOP.CallOpts)
}

// GetMarketAddresses is a free data retrieval call binding the contract method 0x97ce0d31.
//
// Solidity: function getMarketAddresses() view returns((address,address,address,address,address,address,address,address,address,address)[])
func (_MarketViewerOP *MarketViewerOPCallerSession) GetMarketAddresses() ([]OptionMarketViewerOptionMarketAddresses, error) {
	return _MarketViewerOP.Contract.GetMarketAddresses(&_MarketViewerOP.CallOpts)
}

// GetMarketForBaseKey is a free data retrieval call binding the contract method 0x6cc9023f.
//
// Solidity: function getMarketForBaseKey(bytes32 baseKey) view returns((bool,uint256,uint256,uint256,(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[],(uint256,bytes32,bytes32,uint256,uint256)) market)
func (_MarketViewerOP *MarketViewerOPCaller) GetMarketForBaseKey(opts *bind.CallOpts, baseKey [32]byte) (OptionMarketViewerMarketViewWithBoards, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getMarketForBaseKey", baseKey)

	if err != nil {
		return *new(OptionMarketViewerMarketViewWithBoards), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerMarketViewWithBoards)).(*OptionMarketViewerMarketViewWithBoards)

	return out0, err

}

// GetMarketForBaseKey is a free data retrieval call binding the contract method 0x6cc9023f.
//
// Solidity: function getMarketForBaseKey(bytes32 baseKey) view returns((bool,uint256,uint256,uint256,(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[],(uint256,bytes32,bytes32,uint256,uint256)) market)
func (_MarketViewerOP *MarketViewerOPSession) GetMarketForBaseKey(baseKey [32]byte) (OptionMarketViewerMarketViewWithBoards, error) {
	return _MarketViewerOP.Contract.GetMarketForBaseKey(&_MarketViewerOP.CallOpts, baseKey)
}

// GetMarketForBaseKey is a free data retrieval call binding the contract method 0x6cc9023f.
//
// Solidity: function getMarketForBaseKey(bytes32 baseKey) view returns((bool,uint256,uint256,uint256,(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,int256),(address,uint256,uint256,uint256,uint256,bool,uint256,(int256,int256,int256),(uint256,uint256,uint256,uint256,uint256,(int256,int256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256)[])[],(uint256,bytes32,bytes32,uint256,uint256)) market)
func (_MarketViewerOP *MarketViewerOPCallerSession) GetMarketForBaseKey(baseKey [32]byte) (OptionMarketViewerMarketViewWithBoards, error) {
	return _MarketViewerOP.Contract.GetMarketForBaseKey(&_MarketViewerOP.CallOpts, baseKey)
}

// GetMarkets is a free data retrieval call binding the contract method 0x1139e3f1.
//
// Solidity: function getMarkets(address[] markets) view returns((address,bool,(bool,uint256,uint256,uint256,(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,int256),(uint256,bytes32,bytes32,uint256,uint256))[]) marketsView)
func (_MarketViewerOP *MarketViewerOPCaller) GetMarkets(opts *bind.CallOpts, markets []common.Address) (OptionMarketViewerMarketsView, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getMarkets", markets)

	if err != nil {
		return *new(OptionMarketViewerMarketsView), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionMarketViewerMarketsView)).(*OptionMarketViewerMarketsView)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0x1139e3f1.
//
// Solidity: function getMarkets(address[] markets) view returns((address,bool,(bool,uint256,uint256,uint256,(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,int256),(uint256,bytes32,bytes32,uint256,uint256))[]) marketsView)
func (_MarketViewerOP *MarketViewerOPSession) GetMarkets(markets []common.Address) (OptionMarketViewerMarketsView, error) {
	return _MarketViewerOP.Contract.GetMarkets(&_MarketViewerOP.CallOpts, markets)
}

// GetMarkets is a free data retrieval call binding the contract method 0x1139e3f1.
//
// Solidity: function getMarkets(address[] markets) view returns((address,bool,(bool,uint256,uint256,uint256,(address,address,address,address,address,address,address,address,address,address),((uint256,address,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256,uint256,uint256),(int256,int256,int256),(uint256,bytes32,bytes32,uint256,uint256))[]) marketsView)
func (_MarketViewerOP *MarketViewerOPCallerSession) GetMarkets(markets []common.Address) (OptionMarketViewerMarketsView, error) {
	return _MarketViewerOP.Contract.GetMarkets(&_MarketViewerOP.CallOpts, markets)
}

// GetOwnerPositions is a free data retrieval call binding the contract method 0x845e1259.
//
// Solidity: function getOwnerPositions(address account) view returns((address,(uint256,uint256,uint8,uint256,uint256,uint8)[])[])
func (_MarketViewerOP *MarketViewerOPCaller) GetOwnerPositions(opts *bind.CallOpts, account common.Address) ([]OptionMarketViewerMarketOptionPositions, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getOwnerPositions", account)

	if err != nil {
		return *new([]OptionMarketViewerMarketOptionPositions), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionMarketViewerMarketOptionPositions)).(*[]OptionMarketViewerMarketOptionPositions)

	return out0, err

}

// GetOwnerPositions is a free data retrieval call binding the contract method 0x845e1259.
//
// Solidity: function getOwnerPositions(address account) view returns((address,(uint256,uint256,uint8,uint256,uint256,uint8)[])[])
func (_MarketViewerOP *MarketViewerOPSession) GetOwnerPositions(account common.Address) ([]OptionMarketViewerMarketOptionPositions, error) {
	return _MarketViewerOP.Contract.GetOwnerPositions(&_MarketViewerOP.CallOpts, account)
}

// GetOwnerPositions is a free data retrieval call binding the contract method 0x845e1259.
//
// Solidity: function getOwnerPositions(address account) view returns((address,(uint256,uint256,uint8,uint256,uint256,uint8)[])[])
func (_MarketViewerOP *MarketViewerOPCallerSession) GetOwnerPositions(account common.Address) ([]OptionMarketViewerMarketOptionPositions, error) {
	return _MarketViewerOP.Contract.GetOwnerPositions(&_MarketViewerOP.CallOpts, account)
}

// GetOwnerPositionsInRange is a free data retrieval call binding the contract method 0x9cdaffe2.
//
// Solidity: function getOwnerPositionsInRange(address market, address account, uint256 start, uint256 limit) view returns((uint256,uint256,uint8,uint256,uint256,uint8)[])
func (_MarketViewerOP *MarketViewerOPCaller) GetOwnerPositionsInRange(opts *bind.CallOpts, market common.Address, account common.Address, start *big.Int, limit *big.Int) ([]OptionTokenOptionPosition, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "getOwnerPositionsInRange", market, account, start, limit)

	if err != nil {
		return *new([]OptionTokenOptionPosition), err
	}

	out0 := *abi.ConvertType(out[0], new([]OptionTokenOptionPosition)).(*[]OptionTokenOptionPosition)

	return out0, err

}

// GetOwnerPositionsInRange is a free data retrieval call binding the contract method 0x9cdaffe2.
//
// Solidity: function getOwnerPositionsInRange(address market, address account, uint256 start, uint256 limit) view returns((uint256,uint256,uint8,uint256,uint256,uint8)[])
func (_MarketViewerOP *MarketViewerOPSession) GetOwnerPositionsInRange(market common.Address, account common.Address, start *big.Int, limit *big.Int) ([]OptionTokenOptionPosition, error) {
	return _MarketViewerOP.Contract.GetOwnerPositionsInRange(&_MarketViewerOP.CallOpts, market, account, start, limit)
}

// GetOwnerPositionsInRange is a free data retrieval call binding the contract method 0x9cdaffe2.
//
// Solidity: function getOwnerPositionsInRange(address market, address account, uint256 start, uint256 limit) view returns((uint256,uint256,uint8,uint256,uint256,uint8)[])
func (_MarketViewerOP *MarketViewerOPCallerSession) GetOwnerPositionsInRange(market common.Address, account common.Address, start *big.Int, limit *big.Int) ([]OptionTokenOptionPosition, error) {
	return _MarketViewerOP.Contract.GetOwnerPositionsInRange(&_MarketViewerOP.CallOpts, market, account, start, limit)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_MarketViewerOP *MarketViewerOPCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_MarketViewerOP *MarketViewerOPSession) Initialized() (bool, error) {
	return _MarketViewerOP.Contract.Initialized(&_MarketViewerOP.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_MarketViewerOP *MarketViewerOPCallerSession) Initialized() (bool, error) {
	return _MarketViewerOP.Contract.Initialized(&_MarketViewerOP.CallOpts)
}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address shortCollateral, address poolHedger, address quoteAsset, address baseAsset)
func (_MarketViewerOP *MarketViewerOPCaller) MarketAddresses(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	err := _MarketViewerOP.contract.Call(opts, &out, "marketAddresses", arg0)

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
func (_MarketViewerOP *MarketViewerOPSession) MarketAddresses(arg0 common.Address) (struct {
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
	return _MarketViewerOP.Contract.MarketAddresses(&_MarketViewerOP.CallOpts, arg0)
}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address shortCollateral, address poolHedger, address quoteAsset, address baseAsset)
func (_MarketViewerOP *MarketViewerOPCallerSession) MarketAddresses(arg0 common.Address) (struct {
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
	return _MarketViewerOP.Contract.MarketAddresses(&_MarketViewerOP.CallOpts, arg0)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_MarketViewerOP *MarketViewerOPCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_MarketViewerOP *MarketViewerOPSession) NominatedOwner() (common.Address, error) {
	return _MarketViewerOP.Contract.NominatedOwner(&_MarketViewerOP.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_MarketViewerOP *MarketViewerOPCallerSession) NominatedOwner() (common.Address, error) {
	return _MarketViewerOP.Contract.NominatedOwner(&_MarketViewerOP.CallOpts)
}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_MarketViewerOP *MarketViewerOPCaller) OptionMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "optionMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_MarketViewerOP *MarketViewerOPSession) OptionMarkets(arg0 *big.Int) (common.Address, error) {
	return _MarketViewerOP.Contract.OptionMarkets(&_MarketViewerOP.CallOpts, arg0)
}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_MarketViewerOP *MarketViewerOPCallerSession) OptionMarkets(arg0 *big.Int) (common.Address, error) {
	return _MarketViewerOP.Contract.OptionMarkets(&_MarketViewerOP.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketViewerOP *MarketViewerOPCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketViewerOP *MarketViewerOPSession) Owner() (common.Address, error) {
	return _MarketViewerOP.Contract.Owner(&_MarketViewerOP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketViewerOP *MarketViewerOPCallerSession) Owner() (common.Address, error) {
	return _MarketViewerOP.Contract.Owner(&_MarketViewerOP.CallOpts)
}

// SynthetixAdapter is a free data retrieval call binding the contract method 0xd2f5a5bb.
//
// Solidity: function synthetixAdapter() view returns(address)
func (_MarketViewerOP *MarketViewerOPCaller) SynthetixAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketViewerOP.contract.Call(opts, &out, "synthetixAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SynthetixAdapter is a free data retrieval call binding the contract method 0xd2f5a5bb.
//
// Solidity: function synthetixAdapter() view returns(address)
func (_MarketViewerOP *MarketViewerOPSession) SynthetixAdapter() (common.Address, error) {
	return _MarketViewerOP.Contract.SynthetixAdapter(&_MarketViewerOP.CallOpts)
}

// SynthetixAdapter is a free data retrieval call binding the contract method 0xd2f5a5bb.
//
// Solidity: function synthetixAdapter() view returns(address)
func (_MarketViewerOP *MarketViewerOPCallerSession) SynthetixAdapter() (common.Address, error) {
	return _MarketViewerOP.Contract.SynthetixAdapter(&_MarketViewerOP.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MarketViewerOP *MarketViewerOPTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketViewerOP.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MarketViewerOP *MarketViewerOPSession) AcceptOwnership() (*types.Transaction, error) {
	return _MarketViewerOP.Contract.AcceptOwnership(&_MarketViewerOP.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MarketViewerOP *MarketViewerOPTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MarketViewerOP.Contract.AcceptOwnership(&_MarketViewerOP.TransactOpts)
}

// AddMarket is a paid mutator transaction binding the contract method 0x7f68720f.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_MarketViewerOP *MarketViewerOPTransactor) AddMarket(opts *bind.TransactOpts, newMarketAddresses OptionMarketViewerOptionMarketAddresses) (*types.Transaction, error) {
	return _MarketViewerOP.contract.Transact(opts, "addMarket", newMarketAddresses)
}

// AddMarket is a paid mutator transaction binding the contract method 0x7f68720f.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_MarketViewerOP *MarketViewerOPSession) AddMarket(newMarketAddresses OptionMarketViewerOptionMarketAddresses) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.AddMarket(&_MarketViewerOP.TransactOpts, newMarketAddresses)
}

// AddMarket is a paid mutator transaction binding the contract method 0x7f68720f.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_MarketViewerOP *MarketViewerOPTransactorSession) AddMarket(newMarketAddresses OptionMarketViewerOptionMarketAddresses) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.AddMarket(&_MarketViewerOP.TransactOpts, newMarketAddresses)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _synthetixAdapter) returns()
func (_MarketViewerOP *MarketViewerOPTransactor) Init(opts *bind.TransactOpts, _synthetixAdapter common.Address) (*types.Transaction, error) {
	return _MarketViewerOP.contract.Transact(opts, "init", _synthetixAdapter)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _synthetixAdapter) returns()
func (_MarketViewerOP *MarketViewerOPSession) Init(_synthetixAdapter common.Address) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.Init(&_MarketViewerOP.TransactOpts, _synthetixAdapter)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _synthetixAdapter) returns()
func (_MarketViewerOP *MarketViewerOPTransactorSession) Init(_synthetixAdapter common.Address) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.Init(&_MarketViewerOP.TransactOpts, _synthetixAdapter)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_MarketViewerOP *MarketViewerOPTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _MarketViewerOP.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_MarketViewerOP *MarketViewerOPSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.NominateNewOwner(&_MarketViewerOP.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_MarketViewerOP *MarketViewerOPTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.NominateNewOwner(&_MarketViewerOP.TransactOpts, _owner)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_MarketViewerOP *MarketViewerOPTransactor) RemoveMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _MarketViewerOP.contract.Transact(opts, "removeMarket", market)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_MarketViewerOP *MarketViewerOPSession) RemoveMarket(market common.Address) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.RemoveMarket(&_MarketViewerOP.TransactOpts, market)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_MarketViewerOP *MarketViewerOPTransactorSession) RemoveMarket(market common.Address) (*types.Transaction, error) {
	return _MarketViewerOP.Contract.RemoveMarket(&_MarketViewerOP.TransactOpts, market)
}

// MarketViewerOPMarketAddedIterator is returned from FilterMarketAdded and is used to iterate over the raw logs and unpacked data for MarketAdded events raised by the MarketViewerOP contract.
type MarketViewerOPMarketAddedIterator struct {
	Event *MarketViewerOPMarketAdded // Event containing the contract specifics and raw log

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
func (it *MarketViewerOPMarketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketViewerOPMarketAdded)
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
		it.Event = new(MarketViewerOPMarketAdded)
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
func (it *MarketViewerOPMarketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketViewerOPMarketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketViewerOPMarketAdded represents a MarketAdded event raised by the MarketViewerOP contract.
type MarketViewerOPMarketAdded struct {
	Market OptionMarketViewerOptionMarketAddresses
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketAdded is a free log retrieval operation binding the contract event 0xd3777d9870fa8e7682a422021d48549c180f7b06a647478f5211caef6b6a4ee8.
//
// Solidity: event MarketAdded((address,address,address,address,address,address,address,address,address,address) market)
func (_MarketViewerOP *MarketViewerOPFilterer) FilterMarketAdded(opts *bind.FilterOpts) (*MarketViewerOPMarketAddedIterator, error) {

	logs, sub, err := _MarketViewerOP.contract.FilterLogs(opts, "MarketAdded")
	if err != nil {
		return nil, err
	}
	return &MarketViewerOPMarketAddedIterator{contract: _MarketViewerOP.contract, event: "MarketAdded", logs: logs, sub: sub}, nil
}

// WatchMarketAdded is a free log subscription operation binding the contract event 0xd3777d9870fa8e7682a422021d48549c180f7b06a647478f5211caef6b6a4ee8.
//
// Solidity: event MarketAdded((address,address,address,address,address,address,address,address,address,address) market)
func (_MarketViewerOP *MarketViewerOPFilterer) WatchMarketAdded(opts *bind.WatchOpts, sink chan<- *MarketViewerOPMarketAdded) (event.Subscription, error) {

	logs, sub, err := _MarketViewerOP.contract.WatchLogs(opts, "MarketAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketViewerOPMarketAdded)
				if err := _MarketViewerOP.contract.UnpackLog(event, "MarketAdded", log); err != nil {
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
func (_MarketViewerOP *MarketViewerOPFilterer) ParseMarketAdded(log types.Log) (*MarketViewerOPMarketAdded, error) {
	event := new(MarketViewerOPMarketAdded)
	if err := _MarketViewerOP.contract.UnpackLog(event, "MarketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketViewerOPMarketRemovedIterator is returned from FilterMarketRemoved and is used to iterate over the raw logs and unpacked data for MarketRemoved events raised by the MarketViewerOP contract.
type MarketViewerOPMarketRemovedIterator struct {
	Event *MarketViewerOPMarketRemoved // Event containing the contract specifics and raw log

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
func (it *MarketViewerOPMarketRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketViewerOPMarketRemoved)
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
		it.Event = new(MarketViewerOPMarketRemoved)
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
func (it *MarketViewerOPMarketRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketViewerOPMarketRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketViewerOPMarketRemoved represents a MarketRemoved event raised by the MarketViewerOP contract.
type MarketViewerOPMarketRemoved struct {
	Market common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketRemoved is a free log retrieval operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address market)
func (_MarketViewerOP *MarketViewerOPFilterer) FilterMarketRemoved(opts *bind.FilterOpts) (*MarketViewerOPMarketRemovedIterator, error) {

	logs, sub, err := _MarketViewerOP.contract.FilterLogs(opts, "MarketRemoved")
	if err != nil {
		return nil, err
	}
	return &MarketViewerOPMarketRemovedIterator{contract: _MarketViewerOP.contract, event: "MarketRemoved", logs: logs, sub: sub}, nil
}

// WatchMarketRemoved is a free log subscription operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address market)
func (_MarketViewerOP *MarketViewerOPFilterer) WatchMarketRemoved(opts *bind.WatchOpts, sink chan<- *MarketViewerOPMarketRemoved) (event.Subscription, error) {

	logs, sub, err := _MarketViewerOP.contract.WatchLogs(opts, "MarketRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketViewerOPMarketRemoved)
				if err := _MarketViewerOP.contract.UnpackLog(event, "MarketRemoved", log); err != nil {
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
func (_MarketViewerOP *MarketViewerOPFilterer) ParseMarketRemoved(log types.Log) (*MarketViewerOPMarketRemoved, error) {
	event := new(MarketViewerOPMarketRemoved)
	if err := _MarketViewerOP.contract.UnpackLog(event, "MarketRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketViewerOPOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the MarketViewerOP contract.
type MarketViewerOPOwnerChangedIterator struct {
	Event *MarketViewerOPOwnerChanged // Event containing the contract specifics and raw log

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
func (it *MarketViewerOPOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketViewerOPOwnerChanged)
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
		it.Event = new(MarketViewerOPOwnerChanged)
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
func (it *MarketViewerOPOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketViewerOPOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketViewerOPOwnerChanged represents a OwnerChanged event raised by the MarketViewerOP contract.
type MarketViewerOPOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_MarketViewerOP *MarketViewerOPFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*MarketViewerOPOwnerChangedIterator, error) {

	logs, sub, err := _MarketViewerOP.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &MarketViewerOPOwnerChangedIterator{contract: _MarketViewerOP.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_MarketViewerOP *MarketViewerOPFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *MarketViewerOPOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _MarketViewerOP.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketViewerOPOwnerChanged)
				if err := _MarketViewerOP.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_MarketViewerOP *MarketViewerOPFilterer) ParseOwnerChanged(log types.Log) (*MarketViewerOPOwnerChanged, error) {
	event := new(MarketViewerOPOwnerChanged)
	if err := _MarketViewerOP.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketViewerOPOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the MarketViewerOP contract.
type MarketViewerOPOwnerNominatedIterator struct {
	Event *MarketViewerOPOwnerNominated // Event containing the contract specifics and raw log

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
func (it *MarketViewerOPOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketViewerOPOwnerNominated)
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
		it.Event = new(MarketViewerOPOwnerNominated)
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
func (it *MarketViewerOPOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketViewerOPOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketViewerOPOwnerNominated represents a OwnerNominated event raised by the MarketViewerOP contract.
type MarketViewerOPOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_MarketViewerOP *MarketViewerOPFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*MarketViewerOPOwnerNominatedIterator, error) {

	logs, sub, err := _MarketViewerOP.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &MarketViewerOPOwnerNominatedIterator{contract: _MarketViewerOP.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_MarketViewerOP *MarketViewerOPFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *MarketViewerOPOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _MarketViewerOP.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketViewerOPOwnerNominated)
				if err := _MarketViewerOP.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_MarketViewerOP *MarketViewerOPFilterer) ParseOwnerNominated(log types.Log) (*MarketViewerOPOwnerNominated, error) {
	event := new(MarketViewerOPOwnerNominated)
	if err := _MarketViewerOP.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
