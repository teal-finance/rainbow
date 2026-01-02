// Copyright 2023 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package lyrav2

import (
	"strconv"
	"strings"
	"time"

	"github.com/LynxAIeu/emo"
	"github.com/LynxAIeu/garcon"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone(name)

const (
	// https://api.lyra.finance/public/get_instruments?currency=BTC&instrument_type=option&expired=false
	baseURL = "https://api.lyra.finance/public/"
	uiURL   = "https://www.lyra.finance/options/"
	name    = "Lyrav2"
)

type Provider struct {
	ar garcon.AdaptiveRate
}

func (Provider) Name() string {
	return name
}

// adaptiveMinSleepTime to rate limit the lyra API.
// source: https://v2-docs.lyra.finance/reference/rate-limits
const adaptiveMinSleepTime = 25 * time.Millisecond

// Hour at which the options expires = 8:00 UTC.
const Hour = 8

// maxBytesToRead prevents wasting memory/CPU when receiving an abnormally huge response from Aevo API.
// we put the same param as Deribit.
const maxBytesToRead = 2_000_000

func (p Provider) Options() ([]rainbow.Option, error) {
	if p.ar.Name == "" {
		p.ar = garcon.NewAdaptiveRate(name, adaptiveMinSleepTime)
	}

	result, err := p.query()
	if err != nil {
		return nil, err
	}
	r := result.Instruments // [0:3]

	options, err := p.fillOptions(r)
	if err != nil {
		return nil, log.Error("filloptions", err).Err()
	}
	return options, nil
}

func (p Provider) query() (GetInstrumentsResults, error) {
	const url = baseURL + "get_instruments?currency=BTC&instrument_type=option&expired=false"
	log.Info(name + " " + url)

	var results GetInstrumentsResults
	err := p.ar.Get("", url, &results, maxBytesToRead)
	if err != nil {
		return GetInstrumentsResults{}, log.Error("query", url, err).Err()
	}

	return results, nil
}

func (p Provider) fillOptions(instruments []Instrument) ([]rainbow.Option, error) {
	url := ""
	var ticker GetTickerResult
	var err error
	options := []rainbow.Option{}
	var s, iv, index, OpenInterest float64
	var bidSize, bidPrice, askSize, askPrice float64
	// askiv := 0.0
	// bidiv := 0.0
	optionType := ""
	for _, i := range instruments {
		if !strings.Contains(i.InstrumentName, "20240112") {
			continue
		}
		url = baseURL + "get_ticker?instrument_name=" + i.InstrumentName
		err = p.ar.Get("", url, &ticker, maxBytesToRead)
		if err != nil {
			return nil, log.Error("Get Ticker", i.InstrumentName, err).Err()
		}

		switch i.Details.OptionType {
		case "C":
			optionType = "CALL"
		case "P":
			optionType = "PUT"
		default:
			log.Warn("Unknown option type")
			optionType = "???"
		}
		s, err = strconv.ParseFloat(i.Details.Strike, 64)
		if err != nil {
			return []rainbow.Option{}, log.Error("conversion failure", s, err).Err()
		}

		iv, err = strconv.ParseFloat(ticker.Result.OptionPricing.Iv, 64)
		if err != nil {
			return []rainbow.Option{}, log.Error("conversion failure", s, err).Err()
		}
		index, err = strconv.ParseFloat(ticker.Result.IndexPrice, 64)
		if err != nil {
			return []rainbow.Option{}, log.Error("conversion failure", s, err).Err()
		}
		OpenInterest, err = strconv.ParseFloat(ticker.Result.Stats.OpenInterest, 64)
		if err != nil {
			return []rainbow.Option{}, log.Error("conversion failure", s, err).Err()
		}
		bidPrice, err = strconv.ParseFloat(ticker.Result.BestBidPrice, 64)
		if err != nil {
			return []rainbow.Option{}, log.Error("conversion failure", s, err).Err()
		}
		bidSize, err = strconv.ParseFloat(ticker.Result.BestBidAmount, 64)
		if err != nil {
			return []rainbow.Option{}, log.Error("conversion failure", s, err).Err()
		}
		askPrice, err = strconv.ParseFloat(ticker.Result.BestAskPrice, 64)
		if err != nil {
			return []rainbow.Option{}, log.Error("conversion failure", s, err).Err()
		}
		askSize, err = strconv.ParseFloat(ticker.Result.BestAskAmount, 64)
		if err != nil {
			return []rainbow.Option{}, log.Error("conversion failure", s, err).Err()
		}
		infos := strings.Split(i.InstrumentName, "-")

		expiry := time.Unix(int64(i.Details.Expiry), 0).UTC().Format("2006-01-02 15:04:05")

		options = append(options, rainbow.Option{
			Name:            i.InstrumentName,
			Type:            strings.ToUpper(optionType),
			Asset:           ticker.Result.BaseCurrency,
			UnderlyingAsset: ticker.Result.BaseCurrency,
			Strike:          s,
			Expiry:          expiry,
			ExpiryTime:      i.Details.Expiry,
			ExchangeType:    "DEX",
			Chain:           "Ethereum",
			Layer:           "L2",
			LayerName:       "Lyra",
			Provider:        name,
			UnderlyingQuote: ticker.Result.QuoteCurrency,
			QuoteCurrency:   "USD",
			URL:             uiURL + "btc?drawer=false&buy=true&expiry=" + infos[1] + "&option=" + i.InstrumentName,
			Bid: []rainbow.Order{{
				Price: bidPrice,
				Size:  bidSize,
			}},
			// BidIV:,

			Ask: []rainbow.Order{{
				Price: askPrice,
				Size:  askSize,
			}},
			// AskIV:,
			MarketIV: iv * 100,
			// TODO greeks
			// Greeks:
			OpenInterest: OpenInterest * index,
			ProtocolID:   i.InstrumentName,
		})
	}
	// spew.Dump(options)

	return options, nil
}

/*
func bidAsksToOrders(orders [][]string) ([]rainbow.Order, error) {
	if len(orders) == 0 {
		return []rainbow.Order{}, nil
	}
	convertedOrders := []rainbow.Order{}
	p := 0.0 // price

	s := 0.0 // size
	// i:=0.0 iv
	var err error
	for _, o := range orders {
		p, err = convert(o[0])
		if err != nil {
			return nil, log.Error("bidAsksToOrders", err).Err()
		}
		s, err = convert(o[1])
		if err != nil {
			return nil, log.Error("bidAsksToOrders", err).Err()
		}

		convertedOrders = append(convertedOrders, rainbow.Order{
			Price: p,
			Size:  s,
		})
		// i,err=convert(o[2])

	}

	return convertedOrders, nil
}

func translateGreeks(g greeks) (rainbow.TheGreeks, error) {
	// TODO
	return rainbow.TheGreeks{}, nil
}

*/

type GetInstrumentsResults struct {
	Instruments []Instrument `json:"result"`
	ID          string       `json:"id"`
}

type Instrument struct {
	InstrumentType        string        `json:"instrument_type"`
	InstrumentName        string        `json:"instrument_name"`
	ScheduledActivation   int           `json:"scheduled_activation"`
	ScheduledDeactivation int           `json:"scheduled_deactivation"`
	IsActive              bool          `json:"is_active"`
	TickSize              string        `json:"tick_size"`
	MinimumAmount         string        `json:"minimum_amount"`
	MaximumAmount         string        `json:"maximum_amount"`
	AmountStep            string        `json:"amount_step"`
	MarkPriceFeeRateCap   string        `json:"mark_price_fee_rate_cap"`
	MakerFeeRate          string        `json:"maker_fee_rate"`
	TakerFeeRate          string        `json:"taker_fee_rate"`
	BaseFee               string        `json:"base_fee"`
	BaseCurrency          string        `json:"base_currency"`
	QuoteCurrency         string        `json:"quote_currency"`
	Details               OptionDetails `json:"option_details"`
	PerpDetails           any           `json:"perp_details"`
	BaseAssetAddress      string        `json:"base_asset_address"`
	BaseAssetSubID        string        `json:"base_asset_sub_id"`
}
type OptionDetails struct {
	Index           string `json:"index"`
	Expiry          int    `json:"expiry"`
	Strike          string `json:"strike"`
	OptionType      string `json:"option_type"`
	SettlementPrice any    `json:"settlement_price"`
}

// https://v2-docs.lyra.finance/reference/post_public-get-ticker
type GetTickerResult struct {
	Result struct {
		InstrumentType        string        `json:"instrument_type"`
		InstrumentName        string        `json:"instrument_name"`
		ScheduledActivation   int           `json:"scheduled_activation"`
		ScheduledDeactivation int           `json:"scheduled_deactivation"`
		IsActive              bool          `json:"is_active"`
		TickSize              string        `json:"tick_size"`
		MinimumAmount         string        `json:"minimum_amount"`
		MaximumAmount         string        `json:"maximum_amount"`
		AmountStep            string        `json:"amount_step"`
		MarkPriceFeeRateCap   string        `json:"mark_price_fee_rate_cap"`
		MakerFeeRate          string        `json:"maker_fee_rate"`
		TakerFeeRate          string        `json:"taker_fee_rate"`
		BaseFee               string        `json:"base_fee"`
		BaseCurrency          string        `json:"base_currency"`
		QuoteCurrency         string        `json:"quote_currency"`
		Details               OptionDetails `json:"option_details"`
		PerpDetails           any           `json:"perp_details"`
		BaseAssetAddress      string        `json:"base_asset_address"`
		BaseAssetSubID        string        `json:"base_asset_sub_id"`
		BestAskAmount         string        `json:"best_ask_amount"`
		BestAskPrice          string        `json:"best_ask_price"`
		BestBidAmount         string        `json:"best_bid_amount"`
		BestBidPrice          string        `json:"best_bid_price"`
		OptionPricing         struct {
			Delta        string `json:"delta"`
			Theta        string `json:"theta"`
			Gamma        string `json:"gamma"`
			Vega         string `json:"vega"`
			Iv           string `json:"iv"`
			Rho          string `json:"rho"`
			MarkPrice    string `json:"mark_price"`
			ForwardPrice string `json:"forward_price"`
			BidIv        string `json:"bid_iv"`
			AskIv        string `json:"ask_iv"`
		} `json:"option_pricing"`
		IndexPrice string          `json:"index_price"`
		MarkPrice  string          `json:"mark_price"`
		Stats      InstrumentStats `json:"stats"`
		Timestamp  int64           `json:"timestamp"`
		MinPrice   string          `json:"min_price"`
		MaxPrice   string          `json:"max_price"`
	} `json:"result"`
	ID string `json:"id"`
}

type InstrumentStats struct {
	ContractVolume string `json:"contract_volume"`
	NumTrades      string `json:"num_trades"`
	OpenInterest   string `json:"open_interest"`
	High           string `json:"high"`
	Low            string `json:"low"`
	PercentChange  string `json:"percent_change"`
	UsdChange      string `json:"usd_change"`
}
