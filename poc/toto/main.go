package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/teal-finance/rainbow/pkg/provider/lyra"
)

const alchemy = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m"
const optimism = "https://mainnet.optimism.io"
const deltaex = "https://api.delta.exchange/v2/products?contract_types="
const deltaorder = "https://api.delta.exchange/v2/l2orderbook/"

func main() {
	dex()
}

func dex() {

	log.Print(deltaex + "put_options")

	resp, err := http.Get(deltaex + "put_options")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	//spew.Dump(resp.Body)

	result := struct {
		Result response `json:"result"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	//spew.Dump(result.Result[0])

	log.Print(deltaorder + result.Result[0].Symbol)

	resp, err = http.Get(deltaorder + result.Result[0].Symbol)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	//spew.Dump(resp.Body)

	auto := struct {
		Orders Result `json:"result"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&auto); err != nil {
		log.Fatal(err)
	}
	spew.Dump(auto)

}

func tl() {
	p := lyra.Provider{}
	options, err := p.Options()
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(options)
}

func rpc(count int) {
	client, err := ethclient.Dial(alchemy) //+ INFURA_KEY)
	if err != nil {
		log.Fatal(err)
	}
	client2, err := ethclient.Dial(optimism) //+ INFURA_KEY)
	if err != nil {
		log.Fatal(err)
	}
	amount := big.Int{}
	amount.SetString("1000000000000000000", 10) //1000000000000000000
	v := "0x43592bffCF14f1e0A096091E125f023B2ccC2525"
	address := common.HexToAddress(v)
	viewer, err := lyra.NewLyrap(address, client)
	if err != nil {
		log.Fatal(err)
	}
	viewer2, err := lyra.NewLyrap(address, client2)
	if err != nil {
		log.Fatal(err)
	}
	board := big.NewInt(17)
	oldtbd, _ := viewer.GetOpenPremiumsForBoard(&bind.CallOpts{}, board, 0, &amount)
	oldtbd2, _ := viewer2.GetOpenPremiumsForBoard(&bind.CallOpts{}, board, 0, &amount)

	for i := 1; i < count; i++ {

		tbd, err := viewer.GetOpenPremiumsForBoard(&bind.CallOpts{}, board, 0, &amount)
		tbd2, err2 := viewer2.GetOpenPremiumsForBoard(&bind.CallOpts{}, board, 0, &amount)
		oldtbd = tbd
		oldtbd2 = tbd2
		if err != nil {
			fmt.Println(i)
			//spe
			log.Fatal(err)
		}
		if err2 != nil {
			fmt.Println(i)
			//spe
			log.Fatal(err2)
		}
		tbd, err = viewer.GetOpenPremiumsForBoard(&bind.CallOpts{}, board, 3, &amount)
		tbd2, err2 = viewer2.GetOpenPremiumsForBoard(&bind.CallOpts{}, board, 3, &amount)
		oldtbd = tbd
		oldtbd2 = tbd2
		if err != nil {
			fmt.Println(i)
			//spe
			log.Fatal(err)
		}
		if err2 != nil {
			fmt.Println(i)
			//spe
			log.Fatal(err2)
		}
	}

	spew.Dump(oldtbd)
	spew.Dump(oldtbd2)

	fmt.Println(count, "2 OK")
}

type response []struct {
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

type Result struct {
	Buy []struct {
		Depth string `json:"depth"`
		Price string `json:"price"`
		Size  int    `json:"size"`
	} `json:"buy"`
	LastUpdatedAt int64 `json:"last_updated_at"`
	Sell          []struct {
		Depth string `json:"depth"`
		Price string `json:"price"`
		Size  int    `json:"size"`
	} `json:"sell"`
	Symbol string `json:"symbol"`
}
