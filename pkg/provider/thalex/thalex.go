// Copyright 2023 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package thalex

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/teal-finance/emo"
	"github.com/teal-finance/garcon"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone(name)

const baseURL = "https://www."
const name = "Thalex"

type Provider struct {
	ar garcon.AdaptiveRate
}

func (Provider) Name() string {
	return name
}

// adaptiveMinSleepTime to rate limit the Thalex API.
// https://thalex.com/docs/#section/API-description/Message-rates
// The maximum number of matching engine messages (buy, sell, amend, delete, etc.) per connection per second is 10.
// When the connection is set to non-persistent (private/set_cancel_on_disconnect), this limit is raised to 50.
const adaptiveMinSleepTime = 25 * time.Millisecond

// Hour at which the options expires = 8:00 UTC.
const Hour = 8

// maxBytesToRead prevents wasting memory/CPU when receiving an abnormally huge response from Thalex API.
// we put the same param as Deribit
const maxBytesToRead = 2_000_000

func (p *Provider) Options() ([]rainbow.Option, error) {
	if p.ar.Name == "" {
		p.ar = garcon.NewAdaptiveRate(name, adaptiveMinSleepTime)
	}

	instruments, err := p.query()
	if err != nil {
		return nil, err
	}
	j := 9
	spew.Dump(instruments[j : j+1])
	i := instruments[j]
	var result tickers
	apiurl := "https://thalex.com/api/v2/public/ticker" + "?instrument_name=" + i.InstrumentName
	if err := p.ar.Get(i.InstrumentName, apiurl, &result); err != nil {
		log.Warn(name + " book " + err.Error())
	}

	spew.Dump(result)

	return []rainbow.Option{}, nil
}

func (p *Provider) query() ([]Instruments, error) {
	const api = "https://thalex.com/api/v2/public/instruments"
	url := api
	log.Info(name + url)

	var result instrumentsResult
	err := p.ar.Get("", url, &result, maxBytesToRead)
	if err != nil {
		return nil, err
	}

	return result.Result, nil
}

type tickers struct {
	Result Ticker `json:"result"`
}

type Ticker struct {
	MarkPrice     float64 `json:"mark_price"`
	BestBidPrice  int     `json:"best_bid_price"`
	BestBidAmount float64 `json:"best_bid_amount"`
	BestAskPrice  int     `json:"best_ask_price"`
	BestAskAmount float64 `json:"best_ask_amount"`
	LastPrice     int     `json:"last_price"`
	Iv            float64 `json:"iv"`
	Delta         float64 `json:"delta"`
	Volume24H     float64 `json:"volume_24h"`
	Value24H      float64 `json:"value_24h"`
	LowPrice24H   float64 `json:"low_price_24h"`
	HighPrice24H  float64 `json:"high_price_24h"`
	Change24H     float64 `json:"change_24h"`
	Index         float64 `json:"index"`
	Forward       float64 `json:"forward"`
	Funding_mark  float64 `json:"funding_mark"`
	Funding_rate  float64 `json:"funding_rate"`
	CollarLow     float64 `json:"collar_low"`
	CollarHigh    float64 `json:"collar_high"`
	OpenInterest  float64 `json:"open_interest"`
}

type instrumentsResult struct {
	Result []Instruments `json:"result"`
}

type Instruments struct {
	InstrumentName      string  `json:"instrument_name"`
	Product             string  `json:"product"`
	TickSize            float64 `json:"tick_size"`
	VolumeTickSize      float64 `json:"volume_tick_size"`
	Underlying          string  `json:"underlying"`
	Type                string  `json:"type"`
	OptionType          string  `json:"option_type,omitempty"`
	ExpiryDate          string  `json:"expiry_date"`
	ExpirationTimestamp int     `json:"expiration_timestamp"`
	StrikePrice         float64 `json:"strike_price,omitempty"`
	BaseCurrency        string  `json:"base_currency"`
	CreateTime          float64 `json:"create_time"`
	Legs                []struct {
		InstrumentName string `json:"instrument_name"`
		Quantity       int    `json:"quantity"`
	} `json:"legs,omitempty"`
}
