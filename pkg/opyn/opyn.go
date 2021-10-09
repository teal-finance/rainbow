// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package opyn

import (
	"context"
	"log"

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

	return resp.Otokens
}
