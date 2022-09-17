// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package deltaexchange

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/garcon/gg"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone("Delta")

const (
	deltaProducts  = "https://api.delta.exchange/v2/products?states=live&contract_types=put_options,call_options"
	deltaOrders    = "https://api.delta.exchange/v2/l2orderbook/"
	maxBytesToRead = 20_000_000 // Prevent wasting memory/CPU when receiving an abnormally huge response from Delta API
)

type Provider struct{}

func (Provider) Name() string {
	return "Delta Exchange" // TODO real name but we use the short one to have a nice front for now "Delta Exchange"
}

// Hour is the time in UTC at which the options expires
// TODO: should we add a func in Provider interface?
const Hour = 12 // 12:00 UTC

func (pro Provider) Options() ([]rainbow.Option, error) {
	products, err := queryProducts()
	if err != nil {
		return nil, err
	}

	options := make([]rainbow.Option, 0, len(products))

	expiries := rainbow.Expiries(time.Now(), Hour)
	for i := range products {
		p := &products[i]

		if !rainbow.IsExpiryAvailable(expiries, p.SettlementTime) {
			continue
		}

		expiryStr := p.SettlementTime.Format("2006-01-02 15:04:05")

		strike, err := strconv.ParseFloat(p.StrikePrice, 64)
		if err != nil {
			return nil, fmt.Errorf("strike ParseFloat: %w", err)
		}

		// Ugly fix for front, the 100K strike is too long number so I will filter it out
		// TODO, reduce the size of the front character to fit more things
		if strike > 75000 {
			continue
		}

		if p.ContractUnitCurrency == "BNB" || p.ContractUnitCurrency == "XRP" ||
			p.ContractUnitCurrency == "MATIC" || p.ContractUnitCurrency == "AVAX" {
			continue // We'll add them when those assets are on other providers
		}

		bids, asks, err := p.orderbook()
		if err != nil {
			return nil, err
		}

		options = append(options, rainbow.Option{
			Name:          p.Symbol,
			Type:          p.optionsType(),
			Asset:         p.ContractUnitCurrency,
			Expiry:        expiryStr,
			Strike:        strike,
			ExchangeType:  "CEX",
			Chain:         "–",
			Layer:         "–",
			Provider:      pro.Name(),
			QuoteCurrency: p.QuotingAsset.Symbol,
			Bid:           bids,
			Ask:           asks,
		})
	}

	return options, nil
}

