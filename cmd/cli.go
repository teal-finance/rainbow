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
}

func tryOpyn() {
	instruments := zerox.Instruments()
	// log.Println(markets)
	// spew.Dump(markets[1:2])
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

func testPsyops() {
	serumMarketAddresses := []string{
		"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
		"7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
		"9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
		"ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28",
	}
	/*serumDevnetAddresses := []string{
		"B6VGQiFN7auykkDgyJfNibkqRwth7dWzxS6w6TjUxmmD",
		"49vNyNhxq9wp96hYtKHtGM6dvbVPzH31WMzMmNwfdEEr",
		"A1puENk4nxMwCPhtT9HT28nDU8CTE5uSFLL7RhdrPD4s",
		"2wriXhTCsUwBvfyYch8X6J8pLW797Vm8Thsz4xRZh1Yb",
	}*/
	client := rpc.NewClient(mainnet)
	// client := rpc.NewClient(devnet)

	pubKey := solana.MustPublicKeyFromBase58(serumMarketAddresses[1])
	// pubKey := solana.MustPublicKeyFromBase58(serumDevnetAddresses[2])

	out, err := serum.FetchMarket(context.TODO(), client, pubKey)
	if err != nil {
		panic(err)
	}

	spew.Dump(out)

	log.Print("asks ", out.Market.GetAsks(), "\n\n\n")

	orders, err := serum.FetchOpenOrders(context.TODO(), client, out.Market.GetAsks())
	if err != nil {
		panic(err)
	}

	// spew.Dump(orders)
	log.Println(orders.OpenOrders.GetOrder(0))

	pubKey = solana.MustPublicKeyFromBase58(normalserummarket)

	out, err = serum.FetchMarket(context.TODO(), client, pubKey)
	if err != nil {
		panic(err)
	}

	orders, err = serum.FetchOpenOrders(context.TODO(), client, out.Market.GetAsks())
	if err != nil {
		panic(err)
	}

	var index uint32
	for index = 0; index < 40; index++ {
		order := orders.OpenOrders.GetOrder(index)
		// spew.Dump(firstOrder)
		log.Println("price", order.Price())
		log.Println("seqnum", order.SeqNum())
		log.Println("side ", order.Side)
	}
}
