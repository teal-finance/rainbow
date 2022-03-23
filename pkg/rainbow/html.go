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
	// digitsInFractionalPart is the number of digits to keep in the fractional part.
	digitsInFractionalPart = 2

	// dashRightAlignHTML must show the dash with the number of trailing digits defined by digitsInfractionalPart.
	dashRightAlignHTML = "&mdash;&numsp;&numsp;"

	// dashLeftAlignHTML visual width must be same as the the number of digits defined by len(spaces)-1.
	dashLeftAlignHTML = "&numsp;&numsp;&mdash;"

	// dashLeftAlign must contain the same number of runes as len(spaces)-1.
	dashLeftAlign = "   —"
)

// these variables are used as const to avoid unnecessary allocations.
var (
	// "&nbsp;" --> []byte{byte('&'), byte('n'), byte('b'), byte('s'), byte('p'), byte(';')}.
	narrowSp = []byte{byte('&'), byte('#'), byte('x'), byte('2'), byte('0'), byte('2'), byte('f'), byte(';')}
	boldTag  = []byte{byte('<'), byte('b'), byte('>')}
	boldEnd  = []byte{byte('<'), byte('/'), byte('b'), byte('>')}
	spaces   = [5]byte{32, 32, 32, 32, 32} // len(spaces) must be the max digits wanted before the decimal point
	buffer   = make([]byte, 0, 10)
	htmlBuf  = boldTag
)

func BestLimitHTML(option Option) (bidPx, bidSz, askPx, askSz string) {
	if len(option.Bid) > 0 && option.Bid[0].Size != 0 {
		bidPx = rightAlignFloatOnDecimalPointHTML(option.Bid[0].Price)
		bidSz = leftAlignFloatOnDecimalPointHTML(option.Bid[0].Size)
	} else {
		bidPx, bidSz = dashRightAlignHTML, dashLeftAlignHTML
	}

	if len(option.Ask) > 0 && option.Ask[0].Size != 0 {
		askPx = rightAlignFloatOnDecimalPointHTML(option.Ask[0].Price)
		askSz = leftAlignFloatOnDecimalPointHTML(option.Ask[0].Size)
	} else {
		askPx, askSz = dashRightAlignHTML, dashLeftAlignHTML
	}

	return bidPx, bidSz, askPx, askSz
}

func BestLimitStr(option Option) (bidPx, bidSz, askPx, askSz string) {
	if len(option.Bid) > 0 && option.Bid[0].Size != 0 {
		bidPx = leftAlignFloatOnDecimalPointStr(option.Bid[0].Price)
		bidSz = leftAlignFloatOnDecimalPointStr(option.Bid[0].Size)
	} else {
		bidPx, bidSz = dashLeftAlign, dashLeftAlign
	}

	if len(option.Ask) > 0 && option.Ask[0].Size != 0 {
		askPx = leftAlignFloatOnDecimalPointStr(option.Ask[0].Price)
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
	return string(RightAlign(f, true))
}

func leftAlignFloatOnDecimalPoint(f float64) []byte {
	buffer = strconv.AppendFloat(buffer[:0], f, 'f', 2, 64)

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

func RightAlign(f float64, addMissingDot bool) []byte {
	htmlBuf = strconv.AppendFloat(htmlBuf[:len(boldTag)], f, 'f', -1, 64)

	// i is the position of the '.'
	for i := 1 + len(boldTag); i < len(htmlBuf); i++ {
		if htmlBuf[i] == '.' {
			tinyIntegralPart := i == 1+len(boldTag)

			if end := i + 1 + digitsInFractionalPart; end > len(htmlBuf) {
				htmlBuf = append(htmlBuf, byte('0'))
			} else {
				htmlBuf = htmlBuf[:end]
			}

			if i > 3+len(boldTag) {
				htmlBuf = insertThousandSeparator(htmlBuf, i-3)
				i += len(narrowSp)
			}

			if tinyIntegralPart {
				return append(htmlBuf, byte('<'), byte('/'), byte('b'), byte('>'))
			}

			return append(htmlBuf[:i], append(boldEnd, htmlBuf[i:]...)...)
		}
	}

	// Missing dot

	if len(htmlBuf) > 3+len(boldTag) {
		htmlBuf = insertThousandSeparator(htmlBuf, len(htmlBuf)-3)
	}

	htmlBuf = append(htmlBuf, byte('<'), byte('/'), byte('b'), byte('>'))

	if addMissingDot {
		htmlBuf = append(htmlBuf, byte('.'), byte('0'), byte('0'))
	}

	return htmlBuf
}

// insertThousandSeparator inserts "&#x202f;" (narrow no-break space) that is thinner than "&nbsp;".
// Recommendation is to use "&thinsp;" but it may break the line and "&numsp;" is too wide.
func insertThousandSeparator(b []byte, i int) []byte {
	return append(b[:i], append(narrowSp, b[i:]...)...)
}
