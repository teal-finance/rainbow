// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/garcon"
)

var log = emo.NewZone("client", true, true, true, true)

func main() {
	garcon.LogVersion()
	log.PrintAll = true

	parseFlags()

	if *asset != "" {
		*asset += "/"
	}
	*rURL = strings.Trim(*rURL, "/")
	url := *rURL + "/v0/options/" + *asset + "tsv"

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, http.NoBody)
	if err != nil {
		log.Error(err)
		os.Exit(10)
	}

	if *hmac != "" {
		*access = garcon.NewAccessToken(*ttl, *usr, groups, orgs, *hmac)
		log.AccessToken(*access)
	}

	req.Header.Set("Authorization", "Bearer "+*access)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Error(err)
		os.Exit(11)
	}

	maxBytes := 1_000_000
	buf, err := garcon.ReadResponse(resp, maxBytes)
	resp.Body.Close()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if *verbose {
		log.Result(string(buf))
	}

	log.Ok("Fetched " + garcon.ConvertSize(len(buf)) + " from " + url)
}
