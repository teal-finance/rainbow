package api

import (
	"log"
	"math"
	"time"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

type APIHandler struct {
	Service *rainbow.Service
}

type CallPut struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	Put      Limit   `json:"put"`
	Call     Limit   `json:"call"`
	Date     string  `json:"date"`
	Expiry   string  `json:"expiry"`
	Provider string  `json:"provider"`
	Asset    string  `json:"asset"`
	Strike   float64 `json:"strike"`
}

type Limit struct {
	Bid StrOrder `json:"bid"`
	Ask StrOrder `json:"ask"`
}

type StrOrder struct {
	// Price is often abbreviate "px" used by many Centralized Exchanges
	// such as in the FIX protocol: https://fiximate.fixtrading.org/legacy/en/FIX.5.0SP2/abbreviations.html
	Price string `json:"px"`
	// Size is a shorter synonym for "Quantity".
	Size string `json:"size"`
}

func buildCallPut(options []rainbow.Option) CallPut {
	rows := make([]Row, 0, len(options)/2)

	for asset, optionsSameAsset := range groupByAsset(options) {
		asset = sanitizeAsset(asset)

		for date, optionsSameExpiry := range groupByExpiry(optionsSameAsset) {
			expiry := sanitizeDate(date)

			for strike, optionsSameStrike := range groupByStrike(optionsSameExpiry) {
				for provider, optionsSameProvider := range groupByProvider(optionsSameStrike) {
					call := Limit{}
					put := Limit{}

					for _, o := range optionsSameProvider {
						if o.Type == "PUT" {
							put = newLimit(o)
						} else {
							call = newLimit(o)
						}
					}

					rows = append(rows, Row{
						Asset:    asset,
						Date:     date,
						Expiry:   expiry,
						Provider: provider,
						Call:     call,
						Strike:   math.Round(strike*100) / 100, // rounding to the nearest
						Put:      put,
					})
				}
			}
		}
	}

	return CallPut{Rows: rows}
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
		slice, ok := expiryToOptions[o.Expiry]
		if ok {
			expiryToOptions[o.Expiry] = append(slice, o)
		} else {
			expiryToOptions[o.Expiry] = []rainbow.Option{o}
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

func groupByProvider(options []rainbow.Option) (providerToOptions map[string][]rainbow.Option) {
	providerToOptions = map[string][]rainbow.Option{}

	for _, o := range options {
		slice, ok := providerToOptions[o.Provider]
		if ok {
			providerToOptions[o.Provider] = append(slice, o)
		} else {
			providerToOptions[o.Provider] = []rainbow.Option{o}
		}
	}

	return providerToOptions
}

func newLimit(o rainbow.Option) Limit {
	bPx, bSz, aPx, aSz := rainbow.BestLimitHTML(o)

	return Limit{
		Bid: StrOrder{Price: bPx, Size: bSz},
		Ask: StrOrder{Price: aPx, Size: aSz},
	}
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
		log.Printf("WARN prettyDate() cannot parse %q", date)

		return date
	}

	return t.Format("Jan _2")
}
