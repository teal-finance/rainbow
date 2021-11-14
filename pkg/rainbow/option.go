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

// digitsInfractionalPart is the number of digits to keep in the fractional part.
const digitsInfractionalPart = 2

// dashRightAlignHTML must show the dash with the number of trailing digits defined by digitsInfractionalPart.
const dashRightAlignHTML = "&mdash;&numsp;&numsp;"

// dashLeftAlignHTML visual width must be same as the the number of digits defined by len(spaces)-1.
const dashLeftAlignHTML = "&numsp;&numsp;&mdash;"

// dashLeftAlign must contain the same number of runes as len(spaces)-1.
const dashLeftAlign = "   —"

// these two variables avoid unnecessary allocation by alignFloatOnDecimalPoint().
var (
	dotZrs = []byte{byte('.'), byte('0'), byte('0')} // len(dotZrs) must be 1+digitsInfractionalPart
	spaces = [5]byte{32, 32, 32, 32, 32}             // len(spaces) must be the max digits wanted before the decimal point
	buffer = make([]byte, 0, 20)
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
		bidPx = rightAlignFloatOnDecimalPointHTML(option.Bid[0].Px)
		bidSz = leftAlignFloatOnDecimalPointHTML(option.Bid[0].Size)
	} else {
		bidPx, bidSz = dashRightAlignHTML, dashLeftAlignHTML
	}

	if len(option.Ask) > 0 && option.Ask[0].Size != 0 {
		askPx = rightAlignFloatOnDecimalPointHTML(option.Ask[0].Px)
		askSz = leftAlignFloatOnDecimalPointHTML(option.Ask[0].Size)
	} else {
		askPx, askSz = dashRightAlignHTML, dashLeftAlignHTML
	}

	return bidPx, bidSz, askPx, askSz
}

func BestLimitStr(option Option) (bidPx, bidSz, askPx, askSz string) {
	if len(option.Bid) > 0 && option.Bid[0].Size != 0 {
		bidPx = leftAlignFloatOnDecimalPointStr(option.Bid[0].Px)
		bidSz = leftAlignFloatOnDecimalPointStr(option.Bid[0].Size)
	} else {
		bidPx, bidSz = dashLeftAlign, dashLeftAlign
	}

	if len(option.Ask) > 0 && option.Ask[0].Size != 0 {
		askPx = leftAlignFloatOnDecimalPointStr(option.Ask[0].Px)
		askSz = leftAlignFloatOnDecimalPointStr(option.Ask[0].Size)
	} else {
		askPx, askSz = dashLeftAlign, dashLeftAlign
	}

	return bidPx, bidSz, askPx, askSz
}

func leftAlignFloatOnDecimalPointStr(f float64) string {
	return string(leftAlignFloatOnDecimalPoint(f))
}

func leftAlignFloatOnDecimalPointHTML(f float64) string {
	s := string(leftAlignFloatOnDecimalPoint(f))
	return strings.ReplaceAll(s, " ", "&numsp;")
}

func rightAlignFloatOnDecimalPointHTML(f float64) string {
	return string(rightAlign(f))
}

func leftAlignFloatOnDecimalPoint(f float64) []byte {
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

func rightAlign(f float64) []byte {
	buffer = strconv.AppendFloat(buffer[:0], f, 'f', -1, 64)

	var i int // position of the '.'
	for i = 1; i < len(buffer); i++ {
		if buffer[i] == '.' {
			end := i + 1 + digitsInfractionalPart

			if end > len(buffer) {
				return append(buffer, byte('0'))
			}

			return buffer[:end]
		}
	}

	// missing dot
	return append(buffer, dotZrs...)
}
