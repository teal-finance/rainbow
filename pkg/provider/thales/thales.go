// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package thales

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/rainbow/pkg/provider/the-graph/thales"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets/anchor"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone("Thales")

const (
	// thegraph urls.

	urlOptimism = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-markets"
	urlPolygon  = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
	urlArbitrum = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-arbitrum"
	urlBsc      = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-bsc"

	// rpc.
	rpcOptimism = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m"
	rpcPolygon  = "https://polygon-mainnet.g.alchemy.com/v2/7MGFstWkvX-GscRyBQxehyisRlEoQWyu"
	rpcArbitrum = "https://arb-mainnet.g.alchemy.com/v2/hnBqLngSXPbAdvXHjcstEHkvWXV7RzEJ"
	rpcBsc      = "https://nd-747-272-879.p2pify.com/934be2335664b8002d732e48a40b1e43"

	// amm.
	ammPolygon  = "0x9b6d76B1C6140FbB0ABc9C4a348BFf4e4e8a1213"
	ammOptimism = "0x5ae7454827D83526261F3871C1029792644Ef1B1"
	ammArbitrum = "0x2b89275efB9509c33d9AD92A4586bdf8c4d21505"
	ammBsc      = "0x465B66A3e33088F0666dB1836652fBcF037c7319"

	// other.
	name           = "Thales"
	skip           = 0
	first          = 500
	UP       uint8 = 0
	DOWN     uint8 = 1
	amount         = 1
	baseUrl        = "https://thalesmarket.io/markets/"
	referral       = "?referralId=0xb3ac309aee5780d951082731ff2cc7f94f9488fd"
)

func LayerRPC(layer string) string {
	switch layer {
	case "Optimism":
		return rpcOptimism
	case "Polygon":
		return rpcPolygon
	case "Arbitrum":
		return rpcArbitrum
	case "Bsc":
		return rpcBsc
	}
	log.Panic("Unexpected layer", layer)
	return ""
}

func LayerURL(layer string) string {
	switch layer {
	case "Optimism":
		return urlOptimism
	case "Polygon":
		return urlPolygon
	case "Arbitrum":
		return urlArbitrum
	case "Bsc":
		return urlBsc
	}
	log.Panic("Unexpected layer", layer)
	return ""
}

func LayerAMM(layer string) string {
	switch layer {
	case "Optimism":
		return ammOptimism
	case "Polygon":
		return ammPolygon
	case "Arbitrum":
		return ammArbitrum
	case "Bsc":
		return ammBsc
	}
	log.Panic("Unexpected layer", layer)
	return ""
}

func LayerDecimals(layer string) int64 {
	switch layer {
	case "Optimism":
		return rainbow.DefaultEthereumDecimals
	case "Polygon":
		return anchor.USDCDecimals
	case "Arbitrum":
		return anchor.USDCDecimals
	case "Bsc":
		return rainbow.DefaultEthereumDecimals

	}
	log.Panic("Unexpected layer", layer)
	return 0
}

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	marketsOptimism, err := QueryAllMarkets(LayerURL("Optimism"))
	if err != nil {
		return nil, err
	}
	marketsPolygon, err := QueryAllMarkets(LayerURL("Polygon"))
	if err != nil {
		return nil, err
	}
	marketsArbitrum, err := QueryAllMarkets(LayerURL("Arbitrum"))
	if err != nil {
		return nil, err
	}
	// uncomment when BSC is ready
	/*marketsBsc, err := QueryAllMarkets(LayerURL("Bsc"))
	if err != nil {
		return nil, err
	}*/

	//options := make([]rainbow.Option, 0, 2*len(marketsOptimism)+2*len(marketsPolygon)+2*len(marketsArbitrum)+2*len(marketsBsc))
	options := make([]rainbow.Option, 0, 2*len(marketsOptimism)+2*len(marketsPolygon)+2*len(marketsArbitrum))
	err = ProcessMarkets(&options, marketsOptimism, "Optimism")
	if err != nil {
		return nil, err
	}
	err = ProcessMarkets(&options, marketsPolygon, "Polygon")

	if err != nil {
		return nil, err
	}
	err = ProcessMarkets(&options, marketsArbitrum, "Arbitrum")
	if err != nil {
		return nil, err
	}
	/*err = ProcessMarkets(&options, marketsBsc, "Bsc")
	if err != nil {
		return nil, err
	}*/

	return options, err
}

