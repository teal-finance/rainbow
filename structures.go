// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

type Option struct {
	Name         string // ASSET-DATE-Strike-OptionsType
	Type         string // CALL / PUT
	Asset        string // ETH, BTC, SOL
	Expiry       string // Expiry date in format 2021-12-31
	Strike       float64
	ExchangeType string // CEX / DEX
	Chain        string // Ethereum, Solana, None for CEX
	Layer        string // L1 or L2, for Deribit we put "None"
	Provider     string // Opyn, Lyra, Thales, Deribit, Psyoptions
	Bid          []Order
	Ask          []Order
}

type Order struct {
	Price         float64
	Quantity      float64
	QuoteCurrency string // ETH, BTC, USDT...
}
