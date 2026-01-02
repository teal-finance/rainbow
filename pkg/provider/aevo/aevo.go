// Copyright 2023 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package aevo

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
	baseURL = "https://api.aevo.xyz/"
	name    = "Aevo"
)

type Provider struct {
	ar garcon.AdaptiveRate
}

func (Provider) Name() string {
	return name
}

// adaptiveMinSleepTime to rate limit the Aevo API.
// link docs and check it.
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

	instruments, err := p.query()
	if err != nil {
		return nil, err
	}

	options, err := p.fillOptions(instruments)
	if err != nil {
		return nil, log.Error("filloptions", err).Err()
	}
	return options, nil
}

func (p Provider) query() (Markets, error) {
	const api = baseURL + "markets?instrument_type=OPTION"
	url := api
	log.Info(name + " " + url)

	var result Markets
	err := p.ar.Get("", url, &result, maxBytesToRead)
	if err != nil {
		return nil, log.Error("query", url, err).Err()
	}

	return result, nil
}

func (p Provider) fillOptions(instruments Markets) ([]rainbow.Option, error) {
	url := ""
	var result Orderbook
	var err error
	options := []rainbow.Option{}
	s := 0.0
	iv := 0.0
	for _, i := range instruments {
		url = baseURL + "orderbook?instrument_name=" + i.InstrumentName
		err = p.ar.Get("", url, &result, maxBytesToRead)
		if err != nil {
			return nil, log.Error("ordebook", i.InstrumentName, err).Err()
		}
		s, err = convert(i.Strike)
		if err != nil {
			return nil, err
		}
		iv, err = convert(i.Greeks.Iv)
		if err != nil {
			return nil, err
		}
		bid, err := bidAsksToOrders(result.Bids)
		if err != nil {
			return nil, log.Error("Bid conversion", err).Err()
		}
		ask, err := bidAsksToOrders(result.Asks)
		if err != nil {
			return nil, log.Error("Bid conversion", err).Err()
		}
		e, err := strconv.ParseInt(i.Expiry, 10, 64)
		if err != nil {
			return nil, log.Error("convert date", i.Expiry, err).Err()
		}

		seconds := e / 1000000000
		ns := (e % 1000_000_000) * 1000_000
		expiryTime := time.Unix(seconds, ns).UTC()
		expiryStr := expiryTime.Format("2006-01-02 15:04:05")

		options = append(options, rainbow.Option{
			Name:            i.InstrumentName,
			Type:            strings.ToUpper(i.OptionType),
			Asset:           i.UnderlyingAsset,
			UnderlyingAsset: i.UnderlyingAsset,
			Strike:          s,
			Expiry:          expiryStr,
			ExchangeType:    "DEX",
			Chain:           "Ethereum",
			Layer:           "L2",
			LayerName:       "Aevo",
			Provider:        name,
			UnderlyingQuote: i.QuoteAsset,
			QuoteCurrency:   "USD",
			URL:             baseURL + "option/" + i.UnderlyingAsset + "/" + i.InstrumentName,
			MarketIV:        iv * 100,
			// Greeks: ,
			Bid: bid,
			Ask: ask,
			// OpenInterest: , TODO https://api-docs.aevo.xyz/reference/getinstrumentinstrumentname
			ProtocolID: i.InstrumentID,
		})
	}

	return options, nil
}

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

func convert(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, log.Error("conversion failure", s, err).Err()
	}
	return f, nil
}

// https://api-docs.aevo.xyz/reference/getmarkets
//	url := "https://api.aevo.xyz/markets?instrument_type=OPTION"

// https://api-docs.aevo.xyz/reference/getexpiries

// https://api-docs.aevo.xyz/reference/getorderbook

type (
	Markets []Market
	Market  struct {
		InstrumentID    string `json:"instrument_id"`
		InstrumentName  string `json:"instrument_name"`
		InstrumentType  string `json:"instrument_type"`
		UnderlyingAsset string `json:"underlying_asset"`
		QuoteAsset      string `json:"quote_asset"`
		PriceStep       string `json:"price_step"`
		AmountStep      string `json:"amount_step"`
		MinOrderValue   string `json:"min_order_value"`
		MarkPrice       string `json:"mark_price"`
		ForwardPrice    string `json:"forward_price"`
		IndexPrice      string `json:"index_price"`
		IsActive        bool   `json:"is_active"`
		OptionType      string `json:"option_type"`
		Expiry          string `json:"expiry"`
		Strike          string `json:"strike"`
		Greeks          greeks `json:"greeks"`
	}
)

type greeks struct {
	Delta string `json:"delta"`
	Theta string `json:"theta"`
	Gamma string `json:"gamma"`
	Rho   string `json:"rho"`
	Vega  string `json:"vega"`
	Iv    string `json:"iv"`
}

type Orderbook struct {
	Type           string     `json:"type"`
	InstrumentID   string     `json:"instrument_id"`
	InstrumentName string     `json:"instrument_name"`
	InstrumentType string     `json:"instrument_type"`
	Bids           [][]string `json:"bids"`
	Asks           [][]string `json:"asks"`
	LastUpdated    string     `json:"last_updated"`
	Checksum       string     `json:"checksum"`
}
