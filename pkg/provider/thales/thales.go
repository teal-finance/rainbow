// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package thales

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/davecgh/go-spew/spew"

	"github.com/teal-finance/rainbow/pkg/provider/the-graph/thales"
)

const (
	selection   = "polygon"
	urlOptimism = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
	urlPolygon  = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
	skip        = 0
	first       = 100
)

var url = selectURL()

func selectURL() string {
	switch selection {
	case "optimism":
		return urlOptimism
	default:
		return urlPolygon
	}
}

type Provider struct{}

func (Provider) Name() string {
	return "Thales"
}

func twoWeeksInThePast() string {
	const twoWeeks = -14 * 24 * time.Hour
	t := time.Now().Add(-twoWeeks).Unix()
	return strconv.FormatInt(t, 10)
}

func QueryAllMarkets() []thales.AllMarketsMarketsMarket {
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
	spew.Dump(resp)
	return resp.Markets
}

func QueryMarkets() []thales.MarketsMarketsMarket {
	graphqlClient := graphql.NewClient(url, nil)
	minExpiry := twoWeeksInThePast()
	resp, err := thales.Markets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Print("ERR thales.Markets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.Markets: resp=nil")
		return nil
	}
	spew.Dump(resp)
	return resp.Markets
}

func QueryMarket(id string) *thales.MarketMarket {
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

func QueryAllRangedMarkets() []thales.AllRangedMarketsRangedMarketsRangedMarket {
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
	spew.Dump(resp)
	return resp.RangedMarkets
}

func QueryRangedMarkets() []thales.RangedMarketsRangedMarketsRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	minExpiry := twoWeeksInThePast()
	resp, err := thales.RangedMarkets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Print("ERR thales.RangedMarkets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.RangedMarkets: resp=nil")
		return nil
	}
	spew.Dump(resp)
	return resp.RangedMarkets
}

func QueryRangedMarket(id string) *thales.RangedMarketRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.RangedMarket(context.TODO(), graphqlClient, id)
	spew.Dump(resp)
	if err != nil {
		log.Print("ERR thales.RangedMarket: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.RangedMarket: resp=nil")
		return nil
	}
	spew.Dump(resp)
	return &resp.RangedMarket
}
