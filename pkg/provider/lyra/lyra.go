// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package lyra

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	optimismrpc = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m" // "https://mainnet.optimism.io"
	name        = "Lyra"
)

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	client, err := ethclient.Dial(optimismrpc)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	for i := 0; i < len(OptionMarkets); i++ {
		market, err := NewLyra(common.HexToAddress(OptionMarkets[i]), client)
		if err != nil {
			return nil, err
		}

		viewer, err := NewLyrap(common.HexToAddress(OptionMarketViewers[i]), client)
		if err != nil {
			return nil, err
		}

		boards, err := market.GetLiveBoards(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}

		for _, j := range boards {
			boardListings, err := market.GetBoardListings(&bind.CallOpts{}, j)
			if err != nil {
				return nil, err
			}

			sum += len(boardListings)

			for _, k := range boardListings {
				vlist, err := viewer.GetListingView(&bind.CallOpts{}, k)
				if err != nil {
					return nil, err
				}

				amount := 1

				ammOrder, err := viewer.getBidsAsks(k, amount)
				if err != nil {
					return nil, err
				}

				callPut := processOption(vlist, ammOrder, amount, Assets[i])
				options = append(options, callPut...)
			}
		}
	}

	log.Print("total markets ", sum)
	return options, nil
}

func (v *Lyrap) getBidsAsks(boardListing *big.Int, amount int) ([]OptionMarketViewerTradePremiumView, error) {
	ammOrder := []OptionMarketViewerTradePremiumView{}
	a := big.NewInt(10)
	a.Exp(big.NewInt(10), big.NewInt(18), nil)
	a.Mul(a, big.NewInt(int64(amount)))

	// Call BID
	trade, err := v.GetPremiumForOpen(&bind.CallOpts{}, boardListing, 1, a)
	if err != nil {
		return nil, err
	}

	ammOrder = append(ammOrder, trade)

	// Call ASK
	trade, err = v.GetPremiumForOpen(&bind.CallOpts{}, boardListing, 0, a)
	if err != nil {
		return nil, err
	}

	ammOrder = append(ammOrder, trade)

	// Put BID
	trade, err = v.GetPremiumForOpen(&bind.CallOpts{}, boardListing, 3, a)
	if err != nil {
		return nil, err
	}

	ammOrder = append(ammOrder, trade)

	// Put ASK
	trade, err = v.GetPremiumForOpen(&bind.CallOpts{}, boardListing, 2, a)
	if err != nil {
		return nil, err
	}

	ammOrder = append(ammOrder, trade)
	return ammOrder, err
}

func processOption(listing OptionMarketViewerListingView, ammOrder []OptionMarketViewerTradePremiumView, amount int, asset string) []rainbow.Option {
	options := []rainbow.Option{}
	call := rainbow.Option{
		Name:          "",
		Type:          "CALL",
		Asset:         asset,
		Expiry:        expiration(listing.Expiry),
		ExchangeType:  "DEX",
		Chain:         "Ethereum",
		Layer:         "L2",
		Provider:      name,
		QuoteCurrency: "USD", // sUSD but anyway
		Bid:           nil,
		Ask:           nil,
		Strike:        ToFloat(listing.Strike),
	}
	put := rainbow.Option{
		Name:          "",
		Type:          "PUT",
		Asset:         asset,
		Expiry:        expiration(listing.Expiry),
		ExchangeType:  "DEX",
		Chain:         "Ethereum",
		Layer:         "L2",
		Provider:      name,
		QuoteCurrency: "USD", // sUSD but anyway
		Bid:           nil,
		Ask:           nil,
		Strike:        ToFloat(listing.Strike),
	}

	call.Name = optionName(call)
	put.Name = optionName(put)

	call.Bid = append(call.Bid, rainbow.Order{
		Price: ToFloat(ammOrder[0].Premium),
		Size:  float64(amount),
	})

	call.Ask = append(call.Ask, rainbow.Order{
		Price: ToFloat(ammOrder[1].Premium),
		Size:  float64(amount),
	})

	put.Bid = append(put.Bid, rainbow.Order{
		Price: ToFloat(ammOrder[2].Premium),
		Size:  float64(amount),
	})

	put.Ask = append(put.Ask, rainbow.Order{
		Price: ToFloat(ammOrder[3].Premium),
		Size:  float64(amount),
	})

	options = append(options, call, put)
	return options
}

func optionName(o rainbow.Option) string {
	return o.Asset + "-" + o.Expiry + "-" +
		fmt.Sprintf("%.2f", o.Strike) + "-" + o.Type
}

func expiration(e *big.Int) string {
	seconds := e.Int64()
	expiryTime := time.Unix(seconds, 0).UTC()
	return expiryTime.Format("2006-01-02 15:04:05")
}

// because Ethereum use 10^18 to represent 1
// what we do is that we cut 10^13 to keep 5 decimals
// (because IV is a percentage an we want to be accurate)
// then convert the remainder to float64 and divide by 1000.
func ToFloat(n *big.Int) float64 {
	q := common.Big0
	q.Quo(n, big.NewInt(10000000000000)) // divided by 10^13
	return float64(q.Int64()) / 100000.0
}
