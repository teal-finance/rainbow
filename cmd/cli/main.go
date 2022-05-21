// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"fmt"
	"math"
	"os"

	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
)

func main() {
	service := rainbow.NewService(provider.AllProviders(), dbram.NewDB())
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
		"Provider", "Asset", "Type", "Bid IV", "Size", green(" Bid"), "Strike", red(" Ask"), "Size", "Ask IV", "Instrument", "Greeks",
	})

	for _, option := range options {
		bidIv, bestBidPx, bestBidQty, askIv, bestAskPx, bestAskQty := rainbow.BestLimitStr(option)

		t.AppendRows([]table.Row{{
			highlight(option.Provider), option.Asset, option.Type,
			bidIv, bestBidQty, green(bestBidPx), math.Round(option.Strike*1000) / 1000,
			red(bestAskPx), bestAskQty, askIv, option.Name,
			option.Greeks,
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
