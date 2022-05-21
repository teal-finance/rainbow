// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

type Option struct {
	Name          string    `json:"name"`     // ASSET-DATE-Strike-OptionsType
	Type          string    `json:"type"`     // CALL / PUT
	Asset         string    `json:"asset"`    // ETH, BTC, SOL, ...
	Expiry        string    `json:"expiry"`   // Expiry date in format 2021-12-31
	ExchangeType  string    `json:"exchange"` // CEX / DEX
	Chain         string    `json:"chain"`    // Ethereum, Solana and "–" for CEX (Deribit, Delta)
	Layer         string    `json:"layer"`    // L1, L2 and "–" for CEX (Deribit, Delta)
	Provider      string    `json:"provider"` // Opyn, Lyra, Thales, Deribit, Psyoptions, Zeta, Delta
	QuoteCurrency string    `json:"currency"` // ETH, BTC, USDT...
	Bid           []Order   `json:"bid"`      // Bid = buying offers
	BidIV         float64   `json:"bidIV"`    // Implied Volatility computed from the bid
	Ask           []Order   `json:"ask"`      // Ask = selling offers
	AskIV         float64   `json:"askIV"`    // Implied Volatility computed from the ask
	Greeks        TheGreeks `json:"greeks"`   // Greeks measure the sensitivity of an option’s premium to changes in the underlying variables.
	Strike        float64   `json:"strike"`   //
}

type Order struct {
	Price float64 `json:"px"`
	Size  float64 `json:"size"`
}

type TheGreeks struct {
	Vega  float64 `json:"vega"`
	Theta float64 `json:"theta"`
	Rho   float64 `json:"rho"`
	Gamma float64 `json:"gamma"`
	Delta float64 `json:"delta"`
}
