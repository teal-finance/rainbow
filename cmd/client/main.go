// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/teal-finance/garcon"
	"github.com/teal-finance/quid/quidlib/tokens"
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

	if *hmac != "" {
		*jwt = newJWT(*hmac)
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

func newJWT(hexKey string) string {
	if len(hexKey) != 64 {
		log.Panic("Want HMAC-SHA256 key composed by 64 hexadecimal digits, but got ", len(hexKey))
	}

	binKey, err := hex.DecodeString(hexKey)
	if err != nil {
		log.Panic("Cannot decode the HMAC-SHA256 key, please provide 64 hexadecimal digits: ", err)
	}

	token, err := tokens.GenRefreshToken("1y", "1y", "FreePlan", "client", binKey)
	if err != nil || token == "" {
		log.Panic("Cannot create JWT: ", err)
	}
	return token
}
