// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package zerox

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

var decimalsToPower10 = map[int]float64{
	1:  1,
	2:  10,
	3:  100,
	4:  1000,
	5:  10000,
	6:  100000,
	7:  1000000,
	8:  10000000,
	9:  100000000,
	10: 1000000000,
	11: 10000000000,
	12: 100000000000,
}

func extract(i getOptionsOtokensOToken) (tipe, expiry string, strike float64) {
	tipe = "CALL"
	if i.IsPut {
		tipe = "PUT"
	}

	seconds, err := strconv.ParseInt(i.ExpiryTimestamp, 10, 0)
	expiryTime := time.Unix(seconds, 0).UTC()
	expiry = expiryTime.Format("2006-01-02 15:04:05")
	if err != nil {
		log.Printf("WARN Expiry: %v from %+v", err, i)
		expiry = ""
	}

	strikeInt, err := strconv.ParseInt(i.StrikePrice, 10, 0)
	if err != nil {
		log.Printf("WARN Strike: %v from %+v", err, i)
	} else {
		denominateur, ok := decimalsToPower10[i.StrikeAsset.Decimals]
		if ok {
			strike = float64(strikeInt) / denominateur
		}
	}

	return
}

func GetOrderBook(instruments []getOptionsOtokensOToken, provider string) ([]rainbow.Options, error) {
	baseURL := "https://api.0x.org/orderbook/v1?quoteToken="
	opts := "&baseToken=" + USDC

	orderBook := []rainbow.Options{}

	for _, i := range instruments {
		resp, err := http.Get(baseURL + i.Id + opts)
		if err != nil {
			log.Printf("WARN Get base=%v id=%v opts=%v: %v", baseURL, i.Id, opts, err)
			continue
		}

		j, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			log.Print("WARN cannot ReadAll: ", err)
			continue
		}

		var result OrderBook
		if err = json.Unmarshal(j, &result); err != nil {
			log.Printf("WARN %v when json.Unmarshal: %v", err, string(j))
			continue
		}

		tipe, expiry, strike := extract(i)

		o := rainbow.Options{
			Name:         i.Name,
			Type:         tipe,
			Asset:        i.CollateralAsset.Symbol,
			Expiry:       expiry,
			Strike:       strike,
			ExchangeType: "DEX",
			Chain:        "Ethereum",
			Layer:        "L1",
			Provider:     provider,
			Offers:       nil,
		}

		b, err := BidsAsksToOffers(result.Bids.Records, "BUY", i.UnderlyingAsset.Symbol)
		if err != nil {
			return []rainbow.Options{}, err
		}

		a, err := BidsAsksToOffers(result.Asks.Records, "SELL", i.UnderlyingAsset.Symbol)
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

type OrderBook struct {
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

// The result is false because I don't properly take into account the decimals
// Use GetQuote instead.
func BidsAsksToOffers(records Records, side, quote string) ([]rainbow.Offer, error) {
	offers := make([]rainbow.Offer, 0, len(records))

	for _, r := range records {
		takerAmount, err := strconv.ParseFloat(r.Order.TakerAmount, 64)
		if err != nil {
			return []rainbow.Offer{}, err
		}

		makerAmount, err := strconv.ParseFloat(r.Order.MakerAmount, 64)
		if err != nil {
			return []rainbow.Offer{}, err
		}

		offers = append(
			offers,
			rainbow.Offer{
				Side:          side,
				Price:         makerAmount / takerAmount,
				Quantity:      takerAmount * math.Pow(10, -float64(OTokensDecimals)),
				QuoteCurrency: quote,
			})
	}

	return offers, nil
}

// GetQuote you can get the quote in the side you want with 0x
// but for us we will focus on always buying or selling a certain amount of options
// so Asks/SELL side (so you send a buy inquiry): sellToken=usdc, buyToken=option address
// so Bids/BUY side (so you send a sell inquiry): sellToken=option address, buyToken=usdc.
func GetQuote(side, sellToken, buyToken string, amount float64, decimals int) (Quote, error) {
	baseURL := "https://api.0x.org/swap/v1/quote?sellToken=" + sellToken + "&buyToken=" + buyToken
	if side == "SELL" {
		baseURL += "&buyAmount="
	} else {
		baseURL += "&sellAmount="
	}
	baseURL += ConvertToSolidity(amount, decimals)

	fmt.Println("swap url ", baseURL)

	resp, err := http.Get(baseURL)
	if err != nil {
		return Quote{}, err
	}

	j, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return Quote{}, fmt.Errorf("cannot ReadAll quote from Ox: %w", err)
	}

	var result Quote
	if err = json.Unmarshal(j, &result); err != nil {
		return Quote{}, fmt.Errorf("cannot json.Unmarshal quote from Ox: %w in: %v", err, string(j))
	}

	return result, nil
}

func GetAggregatedOrderBook(instruments []getOptionsOtokensOToken, provider string, amount float64) ([]rainbow.Options, error) {
	orderBook := []rainbow.Options{}

	for _, i := range instruments {
		tipe, expiry, strike := extract(i)

		var decimals int
		var quote string
		if provider == "Opyn" {
			decimals = OTokensDecimals
			quote = i.UnderlyingAsset.Symbol
		}

		o := rainbow.Options{
			Name:         i.Name,
			Type:         tipe,
			Asset:        i.CollateralAsset.Symbol,
			Expiry:       expiry,
			Strike:       strike,
			ExchangeType: "DEX",
			Chain:        "Ethereum",
			Layer:        "L1",
			Provider:     provider,
			Offers:       nil,
		}

		b, err := GetQuote("BUY", i.Id, USDC, amount, decimals)
		if err != nil {
			return []rainbow.Options{}, err
		}

		if b.Price != "" {
			price, err := strconv.ParseFloat(b.Price, 64)
			if err != nil {
				fmt.Println("price ")
				return []rainbow.Options{}, err
			}

			o.Offers = append(o.Offers, rainbow.Offer{
				Side:          "BUY",
				Price:         price,
				Quantity:      amount,
				QuoteCurrency: quote,
			})
		} else {
			o.Offers = append(o.Offers, rainbow.Offer{
				Side:          "BUY",
				Price:         0.0,
				Quantity:      0.0,
				QuoteCurrency: quote,
			})
		}

		a, err := GetQuote("SELL", USDC, i.Id, amount, decimals)
		if err != nil {
			return []rainbow.Options{}, err
		}

		if a.Price != "" {
			price, err := strconv.ParseFloat(a.Price, 64)
			if err != nil {
				return []rainbow.Options{}, err
			}

			o.Offers = append(o.Offers, rainbow.Offer{
				Side:          "SELL",
				Price:         price,
				Quantity:      amount,
				QuoteCurrency: quote,
			})
		} else {
			o.Offers = append(o.Offers, rainbow.Offer{
				Side:          "SELL",
				Price:         0.0,
				Quantity:      0.0,
				QuoteCurrency: quote,
			})
		}

		orderBook = append(orderBook, o)
	}

	return orderBook, nil
}

func ConvertToSolidity(value float64, decimals int) string {
	return strconv.Itoa(int(value * math.Pow10(decimals)))
}

func ConvertFromSolidity(s string, decimals int) (float64, error) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}

	value *= math.Pow10(-decimals)

	return value, nil
}

