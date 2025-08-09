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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/rainbow/pkg/provider/the-graph/thales"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets/anchor"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone(name)

const (
	// thegraph urls.

	urlOptimism = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-markets"
	urlPolygon  = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
	urlArbitrum = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-arbitrum"
	// urlBsc      = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-bsc"

	// rpc.
	// TODO use teal on Optimism
	// TODO use perso on Arbitrum and Polygon
	// TODO do that on the 1st Nov
	// Sorry alchemy but we r too poor.

	// teal.
	rpcOptimism = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m"
	// perso
	// rpcOptimism = "https://opt-mainnet.g.alchemy.com/v2/uksZH_SjXAaBnIw95hZcBoCWGCXs9VXI"
	// teal
	// rpcPolygon = "https://polygon-mainnet.g.alchemy.com/v2/7MGFstWkvX-GscRyBQxehyisRlEoQWyu"
	// perso.
	rpcPolygon = "https://polygon-mainnet.g.alchemy.com/v2/uQ-knqUJnSNM61nlSnOxpGfx9cqPPfos"
	// teal
	// rpcArbitrum = "https://arb-mainnet.g.alchemy.com/v2/hnBqLngSXPbAdvXHjcstEHkvWXV7RzEJ"
	// perso.
	rpcArbitrum = "https://arb-mainnet.g.alchemy.com/v2/4TQ_6stSP__V97XUQQC07AV23f_XOemr"

	// rpcBsc      = "https://bsc-dataseed1.ninicoin.io/" // free bscscan rpc
	// binance https://bsc-dataseed1.binance.org/

	// amm.
	ammPolygon  = "0xd52B865584c25FEBfcB676B9A87F32683356A063"
	ammOptimism = "0x278B5A44397c9D8E52743fEdec263c4760dc1A1A"
	ammArbitrum = "0x2b89275efB9509c33d9AD92A4586bdf8c4d21505"

	// other.
	name           = "Thales"
	skip           = 0
	first          = 500
	UP       uint8 = 0
	DOWN     uint8 = 1
	amount         = 1
	baseURL        = "https://thalesmarket.io/markets/"
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
	}
	log.Panic("Unexpected layer", layer)
	return 0
}

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	var err error

	marketsOptimism, err := queryAllMarkets("Optimism")

	marketsPolygon, e := queryAllMarkets("Polygon")
	if e != nil {
		if err == nil {
			err = e
		} else {
			err = fmt.Errorf("%w, %w", err, e)
		}
	}

	marketsArbitrum, e := queryAllMarkets("Arbitrum")
	if e != nil {
		if err == nil {
			err = e
		} else {
			err = fmt.Errorf("%w, %w", err, e)
		}
	}

	options := make([]rainbow.Option, 0, 2*len(marketsOptimism)+2*len(marketsPolygon)+2*len(marketsArbitrum))

	e = processMarkets(&options, marketsOptimism, "Optimism")
	if e != nil {
		if err == nil {
			err = e
		} else {
			err = fmt.Errorf("%w, %w", err, e)
		}
	}

	e = processMarkets(&options, marketsPolygon, "Polygon")
	if e != nil {
		if err == nil {
			err = e
		} else {
			err = fmt.Errorf("%w, %w", err, e)
		}
	}

	e = processMarkets(&options, marketsArbitrum, "Arbitrum")
	if e != nil {
		if err == nil {
			err = e
		} else {
			err = fmt.Errorf("%w, %w", err, e)
		}
	}

	return options, err
}

func queryAllMarkets(layer string) ([]thales.AllMarketsMarketsMarket, error) {
	graphqlClient := graphql.NewClient(LayerURL(layer), nil)
	resp, err := thales.AllMarkets(context.TODO(), graphqlClient, skip, first)
	if err != nil {
		return nil, log.Error("queryAllMarkets(%s) %s", layer, err).Err()
	}
	if resp == nil {
		return nil, log.Error("queryAllMarkets(%s) resp=nil", layer).Err()
	}
	return resp.Markets, nil
}

