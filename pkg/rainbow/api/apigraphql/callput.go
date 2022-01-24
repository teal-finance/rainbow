package apigraphql

import (
	"fmt"
	"log"
	"time"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

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
	Price string `json:"px"`
	Size  string `json:"size"`
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
						Strike:   strike,
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
		slice, ok := assetToOptions[o.Asset]
		if ok {
			assetToOptions[o.Asset] = append(slice, o)
		} else {
			assetToOptions[o.Asset] = []rainbow.Option{o}
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
	l := Limit{}
	if len(o.Bid) > 0 && o.Bid[0].Size != 0 {
		l.Bid.Price = fmt.Sprintf("%.2f", o.Bid[0].Price)
		l.Bid.Size = fmt.Sprintf("%.2f", o.Bid[0].Size)
	}
	if len(o.Ask) > 0 && o.Ask[0].Size != 0 {
		l.Ask.Price = fmt.Sprintf("%.2f", o.Ask[0].Price)
		l.Ask.Size = fmt.Sprintf("%.2f", o.Ask[0].Size)
	}
	return l
}

// sanitizeAsset removes "W" (or "w") in "WETC" and "WBTC".
func sanitizeAsset(asset string) string {
	if len(asset) >= 4 && (asset[0] == 'W' || asset[0] == 'w') {
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
	return t.Format("Jan 02")
}
