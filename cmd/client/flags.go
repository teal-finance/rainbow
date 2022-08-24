// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"flag"

	"github.com/teal-finance/garcon"
)

const (
	rainbowURL  = "https://teal.finance/rainbow"
	quidURL     = "https://teal.finance/quid"
	defaultUser = "client"
	defaultNs   = "FreePlan"
)

var (
	rURL = flag.String("rainbow", garcon.EnvStr("RAINBOW_URL", rainbowURL), "URL of the Rainbow website. The RAINBOW_URL environment variable can also be used.")
	qURL = flag.String("quid", garcon.EnvStr("QUID_URL", quidURL), "URL of the Quid website. The QUID_URL environment variable can also be used.")

	btc   = flag.Bool("btc", false, "Fetch options having underlying=BTC, same as -asset BTC")
	eth   = flag.Bool("eth", false, "Fetch options having underlying=ETH, same as -asset ETH")
	asset = flag.String("asset", "ALL", "underlying asset of the options to query")

	jwtRefresh = flag.String("jwt", garcon.EnvStr("JWT"), "Refresh JSON-Web-Token to authenticate to the Rainbow API. "+
		"The JWT environment variable can also be used.")
	hmac = flag.String("hmac", garcon.EnvStr("HMAC_SHA256"), "HMAC-SHA256 key (64 hex digits) to generate JWT tokens. "+
		"The HMAC_SHA256 environment variable can also be used.")
	ns = flag.String("ns", garcon.EnvStr("JWT_NS", defaultNs), "Namespace to set when generating a new JWT "+
		"from the HMAC-SHA256 key provided by the -hmac flag. "+
		"The JWT_NS environment variable can also be used.")
	user = flag.String("user", garcon.EnvStr("JWT_USER", defaultUser), "User name to set when generating a new JWT "+
		"from the HMAC-SHA256 key provided by the -hmac flag. "+
		"The JWT_USER environment variable can also be used.")

	verbose = flag.Bool("v", false, "Verbose: print the fetch data from Rainbow API")
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

	log.Param("-rainbow RAINBOW_URL =", *rURL)
	log.Param("-quid    QUID_URL    =", *qURL)
	log.Param("-btc                 =", *btc)
	log.Param("-eth                 =", *eth)
	log.Param("-asset               =", *asset)
	log.Param("-ns      JWT_NS      =", *user)
	log.Param("-user    JWT_USER    =", *user)
	log.Param("-jwt     JWT         len=", len(*jwtRefresh))
	log.Param("-hmac    HMAC_SHA256 len=", len(*hmac), "(need 64 hexadecimal digits)")
	log.Param("-v                   =", *verbose)
}
