// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

type CPFormat struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	Asset    string `json:"asset"`
	Expiry   string `json:"expiry"`
	Provider string `json:"provider"`

	Call OptionIndicators `json:"call"`

	Strike float64 `json:"strike"`

	Put OptionIndicators `json:"put"`
}

type OptionIndicators struct {
	Bid              SimpleOrder `json:"bid"`
	Ask              SimpleOrder `json:"ask"`
	OthersIndicators float64     `json:"other_indicators_will_be_added_soon"`
}

type SimpleOrder struct {
	Price float64 `json:"px"`
	Size  float64 `json:"size"`
}

func buildCPFormat(options []Option) CPFormat {
	rows := make([]Row, 0, len(options)/2)

	for asset, optionsWithSameAsset := range tidyByAsset(options) {
		for expiry, optionsWithSameExpiry := range tidyByExpiry(optionsWithSameAsset) {
			for strike, optionsWithSameStrike := range tidyByStrike(optionsWithSameExpiry) {
				for provider, optionsWithSameProvider := range tidyByProvider(optionsWithSameStrike) {
					r := Row{
						Asset:    asset,
						Expiry:   expiry,
						Provider: provider,
						Strike:   strike,
					}

					for _, o := range optionsWithSameProvider {
						if o.Type == "PUT" {
							r.Put = NewOptionIndicators(o)
						} else {
							r.Call = NewOptionIndicators(o)
						}
					}

					rows = append(rows, r)
				}
			}
		}
	}

	return CPFormat{Rows: rows}
}

func tidyByAsset(options []Option) (assetToOptions map[string][]Option) {
	assetToOptions = map[string][]Option{}

	for _, o := range options {
		slice, ok := assetToOptions[o.Asset]
		if ok {
			assetToOptions[o.Asset] = append(slice, o)
		} else {
			assetToOptions[o.Asset] = []Option{o}
		}
	}

	return assetToOptions
}

func tidyByExpiry(options []Option) (expiryToOptions map[string][]Option) {
	expiryToOptions = map[string][]Option{}

	for _, o := range options {
		slice, ok := expiryToOptions[o.Expiry]
		if ok {
			expiryToOptions[o.Expiry] = append(slice, o)
		} else {
			expiryToOptions[o.Expiry] = []Option{o}
		}
	}

	return expiryToOptions
}

func tidyByStrike(options []Option) (strikeToOptions map[float64][]Option) {
	strikeToOptions = map[float64][]Option{}

	for _, o := range options {
		slice, ok := strikeToOptions[o.Strike]
		if ok {
			strikeToOptions[o.Strike] = append(slice, o)
		} else {
			strikeToOptions[o.Strike] = []Option{o}
		}
	}

	return strikeToOptions
}

func tidyByProvider(options []Option) (providerToOptions map[string][]Option) {
	providerToOptions = map[string][]Option{}

	for _, o := range options {
		slice, ok := providerToOptions[o.Provider]
		if ok {
			providerToOptions[o.Provider] = append(slice, o)
		} else {
			providerToOptions[o.Provider] = []Option{o}
		}
	}

	return providerToOptions
}

func NewOptionIndicators(o Option) OptionIndicators {
	oi := OptionIndicators{
		Bid: SimpleOrder{
			Price: 0,
			Size:  0,
		},
		Ask: SimpleOrder{
			Price: 0,
			Size:  0,
		},
		OthersIndicators: 0,
	}

	if len(o.Bid) > 0 {
		oi.Bid.Price = o.Bid[0].Price
		oi.Bid.Size = o.Bid[0].Quantity
	}

	if len(o.Ask) > 0 {
		oi.Ask.Price = o.Ask[0].Price
		oi.Ask.Size = o.Ask[0].Quantity
	}

	return oi
}
