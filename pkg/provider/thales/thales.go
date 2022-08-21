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
	markets1, err1 := QueryAllMarkets(LayerURL("Optimism"))
	markets2, err2 := QueryAllMarkets(LayerURL("Polygon"))

	options := make([]rainbow.Option, 0, 2*len(markets1)+2*len(markets2))
	if err1 != nil {
		err1 = ProcessMarkets(options, markets1, "Optimism")
	}
	if err2 != nil {
		err2 = ProcessMarkets(options, markets2, "Polygon")
	}

	if err1 != nil {
		err2 = err1 // return err1 if any, else err2
	}

	return options, err2
}

func ProcessMarkets(options []rainbow.Option, markets []thales.AllMarketsMarketsMarket, layer string) error {
	spew.Dump(len(markets))

	for i := range markets {
		// HOTFIX for bug on Polygon
		// 3 markets for BTC with very low strike
		// TODO properly understand this "ERR: execution reverted: uint overflow from multiplication"
		// remove annoying market
		if markets[i].Id == "0xa0692fa1040200ac4e4818b460055753855fd623" ||
			markets[i].Id == "0x419bf5bfaf543c1a6d9db5fbd8da8fe24a05c31c" ||
			markets[i].Id == "0x08baf8b8791bb39c4f677eb4b2023665f0a46df8" ||
			markets[i].Id == "0x5a14ad0a5b9108a8c557fa68cab4c2f44005f6ac" ||
			markets[i].Id == "0xbef5d8d4e8f0e86b7c24b1b6f224020c55b65af1" ||
			markets[i].Id == "0xd0792be5111fd1ac4da4da106db53a82d967a41b" {
			continue
		}
		up, err := getOption(&markets[i], UP, layer)
		if err != nil {
			log.Print("ERR #", i, " getOption: ", markets[i].Id, " UP: ", err)
			spew.Dump(markets[i])
			return err
		}
		down, err := getOption(&markets[i], DOWN, layer)
		if err != nil {
			log.Print("ERR #", i, " getOption: ", markets[i].Id, " DOWN: ", err)
			spew.Dump(markets[i])
			return err
		}
		options = append(options, up, down)
		if i%10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}
	spew.Dump(len(options))
	return nil
}

func getOption(m *thales.AllMarketsMarketsMarket, side uint8, layer string) (rainbow.Option, error) {
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
		log.Print("ERR Thales ", err)
		return rainbow.Option{}, err
	}

	strikeInt := new(big.Int)
	_, err = fmt.Sscan(m.StrikePrice, strikeInt)
	if err != nil {
		log.Print("ERR Thales ", err)
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
		log.Print("ERR Thales ", err)
		return rainbow.Option{}, err
	}

	sell, err := getQuote(m, side, "SELL", layer)
	if err != nil {
		log.Print("ERR Thales ", err)
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

func LayerInfo(layer string) (rpc, thegraphURL, amm string, decimals int64) {
	if layer == "Optimism" {
		rpc = rpcOptimism
		thegraphURL = urlOptimism
		amm = ammOptimism
		decimals = rainbow.DefaultEthereumDecimals
	} else if layer == "Polygon" {
		rpc = rpcPolygon
		thegraphURL = urlPolygon
		amm = ammPolygon
		decimals = zerox.USDCDecimals
	}
	return rpc, thegraphURL, amm, decimals
}

func LayerURL(layer string) string {
	_, thegraphURL, _, _ := LayerInfo(layer)
	return thegraphURL
}

func getQuote(m *thales.AllMarketsMarketsMarket, side uint8, action, layer string) (float64, error) {
	rpc, _, amm, decimals := LayerInfo(layer)
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Print("ERR Thales ", err)
		return 0, err
	}

	address := common.HexToAddress(amm)
	instance, err := NewThales(address, client)
	if err != nil {
		log.Print("ERR Thales ", err)
		return 0, err
	}
	amountToQuote := rainbow.IntToEthereumUint256(amount, rainbow.DefaultEthereumDecimals)
	quote := new(big.Int)
	if action == "BUY" {
		quote, err = instance.BuyFromAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), side, amountToQuote)
		if err != nil {
			log.Print("ERR Thales ", err)
			return 0, err
		}
	} else if action == "SELL" {
		quote, err = instance.SellToAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), side, amountToQuote)
		if err != nil {
			log.Print("ERR Thales ", err)
			return 0, err
		}
	}

	return rainbow.ToFloat(quote, decimals), nil
}

func QueryAllLiveMarkets(url string) ([]thales.AllLiveMarketsMarket, error) {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllLive(context.TODO(), graphqlClient)
	if err != nil {
		log.Print("ERR Thales ", err)
		return nil, err
	}

	if resp == nil {
		log.Print("ERR Thales resp=nil")
		return nil, err
	}
	return resp.Markets, err
}

// TODO add err.
func QueryAllMarkets(url string) ([]thales.AllMarketsMarketsMarket, error) {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllMarkets(context.TODO(), graphqlClient, skip, first)
	if err != nil {
		return nil, fmt.Errorf("AllMarkets: %w", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("AllMarkets resp=nil")
	}
	return resp.Markets, nil
}

func QueryMarkets(url string) ([]thales.MarketsMarketsMarket, error) {
	graphqlClient := graphql.NewClient(url, nil)
	minExpiry := rainbow.TwoWeeksInThePast()
	resp, err := thales.Markets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		return nil, fmt.Errorf("Markets: %w", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("Markets: resp=nil")
	}
	return resp.Markets, nil
}

func QueryMarket(id, url string) (*thales.MarketMarket, error) {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.Market(context.TODO(), graphqlClient, id)
	if err != nil {
		return nil, fmt.Errorf("Market: %w", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("Market: resp=nil")
	}
	return &resp.Market, nil
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
