package deltaexchange

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	detlaProducts = "https://api.delta.exchange/v2/products?states=live&contract_types="
	deltaOrders   = "https://api.delta.exchange/v2/l2orderbook/"
	deltaTickers  = "https://api.delta.exchange/v2/tickers/"
)

type Provider struct{}

func (Provider) Name() string {
	return "Delta Exchange" // TODO real name but we use the short one to have a nice front for now "Delta Exchange"
}

// return the hour (UTC) at which the options expires
// 12:00 UTC
// should that be a "func (Provider)"?
func Hour() int {
	return 12
}

func (pro Provider) Options() ([]rainbow.Option, error) {
	options := []rainbow.Option{}
	products, err := queryProducts()
	if err != nil {
		return nil, fmt.Errorf("delta options collect : %w", err)
	}

	expiries := rainbow.Expiries(time.Now(), Hour())
	for _, p := range products {
		resp, err := http.Get(deltaOrders + p.Symbol)
		if err != nil {
			return nil, fmt.Errorf("delta options orders : %w", err)
		}

		defer resp.Body.Close()

		result := struct {
			Result OrderbookResult `json:"result"`
		}{}

		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf(" order book : %w", err)
		}

		if !rainbow.IsExpiryAvailable(expiries, p.SettlementTime) {
			continue
		}

		expiryStr := p.SettlementTime.Format("2006-01-02 15:04:05")

		strike, err := strconv.ParseFloat(p.StrikePrice, 64)
		if err != nil {
			return nil, fmt.Errorf(" strike conversion : %w", err)
		}
		// Ugly fix for front, the 100K strike is too long number so I will filter it out
		// TODO, reduce the size of the front character to fit more things
		// Anyway BTC is well below that so it's fine for now lol
		if strike > 75000 {
			continue
		}
		if p.ContractUnitCurrency == "BNB" || p.ContractUnitCurrency == "XRP" ||
			p.ContractUnitCurrency == "MATIC" || p.ContractUnitCurrency == "AVAX" {
			continue // We'll add them when those assets are on other providers
		}

		bids, asks, err := p.Orderbook(result.Result)
		if err != nil {
			return nil, fmt.Errorf(" Order Book issue : %w", err)
		}

		askIV, bidIV, markIV, greeks, err := p.Greeks()
		if err != nil {
			return nil, fmt.Errorf(" Greeks issue : %w", err)

		}

		options = append(options, rainbow.Option{
			Name:          p.Symbol,
			Type:          p.OptionsType(),
			Asset:         p.ContractUnitCurrency,
			Expiry:        expiryStr,
			Strike:        strike,
			ExchangeType:  "CEX",
			Chain:         "–",
			Layer:         "–",
			Provider:      pro.Name(),
			QuoteCurrency: p.QuotingAsset.Symbol,
			Bid:           bids,
			BidIV:         bidIV,
			Ask:           asks,
			AskIV:         askIV,
			MarketIV:      markIV,
			Greeks:        greeks,
		})
	}

	return options, nil
}

func queryProducts() (ProductResult, error) {
	options := []string{"put_options", "call_options"}
	var products ProductResult

	for _, o := range options {
		log.Print(detlaProducts + o)

		resp, err := http.Get(detlaProducts + o)
		if err != nil {
			return nil, fmt.Errorf(" Fetch options issue : %w", err)
		}
		defer resp.Body.Close()

		result := struct {
			Result ProductResult `json:"result"`
		}{}

		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf(" Query Options json issue : %w", err)
		}
		products = append(products, result.Result...)
	}
	return products, nil
}

func (p Product) Orderbook(orders OrderbookResult) ([]rainbow.Order, []rainbow.Order, error) {

	contractSize, err := strconv.ParseFloat(p.ContractValue, 64)
	if err != nil {
		return []rainbow.Order{}, []rainbow.Order{}, fmt.Errorf(" contract value conversion: %w", err)
	}

	bids, err := orderConversion(orders.Buy, contractSize)
	if err != nil {
		return []rainbow.Order{}, []rainbow.Order{}, err
	}
	asks, err := orderConversion(orders.Sell, contractSize)
	if err != nil {
		return []rainbow.Order{}, []rainbow.Order{}, err
	}
	return bids, asks, nil
}

func orderConversion(orders Orders, contractSize float64) ([]rainbow.Order, error) {
	rOrders := []rainbow.Order{}
	for _, o := range orders {
		price, err := strconv.ParseFloat(o.Price, 64)
		if err != nil {
			return []rainbow.Order{}, err
		}
		rOrders = append(rOrders, rainbow.Order{
			Price: price,
			Size:  float64(o.Size) * contractSize,
		})
	}
	return rOrders, nil
}

func (p Product) OptionsType() string {
	if p.ContractType == "call_options" {
		return "CALL"
	}
	return "PUT"
}

