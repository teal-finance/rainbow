// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"

	"github.com/teal-finance/rainbow/pkg/deribit"
	"github.com/teal-finance/rainbow/pkg/psyoptions"
	"github.com/teal-finance/rainbow/pkg/zerox"
)

const (
	mainnet           = "https://api.mainnet-beta.solana.com"
	devnet            = "https://api.devnet.solana.com"
	normalserummarket = "ByRys5tuUWDgL73G8JBAEfkdFf8JWBzPBDHsBVQ5vbQA"
)

func main() {
	parseFlags()

	if *port != "" && *port != "no" {
		runAPIServer()
	}

	tryOpyn()
	// tryPsyops()
	// all()
}

func tryOpyn() {
	instruments := zerox.Instruments()
	orderBook, err := zerox.GetOrderBook(instruments, "Opyn")
	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(orderBook[0])

	log.Println(zerox.ConvertToSolidity(10.0, 8))

	orderBook, err = zerox.GetAggregatedOrderBook(instruments, "Opyn", 2.0)
	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(orderBook)
}

func tryDeribit() {
	instruments, err := deribit.Instruments("BTC")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(instruments[10])
	spew.Dump(instruments[10])

	orderBook, err := deribit.GetOrderBook(instruments[10:15], 5)
	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(orderBook[0].Offers)
}

func tryPsyops() {
	serumMarketAddresses := psyoptions.Instruments("BTC")
	client := rpc.NewClient(mainnet)
	// client := rpc.NewClient(devnet)

	pubKey := solana.MustPublicKeyFromBase58(serumMarketAddresses[2])

	ctx := context.TODO()
	out, err := serum.FetchMarket(ctx, client, pubKey)
	if err != nil {
		panic(err)
	}

	bids, totalBids, err := psyoptions.BidsAsksToOffers(ctx, out, client, out.Market.GetBids(), false, "BUY")
	if err != nil {
		panic(err)
	}

	asks, totalAsks, err := psyoptions.BidsAsksToOffers(context.TODO(), out, client, out.Market.GetAsks(), true, "SELL")
	if err != nil {
		panic(err)
	}

	log.Println("total ", totalAsks+totalBids)
	offers := append(bids, asks...)
	spew.Dump(offers)
}

func all() {
	spew.Dump(psyoptions.InstrumentsFromAllMarkets())
}
