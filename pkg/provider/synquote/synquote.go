package synquote

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/rainbow/pkg/provider/deribit"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const synquoteApi = "https://api.synquote.com/listings-with-quotes?network=polygon&market="
const marketURL = "https://app.synquote.com/trade/"

var log = emo.NewZone("Synquote")

type Provider struct{}

func (Provider) Name() string {
	return "Synquote"
}

func (p Provider) Options() ([]rainbow.Option, error) {
	markets := []string{"BTC", "ETH"} //, "MATIC", "LINK"}// TO ADD when more liquid
	var quotes []Quote
	for _, m := range markets {
		url := synquoteApi + m
		log.Info(url)
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("query market %s: %w", m, err)
		}
		defer resp.Body.Close()

		var result struct {
			Result []Quote `json:"result"`
		}
		body, _ := io.ReadAll(resp.Body)

		if err = json.Unmarshal(body, &result.Result); err != nil {
			return nil, fmt.Errorf("query market %s decode: %w", m, err)
		}
		quotes = append(quotes, result.Result...)

	}

	return extractOptions(quotes)
}

func extractOptions(quotes []Quote) ([]rainbow.Option, error) {
	var options []rainbow.Option
	//spew.Dump(quotes)
	for _, q := range quotes {
		if !q.check() {
			continue
		}
		asset, optionType, expiry, strike, bidPrice, bidSize, askPrice, askSize := q.Split()
		options = append(options, rainbow.Option{
			Name:          "Synquote-" + q.ID,
			Type:          optionType,
			Asset:         asset,
			Expiry:        expiry,
			Strike:        strike,
			ExchangeType:  "DEX",
			Chain:         "Ethereum",
			Layer:         "Polygon",
			Provider:      "Synquote",
			QuoteCurrency: "USD",
			//TODO add underlyingquote currency
			// e.g. here the quote are in USD
			// but the token used is USDC
			Bid: []rainbow.Order{{
				Price: bidPrice,
				Size:  bidSize,
			},
			},
			Ask: []rainbow.Order{{
				Price: askPrice,
				Size:  askSize,
			},
			},
			URL: marketURL + asset,
		})
	}

	return options, nil
}

type Quote struct {
	ID  string  `json:"symbol"`
	Bid float64 `json:"bid"`
	Ask float64 `json:"ask"`
}

func (q Quote) Split() (Asset, optionType, expiry string, strike, bidPrice, bidSize, askPrice, askSize float64) {
	s := strings.Split(q.ID, "-")
	if len(s) != 4 {
		log.Warn("Option with bad name format", q.ID)
	}
	optionType = "PUT"
	if s[3] == "C" {
		optionType = "CALL"
	}
	bidPrice, bidSize = priceSize(q.Bid)
	askPrice, askSize = priceSize(q.Ask)
	strike, _ = strconv.ParseFloat(s[2], 64)

	return s[0], optionType, Expiry(s[1]), strike, bidPrice, bidSize, askPrice, askSize
}

// the API return the quote for size=1
// when there is no order, the API sends "-1"
func priceSize(bidAsk float64) (price, size float64) {
	if bidAsk == -1 {
		return 0, 0
	}
	return bidAsk, 1
}

// Because I couldn't see a way to do that directly, here is my date parsor
// Expiry parse the date
// Also expiry hour is the same as deribit (like everyon except delta)
// string in format 27JAN23
// Note: failure is silent here because I'm lazy
func Expiry(s string) string {
	length := len(s)
	y := s[length-2 : length]
	m := s[length-5 : length-2]
	d := s[0 : length-5]
	//spew.Dump(s, d, m, y)
	//dummyDate:=
	yy, _ := strconv.Atoi(y)
	dd, _ := strconv.Atoi(d)
	t := time.Date(2000+yy, month(m), dd, deribit.Hour, 0, 0, 0, time.UTC)

	return t.Format("2006-01-02 15:04:05")

}

// ¯\_(ツ)_/¯
func month(m string) time.Month {
	switch m {
	case "JAN":
		return time.January
	case "FEB":
		return time.February
	case "MAR":
		return time.March
	case "APR":
		return time.April
	case "MAY":
		return time.May
	case "JUN":
		return time.June
	case "JUL":
		return time.July
	case "AUG":
		return time.August
	case "SEP":
		return time.September
	case "OCT":
		return time.October
	case "NOV":
		return time.November
	case "DEC":
		return time.December
	}

	//if there is something weird, hopefully February is a weirsd month enough to notice lol
	return time.February

}

// quick restrictions until OK from synquote team
func (q Quote) check() bool {

	a := strings.Contains(q.ID, "NOV")
	return a
}
