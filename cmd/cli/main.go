// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"math"
	"os"

	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/api"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
)

func main() {
	service := rainbow.NewService(provider.AllProvidersNoAlerter(), dbram.NewDB())
	service.FetchOptionsFromProviders()

	options, err := service.Options(rainbow.StoreArgs{})
	if err != nil {
		panic(err)
	}

	printTable(options)
}

func printTable(options []rainbow.Option) {
	green := color.FgGreen.Render
	red := color.FgRed.Render

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.SetTitle(fmt.Sprint(" Available options: ", len(options)))

	t.AppendHeader(table.Row{
		"Provider", "Asset", "Type", "Size", green(" Bid"), "Strike", red(" Ask"), "Size", "Instrument",
	})

	align := api.NewAlign()
	for i := range options {
		o := &options[i]
		bestBidPx, bestBidQty, bestAskPx, bestAskQty := align.BestLimitStr(o)
		t.AppendRows([]table.Row{{
			highlight(o.Provider), o.Asset, o.Type,
			bestBidQty, green(bestBidPx), math.Round(o.Strike*100) / 100,
			red(bestAskPx), bestAskQty, o.Name,
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
	lightGreen := color.FgLightGreen.Render
	blue := color.FgCyan.Render
	darkGray := color.FgDarkGray.Render
	darkGreen := color.FgGreen.Render
	darkBlue := color.FgBlue.Render

	switch p {
	case "Opyn":
		return blue(p)
	case "Deribit":
		return lightGreen(p)
	case "PsyOptions":
		return magenta(p)
	case "Zeta":
		return darkGray(p)
	case "Lyra":
		return darkGreen(p)
	case "Delta Exchange":
		return darkBlue(p)
	default:
		return p
	}
}
