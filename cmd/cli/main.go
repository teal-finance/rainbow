// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
)

func main() {
	service := rainbow.NewService(provider.AllProvider{}, &dbram.DB{})

	options, err := service.OptionsFromProviders()
	if err != nil {
		log.Print("ERROR: ", err)
		return
	}

	printTable(options)
}

func printTable(options []rainbow.Option) {
	green := color.FgGreen.Render
	red := color.FgRed.Render

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.SetTitle(fmt.Sprint("\t\t CeDeFi options: ", len(options)))

	t.AppendHeader(table.Row{
		"Provider", "Asset", "Type",
		"Bid size", "Bid px", "Strike",
		"Ask px", "Ask size", "Instrument",
	})

	for _, option := range options {
		bestBidPx, bestBidQty, bestAskPx, bestAskQty := rainbow.BestLimitStr(option)

		t.AppendRows([]table.Row{{
			highlight(option.Provider), option.Asset, option.Type,
			bestBidQty, green(bestBidPx), option.Strike,
			red(bestAskPx), bestAskQty, option.Name,
		}})
	}

	t.SortBy([]table.SortBy{
		{Name: "Strike", Mode: table.AscNumeric},
		{Name: "Type", Mode: table.Dsc},
	})

	t.Render()
}

func highlight(p string) string {
	magenta := color.FgMagenta.Render
	green := color.FgGreen.Render
	blue := color.FgCyan.Render

	switch p {
	case "Opyn":
		return blue(p)
	case "Deribit":
		return green(p)
	case "PsyOptions":
		return magenta(p)
	default:
		return p
	}
}
