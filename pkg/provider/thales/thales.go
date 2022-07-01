package thales

import (
	"context"
	"log"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/davecgh/go-spew/spew"

	"github.com/teal-finance/rainbow/pkg/provider/the-graph/thales"
)

const (
	urlOptimism = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
	urlPolygon  = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
	skip        = 0
	first       = 100
)

type Provider struct{}

func (Provider) Name() string {
	return "Thales"
}

func twoWeeksInThePast() int64 {
	const twoWeeks = -14 * 24 * time.Hour
	return time.Now().Add(-twoWeeks).Unix()
}

func QueryMarket(id string) *thales.MarketMarket {
	graphqlClient := graphql.NewClient(urlOptimism, nil)
	resp, err := thales.Market(context.TODO(), graphqlClient, id)
	if err != nil {
		log.Print("ERR QueryMarket: ", err)
	}
	if resp == nil {
		log.Print("ERR QueryMarket: resp=nil")
		return nil
	}
	return &resp.Market
}

func QueryRangedMarket(id string) *thales.RangedMarketRangedMarket {
	graphqlClient := graphql.NewClient(urlOptimism, nil)
	resp, err := thales.RangedMarket(context.TODO(), graphqlClient, id)
	if err != nil {
		log.Print("ERR QueryRangedMarket: ", err)
	}
	if resp == nil {
		log.Print("ERR QueryRangedMarket: resp=nil")
		return nil
	}
	return &resp.RangedMarket
}

func QueryMarkets() []thales.MarketsMarketsMarket {
	graphqlClient := graphql.NewClient(urlOptimism, nil)
	minExpiry := twoWeeksInThePast()
	resp, err := thales.Markets(context.TODO(), graphqlClient, skip, first, minExpiry)
	spew.Dump(resp)
	if err != nil {
		log.Print("ERR QueryMarkets: ", err)
	}
	if resp == nil {
		log.Print("ERR QueryMarkets: resp=nil")
		return nil
	}
	return resp.Markets
}

func QueryRangedMarkets() []thales.RangedMarketsRangedMarketsRangedMarket {
	graphqlClient := graphql.NewClient(urlOptimism, nil)
	minExpiry := twoWeeksInThePast()
	resp, err := thales.RangedMarkets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Print("ERR QueryRangedMarkets: ", err)
	}
	if resp == nil {
		log.Print("ERR QueryRangedMarkets: resp=nil")
		return nil
	}
	return resp.RangedMarkets
}
