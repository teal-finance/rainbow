// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"context"

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
	ctx := context.TODO()
	serumMarketAddresses := []string{
		"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
		"7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
		"9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
		"ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28",
	}

	err := runE(ctx, serumMarketAddresses)
	if err != nil {
		panic(err)
	}
}
