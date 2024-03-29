// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package rainbow

import (
	"fmt"
)

// TODO
// ADD a "quote" type for data that change often
// ADD a timestamp to this "quote"
// Do a proper Name
// ADD a hash on the name
// ADD dump of db on the cli
// Work on the csv

type Option struct {
	Name            string  `json:"name"`               // ASSET-DATE-Strike-OptionsType
	Type            string  `json:"type"`               // CALL / PUT  // TODO add exotic like binary and perp(squeeth)
	Asset           string  `json:"asset"`              // ETH, BTC, SOL, ... the crypto we are exposed to
	UnderlyingAsset string  `json:"underlyingasset"`    // sETH, WETH, sBTC, WBTC ... the actual asset token we track
	Strike          float64 `json:"strike"`             // Option strike
	Expiry          string  `json:"expiry"`             // Expiry date in Format("2006-01-02 15:04:05")
	ExpiryTime      int     `json:"expirytime"`         // Expiry in timestamp format (1136241533) seconds since the UNIX epoch
	ExchangeType    string  `json:"exchange"`           // CEX / DEX
	Chain           string  `json:"chain"`              // Ethereum, Solana and "–" for CEX (Deribit)
	Layer           string  `json:"layer"`              // L1, L2 and "–" for CEX (Deribit)
	LayerName       string  `json:"layername"`          // name of the layer 1/2. "-" for CEX
	Provider        string  `json:"provider"`           // Lyra, Thales, Deribit, Synquote, Zeta, Delta
	QuoteCurrency   string  `json:"currency"`           // USD
	UnderlyingQuote string  `json:"underlyingcurrency"` // sUSD, USDC, USDT...   the actual quote token we track
	URL             string  `json:"url"`                // link to market @ provider.
	Bid             []Order `json:"bid"`                // Bid = buying offers
	BidIV           float64 `json:"bidIV"`              // Implied Volatility computed from the bid
	Ask             []Order `json:"ask"`                // Ask = selling offers
	AskIV           float64 `json:"askIV"`              // Implied Volatility computed from the ask
	MarketIV        float64 `json:"markIV"`             // When it is present on the provider, we store their Market IV
	// see https://corporatefinanceinstitute.com/resources/knowledge/trading-investing/option-greeks/
	Greeks       TheGreeks `json:"greeks"`       // Greeks measure the sensitivity of an option’s price to its the underlying determining parameters.
	OpenInterest float64   `json:"openinterest"` // TODO ADD volume too
	ProtocolID   string    `json:"protocolID"`   // when present log the ID of that instrument on the provider
}

// TODO put standard name here.
func (o *Option) OptionName() string {
	if o.Name != "" {
		return o.Name
	}
	return o.Provider + "-" + o.LayerName + "-" + o.Asset + "-" + fmt.Sprintf("%.2f", o.Strike) + "-" + o.Type + "-" + o.Expiry
}

type Order struct {
	Price float64 `json:"px"`
	Size  float64 `json:"size"`
	// TODO add the IV if provided for the order here
}

type TheGreeks struct {
	Vega  float64 `json:"vega"`
	Theta float64 `json:"theta"`
	Rho   float64 `json:"rho"`
	Gamma float64 `json:"gamma"`
	Delta float64 `json:"delta"`
}

func ListProviders(ro []Option) map[string]int {
	providerToOptions := make(map[string]int)
	for i := range ro {
		_, ok := providerToOptions[ro[i].Provider]
		if ok {
			providerToOptions[ro[i].Provider]++
		} else {
			providerToOptions[ro[i].Provider] = 1
		}
	}
	return providerToOptions
}

/*
for easy copy-pasta

 rainbow.Option {
	Name:
	Type:
	Asset:
	UnderlyingAsset:
	Strike:
	Expiry:
	ExpiryTime:
	ExchangeType:
	Chain:
	Layer:
	LayerName:
	Provider:
	QuoteCurrency:
	UnderlyingQuote:
	URL:
	Bid:
	BidIV:
	Ask:
	AskIV:
	MarketIV:
	Greeks:
	OpenInterest:
	ProtocolID:
}
*/
