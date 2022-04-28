// Copyright (c) 2021 Teal.Finance
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
	"strings"
	"time"

	"github.com/teal-finance/rainbow/pkg/rainbow"
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

func extract(i getOptionsOtokensOToken) (optionType, expiry string, strike float64) {
	optionType = "CALL"
	if i.IsPut {
		optionType = "PUT"
	}

	seconds, err := strconv.ParseInt(i.ExpiryTimestamp, 10, 0)
	if err == nil {
		expiryTime := time.Unix(seconds, 0).UTC()
		expiry = expiryTime.Format("2006-01-02 15:04:05")
	} else {
		log.Printf("WARN Expiry: %v from %+v", err, i)
	}

	// thought the USDCdecimals were correct but apparently not (whatever)
	strike, err = convertFromSolidity(i.StrikePrice, OTokensDecimals)
	if err != nil {
		log.Printf("WARN Strike: %v from %+v", err, i) // TODO fail better
	}

	return optionType, expiry, strike
}

// getQuote you can get the quote in the side you want with 0x
// but for us we will focus on always buying or selling a certain amount of options
// so Asks/SELL side (so you send a buy inquiry): sellToken=usdc, buyToken=option address
// so Bids/BUY side (so you send a sell inquiry): sellToken=option address, buyToken=usdc.
func getQuote(side, sellToken, buyToken string, amount float64, decimals int) (Quote, error) {
	url := "https://api.0x.org/swap/v1/quote?sellToken=" + sellToken + "&buyToken=" + buyToken
	if side == "SELL" {
		url += "&buyAmount="
	} else {
		url += "&sellAmount="
	}

	url += convertToSolidity(amount, decimals)

	resp, err := http.Get(url)
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

type stubbornRequester struct {
	sleep  time.Duration
	maxBad time.Duration
	ok     int
	ko     int
}

var defaultStubbornRequester = stubbornRequester{
	sleep:  500 * time.Millisecond,
	maxBad: 250 * time.Millisecond,
	ok:     0,
	ko:     0,
}

func (sr *stubbornRequester) success(n int) {
	sr.ok++

	if sr.ok > 1 {
		var newSleep time.Duration
		if sr.sleep > sr.maxBad {
			newSleep = sr.sleep - (sr.sleep-sr.maxBad)/32
		} else {
			newSleep = sr.maxBad
		}

		if n > 0 {
			log.Printf("Success #%v n=%v bad=%v sleep=%v (%+v)",
				sr.ok, n, sr.maxBad, newSleep, newSleep-sr.sleep)
		}

		sr.sleep = newSleep
		sr.maxBad -= sr.maxBad / 128
	}
}

func (sr *stubbornRequester) failure(err error) {
	newSleep := sr.maxBad / 2
	sr.ko++

	log.Printf("Failure #%v bad=%v sleep=%v + %v %v",
		sr.ko, sr.maxBad, sr.sleep, newSleep,
		strings.ReplaceAll(strings.ReplaceAll(err.Error(), "\n", ""), "\r", ""))

	sr.maxBad = sr.sleep + newSleep
	sr.sleep = newSleep
}

func (sr *stubbornRequester) getQuote(side, sellToken, buyToken string, amount float64, decimals int) (q Quote, err error) {
	for n := 0; n < 22; n++ {
		if sr.ok+sr.ko > 0 {
			time.Sleep(sr.sleep)
		}

		q, err = getQuote(side, sellToken, buyToken, amount, decimals)

		if err == nil {
			sr.success(n)
			break
		}

		sr.failure(err)
	}

	return q, err
}

func (sr *stubbornRequester) logStats() {
	log.Printf("stats: bad=%v sleep=%v ok=%v ko=%v",
		sr.maxBad, sr.sleep, sr.ok, sr.ko)
}

type Orders struct {
	Records []Record `json:"records"`
	Total   int      `json:"total"`
	Page    int      `json:"page"`
	PerPage int      `json:"perPage"`
}

type OrderBook struct {
	Bids Orders `json:"bids"`
	Asks Orders `json:"asks"`
}

type Record struct {
	MetaData struct {
		CreatedAt                    time.Time `json:"createdAt"`
		OrderHash                    string    `json:"orderHash"`
		RemainingFillableTakerAmount string    `json:"remainingFillableTakerAmount"`
	} `json:"metaData"`
	Order struct {
		Pool                string `json:"pool"`
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
		Signature           struct {
			R             string `json:"r"`
			S             string `json:"s"`
			SignatureType int    `json:"signatureType"`
			V             int    `json:"v"`
		} `json:"signature"`
		ChainID int `json:"chainId"`
	} `json:"order"`
}

func normalize(instruments []getOptionsOtokensOToken, provider string, amount float64) ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	sr := defaultStubbornRequester

	for _, i := range instruments {
		optionType, expiry, strike := extract(i)

		var decimals int
		var quoteCcy string
		if provider == "Opyn" {
			decimals = OTokensDecimals
			quoteCcy = i.StrikeAsset.Symbol
		}

		o := rainbow.Option{
			Name:          i.Name,
			Type:          optionType,
			Asset:         i.UnderlyingAsset.Symbol,
			Expiry:        expiry,
			Strike:        strike,
			ExchangeType:  "DEX",
			Chain:         "Ethereum",
			Layer:         "L1",
			Provider:      provider,
			QuoteCurrency: quoteCcy,
			Bid:           nil,
			Ask:           nil,
		}

		b, err := sr.getQuote("BUY", i.Id, USDC, amount, decimals)
		if err != nil {
			log.Print("getQuote BUY ", err)
			return nil, err
		}

		if b.Price != "" {
			price, err := strconv.ParseFloat(b.Price, 64)
			if err != nil {
				log.Print("WARN price ", err)
				continue
			}

			o.Bid = append(o.Bid, rainbow.Order{
				Price: price,
				Size:  amount,
			})
		} else {
			o.Bid = append(o.Bid, rainbow.Order{
				Price: 0.0,
				Size:  0.0,
			})
		}

		a, err := sr.getQuote("SELL", USDC, i.Id, amount, decimals)
		if err != nil {
			log.Print("getQuote SELL ", err)
			return nil, err
		}

		if a.Price != "" {
			price, err := strconv.ParseFloat(a.Price, 64)
			if err != nil {
				return nil, err
			}

			o.Ask = append(o.Ask, rainbow.Order{
				Price: price,
				Size:  amount,
			})
		} else {
			o.Ask = append(o.Ask, rainbow.Order{
				Price: 0.0,
				Size:  0.0,
			})
		}

		options = append(options, o)
	}

	sr.logStats()

	return options, nil
}

