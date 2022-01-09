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
	PsyQuoteCurrency = "USDC"
	mainnet          = "https://solana-api.projectserum.com" // https://api.mainnet-beta.solana.com"
)

type Provider struct{}

func (Provider) Name() string {
	return "PsyOptions"
}

func (p Provider) Options() ([]rainbow.Option, error) {
	// instruments := append(oldInstruments("ETH"), oldInstruments("BTC")...)
	instruments := query()
	client := rpc.NewClient(mainnet)

	options := make([]rainbow.Option, 0, len(instruments))

	for _, i := range instruments {
		pubKey := solana.MustPublicKeyFromBase58(i.SerumMarketAddress)

		ctx := context.TODO()

		out, err := serum.FetchMarket(ctx, client, pubKey)
		if err != nil {
			return nil, fmt.Errorf("serum.FetchMarket: %w", err)
		}
		// inversing the order to be able to quickly find the best bid (bids[0]) and ask (asks[len(offer)-1])
		bids, _, err := normalizeOrders(ctx, out, client, out.Market.GetBids(), true, i.contractSize())
		if err != nil {
			return nil, err
		}

		asks, _, err := normalizeOrders(ctx, out, client, out.Market.GetAsks(), false, i.contractSize())
		if err != nil {
			return nil, err
		}

		options = append(options, rainbow.Option{
			Name:          i.name(),
			Type:          i.optionType(),
			Asset:         i.asset(),
			Expiry:        i.expiration(),
			Strike:        i.strike(),
			ExchangeType:  "DEX",
			Chain:         "Solana",
			Layer:         "L1",
			Provider:      "PsyOptions",
			QuoteCurrency: PsyQuoteCurrency,
			Bid:           bids,
			Ask:           asks,
		})
	}

	return options, nil
}

// I don't really need the totalsize but I am keeping it since it was in the original func:
//     - ASK on the top so desc=true
//     - BID down so desc=false
func normalizeOrders(ctx context.Context, market *serum.MarketMeta, cli *rpc.Client, address solana.PublicKey, desc bool, contractSize float64) (offers []rainbow.Order, totalSize float64, err error) {
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
				Price: price / contractSize, //to get the price for 1 asset since psyoptions has <1 contract size
				Size:  qty * contractSize,   // to convert the right quantity
			},
		)
	}

	return offers, totalSize, nil
}
