package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"
)

const url = "https://api.mainnet-beta.solana.com"
const normalserummarket = "ByRys5tuUWDgL73G8JBAEfkdFf8JWBzPBDHsBVQ5vbQA"

func main() {
	serumMarketAddresses := []string{"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
		"7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
		"9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
		"ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28"}
	client := rpc.NewClient(url)

	pubKey := solana.MustPublicKeyFromBase58(serumMarketAddresses[2])

	out, err := serum.FetchMarket(context.TODO(), client, pubKey)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
