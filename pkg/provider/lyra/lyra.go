// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package lyra

import (
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/pkg/provider/lyra/arbitrum"
)

var log = emo.NewZone("Lyra")

// DOC: https://docs.lyra.finance/developers/deployed-contracts

const (
	rpcOP                 = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m" // "https://mainnet.optimism.io"
	name                  = "Lyra"
	lyraRegistryOP        = "0xF5A0442D4753cA1Ea36427ec071aa5E786dA5916"
	optionMarketViewerOP  = "0xEAf788AD8abd9C98bA05F6802a62B8DbC673D76B"
	quoterAddressOP       = "0xea83ee73eB397c5974CB6b5220AE0A32fbE48B2B"
	rpcARB                = "https://arb-mainnet.g.alchemy.com/v2/hnBqLngSXPbAdvXHjcstEHkvWXV7RzEJ"
	lyraRegistryARB       = "0x6c87e4364Fd44B0D425ADfD0328e56b89b201329"
	optionMarketViewerARB = "0x0527A05c3CEBc9Ef8171FeD29DE5900A7ea093a4"
	quoterAddressARB      = "0x4CdB992fDEFcb840125dd344f87BDCbb8fEfA3e7"

	oneOption = 1
)

// TODO remove OptionMarket.go/abi if not necessary
// I guess it was there for testing in the beginning.

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	marketsOP, clientOP, err := GetOptionsFromLayer("Optimism")
	if err != nil {
		return nil, log.Error("GetOptionsFromLayer", "Optimism", err).Err()
	}
	spew.Dump(marketsOP)
	optionsOP, sumOP, err := processMarketsFromLayer("Optimism", marketsOP, clientOP)
	if err != nil {
		return nil, log.Error("processMarketsFromLayer", "Optimism", err).Err()
	}

	marketsARB, clientARB, err := GetOptionsFromLayer("Arbitrum")
	if err != nil {
		return nil, log.Error("GetOptionsFromLayer", "Arbitrum", err).Err()
	}
	spew.Dump(marketsARB)

	optionsARB, sumARB, err := processMarketsFromLayer("Arbitrum", marketsARB, clientARB)
	if err != nil {
		return nil, log.Error("processMarketsFromLayer", "Arbitrum", err).Err()
	}
	options = append(options, optionsOP...)
	options = append(options, optionsARB...)
	sum := sumOP + sumARB
	log.Info("Lyra total markets", sum)

	return options, nil
}

func GetOptionsFromLayer(layer string) (*[]common.Address, *ethclient.Client, error) {
	rpc := LayerRPC(layer)
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, &ethclient.Client{}, log.Error("Lyra ethclient", layer, err).Err()
	}
	lyraRegistry := LayerRegistry(layer)
	address := common.HexToAddress(lyraRegistry)
	registry, err := NewLyrar(address, client)
	if err != nil {
		return nil, &ethclient.Client{}, log.Error("Lyra registry", layer, err).Err()
	}

	// DO NOT use make() here!
	// we don't want to have 0x00...0 initialize here.
	// and we don't know how many market there will be.
	markets := []common.Address{}
	var i int64
	var market common.Address
	err = nil
	for ; err == nil; i++ {
		market, err = registry.OptionMarkets(&bind.CallOpts{}, big.NewInt(i))
		if err == nil {
			markets = append(markets, market)
		}
	}
	if len(markets) == 0 {
		log.Error("registry.OptionMarkets return empty array")
	}
	return &markets, client, nil
}

func processMarketsFromLayer(layer string, markets *[]common.Address, client *ethclient.Client) ([]rainbow.Option, int, error) {
	log.Printf("Processing options on %s  \n", layer)

	options := []rainbow.Option{}
	sum := 0

	q, err := LayerQuoter(layer, client)
	if err != nil {
		return nil, 0, err
	}

	viewer, err := LayerMarketViewer(layer, client)
	if err != nil {
		return nil, 0, err
	}

	for i := 0; i < len(*markets); i++ {
		marketAddresses, err := viewer.MarketAddresses(&bind.CallOpts{}, (*markets)[i])
		if err != nil {
			return nil, 0, log.Error("MarketAddresses", layer, err).Err()
		}
		baseAsset := Asset(marketAddresses.BaseAsset)

		boards, err := viewer.GetLiveBoards(&bind.CallOpts{}, (*markets)[i])
		if err != nil {
			spew.Dump((*markets)[i])
			return nil, 0, log.Error("GetLiveBoards", layer, err).Err()
		}

		for _, b := range boards {
			sum += len(b.Strikes)

			for ii := range b.Strikes {
				callPut, err := b.process(ii, baseAsset, q, layer)
				if err != nil {
					return nil, 0, log.Error("process", layer, err).Err()
				}
				options = append(options, callPut...)
			}
		}
	}
	log.Printf("Processed  %v options on %s\n", len(options), layer)

	return options, sum, nil
}

