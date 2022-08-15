// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package thales

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/teal-finance/rainbow/pkg/provider/the-graph/thales"
	"github.com/teal-finance/rainbow/pkg/provider/zerox"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	// thegraph urls.
	urlOptimism = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
	urlPolygon  = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"

	// rpc.
	rpcOptimism = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m"
	rpcPolygon  = "https://polygon-mainnet.g.alchemy.com/v2/7MGFstWkvX-GscRyBQxehyisRlEoQWyu"

	// amm.
	ammPolygon  = "0x9b6d76B1C6140FbB0ABc9C4a348BFf4e4e8a1213"
	ammOptimism = "0x5ae7454827D83526261F3871C1029792644Ef1B1"

	// other.
	name         = "Thales"
	skip         = 0
	first        = 500
	UP     uint8 = 0
	DOWN   uint8 = 1
	amount       = 1
)

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	layer := []string{"Optimism", "Polygon"}
	var o []rainbow.Option
	for _, l := range layer {
		_, url, _, _ := LayerInfo(l)
		markets, err := QueryAllMarkets(url)
		if err != nil {
			log.Print("ERR: ", err)
			return nil, err
		}
		options, errmione := ProcessMarkets(markets, l)
		if errmione != nil {
			log.Print("ERR: ", err)
			return nil, err
		}
		o = append(o, options...)
	}

	return o, nil
}

func ProcessMarkets(markets []thales.AllMarketsMarketsMarket, layer string) ([]rainbow.Option, error) {
	spew.Dump(len(markets))

	r := make([]rainbow.Option, 0, 2*len(markets))

	for count, m := range markets {
		if m.Id == "0xa0692fa1040200ac4e4818b460055753855fd623" {
			continue
		}
		up, err := getOption(m, UP, layer)
		if err != nil {
			log.Print(count, " ERR: getOption: ", m.Id, " UP: ", err)
			spew.Dump(m)
			return nil, err
		}
		down, err := getOption(m, DOWN, layer)
		if err != nil {
			log.Print(count, " ERR: getOption: ", m.Id, " DOWN: ", err)
			spew.Dump(m)
			return nil, err
		}
		r = append(r, up, down)
		if count%10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}
	spew.Dump(len(r))
	return r, nil
}

func getOption(m thales.AllMarketsMarketsMarket, side uint8, layer string) (rainbow.Option, error) {
	// TODO change front to take care of UP/DOWN properly
	// here we do a hack where
	// DOWN == PUT
	// UP   == CALL
	binaryType := "CALL" // "DOWN"
	if side != 0 {
		binaryType = "PUT" // "UP"
	}

	expiry, err := rainbow.TimeStringConvert(m.MaturityDate)
	if err != nil {
		log.Print("ERR: ", err)
		return rainbow.Option{}, err
	}

	strikeInt := new(big.Int)
	_, err = fmt.Sscan(m.StrikePrice, strikeInt)
	if err != nil {
		log.Print("ERR: ", err)
		return rainbow.Option{}, err
	}

	binary := rainbow.Option{
		Name:          "",
		Type:          binaryType,
		Asset:         Underlying(m.CurrencyKey),
		Expiry:        expiry,
		ExchangeType:  "DEX",
		Chain:         "Ethereum",
		Layer:         "L2",
		LayerName:     layer,
		Provider:      name + "::" + layer,
		QuoteCurrency: "USD", // sUSD for optimism, usdc for polygon
		// TODO add underlying quote currency to be able to specify the token
		Bid:    nil,
		Ask:    nil,
		Strike: rainbow.ToFloat(strikeInt, rainbow.DefaultEthereumDecimals),
	}
	binary.Name = binary.OptionName()
	buy, err := getQuote(m, side, "BUY", layer)
	if err != nil {
		log.Print("ERR: ", err)
		return rainbow.Option{}, err
	}

	sell, err := getQuote(m, side, "SELL", layer)
	if err != nil {
		log.Print("ERR: ", err)
		return rainbow.Option{}, err
	}
	binary.Bid = append(binary.Bid, rainbow.Order{
		Price: sell,
		Size:  float64(amount),
	})

	binary.Ask = append(binary.Ask, rainbow.Order{
		Price: buy,
		Size:  float64(amount),
	})

	return binary, nil
}

