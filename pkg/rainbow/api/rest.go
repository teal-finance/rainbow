// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mpvl/unique"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/garcon/gg"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone("api")

type Handler struct {
	Service *rainbow.Service
	align   Align
}

func NewHandler(s *rainbow.Service) Handler {
	return Handler{
		Service: s,
		align:   NewAlign(),
	}
}

func (h Handler) Options(w http.ResponseWriter, r *http.Request) {
	sa, format, err := query(r)
	if err != nil {
		log.Warn("Options", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	options, _ := h.Service.Options(sa)
	if len(options) == 0 {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}

	err = h.writeOptions(w, options, sa, format)
	if err != nil {
		log.Warnf("Options options=%v sa=%v format=%v err=%v", len(options), sa, format, err)
	}
}

func query(r *http.Request) (sa rainbow.StoreArgs, format string, _ error) {
	if err := r.ParseForm(); err != nil {
		return sa, "", err
	}

	assets, err := gg.Values(r, "asset")
	if err != nil {
		return sa, "", err
	}

	expiries, err := gg.Values(r, "expiry")
	if err != nil {
		return sa, "", err
	}

	providers, err := gg.Values(r, "provider")
	if err != nil {
		return sa, "", err
	}

	format, err = gg.Value(r, "format", "Accept")
	if err != nil {
		return sa, "", err
	}

	unique.Sort(unique.StringSlice{P: &assets})
	unique.Sort(unique.StringSlice{P: &expiries})
	unique.Sort(unique.StringSlice{P: &providers})

	specialValues(&format, &assets)
	specialValues(&format, &expiries)
	specialValues(&format, &providers)

	sa = rainbow.StoreArgs{
		Assets:    assets,
		Expiries:  expiries,
		Providers: providers,
	}
	return sa, format, nil
}

// specialValues removes the special values from the values slide.
// Special values are the empty string, "ALL"
// and the supported formats ("csv", "tsv", "json").
// specialValues also sets the format.
func specialValues(format *string, values *[]string) {
	i := 0
	for _, v := range *values {
		switch v {
		case "", "ALL":
			continue

		case "csv", "tsv", "json":
			*format = v
			continue
		}

		(*values)[i] = v
		i++
	}

	if i == 0 {
		*values = nil
	} else {
		*values = (*values)[:i]
	}
}

func (h Handler) writeOptions(w http.ResponseWriter, options []rainbow.Option, sa rainbow.StoreArgs, format string) error {
	switch {
	case format == "": // The most frequent first
		return h.replyJSON(w, options)

	// We use strings.Contains() because format can be
	// "text/csv" when provided by the "Accept" header.
	// This looks stupid but it's just simple.
	case strings.Contains(format, "csv"):
		setFilename(w, sa, ".csv")
		return h.replyCSV(w, options, ',')

	case strings.Contains(format, "tsv"):
		setFilename(w, sa, ".tsv")
		return h.replyCSV(w, options, '\t')

	case strings.Contains(format, "json"):
		setFilename(w, sa, ".json")
		fallthrough

	default:
		return h.replyJSON(w, options)
	}
}

func (h Handler) replyJSON(w http.ResponseWriter, options []rainbow.Option) error {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(options)
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return err
	}

	return nil
}

func (h Handler) replyCSV(w http.ResponseWriter, options []rainbow.Option, comma rune) error {
	// Write in CSV format.
	csvWriter := csv.NewWriter(w)
	csvWriter.Comma = comma

	w.Header().Set("Content-Type", "text/csv")

	// Header
	row := []string{ // #xx is the number of the CSV column (used below has row[xx])
		"Name",         // #00
		"Expiry",       // #01
		"Asset",        // #02
		"Currency",     // #03
		"Strike",       // #04
		"Type",         // #05
		"Provider",     // #06
		"Layer",        // #07
		"Chain",        // #08 // if empty --> ExchangeType
		"Bid Price #1", // #09
		"Bid Size #1",  // #10
		"Bid Price #2", // #11
		"Bid Size #2",  // #12
		"Ask Price #1", // #13
		"Ask Size #1",  // #14
		"Ask Price #2", // #15
		"Ask Size #2",  // #16
	}

	for i := 0; ; i++ {
		err := csvWriter.Write(row)
		if err != nil {
			return err
		}

		if i >= len(options) {
			break
		}

		row[0] = options[i].Name
		row[1] = options[i].Expiry
		row[2] = options[i].Asset
		row[3] = options[i].QuoteCurrency
		row[4] = fmt.Sprint(options[i].Strike)
		row[5] = options[i].Type
		row[6] = options[i].Provider
		row[7] = options[i].Layer

		if options[i].Chain == "" {
			row[8] = options[i].ExchangeType
		} else {
			row[8] = options[i].Chain
		}

		if len(options[i].Bid) > 0 {
			row[9] = fmt.Sprint(options[i].Bid[0].Price)
			row[10] = fmt.Sprint(options[i].Bid[0].Size)
		} else {
			row[9] = ""
			row[10] = ""
		}

		if len(options[i].Bid) > 1 {
			row[11] = fmt.Sprint(options[i].Bid[1].Price)
			row[12] = fmt.Sprint(options[i].Bid[1].Size)
		} else {
			row[11] = ""
			row[12] = ""
		}

		if len(options[i].Ask) > 0 {
			row[13] = fmt.Sprint(options[i].Ask[0].Price)
			row[14] = fmt.Sprint(options[i].Ask[0].Size)
		} else {
			row[13] = ""
			row[14] = ""
		}

		if len(options[i].Ask) > 1 {
			row[15] = fmt.Sprint(options[i].Ask[1].Price)
			row[16] = fmt.Sprint(options[i].Ask[1].Size)
		} else {
			row[15] = ""
			row[16] = ""
		}
	}

	csvWriter.Flush()
	err := csvWriter.Error()
	if err != nil {
		return err
	}

	return nil
}

// setFilename builds the CSV filename.
func setFilename(w http.ResponseWriter, sa rainbow.StoreArgs, extension string) {
	name := "AllOptions"

	if len(sa.Assets) > 0 || len(sa.Expiries) > 0 || len(sa.Providers) > 0 {
		name = "AllAssets"
		if len(sa.Assets) > 0 {
			name = strings.Join(sa.Assets, "+")
		}

		if len(sa.Expiries) > 0 {
			name += "-" + strings.Join(sa.Expiries, "+")
		}

		if len(sa.Providers) > 0 {
			name += "-" + strings.Join(sa.Providers, "+")
		}
	}

	name += extension

	w.Header().Set("Content-Disposition", "attachment; filename="+name)
}
