package main

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/davecgh/go-spew/spew"
	"github.com/streamingfast/solana-go"

	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"
	"github.com/teal-finance/rainbow/pkg/provider/psyoptions/anchor"
)

const (
	mainnet           = "https://api.mainnet-beta.solana.com"
	serummain         = "https://solana-api.projectserum.com"
	devnet            = "https://api.devnet.solana.com"
	normalserummarket = "ByRys5tuUWDgL73G8JBAEfkdFf8JWBzPBDHsBVQ5vbQA"
)

func getClient() *rpc.Client {
	return rpc.NewClient(serummain)
}
func main() {

	//spew.Dump(instruments)
	rawOptions, err := anchor.Query()
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(rawOptions)
	client := getClient()
	for _, i := range rawOptions {
		pubKey := solana.PublicKey(i.SerumMarketAddress())

		ctx := context.TODO()

		_, err := serum.FetchMarket(ctx, client, pubKey)
		if err != nil {
			//spew.Dump(pubKey)
			fmt.Println(pubKey, err)
		} else {
			fmt.Println(pubKey, " OK")
		}
	}
	x := 7.60000005                      //12.3456
	fmt.Println(math.Floor(x*100) / 100) // 12.34 (round down)
	fmt.Println(math.Round(x*100) / 100) // 12.35 (round to nearest)
	fmt.Println(math.Ceil(x*100) / 100)  // 12.35 (round up)
	x = math.Floor(x*100) / 100
	fmt.Println(x)

}
