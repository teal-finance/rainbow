// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package opyn

import (
	"context"
	"log"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

const opynURL = "https://api.thegraph.com/subgraphs/name/opynfinance/gamma-mainnet"

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

func Instruments() []getOptionsOtokensOToken {
	// -- key := os.Getenv("OPYN_API_KEY")
	// -- if key == "" {
	// -- 	log.Print("Missing OPYN_API_KEY=<secret>")
	// -- }
	// --
	// -- httpClient := http.Client{
	// -- 	Transport: &authedTransport{
	// -- 		key:     key,
	// -- 		wrapped: http.DefaultTransport,
	// -- 	},
	// -- }

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
