// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/teal-finance/garcon"
)

func main() {
	parseFlags()

	if *asset != "" {
		*asset += "/"
	}

	url := "https://teal.finance/rainbow/v0/options/" + *asset + "tsv"

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, http.NoBody)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+*jwt)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	maxBytes := 1_000_000
	buf, err := garcon.ReadResponse(resp, maxBytes)
	if err != nil {
		log.Print(err)
	}

	log.Print(string(buf))
}
