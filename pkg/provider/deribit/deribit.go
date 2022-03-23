// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package deribit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

type Provider struct{}

func (Provider) Name() string {
	return "Deribit"
}

func (Provider) Options() ([]rainbow.Option, error) {
	instruments, err := query("BTC")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	optionsBTC, err := fillOptions(instruments, 5)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	instruments, err = query("ETH")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	optionsETH, err := fillOptions(instruments, 5)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return append(optionsBTC, optionsETH...), nil
}

func query(coin string) ([]instrument, error) {
	baseURL := "https://deribit.com/api/v2/public/get_instruments?currency="
	opts := "&expired=false&kind=option"
	log.Print(baseURL + coin + opts)

	resp, err := http.Get(baseURL + coin + opts)
	if err != nil {
		return []instrument{}, err
	}

	defer resp.Body.Close()

	result := struct {
		Result []instrument `json:"result"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []instrument{}, fmt.Errorf("deribit options collect : %w", err)
	}

	return filterTooFar(result.Result), nil
}

type instrument struct {
	OptionType           string  `json:"option_type"`
	InstrumentName       string  `json:"instrument_name"`
	Kind                 string  `json:"kind"`
	SettlementPeriod     string  `json:"settlement_period"`
	QuoteCurrency        string  `json:"quote_currency"`
	BaseCurrency         string  `json:"base_currency"`
	MinTradeAmount       float64 `json:"min_trade_amount"`
	MakerCommission      float64 `json:"maker_commission"`
	Strike               float64 `json:"strike"`
	TickSize             float64 `json:"tick_size"`
	TakerCommission      float64 `json:"taker_commission"`
	ExpirationTimestamp  int64   `json:"expiration_timestamp"`
	CreationTimestamp    int64   `json:"creation_timestamp"`
	ContractSize         float64 `json:"contract_size"`
	BlockTradeCommission float64 `json:"block_trade_commission"`
	IsActive             bool    `json:"is_active"`
}

func filterTooFar(instruments []instrument) (filtered []instrument) {
	for _, i := range instruments {
		seconds := i.ExpirationTimestamp / 1000
		ns := (i.ExpirationTimestamp % 1000) * 1000_000
		expiryTime := time.Unix(seconds, ns).UTC()
		// we should filter by taking what is available elsewhere and then
		// only fetch those
		if isExpiryAvailable(expiryTime) && isStrikeAvailable(i) {
			filtered = append(filtered, i)
		}
	}

	return filtered
}

//TODO change this quick and dirty way of filtering date from deribit.
func isExpiryAvailable(expiry time.Time) bool {
	dates := []string{
		"2022-03-25T08:00:00Z",
		"2022-04-01T08:00:00Z",
		"2022-04-08T08:00:00Z",
		"2022-04-29T08:00:00Z",
		"2022-06-24T08:00:00Z",
	}
	for _, d := range dates {
		t, _ := time.Parse(time.RFC3339, d)
		if expiry.Equal(t) {
			return true
		}
	}

	return false
}

// TODO change this quick and dirty way of filtering strikes from deribit.
func isStrikeAvailable(i instrument) bool {
	ethStrike := []float64{1800, 2000, 2100, 2200, 2300, 2400, 2500, 2600, 2700, 2800, 2900, 3000, 3100, 3200, 3300, 3400, 3500, 3800}
	btcStrike := []float64{20000, 25000, 29000, 30000, 32000, 33000, 34000, 35000, 36000, 37000, 38000, 39000, 40000, 41000, 42000, 43000, 44000, 45000, 46000, 47000, 48000, 50000}
	strikes := ethStrike

	if i.BaseCurrency == "BTC" {
		strikes = btcStrike
	}

	for _, s := range strikes {
		// if strike on deribit around strike anywhere else, keep it
		if i.Strike >= s*0.99 && i.Strike <= s*1.01 {
			return true
		}
	}

	return false
}

func fillOptions(instruments []instrument, depth uint32) ([]rainbow.Option, error) {
	options := []rainbow.Option{}
	baseURL := "https://www.deribit.com/api/v2/public/get_order_book?depth=" + strconv.Itoa(int(depth)) + "&instrument_name="

	for _, i := range instruments {
		resp, err := http.Get(baseURL + i.InstrumentName)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		result := struct {
			Result OrderBook `json:"result"`
		}{}

		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf(" order book : %w", err)
		}

		// API doc: https://docs.deribit.com/#public-get_index_price_names
		// ExpirationTimestamp = The time when the instrument will expire (milliseconds since the UNIX epoch)
		seconds := i.ExpirationTimestamp / 1000
		ns := (i.ExpirationTimestamp % 1000) * 1000_000
		expiryTime := time.Unix(seconds, ns).UTC()
		expiryStr := expiryTime.Format("2006-01-02 15:04:05")

		bids := normalizeOrders(result.Result.Bids, result.Result.IndexPrice)
		sort.Slice(bids, func(i, j int) bool {
			return bids[i].Price > bids[j].Price
		})

		asks := normalizeOrders(result.Result.Asks, result.Result.IndexPrice)
		sort.Slice(asks, func(i, j int) bool {
			return asks[i].Price < asks[j].Price
		})

		options = append(options, rainbow.Option{
			Name:          i.InstrumentName,
			Type:          strings.ToUpper(i.OptionType),
			Asset:         i.BaseCurrency,
			Expiry:        expiryStr,
			Strike:        i.Strike,
			ExchangeType:  "CEX",
			Chain:         "–",
			Layer:         "–",
			Provider:      "Deribit",
			QuoteCurrency: i.QuoteCurrency,
			Bid:           bids,
			Ask:           asks,
		})
	}

	return options, nil
}

type OrderBook struct {
	InstrumentName  string      `json:"instrument_name"`
	UnderlyingIndex string      `json:"underlying_index"`
	State           string      `json:"state"`
	Asks            [][]float64 `json:"asks"`
	Bids            [][]float64 `json:"bids"`
	Greeks          struct {
		Vega  float64 `json:"vega"`
		Theta float64 `json:"theta"`
		Rho   float64 `json:"rho"`
		Gamma float64 `json:"gamma"`
		Delta float64 `json:"delta"`
	} `json:"greeks"`
	Stats struct {
		Volume      float64 `json:"volume"`
		PriceChange float64 `json:"price_change"`
		Low         float64 `json:"low"`
		High        float64 `json:"high"`
	} `json:"stats"`
	UnderlyingPrice        float64 `json:"underlying_price"`
	MaxPrice               float64 `json:"max_price"`
	MarkPrice              float64 `json:"mark_price"`
	MarkIv                 float64 `json:"mark_iv"`
	LastPrice              float64 `json:"last_price"`
	InterestRate           float64 `json:"interest_rate"`
	MinPrice               float64 `json:"min_price"`
	IndexPrice             float64 `json:"index_price"`
	OpenInterest           float64 `json:"open_interest"`
	EstimatedDeliveryPrice float64 `json:"estimated_delivery_price"`
	ChangeID               int64   `json:"change_id"`
	SettlementPrice        float64 `json:"settlement_price"`
	BidIv                  float64 `json:"bid_iv"`
	BestBidPrice           float64 `json:"best_bid_price"`
	BestBidAmount          float64 `json:"best_bid_amount"`
	BestAskPrice           float64 `json:"best_ask_price"`
	BestAskAmount          float64 `json:"best_ask_amount"`
	Timestamp              int64   `json:"timestamp"`
	AskIv                  float64 `json:"ask_iv"`
}

// Prices are not in $ but in crypto so we need the coin (index) price to multiply
// and get the USD price
func normalizeOrders(orders [][]float64, assetPrice float64) []rainbow.Order {
	// if there is no offer, send price=0.0, quant=0.0
	// hopefully we never an array of empty array
	if len(orders) == 0 {
		return []rainbow.Order{{Price: 0.0, Size: 0.0}}
	}

	offers := make([]rainbow.Order, 0, len(orders))
	for _, ord := range orders {
		offers = append(offers, rainbow.Order{Price: ord[0] * assetPrice, Size: ord[1]})
	}

	return offers
}
