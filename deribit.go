package rainbow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetMarkets(coin string) ([]Instruments, error) {
	baseURL := "https://deribit.com/api/v2/public/get_instruments?currency="
	opts := "&expired=false&kind=option"
	fmt.Println(baseURL + coin + opts)
	clt := http.Client{}
	resp, err := clt.Get(baseURL + coin + opts)
	if err != nil {
		return []Instruments{}, err
	}

	result := struct {
		Result []Instruments `json:"result"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []Instruments{}, fmt.Errorf("nft collections : %w", err)
	}
	return result.Result, nil
}

type listResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	Result  []Instruments `json:"result"`
	UsIn    int64         `json:"usIn"`
	UsOut   int64         `json:"usOut"`
	UsDiff  int           `json:"usDiff"`
	Testnet bool          `json:"testnet"`
}

type Instruments struct {
	TickSize             float64 `json:"tick_size"`
	TakerCommission      float64 `json:"taker_commission"`
	Strike               float64 `json:"strike"`
	SettlementPeriod     string  `json:"settlement_period"`
	QuoteCurrency        string  `json:"quote_currency"`
	OptionType           string  `json:"option_type"`
	MinTradeAmount       float64 `json:"min_trade_amount"`
	MakerCommission      float64 `json:"maker_commission"`
	Kind                 string  `json:"kind"`
	IsActive             bool    `json:"is_active"`
	InstrumentName       string  `json:"instrument_name"`
	ExpirationTimestamp  int64   `json:"expiration_timestamp"`
	CreationTimestamp    int64   `json:"creation_timestamp"`
	ContractSize         float64 `json:"contract_size"`
	BlockTradeCommission float64 `json:"block_trade_commission"`
	BaseCurrency         string  `json:"base_currency"`
}
