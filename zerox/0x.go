// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package zerox

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/teal-finance/rainbow"
)

const (
	USDC            = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	WETH            = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
	WBTC            = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
	USDCDecimals    = 6
	OTokensDecimals = 8
	WETHDecimals    = 18
	WBTCDEcimals    = 8
)

func GetMarkets(coin string) []Opyn {

	return Opynmarket(coin)
}

func GetOrderBook(instruments []Opyn, provider string) ([]rainbow.Options, error) {
	baseURL := "https://api.0x.org/orderbook/v1?quoteToken="
	opts := "&baseToken=" + USDC

	clt := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	orderBook := []rainbow.Options{}

	for _, i := range instruments {
		resp, err := clt.Get(baseURL + i.ID + opts)
		if err != nil {
			return []rainbow.Options{}, err
		}

		defer resp.Body.Close()

		result := ZeroxOrderBook{}

		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return []rainbow.Options{}, fmt.Errorf(" order book : %w", err)
		}

		o := rainbow.Options{
			Name:                i.Name,
			Type:                i.Type,
			Asset:               i.Asset,
			ExpirationTimestamp: i.ExpirationTimestamp,
			Strike:              i.Strike,
			ExchangeType:        "DEX",
			Chain:               "Ethereum",
			Layer:               "L1",
			Provider:            provider,
			Offers:              nil,
		}
		b, err := BidsAsksToOffers(result.Bids.Records, "BUY", OpynQuoteCurrency)
		if err != nil {
			return []rainbow.Options{}, err
		}
		a, err := BidsAsksToOffers(result.Asks.Records, "SELL", OpynQuoteCurrency)
		if err != nil {
			return []rainbow.Options{}, err
		}

		o.Offers = append(o.Offers, b...)
		o.Offers = append(o.Offers, a...)
		// fmt.Println("o ", o)

		orderBook = append(orderBook, o)
	}

	return orderBook, nil
}

type ZeroxOrderBook struct {
	Bids struct {
		Total   int     `json:"total"`
		Page    int     `json:"page"`
		PerPage int     `json:"perPage"`
		Records Records `json:"records"`
	} `json:"bids"`
	Asks struct {
		Total   int     `json:"total"`
		Page    int     `json:"page"`
		PerPage int     `json:"perPage"`
		Records Records `json:"records"`
	} `json:"asks"`
}

type Records []struct {
	Order struct {
		Signature struct {
			SignatureType int    `json:"signatureType"`
			R             string `json:"r"`
			S             string `json:"s"`
			V             int    `json:"v"`
		} `json:"signature"`
		Sender              string `json:"sender"`
		Maker               string `json:"maker"`
		Taker               string `json:"taker"`
		TakerTokenFeeAmount string `json:"takerTokenFeeAmount"`
		MakerAmount         string `json:"makerAmount"`
		TakerAmount         string `json:"takerAmount"`
		MakerToken          string `json:"makerToken"`
		TakerToken          string `json:"takerToken"`
		Salt                string `json:"salt"`
		VerifyingContract   string `json:"verifyingContract"`
		FeeRecipient        string `json:"feeRecipient"`
		Expiry              string `json:"expiry"`
		ChainID             int    `json:"chainId"`
		Pool                string `json:"pool"`
	} `json:"order"`
	MetaData struct {
		OrderHash                    string    `json:"orderHash"`
		RemainingFillableTakerAmount string    `json:"remainingFillableTakerAmount"`
		CreatedAt                    time.Time `json:"createdAt"`
	} `json:"metaData"`
}

func BidsAsksToOffers(records Records, side, quote string) ([]rainbow.Offer, error) {
	offers := []rainbow.Offer{}
	for _, r := range records {
		takerAmount, err := strconv.ParseFloat(r.Order.TakerAmount, 64)
		if err != nil {
			return []rainbow.Offer{}, err
		}
		makerAmount, err := strconv.ParseFloat(r.Order.MakerAmount, 64)
		if err != nil {
			return []rainbow.Offer{}, err
		}
		offers = append(offers, rainbow.Offer{Side: side, Price: makerAmount / takerAmount, Quantity: takerAmount * math.Pow(10, -float64(USDCDecimal)), QuoteCurrency: quote})

	}
	return offers, nil

}
