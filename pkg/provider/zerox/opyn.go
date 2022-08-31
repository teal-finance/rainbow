// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package zerox

import (
	"context"
	"strconv"
	"time"

	"github.com/Khan/genqlient/graphql"

	"github.com/teal-finance/rainbow/pkg/provider/the-graph/opyn"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const url = "https://api.thegraph.com/subgraphs/name/opynfinance/gamma-mainnet"

type Provider struct{}

func (Provider) Name() string {
	return "Opyn"
}

func (Provider) Options() ([]rainbow.Option, error) {
	instruments := QueryTheGraph()

	return normalize(instruments, "Opyn", 10.0)
}

func QueryTheGraph() []opyn.OptionsOtokensOToken {
	const skip = 0
	const first = 100

	graphqlClient := graphql.NewClient(url, nil)

	// we keep an option even 2 days after expiry
	// mainly because not all protocol stop at expiry or right before
	minExpiry := time.Now().Add(-time.Hour * 48)
	minExpiryStr := strconv.FormatInt(minExpiry.Unix(), 10)

	resp, err := opyn.Options(context.TODO(), graphqlClient, skip, first, minExpiryStr)
	if err != nil {
		log.Print("ERR Opyn: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR Opyn: resp=nil")
		return nil
	}

	log.Printf("INF Query Opyn: minExpiry=%v => %v options", minExpiry, len(resp.Otokens))

	return resp.Otokens
}
