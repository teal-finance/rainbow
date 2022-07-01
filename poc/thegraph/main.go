// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"context"
	"log"

	"github.com/Khan/genqlient/graphql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spewerspew/spew"

	"github.com/teal-finance/rainbow/pkg/provider/generated/thales"
)

func main() {
	/*options := zerox.QueryTheGraph()
	// zerox.QueryTheGraph()
	spew.Dump(options)*/
	const url = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
	//const url = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
	graphqlClient := graphql.NewClient(url, nil)

	resp, err := thales.Markets(context.TODO(), graphqlClient)
	if err != nil {
		log.Print("ERR: ", err)
	}

	if resp == nil {
		log.Print("ERR: resp=nil")
		return
	}
	spew.Dump(resp.Markets[0:4])
	//name("0x4c554e4100000000000000000000000000000000000000000000000000000000")
	for _, m := range resp.Markets {
		if name(m.CurrencyKey) == "BTC" && m.StrikePrice == "20000000000000000000000" {
			spew.Dump(m)
		}

	}
}

func name(s string) string {
	l := common.HexToHash(s)
	b := l.Bytes()
	b = bytes.Trim(b, "\x00")
	return string(b)
}
