package psyoptions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/serum"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/teal-finance/rainbow/pkg/provider/psyoptions/anchor"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

type Provider struct{}

func (Provider) Name() string {
	return "PsyOptions"
}

func (p Provider) Options() ([]rainbow.Option, error) {
	rawOptions, err := anchor.Query()
	if err != nil {
		return nil, fmt.Errorf("anchor.query: %w", err)
	}
	client := rpc.New(rpc.MainNetBeta_RPC)

	options := make([]rainbow.Option, 0, len(rawOptions))

	for _, i := range rawOptions {
		pubKey := i.SerumMarketAddress()

		ctx := context.TODO()

		out, err := serum.FetchMarket(ctx, client, pubKey)
		if err != nil {
			return nil, fmt.Errorf("serum.FetchMarket: %w", err)
		}
		// inversing the order to be able to quickly find the best bid (bids[0]) and ask (asks[len(offer)-1])
		bids, _, err := normalizeOrders(ctx, out, client, out.MarketV2.Bids, true, i.ContractSize())
		if err != nil {
			return nil, err
		}

		asks, _, err := normalizeOrders(ctx, out, client, out.MarketV2.Asks, false, i.ContractSize())
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
			Provider:      "PsyOptions",
			QuoteCurrency: i.Quote(),
			Bid:           bids,
			Ask:           asks,
		})

	}
	return options, nil
}

// I don't really need the totalsize but I am keeping it since it was in the original func:
//     - ASK on the top so desc=true
//     - BID down so desc=false
func normalizeOrders(ctx context.Context, market *serum.MarketMeta, client *rpc.Client, address solana.PublicKey, desc bool, contractSize float64) (offers []rainbow.Order, totalSize float64, err error) {
	var o serum.Orderbook
	if err := client.GetAccountDataInto(ctx, address, &o); err != nil {
		return nil, 0, fmt.Errorf("getting orderbook: %w", err)
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

	for _, level := range levels {
		p := market.PriceLotsToNumber(level[0])
		price, _ := p.Float64()
		q := market.BaseSizeLotsToNumber(level[1])
		qty, _ := q.Float64()
		totalSize += qty

		offers = append(offers,
			rainbow.Order{
				Px:   price / contractSize, //to get the price for 1 asset since psyoptions has <1 contract size
				Size: qty * contractSize,   // to convert the right quantity
			},
		)
	}

	return offers, totalSize, nil
}
