// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"strings"

	"github.com/LynxAIeu/garcon/gg"
	"github.com/LynxAIeu/garcon/vv"
)

const (
	hmacSHA256DevMod = "9d2e0a02121179a3c3de1b035ae1355b1548781c8ce8538a1dc0853a12dfb13d"
	rainbowURL       = "https://teal.finance/rainbow"
	rainbowLocal     = "http://localhost:8090"
	quidURL          = "https://teal.finance/quid"
	defaultUser      = "client"
	defaultNs        = "FreePlan"
	defaultGrp       = "FreePlan"
	defaultTTL       = "1y"
)

var (
	verbose = flag.Bool("v", false, "Verbose: print the fetch data from Rainbow API")

	rURL = flag.String("rainbow", gg.EnvStr("RAINBOW_URL", rainbowURL), "URL of the Rainbow website. The RAINBOW_URL environment variable can also be used.")
	qURL = flag.String("quid", gg.EnvStr("QUID_URL", quidURL), "URL of the Quid website. The QUID_URL environment variable can also be used.")

	btc   = flag.Bool("btc", false, "Fetch options having underlying=BTC, same as -asset BTC")
	eth   = flag.Bool("eth", false, "Fetch options having underlying=ETH, same as -asset ETH")
	asset = flag.String("asset", "ALL", "underlying asset of the options to query")

	access  = flag.String("access", gg.EnvStr("JWT_ACCESS"), "Access Token (JWT) to be access to the Rainbow API. The JWT_ACCESS environment variable can also be used.")
	refresh = flag.String("refresh", gg.EnvStr("JWT_REFRESH"), "Refresh Token (JWT) to get a temporary Access Token from Quid API. The JWT_REFRESH environment variable can also be used.")
	hmac    = flag.String("hmac", gg.EnvStr("HMAC_SHA256"), "HMAC-SHA256 key (64 hex digits) required to generate a JWT. The HMAC_SHA256 environment variable can also be used.")
	ttl     = flag.String("ttl", gg.EnvStr("JWT_TTL", defaultTTL), "Max TTL (Time To Live) to set when generating a new JWT (similar to the cookie Max-Age). The JWT_TTL environment variable can also be used.")
	ns      = flag.String("ns", gg.EnvStr("JWT_NS", defaultNs), "Namespace to set when generating a new Refresh JWT. The JWT_NS environment variable can also be used.")
	usr     = flag.String("usr", gg.EnvStr("JWT_USR", defaultUser), "User name to set when generating a new JWT. The JWT_USR environment variable can also be used.")
	grp     = flag.String("grp", gg.EnvStr("JWT_GRP", defaultGrp), "List of groups (separated by coma) to set when generating a new Access JWT. The JWT_GRP environment variable can also be used.")
	org     = flag.String("orgs", gg.EnvStr("JWT_ORG"), "List of organizations (separated by coma) to set when generating a new Access JWT. The JWT_ORG environment variable can also be used.")

	groups, orgs []string
)

func parseFlags() {
	vv.SetVersionFlag()
	flag.Parse()

	switch {
	case *btc:
		*asset = "BTC"
	case *eth:
		*asset = "ETH"
	}

	groups = strings.Split(*grp, ",")
	orgs = strings.Split(*org, ",")

	if *access == "" && *hmac == "" && *rURL == rainbowURL {
		log.Param("neither -access nor -hmac nor -rainbow flags used: use dev. mode HMAC-SHA256 key and localhost for rainbow API")
		*hmac = hmacSHA256DevMod
		*rURL = rainbowLocal
	}

	log.Param("-v                   =", *verbose)
	log.Param("-rainbow RAINBOW_URL =", *rURL)
	log.Param("-quid    QUID_URL    =", *qURL)
	log.Param("-btc                 =", *btc)
	log.Param("-eth                 =", *eth)
	log.Param("-asset               =", *asset)
	log.Param("-ttl     JWT_TTL     =", *ttl)
	log.Param("-ns      JWT_NS      =", *ns)
	log.Param("-usr     JWT_USR     =", *usr)
	log.Param("-grp     JWT_GRP     =", groups)
	log.Param("-org     JWT_ORG     =", orgs)
	log.Param("-access  JWT_ACCESS  len=", len(*access))
	log.Param("-refresh JWT_REFRESH len=", len(*refresh))
	log.Param("-hmac    HMAC_SHA256 len=", len(*hmac), "(need 64 hexadecimal digits)")
}
