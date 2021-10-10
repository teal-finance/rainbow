// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package zerox

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Khan/genqlient/graphql"
)

const opynURL = "https://api.thegraph.com/subgraphs/name/opynfinance/gamma-mainnet"

func Instruments() []getOptionsOtokensOToken {
	graphqlClient := graphql.NewClient(opynURL, nil)

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
	for _, i := range instruments {
		seconds, err := strconv.ParseInt(i.ExpiryTimestamp, 10, 0)
		if err != nil {
			fmt.Println("Oh Sh*t ", i.ExpiryTimestamp)
			continue //TODO should do much better than failing silently
		}
		expiryTime := time.Unix(seconds, 0).Add(48 * time.Hour) // we keep a matket for two days after expiry
		if expiryTime.After(time.Now()) {
			filtered = append(filtered, i)
		}

	}
	return filtered
}
