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

const (
	// digitsInfractionalPart is the number of digits to keep in the fractional part.
	digitsInfractionalPart = 2

	// dashRightAlignHTML must show the dash with the number of trailing digits defined by digitsInfractionalPart.
	dashRightAlignHTML = "&mdash;&numsp;&numsp;"

	// dashLeftAlignHTML visual width must be same as the the number of digits defined by len(spaces)-1.
	dashLeftAlignHTML = "&numsp;&numsp;&mdash;"

	// dashLeftAlign must contain the same number of runes as len(spaces)-1.
	dashLeftAlign = "   —"
)

// these two variables avoid unnecessary allocation by alignFloatOnDecimalPoint().
var (
	dotZrs = []byte{byte('.'), byte('0'), byte('0')} // len(dotZrs) must be 1+digitsInfractionalPart
	spaces = [5]byte{32, 32, 32, 32, 32}             // len(spaces) must be the max digits wanted before the decimal point
	buffer = make([]byte, 0, 20)
)

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
