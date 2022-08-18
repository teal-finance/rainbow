// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package zetamarkets

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"

	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets/anchor"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// const serumMainnet = "https://solana-api.projectserum.com" // "https://api.mainnet-beta.solana.com"

type Provider struct {
	Endpoint string
}

func (Provider) Name() string {
	return "Zeta"
}

func (p Provider) Options() ([]rainbow.Option, error) {
	rawOptions, err := anchor.Query()
	if err != nil {
		return nil, fmt.Errorf("anchor.query: %w", err)
	}

	client := rpc.NewClient(p.Endpoint)

	options := make([]rainbow.Option, 0, len(rawOptions))

	for _, i := range rawOptions {
		pubKey := solana.PublicKey(i.SerumAddress())

		out, err := serum.FetchMarket(context.Background(), client, pubKey)
		if err != nil {
			// for now because error in serumAddress generated
			continue
		}

		// inverting the order to be able to quickly find the best bid (bids[0]) and ask (asks[len(offer)-1])
		bids, err := normalizeOrders(out, client, out.Market.GetBids(), true, anchor.ContractSize)
		if err != nil {
			return nil, err
		}

		asks, err := normalizeOrders(out, client, out.Market.GetAsks(), false, anchor.ContractSize)
		if err != nil {
			return nil, err
		}

		options = append(options, rainbow.Option{
			Name:          i.Name(),
			Type:          i.OptionType(),
			Asset:         i.Asset(),
			Expiry:        i.Expiration(),
			Strike:        i.Strike(),
			ExchangeType:  "DEX",
			Chain:         "Solana",
			Layer:         "L1",
			Provider:      p.Name(),
			QuoteCurrency: i.Quote(),
			Bid:           bids,
			Ask:           asks,
		})
	}

	if len(options) == 0 {
		return nil, errors.New("empty options lists")
	}
	return options, nil
}

// I don't really need the totalSize but I am keeping it since it was in the original func:
//   - ASK on the top so desc=true
//   - BID down so desc=false
func normalizeOrders(
	market *serum.MarketMeta,
	cli *rpc.Client,
	address solana.PublicKey,
	desc bool,
	contractSize float64,
) ([]rainbow.Order, error) {
	// quick & dirty extra rate limit
	time.Sleep(300 * time.Microsecond)

	var book serum.Orderbook
	if err := cli.GetAccountDataIn(address, &book); err != nil {
		return nil, fmt.Errorf("cli.GetAccountDataIn: %w", err)
	}

	limit := 20
	levels := [][]*big.Int{}

	err := book.Items(desc, func(node *serum.SlabLeafNode) error {
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
		return nil, fmt.Errorf("cli.GetAccountDataIn: %w", err)
	}

	orders := make([]rainbow.Order, 0, len(levels))

	for _, level := range levels {
		pBig := market.PriceLotsToNumber(level[0])
		qBig := market.BaseSizeLotsToNumber(level[1])
		px, _ := pBig.Float64()
		sz, _ := qBig.Float64()

		orders = append(orders,
			rainbow.Order{
				Price: px,
				Size:  sz / contractSize, // to convert the right quantity since there 3 decimals in the raw output
			},
		)
	}

	return orders, nil
}
