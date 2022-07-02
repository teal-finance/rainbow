// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"bytes"
	"log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/teal-finance/rainbow/pkg/provider/thales"
)

func main() {
	markets := thales.QueryAllMarkets()
	log.Print(markets)

	m := thales.QueryMarket("0x5416c2ab11c7852ed9648aa006ee69d412c735c9")
	log.Print(m)
}

func name(c []byte) string {
	l := common.HexToHash(string(c))
	b := l.Bytes()
	b = bytes.Trim(b, "\x00")
	return string(b)
}
