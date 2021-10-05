package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"
)

const mainnet = "https://api.mainnet-beta.solana.com"
const devnet = "https://api.devnet.solana.com"
const normalserummarket = "ByRys5tuUWDgL73G8JBAEfkdFf8JWBzPBDHsBVQ5vbQA"

func main() {
	serumMarketAddresses := []string{"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
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
	//client := rpc.NewClient(devnet)

	pubKey := solana.MustPublicKeyFromBase58(serumMarketAddresses[1])
	//pubKey := solana.MustPublicKeyFromBase58(serumDevnetAddresses[2])

	out, err := serum.FetchMarket(context.TODO(), client, pubKey)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)

	fmt.Println("asks ", out.Market.GetAsks(), "\n\n\n")
	orders, err := serum.FetchOpenOrders(context.TODO(), client, out.Market.GetAsks())
	//spew.Dump(orders)
	fmt.Println(orders.OpenOrders.GetOrder(0))
	pubKey = solana.MustPublicKeyFromBase58(normalserummarket)
	out, err = serum.FetchMarket(context.TODO(), client, pubKey)
	orders, err = serum.FetchOpenOrders(context.TODO(), client, out.Market.GetAsks())

	var index uint32
	for index = 0; index < 40; index++ {

		order := orders.OpenOrders.GetOrder(index)
		//spew.Dump(firstOrder)
		fmt.Println("price", order.Price())
		fmt.Println("seqnum", order.SeqNum())
		fmt.Println("side ", order.Side)
	}

}