// Greeks retrieve the iv(in %) and greeks of the options
func (p Product) Greeks() (float64, float64, float64, rainbow.TheGreeks, error) {
	resp, err := http.Get(deltaTickers + p.Symbol)
	if err != nil {
		return 0, 0, 0, rainbow.TheGreeks{}, fmt.Errorf("delta options greeks request: %w", err)
	}

	defer resp.Body.Close()

	result := struct {
		Result TickerResult `json:"result"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, 0, rainbow.TheGreeks{}, fmt.Errorf("delta options greeks decode: %w", err)
	}

	greeks, err := convertGreeks(result.Result.Greeks)
	if err != nil {
		return 0, 0, 0, rainbow.TheGreeks{}, fmt.Errorf("delta options greeks conversion: %w", err)
	}

	askIV, bidIV, markIV, err := convertIV(result.Result.Quotes)
	if err != nil {
		return 0, 0, 0, rainbow.TheGreeks{}, fmt.Errorf("delta options iv conversion: %w", err)
	}

	//*100 to have the IV as %
	return 100 * askIV, 100 * bidIV, 100 * markIV, greeks, nil
}

func convertGreeks(d deltaGreeks) (rainbow.TheGreeks, error) {
	g := rainbow.TheGreeks{}
	greek, err := strconv.ParseFloat(d.Delta, 64)
	if err != nil {
		return rainbow.TheGreeks{}, fmt.Errorf(" delta : %w", err)
	}
	g.Delta = greek

	greek, err = strconv.ParseFloat(d.Gamma, 64)
	if err != nil {
		return rainbow.TheGreeks{}, fmt.Errorf(" gamma : %w", err)
	}
	g.Gamma = greek

	greek, err = strconv.ParseFloat(d.Theta, 64)
	if err != nil {
		return rainbow.TheGreeks{}, fmt.Errorf(" theta : %w", err)
	}
	g.Theta = greek

	greek, err = strconv.ParseFloat(d.Vega, 64)
	if err != nil {
		return rainbow.TheGreeks{}, fmt.Errorf(" vega : %w", err)
	}
	g.Vega = greek

	greek, err = strconv.ParseFloat(d.Rho, 64)
	if err != nil {
		return rainbow.TheGreeks{}, fmt.Errorf(" rho : %w", err)
	}
	g.Rho = greek

	return g, nil
}

func convertIV(q deltaQuotes) (float64, float64, float64, error) {
	askIV, err := strconv.ParseFloat(q.AskIv, 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf(" ask iv : %w", err)
	}
	bidIV, err := strconv.ParseFloat(q.BidIv, 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf(" bid iv : %w", err)
	}
	markIV, err := strconv.ParseFloat(q.MarkIv, 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf(" mark iv : %w", err)
	}
	return askIV, bidIV, markIV, nil

}

type (
	ProductResult []Product
	Product       struct {
		UIConfig struct {
			DefaultTradingViewCandle string    `json:"default_trading_view_candle"`
			LeverageSliderValues     []int     `json:"leverage_slider_values"`
			PriceClubbingValues      []float64 `json:"price_clubbing_values"`
			ShowBracketOrders        bool      `json:"show_bracket_orders"`
			SortPriority             int       `json:"sort_priority"`
		} `json:"ui_config,omitempty"`
		ContractValue                   string      `json:"contract_value"`
		InitialMarginScalingFactor      string      `json:"initial_margin_scaling_factor"`
		ImpactSize                      int         `json:"impact_size"`
		ID                              int         `json:"id"`
		BasisFactorMaxLimit             string      `json:"basis_factor_max_limit"`
		InsuranceFundMarginContribution string      `json:"insurance_fund_margin_contribution"`
		LiquidationPenaltyFactor        string      `json:"liquidation_penalty_factor"`
		DisruptionReason                interface{} `json:"disruption_reason"`
		AuctionFinishTime               interface{} `json:"auction_finish_time"`
		SettlingAsset                   struct {
			BaseWithdrawalFee   string      `json:"base_withdrawal_fee"`
			DepositStatus       string      `json:"deposit_status"`
			ID                  int         `json:"id"`
			InterestCredit      bool        `json:"interest_credit"`
			InterestSlabs       interface{} `json:"interest_slabs"`
			KycDepositLimit     string      `json:"kyc_deposit_limit"`
			KycWithdrawalLimit  string      `json:"kyc_withdrawal_limit"`
			MinWithdrawalAmount string      `json:"min_withdrawal_amount"`
			MinimumPrecision    int         `json:"minimum_precision"`
			Name                string      `json:"name"`
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
		MaxLeverageNotional  string      `json:"max_leverage_notional"`
		IsQuanto             bool        `json:"is_quanto"`
		ContractUnitCurrency string      `json:"contract_unit_currency"`
		TickSize             string      `json:"tick_size"`
		ContractType         string      `json:"contract_type"`
		SettlementPrice      interface{} `json:"settlement_price"`
		ShortDescription     string      `json:"short_description"`
		MaintenanceMargin    string      `json:"maintenance_margin"`
		MakerCommissionRate  string      `json:"maker_commission_rate"`
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
			ConstituentIndices interface{} `json:"constituent_indices"`
			Description        string      `json:"description"`
			HealthInterval     int         `json:"health_interval"`
			ID                 int         `json:"id"`
			ImpactSize         string      `json:"impact_size"`
			IndexType          string      `json:"index_type"`
			IsComposite        bool        `json:"is_composite"`
			PriceMethod        string      `json:"price_method"`
			QuotingAssetID     int         `json:"quoting_asset_id"`
			Symbol             string      `json:"symbol"`
			TickSize           string      `json:"tick_size"`
			UnderlyingAssetID  int         `json:"underlying_asset_id"`
		} `json:"spot_index"`
		MaintenanceMarginScalingFactor string `json:"maintenance_margin_scaling_factor"`
		UnderlyingAsset                struct {
			BaseWithdrawalFee   string      `json:"base_withdrawal_fee"`
			DepositStatus       string      `json:"deposit_status"`
			ID                  int         `json:"id"`
			InterestCredit      bool        `json:"interest_credit"`
			InterestSlabs       interface{} `json:"interest_slabs"`
			KycDepositLimit     string      `json:"kyc_deposit_limit"`
			KycWithdrawalLimit  string      `json:"kyc_withdrawal_limit"`
			MinWithdrawalAmount string      `json:"min_withdrawal_amount"`
			MinimumPrecision    int         `json:"minimum_precision"`
			Name                string      `json:"name"`
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
			BaseWithdrawalFee   string      `json:"base_withdrawal_fee"`
			DepositStatus       string      `json:"deposit_status"`
			ID                  int         `json:"id"`
			InterestCredit      bool        `json:"interest_credit"`
			InterestSlabs       interface{} `json:"interest_slabs"`
			KycDepositLimit     string      `json:"kyc_deposit_limit"`
			KycWithdrawalLimit  string      `json:"kyc_withdrawal_limit"`
			MinWithdrawalAmount string      `json:"min_withdrawal_amount"`
			MinimumPrecision    int         `json:"minimum_precision"`
			Name                string      `json:"name"`
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
		NotionalType     string      `json:"notional_type"`
		Description      string      `json:"description"`
		AuctionStartTime interface{} `json:"auction_start_time"`
		Symbol           string      `json:"symbol"`
		State            string      `json:"state"`
		ProductSpecs     struct {
			BackupVolExpiryTime         int     `json:"backup_vol_expiry_time"`
			MaxDeviationFromExternalVol float64 `json:"max_deviation_from_external_vol"`
			MaxVolatility               int     `json:"max_volatility"`
			MinVolatility               float64 `json:"min_volatility"`
			PremiumCommissionRate       float64 `json:"premium_commission_rate"`
			VolCalculationMethod        string  `json:"vol_calculation_method"`
			VolExpiryTime               int     `json:"vol_expiry_time"`
		} `json:"product_specs,omitempty"`
		FundingMethod     string      `json:"funding_method"`
		StrikePrice       string      `json:"strike_price"`
		PriceBand         string      `json:"price_band"`
		DefaultLeverage   string      `json:"default_leverage"`
		TradingStatus     string      `json:"trading_status"`
		BarrierPrice      interface{} `json:"barrier_price"`
		AnnualizedFunding string      `json:"annualized_funding"`
	}
)

type OrderbookResult struct {
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

type TickerResult struct {
	Close         float64     `json:"close"`
	ContractType  string      `json:"contract_type"`
	Greeks        deltaGreeks `json:"greeks"`
	High          float64     `json:"high"`
	Low           float64     `json:"low"`
	MarkPrice     string      `json:"mark_price"`
	MarkVol       string      `json:"mark_vol"`
	Oi            string      `json:"oi"`
	OiValue       string      `json:"oi_value"`
	OiValueSymbol string      `json:"oi_value_symbol"`
	OiValueUsd    string      `json:"oi_value_usd"`
	Open          float64     `json:"open"`
	PriceBand     struct {
		LowerLimit string `json:"lower_limit"`
		UpperLimit string `json:"upper_limit"`
	} `json:"price_band"`
	ProductID      int         `json:"product_id"`
	Quotes         deltaQuotes `json:"quotes"`
	Size           int         `json:"size"`
	SpotPrice      string      `json:"spot_price"`
	StrikePrice    string      `json:"strike_price"`
	Symbol         string      `json:"symbol"`
	Timestamp      int64       `json:"timestamp"`
	Turnover       float64     `json:"turnover"`
	TurnoverSymbol string      `json:"turnover_symbol"`
	TurnoverUsd    float64     `json:"turnover_usd"`
	Volume         float64     `json:"volume"`
}

type deltaGreeks struct {
	Delta string `json:"delta"`
	Gamma string `json:"gamma"`
	Rho   string `json:"rho"`
	Theta string `json:"theta"`
	Vega  string `json:"vega"`
}

type deltaQuotes struct {
	AskIv   string `json:"ask_iv"`
	AskSize string `json:"ask_size"`
	BestAsk string `json:"best_ask"`
	BestBid string `json:"best_bid"`
	BidIv   string `json:"bid_iv"`
	BidSize string `json:"bid_size"`
	MarkIv  string `json:"mark_iv"`
}
