// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import "math"

type Expiries struct {
	AssetToExpiries map[string]map[string]int `json:"expiries"`
}

type Table struct {
	Asset  string `json:"asset"`
	Expiry string `json:"expiry"`
	Rows   []Row  `json:"table"`
}

type Row struct {
	Provider string `json:"p"`

	CallBid    float64 `json:"cb"`
	CallPrice  float64 `json:"cp"`
	CallAsk    float64 `json:"ca"`
	CallChange float64 `json:"cc"`
	CallVolume float64 `json:"cv"`
	CallOpen   float64 `json:"co"`

	Strike float64 `json:"s"`

	PutBid    float64 `json:"pb"`
	PutPrice  float64 `json:"pp"`
	PutAsk    float64 `json:"pa"`
	PutChange float64 `json:"pc"`
	PutVolume float64 `json:"pv"`
	PutOpen   float64 `json:"po"`
}

func buildExpiries(options []Option) Expiries {
	assetToExpiries := map[string]map[string]int{}

	for _, o := range options {
		if assetToExpiries[o.Asset] == nil {
			assetToExpiries[o.Asset] = map[string]int{o.Expiry: 1}
		} else {
			assetToExpiries[o.Asset][o.Expiry]++
		}
	}

	return Expiries{assetToExpiries}
}

func buildTables(options []Option) map[string]Table {
	tables := map[string]Table{}

	for asset, expiryProviderStrikeToOptions := range tidy(options) {
		for expiry, providerStrikeToOptions := range expiryProviderStrikeToOptions {
			t := Table{
				Asset:  asset,
				Expiry: expiry,
				Rows:   buildRows(providerStrikeToOptions),
			}

			tables[asset+expiry] = t

			if _, ok := tables[""]; !ok {
				tables[""] = t
			} else if len(tables[""].Rows) < len(t.Rows) {
				tables[""] = t
			}
		}
	}

	return tables
}

func tidy(options []Option) map[string]map[string]map[string]map[float64][]Option {
	assetExpiryProviderStrikeToOptions := map[string]map[string]map[string]map[float64][]Option{}

	for _, o := range options {
		if _, ok := assetExpiryProviderStrikeToOptions[o.Asset]; !ok {
			assetExpiryProviderStrikeToOptions[o.Asset] = map[string]map[string]map[float64][]Option{}
		}

		if _, ok := assetExpiryProviderStrikeToOptions[o.Asset][o.Expiry]; !ok {
			assetExpiryProviderStrikeToOptions[o.Asset][o.Expiry] = map[string]map[float64][]Option{}
		}

		if _, ok := assetExpiryProviderStrikeToOptions[o.Asset][o.Expiry][o.Provider]; !ok {
			assetExpiryProviderStrikeToOptions[o.Asset][o.Expiry][o.Provider] = map[float64][]Option{}
		}

		if _, ok := assetExpiryProviderStrikeToOptions[o.Asset][o.Expiry][o.Provider][o.Strike]; !ok {
			assetExpiryProviderStrikeToOptions[o.Asset][o.Expiry][o.Provider][o.Strike] = []Option{}
		}

		assetExpiryProviderStrikeToOptions[o.Asset][o.Expiry][o.Provider][o.Strike] = append(
			assetExpiryProviderStrikeToOptions[o.Asset][o.Expiry][o.Provider][o.Strike], o)
	}

	return assetExpiryProviderStrikeToOptions
}

func buildRows(providerStrikeToOptions map[string]map[float64][]Option) []Row {
	rows := []Row{}

	for provider, strikeToOptions := range providerStrikeToOptions {
		for strike, options := range strikeToOptions {
			r := Row{
				Provider: provider,
				Strike:   strike,
			}

			for _, o := range options {
				vol := 0.0

				bid := math.NaN()
				if len(o.Bid) > 0 {
					bid = o.Bid[0].Price
					vol += o.Bid[0].Quantity
				}

				ask := math.NaN()
				if len(o.Ask) > 0 {
					ask = o.Ask[0].Price
					vol += o.Ask[0].Quantity
				}

				if o.Type == "PUT" {
					r.PutBid = bid
					r.PutPrice = 0
					r.PutAsk = ask
					r.PutChange = 0
					r.PutVolume = vol
					r.PutOpen = 0
				} else {
					r.CallBid = bid
					r.CallPrice = 0
					r.CallAsk = ask
					r.CallChange = 0
					r.CallVolume = vol
					r.CallOpen = 0
				}
			}

			rows = append(rows, r)
		}
	}

	return rows
}
