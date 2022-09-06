// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/garcon"
)

var log = emo.NewZone("main")

func main() {
	garcon.LogVersion()

	parseFlags()

	if *asset != "" {
		*asset += "/"
	}
	*rURL = strings.Trim(*rURL, "/")
	url := *rURL + "/v0/options/" + *asset + "tsv"

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	if *hmac != "" {
		*access = garcon.NewAccessToken(*ttl, *usr, groups, orgs, *hmac)
		log.AccessToken(*access)
	}

	req.Header.Set("Authorization", "Bearer "+*access)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Fatal(err)
	}

	maxBytes := 1_000_000
	buf, err := garcon.ReadResponse(resp, maxBytes)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if *verbose {
		log.Result(string(buf))
	}

	log.Ok("Fetched " + garcon.ConvertSize(len(buf)) + " from " + url)
}
