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

	"github.com/davecgh/go-spew/spew"
	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets/anchor"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const ZetaAPIUrl = "https://api.zeta.markets/markets/"

var log = emo.NewZone("Zeta")

type Provider struct{}

func (Provider) Name() string {
	return "Zeta"
}

func (p Provider) Options() ([]rainbow.Option, error) {
	rawOptions, m, err := anchor.Query()
	if err != nil {
		return nil, fmt.Errorf("anchor.query: %w", err)
	}
	//spew.Dump(m)
	//spew.Dump(&m)
	oi := OpenInterestMap(m)
	spew.Dump(oi)

	client := rpc.NewClient(anchor.SolanaRPC)

	options := make([]rainbow.Option, 0, len(rawOptions))

	for _, i := range rawOptions {
		pubKey := solana.PublicKey(i.SerumAddress())

		out, err := serum.FetchMarket(context.Background(), client, pubKey)
		if err != nil {
			// previously just silently skipped because of a weird error from Serum
			return []rainbow.Option{}, log.Error("serum FetchMarket", err).Err()

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
			MarketIV:      i.Vol(),
			URL:           "https://dex.zeta.markets/",
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
	var levels [][]*big.Int

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

func OpenInterestMap(m map[string][]uint64) map[string]ZetaAPI {
	//if error fail with just a log
	//let's keep it like that for now because evyrything shouldn't fail if zeta api is down
	oi := make(map[string]ZetaAPI)
	/*
		for asset, expiries := range m {

			for _, e := range expiries {
				url := ZetaAPIUrl + asset + "?expiry=" + strconv.FormatUint(e, 10)
				spew.Dump(url)
				resp, err := http.Get(url)
				if err != nil {
					log.Warnf("zeta open interest GET %s: %w", url, err)
					return nil
				}
				defer resp.Body.Close()
				//spew.Dump(resp.Body)

				var result struct {
					Result []ZetaAPI `json:"result"`
				}
				if err = gg.DecodeJSONResponse(resp, &result.Result); err != nil {
					//return Quote{}, fmt.Errorf("quote from Ox: %w", err)
					return nil
				}
				json.NewDecoder(resp.Body).Decode(&result)
				//json.Unmarshal(resp.Body, &result)
				spew.Dump(result)
			}

		}
	*/

	return oi

}

type ZetaAPI struct {
	OI ZetaOI
}
type ZetaOI struct {
	Timestamp         int64   `json:"timestamp"`
	OpenInterest      int     `json:"open_interest"`
	ImpliedVolatility float64 `json:"implied_volatility"`
	Delta             int     `json:"delta"`
	Vega              int     `json:"vega"`
	Theo              float64 `json:"theo"`
}
