// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package rainbow

import "fmt"

type Option struct {
	Name          string  `json:"name"`      // ASSET-DATE-Strike-OptionsType
	Type          string  `json:"type"`      // CALL / PUT
	Asset         string  `json:"asset"`     // ETH, BTC, SOL
	Expiry        string  `json:"expiry"`    // Expiry date in format 2021-12-31
	ExchangeType  string  `json:"exchange"`  // CEX / DEX
	Chain         string  `json:"chain"`     // Ethereum, Solana and "–" for CEX (Deribit)
	Layer         string  `json:"layer"`     // L1, L2 and "–" for CEX (Deribit)
	LayerName     string  `json:"layername"` // name of the layer 1/2. "-" for CEX
	Provider      string  `json:"provider"`  // Opyn, Lyra, Thales, Deribit, Psyoptions
	QuoteCurrency string  `json:"currency"`  // ETH, BTC, USDT...
	Bid           []Order `json:"bid"`
	Ask           []Order `json:"ask"`
	Strike        float64 `json:"strike"`
}

// TODO put standard name here
func (o Option) OptionName() string {
	if o.Name != "" {
		return o.Name
	}
	return o.Provider + "-" + o.LayerName + "-" + o.Asset + "-" + fmt.Sprintf("%.3f", o.Strike) + "-" + o.Type + "-" + o.Expiry

}

type Order struct {
	Price float64 `json:"px"`
	Size  float64 `json:"size"`
}