type Quote struct {
	ChainID            int    `json:"chainId"`
	Price              string `json:"price"`
	GuaranteedPrice    string `json:"guaranteedPrice"`
	To                 string `json:"to"`
	Data               string `json:"data"`
	Value              string `json:"value"`
	Gas                string `json:"gas"`
	EstimatedGas       string `json:"estimatedGas"`
	GasPrice           string `json:"gasPrice"`
	ProtocolFee        string `json:"protocolFee"`
	MinimumProtocolFee string `json:"minimumProtocolFee"`
	BuyTokenAddress    string `json:"buyTokenAddress"`
	SellTokenAddress   string `json:"sellTokenAddress"`
	BuyAmount          string `json:"buyAmount"`
	SellAmount         string `json:"sellAmount"`
	Sources            []struct {
		Name       string `json:"name"`
		Proportion string `json:"proportion"`
	} `json:"sources"`
	Orders []struct {
		Type        int    `json:"type"`
		Source      string `json:"source"`
		MakerToken  string `json:"makerToken"`
		TakerToken  string `json:"takerToken"`
		MakerAmount string `json:"makerAmount"`
		TakerAmount string `json:"takerAmount"`
		FillData    struct {
			Order struct {
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
			Signature struct {
				SignatureType int    `json:"signatureType"`
				R             string `json:"r"`
				S             string `json:"s"`
				V             int    `json:"v"`
			} `json:"signature"`
			Type                   int    `json:"type"`
			FillableTakerAmount    string `json:"fillableTakerAmount"`
			FillableMakerAmount    string `json:"fillableMakerAmount"`
			FillableTakerFeeAmount string `json:"fillableTakerFeeAmount"`
		} `json:"fillData"`
	} `json:"orders"`
	AllowanceTarget    string `json:"allowanceTarget"`
	SellTokenToEthRate string `json:"sellTokenToEthRate"`
	BuyTokenToEthRate  string `json:"buyTokenToEthRate"`
}
