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

	"github.com/Khan/genqlient/graphql"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/teal-finance/rainbow/pkg/provider/the-graph/thales"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	//thegraph urls
	urlOptimism = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
	urlPolygon  = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
	//rpc
	rpcOptimism = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m"
	rpcPolygon  = "https://polygon-mainnet.g.alchemy.com/v2/7MGFstWkvX-GscRyBQxehyisRlEoQWyu"
	//amm
	ammPolygon  = "0x9b6d76B1C6140FbB0ABc9C4a348BFf4e4e8a1213"
	ammOptimism = "0x5ae7454827D83526261F3871C1029792644Ef1B1"
	name        = "Thales"
	skip        = 0
	first       = 100
	UP          = 0
	DOWN        = 1
	amount      = 1
)

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	layer := []string{"Optimism", "Polygon"}
	var o []rainbow.Option
	for _, l := range layer {

		_, url, _ := LayerInfo(l)
		markets, err := QueryAllLiveMarkets(url)
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

func ProcessMarkets(markets []thales.AllLiveMarketsMarket, layer string) ([]rainbow.Option, error) {
	spew.Dump(len(markets))

	r := make([]rainbow.Option, 0, 2*len(markets))

	for _, m := range markets {
		up, err := getOption(m, UP, layer)
		if err != nil {
			log.Print("ERR: ", err)
			return nil, err
		}
		down, err := getOption(m, DOWN, layer)
		if err != nil {
			log.Print("ERR: ", err)
			return nil, err
		}
		r = append(r, up, down)

	}
	spew.Dump(len(r))
	return r, nil

}
func getOption(m thales.AllLiveMarketsMarket, side int8, layer string) (rainbow.Option, error) {
	binaryType := "DOWN"
	if side != 0 {
		binaryType = "UP"
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
		Provider:      name,
		QuoteCurrency: "USD", // sUSD for optimism, usdc for polygon
		//TODO add underlying quote currency to be able to specify the token
		Bid:    nil,
		Ask:    nil,
		Strike: rainbow.ToFloat(strikeInt),
	}
	binary.Name = binary.OptionName()
	return binary, nil
}
func LayerInfo(s string) (rpc, thegraphURL, amm string) {
	if s == "Optimism" {
		rpc = rpcOptimism
		thegraphURL = urlOptimism
		amm = ammOptimism
	} else if s == "Polygon" {
		rpc = rpcPolygon
		thegraphURL = urlPolygon
		amm = ammPolygon
	}
	return rpc, thegraphURL, amm
}

func getQuote(m thales.AllLiveMarketsMarket, side int8, action, layer string) (float64, error) {
	rpc, _, amm := LayerInfo(layer)
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(amm)
	instance, err := NewThales(address, client)
	if err != nil {
		log.Fatal(err)
	}
	amountToQuote := rainbow.IntToEthereumFormat(amount)
	quote := new(big.Int)
	if action == "BUY" {
		quote, err = instance.BuyFromAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), UP, amountToQuote)
		if err != nil {
			log.Fatal(err)

		}
	} else if action == "SELL" {
		quote, err = instance.SellToAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), UP, amountToQuote)
		if err != nil {
			log.Fatal(err)
		}
	}

	return rainbow.ToFloat(quote), nil
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

//TODO add err
func QueryAllMarkets(url string) []thales.AllMarketsMarketsMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllMarkets(context.TODO(), graphqlClient, skip, first)
	if err != nil {
		log.Print("ERR thales.AllMarkets: ", err)
		return nil
	}
	if resp == nil {
		log.Print("ERR thales.AllMarkets: resp=nil")
		return nil
	}
	return resp.Markets
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
