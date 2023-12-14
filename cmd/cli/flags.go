// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"flag"

	"github.com/teal-finance/garcon"
	"github.com/teal-finance/garcon/gg"
	"github.com/teal-finance/rainbow/pkg/provider"
)

const (
	defaultProviders = "ALL" //"deribit,delta,lyra,synquote,aevo,rysk,sdx"
)

var (
	cex    = flag.Bool("cex", false, "Enable the centralized exchanges: Deribit, Delta Exchange, Thalex")
	dex    = flag.Bool("dex", false, "Enable the decentralized exchanges: Lyra, Synquote, Aevo, Rysk and SDX")
	exotic = flag.Bool("exotic", false, "Enable the decentralized exchanges with binary options: Thales")
	infos  = flag.Bool("infos", false, "Show additional infos on instruments like URL ...")

	providers = flag.String("providers", gg.EnvStr("PROVIDERS", defaultProviders), "Coma-separated list of providers, same as the PROVIDERS env. var.")
)

func parseFlags() {
	garcon.SetVersionFlag()
	flag.Parse()

	garcon.LogVersion()
	log.Init("Centralized exchanges     -cex       =", *cex)
	log.Init("Decentralized exchanges   -dex       =", *dex)
	log.Init("Binary options exchanges  -exotic    =", *exotic)
	log.Init("Additional infos          -infos     =", *infos)

	log.Init("PROVIDERS                 -providers =", *providers)
}

// TODO remove duplicates
// listProviderNames is duplicated https://github.com/teal-finance/rainbow/blob/main/cmd/server/flags.go#L88
func listProviderNames() []string {
	if *cex || *dex || *exotic {
		if *providers == defaultProviders {
			*providers = ""
		}
		if *cex {
			*providers += ",deribit,delta,thalex"
		}
		if *dex { // deprecating "opyn" & "psy" & "zeta"
			*providers += ",lyra,synquote,aevo,rysk,sdx"
		}
		if *exotic {
			*providers += ",thales"
		}
	}

	return gg.ExtractWords(*providers, provider.Names())
}
