// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import "github.com/teal-finance/rainbow/pkg/rainbow"

type Arb struct {
	Type         string `json:"type"` // CALL / PUT
	Strike       float64
	Expiry       string `json:"expiry"`
	Asset        string `json:"asset"`
	BuyQ         float64
	BuyPx        float64
	BuyProvider  string
	SellQ        float64
	SellPx       float64
	SellProvider string
	SellM        float64
	ROI          float64 // percentage of benef
}

type Arbs []Arb

func buylowsellhigh(blocks Blocks, limit float64) Arbs {
	arbs := Arbs{}
	for _, block := range blocks {
		if len(block.Options) == 1 { // normally shouldn't be empty but weird error I can't fix
			continue
		}

		for i := 0; i < len(block.Options)-1; i++ {
			opt := block.Options[i]
			// this also test local arbs, if in one place the bid > ask
			for _, o := range block.Options[i:] {
				if conditions(o.Bid, opt.Ask) {
					r := roi(block.Strike, o.Bid, opt.Ask)
					if r >= limit {
						arbs = saveArbitrage(arbs, block, r, o, opt)
					}
				} else if conditions(opt.Bid, o.Ask) {
					r := roi(block.Strike, opt.Bid, o.Ask)
					if r >= limit {
						arbs = saveArbitrage(arbs, block, r, opt, o)
					}
				}
			}
		}
	}

	return arbs
}

func conditions(bid, ask []rainbow.Order) bool {
	return len(bid) != 0 && len(ask) != 0 &&
		ask[0].Price*ask[0].Size != 0 && bid[0].Price*bid[0].Size != 0 &&
		ask[0].Price < bid[0].Price
}

func roi(strike float64, bid, ask []rainbow.Order) float64 {
	return 100 * (bid[0].Price - ask[0].Price) / strike
}

func saveArbitrage(arbs Arbs, block Block, r float64, bidOptions, askOption rainbow.Option) Arbs {
	return append(arbs, Arb{
		Type:         block.Type,
		Strike:       block.Strike,
		Expiry:       block.Expiry,
		Asset:        block.Asset,
		BuyPx:        askOption.Ask[0].Price,
		BuyQ:         askOption.Ask[0].Size,
		BuyProvider:  askOption.Provider,
		SellPx:       bidOptions.Bid[0].Price,
		SellQ:        bidOptions.Bid[0].Size,
		SellProvider: bidOptions.Provider,
		ROI:          r,
	})
}
