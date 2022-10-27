// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/teal-finance/garcon"
	"github.com/teal-finance/garcon/gg"
	"github.com/teal-finance/garcon/timex"
	"github.com/teal-finance/rainbow/pkg/provider"
)

const (
	defaultAddr      = "http://localhost"
	defaultPort      = 8090
	defaultProviders = "ALL"
)

var (
	dev          = flag.Bool("dev", false, "Enable the developer mode (enabled by default if -addr and -port are not used)")
	period       = flag.Duration("period", 10*time.Minute, "Period to fetch market data from providers")
	cex          = flag.Bool("cex", false, "Enable the centralized exchanges: Deribit and Delta Exchange")
	dex          = flag.Bool("dex", false, "Enable the decentralized exchanges: Lyra, Opyn, PsyOptions and Zeta")
	exotic       = flag.Bool("exotic", false, "Enable the decentralized exchanges with binary options: Thales")
	providers    = flag.String("providers", gg.EnvStr("PROVIDERS", defaultProviders), "Coma-separated list of providers, has precedence over PROVIDERS")
	mainAddr     = flag.String("addr", gg.EnvStr("MAIN_ADDR", defaultAddr), "Schema and DNS used for doc URL and CORS, has precedence over MAIN_ADDR")
	mainPort     = flag.Int("port", gg.EnvInt("MAIN_PORT", defaultPort), "API port, has precedence over MAIN_PORT")
	expPort      = flag.Int("exp", gg.EnvInt("EXP_PORT"), "Export port for Prometheus, has precedence over EXP_PORT")
	reqPerMinute = flag.Int("rate", gg.EnvInt("REQ_PER_MINUTE", 88), "Max requests per minute, has precedence over REQ_PER_MINUTE")
	reqBurst     = flag.Int("burst", gg.EnvInt("REQ_BURST", 22), "Max requests during a burst, has precedence over REQ_BURST")
	wwwDir       = flag.String("www", gg.EnvStr("WWW_DIR", "frontend/dist"), "Folder of the web static files, has precedence over WWW_DIR")
	alert        = flag.String("alert", gg.EnvStr("ALERT_URL"), "Webhook endpoint to notify anomalies, has precedence over ALERT_URL")
	form         = flag.String("form", gg.EnvStr("WEBFORM_URL"), "Webhook endpoint to notify filled contact form, has precedence over WEBFORM_URL")
	aes          = flag.String("aes", gg.EnvStr("AES128"), " 128-bit AES key (32 hex digits) for the session cookies, has precedence over AES128")
	hmac         = flag.String("hmac", gg.EnvStr("HMAC_SHA256"), "HMAC-SHA256 key (64 hex digits) for the JWT tokens, has precedence over HMAC_SHA256")
	listenAddr   string
)

func parseFlags() {
	garcon.SetVersionFlag()
	flag.Parse()

	listenAddr = ":" + strconv.Itoa(*mainPort)

	if !*dev && *mainAddr == defaultAddr && *mainPort == defaultPort {
		*dev = true
		log.Init("Enable -dev mode because -addr and -port flags are not used")
	}

	garcon.LogVersion()
	log.Init("Dev. mode      -dev       =", *dev)
	log.Init("Data fetch     -period    =", timex.DStr(*period))
	log.Init("Centralized Ex -cex       =", *cex)
	log.Init("Decentralized  -dex       =", *dex)
	log.Init("Binary options -exotic    =", *exotic)
	log.Init("PROVIDERS      -providers =", *providers)
	log.Init("MAIN_ADDR      -addr      =", *mainAddr)
	log.Init("MAIN_PORT      -port      =", *mainPort)
	log.Init("EXP_PORT       -exp       =", *expPort)
	log.Init("REQ_PER_MINUTE -rate      =", *reqPerMinute)
	log.Init("REQ_BURST      -burst     =", *reqBurst)
	log.Init("WWW_DIR        -www       =", *wwwDir)
	log.Init("ALERT_URL      -alert  len=", len(*alert))
	log.Init("WEBFORM_URL    -form   len=", len(*form))
	log.Init("AES128         -aes    len=", len(*aes), "(need 32 hexadecimal digits)")
	log.Init("HMAC_SHA256    -hmac   len=", len(*hmac), "(need 64 hexadecimal digits)")

	// mandatory: -aes or -hmac
	if *aes == "" && *hmac == "" {
		if *dev {
			*hmac = "9d2e0a02121179a3c3de1b035ae1355b1548781c8ce8538a1dc0853a12dfb13d"
		} else {
			log.Fatal("Missing secret key for the tokens (cookies). " +
				"Please provide it using the -aes or -hmac flag " +
				"(or the AES128 or HMAC_SHA256 env. var.)")
		}
	}
	if len(*aes) > 0 && len(*hmac) > 0 {
		log.Warn("Should use -aes or -hmac, not both in the same time")
	}
}

// listProviderNames is duplicated https://github.com/teal-finance/rainbow/blob/main/cmd/cli/flags.go#L37
func listProviderNames() []string {
	if *cex || *dex || *exotic {
		if *providers == defaultProviders {
			*providers = ""
		}
		if *cex {
			*providers += ",deribit,delta"
		}
		if *dex { // deprecating "opyn" & "psy"
			*providers += ",lyra,zeta"
		}
		if *exotic {
			*providers += ",thales"
		}
	}

	return gg.ExtractWords(*providers, provider.Names())
}
