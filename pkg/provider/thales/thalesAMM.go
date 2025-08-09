// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package thales

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

// ThalesMetaData contains all meta data concerning the Thales contract.
var ThalesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sUSDPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"susd\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"BoughtFromAmm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PauseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refferer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"volume\",\"type\":\"uint256\"}],\"name\":\"ReferrerPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"SetCapPerAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_capPerMarket\",\"type\":\"uint256\"}],\"name\":\"SetCapPerMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_impliedVolatility\",\"type\":\"uint256\"}],\"name\":\"SetImpliedVolatilityPerAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_spread\",\"type\":\"uint256\"}],\"name\":\"SetMaxSpread\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_spread\",\"type\":\"uint256\"}],\"name\":\"SetMaxSupportedPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_spread\",\"type\":\"uint256\"}],\"name\":\"SetMinSpread\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_spread\",\"type\":\"uint256\"}],\"name\":\"SetMinSupportedPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minimalTimeLeftToMaturity\",\"type\":\"uint256\"}],\"name\":\"SetMinimalTimeLeftToMaturity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"SetPositionalMarketManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"SetPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sUSD\",\"type\":\"address\"}],\"name\":\"SetSUSD\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_safeBox\",\"type\":\"address\"}],\"name\":\"SetSafeBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_safeBoxImpact\",\"type\":\"uint256\"}],\"name\":\"SetSafeBoxImpact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sUSDPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"susd\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"SoldToAMM\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_APPROVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"}],\"name\":\"availableToBuyFromAMM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_available\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"}],\"name\":\"availableToSellToAMM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPayout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalSlippage\",\"type\":\"uint256\"}],\"name\":\"buyFromAMM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPayout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalSlippage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"buyFromAMMWithDifferentCollateralAndReferrer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPayout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalSlippage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"buyFromAMMWithReferrer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buyFromAmmQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"buyFromAmmQuoteWithDifferentCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sUSDToPay\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buyPriceImpact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeLeftInDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"volatility\",\"type\":\"uint256\"}],\"name\":\"calculateOdds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"canExerciseMaturedMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_canExercise\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"capPerMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"curveOnrampEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"curveSUSD\",\"outputs\":[{\"internalType\":\"contractICurveSUSD\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deciMath\",\"outputs\":[{\"internalType\":\"contractDeciMath\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"exerciseMaturedMarket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"}],\"name\":\"getCapPerAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"impliedVolatilityPerAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initNonReentrant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIPriceFeed\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_sUSD\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_capPerMarket\",\"type\":\"uint256\"},{\"internalType\":\"contractDeciMath\",\"name\":\"_deciMath\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_min_spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max_spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimalTimeLeftToMaturity\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"isMarketInAMMTrading\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isTrading\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastPauseTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxSupportedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"max_spread\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minSupportedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"min_spread\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimalTimeLeftToMaturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"referrals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"referrerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"retrieveSUSDAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sUSD\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"safeBox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"safeBoxImpact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sellPriceImpact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_impact\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPayout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalSlippage\",\"type\":\"uint256\"}],\"name\":\"sellToAMM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"enumThalesAMM.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sellToAmmQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_available\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"setCapPerAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_curveSUSD\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dai\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_curveOnrampEnabled\",\"type\":\"bool\"}],\"name\":\"setCurveSUSD\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_impliedVolatility\",\"type\":\"uint256\"}],\"name\":\"setImpliedVolatilityPerAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minspread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxspread\",\"type\":\"uint256\"}],\"name\":\"setMinMaxSpread\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSupportedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupportedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_capPerMarket\",\"type\":\"uint256\"}],\"name\":\"setMinMaxSupportedPriceAndCap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimalTimeLeftToMaturity\",\"type\":\"uint256\"}],\"name\":\"setMinimalTimeLeftToMaturity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setPositionalMarketManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_sUSD\",\"type\":\"address\"}],\"name\":\"setPriceFeedAndSUSD\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_safeBox\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_safeBoxImpact\",\"type\":\"uint256\"}],\"name\":\"setSafeBoxData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIStakingThales\",\"name\":\"_stakingThales\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_referrals\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_referrerFee\",\"type\":\"uint256\"}],\"name\":\"setStakingThalesAndReferrals\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"spentOnMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingThales\",\"outputs\":[{\"internalType\":\"contractIStakingThales\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"}],\"name\":\"transferOwnershipAtInit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ThalesABI is the input ABI used to generate the binding from.
// Deprecated: Use ThalesMetaData.ABI instead.
var ThalesABI = ThalesMetaData.ABI

// Thales is an auto generated Go binding around an Ethereum contract.
type Thales struct {
	ThalesCaller     // Read-only binding to the contract
	ThalesTransactor // Write-only binding to the contract
	ThalesFilterer   // Log filterer for contract events
}

// ThalesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThalesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThalesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThalesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThalesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThalesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThalesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThalesSession struct {
	Contract     *Thales           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThalesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThalesCallerSession struct {
	Contract *ThalesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ThalesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThalesTransactorSession struct {
	Contract     *ThalesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThalesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThalesRaw struct {
	Contract *Thales // Generic contract binding to access the raw methods on
}

// ThalesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThalesCallerRaw struct {
	Contract *ThalesCaller // Generic read-only contract binding to access the raw methods on
}

// ThalesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThalesTransactorRaw struct {
	Contract *ThalesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThales creates a new instance of Thales, bound to a specific deployed contract.
func NewThales(address common.Address, backend bind.ContractBackend) (*Thales, error) {
	contract, err := bindThales(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Thales{ThalesCaller: ThalesCaller{contract: contract}, ThalesTransactor: ThalesTransactor{contract: contract}, ThalesFilterer: ThalesFilterer{contract: contract}}, nil
}

// NewThalesCaller creates a new read-only instance of Thales, bound to a specific deployed contract.
func NewThalesCaller(address common.Address, caller bind.ContractCaller) (*ThalesCaller, error) {
	contract, err := bindThales(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThalesCaller{contract: contract}, nil
}

// NewThalesTransactor creates a new write-only instance of Thales, bound to a specific deployed contract.
func NewThalesTransactor(address common.Address, transactor bind.ContractTransactor) (*ThalesTransactor, error) {
	contract, err := bindThales(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThalesTransactor{contract: contract}, nil
}

// NewThalesFilterer creates a new log filterer instance of Thales, bound to a specific deployed contract.
func NewThalesFilterer(address common.Address, filterer bind.ContractFilterer) (*ThalesFilterer, error) {
	contract, err := bindThales(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThalesFilterer{contract: contract}, nil
}

// bindThales binds a generic wrapper to an already deployed contract.
func bindThales(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ThalesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thales *ThalesRaw) Call(opts *bind.CallOpts, result *[]any, method string, params ...any) error {
	return _Thales.Contract.ThalesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thales *ThalesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thales.Contract.ThalesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thales *ThalesRaw) Transact(opts *bind.TransactOpts, method string, params ...any) (*types.Transaction, error) {
	return _Thales.Contract.ThalesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thales *ThalesCallerRaw) Call(opts *bind.CallOpts, result *[]any, method string, params ...any) error {
	return _Thales.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thales *ThalesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thales.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thales *ThalesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...any) (*types.Transaction, error) {
	return _Thales.Contract.contract.Transact(opts, method, params...)
}

// MAXAPPROVAL is a free data retrieval call binding the contract method 0x7f4ed8ee.
//
// Solidity: function MAX_APPROVAL() view returns(uint256)
func (_Thales *ThalesCaller) MAXAPPROVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "MAX_APPROVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXAPPROVAL is a free data retrieval call binding the contract method 0x7f4ed8ee.
//
// Solidity: function MAX_APPROVAL() view returns(uint256)
func (_Thales *ThalesSession) MAXAPPROVAL() (*big.Int, error) {
	return _Thales.Contract.MAXAPPROVAL(&_Thales.CallOpts)
}

// MAXAPPROVAL is a free data retrieval call binding the contract method 0x7f4ed8ee.
//
// Solidity: function MAX_APPROVAL() view returns(uint256)
func (_Thales *ThalesCallerSession) MAXAPPROVAL() (*big.Int, error) {
	return _Thales.Contract.MAXAPPROVAL(&_Thales.CallOpts)
}

// AvailableToBuyFromAMM is a free data retrieval call binding the contract method 0xefc15251.
//
// Solidity: function availableToBuyFromAMM(address market, uint8 position) view returns(uint256 _available)
func (_Thales *ThalesCaller) AvailableToBuyFromAMM(opts *bind.CallOpts, market common.Address, position uint8) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "availableToBuyFromAMM", market, position)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableToBuyFromAMM is a free data retrieval call binding the contract method 0xefc15251.
//
// Solidity: function availableToBuyFromAMM(address market, uint8 position) view returns(uint256 _available)
func (_Thales *ThalesSession) AvailableToBuyFromAMM(market common.Address, position uint8) (*big.Int, error) {
	return _Thales.Contract.AvailableToBuyFromAMM(&_Thales.CallOpts, market, position)
}

// AvailableToBuyFromAMM is a free data retrieval call binding the contract method 0xefc15251.
//
// Solidity: function availableToBuyFromAMM(address market, uint8 position) view returns(uint256 _available)
func (_Thales *ThalesCallerSession) AvailableToBuyFromAMM(market common.Address, position uint8) (*big.Int, error) {
	return _Thales.Contract.AvailableToBuyFromAMM(&_Thales.CallOpts, market, position)
}

// AvailableToSellToAMM is a free data retrieval call binding the contract method 0xc2783f92.
//
// Solidity: function availableToSellToAMM(address market, uint8 position) view returns(uint256)
func (_Thales *ThalesCaller) AvailableToSellToAMM(opts *bind.CallOpts, market common.Address, position uint8) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "availableToSellToAMM", market, position)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableToSellToAMM is a free data retrieval call binding the contract method 0xc2783f92.
//
// Solidity: function availableToSellToAMM(address market, uint8 position) view returns(uint256)
func (_Thales *ThalesSession) AvailableToSellToAMM(market common.Address, position uint8) (*big.Int, error) {
	return _Thales.Contract.AvailableToSellToAMM(&_Thales.CallOpts, market, position)
}

// AvailableToSellToAMM is a free data retrieval call binding the contract method 0xc2783f92.
//
// Solidity: function availableToSellToAMM(address market, uint8 position) view returns(uint256)
func (_Thales *ThalesCallerSession) AvailableToSellToAMM(market common.Address, position uint8) (*big.Int, error) {
	return _Thales.Contract.AvailableToSellToAMM(&_Thales.CallOpts, market, position)
}

// BuyFromAmmQuote is a free data retrieval call binding the contract method 0x270e13ef.
//
// Solidity: function buyFromAmmQuote(address market, uint8 position, uint256 amount) view returns(uint256)
func (_Thales *ThalesCaller) BuyFromAmmQuote(opts *bind.CallOpts, market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "buyFromAmmQuote", market, position, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyFromAmmQuote is a free data retrieval call binding the contract method 0x270e13ef.
//
// Solidity: function buyFromAmmQuote(address market, uint8 position, uint256 amount) view returns(uint256)
func (_Thales *ThalesSession) BuyFromAmmQuote(market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	return _Thales.Contract.BuyFromAmmQuote(&_Thales.CallOpts, market, position, amount)
}

// BuyFromAmmQuote is a free data retrieval call binding the contract method 0x270e13ef.
//
// Solidity: function buyFromAmmQuote(address market, uint8 position, uint256 amount) view returns(uint256)
func (_Thales *ThalesCallerSession) BuyFromAmmQuote(market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	return _Thales.Contract.BuyFromAmmQuote(&_Thales.CallOpts, market, position, amount)
}

// BuyFromAmmQuoteWithDifferentCollateral is a free data retrieval call binding the contract method 0xd4a2641b.
//
// Solidity: function buyFromAmmQuoteWithDifferentCollateral(address market, uint8 position, uint256 amount, address collateral) view returns(uint256 collateralQuote, uint256 sUSDToPay)
func (_Thales *ThalesCaller) BuyFromAmmQuoteWithDifferentCollateral(opts *bind.CallOpts, market common.Address, position uint8, amount *big.Int, collateral common.Address) (struct {
	CollateralQuote *big.Int
	SUSDToPay       *big.Int
}, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "buyFromAmmQuoteWithDifferentCollateral", market, position, amount, collateral)

	outstruct := new(struct {
		CollateralQuote *big.Int
		SUSDToPay       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralQuote = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SUSDToPay = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BuyFromAmmQuoteWithDifferentCollateral is a free data retrieval call binding the contract method 0xd4a2641b.
//
// Solidity: function buyFromAmmQuoteWithDifferentCollateral(address market, uint8 position, uint256 amount, address collateral) view returns(uint256 collateralQuote, uint256 sUSDToPay)
func (_Thales *ThalesSession) BuyFromAmmQuoteWithDifferentCollateral(market common.Address, position uint8, amount *big.Int, collateral common.Address) (struct {
	CollateralQuote *big.Int
	SUSDToPay       *big.Int
}, error) {
	return _Thales.Contract.BuyFromAmmQuoteWithDifferentCollateral(&_Thales.CallOpts, market, position, amount, collateral)
}

// BuyFromAmmQuoteWithDifferentCollateral is a free data retrieval call binding the contract method 0xd4a2641b.
//
// Solidity: function buyFromAmmQuoteWithDifferentCollateral(address market, uint8 position, uint256 amount, address collateral) view returns(uint256 collateralQuote, uint256 sUSDToPay)
func (_Thales *ThalesCallerSession) BuyFromAmmQuoteWithDifferentCollateral(market common.Address, position uint8, amount *big.Int, collateral common.Address) (struct {
	CollateralQuote *big.Int
	SUSDToPay       *big.Int
}, error) {
	return _Thales.Contract.BuyFromAmmQuoteWithDifferentCollateral(&_Thales.CallOpts, market, position, amount, collateral)
}

// BuyPriceImpact is a free data retrieval call binding the contract method 0xbf46c0b4.
//
// Solidity: function buyPriceImpact(address market, uint8 position, uint256 amount) view returns(uint256)
func (_Thales *ThalesCaller) BuyPriceImpact(opts *bind.CallOpts, market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "buyPriceImpact", market, position, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyPriceImpact is a free data retrieval call binding the contract method 0xbf46c0b4.
//
// Solidity: function buyPriceImpact(address market, uint8 position, uint256 amount) view returns(uint256)
func (_Thales *ThalesSession) BuyPriceImpact(market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	return _Thales.Contract.BuyPriceImpact(&_Thales.CallOpts, market, position, amount)
}

// BuyPriceImpact is a free data retrieval call binding the contract method 0xbf46c0b4.
//
// Solidity: function buyPriceImpact(address market, uint8 position, uint256 amount) view returns(uint256)
func (_Thales *ThalesCallerSession) BuyPriceImpact(market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	return _Thales.Contract.BuyPriceImpact(&_Thales.CallOpts, market, position, amount)
}

// CalculateOdds is a free data retrieval call binding the contract method 0xa8cd06e8.
//
// Solidity: function calculateOdds(uint256 _price, uint256 strike, uint256 timeLeftInDays, uint256 volatility) view returns(uint256)
func (_Thales *ThalesCaller) CalculateOdds(opts *bind.CallOpts, _price *big.Int, strike *big.Int, timeLeftInDays *big.Int, volatility *big.Int) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "calculateOdds", _price, strike, timeLeftInDays, volatility)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateOdds is a free data retrieval call binding the contract method 0xa8cd06e8.
//
// Solidity: function calculateOdds(uint256 _price, uint256 strike, uint256 timeLeftInDays, uint256 volatility) view returns(uint256)
func (_Thales *ThalesSession) CalculateOdds(_price *big.Int, strike *big.Int, timeLeftInDays *big.Int, volatility *big.Int) (*big.Int, error) {
	return _Thales.Contract.CalculateOdds(&_Thales.CallOpts, _price, strike, timeLeftInDays, volatility)
}

// CalculateOdds is a free data retrieval call binding the contract method 0xa8cd06e8.
//
// Solidity: function calculateOdds(uint256 _price, uint256 strike, uint256 timeLeftInDays, uint256 volatility) view returns(uint256)
func (_Thales *ThalesCallerSession) CalculateOdds(_price *big.Int, strike *big.Int, timeLeftInDays *big.Int, volatility *big.Int) (*big.Int, error) {
	return _Thales.Contract.CalculateOdds(&_Thales.CallOpts, _price, strike, timeLeftInDays, volatility)
}

// CanExerciseMaturedMarket is a free data retrieval call binding the contract method 0xca1d578e.
//
// Solidity: function canExerciseMaturedMarket(address market) view returns(bool _canExercise)
func (_Thales *ThalesCaller) CanExerciseMaturedMarket(opts *bind.CallOpts, market common.Address) (bool, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "canExerciseMaturedMarket", market)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanExerciseMaturedMarket is a free data retrieval call binding the contract method 0xca1d578e.
//
// Solidity: function canExerciseMaturedMarket(address market) view returns(bool _canExercise)
func (_Thales *ThalesSession) CanExerciseMaturedMarket(market common.Address) (bool, error) {
	return _Thales.Contract.CanExerciseMaturedMarket(&_Thales.CallOpts, market)
}

// CanExerciseMaturedMarket is a free data retrieval call binding the contract method 0xca1d578e.
//
// Solidity: function canExerciseMaturedMarket(address market) view returns(bool _canExercise)
func (_Thales *ThalesCallerSession) CanExerciseMaturedMarket(market common.Address) (bool, error) {
	return _Thales.Contract.CanExerciseMaturedMarket(&_Thales.CallOpts, market)
}

// CapPerMarket is a free data retrieval call binding the contract method 0xfb91d41c.
//
// Solidity: function capPerMarket() view returns(uint256)
func (_Thales *ThalesCaller) CapPerMarket(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "capPerMarket")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CapPerMarket is a free data retrieval call binding the contract method 0xfb91d41c.
//
// Solidity: function capPerMarket() view returns(uint256)
func (_Thales *ThalesSession) CapPerMarket() (*big.Int, error) {
	return _Thales.Contract.CapPerMarket(&_Thales.CallOpts)
}

// CapPerMarket is a free data retrieval call binding the contract method 0xfb91d41c.
//
// Solidity: function capPerMarket() view returns(uint256)
func (_Thales *ThalesCallerSession) CapPerMarket() (*big.Int, error) {
	return _Thales.Contract.CapPerMarket(&_Thales.CallOpts)
}

// CurveOnrampEnabled is a free data retrieval call binding the contract method 0x1fbb38e8.
//
// Solidity: function curveOnrampEnabled() view returns(bool)
func (_Thales *ThalesCaller) CurveOnrampEnabled(opts *bind.CallOpts) (bool, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "curveOnrampEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CurveOnrampEnabled is a free data retrieval call binding the contract method 0x1fbb38e8.
//
// Solidity: function curveOnrampEnabled() view returns(bool)
func (_Thales *ThalesSession) CurveOnrampEnabled() (bool, error) {
	return _Thales.Contract.CurveOnrampEnabled(&_Thales.CallOpts)
}

// CurveOnrampEnabled is a free data retrieval call binding the contract method 0x1fbb38e8.
//
// Solidity: function curveOnrampEnabled() view returns(bool)
func (_Thales *ThalesCallerSession) CurveOnrampEnabled() (bool, error) {
	return _Thales.Contract.CurveOnrampEnabled(&_Thales.CallOpts)
}

// CurveSUSD is a free data retrieval call binding the contract method 0xa5bf660d.
//
// Solidity: function curveSUSD() view returns(address)
func (_Thales *ThalesCaller) CurveSUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "curveSUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurveSUSD is a free data retrieval call binding the contract method 0xa5bf660d.
//
// Solidity: function curveSUSD() view returns(address)
func (_Thales *ThalesSession) CurveSUSD() (common.Address, error) {
	return _Thales.Contract.CurveSUSD(&_Thales.CallOpts)
}

// CurveSUSD is a free data retrieval call binding the contract method 0xa5bf660d.
//
// Solidity: function curveSUSD() view returns(address)
func (_Thales *ThalesCallerSession) CurveSUSD() (common.Address, error) {
	return _Thales.Contract.CurveSUSD(&_Thales.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_Thales *ThalesCaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_Thales *ThalesSession) Dai() (common.Address, error) {
	return _Thales.Contract.Dai(&_Thales.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_Thales *ThalesCallerSession) Dai() (common.Address, error) {
	return _Thales.Contract.Dai(&_Thales.CallOpts)
}

// DeciMath is a free data retrieval call binding the contract method 0x6cfa636e.
//
// Solidity: function deciMath() view returns(address)
func (_Thales *ThalesCaller) DeciMath(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "deciMath")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeciMath is a free data retrieval call binding the contract method 0x6cfa636e.
//
// Solidity: function deciMath() view returns(address)
func (_Thales *ThalesSession) DeciMath() (common.Address, error) {
	return _Thales.Contract.DeciMath(&_Thales.CallOpts)
}

// DeciMath is a free data retrieval call binding the contract method 0x6cfa636e.
//
// Solidity: function deciMath() view returns(address)
func (_Thales *ThalesCallerSession) DeciMath() (common.Address, error) {
	return _Thales.Contract.DeciMath(&_Thales.CallOpts)
}

// GetCapPerAsset is a free data retrieval call binding the contract method 0xf8debeb7.
//
// Solidity: function getCapPerAsset(bytes32 asset) view returns(uint256 _cap)
func (_Thales *ThalesCaller) GetCapPerAsset(opts *bind.CallOpts, asset [32]byte) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "getCapPerAsset", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCapPerAsset is a free data retrieval call binding the contract method 0xf8debeb7.
//
// Solidity: function getCapPerAsset(bytes32 asset) view returns(uint256 _cap)
func (_Thales *ThalesSession) GetCapPerAsset(asset [32]byte) (*big.Int, error) {
	return _Thales.Contract.GetCapPerAsset(&_Thales.CallOpts, asset)
}

// GetCapPerAsset is a free data retrieval call binding the contract method 0xf8debeb7.
//
// Solidity: function getCapPerAsset(bytes32 asset) view returns(uint256 _cap)
func (_Thales *ThalesCallerSession) GetCapPerAsset(asset [32]byte) (*big.Int, error) {
	return _Thales.Contract.GetCapPerAsset(&_Thales.CallOpts, asset)
}

// ImpliedVolatilityPerAsset is a free data retrieval call binding the contract method 0xf502b003.
//
// Solidity: function impliedVolatilityPerAsset(bytes32 ) view returns(uint256)
func (_Thales *ThalesCaller) ImpliedVolatilityPerAsset(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "impliedVolatilityPerAsset", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ImpliedVolatilityPerAsset is a free data retrieval call binding the contract method 0xf502b003.
//
// Solidity: function impliedVolatilityPerAsset(bytes32 ) view returns(uint256)
func (_Thales *ThalesSession) ImpliedVolatilityPerAsset(arg0 [32]byte) (*big.Int, error) {
	return _Thales.Contract.ImpliedVolatilityPerAsset(&_Thales.CallOpts, arg0)
}

// ImpliedVolatilityPerAsset is a free data retrieval call binding the contract method 0xf502b003.
//
// Solidity: function impliedVolatilityPerAsset(bytes32 ) view returns(uint256)
func (_Thales *ThalesCallerSession) ImpliedVolatilityPerAsset(arg0 [32]byte) (*big.Int, error) {
	return _Thales.Contract.ImpliedVolatilityPerAsset(&_Thales.CallOpts, arg0)
}

// IsMarketInAMMTrading is a free data retrieval call binding the contract method 0x5727a0f3.
//
// Solidity: function isMarketInAMMTrading(address market) view returns(bool isTrading)
func (_Thales *ThalesCaller) IsMarketInAMMTrading(opts *bind.CallOpts, market common.Address) (bool, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "isMarketInAMMTrading", market)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketInAMMTrading is a free data retrieval call binding the contract method 0x5727a0f3.
//
// Solidity: function isMarketInAMMTrading(address market) view returns(bool isTrading)
func (_Thales *ThalesSession) IsMarketInAMMTrading(market common.Address) (bool, error) {
	return _Thales.Contract.IsMarketInAMMTrading(&_Thales.CallOpts, market)
}

// IsMarketInAMMTrading is a free data retrieval call binding the contract method 0x5727a0f3.
//
// Solidity: function isMarketInAMMTrading(address market) view returns(bool isTrading)
func (_Thales *ThalesCallerSession) IsMarketInAMMTrading(market common.Address) (bool, error) {
	return _Thales.Contract.IsMarketInAMMTrading(&_Thales.CallOpts, market)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Thales *ThalesCaller) LastPauseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "lastPauseTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Thales *ThalesSession) LastPauseTime() (*big.Int, error) {
	return _Thales.Contract.LastPauseTime(&_Thales.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Thales *ThalesCallerSession) LastPauseTime() (*big.Int, error) {
	return _Thales.Contract.LastPauseTime(&_Thales.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Thales *ThalesCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Thales *ThalesSession) Manager() (common.Address, error) {
	return _Thales.Contract.Manager(&_Thales.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Thales *ThalesCallerSession) Manager() (common.Address, error) {
	return _Thales.Contract.Manager(&_Thales.CallOpts)
}

// MaxSupportedPrice is a free data retrieval call binding the contract method 0x19b844a6.
//
// Solidity: function maxSupportedPrice() view returns(uint256)
func (_Thales *ThalesCaller) MaxSupportedPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "maxSupportedPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupportedPrice is a free data retrieval call binding the contract method 0x19b844a6.
//
// Solidity: function maxSupportedPrice() view returns(uint256)
func (_Thales *ThalesSession) MaxSupportedPrice() (*big.Int, error) {
	return _Thales.Contract.MaxSupportedPrice(&_Thales.CallOpts)
}

// MaxSupportedPrice is a free data retrieval call binding the contract method 0x19b844a6.
//
// Solidity: function maxSupportedPrice() view returns(uint256)
func (_Thales *ThalesCallerSession) MaxSupportedPrice() (*big.Int, error) {
	return _Thales.Contract.MaxSupportedPrice(&_Thales.CallOpts)
}

// MaxSpread is a free data retrieval call binding the contract method 0x316425c3.
//
// Solidity: function max_spread() view returns(uint256)
func (_Thales *ThalesCaller) MaxSpread(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "max_spread")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSpread is a free data retrieval call binding the contract method 0x316425c3.
//
// Solidity: function max_spread() view returns(uint256)
func (_Thales *ThalesSession) MaxSpread() (*big.Int, error) {
	return _Thales.Contract.MaxSpread(&_Thales.CallOpts)
}

// MaxSpread is a free data retrieval call binding the contract method 0x316425c3.
//
// Solidity: function max_spread() view returns(uint256)
func (_Thales *ThalesCallerSession) MaxSpread() (*big.Int, error) {
	return _Thales.Contract.MaxSpread(&_Thales.CallOpts)
}

// MinSupportedPrice is a free data retrieval call binding the contract method 0x931b2040.
//
// Solidity: function minSupportedPrice() view returns(uint256)
func (_Thales *ThalesCaller) MinSupportedPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "minSupportedPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSupportedPrice is a free data retrieval call binding the contract method 0x931b2040.
//
// Solidity: function minSupportedPrice() view returns(uint256)
func (_Thales *ThalesSession) MinSupportedPrice() (*big.Int, error) {
	return _Thales.Contract.MinSupportedPrice(&_Thales.CallOpts)
}

// MinSupportedPrice is a free data retrieval call binding the contract method 0x931b2040.
//
// Solidity: function minSupportedPrice() view returns(uint256)
func (_Thales *ThalesCallerSession) MinSupportedPrice() (*big.Int, error) {
	return _Thales.Contract.MinSupportedPrice(&_Thales.CallOpts)
}

// MinSpread is a free data retrieval call binding the contract method 0x6aaa81b6.
//
// Solidity: function min_spread() view returns(uint256)
func (_Thales *ThalesCaller) MinSpread(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "min_spread")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSpread is a free data retrieval call binding the contract method 0x6aaa81b6.
//
// Solidity: function min_spread() view returns(uint256)
func (_Thales *ThalesSession) MinSpread() (*big.Int, error) {
	return _Thales.Contract.MinSpread(&_Thales.CallOpts)
}

// MinSpread is a free data retrieval call binding the contract method 0x6aaa81b6.
//
// Solidity: function min_spread() view returns(uint256)
func (_Thales *ThalesCallerSession) MinSpread() (*big.Int, error) {
	return _Thales.Contract.MinSpread(&_Thales.CallOpts)
}

// MinimalTimeLeftToMaturity is a free data retrieval call binding the contract method 0xdf8974d0.
//
// Solidity: function minimalTimeLeftToMaturity() view returns(uint256)
func (_Thales *ThalesCaller) MinimalTimeLeftToMaturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "minimalTimeLeftToMaturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimalTimeLeftToMaturity is a free data retrieval call binding the contract method 0xdf8974d0.
//
// Solidity: function minimalTimeLeftToMaturity() view returns(uint256)
func (_Thales *ThalesSession) MinimalTimeLeftToMaturity() (*big.Int, error) {
	return _Thales.Contract.MinimalTimeLeftToMaturity(&_Thales.CallOpts)
}

// MinimalTimeLeftToMaturity is a free data retrieval call binding the contract method 0xdf8974d0.
//
// Solidity: function minimalTimeLeftToMaturity() view returns(uint256)
func (_Thales *ThalesCallerSession) MinimalTimeLeftToMaturity() (*big.Int, error) {
	return _Thales.Contract.MinimalTimeLeftToMaturity(&_Thales.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Thales *ThalesCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Thales *ThalesSession) NominatedOwner() (common.Address, error) {
	return _Thales.Contract.NominatedOwner(&_Thales.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Thales *ThalesCallerSession) NominatedOwner() (common.Address, error) {
	return _Thales.Contract.NominatedOwner(&_Thales.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Thales *ThalesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Thales *ThalesSession) Owner() (common.Address, error) {
	return _Thales.Contract.Owner(&_Thales.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Thales *ThalesCallerSession) Owner() (common.Address, error) {
	return _Thales.Contract.Owner(&_Thales.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Thales *ThalesCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Thales *ThalesSession) Paused() (bool, error) {
	return _Thales.Contract.Paused(&_Thales.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Thales *ThalesCallerSession) Paused() (bool, error) {
	return _Thales.Contract.Paused(&_Thales.CallOpts)
}

// PreviousManager is a free data retrieval call binding the contract method 0x06f58fe4.
//
// Solidity: function previousManager() view returns(address)
func (_Thales *ThalesCaller) PreviousManager(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "previousManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreviousManager is a free data retrieval call binding the contract method 0x06f58fe4.
//
// Solidity: function previousManager() view returns(address)
func (_Thales *ThalesSession) PreviousManager() (common.Address, error) {
	return _Thales.Contract.PreviousManager(&_Thales.CallOpts)
}

// PreviousManager is a free data retrieval call binding the contract method 0x06f58fe4.
//
// Solidity: function previousManager() view returns(address)
func (_Thales *ThalesCallerSession) PreviousManager() (common.Address, error) {
	return _Thales.Contract.PreviousManager(&_Thales.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0x6ed033f8.
//
// Solidity: function price(address market, uint8 position) view returns(uint256 price)
func (_Thales *ThalesCaller) Price(opts *bind.CallOpts, market common.Address, position uint8) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "price", market, position)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0x6ed033f8.
//
// Solidity: function price(address market, uint8 position) view returns(uint256 price)
func (_Thales *ThalesSession) Price(market common.Address, position uint8) (*big.Int, error) {
	return _Thales.Contract.Price(&_Thales.CallOpts, market, position)
}

// Price is a free data retrieval call binding the contract method 0x6ed033f8.
//
// Solidity: function price(address market, uint8 position) view returns(uint256 price)
func (_Thales *ThalesCallerSession) Price(market common.Address, position uint8) (*big.Int, error) {
	return _Thales.Contract.Price(&_Thales.CallOpts, market, position)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Thales *ThalesCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Thales *ThalesSession) PriceFeed() (common.Address, error) {
	return _Thales.Contract.PriceFeed(&_Thales.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Thales *ThalesCallerSession) PriceFeed() (common.Address, error) {
	return _Thales.Contract.PriceFeed(&_Thales.CallOpts)
}

// Referrals is a free data retrieval call binding the contract method 0xd3dc7539.
//
// Solidity: function referrals() view returns(address)
func (_Thales *ThalesCaller) Referrals(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "referrals")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Referrals is a free data retrieval call binding the contract method 0xd3dc7539.
//
// Solidity: function referrals() view returns(address)
func (_Thales *ThalesSession) Referrals() (common.Address, error) {
	return _Thales.Contract.Referrals(&_Thales.CallOpts)
}

// Referrals is a free data retrieval call binding the contract method 0xd3dc7539.
//
// Solidity: function referrals() view returns(address)
func (_Thales *ThalesCallerSession) Referrals() (common.Address, error) {
	return _Thales.Contract.Referrals(&_Thales.CallOpts)
}

// ReferrerFee is a free data retrieval call binding the contract method 0x6e88a7bd.
//
// Solidity: function referrerFee() view returns(uint256)
func (_Thales *ThalesCaller) ReferrerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "referrerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferrerFee is a free data retrieval call binding the contract method 0x6e88a7bd.
//
// Solidity: function referrerFee() view returns(uint256)
func (_Thales *ThalesSession) ReferrerFee() (*big.Int, error) {
	return _Thales.Contract.ReferrerFee(&_Thales.CallOpts)
}

// ReferrerFee is a free data retrieval call binding the contract method 0x6e88a7bd.
//
// Solidity: function referrerFee() view returns(uint256)
func (_Thales *ThalesCallerSession) ReferrerFee() (*big.Int, error) {
	return _Thales.Contract.ReferrerFee(&_Thales.CallOpts)
}

// SUSD is a free data retrieval call binding the contract method 0x9324cac7.
//
// Solidity: function sUSD() view returns(address)
func (_Thales *ThalesCaller) SUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "sUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SUSD is a free data retrieval call binding the contract method 0x9324cac7.
//
// Solidity: function sUSD() view returns(address)
func (_Thales *ThalesSession) SUSD() (common.Address, error) {
	return _Thales.Contract.SUSD(&_Thales.CallOpts)
}

// SUSD is a free data retrieval call binding the contract method 0x9324cac7.
//
// Solidity: function sUSD() view returns(address)
func (_Thales *ThalesCallerSession) SUSD() (common.Address, error) {
	return _Thales.Contract.SUSD(&_Thales.CallOpts)
}

// SafeBox is a free data retrieval call binding the contract method 0x48663e95.
//
// Solidity: function safeBox() view returns(address)
func (_Thales *ThalesCaller) SafeBox(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "safeBox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeBox is a free data retrieval call binding the contract method 0x48663e95.
//
// Solidity: function safeBox() view returns(address)
func (_Thales *ThalesSession) SafeBox() (common.Address, error) {
	return _Thales.Contract.SafeBox(&_Thales.CallOpts)
}

// SafeBox is a free data retrieval call binding the contract method 0x48663e95.
//
// Solidity: function safeBox() view returns(address)
func (_Thales *ThalesCallerSession) SafeBox() (common.Address, error) {
	return _Thales.Contract.SafeBox(&_Thales.CallOpts)
}

// SafeBoxImpact is a free data retrieval call binding the contract method 0xd69fb668.
//
// Solidity: function safeBoxImpact() view returns(uint256)
func (_Thales *ThalesCaller) SafeBoxImpact(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "safeBoxImpact")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SafeBoxImpact is a free data retrieval call binding the contract method 0xd69fb668.
//
// Solidity: function safeBoxImpact() view returns(uint256)
func (_Thales *ThalesSession) SafeBoxImpact() (*big.Int, error) {
	return _Thales.Contract.SafeBoxImpact(&_Thales.CallOpts)
}

// SafeBoxImpact is a free data retrieval call binding the contract method 0xd69fb668.
//
// Solidity: function safeBoxImpact() view returns(uint256)
func (_Thales *ThalesCallerSession) SafeBoxImpact() (*big.Int, error) {
	return _Thales.Contract.SafeBoxImpact(&_Thales.CallOpts)
}

// SellPriceImpact is a free data retrieval call binding the contract method 0x1c37d04b.
//
// Solidity: function sellPriceImpact(address market, uint8 position, uint256 amount) view returns(uint256 _impact)
func (_Thales *ThalesCaller) SellPriceImpact(opts *bind.CallOpts, market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "sellPriceImpact", market, position, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellPriceImpact is a free data retrieval call binding the contract method 0x1c37d04b.
//
// Solidity: function sellPriceImpact(address market, uint8 position, uint256 amount) view returns(uint256 _impact)
func (_Thales *ThalesSession) SellPriceImpact(market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	return _Thales.Contract.SellPriceImpact(&_Thales.CallOpts, market, position, amount)
}

// SellPriceImpact is a free data retrieval call binding the contract method 0x1c37d04b.
//
// Solidity: function sellPriceImpact(address market, uint8 position, uint256 amount) view returns(uint256 _impact)
func (_Thales *ThalesCallerSession) SellPriceImpact(market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	return _Thales.Contract.SellPriceImpact(&_Thales.CallOpts, market, position, amount)
}

// SellToAmmQuote is a free data retrieval call binding the contract method 0x0f13aae8.
//
// Solidity: function sellToAmmQuote(address market, uint8 position, uint256 amount) view returns(uint256 _available)
func (_Thales *ThalesCaller) SellToAmmQuote(opts *bind.CallOpts, market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "sellToAmmQuote", market, position, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellToAmmQuote is a free data retrieval call binding the contract method 0x0f13aae8.
//
// Solidity: function sellToAmmQuote(address market, uint8 position, uint256 amount) view returns(uint256 _available)
func (_Thales *ThalesSession) SellToAmmQuote(market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	return _Thales.Contract.SellToAmmQuote(&_Thales.CallOpts, market, position, amount)
}

// SellToAmmQuote is a free data retrieval call binding the contract method 0x0f13aae8.
//
// Solidity: function sellToAmmQuote(address market, uint8 position, uint256 amount) view returns(uint256 _available)
func (_Thales *ThalesCallerSession) SellToAmmQuote(market common.Address, position uint8, amount *big.Int) (*big.Int, error) {
	return _Thales.Contract.SellToAmmQuote(&_Thales.CallOpts, market, position, amount)
}

// SpentOnMarket is a free data retrieval call binding the contract method 0xc4dc27d7.
//
// Solidity: function spentOnMarket(address ) view returns(uint256)
func (_Thales *ThalesCaller) SpentOnMarket(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "spentOnMarket", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpentOnMarket is a free data retrieval call binding the contract method 0xc4dc27d7.
//
// Solidity: function spentOnMarket(address ) view returns(uint256)
func (_Thales *ThalesSession) SpentOnMarket(arg0 common.Address) (*big.Int, error) {
	return _Thales.Contract.SpentOnMarket(&_Thales.CallOpts, arg0)
}

// SpentOnMarket is a free data retrieval call binding the contract method 0xc4dc27d7.
//
// Solidity: function spentOnMarket(address ) view returns(uint256)
func (_Thales *ThalesCallerSession) SpentOnMarket(arg0 common.Address) (*big.Int, error) {
	return _Thales.Contract.SpentOnMarket(&_Thales.CallOpts, arg0)
}

// StakingThales is a free data retrieval call binding the contract method 0xfd8a8cc6.
//
// Solidity: function stakingThales() view returns(address)
func (_Thales *ThalesCaller) StakingThales(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "stakingThales")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingThales is a free data retrieval call binding the contract method 0xfd8a8cc6.
//
// Solidity: function stakingThales() view returns(address)
func (_Thales *ThalesSession) StakingThales() (common.Address, error) {
	return _Thales.Contract.StakingThales(&_Thales.CallOpts)
}

// StakingThales is a free data retrieval call binding the contract method 0xfd8a8cc6.
//
// Solidity: function stakingThales() view returns(address)
func (_Thales *ThalesCallerSession) StakingThales() (common.Address, error) {
	return _Thales.Contract.StakingThales(&_Thales.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Thales *ThalesCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Thales *ThalesSession) Usdc() (common.Address, error) {
	return _Thales.Contract.Usdc(&_Thales.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Thales *ThalesCallerSession) Usdc() (common.Address, error) {
	return _Thales.Contract.Usdc(&_Thales.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Thales *ThalesCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Thales *ThalesSession) Usdt() (common.Address, error) {
	return _Thales.Contract.Usdt(&_Thales.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Thales *ThalesCallerSession) Usdt() (common.Address, error) {
	return _Thales.Contract.Usdt(&_Thales.CallOpts)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0x06c933d8.
//
// Solidity: function whitelistedAddresses(address ) view returns(bool)
func (_Thales *ThalesCaller) WhitelistedAddresses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []any
	err := _Thales.contract.Call(opts, &out, "whitelistedAddresses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0x06c933d8.
//
// Solidity: function whitelistedAddresses(address ) view returns(bool)
func (_Thales *ThalesSession) WhitelistedAddresses(arg0 common.Address) (bool, error) {
	return _Thales.Contract.WhitelistedAddresses(&_Thales.CallOpts, arg0)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0x06c933d8.
//
// Solidity: function whitelistedAddresses(address ) view returns(bool)
func (_Thales *ThalesCallerSession) WhitelistedAddresses(arg0 common.Address) (bool, error) {
	return _Thales.Contract.WhitelistedAddresses(&_Thales.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Thales *ThalesTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Thales *ThalesSession) AcceptOwnership() (*types.Transaction, error) {
	return _Thales.Contract.AcceptOwnership(&_Thales.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Thales *ThalesTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Thales.Contract.AcceptOwnership(&_Thales.TransactOpts)
}

// BuyFromAMM is a paid mutator transaction binding the contract method 0x8875eb84.
//
// Solidity: function buyFromAMM(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage) returns()
func (_Thales *ThalesTransactor) BuyFromAMM(opts *bind.TransactOpts, market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "buyFromAMM", market, position, amount, expectedPayout, additionalSlippage)
}

// BuyFromAMM is a paid mutator transaction binding the contract method 0x8875eb84.
//
// Solidity: function buyFromAMM(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage) returns()
func (_Thales *ThalesSession) BuyFromAMM(market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.BuyFromAMM(&_Thales.TransactOpts, market, position, amount, expectedPayout, additionalSlippage)
}

// BuyFromAMM is a paid mutator transaction binding the contract method 0x8875eb84.
//
// Solidity: function buyFromAMM(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage) returns()
func (_Thales *ThalesTransactorSession) BuyFromAMM(market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.BuyFromAMM(&_Thales.TransactOpts, market, position, amount, expectedPayout, additionalSlippage)
}

// BuyFromAMMWithDifferentCollateralAndReferrer is a paid mutator transaction binding the contract method 0x6cc5a6ff.
//
// Solidity: function buyFromAMMWithDifferentCollateralAndReferrer(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage, address collateral, address _referrer) returns()
func (_Thales *ThalesTransactor) BuyFromAMMWithDifferentCollateralAndReferrer(opts *bind.TransactOpts, market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int, collateral common.Address, _referrer common.Address) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "buyFromAMMWithDifferentCollateralAndReferrer", market, position, amount, expectedPayout, additionalSlippage, collateral, _referrer)
}

// BuyFromAMMWithDifferentCollateralAndReferrer is a paid mutator transaction binding the contract method 0x6cc5a6ff.
//
// Solidity: function buyFromAMMWithDifferentCollateralAndReferrer(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage, address collateral, address _referrer) returns()
func (_Thales *ThalesSession) BuyFromAMMWithDifferentCollateralAndReferrer(market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int, collateral common.Address, _referrer common.Address) (*types.Transaction, error) {
	return _Thales.Contract.BuyFromAMMWithDifferentCollateralAndReferrer(&_Thales.TransactOpts, market, position, amount, expectedPayout, additionalSlippage, collateral, _referrer)
}

// BuyFromAMMWithDifferentCollateralAndReferrer is a paid mutator transaction binding the contract method 0x6cc5a6ff.
//
// Solidity: function buyFromAMMWithDifferentCollateralAndReferrer(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage, address collateral, address _referrer) returns()
func (_Thales *ThalesTransactorSession) BuyFromAMMWithDifferentCollateralAndReferrer(market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int, collateral common.Address, _referrer common.Address) (*types.Transaction, error) {
	return _Thales.Contract.BuyFromAMMWithDifferentCollateralAndReferrer(&_Thales.TransactOpts, market, position, amount, expectedPayout, additionalSlippage, collateral, _referrer)
}

// BuyFromAMMWithReferrer is a paid mutator transaction binding the contract method 0x9f916c9f.
//
// Solidity: function buyFromAMMWithReferrer(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage, address _referrer) returns()
func (_Thales *ThalesTransactor) BuyFromAMMWithReferrer(opts *bind.TransactOpts, market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "buyFromAMMWithReferrer", market, position, amount, expectedPayout, additionalSlippage, _referrer)
}

// BuyFromAMMWithReferrer is a paid mutator transaction binding the contract method 0x9f916c9f.
//
// Solidity: function buyFromAMMWithReferrer(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage, address _referrer) returns()
func (_Thales *ThalesSession) BuyFromAMMWithReferrer(market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _Thales.Contract.BuyFromAMMWithReferrer(&_Thales.TransactOpts, market, position, amount, expectedPayout, additionalSlippage, _referrer)
}

// BuyFromAMMWithReferrer is a paid mutator transaction binding the contract method 0x9f916c9f.
//
// Solidity: function buyFromAMMWithReferrer(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage, address _referrer) returns()
func (_Thales *ThalesTransactorSession) BuyFromAMMWithReferrer(market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _Thales.Contract.BuyFromAMMWithReferrer(&_Thales.TransactOpts, market, position, amount, expectedPayout, additionalSlippage, _referrer)
}

// ExerciseMaturedMarket is a paid mutator transaction binding the contract method 0x51b9181f.
//
// Solidity: function exerciseMaturedMarket(address market) returns()
func (_Thales *ThalesTransactor) ExerciseMaturedMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "exerciseMaturedMarket", market)
}

// ExerciseMaturedMarket is a paid mutator transaction binding the contract method 0x51b9181f.
//
// Solidity: function exerciseMaturedMarket(address market) returns()
func (_Thales *ThalesSession) ExerciseMaturedMarket(market common.Address) (*types.Transaction, error) {
	return _Thales.Contract.ExerciseMaturedMarket(&_Thales.TransactOpts, market)
}

// ExerciseMaturedMarket is a paid mutator transaction binding the contract method 0x51b9181f.
//
// Solidity: function exerciseMaturedMarket(address market) returns()
func (_Thales *ThalesTransactorSession) ExerciseMaturedMarket(market common.Address) (*types.Transaction, error) {
	return _Thales.Contract.ExerciseMaturedMarket(&_Thales.TransactOpts, market)
}

// InitNonReentrant is a paid mutator transaction binding the contract method 0xebc79772.
//
// Solidity: function initNonReentrant() returns()
func (_Thales *ThalesTransactor) InitNonReentrant(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "initNonReentrant")
}

// InitNonReentrant is a paid mutator transaction binding the contract method 0xebc79772.
//
// Solidity: function initNonReentrant() returns()
func (_Thales *ThalesSession) InitNonReentrant() (*types.Transaction, error) {
	return _Thales.Contract.InitNonReentrant(&_Thales.TransactOpts)
}

// InitNonReentrant is a paid mutator transaction binding the contract method 0xebc79772.
//
// Solidity: function initNonReentrant() returns()
func (_Thales *ThalesTransactorSession) InitNonReentrant() (*types.Transaction, error) {
	return _Thales.Contract.InitNonReentrant(&_Thales.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xad18f0da.
//
// Solidity: function initialize(address _owner, address _priceFeed, address _sUSD, uint256 _capPerMarket, address _deciMath, uint256 _min_spread, uint256 _max_spread, uint256 _minimalTimeLeftToMaturity) returns()
func (_Thales *ThalesTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _priceFeed common.Address, _sUSD common.Address, _capPerMarket *big.Int, _deciMath common.Address, _min_spread *big.Int, _max_spread *big.Int, _minimalTimeLeftToMaturity *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "initialize", _owner, _priceFeed, _sUSD, _capPerMarket, _deciMath, _min_spread, _max_spread, _minimalTimeLeftToMaturity)
}

// Initialize is a paid mutator transaction binding the contract method 0xad18f0da.
//
// Solidity: function initialize(address _owner, address _priceFeed, address _sUSD, uint256 _capPerMarket, address _deciMath, uint256 _min_spread, uint256 _max_spread, uint256 _minimalTimeLeftToMaturity) returns()
func (_Thales *ThalesSession) Initialize(_owner common.Address, _priceFeed common.Address, _sUSD common.Address, _capPerMarket *big.Int, _deciMath common.Address, _min_spread *big.Int, _max_spread *big.Int, _minimalTimeLeftToMaturity *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.Initialize(&_Thales.TransactOpts, _owner, _priceFeed, _sUSD, _capPerMarket, _deciMath, _min_spread, _max_spread, _minimalTimeLeftToMaturity)
}

// Initialize is a paid mutator transaction binding the contract method 0xad18f0da.
//
// Solidity: function initialize(address _owner, address _priceFeed, address _sUSD, uint256 _capPerMarket, address _deciMath, uint256 _min_spread, uint256 _max_spread, uint256 _minimalTimeLeftToMaturity) returns()
func (_Thales *ThalesTransactorSession) Initialize(_owner common.Address, _priceFeed common.Address, _sUSD common.Address, _capPerMarket *big.Int, _deciMath common.Address, _min_spread *big.Int, _max_spread *big.Int, _minimalTimeLeftToMaturity *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.Initialize(&_Thales.TransactOpts, _owner, _priceFeed, _sUSD, _capPerMarket, _deciMath, _min_spread, _max_spread, _minimalTimeLeftToMaturity)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Thales *ThalesTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Thales *ThalesSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Thales.Contract.NominateNewOwner(&_Thales.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Thales *ThalesTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Thales.Contract.NominateNewOwner(&_Thales.TransactOpts, _owner)
}

// RetrieveSUSDAmount is a paid mutator transaction binding the contract method 0xefb1fe35.
//
// Solidity: function retrieveSUSDAmount(address account, uint256 amount) returns()
func (_Thales *ThalesTransactor) RetrieveSUSDAmount(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "retrieveSUSDAmount", account, amount)
}

// RetrieveSUSDAmount is a paid mutator transaction binding the contract method 0xefb1fe35.
//
// Solidity: function retrieveSUSDAmount(address account, uint256 amount) returns()
func (_Thales *ThalesSession) RetrieveSUSDAmount(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.RetrieveSUSDAmount(&_Thales.TransactOpts, account, amount)
}

// RetrieveSUSDAmount is a paid mutator transaction binding the contract method 0xefb1fe35.
//
// Solidity: function retrieveSUSDAmount(address account, uint256 amount) returns()
func (_Thales *ThalesTransactorSession) RetrieveSUSDAmount(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.RetrieveSUSDAmount(&_Thales.TransactOpts, account, amount)
}

// SellToAMM is a paid mutator transaction binding the contract method 0x3ce1108d.
//
// Solidity: function sellToAMM(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage) returns()
func (_Thales *ThalesTransactor) SellToAMM(opts *bind.TransactOpts, market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "sellToAMM", market, position, amount, expectedPayout, additionalSlippage)
}

// SellToAMM is a paid mutator transaction binding the contract method 0x3ce1108d.
//
// Solidity: function sellToAMM(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage) returns()
func (_Thales *ThalesSession) SellToAMM(market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SellToAMM(&_Thales.TransactOpts, market, position, amount, expectedPayout, additionalSlippage)
}

// SellToAMM is a paid mutator transaction binding the contract method 0x3ce1108d.
//
// Solidity: function sellToAMM(address market, uint8 position, uint256 amount, uint256 expectedPayout, uint256 additionalSlippage) returns()
func (_Thales *ThalesTransactorSession) SellToAMM(market common.Address, position uint8, amount *big.Int, expectedPayout *big.Int, additionalSlippage *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SellToAMM(&_Thales.TransactOpts, market, position, amount, expectedPayout, additionalSlippage)
}

// SetCapPerAsset is a paid mutator transaction binding the contract method 0xd1624924.
//
// Solidity: function setCapPerAsset(bytes32 asset, uint256 _cap) returns()
func (_Thales *ThalesTransactor) SetCapPerAsset(opts *bind.TransactOpts, asset [32]byte, _cap *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setCapPerAsset", asset, _cap)
}

// SetCapPerAsset is a paid mutator transaction binding the contract method 0xd1624924.
//
// Solidity: function setCapPerAsset(bytes32 asset, uint256 _cap) returns()
func (_Thales *ThalesSession) SetCapPerAsset(asset [32]byte, _cap *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetCapPerAsset(&_Thales.TransactOpts, asset, _cap)
}

// SetCapPerAsset is a paid mutator transaction binding the contract method 0xd1624924.
//
// Solidity: function setCapPerAsset(bytes32 asset, uint256 _cap) returns()
func (_Thales *ThalesTransactorSession) SetCapPerAsset(asset [32]byte, _cap *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetCapPerAsset(&_Thales.TransactOpts, asset, _cap)
}

// SetCurveSUSD is a paid mutator transaction binding the contract method 0xb42b5798.
//
// Solidity: function setCurveSUSD(address _curveSUSD, address _dai, address _usdc, address _usdt, bool _curveOnrampEnabled) returns()
func (_Thales *ThalesTransactor) SetCurveSUSD(opts *bind.TransactOpts, _curveSUSD common.Address, _dai common.Address, _usdc common.Address, _usdt common.Address, _curveOnrampEnabled bool) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setCurveSUSD", _curveSUSD, _dai, _usdc, _usdt, _curveOnrampEnabled)
}

// SetCurveSUSD is a paid mutator transaction binding the contract method 0xb42b5798.
//
// Solidity: function setCurveSUSD(address _curveSUSD, address _dai, address _usdc, address _usdt, bool _curveOnrampEnabled) returns()
func (_Thales *ThalesSession) SetCurveSUSD(_curveSUSD common.Address, _dai common.Address, _usdc common.Address, _usdt common.Address, _curveOnrampEnabled bool) (*types.Transaction, error) {
	return _Thales.Contract.SetCurveSUSD(&_Thales.TransactOpts, _curveSUSD, _dai, _usdc, _usdt, _curveOnrampEnabled)
}

// SetCurveSUSD is a paid mutator transaction binding the contract method 0xb42b5798.
//
// Solidity: function setCurveSUSD(address _curveSUSD, address _dai, address _usdc, address _usdt, bool _curveOnrampEnabled) returns()
func (_Thales *ThalesTransactorSession) SetCurveSUSD(_curveSUSD common.Address, _dai common.Address, _usdc common.Address, _usdt common.Address, _curveOnrampEnabled bool) (*types.Transaction, error) {
	return _Thales.Contract.SetCurveSUSD(&_Thales.TransactOpts, _curveSUSD, _dai, _usdc, _usdt, _curveOnrampEnabled)
}

// SetImpliedVolatilityPerAsset is a paid mutator transaction binding the contract method 0x5ef85b26.
//
// Solidity: function setImpliedVolatilityPerAsset(bytes32 asset, uint256 _impliedVolatility) returns()
func (_Thales *ThalesTransactor) SetImpliedVolatilityPerAsset(opts *bind.TransactOpts, asset [32]byte, _impliedVolatility *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setImpliedVolatilityPerAsset", asset, _impliedVolatility)
}

// SetImpliedVolatilityPerAsset is a paid mutator transaction binding the contract method 0x5ef85b26.
//
// Solidity: function setImpliedVolatilityPerAsset(bytes32 asset, uint256 _impliedVolatility) returns()
func (_Thales *ThalesSession) SetImpliedVolatilityPerAsset(asset [32]byte, _impliedVolatility *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetImpliedVolatilityPerAsset(&_Thales.TransactOpts, asset, _impliedVolatility)
}

// SetImpliedVolatilityPerAsset is a paid mutator transaction binding the contract method 0x5ef85b26.
//
// Solidity: function setImpliedVolatilityPerAsset(bytes32 asset, uint256 _impliedVolatility) returns()
func (_Thales *ThalesTransactorSession) SetImpliedVolatilityPerAsset(asset [32]byte, _impliedVolatility *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetImpliedVolatilityPerAsset(&_Thales.TransactOpts, asset, _impliedVolatility)
}

// SetMinMaxSpread is a paid mutator transaction binding the contract method 0x82615349.
//
// Solidity: function setMinMaxSpread(uint256 _minspread, uint256 _maxspread) returns()
func (_Thales *ThalesTransactor) SetMinMaxSpread(opts *bind.TransactOpts, _minspread *big.Int, _maxspread *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setMinMaxSpread", _minspread, _maxspread)
}

// SetMinMaxSpread is a paid mutator transaction binding the contract method 0x82615349.
//
// Solidity: function setMinMaxSpread(uint256 _minspread, uint256 _maxspread) returns()
func (_Thales *ThalesSession) SetMinMaxSpread(_minspread *big.Int, _maxspread *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetMinMaxSpread(&_Thales.TransactOpts, _minspread, _maxspread)
}

// SetMinMaxSpread is a paid mutator transaction binding the contract method 0x82615349.
//
// Solidity: function setMinMaxSpread(uint256 _minspread, uint256 _maxspread) returns()
func (_Thales *ThalesTransactorSession) SetMinMaxSpread(_minspread *big.Int, _maxspread *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetMinMaxSpread(&_Thales.TransactOpts, _minspread, _maxspread)
}

// SetMinMaxSupportedPriceAndCap is a paid mutator transaction binding the contract method 0xc8f0b4ec.
//
// Solidity: function setMinMaxSupportedPriceAndCap(uint256 _minSupportedPrice, uint256 _maxSupportedPrice, uint256 _capPerMarket) returns()
func (_Thales *ThalesTransactor) SetMinMaxSupportedPriceAndCap(opts *bind.TransactOpts, _minSupportedPrice *big.Int, _maxSupportedPrice *big.Int, _capPerMarket *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setMinMaxSupportedPriceAndCap", _minSupportedPrice, _maxSupportedPrice, _capPerMarket)
}

// SetMinMaxSupportedPriceAndCap is a paid mutator transaction binding the contract method 0xc8f0b4ec.
//
// Solidity: function setMinMaxSupportedPriceAndCap(uint256 _minSupportedPrice, uint256 _maxSupportedPrice, uint256 _capPerMarket) returns()
func (_Thales *ThalesSession) SetMinMaxSupportedPriceAndCap(_minSupportedPrice *big.Int, _maxSupportedPrice *big.Int, _capPerMarket *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetMinMaxSupportedPriceAndCap(&_Thales.TransactOpts, _minSupportedPrice, _maxSupportedPrice, _capPerMarket)
}

// SetMinMaxSupportedPriceAndCap is a paid mutator transaction binding the contract method 0xc8f0b4ec.
//
// Solidity: function setMinMaxSupportedPriceAndCap(uint256 _minSupportedPrice, uint256 _maxSupportedPrice, uint256 _capPerMarket) returns()
func (_Thales *ThalesTransactorSession) SetMinMaxSupportedPriceAndCap(_minSupportedPrice *big.Int, _maxSupportedPrice *big.Int, _capPerMarket *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetMinMaxSupportedPriceAndCap(&_Thales.TransactOpts, _minSupportedPrice, _maxSupportedPrice, _capPerMarket)
}

// SetMinimalTimeLeftToMaturity is a paid mutator transaction binding the contract method 0x85170209.
//
// Solidity: function setMinimalTimeLeftToMaturity(uint256 _minimalTimeLeftToMaturity) returns()
func (_Thales *ThalesTransactor) SetMinimalTimeLeftToMaturity(opts *bind.TransactOpts, _minimalTimeLeftToMaturity *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setMinimalTimeLeftToMaturity", _minimalTimeLeftToMaturity)
}

// SetMinimalTimeLeftToMaturity is a paid mutator transaction binding the contract method 0x85170209.
//
// Solidity: function setMinimalTimeLeftToMaturity(uint256 _minimalTimeLeftToMaturity) returns()
func (_Thales *ThalesSession) SetMinimalTimeLeftToMaturity(_minimalTimeLeftToMaturity *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetMinimalTimeLeftToMaturity(&_Thales.TransactOpts, _minimalTimeLeftToMaturity)
}

// SetMinimalTimeLeftToMaturity is a paid mutator transaction binding the contract method 0x85170209.
//
// Solidity: function setMinimalTimeLeftToMaturity(uint256 _minimalTimeLeftToMaturity) returns()
func (_Thales *ThalesTransactorSession) SetMinimalTimeLeftToMaturity(_minimalTimeLeftToMaturity *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetMinimalTimeLeftToMaturity(&_Thales.TransactOpts, _minimalTimeLeftToMaturity)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Thales *ThalesTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Thales *ThalesSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Thales.Contract.SetOwner(&_Thales.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Thales *ThalesTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Thales.Contract.SetOwner(&_Thales.TransactOpts, _owner)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Thales *ThalesTransactor) SetPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setPaused", _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Thales *ThalesSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _Thales.Contract.SetPaused(&_Thales.TransactOpts, _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Thales *ThalesTransactorSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _Thales.Contract.SetPaused(&_Thales.TransactOpts, _paused)
}

// SetPositionalMarketManager is a paid mutator transaction binding the contract method 0xbf996ae3.
//
// Solidity: function setPositionalMarketManager(address _manager) returns()
func (_Thales *ThalesTransactor) SetPositionalMarketManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setPositionalMarketManager", _manager)
}

// SetPositionalMarketManager is a paid mutator transaction binding the contract method 0xbf996ae3.
//
// Solidity: function setPositionalMarketManager(address _manager) returns()
func (_Thales *ThalesSession) SetPositionalMarketManager(_manager common.Address) (*types.Transaction, error) {
	return _Thales.Contract.SetPositionalMarketManager(&_Thales.TransactOpts, _manager)
}

// SetPositionalMarketManager is a paid mutator transaction binding the contract method 0xbf996ae3.
//
// Solidity: function setPositionalMarketManager(address _manager) returns()
func (_Thales *ThalesTransactorSession) SetPositionalMarketManager(_manager common.Address) (*types.Transaction, error) {
	return _Thales.Contract.SetPositionalMarketManager(&_Thales.TransactOpts, _manager)
}

// SetPriceFeedAndSUSD is a paid mutator transaction binding the contract method 0xbbdf88ce.
//
// Solidity: function setPriceFeedAndSUSD(address _priceFeed, address _sUSD) returns()
func (_Thales *ThalesTransactor) SetPriceFeedAndSUSD(opts *bind.TransactOpts, _priceFeed common.Address, _sUSD common.Address) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setPriceFeedAndSUSD", _priceFeed, _sUSD)
}

// SetPriceFeedAndSUSD is a paid mutator transaction binding the contract method 0xbbdf88ce.
//
// Solidity: function setPriceFeedAndSUSD(address _priceFeed, address _sUSD) returns()
func (_Thales *ThalesSession) SetPriceFeedAndSUSD(_priceFeed common.Address, _sUSD common.Address) (*types.Transaction, error) {
	return _Thales.Contract.SetPriceFeedAndSUSD(&_Thales.TransactOpts, _priceFeed, _sUSD)
}

// SetPriceFeedAndSUSD is a paid mutator transaction binding the contract method 0xbbdf88ce.
//
// Solidity: function setPriceFeedAndSUSD(address _priceFeed, address _sUSD) returns()
func (_Thales *ThalesTransactorSession) SetPriceFeedAndSUSD(_priceFeed common.Address, _sUSD common.Address) (*types.Transaction, error) {
	return _Thales.Contract.SetPriceFeedAndSUSD(&_Thales.TransactOpts, _priceFeed, _sUSD)
}

// SetSafeBoxData is a paid mutator transaction binding the contract method 0xfc1ed857.
//
// Solidity: function setSafeBoxData(address _safeBox, uint256 _safeBoxImpact) returns()
func (_Thales *ThalesTransactor) SetSafeBoxData(opts *bind.TransactOpts, _safeBox common.Address, _safeBoxImpact *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setSafeBoxData", _safeBox, _safeBoxImpact)
}

// SetSafeBoxData is a paid mutator transaction binding the contract method 0xfc1ed857.
//
// Solidity: function setSafeBoxData(address _safeBox, uint256 _safeBoxImpact) returns()
func (_Thales *ThalesSession) SetSafeBoxData(_safeBox common.Address, _safeBoxImpact *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetSafeBoxData(&_Thales.TransactOpts, _safeBox, _safeBoxImpact)
}

// SetSafeBoxData is a paid mutator transaction binding the contract method 0xfc1ed857.
//
// Solidity: function setSafeBoxData(address _safeBox, uint256 _safeBoxImpact) returns()
func (_Thales *ThalesTransactorSession) SetSafeBoxData(_safeBox common.Address, _safeBoxImpact *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetSafeBoxData(&_Thales.TransactOpts, _safeBox, _safeBoxImpact)
}

// SetStakingThalesAndReferrals is a paid mutator transaction binding the contract method 0xcafb3ca2.
//
// Solidity: function setStakingThalesAndReferrals(address _stakingThales, address _referrals, uint256 _referrerFee) returns()
func (_Thales *ThalesTransactor) SetStakingThalesAndReferrals(opts *bind.TransactOpts, _stakingThales common.Address, _referrals common.Address, _referrerFee *big.Int) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setStakingThalesAndReferrals", _stakingThales, _referrals, _referrerFee)
}

// SetStakingThalesAndReferrals is a paid mutator transaction binding the contract method 0xcafb3ca2.
//
// Solidity: function setStakingThalesAndReferrals(address _stakingThales, address _referrals, uint256 _referrerFee) returns()
func (_Thales *ThalesSession) SetStakingThalesAndReferrals(_stakingThales common.Address, _referrals common.Address, _referrerFee *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetStakingThalesAndReferrals(&_Thales.TransactOpts, _stakingThales, _referrals, _referrerFee)
}

// SetStakingThalesAndReferrals is a paid mutator transaction binding the contract method 0xcafb3ca2.
//
// Solidity: function setStakingThalesAndReferrals(address _stakingThales, address _referrals, uint256 _referrerFee) returns()
func (_Thales *ThalesTransactorSession) SetStakingThalesAndReferrals(_stakingThales common.Address, _referrals common.Address, _referrerFee *big.Int) (*types.Transaction, error) {
	return _Thales.Contract.SetStakingThalesAndReferrals(&_Thales.TransactOpts, _stakingThales, _referrals, _referrerFee)
}

// SetWhitelistedAddress is a paid mutator transaction binding the contract method 0x7b337a36.
//
// Solidity: function setWhitelistedAddress(address _address, bool enabled) returns()
func (_Thales *ThalesTransactor) SetWhitelistedAddress(opts *bind.TransactOpts, _address common.Address, enabled bool) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "setWhitelistedAddress", _address, enabled)
}

// SetWhitelistedAddress is a paid mutator transaction binding the contract method 0x7b337a36.
//
// Solidity: function setWhitelistedAddress(address _address, bool enabled) returns()
func (_Thales *ThalesSession) SetWhitelistedAddress(_address common.Address, enabled bool) (*types.Transaction, error) {
	return _Thales.Contract.SetWhitelistedAddress(&_Thales.TransactOpts, _address, enabled)
}

// SetWhitelistedAddress is a paid mutator transaction binding the contract method 0x7b337a36.
//
// Solidity: function setWhitelistedAddress(address _address, bool enabled) returns()
func (_Thales *ThalesTransactorSession) SetWhitelistedAddress(_address common.Address, enabled bool) (*types.Transaction, error) {
	return _Thales.Contract.SetWhitelistedAddress(&_Thales.TransactOpts, _address, enabled)
}

// TransferOwnershipAtInit is a paid mutator transaction binding the contract method 0xc3b83f5f.
//
// Solidity: function transferOwnershipAtInit(address proxyAddress) returns()
func (_Thales *ThalesTransactor) TransferOwnershipAtInit(opts *bind.TransactOpts, proxyAddress common.Address) (*types.Transaction, error) {
	return _Thales.contract.Transact(opts, "transferOwnershipAtInit", proxyAddress)
}

// TransferOwnershipAtInit is a paid mutator transaction binding the contract method 0xc3b83f5f.
//
// Solidity: function transferOwnershipAtInit(address proxyAddress) returns()
func (_Thales *ThalesSession) TransferOwnershipAtInit(proxyAddress common.Address) (*types.Transaction, error) {
	return _Thales.Contract.TransferOwnershipAtInit(&_Thales.TransactOpts, proxyAddress)
}

// TransferOwnershipAtInit is a paid mutator transaction binding the contract method 0xc3b83f5f.
//
// Solidity: function transferOwnershipAtInit(address proxyAddress) returns()
func (_Thales *ThalesTransactorSession) TransferOwnershipAtInit(proxyAddress common.Address) (*types.Transaction, error) {
	return _Thales.Contract.TransferOwnershipAtInit(&_Thales.TransactOpts, proxyAddress)
}

// ThalesBoughtFromAmmIterator is returned from FilterBoughtFromAmm and is used to iterate over the raw logs and unpacked data for BoughtFromAmm events raised by the Thales contract.
type ThalesBoughtFromAmmIterator struct {
	Event *ThalesBoughtFromAmm // Event containing the contract specifics and raw log

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
func (it *ThalesBoughtFromAmmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesBoughtFromAmm)
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
		it.Event = new(ThalesBoughtFromAmm)
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
func (it *ThalesBoughtFromAmmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesBoughtFromAmmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesBoughtFromAmm represents a BoughtFromAmm event raised by the Thales contract.
type ThalesBoughtFromAmm struct {
	Buyer    common.Address
	Market   common.Address
	Position uint8
	Amount   *big.Int
	SUSDPaid *big.Int
	Susd     common.Address
	Asset    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBoughtFromAmm is a free log retrieval operation binding the contract event 0xf3bfbc0822d1ed667a2b298e71e0304f2c1f4685398189d7c39e412f733150f4.
//
// Solidity: event BoughtFromAmm(address buyer, address market, uint8 position, uint256 amount, uint256 sUSDPaid, address susd, address asset)
func (_Thales *ThalesFilterer) FilterBoughtFromAmm(opts *bind.FilterOpts) (*ThalesBoughtFromAmmIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "BoughtFromAmm")
	if err != nil {
		return nil, err
	}
	return &ThalesBoughtFromAmmIterator{contract: _Thales.contract, event: "BoughtFromAmm", logs: logs, sub: sub}, nil
}

// WatchBoughtFromAmm is a free log subscription operation binding the contract event 0xf3bfbc0822d1ed667a2b298e71e0304f2c1f4685398189d7c39e412f733150f4.
//
// Solidity: event BoughtFromAmm(address buyer, address market, uint8 position, uint256 amount, uint256 sUSDPaid, address susd, address asset)
func (_Thales *ThalesFilterer) WatchBoughtFromAmm(opts *bind.WatchOpts, sink chan<- *ThalesBoughtFromAmm) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "BoughtFromAmm")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesBoughtFromAmm)
				if err := _Thales.contract.UnpackLog(event, "BoughtFromAmm", log); err != nil {
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

// ParseBoughtFromAmm is a log parse operation binding the contract event 0xf3bfbc0822d1ed667a2b298e71e0304f2c1f4685398189d7c39e412f733150f4.
//
// Solidity: event BoughtFromAmm(address buyer, address market, uint8 position, uint256 amount, uint256 sUSDPaid, address susd, address asset)
func (_Thales *ThalesFilterer) ParseBoughtFromAmm(log types.Log) (*ThalesBoughtFromAmm, error) {
	event := new(ThalesBoughtFromAmm)
	if err := _Thales.contract.UnpackLog(event, "BoughtFromAmm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Thales contract.
type ThalesOwnerChangedIterator struct {
	Event *ThalesOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ThalesOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesOwnerChanged)
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
		it.Event = new(ThalesOwnerChanged)
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
func (it *ThalesOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesOwnerChanged represents a OwnerChanged event raised by the Thales contract.
type ThalesOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Thales *ThalesFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ThalesOwnerChangedIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ThalesOwnerChangedIterator{contract: _Thales.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Thales *ThalesFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ThalesOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesOwnerChanged)
				if err := _Thales.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_Thales *ThalesFilterer) ParseOwnerChanged(log types.Log) (*ThalesOwnerChanged, error) {
	event := new(ThalesOwnerChanged)
	if err := _Thales.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Thales contract.
type ThalesOwnerNominatedIterator struct {
	Event *ThalesOwnerNominated // Event containing the contract specifics and raw log

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
func (it *ThalesOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesOwnerNominated)
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
		it.Event = new(ThalesOwnerNominated)
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
func (it *ThalesOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesOwnerNominated represents a OwnerNominated event raised by the Thales contract.
type ThalesOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Thales *ThalesFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*ThalesOwnerNominatedIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &ThalesOwnerNominatedIterator{contract: _Thales.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Thales *ThalesFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *ThalesOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesOwnerNominated)
				if err := _Thales.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_Thales *ThalesFilterer) ParseOwnerNominated(log types.Log) (*ThalesOwnerNominated, error) {
	event := new(ThalesOwnerNominated)
	if err := _Thales.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesPauseChangedIterator is returned from FilterPauseChanged and is used to iterate over the raw logs and unpacked data for PauseChanged events raised by the Thales contract.
type ThalesPauseChangedIterator struct {
	Event *ThalesPauseChanged // Event containing the contract specifics and raw log

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
func (it *ThalesPauseChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesPauseChanged)
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
		it.Event = new(ThalesPauseChanged)
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
func (it *ThalesPauseChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesPauseChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesPauseChanged represents a PauseChanged event raised by the Thales contract.
type ThalesPauseChanged struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauseChanged is a free log retrieval operation binding the contract event 0x8fb6c181ee25a520cf3dd6565006ef91229fcfe5a989566c2a3b8c115570cec5.
//
// Solidity: event PauseChanged(bool isPaused)
func (_Thales *ThalesFilterer) FilterPauseChanged(opts *bind.FilterOpts) (*ThalesPauseChangedIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "PauseChanged")
	if err != nil {
		return nil, err
	}
	return &ThalesPauseChangedIterator{contract: _Thales.contract, event: "PauseChanged", logs: logs, sub: sub}, nil
}

// WatchPauseChanged is a free log subscription operation binding the contract event 0x8fb6c181ee25a520cf3dd6565006ef91229fcfe5a989566c2a3b8c115570cec5.
//
// Solidity: event PauseChanged(bool isPaused)
func (_Thales *ThalesFilterer) WatchPauseChanged(opts *bind.WatchOpts, sink chan<- *ThalesPauseChanged) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "PauseChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesPauseChanged)
				if err := _Thales.contract.UnpackLog(event, "PauseChanged", log); err != nil {
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

// ParsePauseChanged is a log parse operation binding the contract event 0x8fb6c181ee25a520cf3dd6565006ef91229fcfe5a989566c2a3b8c115570cec5.
//
// Solidity: event PauseChanged(bool isPaused)
func (_Thales *ThalesFilterer) ParsePauseChanged(log types.Log) (*ThalesPauseChanged, error) {
	event := new(ThalesPauseChanged)
	if err := _Thales.contract.UnpackLog(event, "PauseChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesReferrerPaidIterator is returned from FilterReferrerPaid and is used to iterate over the raw logs and unpacked data for ReferrerPaid events raised by the Thales contract.
type ThalesReferrerPaidIterator struct {
	Event *ThalesReferrerPaid // Event containing the contract specifics and raw log

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
func (it *ThalesReferrerPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesReferrerPaid)
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
		it.Event = new(ThalesReferrerPaid)
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
func (it *ThalesReferrerPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesReferrerPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesReferrerPaid represents a ReferrerPaid event raised by the Thales contract.
type ThalesReferrerPaid struct {
	Refferer common.Address
	Trader   common.Address
	Amount   *big.Int
	Volume   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReferrerPaid is a free log retrieval operation binding the contract event 0x8fa68a6a8e2fc9ff758a6e64afba8bc2f66fb082999a2c5225c8c49633faded4.
//
// Solidity: event ReferrerPaid(address refferer, address trader, uint256 amount, uint256 volume)
func (_Thales *ThalesFilterer) FilterReferrerPaid(opts *bind.FilterOpts) (*ThalesReferrerPaidIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "ReferrerPaid")
	if err != nil {
		return nil, err
	}
	return &ThalesReferrerPaidIterator{contract: _Thales.contract, event: "ReferrerPaid", logs: logs, sub: sub}, nil
}

// WatchReferrerPaid is a free log subscription operation binding the contract event 0x8fa68a6a8e2fc9ff758a6e64afba8bc2f66fb082999a2c5225c8c49633faded4.
//
// Solidity: event ReferrerPaid(address refferer, address trader, uint256 amount, uint256 volume)
func (_Thales *ThalesFilterer) WatchReferrerPaid(opts *bind.WatchOpts, sink chan<- *ThalesReferrerPaid) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "ReferrerPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesReferrerPaid)
				if err := _Thales.contract.UnpackLog(event, "ReferrerPaid", log); err != nil {
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

// ParseReferrerPaid is a log parse operation binding the contract event 0x8fa68a6a8e2fc9ff758a6e64afba8bc2f66fb082999a2c5225c8c49633faded4.
//
// Solidity: event ReferrerPaid(address refferer, address trader, uint256 amount, uint256 volume)
func (_Thales *ThalesFilterer) ParseReferrerPaid(log types.Log) (*ThalesReferrerPaid, error) {
	event := new(ThalesReferrerPaid)
	if err := _Thales.contract.UnpackLog(event, "ReferrerPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetCapPerAssetIterator is returned from FilterSetCapPerAsset and is used to iterate over the raw logs and unpacked data for SetCapPerAsset events raised by the Thales contract.
type ThalesSetCapPerAssetIterator struct {
	Event *ThalesSetCapPerAsset // Event containing the contract specifics and raw log

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
func (it *ThalesSetCapPerAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetCapPerAsset)
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
		it.Event = new(ThalesSetCapPerAsset)
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
func (it *ThalesSetCapPerAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetCapPerAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetCapPerAsset represents a SetCapPerAsset event raised by the Thales contract.
type ThalesSetCapPerAsset struct {
	Asset [32]byte
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetCapPerAsset is a free log retrieval operation binding the contract event 0x5af395595015797b4d0f26b77c38dd4831298dabc7906ee3d62b80fa75d35c1e.
//
// Solidity: event SetCapPerAsset(bytes32 asset, uint256 _cap)
func (_Thales *ThalesFilterer) FilterSetCapPerAsset(opts *bind.FilterOpts) (*ThalesSetCapPerAssetIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetCapPerAsset")
	if err != nil {
		return nil, err
	}
	return &ThalesSetCapPerAssetIterator{contract: _Thales.contract, event: "SetCapPerAsset", logs: logs, sub: sub}, nil
}

// WatchSetCapPerAsset is a free log subscription operation binding the contract event 0x5af395595015797b4d0f26b77c38dd4831298dabc7906ee3d62b80fa75d35c1e.
//
// Solidity: event SetCapPerAsset(bytes32 asset, uint256 _cap)
func (_Thales *ThalesFilterer) WatchSetCapPerAsset(opts *bind.WatchOpts, sink chan<- *ThalesSetCapPerAsset) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetCapPerAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetCapPerAsset)
				if err := _Thales.contract.UnpackLog(event, "SetCapPerAsset", log); err != nil {
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

// ParseSetCapPerAsset is a log parse operation binding the contract event 0x5af395595015797b4d0f26b77c38dd4831298dabc7906ee3d62b80fa75d35c1e.
//
// Solidity: event SetCapPerAsset(bytes32 asset, uint256 _cap)
func (_Thales *ThalesFilterer) ParseSetCapPerAsset(log types.Log) (*ThalesSetCapPerAsset, error) {
	event := new(ThalesSetCapPerAsset)
	if err := _Thales.contract.UnpackLog(event, "SetCapPerAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetCapPerMarketIterator is returned from FilterSetCapPerMarket and is used to iterate over the raw logs and unpacked data for SetCapPerMarket events raised by the Thales contract.
type ThalesSetCapPerMarketIterator struct {
	Event *ThalesSetCapPerMarket // Event containing the contract specifics and raw log

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
func (it *ThalesSetCapPerMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetCapPerMarket)
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
		it.Event = new(ThalesSetCapPerMarket)
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
func (it *ThalesSetCapPerMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetCapPerMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetCapPerMarket represents a SetCapPerMarket event raised by the Thales contract.
type ThalesSetCapPerMarket struct {
	CapPerMarket *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetCapPerMarket is a free log retrieval operation binding the contract event 0xcc72495e91bfc45ec11465a752fc2866b8a5ef8960b0925774ede927ee7cbdbe.
//
// Solidity: event SetCapPerMarket(uint256 _capPerMarket)
func (_Thales *ThalesFilterer) FilterSetCapPerMarket(opts *bind.FilterOpts) (*ThalesSetCapPerMarketIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetCapPerMarket")
	if err != nil {
		return nil, err
	}
	return &ThalesSetCapPerMarketIterator{contract: _Thales.contract, event: "SetCapPerMarket", logs: logs, sub: sub}, nil
}

// WatchSetCapPerMarket is a free log subscription operation binding the contract event 0xcc72495e91bfc45ec11465a752fc2866b8a5ef8960b0925774ede927ee7cbdbe.
//
// Solidity: event SetCapPerMarket(uint256 _capPerMarket)
func (_Thales *ThalesFilterer) WatchSetCapPerMarket(opts *bind.WatchOpts, sink chan<- *ThalesSetCapPerMarket) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetCapPerMarket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetCapPerMarket)
				if err := _Thales.contract.UnpackLog(event, "SetCapPerMarket", log); err != nil {
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

// ParseSetCapPerMarket is a log parse operation binding the contract event 0xcc72495e91bfc45ec11465a752fc2866b8a5ef8960b0925774ede927ee7cbdbe.
//
// Solidity: event SetCapPerMarket(uint256 _capPerMarket)
func (_Thales *ThalesFilterer) ParseSetCapPerMarket(log types.Log) (*ThalesSetCapPerMarket, error) {
	event := new(ThalesSetCapPerMarket)
	if err := _Thales.contract.UnpackLog(event, "SetCapPerMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetImpliedVolatilityPerAssetIterator is returned from FilterSetImpliedVolatilityPerAsset and is used to iterate over the raw logs and unpacked data for SetImpliedVolatilityPerAsset events raised by the Thales contract.
type ThalesSetImpliedVolatilityPerAssetIterator struct {
	Event *ThalesSetImpliedVolatilityPerAsset // Event containing the contract specifics and raw log

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
func (it *ThalesSetImpliedVolatilityPerAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetImpliedVolatilityPerAsset)
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
		it.Event = new(ThalesSetImpliedVolatilityPerAsset)
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
func (it *ThalesSetImpliedVolatilityPerAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetImpliedVolatilityPerAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetImpliedVolatilityPerAsset represents a SetImpliedVolatilityPerAsset event raised by the Thales contract.
type ThalesSetImpliedVolatilityPerAsset struct {
	Asset             [32]byte
	ImpliedVolatility *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetImpliedVolatilityPerAsset is a free log retrieval operation binding the contract event 0x715e0a52c0b74c77d2d2012a363ac95b494302ad2abb78ac7406ec93451f1adb.
//
// Solidity: event SetImpliedVolatilityPerAsset(bytes32 asset, uint256 _impliedVolatility)
func (_Thales *ThalesFilterer) FilterSetImpliedVolatilityPerAsset(opts *bind.FilterOpts) (*ThalesSetImpliedVolatilityPerAssetIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetImpliedVolatilityPerAsset")
	if err != nil {
		return nil, err
	}
	return &ThalesSetImpliedVolatilityPerAssetIterator{contract: _Thales.contract, event: "SetImpliedVolatilityPerAsset", logs: logs, sub: sub}, nil
}

// WatchSetImpliedVolatilityPerAsset is a free log subscription operation binding the contract event 0x715e0a52c0b74c77d2d2012a363ac95b494302ad2abb78ac7406ec93451f1adb.
//
// Solidity: event SetImpliedVolatilityPerAsset(bytes32 asset, uint256 _impliedVolatility)
func (_Thales *ThalesFilterer) WatchSetImpliedVolatilityPerAsset(opts *bind.WatchOpts, sink chan<- *ThalesSetImpliedVolatilityPerAsset) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetImpliedVolatilityPerAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetImpliedVolatilityPerAsset)
				if err := _Thales.contract.UnpackLog(event, "SetImpliedVolatilityPerAsset", log); err != nil {
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

// ParseSetImpliedVolatilityPerAsset is a log parse operation binding the contract event 0x715e0a52c0b74c77d2d2012a363ac95b494302ad2abb78ac7406ec93451f1adb.
//
// Solidity: event SetImpliedVolatilityPerAsset(bytes32 asset, uint256 _impliedVolatility)
func (_Thales *ThalesFilterer) ParseSetImpliedVolatilityPerAsset(log types.Log) (*ThalesSetImpliedVolatilityPerAsset, error) {
	event := new(ThalesSetImpliedVolatilityPerAsset)
	if err := _Thales.contract.UnpackLog(event, "SetImpliedVolatilityPerAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetMaxSpreadIterator is returned from FilterSetMaxSpread and is used to iterate over the raw logs and unpacked data for SetMaxSpread events raised by the Thales contract.
type ThalesSetMaxSpreadIterator struct {
	Event *ThalesSetMaxSpread // Event containing the contract specifics and raw log

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
func (it *ThalesSetMaxSpreadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetMaxSpread)
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
		it.Event = new(ThalesSetMaxSpread)
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
func (it *ThalesSetMaxSpreadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetMaxSpreadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetMaxSpread represents a SetMaxSpread event raised by the Thales contract.
type ThalesSetMaxSpread struct {
	Spread *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxSpread is a free log retrieval operation binding the contract event 0x4af6d03c4624e0a6b868b8f6453e047f23f3ea15e9d08c938bd4c445d7ef19b3.
//
// Solidity: event SetMaxSpread(uint256 _spread)
func (_Thales *ThalesFilterer) FilterSetMaxSpread(opts *bind.FilterOpts) (*ThalesSetMaxSpreadIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetMaxSpread")
	if err != nil {
		return nil, err
	}
	return &ThalesSetMaxSpreadIterator{contract: _Thales.contract, event: "SetMaxSpread", logs: logs, sub: sub}, nil
}

// WatchSetMaxSpread is a free log subscription operation binding the contract event 0x4af6d03c4624e0a6b868b8f6453e047f23f3ea15e9d08c938bd4c445d7ef19b3.
//
// Solidity: event SetMaxSpread(uint256 _spread)
func (_Thales *ThalesFilterer) WatchSetMaxSpread(opts *bind.WatchOpts, sink chan<- *ThalesSetMaxSpread) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetMaxSpread")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetMaxSpread)
				if err := _Thales.contract.UnpackLog(event, "SetMaxSpread", log); err != nil {
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

// ParseSetMaxSpread is a log parse operation binding the contract event 0x4af6d03c4624e0a6b868b8f6453e047f23f3ea15e9d08c938bd4c445d7ef19b3.
//
// Solidity: event SetMaxSpread(uint256 _spread)
func (_Thales *ThalesFilterer) ParseSetMaxSpread(log types.Log) (*ThalesSetMaxSpread, error) {
	event := new(ThalesSetMaxSpread)
	if err := _Thales.contract.UnpackLog(event, "SetMaxSpread", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetMaxSupportedPriceIterator is returned from FilterSetMaxSupportedPrice and is used to iterate over the raw logs and unpacked data for SetMaxSupportedPrice events raised by the Thales contract.
type ThalesSetMaxSupportedPriceIterator struct {
	Event *ThalesSetMaxSupportedPrice // Event containing the contract specifics and raw log

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
func (it *ThalesSetMaxSupportedPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetMaxSupportedPrice)
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
		it.Event = new(ThalesSetMaxSupportedPrice)
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
func (it *ThalesSetMaxSupportedPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetMaxSupportedPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetMaxSupportedPrice represents a SetMaxSupportedPrice event raised by the Thales contract.
type ThalesSetMaxSupportedPrice struct {
	Spread *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxSupportedPrice is a free log retrieval operation binding the contract event 0xafb2511d843afc28083da6c608f400d0262aef83b3cd885abd205b3e64f544bb.
//
// Solidity: event SetMaxSupportedPrice(uint256 _spread)
func (_Thales *ThalesFilterer) FilterSetMaxSupportedPrice(opts *bind.FilterOpts) (*ThalesSetMaxSupportedPriceIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetMaxSupportedPrice")
	if err != nil {
		return nil, err
	}
	return &ThalesSetMaxSupportedPriceIterator{contract: _Thales.contract, event: "SetMaxSupportedPrice", logs: logs, sub: sub}, nil
}

// WatchSetMaxSupportedPrice is a free log subscription operation binding the contract event 0xafb2511d843afc28083da6c608f400d0262aef83b3cd885abd205b3e64f544bb.
//
// Solidity: event SetMaxSupportedPrice(uint256 _spread)
func (_Thales *ThalesFilterer) WatchSetMaxSupportedPrice(opts *bind.WatchOpts, sink chan<- *ThalesSetMaxSupportedPrice) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetMaxSupportedPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetMaxSupportedPrice)
				if err := _Thales.contract.UnpackLog(event, "SetMaxSupportedPrice", log); err != nil {
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

// ParseSetMaxSupportedPrice is a log parse operation binding the contract event 0xafb2511d843afc28083da6c608f400d0262aef83b3cd885abd205b3e64f544bb.
//
// Solidity: event SetMaxSupportedPrice(uint256 _spread)
func (_Thales *ThalesFilterer) ParseSetMaxSupportedPrice(log types.Log) (*ThalesSetMaxSupportedPrice, error) {
	event := new(ThalesSetMaxSupportedPrice)
	if err := _Thales.contract.UnpackLog(event, "SetMaxSupportedPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetMinSpreadIterator is returned from FilterSetMinSpread and is used to iterate over the raw logs and unpacked data for SetMinSpread events raised by the Thales contract.
type ThalesSetMinSpreadIterator struct {
	Event *ThalesSetMinSpread // Event containing the contract specifics and raw log

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
func (it *ThalesSetMinSpreadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetMinSpread)
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
		it.Event = new(ThalesSetMinSpread)
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
func (it *ThalesSetMinSpreadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetMinSpreadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetMinSpread represents a SetMinSpread event raised by the Thales contract.
type ThalesSetMinSpread struct {
	Spread *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinSpread is a free log retrieval operation binding the contract event 0x1e6a338a58debcc786781a079c4459466b102ad0156cc84f51d25ef7dd8cb9b0.
//
// Solidity: event SetMinSpread(uint256 _spread)
func (_Thales *ThalesFilterer) FilterSetMinSpread(opts *bind.FilterOpts) (*ThalesSetMinSpreadIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetMinSpread")
	if err != nil {
		return nil, err
	}
	return &ThalesSetMinSpreadIterator{contract: _Thales.contract, event: "SetMinSpread", logs: logs, sub: sub}, nil
}

// WatchSetMinSpread is a free log subscription operation binding the contract event 0x1e6a338a58debcc786781a079c4459466b102ad0156cc84f51d25ef7dd8cb9b0.
//
// Solidity: event SetMinSpread(uint256 _spread)
func (_Thales *ThalesFilterer) WatchSetMinSpread(opts *bind.WatchOpts, sink chan<- *ThalesSetMinSpread) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetMinSpread")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetMinSpread)
				if err := _Thales.contract.UnpackLog(event, "SetMinSpread", log); err != nil {
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

// ParseSetMinSpread is a log parse operation binding the contract event 0x1e6a338a58debcc786781a079c4459466b102ad0156cc84f51d25ef7dd8cb9b0.
//
// Solidity: event SetMinSpread(uint256 _spread)
func (_Thales *ThalesFilterer) ParseSetMinSpread(log types.Log) (*ThalesSetMinSpread, error) {
	event := new(ThalesSetMinSpread)
	if err := _Thales.contract.UnpackLog(event, "SetMinSpread", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetMinSupportedPriceIterator is returned from FilterSetMinSupportedPrice and is used to iterate over the raw logs and unpacked data for SetMinSupportedPrice events raised by the Thales contract.
type ThalesSetMinSupportedPriceIterator struct {
	Event *ThalesSetMinSupportedPrice // Event containing the contract specifics and raw log

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
func (it *ThalesSetMinSupportedPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetMinSupportedPrice)
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
		it.Event = new(ThalesSetMinSupportedPrice)
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
func (it *ThalesSetMinSupportedPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetMinSupportedPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetMinSupportedPrice represents a SetMinSupportedPrice event raised by the Thales contract.
type ThalesSetMinSupportedPrice struct {
	Spread *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinSupportedPrice is a free log retrieval operation binding the contract event 0xa36b45ea010fcb2e70d33048b7d3f27c265915709bf40fc7e239459424933d9f.
//
// Solidity: event SetMinSupportedPrice(uint256 _spread)
func (_Thales *ThalesFilterer) FilterSetMinSupportedPrice(opts *bind.FilterOpts) (*ThalesSetMinSupportedPriceIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetMinSupportedPrice")
	if err != nil {
		return nil, err
	}
	return &ThalesSetMinSupportedPriceIterator{contract: _Thales.contract, event: "SetMinSupportedPrice", logs: logs, sub: sub}, nil
}

// WatchSetMinSupportedPrice is a free log subscription operation binding the contract event 0xa36b45ea010fcb2e70d33048b7d3f27c265915709bf40fc7e239459424933d9f.
//
// Solidity: event SetMinSupportedPrice(uint256 _spread)
func (_Thales *ThalesFilterer) WatchSetMinSupportedPrice(opts *bind.WatchOpts, sink chan<- *ThalesSetMinSupportedPrice) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetMinSupportedPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetMinSupportedPrice)
				if err := _Thales.contract.UnpackLog(event, "SetMinSupportedPrice", log); err != nil {
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

// ParseSetMinSupportedPrice is a log parse operation binding the contract event 0xa36b45ea010fcb2e70d33048b7d3f27c265915709bf40fc7e239459424933d9f.
//
// Solidity: event SetMinSupportedPrice(uint256 _spread)
func (_Thales *ThalesFilterer) ParseSetMinSupportedPrice(log types.Log) (*ThalesSetMinSupportedPrice, error) {
	event := new(ThalesSetMinSupportedPrice)
	if err := _Thales.contract.UnpackLog(event, "SetMinSupportedPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetMinimalTimeLeftToMaturityIterator is returned from FilterSetMinimalTimeLeftToMaturity and is used to iterate over the raw logs and unpacked data for SetMinimalTimeLeftToMaturity events raised by the Thales contract.
type ThalesSetMinimalTimeLeftToMaturityIterator struct {
	Event *ThalesSetMinimalTimeLeftToMaturity // Event containing the contract specifics and raw log

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
func (it *ThalesSetMinimalTimeLeftToMaturityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetMinimalTimeLeftToMaturity)
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
		it.Event = new(ThalesSetMinimalTimeLeftToMaturity)
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
func (it *ThalesSetMinimalTimeLeftToMaturityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetMinimalTimeLeftToMaturityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetMinimalTimeLeftToMaturity represents a SetMinimalTimeLeftToMaturity event raised by the Thales contract.
type ThalesSetMinimalTimeLeftToMaturity struct {
	MinimalTimeLeftToMaturity *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetMinimalTimeLeftToMaturity is a free log retrieval operation binding the contract event 0xdc469b5583fa9b7ebd3245e1665334cd758c4bef4c5a132c62baca85effacfec.
//
// Solidity: event SetMinimalTimeLeftToMaturity(uint256 _minimalTimeLeftToMaturity)
func (_Thales *ThalesFilterer) FilterSetMinimalTimeLeftToMaturity(opts *bind.FilterOpts) (*ThalesSetMinimalTimeLeftToMaturityIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetMinimalTimeLeftToMaturity")
	if err != nil {
		return nil, err
	}
	return &ThalesSetMinimalTimeLeftToMaturityIterator{contract: _Thales.contract, event: "SetMinimalTimeLeftToMaturity", logs: logs, sub: sub}, nil
}

// WatchSetMinimalTimeLeftToMaturity is a free log subscription operation binding the contract event 0xdc469b5583fa9b7ebd3245e1665334cd758c4bef4c5a132c62baca85effacfec.
//
// Solidity: event SetMinimalTimeLeftToMaturity(uint256 _minimalTimeLeftToMaturity)
func (_Thales *ThalesFilterer) WatchSetMinimalTimeLeftToMaturity(opts *bind.WatchOpts, sink chan<- *ThalesSetMinimalTimeLeftToMaturity) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetMinimalTimeLeftToMaturity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetMinimalTimeLeftToMaturity)
				if err := _Thales.contract.UnpackLog(event, "SetMinimalTimeLeftToMaturity", log); err != nil {
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

// ParseSetMinimalTimeLeftToMaturity is a log parse operation binding the contract event 0xdc469b5583fa9b7ebd3245e1665334cd758c4bef4c5a132c62baca85effacfec.
//
// Solidity: event SetMinimalTimeLeftToMaturity(uint256 _minimalTimeLeftToMaturity)
func (_Thales *ThalesFilterer) ParseSetMinimalTimeLeftToMaturity(log types.Log) (*ThalesSetMinimalTimeLeftToMaturity, error) {
	event := new(ThalesSetMinimalTimeLeftToMaturity)
	if err := _Thales.contract.UnpackLog(event, "SetMinimalTimeLeftToMaturity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetPositionalMarketManagerIterator is returned from FilterSetPositionalMarketManager and is used to iterate over the raw logs and unpacked data for SetPositionalMarketManager events raised by the Thales contract.
type ThalesSetPositionalMarketManagerIterator struct {
	Event *ThalesSetPositionalMarketManager // Event containing the contract specifics and raw log

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
func (it *ThalesSetPositionalMarketManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetPositionalMarketManager)
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
		it.Event = new(ThalesSetPositionalMarketManager)
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
func (it *ThalesSetPositionalMarketManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetPositionalMarketManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetPositionalMarketManager represents a SetPositionalMarketManager event raised by the Thales contract.
type ThalesSetPositionalMarketManager struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetPositionalMarketManager is a free log retrieval operation binding the contract event 0x9987372437ace1af79923f26b948aa04afef92b2b7786144c5aae621ea84eb0a.
//
// Solidity: event SetPositionalMarketManager(address _manager)
func (_Thales *ThalesFilterer) FilterSetPositionalMarketManager(opts *bind.FilterOpts) (*ThalesSetPositionalMarketManagerIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetPositionalMarketManager")
	if err != nil {
		return nil, err
	}
	return &ThalesSetPositionalMarketManagerIterator{contract: _Thales.contract, event: "SetPositionalMarketManager", logs: logs, sub: sub}, nil
}

// WatchSetPositionalMarketManager is a free log subscription operation binding the contract event 0x9987372437ace1af79923f26b948aa04afef92b2b7786144c5aae621ea84eb0a.
//
// Solidity: event SetPositionalMarketManager(address _manager)
func (_Thales *ThalesFilterer) WatchSetPositionalMarketManager(opts *bind.WatchOpts, sink chan<- *ThalesSetPositionalMarketManager) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetPositionalMarketManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetPositionalMarketManager)
				if err := _Thales.contract.UnpackLog(event, "SetPositionalMarketManager", log); err != nil {
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

// ParseSetPositionalMarketManager is a log parse operation binding the contract event 0x9987372437ace1af79923f26b948aa04afef92b2b7786144c5aae621ea84eb0a.
//
// Solidity: event SetPositionalMarketManager(address _manager)
func (_Thales *ThalesFilterer) ParseSetPositionalMarketManager(log types.Log) (*ThalesSetPositionalMarketManager, error) {
	event := new(ThalesSetPositionalMarketManager)
	if err := _Thales.contract.UnpackLog(event, "SetPositionalMarketManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetPriceFeedIterator is returned from FilterSetPriceFeed and is used to iterate over the raw logs and unpacked data for SetPriceFeed events raised by the Thales contract.
type ThalesSetPriceFeedIterator struct {
	Event *ThalesSetPriceFeed // Event containing the contract specifics and raw log

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
func (it *ThalesSetPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetPriceFeed)
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
		it.Event = new(ThalesSetPriceFeed)
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
func (it *ThalesSetPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetPriceFeed represents a SetPriceFeed event raised by the Thales contract.
type ThalesSetPriceFeed struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetPriceFeed is a free log retrieval operation binding the contract event 0xf724a45d041687842411f2b977ef22ab8f43c8f1104f4592b42a00f9b34a643d.
//
// Solidity: event SetPriceFeed(address _priceFeed)
func (_Thales *ThalesFilterer) FilterSetPriceFeed(opts *bind.FilterOpts) (*ThalesSetPriceFeedIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetPriceFeed")
	if err != nil {
		return nil, err
	}
	return &ThalesSetPriceFeedIterator{contract: _Thales.contract, event: "SetPriceFeed", logs: logs, sub: sub}, nil
}

// WatchSetPriceFeed is a free log subscription operation binding the contract event 0xf724a45d041687842411f2b977ef22ab8f43c8f1104f4592b42a00f9b34a643d.
//
// Solidity: event SetPriceFeed(address _priceFeed)
func (_Thales *ThalesFilterer) WatchSetPriceFeed(opts *bind.WatchOpts, sink chan<- *ThalesSetPriceFeed) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetPriceFeed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetPriceFeed)
				if err := _Thales.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
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

// ParseSetPriceFeed is a log parse operation binding the contract event 0xf724a45d041687842411f2b977ef22ab8f43c8f1104f4592b42a00f9b34a643d.
//
// Solidity: event SetPriceFeed(address _priceFeed)
func (_Thales *ThalesFilterer) ParseSetPriceFeed(log types.Log) (*ThalesSetPriceFeed, error) {
	event := new(ThalesSetPriceFeed)
	if err := _Thales.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetSUSDIterator is returned from FilterSetSUSD and is used to iterate over the raw logs and unpacked data for SetSUSD events raised by the Thales contract.
type ThalesSetSUSDIterator struct {
	Event *ThalesSetSUSD // Event containing the contract specifics and raw log

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
func (it *ThalesSetSUSDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetSUSD)
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
		it.Event = new(ThalesSetSUSD)
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
func (it *ThalesSetSUSDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetSUSDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetSUSD represents a SetSUSD event raised by the Thales contract.
type ThalesSetSUSD struct {
	SUSD common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetSUSD is a free log retrieval operation binding the contract event 0x74a8764fc8d62d2d844c8c54426bd94ad034e0e92abdf5280ff75e2cbd678fb6.
//
// Solidity: event SetSUSD(address sUSD)
func (_Thales *ThalesFilterer) FilterSetSUSD(opts *bind.FilterOpts) (*ThalesSetSUSDIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetSUSD")
	if err != nil {
		return nil, err
	}
	return &ThalesSetSUSDIterator{contract: _Thales.contract, event: "SetSUSD", logs: logs, sub: sub}, nil
}

// WatchSetSUSD is a free log subscription operation binding the contract event 0x74a8764fc8d62d2d844c8c54426bd94ad034e0e92abdf5280ff75e2cbd678fb6.
//
// Solidity: event SetSUSD(address sUSD)
func (_Thales *ThalesFilterer) WatchSetSUSD(opts *bind.WatchOpts, sink chan<- *ThalesSetSUSD) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetSUSD")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetSUSD)
				if err := _Thales.contract.UnpackLog(event, "SetSUSD", log); err != nil {
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

// ParseSetSUSD is a log parse operation binding the contract event 0x74a8764fc8d62d2d844c8c54426bd94ad034e0e92abdf5280ff75e2cbd678fb6.
//
// Solidity: event SetSUSD(address sUSD)
func (_Thales *ThalesFilterer) ParseSetSUSD(log types.Log) (*ThalesSetSUSD, error) {
	event := new(ThalesSetSUSD)
	if err := _Thales.contract.UnpackLog(event, "SetSUSD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetSafeBoxIterator is returned from FilterSetSafeBox and is used to iterate over the raw logs and unpacked data for SetSafeBox events raised by the Thales contract.
type ThalesSetSafeBoxIterator struct {
	Event *ThalesSetSafeBox // Event containing the contract specifics and raw log

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
func (it *ThalesSetSafeBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetSafeBox)
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
		it.Event = new(ThalesSetSafeBox)
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
func (it *ThalesSetSafeBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetSafeBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetSafeBox represents a SetSafeBox event raised by the Thales contract.
type ThalesSetSafeBox struct {
	SafeBox common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetSafeBox is a free log retrieval operation binding the contract event 0x49a4e3b1ee2ad2bfb1455dfad2d72e84e71a35c756d2c8296f305f12e1828bf5.
//
// Solidity: event SetSafeBox(address _safeBox)
func (_Thales *ThalesFilterer) FilterSetSafeBox(opts *bind.FilterOpts) (*ThalesSetSafeBoxIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetSafeBox")
	if err != nil {
		return nil, err
	}
	return &ThalesSetSafeBoxIterator{contract: _Thales.contract, event: "SetSafeBox", logs: logs, sub: sub}, nil
}

// WatchSetSafeBox is a free log subscription operation binding the contract event 0x49a4e3b1ee2ad2bfb1455dfad2d72e84e71a35c756d2c8296f305f12e1828bf5.
//
// Solidity: event SetSafeBox(address _safeBox)
func (_Thales *ThalesFilterer) WatchSetSafeBox(opts *bind.WatchOpts, sink chan<- *ThalesSetSafeBox) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetSafeBox")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetSafeBox)
				if err := _Thales.contract.UnpackLog(event, "SetSafeBox", log); err != nil {
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

// ParseSetSafeBox is a log parse operation binding the contract event 0x49a4e3b1ee2ad2bfb1455dfad2d72e84e71a35c756d2c8296f305f12e1828bf5.
//
// Solidity: event SetSafeBox(address _safeBox)
func (_Thales *ThalesFilterer) ParseSetSafeBox(log types.Log) (*ThalesSetSafeBox, error) {
	event := new(ThalesSetSafeBox)
	if err := _Thales.contract.UnpackLog(event, "SetSafeBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSetSafeBoxImpactIterator is returned from FilterSetSafeBoxImpact and is used to iterate over the raw logs and unpacked data for SetSafeBoxImpact events raised by the Thales contract.
type ThalesSetSafeBoxImpactIterator struct {
	Event *ThalesSetSafeBoxImpact // Event containing the contract specifics and raw log

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
func (it *ThalesSetSafeBoxImpactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSetSafeBoxImpact)
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
		it.Event = new(ThalesSetSafeBoxImpact)
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
func (it *ThalesSetSafeBoxImpactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSetSafeBoxImpactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSetSafeBoxImpact represents a SetSafeBoxImpact event raised by the Thales contract.
type ThalesSetSafeBoxImpact struct {
	SafeBoxImpact *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetSafeBoxImpact is a free log retrieval operation binding the contract event 0x01edd423db862fb00774918e3b06d9c1dd3db9a99b5a194c439d2f141876f444.
//
// Solidity: event SetSafeBoxImpact(uint256 _safeBoxImpact)
func (_Thales *ThalesFilterer) FilterSetSafeBoxImpact(opts *bind.FilterOpts) (*ThalesSetSafeBoxImpactIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SetSafeBoxImpact")
	if err != nil {
		return nil, err
	}
	return &ThalesSetSafeBoxImpactIterator{contract: _Thales.contract, event: "SetSafeBoxImpact", logs: logs, sub: sub}, nil
}

// WatchSetSafeBoxImpact is a free log subscription operation binding the contract event 0x01edd423db862fb00774918e3b06d9c1dd3db9a99b5a194c439d2f141876f444.
//
// Solidity: event SetSafeBoxImpact(uint256 _safeBoxImpact)
func (_Thales *ThalesFilterer) WatchSetSafeBoxImpact(opts *bind.WatchOpts, sink chan<- *ThalesSetSafeBoxImpact) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SetSafeBoxImpact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSetSafeBoxImpact)
				if err := _Thales.contract.UnpackLog(event, "SetSafeBoxImpact", log); err != nil {
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

// ParseSetSafeBoxImpact is a log parse operation binding the contract event 0x01edd423db862fb00774918e3b06d9c1dd3db9a99b5a194c439d2f141876f444.
//
// Solidity: event SetSafeBoxImpact(uint256 _safeBoxImpact)
func (_Thales *ThalesFilterer) ParseSetSafeBoxImpact(log types.Log) (*ThalesSetSafeBoxImpact, error) {
	event := new(ThalesSetSafeBoxImpact)
	if err := _Thales.contract.UnpackLog(event, "SetSafeBoxImpact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThalesSoldToAMMIterator is returned from FilterSoldToAMM and is used to iterate over the raw logs and unpacked data for SoldToAMM events raised by the Thales contract.
type ThalesSoldToAMMIterator struct {
	Event *ThalesSoldToAMM // Event containing the contract specifics and raw log

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
func (it *ThalesSoldToAMMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThalesSoldToAMM)
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
		it.Event = new(ThalesSoldToAMM)
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
func (it *ThalesSoldToAMMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThalesSoldToAMMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThalesSoldToAMM represents a SoldToAMM event raised by the Thales contract.
type ThalesSoldToAMM struct {
	Seller   common.Address
	Market   common.Address
	Position uint8
	Amount   *big.Int
	SUSDPaid *big.Int
	Susd     common.Address
	Asset    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSoldToAMM is a free log retrieval operation binding the contract event 0x1d6ff70c632edb1e6aba7fbc0148db68c8392e30f9dfaadae2543a2543757cf6.
//
// Solidity: event SoldToAMM(address seller, address market, uint8 position, uint256 amount, uint256 sUSDPaid, address susd, address asset)
func (_Thales *ThalesFilterer) FilterSoldToAMM(opts *bind.FilterOpts) (*ThalesSoldToAMMIterator, error) {

	logs, sub, err := _Thales.contract.FilterLogs(opts, "SoldToAMM")
	if err != nil {
		return nil, err
	}
	return &ThalesSoldToAMMIterator{contract: _Thales.contract, event: "SoldToAMM", logs: logs, sub: sub}, nil
}

// WatchSoldToAMM is a free log subscription operation binding the contract event 0x1d6ff70c632edb1e6aba7fbc0148db68c8392e30f9dfaadae2543a2543757cf6.
//
// Solidity: event SoldToAMM(address seller, address market, uint8 position, uint256 amount, uint256 sUSDPaid, address susd, address asset)
func (_Thales *ThalesFilterer) WatchSoldToAMM(opts *bind.WatchOpts, sink chan<- *ThalesSoldToAMM) (event.Subscription, error) {

	logs, sub, err := _Thales.contract.WatchLogs(opts, "SoldToAMM")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThalesSoldToAMM)
				if err := _Thales.contract.UnpackLog(event, "SoldToAMM", log); err != nil {
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

// ParseSoldToAMM is a log parse operation binding the contract event 0x1d6ff70c632edb1e6aba7fbc0148db68c8392e30f9dfaadae2543a2543757cf6.
//
// Solidity: event SoldToAMM(address seller, address market, uint8 position, uint256 amount, uint256 sUSDPaid, address susd, address asset)
func (_Thales *ThalesFilterer) ParseSoldToAMM(log types.Log) (*ThalesSoldToAMM, error) {
	event := new(ThalesSoldToAMM)
	if err := _Thales.contract.UnpackLog(event, "SoldToAMM", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
