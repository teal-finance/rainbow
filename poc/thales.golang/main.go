// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"log"

	"github.com/spewerspew/spew"

	"github.com/teal-finance/rainbow/pkg/provider/thales"
)

func main() {
	log.Print("-------- thales.QueryAllMarkets --------")
	url := thales.LayerURL("Optimism")
	all, err := thales.QueryAllMarkets(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("len =", len(all))
	if len(all) > 0 {
		spew.Dump(all[0])
	}

	log.Print("-------- thales.QueryMarkets --------")
	markets, err := thales.QueryMarkets(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("len =", len(markets))
	if len(markets) > 0 {
		spew.Dump(markets[0])
	}

	log.Print("-------- thales.QueryMarket --------")
	m, err := thales.QueryMarket("0x5416c2ab11c7852ed9648aa006ee69d412c735c9", url)
	if err != nil {
		log.Fatal(err.Error())
	}

	spew.Dump(m)
}