func convertToSolidity(value float64, decimals int) string {
	return strconv.Itoa(int(value * math.Pow10(decimals)))
}

func convertFromSolidity(s string, decimals int) (float64, error) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}

	value *= math.Pow10(-decimals)
	return value, nil
}

type Quote struct {
	BuyTokenToEthRate  string `json:"buyTokenToEthRate"`
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
	SellTokenToEthRate string `json:"sellTokenToEthRate"`
	AllowanceTarget    string `json:"allowanceTarget"`
	Orders             []struct {
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
				Pool                string `json:"pool"`
				ChainID             int    `json:"chainId"`
			} `json:"order"`
			FillableTakerAmount    string `json:"fillableTakerAmount"`
			FillableMakerAmount    string `json:"fillableMakerAmount"`
			FillableTakerFeeAmount string `json:"fillableTakerFeeAmount"`
			Signature              struct {
				R             string `json:"r"`
				S             string `json:"s"`
				SignatureType int    `json:"signatureType"`
				V             int    `json:"v"`
			} `json:"signature"`
			Type int `json:"type"`
		} `json:"fillData"`
		Type int `json:"type"`
	} `json:"orders"`
	Sources []struct {
		Name       string `json:"name"`
		Proportion string `json:"proportion"`
	} `json:"sources"`
	ChainID int `json:"chainId"`
}
