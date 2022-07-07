// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package api

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

func (h Handler) CallPut(w http.ResponseWriter, r *http.Request) {
	options, err := h.Service.Options(rainbow.StoreArgs{})
	if err != nil {
		log.Print("ERROR Options ", err)
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}

	cp := h.align.buildCallPut(options)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(cp); err != nil {
		log.Print("ERROR CallPut ", err)
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
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

func (row *Row) less(other *Row) bool {
	if row.Strike < other.Strike {
		return true
	}
	if row.Strike > other.Strike {
		return false
	}
	if row.Asset < other.Asset {
		return true
	}
	if row.Asset > other.Asset {
		return false
	}
	if row.Expiry < other.Expiry {
		return true
	}
	if row.Expiry > other.Expiry {
		return false
	}
	return row.Provider < other.Provider
}

func IsRowEmpty(row []rainbow.Option) bool {
	//fmt.Println(len(row))
	for _, o := range row {
		if !IsOffersEmpty(o.Ask) || !IsOffersEmpty(o.Bid) {
			return false
		}
	}
	//fmt.Println(row[0].Name, " ", row[1].Name)
	return true
}

func IsOffersEmpty(order []rainbow.Order) bool {

	return len(order) == 0 || order[0].Size == 0 || order[0].Price < 0.001
}

func (a *Align) buildCallPut(options []rainbow.Option) []Row {
	rows := make([]Row, 0, len(options)/2)

	for asset, optionsSameAsset := range groupByAsset(options) {
		asset = sanitizeAsset(asset)

		for date, optionsSameExpiry := range groupByExpiry(optionsSameAsset) {
			expiry := sanitizeDate(date)

			for strike, optionsSameStrike := range groupByStrike(optionsSameExpiry) {
				for provider, optionsSameProvider := range groupByProvider(optionsSameStrike) {
					call := Limit{}
					put := Limit{}
					if IsRowEmpty(optionsSameProvider) {
						continue
					}
					for i := range optionsSameProvider {
						o := &optionsSameProvider[i]
						if o.Type == "PUT" {
							put = a.newLimit(o)
						} else {
							call = a.newLimit(o)
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

	return rows
}

func groupByAsset(options []rainbow.Option) (assetToOptions map[string][]rainbow.Option) {
	assetToOptions = map[string][]rainbow.Option{}

	for i := range options {
		// Sanitize to properly group assets
		asset := sanitizeAsset(options[i].Asset)

		slice, ok := assetToOptions[asset]
		if ok {
			assetToOptions[asset] = append(slice, options[i])
		} else {
			assetToOptions[asset] = []rainbow.Option{options[i]}
		}
	}

	return assetToOptions
}

func groupByExpiry(options []rainbow.Option) (expiryToOptions map[string][]rainbow.Option) {
	expiryToOptions = map[string][]rainbow.Option{}

	for i := range options {
		slice, ok := expiryToOptions[options[i].Expiry]
		if ok {
			expiryToOptions[options[i].Expiry] = append(slice, options[i])
		} else {
			expiryToOptions[options[i].Expiry] = []rainbow.Option{options[i]}
		}
	}

	return expiryToOptions
}

func groupByStrike(options []rainbow.Option) (strikeToOptions map[float64][]rainbow.Option) {
	strikeToOptions = map[float64][]rainbow.Option{}
	for i := range options {
		slice, ok := strikeToOptions[options[i].Strike]
		if ok {
			strikeToOptions[options[i].Strike] = append(slice, options[i])
		} else {
			strikeToOptions[options[i].Strike] = []rainbow.Option{options[i]}
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

func (a *Align) newLimit(o *rainbow.Option) Limit {
	bPx, bSz, aPx, aSz := a.BestLimitHTML(o)
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
		log.Printf("WRN prettyDate() cannot parse %q", date)
		return date
	}
	return t.Format("Jan _2")
}
