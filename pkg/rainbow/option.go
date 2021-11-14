// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"strconv"
	"strings"
)

type Option struct {
	Name          string  `json:"name"`     // ASSET-DATE-Strike-OptionsType
	Type          string  `json:"type"`     // CALL / PUT
	Asset         string  `json:"asset"`    // ETH, BTC, SOL
	Expiry        string  `json:"expiry"`   // Expiry date in format 2021-12-31
	ExchangeType  string  `json:"exchange"` // CEX / DEX
	Chain         string  `json:"chain"`    // Ethereum, Solana and "–" for CEX (Deribit)
	Layer         string  `json:"layer"`    // L1, L2 and "–" for CEX (Deribit)
	Provider      string  `json:"provider"` // Opyn, Lyra, Thales, Deribit, Psyoptions
	QuoteCurrency string  `json:"currency"` // ETH, BTC, USDT...
	Bid           []Order `json:"bid"`
	Ask           []Order `json:"ask"`
	Strike        float64 `json:"strike"`
}

type Order struct {
	// Px is a known abbreviation for "Price" used by many Centralized Exchanges
	// such as in the FIX protocol: https://fiximate.fixtrading.org/legacy/en/FIX.5.0SP2/abbreviations.html
	// Independently of the FIX protocol, "Px" is intensively used at Euronext.
	// "Px" is also a French abbreviation for "Prix". :-)
	// Rainbow uses "px" in the API (JSON)) but uses "price" on the front-end side.
	Px float64 `json:"px"`

	// Size is often used in lieu of the longer word "Quantity".
	Size float64 `json:"size"`
}

func BestLimitHTML(option Option) (bidPx, bidSz, askPx, askSz string) {
	if len(option.Bid) > 0 && option.Bid[0].Size != 0 {
		bidPx = alignFloatOnDecimalPointHTML(option.Bid[0].Px)
		bidSz = alignFloatOnDecimalPointHTML(option.Bid[0].Size)
	} else {
		bidPx, bidSz = dashHTML, dashHTML
	}

	if len(option.Ask) > 0 && option.Ask[0].Size != 0 {
		askPx = alignFloatOnDecimalPointHTML(option.Ask[0].Px)
		askSz = alignFloatOnDecimalPointHTML(option.Ask[0].Size)
	} else {
		askPx, askSz = dashHTML, dashHTML
	}

	return bidPx, bidSz, askPx, askSz
}

func BestLimitStr(option Option) (bidPx, bidSz, askPx, askSz string) {
	if len(option.Bid) > 0 && option.Bid[0].Size != 0 {
		bidPx = alignFloatOnDecimalPointStr(option.Bid[0].Px)
		bidSz = alignFloatOnDecimalPointStr(option.Bid[0].Size)
	} else {
		bidPx, bidSz = dash, dash
	}

	if len(option.Ask) > 0 && option.Ask[0].Size != 0 {
		askPx = alignFloatOnDecimalPointStr(option.Ask[0].Px)
		askSz = alignFloatOnDecimalPointStr(option.Ask[0].Size)
	} else {
		askPx, askSz = dash, dash
	}

	return bidPx, bidSz, askPx, askSz
}

// dashHTML visual width must same the number of digits defined by len(spaces)-1.
const dashHTML = "&numsp;&numsp;&mdash;"

// dash must contain the same number of runes as len(spaces)-1.
const dash = "   —"

// these two variables avoid unnecessary allocation by alignFloatOnDecimalPoint().
var (
	spaces = [5]byte{32, 32, 32, 32, 32} // len(spaces) must be the max digits wanted before the decimal point
	buffer = make([]byte, 0, 10)
)

func alignFloatOnDecimalPointStr(f float64) string {
	return string(alignFloatOnDecimalPoint(f))
}

func alignFloatOnDecimalPointHTML(f float64) string {
	s := string(alignFloatOnDecimalPoint(f))
	return strings.ReplaceAll(s, " ", "&numsp;")
}

func alignFloatOnDecimalPoint(f float64) []byte {
	buffer = strconv.AppendFloat(buffer[:0], f, 'f', -1, 64)

	var i int // position of the '.' (if no '.' => i = len(b))
	for i = 1; i < len(buffer); i++ {
		if buffer[i] == '.' {
			break
		}
	}

	if i >= len(spaces) {
		return buffer
	}

	return append(spaces[:len(spaces)-i-1], buffer...)
}
