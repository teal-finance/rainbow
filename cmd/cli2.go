package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"rainbow/deribit"

	"github.com/davecgh/go-spew/spew"
	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"
)

const mainnet = "https://api.mainnet-beta.solana.com"
const devnet = "https://api.devnet.solana.com"
const normalserummarket = "ByRys5tuUWDgL73G8JBAEfkdFf8JWBzPBDHsBVQ5vbQA"
const thegraphopyn = "https://api.thegraph.com/subgraphs/name/opynfinance/gamma-mainnet"

func main() {
	values := map[string]string{"Referer": "https://v2.opyn.co/",
		"Origin": "https://v2.opyn.co",
		//Host: api.thegraph.com
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:93.0) Gecko/20100101 Firefox/93.0",
		"Accept":     "*/*",
		//Accept-Language: fr-FR,fr;q=0.8,en-US;q=0.5,en;q=0.3
		"Accept-Encoding": "gzip, deflate, br",
		//Referer: https://v2.opyn.co/
		//Content-Type: application/json
		//Origin: https://v2.opyn.co
		"Content-Length": "451",
		"DNT":            "1",
		"Connection":     "keep-alive",
		"Sec-Fetch-Dest": "empty",
		"Sec-Fetch-Mode": "cors",
		"Sec-Fetch-Site": "cross-site",
		"Sec-GPC":        "1",
		"TE":             "trailers",
	}
	json_data, _ := json.Marshal(values)
	clt := http.Client{}
	resp, err := clt.Post(thegraphopyn, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println("errman ", err)
		return
	}
	result := graph{}
	fmt.Println("OK")
	//spew.Dump(resp.Body)

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("json ", err)
		return
	}
	fmt.Println("KO")

	fmt.Println("result ", result)
	spew.Dump(result)
	return
}

func tryDeribit() {
	list, err := deribit.GetMarkets("BTC")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(list[10])
	spew.Dump(list[10])

	orderBook, err := deribit.GetOrderBook(list[10:15], 5)

	if err != nil {
		fmt.Println(err)
		return
	}
	spew.Dump(orderBook[0].Offers)

}

type graph struct {
	Data struct {
		Otokens []struct {
			CollateralAsset struct {
				Decimals int    `json:"decimals"`
				ID       string `json:"id"`
				Symbol   string `json:"symbol"`
			} `json:"collateralAsset"`
			Decimals        int    `json:"decimals"`
			ExpiryTimestamp string `json:"expiryTimestamp"`
			ID              string `json:"id"`
			Implementation  string `json:"implementation"`
			IsPut           bool   `json:"isPut"`
			Name            string `json:"name"`
			StrikeAsset     struct {
				Decimals int    `json:"decimals"`
				ID       string `json:"id"`
				Symbol   string `json:"symbol"`
			} `json:"strikeAsset"`
			StrikePrice     string `json:"strikePrice"`
			Symbol          string `json:"symbol"`
			UnderlyingAsset struct {
				Decimals int    `json:"decimals"`
				ID       string `json:"id"`
				Symbol   string `json:"symbol"`
			} `json:"underlyingAsset"`
		} `json:"otokens"`
	} `json:"data"`
}

func testPsyops() {
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

	fmt.Println("asks ", out.Market.GetAsks(), "\n\n")
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
