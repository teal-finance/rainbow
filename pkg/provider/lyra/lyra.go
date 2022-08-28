// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package lyra

import (
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
	client, err := ethclient.Dial(optimismrpc)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	var options []rainbow.Option

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

				callPut := processOption(&vlist, ammOrder, amount, Assets[i])
				options = append(options, callPut...)
			}
		}
	}

	log.Print("INF Lyra total markets ", sum)
	return options, nil
}

func (v *Lyrap) getBidsAsks(boardListing *big.Int, amount int) ([]OptionMarketViewerTradePremiumView, error) {
	var ammOrder []OptionMarketViewerTradePremiumView
	a := rainbow.IntToEthereumUint256(amount, rainbow.DefaultEthereumDecimals)

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

func processOption(listing *OptionMarketViewerListingView, ammOrder []OptionMarketViewerTradePremiumView, amount int, asset string) []rainbow.Option {
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
		Strike:        rainbow.ToFloat(listing.Strike, rainbow.DefaultEthereumDecimals),
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
		Strike:        rainbow.ToFloat(listing.Strike, rainbow.DefaultEthereumDecimals),
	}

	call.Name = call.OptionName()
	put.Name = put.OptionName()

	call.Bid = append(call.Bid, rainbow.Order{
		Price: rainbow.ToFloat(ammOrder[0].Premium, rainbow.DefaultEthereumDecimals),
		Size:  float64(amount),
	})

	call.Ask = append(call.Ask, rainbow.Order{
		Price: rainbow.ToFloat(ammOrder[1].Premium, rainbow.DefaultEthereumDecimals),
		Size:  float64(amount),
	})

	put.Bid = append(put.Bid, rainbow.Order{
		Price: rainbow.ToFloat(ammOrder[2].Premium, rainbow.DefaultEthereumDecimals),
		Size:  float64(amount),
	})

	put.Ask = append(put.Ask, rainbow.Order{
		Price: rainbow.ToFloat(ammOrder[3].Premium, rainbow.DefaultEthereumDecimals),
		Size:  float64(amount),
	})

	return []rainbow.Option{call, put}
}

func expiration(e *big.Int) string {
	seconds := e.Int64()
	expiryTime := time.Unix(seconds, 0).UTC()
	return expiryTime.Format("2006-01-02 15:04:05")
}
