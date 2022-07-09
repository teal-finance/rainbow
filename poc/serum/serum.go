// Copyright 2021-2022 Teal.Finance contributors
// This file is part of Teal.Finance/Rainbow,
// a screener fo DeFi options, under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"log"

	"github.com/streamingfast/solana-go/rpc"
)

const (
	mainnet           = "https://api.mainnet-beta.solana.com"
	devnet            = "https://api.devnet.solana.com"
	normalserummarket = "ByRys5tuUWDgL73G8JBAEfkdFf8JWBzPBDHsBVQ5vbQA"
)

func getClient() *rpc.Client {
	return rpc.NewClient(mainnet)
}

func main() {
	serumMarketAddresses := []string{
		"2fH6seyt7sHqEPyPtwsgxoTeca3EVr2ndajcN1YwvQpF",
		"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
		"7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
		"9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
		"ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28",
	}

	err := runE(context.TODO(), serumMarketAddresses)
	if err != nil {
		log.Panic(err)
	}
}
