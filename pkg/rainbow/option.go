// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

type Option struct {
	Name          string  `json:"name"`     // ASSET-DATE-Strike-OptionsType
	Type          string  `json:"type"`     // CALL / PUT
	Asset         string  `json:"asset"`    // ETH, BTC, SOL
	Expiry        string  `json:"expiry"`   // Expiry date in format 2021-12-31
	ExchangeType  string  `json:"exchange"` // CEX / DEX
	Chain         string  `json:"chain"`    // Ethereum, Solana and "–" for CEX (Deribit)
	Layer         string  `json:"layer"`    // L1, L2 and "–" for CEX (Deribit)
	Provider      string  `json:"provider"` // Opyn, Lyra, Thales, Deribit, Psyoptions
	QuoteCurrency string  `json:"currency"` // ETH, BTC, USDT...
	Bid           []Order `json:"bid"`
	Ask           []Order `json:"ask"`
	Strike        float64 `json:"strike"`
}

type Order struct {
	Price     float64 `json:"px"`
	Size      float64 `json:"size"`
	Timestamp string  `json:"timestamp"`
}