func ProcessMarkets(options *[]rainbow.Option, markets []thales.AllMarketsMarketsMarket, layer string) error {
	log.Printf("%s %v options\n", layer, len(markets))
	for i := range markets {
		// HOTFIX for bug on Polygon
		// 3 markets for BTC with very low strike
		// TODO properly understand this error "execution reverted: uint overflow from multiplication"
		// remove annoying market
		/*if markets[i].Id == "0xa0692fa1040200ac4e4818b460055753855fd623" ||
			markets[i].Id == "0x419bf5bfaf543c1a6d9db5fbd8da8fe24a05c31c" ||
			markets[i].Id == "0x08baf8b8791bb39c4f677eb4b2023665f0a46df8" ||
			markets[i].Id == "0x5a14ad0a5b9108a8c557fa68cab4c2f44005f6ac" ||
			markets[i].Id == "0xbef5d8d4e8f0e86b7c24b1b6f224020c55b65af1" ||
			markets[i].Id == "0xd0792be5111fd1ac4da4da106db53a82d967a41b" {
			continue
		}*/

		poolSize := new(big.Int)
		_, err := fmt.Sscan(markets[i].PoolSize, poolSize)
		if err != nil {
			log.Error(err)
			return err
		}

		// check if big int iszero
		// https://stackoverflow.com/questions/64257065/is-there-another-way-of-testing-if-a-big-int-is-0
		// TODO make a function maybe
		if len(poolSize.Bits()) == 0 {
			continue
		}
		up, err := getOption(&markets[i], UP, layer)
		if err != nil {
			log.Error("#", i, "getOption: "+markets[i].Id+" UP:", err)
			spew.Dump(markets[i])
			return err
		}
		down, err := getOption(&markets[i], DOWN, layer)
		if err != nil {
			log.Error("#", i, "getOption:"+markets[i].Id+" DOWN:", err)
			spew.Dump(markets[i])
			return err
		}
		*options = append(*options, up, down)
		if i%10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}

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
		log.Error(err)
		return rainbow.Option{}, err
	}

	strikeInt := new(big.Int)
	_, err = fmt.Sscan(m.StrikePrice, strikeInt)
	if err != nil {
		log.Error(err)
		return rainbow.Option{}, err
	}

	binary := rainbow.Option{
		Name:         "",
		Type:         binaryType,
		Asset:        Underlying(m.CurrencyKey),
		Expiry:       expiry,
		ExchangeType: "DEX",
		Chain:        "Ethereum",
		Layer:        "L2",
		LayerName:    layer,
		Provider:     name + "::" + layer,
		// Link: url(m.Id)
		QuoteCurrency: "USD", // sUSD for optimism, usdc for polygon/arbitrum,busd for binance
		// TODO add underlying quote currency to be able to specify the token
		Bid:    nil,
		Ask:    nil,
		Strike: rainbow.ToFloat(strikeInt, rainbow.DefaultEthereumDecimals),
	}
	binary.Name = binary.OptionName()

	// here I had to do a weird choice. basically, if the orderbook is empty on one side,
	// I obviously can't get a quote. I think I should get something like this:
	// execution reverted: SafeMath: subtraction overflow from ...
	// the real error is mainly when both fail, i.e., the whole AMM is empty or something worth logging
	//
	buy, err1 := getQuote(m, side, "BUY", layer)
	sell, err2 := getQuote(m, side, "SELL", layer)

	if err1 != nil && err2 != nil {
		log.Error(err1)
		log.Error(err2)
		return rainbow.Option{}, err1
	}

	if err2 != nil {
		binary.Bid = append(binary.Bid, rainbow.Order{
			Price: sell,
			Size:  float64(amount),
		})
	}

	if err1 != nil {
		binary.Ask = append(binary.Ask, rainbow.Order{
			Price: buy,
			Size:  float64(amount),
		})
	}

	return binary, nil
}

func getQuote(m *thales.AllMarketsMarketsMarket, side uint8, action, layer string) (float64, error) {
	rpc := LayerRPC(layer)
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Error("Thales ethclient.Dial", err)
		return 0, err
	}

	amm := LayerAMM(layer)
	address := common.HexToAddress(amm)
	instance, err := NewThales(address, client)
	if err != nil {
		log.Error("NewThales", err)
		return 0, err
	}
	amountToQuote := rainbow.IntToEthereumUint256(amount, rainbow.DefaultEthereumDecimals)
	quote := new(big.Int)
	if action == "BUY" {
		quote, err = instance.BuyFromAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), side, amountToQuote)
		if err != nil {
			log.Error("Thales BuyFromAmmQuote on ", layer, err)
			return 0, err
		}
	} else if action == "SELL" {
		quote, err = instance.SellToAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), side, amountToQuote)
		if err != nil {
			log.Error("Thales SellToAmmQuote on ", layer, err)
			return 0, err
		}
	}

	decimals := LayerDecimals(layer)
	return rainbow.ToFloat(quote, decimals), nil
}

func QueryAllLiveMarkets(url string) ([]thales.AllLiveMarketsMarket, error) {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.AllLive(context.TODO(), graphqlClient)
	if err != nil {
		log.Error("Thales AllLive", err)
		return nil, err
	}

	if resp == nil {
		log.Error("Thales resp=nil")
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
		log.Error("thales.AllRangedMarkets", err)
		return nil
	}
	if resp == nil {
		log.Error("thales.AllRangedMarkets resp=nil")
		return nil
	}
	return resp.RangedMarkets
}

func QueryRangedMarkets(url string) []thales.RangedMarketsRangedMarketsRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	minExpiry := rainbow.TwoWeeksInThePast()
	resp, err := thales.RangedMarkets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		log.Error("thales.RangedMarkets", err)
		return nil
	}
	if resp == nil {
		log.Error("thales.RangedMarkets resp=nil")
		return nil
	}
	return resp.RangedMarkets
}

func QueryRangedMarket(id, url string) *thales.RangedMarketRangedMarket {
	graphqlClient := graphql.NewClient(url, nil)
	resp, err := thales.RangedMarket(context.TODO(), graphqlClient, id)
	if err != nil {
		log.Error("thales.RangedMarket", err)
		return nil
	}
	if resp == nil {
		log.Error("thales.RangedMarket resp=nil")
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

func url(id string) string {
	return baseUrl + id + referral
}