func LayerInfo(s string) (rpc, thegraphURL, amm string, decimals int64) {
	if s == "Optimism" {
		rpc = rpcOptimism
		thegraphURL = urlOptimism
		amm = ammOptimism
		decimals = rainbow.DefaultEthereumDecimals
	} else if s == "Polygon" {
		rpc = rpcPolygon
		thegraphURL = urlPolygon
		amm = ammPolygon
		decimals = zerox.USDCDecimals
	}
	return rpc, thegraphURL, amm, decimals
}

func getQuote(m thales.AllMarketsMarketsMarket, side uint8, action, layer string) (float64, error) {
	rpc, _, amm, decimals := LayerInfo(layer)
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Print("ERR: ", err)
		return 0, err
	}

	address := common.HexToAddress(amm)
	instance, err := NewThales(address, client)
	if err != nil {
		log.Print("ERR: ", err)
		return 0, err
	}
	amountToQuote := rainbow.IntToEthereumUint256(amount, rainbow.DefaultEthereumDecimals)
	quote := new(big.Int)
	if action == "BUY" {
		quote, err = instance.BuyFromAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), side, amountToQuote)
		if err != nil {
			log.Print("ERR: ", err)
			return 0, err
		}
	} else if action == "SELL" {
		quote, err = instance.SellToAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), side, amountToQuote)
		if err != nil {
			log.Print("ERR: ", err)
			return 0, err
		}
	}

	return rainbow.ToFloat(quote, decimals), nil
}

func QueryAllLiveMarkets(url string) ([]thales.AllLiveMarketsMarket, error) {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllLive(context.TODO(), graphqlClient)
	if err != nil {
		log.Print("ERR: ", err)
		return nil, err
	}

	if resp == nil {
		log.Print("ERR: resp=nil")
		return nil, err
	}
	return resp.Markets, err
}

// TODO add err.
func QueryAllMarkets(url string) ([]thales.AllMarketsMarketsMarket, error) {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllMarkets(context.TODO(), graphqlClient, skip, first)
	if err != nil {
		log.Print("ERR thales.AllMarkets: ", err)
		return nil, err
	}
	if resp == nil {
		log.Print("ERR thales.AllMarkets: resp=nil")
		return nil, err
	}
	return resp.Markets, nil
}

func QueryMarkets(url string) []thales.MarketsMarketsMarket {
	graphqlClient := graphql.NewClient(url, nil)
	minExpiry := rainbow.TwoWeeksInThePast()
	resp, err := thales.Markets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Print("ERR thales.Markets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.Markets: resp=nil")
		return nil
	}
	return resp.Markets
}

func QueryMarket(id, url string) *thales.MarketMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.Market(context.TODO(), graphqlClient, id)
	if err != nil {
		log.Print("ERR thales.Market: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.Market: resp=nil")
		return nil
	}
	return &resp.Market
}

func QueryAllRangedMarkets(url string) []thales.AllRangedMarketsRangedMarketsRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllRangedMarkets(context.TODO(), graphqlClient, skip, first)
	if err != nil {
		log.Print("ERR thales.AllRangedMarkets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.AllRangedMarkets: resp=nil")
		return nil
	}
	return resp.RangedMarkets
}

func QueryRangedMarkets(url string) []thales.RangedMarketsRangedMarketsRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	minExpiry := rainbow.TwoWeeksInThePast()
	resp, err := thales.RangedMarkets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Print("ERR thales.RangedMarkets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.RangedMarkets: resp=nil")
		return nil
	}
	return resp.RangedMarkets
}

func QueryRangedMarket(id, url string) *thales.RangedMarketRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.RangedMarket(context.TODO(), graphqlClient, id)
	if err != nil {
		log.Print("ERR thales.RangedMarket: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.RangedMarket: resp=nil")
		return nil
	}
	return &resp.RangedMarket
}

func Underlying(s string) string {
	l := common.HexToHash(s)
	b := l.Bytes()
	b = bytes.Trim(b, "\x00")
	return string(b)
}
