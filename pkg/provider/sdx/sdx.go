// Copyright 2023 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package sdx

import (
	"strconv"
	"strings"
	"time"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/garcon"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	Hour    = 8
	name    = "SDX"
	baseURL = "https://bidasklive-d2jplirhdq-as.a.run.app/?"
	asset   = "MSOL"
	dex     = "https://sdx.markets/trade/"
	network = "mainnet-beta"
)

var log = emo.NewZone(name)

type Provider struct {
	ar garcon.AdaptiveRate
}

func (Provider) Name() string {
	return name
}

/*
--Error handling
400 Bad Request: Invalid assets or network provided
429 Too Many Requests: Rate limit exceeded.
500 Internal Server Error: Unexpected server-side errors.
--Caching
The API internally caches responses for 60 seconds.
--Rate limiting
Be aware of the rate limiting policy of 5 requests / min to avoid service interruptions.
The API enforces rate limits to ensure equitable resource usage.
*/

// adaptiveMinSleepTime to rate limit the SDX API.
// https://docs.sdx.markets/developers/api-endpoints
const adaptiveMinSleepTime = 1 * time.Second

// maxBytesToRead prevents wasting memory/CPU when receiving an abnormally huge response from SDX API.
// taking Deribit value
const maxBytesToRead = 2_000_000

func (p *Provider) Options() ([]rainbow.Option, error) {
	if p.ar.Name == "" {
		p.ar = garcon.NewAdaptiveRate(name, adaptiveMinSleepTime)
	}
	p.ar.LogStats()

	assets, err := p.query(asset)
	if err != nil {
		return nil, err
	}
	options, _ := process(assets)
	return options, nil
}

func (p *Provider) query(coin string) ([]UnderlyingAsset, error) {
	url := baseURL + "assets=" + coin + "&network=" + network
	log.Info(name + " " + url)

	var results APIResults
	err := p.ar.Get(coin, url, &results, maxBytesToRead)
	if err != nil {
		return nil, err
	}

	return results.UnderlyingAssets, nil
}

func process(under []UnderlyingAsset) ([]rainbow.Option, error) {
	// TODO Loop when more markets :)
	markets := under[0].Markets
	var options []rainbow.Option

	for _, m := range markets {
		optionType := m.OptionType
		// fmt.Println(optionType)
		if len(optionType) >= 5 {
			continue
		}
		coin, expiryTime, expiry, strike, err := m.Infos()
		if err != nil {
			return []rainbow.Option{}, log.Error("Processing error", m.MarketName, err).Err()
		}
		// fmt.Println(coin, expiry, typ, strike)
		option := rainbow.Option{
			Name:            m.MarketName,
			Type:            optionType,
			Asset:           coin,
			UnderlyingAsset: coin,
			Strike:          strike,
			Expiry:          expiry,
			ExpiryTime:      int(expiryTime),
			ExchangeType:    "DEX",
			Chain:           "Solana",
			Layer:           "L1",
			LayerName:       "Solana",
			Provider:        name,
			QuoteCurrency:   "USD",
			UnderlyingQuote: "USDC",
			URL:             dex + asset,
			Bid: []rainbow.Order{{
				Price: m.BidStable,
				Size:  1,
			}},
			Ask: []rainbow.Order{{
				Price: m.AskStable,
				Size:  1,
			}},
			ProtocolID: strconv.Itoa(m.SeriesID),
		}

		if len(m.MarkPriceIV) != 0 {
			option.MarketIV = m.MarkPriceIV[0] * 100
		}
		options = append(options, option)
	}
	return options, nil
}

type APIResults struct {
	UnderlyingAssets []UnderlyingAsset `json:"underlyingAssets"`
}

type UnderlyingAsset struct {
	AssetName    string   `json:"assetName"`
	VaultAddress string   `json:"vaultAddress"`
	RetrievedAt  int      `json:"retrievedAt"`
	Markets      []Market `json:"markets"`
}
type Market struct {
	MarketName          string    `json:"marketName"`
	SeriesID            int       `json:"seriesId"`
	OptionType          string    `json:"optionType"`
	UnderlyingPrice     float64   `json:"underlyingPrice"`
	MarkPriceStable     float64   `json:"markPriceStable"`
	MarkPriceUnderlying float64   `json:"markPriceUnderlying"`
	BidStable           float64   `json:"bidStable"`
	BidUnderlying       float64   `json:"bidUnderlying"`
	AskStable           float64   `json:"askStable"`
	AskUnderlying       float64   `json:"askUnderlying"`
	BidIV               []float64 `json:"bidIV"`
	AskIV               []float64 `json:"askIV"`
	MarkPriceIV         []float64 `json:"markPriceIV"`
}

// Infos returns coin, expiry(Unix),expiry (formatted string), strike, err
// expiry: 0800 UTC each Friday
func (m *Market) Infos() (string, int64, string, float64, error) {
	/* example of names from API
	"MSOL-08Dec23-$64-Put"
	"MSOL-08Dec23-$70/$100-Call Spread"
	"MSOL-01Dec23-$76-Call"
	*/
	infos := strings.Split(m.MarketName, "-")

	expiry, err := time.Parse("02Jan06", infos[1])
	if err != nil {
		return "", 0, "", 0, log.Error("Expiry Parsing", m.MarketName, err).Err()
	}
	expiryStr := expiry.Add(Hour * time.Hour).Format("2006-01-02 15:04:05")

	strike, err := strconv.ParseFloat(strings.Trim(infos[2], "$"), 64)
	if err != nil {
		return "", 0, "", 0, log.Error("Strike Parsing", m.MarketName, err).Err()
	}

	return infos[0], expiry.Unix(), expiryStr, strike, nil
}
