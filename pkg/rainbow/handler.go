// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	c *Cache
}

func (h *handler) router() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/options", h.getOptions)
		r.Get("/options/cp", h.getCallPut)
	})

	return r
}

func (h handler) getOptions(w http.ResponseWriter, r *http.Request) {
	options, err := h.c.Options()
	if err != nil {
		log.Print("ERROR getOptions ", err)
		http.Error(w, "No Content", http.StatusNoContent)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(options); err != nil {
		log.Print("ERROR getOptions ", err)
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)

		return
	}
}

func (h handler) getCallPut(w http.ResponseWriter, r *http.Request) {
	cp, err := h.c.CallPut()
	if err != nil {
		log.Print("ERROR getCallPut ", err)
		http.Error(w, "No Content", http.StatusNoContent)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(cp); err != nil {
		log.Print("ERROR getCallPut ", err)
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)

		return
	}
}

type CallPut struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	Call     Limit   `json:"call"`
	Put      Limit   `json:"put"`
	Asset    string  `json:"asset"`
	Expiry   string  `json:"expiry"`
	Provider string  `json:"provider"`
	Strike   float64 `json:"strike"`
}

type Limit struct {
	Bid StrOrder `json:"bid"`
	Ask StrOrder `json:"ask"`
}

func NoneLimit() Limit {
	return Limit{
		Bid: StrOrder{Px: dashRightAlignHTML, Size: dashLeftAlignHTML},
		Ask: StrOrder{Px: dashRightAlignHTML, Size: dashLeftAlignHTML},
	}
}

type StrOrder struct {
	// Px is a known abbreviation for "Price" used by many Centralized Exchanges
	// such as in the FIX protocol: https://fiximate.fixtrading.org/legacy/en/FIX.5.0SP2/abbreviations.html
	// Independently of the FIX protocol, "Px" is intensively used at Euronext.
	// "Px" is also a French abbreviation for "Prix". :-)
	// Rainbow uses "px" in the API (JSON)) but uses "price" on the front-end side.
	Px string `json:"px"`

	// Size is often used in lieu of the longer word "Quantity".
	Size string `json:"size"`
}

func buildCallPut(options []Option) CallPut {
	rows := make([]Row, 0, len(options)/2)

	for asset, optionsSameAsset := range groupByAsset(options) {
		for expiry, optionsSameExpiry := range groupByExpiry(optionsSameAsset) {
			for strike, optionsSameStrike := range groupByStrike(optionsSameExpiry) {
				for provider, optionsSameProvider := range groupByProvider(optionsSameStrike) {
					call := NoneLimit()
					put := NoneLimit()

					for _, o := range optionsSameProvider {
						if o.Type == "PUT" {
							put = newOptionIndicators(o)
						} else {
							call = newOptionIndicators(o)
						}
					}

					rows = append(rows, Row{
						Asset:    asset,
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

func groupByAsset(options []Option) (assetToOptions map[string][]Option) {
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

func groupByExpiry(options []Option) (expiryToOptions map[string][]Option) {
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

func groupByStrike(options []Option) (strikeToOptions map[float64][]Option) {
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

func groupByProvider(options []Option) (providerToOptions map[string][]Option) {
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

func newOptionIndicators(o Option) Limit {
	bPx, bSz, aPx, aSz := BestLimitHTML(o)

	return Limit{
		Bid: StrOrder{Px: bPx, Size: bSz},
		Ask: StrOrder{Px: aPx, Size: aSz},
	}
}
