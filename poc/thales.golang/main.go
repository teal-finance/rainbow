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

	m := thales.QueryMarket("0xd9e137e43fb366fb500553c2066b805616d9b0d4")
	log.Print(m)

	httpResponse := `{"data":{"markets":[{
		"id":"0xd9e137e43fb366fb500553c2066b805616d9b0d4",
		"timestamp":"1656423710",
		"creator":"0x5027ce356c375a934b4d1de9240ba789072a5af1",
		"currencyKey":"0x4f484d0000000000000000000000000000000000000000000000000000000000",
		"strikePrice":"18000000000000000000",
		"maturityDate":"1657526400",
		"expiryDate":"1673251200",
		"isOpen":true,
		"poolSize":"10000000000000000",
		"longAddress":"0x252b0c0d3268ca8c54eb0b4ddfdb252bb2681f0d",
		"shortAddress":"0xa3487b9c100c25a162d2df56dc00b4fdf4cb4d15",
		"result":null,
		...`
	log.Print("HTTP response from TheGraph:", httpResponse)
}

func name(c []byte) string {
	l := common.HexToHash(string(c))
	b := l.Bytes()
	b = bytes.Trim(b, "\x00")
	return string(b)
}
