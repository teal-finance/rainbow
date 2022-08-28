// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"time"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

type (
	Blocks []Block
	Block  struct {
		Expiry  string  `json:"expiry"`
		Asset   string  `json:"asset"`
		Strike  float64 `json:"strike"`
		Type    string  `json:"type"`
		Options []rainbow.Option
	}
)

func buildCallPut(options []rainbow.Option) Blocks {
	var blocks Blocks
	for asset, optionsSameAsset := range groupByAsset(options) {
		asset = sanitizeAsset(asset)

		for date, optionsSameExpiry := range groupByExpiry(optionsSameAsset) {
			expiry := date // sanitizeDate(date)

			for strike, optionsSameStrike := range groupByStrike(optionsSameExpiry) {
				for t, optionsSameType := range groupByType(optionsSameStrike) {
					blocks = append(blocks, Block{
						Expiry:  expiry,
						Asset:   asset,
						Strike:  strike,
						Type:    t,
						Options: optionsSameType,
					})
				}
			}
		}
	}

	return blocks
}

func groupByAsset(options []rainbow.Option) (assetToOptions map[string][]rainbow.Option) {
	assetToOptions = map[string][]rainbow.Option{}

	for _, o := range options {
		// Sanitize to properly group assets
		asset := sanitizeAsset(o.Asset)

		slice, ok := assetToOptions[asset]
		if ok {
			assetToOptions[asset] = append(slice, o)
		} else {
			assetToOptions[asset] = []rainbow.Option{o}
		}
	}

	return assetToOptions
}

func groupByExpiry(options []rainbow.Option) (expiryToOptions map[string][]rainbow.Option) {
	expiryToOptions = map[string][]rainbow.Option{}

	for _, o := range options {
		expiry := sanitizeDate(o.Expiry)
		slice, ok := expiryToOptions[expiry]
		if ok {
			expiryToOptions[expiry] = append(slice, o)
		} else {
			expiryToOptions[expiry] = []rainbow.Option{o}
		}
	}

	return expiryToOptions
}

func groupByStrike(options []rainbow.Option) (strikeToOptions map[float64][]rainbow.Option) {
	strikeToOptions = map[float64][]rainbow.Option{}

	for _, o := range options {
		slice, ok := strikeToOptions[o.Strike]
		if ok {
			strikeToOptions[o.Strike] = append(slice, o)
		} else {
			strikeToOptions[o.Strike] = []rainbow.Option{o}
		}
	}

	return strikeToOptions
}

func groupByType(options []rainbow.Option) (typeToOptions map[string][]rainbow.Option) {
	typeToOptions = map[string][]rainbow.Option{}

	for _, o := range options {
		slice, ok := typeToOptions[o.Type]
		if ok {
			typeToOptions[o.Type] = append(slice, o)
		} else {
			typeToOptions[o.Type] = []rainbow.Option{o}
		}
	}

	return typeToOptions
}

// sanitizeAsset removes
// "W" (or "w") in "WETH" and "WBTC"
// "s" for Lyra assets from Synthethix: sETH, sBTC, sLINK, sSOL.
func sanitizeAsset(asset string) string {
	if len(asset) >= 4 && (asset[0] == 'W' || asset[0] == 'w' || asset[0] == 's') {
		return asset[1:]
	}

	return asset
}

// prettyDate converts "2021-12-31 23:59:59" into "Dec 31".
func sanitizeDate(date string) string {
	t, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		log.Printf("WRN sanitizeDate cannot parse %q", date)
		return date // TODO: notify Mattermost
	}

	return t.Format("Jan _2")
}
