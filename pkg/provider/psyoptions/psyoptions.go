// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package psyoptions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	listMarketsURL   = "wss://api.psyoptions.io/v1/graphql"
	PsyQuoteCurrency = "USDC"
	mainnet          = "https://solana-api.projectserum.com" // https://api.mainnet-beta.solana.com"
	devnet           = "https://api.devnet.solana.com"
	Expiration       = "2021-10-29 23:59:59"
)

func oldInstruments(coin string) []string {
	switch coin {
	case "ETH":
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

	case "BTC":
		return []string{
			"8fhiAYm41RwtiX8WusCSpY617GWPt2LwUnCQcEeer78o",
			"6at26sVk8vTYtLh4YDKvje4PDdgFJsNHHyoGw87WNszP",
			"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
			"7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
			"9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
			"ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28",
			"2q5f1H8xT3tsBzQhwZC3BKnbKMb44fTuDGamZ6xUdZz2",
			"DvohGwDZR9Z2siWBj2Xhgxd1qRScVCpywL3EoRbpon3p",
		}
	default:
		return []string{}
	}
}

func Options() (options []rainbow.Option, err error) {
	// instruments := append(oldInstruments("ETH"), oldInstruments("BTC")...)
	instruments := query()
	client := rpc.NewClient(mainnet)

	for _, i := range instruments {
		pubKey := solana.MustPublicKeyFromBase58(i.SerumMarketAddress)

		ctx := context.TODO()

		out, err := serum.FetchMarket(ctx, client, pubKey)
		if err != nil {
			return nil, fmt.Errorf("serum.FetchMarket: %w", err)
		}
		// inversing the order to be able to quickly find the best bid (bids[0]) and ask (asks[len(offer)-1])
		bids, _, err := normalizeOrders(ctx, out, client, out.Market.GetBids(), true)
		if err != nil {
			return nil, err
		}

		asks, _, err := normalizeOrders(ctx, out, client, out.Market.GetAsks(), false)
		if err != nil {
			return nil, err
		}

		options = append(options, rainbow.Option{
			Name:         i.name(),
			Type:         i.optionType(),
			Asset:        i.asset(),
			Expiry:       Expiration,
			Strike:       i.strike(),
			ExchangeType: "DEX",
			Chain:        "Solana",
			Layer:        "L1",
			Provider:     "PsyOptions",
			Bid:          bids,
			Ask:          asks,
		})
	}

	return options, err
}

// I don't really need the totalsize but I am keeping it since it was in the original func:
//     - ASK on the top so desc=true
//     - BID down so desc=false
func normalizeOrders(ctx context.Context, market *serum.MarketMeta, cli *rpc.Client, address solana.PublicKey, desc bool) (offers []rainbow.Order, totalSize float64, err error) {
	var o serum.Orderbook
	if err := cli.GetAccountDataIn(ctx, address, &o); err != nil {
		return nil, 0, fmt.Errorf("cli.GetAccountDataIn: %w", err)
	}

	limit := 20
	levels := [][]*big.Int{}

	err = o.Items(desc, func(node *serum.SlabLeafNode) error {
		quantity := big.NewInt(int64(node.Quantity))
		price := node.GetPrice()
		if len(levels) > 0 && levels[len(levels)-1][0].Cmp(price) == 0 {
			current := levels[len(levels)-1][1]
			levels[len(levels)-1][1] = new(big.Int).Add(current, quantity)
		} else if len(levels) != limit {
			levels = append(levels, []*big.Int{price, quantity})
		}
		return nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("cli.GetAccountDataIn: %w", err)
	}

	for _, level := range levels {
		p := market.PriceLotsToNumber(level[0])
		price, _ := p.Float64()
		q := market.BaseSizeLotsToNumber(level[1])
		qty, _ := q.Float64()
		totalSize += qty

		offers = append(offers,
			rainbow.Order{
				Price:         price,
				Quantity:      qty,
				QuoteCurrency: PsyQuoteCurrency,
			},
		)
	}

	return offers, totalSize, nil
}
