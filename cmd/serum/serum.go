// copied from https://github.com/streamingfast/solana-go/blob/master/cmd/slnc/cmd/serum_get_market.go

// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"

	"github.com/streamingfast/solana-go/rpc"
)

const mainnet = "https://api.mainnet-beta.solana.com"
const devnet = "https://api.devnet.solana.com"
const normalserummarket = "ByRys5tuUWDgL73G8JBAEfkdFf8JWBzPBDHsBVQ5vbQA"

func getClient() *rpc.Client {
	return rpc.NewClient(mainnet)
}

func main() {
	ctx := context.TODO()
	serumMarketAddresses := []string{
		"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
		"7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
		"9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
		"ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28",
	}

	err := runE(ctx, serumMarketAddresses)
	if err != nil {
		panic(err)
	}
}
