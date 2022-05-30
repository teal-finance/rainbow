package thales

import (
	"context"
	"log"

	"github.com/Khan/genqlient/graphql"

	"github.com/teal-finance/rainbow/pkg/provider/the-graph/thales"
)

type Provider struct{}

func (Provider) Name() string {
	return "Thales"
}

func Query() []thales.MarketsMarketsMarket {
	const url = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
	const skip = 0
	const first = 100
	const minExpiry = 1651300000

	graphqlClient := graphql.NewClient(url, nil)

	resp, err := thales.Markets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Print("ERR: ", err)
	}
	if resp == nil {
		log.Print("ERR: resp=nil")
		return nil
	}

	return resp.Markets
}