func processMarkets(options *[]rainbow.Option, markets []thales.AllMarketsMarketsMarket, layer string) error {
	log.Printf("Processing %s %v options\n", layer, len(markets))
	var err error
	iv := make(map[string]float64)

	for i := range markets {
		// HOTFIX for bug on Polygon
		// 3 markets for BTC with very low strike
		// TODO properly understand this error "execution reverted: uint overflow from multiplication"
		// remove annoying market
		up, errUp := getOption(&markets[i], UP, layer, iv)
		if errUp != nil {
			errUp = log.Error("#", i, " getOption: "+markets[i].Id+" UP:", err).Err()
			if err == nil {
				err = errUp
			} else {
				err = fmt.Errorf("%w, %w", err, errUp)
			}
		} else {
			*options = append(*options, up)
		}

		down, errDown := getOption(&markets[i], DOWN, layer, iv)
		if errDown != nil {
			errDown = log.Error("#", i, " getOption:"+markets[i].Id+" DOWN:", err).Err()
			if err == nil {
				err = errDown
			} else {
				err = fmt.Errorf("%w, %w", err, errUp)
			}
		} else {
			*options = append(*options, down)
		}

		// pause because we don't have a premium RPC (we too poor)
		// TODO becom rich enough to afford a RPC
		if i%10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}

	return err
}

func getOption(m *thales.AllMarketsMarketsMarket, side uint8, layer string, iv map[string]float64) (rainbow.Option, error) {
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

	underlyingAssetIV := iv[m.CurrencyKey]
	if underlyingAssetIV == 0.0 {
		underlyingAssetIV, err = getIV(m, layer)
		if err != nil {
			log.Error(err)
			return rainbow.Option{}, err
		}
		iv[m.CurrencyKey] = underlyingAssetIV
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
		Layer:        l1orl2(layer),
		LayerName:    layer,
		// Thales store the IV of the underlying asset, which we just store in MarketIV
		// This is not the Market IV of the options
		// it's update by whitelist addresses
		// TODO learn how thales compute that IV
		MarketIV:      underlyingAssetIV,
		ProtocolID:    m.Id,
		Provider:      name + "::" + layer,
		URL:           url(m.Id),
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
	err = err1

	if err1 != nil && err2 != nil {
		log.Error("market error buy/sell")
		// log.Error(err1)
		// log.Error(err2)
		// both error are logged so doesn't really matter which I send back
		// I'm assuming that there is a real problem is both side have issue, if not
		// as long as one works, we store that data
		// that's wht I'm going with right now
		return rainbow.Option{}, err2
	}

	if err2 == nil {
		binary.Bid = append(binary.Bid, rainbow.Order{
			Price: sell,
			Size:  float64(amount),
		})
	} else {
		err = err2
	}

	if err1 == nil {
		binary.Ask = append(binary.Ask, rainbow.Order{
			Price: buy,
			Size:  float64(amount),
		})
	}

	return binary, err
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
	switch action {
	case "BUY":
		quote, err = instance.BuyFromAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), side, amountToQuote)
		if err != nil {
			log.Error("Thales BuyFromAmmQuote on", layer, err)
			return 0, err
		}
	case "SELL":
		quote, err = instance.SellToAmmQuote(&bind.CallOpts{}, common.HexToAddress(m.Id), side, amountToQuote)
		if err != nil {
			log.Error("Thales SellToAmmQuote on", layer, err)
			return 0, err
		}
	}

	decimals := LayerDecimals(layer)
	return rainbow.ToFloat(quote, decimals), nil
}

func getIV(m *thales.AllMarketsMarketsMarket, layer string) (float64, error) {
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
	var key [32]byte
	copy(key[:], common.FromHex(m.CurrencyKey))

	quote, err := instance.ImpliedVolatilityPerAsset(&bind.CallOpts{}, key)
	if err != nil {
		log.Error("Thales ImpliedVolatilityPerAsset on ", layer, err)
		return 0, err
	}
	return rainbow.ToFloat(quote, rainbow.DefaultEthereumDecimals), nil
}

func l1orl2(layer string) string {
	if layer == "Bsc" {
		return "L1"
	}
	return "L2"
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

func QueryMarkets(url string) ([]thales.MarketsMarketsMarket, error) {
	graphqlClient := graphql.NewClient(url, nil)
	minExpiry := rainbow.TwoWeeksInThePast()
	resp, err := thales.Markets(context.TODO(), graphqlClient, skip, first, minExpiry)
	if err != nil {
		return nil, fmt.Errorf("Markets: %w", err)
	}
	if resp == nil {
		return nil, errors.New("Markets: resp=nil")
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
		return nil, errors.New("Market: resp=nil")
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

// url to the option
// useful to put thereferral on the front.
func url(id string) string {
	return baseURL + id + referral
}