func queryProducts() (ProductResult, error) {
	log.Info(deltaProducts)
	resp, err := http.Get(deltaProducts)
	if err != nil {
		return nil, fmt.Errorf("queryProducts GET: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Result ProductResult `json:"result"`
	}
	// this DeltaExchange response can be more than 100 KB, thus we accept much more
	if err = gg.DecodeJSONResponse(resp, &result, maxBytesToRead); err != nil {
		return nil, fmt.Errorf("queryProducts decode: %w", err)
	}

	return result.Result, nil
}

type (
	ProductResult []Product

	Product struct {
		UIConfig struct {
			DefaultTradingViewCandle string    `json:"default_trading_view_candle"`
			LeverageSliderValues     []int     `json:"leverage_slider_values"`
			PriceClubbingValues      []float64 `json:"price_clubbing_values"`
			ShowBracketOrders        bool      `json:"show_bracket_orders"`
			SortPriority             int       `json:"sort_priority"`
		} `json:"ui_config,omitempty"`
		ContractValue                   string `json:"contract_value"`
		InitialMarginScalingFactor      string `json:"initial_margin_scaling_factor"`
		ImpactSize                      int    `json:"impact_size"`
		ID                              int    `json:"id"`
		BasisFactorMaxLimit             string `json:"basis_factor_max_limit"`
		InsuranceFundMarginContribution string `json:"insurance_fund_margin_contribution"`
		LiquidationPenaltyFactor        string `json:"liquidation_penalty_factor"`
		DisruptionReason                any    `json:"disruption_reason"`
		AuctionFinishTime               any    `json:"auction_finish_time"`
		SettlingAsset                   struct {
			BaseWithdrawalFee   string `json:"base_withdrawal_fee"`
			DepositStatus       string `json:"deposit_status"`
			ID                  int    `json:"id"`
			InterestCredit      bool   `json:"interest_credit"`
			InterestSlabs       any    `json:"interest_slabs"`
			KycDepositLimit     string `json:"kyc_deposit_limit"`
			KycWithdrawalLimit  string `json:"kyc_withdrawal_limit"`
			MinWithdrawalAmount string `json:"min_withdrawal_amount"`
			MinimumPrecision    int    `json:"minimum_precision"`
			Name                string `json:"name"`
			Networks            []struct {
				BaseWithdrawalFee     string `json:"base_withdrawal_fee"`
				DepositStatus         string `json:"deposit_status"`
				MemoRequired          bool   `json:"memo_required"`
				Network               string `json:"network"`
				VariableWithdrawalFee string `json:"variable_withdrawal_fee"`
				WithdrawalStatus      string `json:"withdrawal_status"`
			} `json:"networks"`
			Precision             int    `json:"precision"`
			SortPriority          int    `json:"sort_priority"`
			Symbol                string `json:"symbol"`
			VariableWithdrawalFee string `json:"variable_withdrawal_fee"`
			WithdrawalStatus      string `json:"withdrawal_status"`
		} `json:"settling_asset"`
		MaxLeverageNotional  string `json:"max_leverage_notional"`
		IsQuanto             bool   `json:"is_quanto"`
		ContractUnitCurrency string `json:"contract_unit_currency"`
		TickSize             string `json:"tick_size"`
		ContractType         string `json:"contract_type"`
		SettlementPrice      any    `json:"settlement_price"`
		ShortDescription     string `json:"short_description"`
		MaintenanceMargin    string `json:"maintenance_margin"`
		MakerCommissionRate  string `json:"maker_commission_rate"`
		SpotIndex            struct {
			Config struct {
				QuotingAsset    string `json:"quoting_asset"`
				ServiceID       int    `json:"service_id"`
				UnderlyingAsset string `json:"underlying_asset"`
			} `json:"config"`
			ConstituentExchanges []struct {
				Exchange       string `json:"exchange"`
				HealthInterval int    `json:"health_interval"`
				HealthPriority int    `json:"health_priority"`
				Weight         int    `json:"weight"`
			} `json:"constituent_exchanges"`
			ConstituentIndices any    `json:"constituent_indices"`
			Description        string `json:"description"`
			HealthInterval     int    `json:"health_interval"`
			ID                 int    `json:"id"`
			ImpactSize         string `json:"impact_size"`
			IndexType          string `json:"index_type"`
			IsComposite        bool   `json:"is_composite"`
			PriceMethod        string `json:"price_method"`
			QuotingAssetID     int    `json:"quoting_asset_id"`
			Symbol             string `json:"symbol"`
			TickSize           string `json:"tick_size"`
			UnderlyingAssetID  int    `json:"underlying_asset_id"`
		} `json:"spot_index"`
		MaintenanceMarginScalingFactor string `json:"maintenance_margin_scaling_factor"`
		UnderlyingAsset                struct {
			BaseWithdrawalFee   string `json:"base_withdrawal_fee"`
			DepositStatus       string `json:"deposit_status"`
			ID                  int    `json:"id"`
			InterestCredit      bool   `json:"interest_credit"`
			InterestSlabs       any    `json:"interest_slabs"`
			KycDepositLimit     string `json:"kyc_deposit_limit"`
			KycWithdrawalLimit  string `json:"kyc_withdrawal_limit"`
			MinWithdrawalAmount string `json:"min_withdrawal_amount"`
			MinimumPrecision    int    `json:"minimum_precision"`
			Name                string `json:"name"`
			Networks            []struct {
				BaseWithdrawalFee     string `json:"base_withdrawal_fee"`
				DepositStatus         string `json:"deposit_status"`
				MemoRequired          bool   `json:"memo_required"`
				Network               string `json:"network"`
				VariableWithdrawalFee string `json:"variable_withdrawal_fee"`
				WithdrawalStatus      string `json:"withdrawal_status"`
			} `json:"networks"`
			Precision             int    `json:"precision"`
			SortPriority          int    `json:"sort_priority"`
			Symbol                string `json:"symbol"`
			VariableWithdrawalFee string `json:"variable_withdrawal_fee"`
			WithdrawalStatus      string `json:"withdrawal_status"`
		} `json:"underlying_asset"`
		TakerCommissionRate string    `json:"taker_commission_rate"`
		LaunchTime          time.Time `json:"launch_time"`
		PositionSizeLimit   int       `json:"position_size_limit"`
		SettlementTime      time.Time `json:"settlement_time"`
		InitialMargin       string    `json:"initial_margin"`
		QuotingAsset        struct {
			BaseWithdrawalFee   string `json:"base_withdrawal_fee"`
			DepositStatus       string `json:"deposit_status"`
			ID                  int    `json:"id"`
			InterestCredit      bool   `json:"interest_credit"`
			InterestSlabs       any    `json:"interest_slabs"`
			KycDepositLimit     string `json:"kyc_deposit_limit"`
			KycWithdrawalLimit  string `json:"kyc_withdrawal_limit"`
			MinWithdrawalAmount string `json:"min_withdrawal_amount"`
			MinimumPrecision    int    `json:"minimum_precision"`
			Name                string `json:"name"`
			Networks            []struct {
				BaseWithdrawalFee     string `json:"base_withdrawal_fee"`
				DepositStatus         string `json:"deposit_status"`
				MemoRequired          bool   `json:"memo_required"`
				Network               string `json:"network"`
				VariableWithdrawalFee string `json:"variable_withdrawal_fee"`
				WithdrawalStatus      string `json:"withdrawal_status"`
			} `json:"networks"`
			Precision             int    `json:"precision"`
			SortPriority          int    `json:"sort_priority"`
			Symbol                string `json:"symbol"`
			VariableWithdrawalFee string `json:"variable_withdrawal_fee"`
			WithdrawalStatus      string `json:"withdrawal_status"`
		} `json:"quoting_asset"`
		NotionalType     string `json:"notional_type"`
		Description      string `json:"description"`
		AuctionStartTime any    `json:"auction_start_time"`
		Symbol           string `json:"symbol"`
		State            string `json:"state"`
		ProductSpecs     struct {
			BackupVolExpiryTime         int     `json:"backup_vol_expiry_time"`
			MaxDeviationFromExternalVol float64 `json:"max_deviation_from_external_vol"`
			MaxVolatility               int     `json:"max_volatility"`
			MinVolatility               float64 `json:"min_volatility"`
			PremiumCommissionRate       float64 `json:"premium_commission_rate"`
			VolCalculationMethod        string  `json:"vol_calculation_method"`
			VolExpiryTime               int     `json:"vol_expiry_time"`
		} `json:"product_specs,omitempty"`
		FundingMethod     string `json:"funding_method"`
		StrikePrice       string `json:"strike_price"`
		PriceBand         string `json:"price_band"`
		DefaultLeverage   string `json:"default_leverage"`
		TradingStatus     string `json:"trading_status"`
		BarrierPrice      any    `json:"barrier_price"`
		AnnualizedFunding string `json:"annualized_funding"`
	}
)

type OrderBookResult struct {
	Buy           Orders `json:"buy"`
	LastUpdatedAt int64  `json:"last_updated_at"`
	Sell          Orders `json:"sell"`
	Symbol        string `json:"symbol"`
}

type Orders []struct {
	Depth string `json:"depth"`
	Price string `json:"price"`
	Size  int    `json:"size"`
}

func (p *Product) orderbook() (bid, ask []rainbow.Order, err error) {
	resp, err := http.Get(deltaOrders + p.Symbol)
	if err != nil {
		return nil, nil, fmt.Errorf("book GET: %w", err)
	}
	defer resp.Body.Close()

	var orders struct {
		Orders OrderBookResult `json:"result"`
	}
	if err = gg.DecodeJSONResponse(resp, &orders); err != nil {
		return nil, nil, fmt.Errorf("book decode response: %w", err)
	}

	contractSize, err := strconv.ParseFloat(p.ContractValue, 64)
	if err != nil {
		return nil, nil, fmt.Errorf("book ParseFloat: %w", err)
	}

	bid, err = orderConversion(orders.Orders.Buy, contractSize)
	if err != nil {
		return nil, nil, err
	}

	ask, err = orderConversion(orders.Orders.Sell, contractSize)
	if err != nil {
		return nil, nil, err
	}

	return bid, ask, nil
}

func (p *Product) optionsType() string {
	if p.ContractType == "call_options" {
		return "CALL"
	}
	return "PUT"
}

func orderConversion(orders Orders, contractSize float64) ([]rainbow.Order, error) {
	rOrders := make([]rainbow.Order, 0, len(orders))

	for _, o := range orders {
		price, err := strconv.ParseFloat(o.Price, 64)
		if err != nil {
			return nil, err
		}
		rOrders = append(rOrders, rainbow.Order{
			Price: price,
			Size:  float64(o.Size) * contractSize,
		})
	}

	return rOrders, nil
}
