// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"log"

	"github.com/teal-finance/garcon"
)

var (
	btc   = flag.Bool("btc", false, "Fetch options having underlying=BTC, same as -asset BTC")
	eth   = flag.Bool("eth", false, "Fetch options having underlying=ETH, same as -asset ETH")
	asset = flag.String("asset", "ALL", "underlying asset of the options to query")
	jwt   = flag.String("jwt", garcon.EnvStr("JWT"), "JSON-Web-Token to authenticate to the Rainbow API. "+
		"The JWT environment variable can also be used.")
	hmac = flag.String("hmac", garcon.EnvStr("HMAC_SHA256"), "HMAC-SHA256 key (64 hex digits) to generate JWT tokens. "+
		"The HMAC_SHA256 environment variable can also be used.")
)

func parseFlags() {
	garcon.SetVersionFlag()
	flag.Parse()

	switch {
	case *btc:
		*asset = "BTC"
	case *eth:
		*asset = "ETH"
	}

	garcon.LogVersion()
	log.Print("-btc    = ", *btc)
	log.Print("-eth    = ", *eth)
	log.Print("-asset  = ", *asset)
	log.Print("-jwt  len=", len(*jwt))
	log.Print("-hmac len=", len(*hmac), " (need 64 hexadecimal digits) env. var. HMAC_SHA256")
}
