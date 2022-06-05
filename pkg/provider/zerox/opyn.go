// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package zerox

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Khan/genqlient/graphql"

	"github.com/teal-finance/rainbow/pkg/provider/the-graph/opyn"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

type Provider struct{}

func (Provider) Name() string {
	return "Opyn"
}

func (Provider) Options() ([]rainbow.Option, error) {
	instruments := QueryTheGraph()

	options, err := normalize(instruments, "Opyn", 10.0)
	if err != nil {
		log.Print("ERR: ", err)
		return nil, err
	}

	return options, err
}

func QueryTheGraph() []opyn.OptionsOtokensOToken {
	const url = "https://api.thegraph.com/subgraphs/name/opynfinance/gamma-mainnet"
	const skip = 0
	const first = 100
	const minExpiry = "1651300000"

	graphqlClient := graphql.NewClient(url, nil)

	resp, err := opyn.Options(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Print("ERR Opyn: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR Opyn: resp=nil")
		return nil
	}

	log.Printf("Query Opyn: minExpiry=%v => %+v", minExpiry, resp)

	// always filter to not get old options
	return resp.Otokens // filterExpired(resp.Otokens, time.Now())
}

func filterExpired(instruments []opyn.OptionsOtokensOToken, date time.Time) (filtered []opyn.OptionsOtokensOToken) {
	for _, i := range instruments {
		seconds, err := strconv.ParseInt(i.ExpiryTimestamp, 10, 64)
		if err != nil {
			log.Print("ERR Opyn ExpiryTimestamp: ", err)
			return nil
		}

		expiryTime := time.Unix(seconds, 0)
		// we keep an option even 2 days after expiry
		// mainly because not all protocol stop at expiry or right before
		// TODO re-check later
		if expiryTime.After(date.Add(-time.Hour * 48)) {
			filtered = append(filtered, i)
		}
	}

	return filtered
}
