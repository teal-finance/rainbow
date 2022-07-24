// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package api

import (
	"strconv"
	"strings"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	// digitsInFractionalPart is the number of digits to keep in the fractional part.
	digitsInFractionalPart = 2

	// dashRightAlignHTML must show the dash with the number of trailing digits defined by digitsInFractionalPart.
	dashRightAlignHTML = "&mdash;&numsp;&numsp;"

	// dashLeftAlignHTML visual width must be same as the the number of digits defined by len(spaces)-1.
	dashLeftAlignHTML = "&numsp;&numsp;&mdash;"

	// dashLeftAlign must contain the same number of runes as len(spaces)-1.
	dashLeftAlign = "   —"
)

// Variables used as const.
var (
	// "&nbsp;" --> []byte{byte('&'), byte('n'), byte('b'), byte('s'), byte('p'), byte(';')}.
	narrowNBSp = []byte{byte('&'), byte('#'), byte('x'), byte('2'), byte('0'), byte('2'), byte('f'), byte(';')}
	boldEnd    = []byte{byte('<'), byte('/'), byte('b'), byte('>')}
	boldTag    = [3]byte{byte('<'), byte('b'), byte('>')}
	spaces     = [5]byte{32, 32, 32, 32, 32} // len(spaces) must be the max digits wanted before the decimal point
)

type Align struct {
	buf []byte
}

// BestLimitHTML is used by the web/API server to align numbers with HTML <tags>.
func (a *Align) BestLimitHTML(option *rainbow.Option) (bidPx, bidSz, askPx, askSz string) {
	if len(option.Bid) > 0 && option.Bid[0].Size != 0 && option.Bid[0].Price != 0 {
		bidPx = a.rightAlignOnDecimalPointHTML(option.Bid[0].Price, true)
		bidSz = a.leftAlignOnDecimalPointHTML(option.Bid[0].Size)
	} else {
		bidPx, bidSz = dashRightAlignHTML, dashLeftAlignHTML
	}

	if len(option.Ask) > 0 && option.Ask[0].Size != 0 && option.Ask[0].Price != 0 {
		askPx = a.rightAlignOnDecimalPointHTML(option.Ask[0].Price, true)
		askSz = a.leftAlignOnDecimalPointHTML(option.Ask[0].Size)
	} else {
		askPx, askSz = dashRightAlignHTML, dashLeftAlignHTML
	}

	return bidPx, bidSz, askPx, askSz
}

// BestLimitStr is used by the CLI to align numbers with whitespaces.
func (a *Align) BestLimitStr(option *rainbow.Option) (bidPx, bidSz, askPx, askSz string) {
	if len(option.Bid) > 0 && option.Bid[0].Size != 0 && option.Bid[0].Price != 0 {
		bidPx = a.leftAlignOnDecimalPoint(option.Bid[0].Price)
		bidSz = a.leftAlignOnDecimalPoint(option.Bid[0].Size)
	} else {
		bidPx, bidSz = dashLeftAlign, dashLeftAlign
	}

	if len(option.Ask) > 0 && option.Ask[0].Size != 0 && option.Ask[0].Price != 0 {
		askPx = a.leftAlignOnDecimalPoint(option.Ask[0].Price)
		askSz = a.leftAlignOnDecimalPoint(option.Ask[0].Size)
	} else {
		askPx, askSz = dashLeftAlign, dashLeftAlign
	}

	return bidPx, bidSz, askPx, askSz
}

func (a *Align) leftAlignOnDecimalPointHTML(f float64) string {
	s := a.leftAlignOnDecimalPoint(f)
	return strings.ReplaceAll(s, " ", "&numsp;")
}

func NewAlign() Align {
	return Align{buf: append(make([]byte, 0, 16), boldTag[:]...)}
}

func (a *Align) leftAlignOnDecimalPoint(f float64) string {
	// reduce allocation: reuse buffer
	a.buf = a.buf[:len(boldTag)]
	a.buf = strconv.AppendFloat(a.buf, f, 'f', 2, 64)

	var i int // position of the '.' (if no '.' => i = len(b))
	for i = len(boldTag) + 1; i < len(a.buf); i++ {
		if a.buf[i] == '.' {
			break
		}
	}

	if i >= len(boldTag)+len(spaces) {
		return string(a.buf[len(boldTag):])
	}

	return string(append(spaces[:len(boldTag)+len(spaces)-i-1], a.buf[len(boldTag):]...))
}

func (a *Align) rightAlignOnDecimalPointHTML(f float64, addMissingDot bool) string {
	// reduce allocation: reuse buffer
	a.buf = a.buf[:len(boldTag)]
	a.buf = strconv.AppendFloat(a.buf, f, 'f', -1, 64)

	// i is the position of the '.'
	for i := 1 + len(boldTag); i < len(a.buf); i++ {
		if a.buf[i] != '.' {
			continue
		}

		tinyIntegralPart := i == 1+len(boldTag)

		if end := i + 1 + digitsInFractionalPart; end > len(a.buf) {
			a.buf = append(a.buf, byte('0'))
		} else {
			a.buf = a.buf[:end]
		}

		if i > 3+len(boldTag) {
			a.buf = insertThousandSeparator(a.buf, i-3)
			i += len(narrowNBSp)
		}

		if tinyIntegralPart {
			return string(append(a.buf, byte('<'), byte('/'), byte('b'), byte('>')))
		}

		return string(append(a.buf[:i], append(boldEnd, a.buf[i:]...)...))
	}

	// Missing dot

	if len(a.buf) > 3+len(boldTag) {
		a.buf = insertThousandSeparator(a.buf, len(a.buf)-3)
	}

	a.buf = append(a.buf, byte('<'), byte('/'), byte('b'), byte('>'))

	if addMissingDot {
		a.buf = append(a.buf, byte('.'), byte('0'), byte('0'))
	}

	return string(a.buf)
}

// insertThousandSeparator inserts "&#x202f;" (narrow no-break space) that is thinner than "&nbsp;".
// Alternative are not better: "&thinsp;" breaks the line, "&numsp;" is too wide.
func insertThousandSeparator(buf []byte, i int) []byte {
	return append(buf[:i], append(narrowNBSp, buf[i:]...)...)
}
