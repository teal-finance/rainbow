// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"encoding/hex"
	"net/http"
	"os"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/garcon"
	"github.com/teal-finance/quid/quidlib/tokens"
)

var log = emo.NewZone("client")

func main() {
	log.Print = true

	parseFlags()

	if *asset != "" {
		*asset += "/"
	}
	url := "https://teal.finance/rainbow/v0/options/" + *asset + "tsv"

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, http.NoBody)
	if err != nil {
		log.Error(err)
		return
	}

	if *hmac != "" {
		*jwt = newJWT(*hmac, *ns, *user)
		log.RefreshToken(*jwt)
	}

	req.Header.Set("Authorization", "Bearer "+*jwt)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()

	maxBytes := 1_000_000
	buf, err := garcon.ReadResponse(resp, maxBytes)
	if err != nil {
		log.Error(err)
		return
	}

	log.Result(string(buf))
}

func newJWT(hexKey, namespace, username string) string {
	if len(hexKey) != 64 {
		log.Error("Want HMAC-SHA256 key composed by 64 hexadecimal digits, but got ", len(hexKey))
		os.Exit(1)
	}

	binKey, err := hex.DecodeString(hexKey)
	if err != nil {
		log.Error("Cannot decode the HMAC-SHA256 key, please provide 64 hexadecimal digits: ", err)
		os.Exit(2)
	}

	token, err := tokens.GenRefreshToken("1y", "1y", namespace, username, binKey)
	if err != nil || token == "" {
		log.Error("Cannot create JWT: ", err)
		os.Exit(3)
	}

	return token
}
