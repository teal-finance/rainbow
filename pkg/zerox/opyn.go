// Copyright (c) 2021 teal.finance
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

	"github.com/teal-finance/rainbow"
)

func Options() ([]rainbow.Option, error) {
	instruments := QueryTheGraph()

	options, err := normalize(instruments, "Opyn", 2.0)
	if err != nil {
		log.Print("ERROR: ", err)
		return nil, err
	}

	return options, err
}

func QueryTheGraph() []getOptionsOtokensOToken {
	const url = "https://api.thegraph.com/subgraphs/name/opynfinance/gamma-mainnet"

	graphqlClient := graphql.NewClient(url, nil)

	resp, err := getOptions(context.Background(), graphqlClient)
	if err != nil {
		log.Print("ERROR: ", err)
	}

	if resp == nil {
		log.Print("ERROR: resp=nil")
		return nil
	}

	return filterExpired(resp.Otokens)
}

func filterExpired(instruments []getOptionsOtokensOToken) (filtered []getOptionsOtokensOToken) {
	oct29th, _ := time.Parse(time.RFC3339, "2021-10-29T08:00:00Z")

	for _, i := range instruments {
		seconds, err := strconv.ParseInt(i.ExpiryTimestamp, 10, 0)
		if err != nil {
			log.Print("Oh Sh*t ", i.ExpiryTimestamp)
			continue // TODO should do much better than failing silently
		}

		expiryTime := time.Unix(seconds, 0)
		if expiryTime.Equal(oct29th) {
			filtered = append(filtered, i)
		}
	}

	return filtered
}
