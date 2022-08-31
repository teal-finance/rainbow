// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package zerox

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/garcon"
	"github.com/teal-finance/rainbow/pkg/provider/the-graph/opyn"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone("0x")

const (
	USDC            = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	WETH            = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
	WBTC            = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
	USDCDecimals    = 6
	OTokensDecimals = 8
	WETHDecimals    = 18
	WBTCDEcimals    = 8
)

func extract(o *opyn.OptionsOtokensOToken) (optionType, expiry string, strike float64) {
	optionType = "CALL"
	if o.IsPut {
		optionType = "PUT"
	}

	expiry, err := rainbow.TimeStringConvert(o.ExpiryTimestamp)
	if err != nil {
		log.Printf("WRN Opyn ExpiryTimestamp: %v from %+v", err, o)
		expiry = ""
	}

	// thought the USDCdecimals were correct but apparently not (whatever)
	strike, err = convertFromSolidity(o.StrikePrice, OTokensDecimals)
	if err != nil {
		log.Printf("WRN Opyn Strike: %v from %+v", err, o)
		strike = 0
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
	defer resp.Body.Close()

	var result Quote
	if err = garcon.DecodeJSONResponse(resp, &result); err != nil {
		return Quote{}, fmt.Errorf("quote from Ox: %w", err)
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
			log.Printf("INF Opyn success #%v n=%v bad=%v sleep=%v (%+v)",
				sr.ok, n, sr.maxBad, newSleep, newSleep-sr.sleep)
		}

		sr.sleep = newSleep
		sr.maxBad -= sr.maxBad / 128
	}
}

func (sr *stubbornRequester) failure(err error) {
	newSleep := sr.maxBad / 2
	sr.ko++

	log.Printf("WRN Opyn failure #%v bad=%v sleep=%v + %v %v",
		sr.ko, sr.maxBad, sr.sleep, newSleep, garcon.Sanitize(err.Error()))

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
	log.Printf("INF Opyn stats: bad=%v sleep=%v ok=%v ko=%v",
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

func normalize(instruments []opyn.OptionsOtokensOToken, provider string, amount float64) ([]rainbow.Option, error) {
	options := make([]rainbow.Option, 0, len(instruments))

	requester := defaultStubbornRequester

	for i := range instruments {
		instr := &instruments[i]
		optionType, expiry, strike := extract(instr)

		var decimals int
		var quoteCcy string
		if provider == "Opyn" {
			decimals = OTokensDecimals
			quoteCcy = instr.StrikeAsset.Symbol
		}

		o := rainbow.Option{
			Name:          instr.Name,
			Type:          optionType,
			Asset:         instr.UnderlyingAsset.Symbol,
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

		bid, err := requester.getQuote("BUY", instr.Id, USDC, amount, decimals)
		if err != nil {
			log.Print("WRN Opyn getQuote BUY ", err)
			return nil, err
		}

		if bid.Price != "" {
			price, e := strconv.ParseFloat(bid.Price, 64)
			if e != nil {
				log.Print("WRN Opyn price ", e)
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

		ask, err := requester.getQuote("SELL", USDC, instr.Id, amount, decimals)
		if err != nil {
			log.Print("INF Opyn getQuote SELL ", err)
			return nil, err
		}

		if ask.Price != "" {
			price, err := strconv.ParseFloat(ask.Price, 64)
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

	requester.logStats()

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
