// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"log"

	"github.com/teal-finance/garcon"
	"github.com/teal-finance/rainbow/pkg/provider"
)

const (
	defaultProviders = "deribit,delta,lyra,opyn,psy,zeta"
)

var (
	cex       = flag.Bool("cex", false, "Enable the centralized exchanges: Deribit and Delta Exchange")
	dex       = flag.Bool("dex", false, "Enable the decentralized exchanges: Lyra, Opyn, PsyOptions and Zeta")
	exotic    = flag.Bool("exotic", false, "Enable the decentralized exchanges with binary options: Thales")
	providers = flag.String("providers", garcon.EnvStr("PROVIDERS", defaultProviders), "Coma-separated list of providers, same as the PROVIDERS env. var.")
)

func parseFlags() {
	garcon.SetVersionFlag()
	flag.Parse()

	garcon.LogVersion()
	log.Print("Centralized exchanges     -cex       = ", *cex)
	log.Print("Decentralized exchanges   -dex       = ", *dex)
	log.Print("Binary options exchanges  -exotic    = ", *exotic)
	log.Print("PROVIDERS                 -providers = ", *providers)
}

func listProviderNames() []string {
	if *cex || *dex || *exotic {
		if *providers == defaultProviders {
			*providers = ""
		}
		if *cex {
			*providers = "deribit,delta"
		}
		if *dex {
			*providers += ",lyra,opyn,psy,zeta"
		}
		if *exotic {
			*providers += ",thales"
		}
	}

	return garcon.ExtractWords(*providers, provider.Names())
}
