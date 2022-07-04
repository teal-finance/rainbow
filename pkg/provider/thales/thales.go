// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package thales

import (
	"bytes"
	"context"
	"log"

	"github.com/Khan/genqlient/graphql"
	"github.com/ethereum/go-ethereum/common"

	"github.com/teal-finance/rainbow/pkg/provider/the-graph/thales"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	selection   = "optimism"
	urlOptimism = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
	urlPolygon  = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
	OptimismRPC = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m"
	PolygonRPC  = "https://polygon-mainnet.g.alchemy.com/v2/7MGFstWkvX-GscRyBQxehyisRlEoQWyu"
	name        = "Thales"
	skip        = 0
	first       = 100
)

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	return nil, nil
}
func QueryAllLiveMarkets(url string) []thales.AllLiveMarketsMarket {

	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllLive(context.TODO(), graphqlClient)
	if err != nil {
		log.Print("ERR: ", err)
	}

	if resp == nil {
		log.Print("ERR: resp=nil")
		return nil
	}
	return resp.Markets
}

func QueryAllMarkets(url string) []thales.AllMarketsMarketsMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllMarkets(context.TODO(), graphqlClient, skip, first)
	if err != nil {
		log.Print("ERR thales.AllMarkets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.AllMarkets: resp=nil")
		return nil
	}
	return resp.Markets
}

func QueryMarkets(url string) []thales.MarketsMarketsMarket {
	graphqlClient := graphql.NewClient(url, nil)
	minExpiry := rainbow.TwoWeeksInThePast()
	resp, err := thales.Markets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Print("ERR thales.Markets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.Markets: resp=nil")
		return nil
	}
	return resp.Markets
}

func QueryMarket(id, url string) *thales.MarketMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.Market(context.TODO(), graphqlClient, id)
	if err != nil {
		log.Print("ERR thales.Market: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.Market: resp=nil")
		return nil
	}
	return &resp.Market
}

func QueryAllRangedMarkets(url string) []thales.AllRangedMarketsRangedMarketsRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllRangedMarkets(context.TODO(), graphqlClient, skip, first)
	if err != nil {
		log.Print("ERR thales.AllRangedMarkets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.AllRangedMarkets: resp=nil")
		return nil
	}
	return resp.RangedMarkets
}

func QueryRangedMarkets(url string) []thales.RangedMarketsRangedMarketsRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	minExpiry := rainbow.TwoWeeksInThePast()
	resp, err := thales.RangedMarkets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Print("ERR thales.RangedMarkets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.RangedMarkets: resp=nil")
		return nil
	}
	return resp.RangedMarkets
}

func QueryRangedMarket(id, url string) *thales.RangedMarketRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.RangedMarket(context.TODO(), graphqlClient, id)
	if err != nil {
		log.Print("ERR thales.RangedMarket: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.RangedMarket: resp=nil")
		return nil
	}
	return &resp.RangedMarket
}

func Underlying(s string) string {
	l := common.HexToHash(s)
	b := l.Bytes()
	b = bytes.Trim(b, "\x00")
	return string(b)
}