func (b *OptionMarketViewerBoardView) process(i int, asset string, quoter Quoter, layer string) ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	call := rainbow.Option{
		Name:            "",
		Type:            "CALL",
		Asset:           asset,
		Expiry:          expiration(b.Expiry),
		ExchangeType:    "DEX",
		Chain:           "Ethereum",
		Layer:           "L2",
		LayerName:       layer,
		Provider:        name + "::" + layer,
		QuoteCurrency:   "USD", // sUSD or USDC
		UnderlyingQuote: LayerUnderlyingQuoteCurrency(layer),
		Bid:             nil,
		Ask:             nil,
		Strike:          rainbow.ToFloat(b.Strikes[i].StrikePrice, rainbow.DefaultEthereumDecimals),
	}
	put := rainbow.Option{
		Name:            "",
		Type:            "PUT",
		Asset:           asset,
		Expiry:          expiration(b.Expiry),
		ExchangeType:    "DEX",
		Chain:           "Ethereum",
		Layer:           "L2",
		LayerName:       layer,
		Provider:        name + "::" + layer,
		QuoteCurrency:   "USD", // sUSD or USDC
		UnderlyingQuote: LayerUnderlyingQuoteCurrency(layer),
		Bid:             nil,
		Ask:             nil,
		Strike:          rainbow.ToFloat(b.Strikes[i].StrikePrice, rainbow.DefaultEthereumDecimals),
	}

	call.Name = call.OptionName()
	put.Name = put.OptionName()

	call.URL = url(&call) // , b.Strikes[i].StrikeId)
	put.URL = url(&put)   // , b.Strikes[i].StrikeId)

	// Market IV = board IV (baseIV) * Skew
	call.MarketIV = rainbow.ToFloat(b.BaseIv, rainbow.DefaultEthereumDecimals) *
		rainbow.ToFloat(b.Strikes[i].Skew, rainbow.DefaultEthereumDecimals)
	// keep only 5 decimals (0.XXXXX) and show it as XX.XXX%
	call.MarketIV = math.Floor(call.MarketIV*10_000_000) / 100_000
	put.MarketIV = call.MarketIV

	bidasks, err := getBidsAsks(b.Strikes[i].StrikeId, b.Market, oneOption, quoter)
	if err != nil {
		return options, log.Error("getBidsAsks", err).Err()
	}

	call.Bid = append(call.Bid, rainbow.Order{
		Price: bidasks[3],
		Size:  float64(oneOption),
	})

	call.Ask = append(call.Ask, rainbow.Order{
		Price: bidasks[0],
		Size:  float64(oneOption),
	})
	put.Bid = append(put.Bid, rainbow.Order{
		Price: bidasks[4],
		Size:  float64(oneOption),
	})

	put.Ask = append(put.Ask, rainbow.Order{
		Price: bidasks[1],
		Size:  float64(oneOption),
	})
	options = append(options, call, put)

	return options, err
}

// TODO use fullquotes on arbitrum
func getBidsAsks(strikeID *big.Int, market common.Address, amount int, quoter Quoter) ([]float64, error) {
	// We use 3 iterations (common.Big3) here because  that's what they do in the UI
	premium, _, err := quoter.FullQuotes(&bind.CallOpts{}, market, strikeID, common.Big3,
		rainbow.IntToEthereumUint256(amount, 18))
	if err != nil {
		return nil, log.Error("FullQuotes market=", market, "strikeID=", strikeID, err).Err()
	}
	return rainbow.ToFloats(premium, rainbow.DefaultEthereumDecimals), nil
}

