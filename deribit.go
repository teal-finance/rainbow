package rainbow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func GetOrderBook(instruments []Instruments, depth uint32) ([]Options, error) {
	orderBook := []Options{}
	//deribitOrderBook := []DeribitOrderBook{}
	baseURL := "https://www.deribit.com/api/v2/public/get_order_book?depth=" + strconv.Itoa(int(depth)) + "&instrument_name="
	clt := http.Client{}

	for _, i := range instruments {
		resp, err := clt.Get(baseURL + i.InstrumentName)
		if err != nil {
			return []Options{}, err
		}
		result := struct {
			Result DeribitOrderBook `json:"result"`
		}{}
		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return []Options{}, fmt.Errorf(" order book : %w", err)
		}

		o := Options{}
		o.Asset = i.BaseCurrency
		o.Chain = "None"
		o.Layer = "None"
		o.ExchangeType = "CEX"
		o.Provider = "Deribit"
		o.ExpirationTimestamp = i.ExpirationTimestamp
		o.Strike = i.Strike
		o.Name = i.InstrumentName
		o.Type = strings.ToUpper(i.OptionType)
		o.Offers = BidsAsksToOffers(result.Result.Bids, "BUY", i.QuoteCurrency)
		o.Offers = append(o.Offers, BidsAsksToOffers(result.Result.Asks, "SELL", i.QuoteCurrency)...)
		//fmt.Println("o ", o)

		orderBook = append(orderBook, o)

	}

	return orderBook, nil
}

type DeribitOrderBook struct {
	UnderlyingPrice float64 `json:"underlying_price"`
	UnderlyingIndex string  `json:"underlying_index"`
	Timestamp       int64   `json:"timestamp"`
	Stats           struct {
		Volume      float64 `json:"volume"`
		PriceChange float64 `json:"price_change"`
		Low         float64 `json:"low"`
		High        float64 `json:"high"`
	} `json:"stats"`
	State           string  `json:"state"`
	SettlementPrice float64 `json:"settlement_price"`
	OpenInterest    float64 `json:"open_interest"`
	MinPrice        float64 `json:"min_price"`
	MaxPrice        float64 `json:"max_price"`
	MarkPrice       float64 `json:"mark_price"`
	MarkIv          float64 `json:"mark_iv"`
	LastPrice       float64 `json:"last_price"`
	InterestRate    float64 `json:"interest_rate"`
	InstrumentName  string  `json:"instrument_name"`
	IndexPrice      float64 `json:"index_price"`
	Greeks          struct {
		Vega  float64 `json:"vega"`
		Theta float64 `json:"theta"`
		Rho   float64 `json:"rho"`
		Gamma float64 `json:"gamma"`
		Delta float64 `json:"delta"`
	} `json:"greeks"`
	EstimatedDeliveryPrice float64     `json:"estimated_delivery_price"`
	ChangeID               int64       `json:"change_id"`
	Bids                   [][]float64 `json:"bids"`
	BidIv                  float64     `json:"bid_iv"`
	BestBidPrice           float64     `json:"best_bid_price"`
	BestBidAmount          float64     `json:"best_bid_amount"`
	BestAskPrice           float64     `json:"best_ask_price"`
	BestAskAmount          float64     `json:"best_ask_amount"`
	Asks                   [][]float64 `json:"asks"`
	AskIv                  float64     `json:"ask_iv"`
}

func BidsAsksToOffers(orders [][]float64, side, quote string) []Offer {
	offers := []Offer{}
	for _, ord := range orders {
		offers = append(offers, Offer{side, ord[1], ord[0], quote})

	}
	return offers

}
