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
	"strings"

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
	*rURL = strings.Trim(*rURL, "/")
	url := *rURL + "/v0/options/" + *asset + "tsv"

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, http.NoBody)
	if err != nil {
		log.Error(err)
		os.Exit(10)
	}

	if *hmac != "" {
		*jwtRefresh = newJWT(*hmac, *ns, *user)
		log.RefreshToken(*jwtRefresh)
	}

	req.Header.Set("Authorization", "Bearer "+*jwtRefresh)

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

func newJWT(hexKey, namespace, username string) string {
	if len(hexKey) != 64 {
		log.Error("Want HMAC-SHA256 key composed by 64 hexadecimal digits, but got ", len(hexKey))
		os.Exit(20)
	}

	binKey, err := hex.DecodeString(hexKey)
	if err != nil {
		log.Error("Cannot decode the HMAC-SHA256 key, please provide 64 hexadecimal digits: ", err)
		os.Exit(21)
	}

	token, err := tokens.GenRefreshToken("1y", "1y", namespace, username, binKey)
	if err != nil || token == "" {
		log.Error("Cannot create JWT: ", err)
		os.Exit(22)
	}

	return token
}
