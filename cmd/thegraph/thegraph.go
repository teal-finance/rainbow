// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
)

const opynURL = "https://gateway.thegraph.com/api/[api-key]/subgraphs/id/0xfc3ac80003d8a5181e554d03983284e4341a7610-0"

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

func main() {
	key := os.Getenv("OPYN_API_KEY")
	if key == "" {
		log.Print("Missing GITHUB_TOKEN=<github token>")
	}

	httpClient := http.Client{
		Transport: &authedTransport{
			key:     key,
			wrapped: http.DefaultTransport,
		},
	}

	graphqlClient := graphql.NewClient(opynURL, &httpClient)

	resp, err := getOptions(context.Background(), graphqlClient)
	if err != nil {
		log.Print("ERROR:", err)
	}

	log.Print(resp)
}