func Asset(address common.Address) string {
	// those asset are part of Synthetix so if it's not recognized
	// it is an unknown synthetic asset.

	switch {
	case address.String() == sETH:
		return "sETH"
	case address.String() == sBTC:
		return "sBTC"
	case address.String() == sSOL:
		return "sSOL"
	case address.String() == sLINK:
		return "sLINK"
	case address.String() == WBTC:
		return "WBTC"
	case address.String() == WETH:
		return "WETH"
	default:
		log.Warn("Lyra Unknown token:", address.String())
		return "LLLL"
	}
}

// TODO check function on their frontend.
func url(o *rainbow.Option /*, strikeID *big.Int*/) string {
	// base := "https://app.lyra.finance/position/?"
	base := "https://app.lyra.finance/trade"
	asset := strings.ToLower(o.Asset[1:])
	// TODO if they include asset with decimal, modify this
	// most likely example: SOL, LINK
	// strike := fmt.Sprintf("%.f", rainbow.ToFloat(strikeID, 0))
	// t := strings.ToLower(o.Type)

	return base + "/" + asset //+ "/" + strike + "/" + t
	// return base + "market=" + asset + "&id=" + ""
}

func expiration(e *big.Int) string {
	seconds := e.Int64()
	expiryTime := time.Unix(seconds, 0).UTC()
	return expiryTime.Format("2006-01-02 15:04:05")
}

func LayerRPC(layer string) string {
	switch layer {
	case "Optimism":
		return rpcOP
	case "Arbitrum":
		return rpcARB
	}
	log.Panic("Unexpected layer", layer)
	return ""
}

// LayerRegistry returns lyraRegistry Address for the choosen layer
// This is fine since it does not seem like the code is different between layers
func LayerRegistry(layer string) string {
	switch layer {
	case "Optimism":
		return lyraRegistryOP
	case "Arbitrum":
		return lyraRegistryARB
	}
	log.Panic("Unexpected layer", layer)
	return ""
}

func LayerQuoter(layer string, client *ethclient.Client) (Quoter, error) {
	if layer == "Optimism" {
		quoter, err := NewQuoterOP(common.HexToAddress(quoterAddressOP), client)
		if err != nil {
			return nil, log.Error("quoter contract", layer, err).Err()
		}
		return quoter, nil
	} else if layer == "Arbitrum" {
		quoter, err := NewQuoterARB(common.HexToAddress(quoterAddressARB), client)
		if err != nil {
			return nil, log.Error("quoter contract", layer, err).Err()
		}
		return quoter, nil
	}
	//just to remove warning
	return nil, nil
}

// LayerMarketViewer to view market, what did you expect
// I am at the "I'm too tired to properly code/comment"
func LayerMarketViewer(layer string, client *ethclient.Client) (MarketViewer, error) {
	if layer == "Optimism" {
		viewer, err := NewMarketViewerOP(common.HexToAddress(optionMarketViewerOP), client)
		if err != nil {
			return nil, log.Error("quoter contract", layer, err).Err()
		}
		return viewer, nil
	} else if layer == "Arbitrum" {
		viewer, err := arbitrum.NewMarketViewerARB(common.HexToAddress(optionMarketViewerARB), client)
		if err != nil {
			return nil, log.Error("quoter contract", layer, err).Err()
		}
		return viewer, nil
	}
	//just to remove warning
	return nil, nil
}

func LayerUnderlyingQuoteCurrency(layer string) string {
	switch layer {
	case "Optimism":
		return "sUSD"
	case "Arbitrum":
		return "USDC"
	}
	log.Panic("Unexpected layer", layer)
	return "UUU"

}

type Quoter interface {
	FullQuotes(opts *bind.CallOpts, _optionMarket common.Address, strikeId *big.Int, iterations *big.Int, amount *big.Int) ([]*big.Int, []*big.Int, error)
}

type MarketViewer interface {
	MarketAddresses(opts *bind.CallOpts, arg0 common.Address) (struct {
		LiquidityPool      common.Address
		LiquidityToken     common.Address
		GreekCache         common.Address
		OptionMarket       common.Address
		OptionMarketPricer common.Address
		OptionToken        common.Address
		ShortCollateral    common.Address
		PoolHedger         common.Address
		QuoteAsset         common.Address
		BaseAsset          common.Address
	}, error)
	GetLiveBoards(opts *bind.CallOpts, market common.Address) ([]OptionMarketViewerBoardView, error)
}
