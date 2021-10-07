// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package psyoptions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/teal-finance/rainbow"

	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"
)

const listMarketsURL = "wss://api.psyoptions.io/v1/graphql"

func GetListMarkets(coin string) []string {
	if coin == "ETH" {
		return []string{
			"8A493gU55NfS4fCjDoLAiN57zPzWf6QQw31QQf1fd6iX",
			"5pHcU2Gz8eCMwynLvz1AHSFoFbKkeTc7ufeqeG4spb99",
			"8qgAVVE6eeJ9u32LvJupgW6NyNWPPFFK4xnXfGtDNeP4",
			"FfkVR6ha6N9fcefGxfDFP97xL54Pc5ipiwnt1uuTpuyX",
			"TLAzx53rb6pEDT3oWitTzg6YNe7BTERLDzujMdU8RQx",
			"9A1m1JXgnMrdq3JkqEYx9rK3CcYm9EHDLeA6sn1hH3PP",
			"HYmPvo8szh62QVaAfUAXR1eppvCfouUPpH68yE87UYmy",
			"8fFcWuVaZSKoge4DCpMcNrR5nNXF2pbXCfBUxkMomgr5",
		}
	}
	//"BTC"
	return []string{"8fhiAYm41RwtiX8WusCSpY617GWPt2LwUnCQcEeer78o",
		"6at26sVk8vTYtLh4YDKvje4PDdgFJsNHHyoGw87WNszP",
		"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
		"7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
		"9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
		"ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28",
		"2q5f1H8xT3tsBzQhwZC3BKnbKMb44fTuDGamZ6xUdZz2",
		"DvohGwDZR9Z2siWBj2Xhgxd1qRScVCpywL3EoRbpon3p",
	}
}

func GetOrderBook(ctx context.Context, market *serum.MarketMeta, cli *rpc.Client, address solana.PublicKey, desc bool, side string) (offers []rainbow.Offer, totalSize *big.Float, err error) {
	var o serum.Orderbook
	if err := cli.GetAccountDataIn(ctx, address, &o); err != nil {
		return nil, nil, fmt.Errorf("getting orderbook: %w", err)
	}

	limit := 20
	levels := [][]*big.Int{}

	o.Items(desc, func(node *serum.SlabLeafNode) error {
		quantity := big.NewInt(int64(node.Quantity))
		price := node.GetPrice()
		if len(levels) > 0 && levels[len(levels)-1][0].Cmp(price) == 0 {
			current := levels[len(levels)-1][1]
			levels[len(levels)-1][1] = new(big.Int).Add(current, quantity)
		} else if len(levels) == limit {
			return fmt.Errorf("done")
		} else {
			levels = append(levels, []*big.Int{price, quantity})
		}
		return nil
	})

	totalSize = big.NewFloat(0)
	for _, level := range levels {
		//price := market.PriceLotsToNumber(level[0]) TODO
		qty := market.BaseSizeLotsToNumber(level[1])
		totalSize = new(big.Float).Add(totalSize, qty)
		offers = append(offers,
			rainbow.Offer{
				Price:         1.0, //price, TODO
				Quantity:      2.0, //qty, TODO
				QuoteCurrency: "USD",
				Side:          side,
			},
		)
	}
	return offers, totalSize, nil
}
