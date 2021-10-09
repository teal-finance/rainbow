// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"context"
	"fmt"

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
	tryPsyops()
}

func tryOpyn() {
	instruments := zerox.Instruments("ETH")
	// fmt.Println(markets)
	// spew.Dump(markets[1:2])
	orderBook, err := zerox.GetOrderBook(instruments[1:2], "Opyn")
	if err != nil {
		fmt.Println(err)
		return
	}

	spew.Dump(orderBook[0])

	fmt.Println(zerox.ConvertToSolidity(10.0, 8))

	orderBook, err = zerox.GetAggregatedOrderBook(instruments, "Opyn", 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	spew.Dump(orderBook)
}

func tryDeribit() {
	instruments, err := deribit.Instruments("BTC")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(instruments[10])
	spew.Dump(instruments[10])

	orderBook, err := deribit.GetOrderBook(instruments[10:15], 5)
	if err != nil {
		fmt.Println(err)
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

	//spew.Dump(out)

	fmt.Print("asks ", out.Market.GetAsks(), "\n\n\n")

	orders, err := serum.FetchOpenOrders(ctx, client, out.Market.GetAsks())
	if err != nil {
		panic(err)
	}

	// spew.Dump(orders)
	fmt.Println(orders.OpenOrders.GetOrder(0))

	pubKey = solana.MustPublicKeyFromBase58(normalserummarket)

	out, err = serum.FetchMarket(ctx, client, pubKey)
	if err != nil {
		panic(err)
	}

	orders, err = serum.FetchOpenOrders(ctx, client, out.Market.GetAsks())
	if err != nil {
		panic(err)
	}
	/*
		var index uint32
		for index = 0; index < 2; index++ {
			order := orders.OpenOrders.GetOrder(index)
			// spew.Dump(firstOrder)
			fmt.Println("price", order.Price())
			fmt.Println("seqnum", order.SeqNum())
			fmt.Println("side ", order.Side)
		}*/
	offers, total, err := psyoptions.BidsAsksToOffers(ctx, out, client, out.Market.GetAsks(), true, "SELL")
	if err != nil {
		panic(err)
	}
	fmt.Println("total ", total)
	spew.Dump(offers[0:2])

}
