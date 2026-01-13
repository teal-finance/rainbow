// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package lyra

import (
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lynxai-team/emo"
	"github.com/spewerspew/spew"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone("Lyra")

// DOC: https://docs.lyra.finance/developers/deployed-contracts

const (
	// TODO use teal on Optimism
	// TODO use perso on Arbitrum
	// TODO do that on the 1st Nov
	// Sorry alchemy but we r too poor.

	// teal.
	rpcOP = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m"
	// perso
	// rpcOP = "https://opt-mainnet.g.alchemy.com/v2/uksZH_SjXAaBnIw95hZcBoCWGCXs9VXI"
	name = "Lyra"
	// teal
	// rpcARB = "https://arb-mainnet.g.alchemy.com/v2/hnBqLngSXPbAdvXHjcstEHkvWXV7RzEJ"
	// perso.
	rpcARB = "https://arb-mainnet.g.alchemy.com/v2/4TQ_6stSP__V97XUQQC07AV23f_XOemr"

	oneOption = 1
)

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	options := []rainbow.Option{}
	marketsARB, clientARB, err := GetOptionsFromLayer("Arbitrum")
	if err != nil {
		return nil, log.Error("GetOptionsFromLayer", "Arbitrum", err).Err()
	}

	optionsARB, err := processMarketsFromLayer("Arbitrum", marketsARB, clientARB)
	if err != nil {
		return nil, log.Error("processMarketsFromLayer", "Arbitrum", err).Err()
	}

	marketsOP, clientOP, err := GetOptionsFromLayer("Optimism")
	if err != nil {
		return nil, log.Error("GetOptionsFromLayer", "Optimism", err).Err()
	}
	optionsOP, err := processMarketsFromLayer("Optimism", marketsOP, clientOP)
	if err != nil {
		return nil, log.Error("processMarketsFromLayer", "Optimism", err).Err()
	}

	options = append(options, optionsARB...)
	options = append(options, optionsOP...)
	log.Info("Lyra total markets", len(options))

	return options, nil
}

func GetOptionsFromLayer(layer string) (*[]common.Address, *ethclient.Client, error) {
	rpc := ""
	lyraRegistry := ""

	switch layer {
	case "Optimism":
		rpc = rpcOP
		lyraRegistry = lyraRegistryOP
	case "Arbitrum":
		rpc = rpcARB
		lyraRegistry = lyraRegistryARB
	default:
		return nil, &ethclient.Client{}, log.Error("Unknown layer", layer).Err()
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, &ethclient.Client{}, log.Error("Lyra ethclient", layer, err).Err()
	}

	address := common.HexToAddress(lyraRegistry)
	registry, err := NewRegistry(address, client)
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
		log.Error("registry.OptionMarkets return empty array on + ", layer)
	}
	return &markets, client, nil
}

func processMarketsFromLayer(layer string, markets *[]common.Address, client *ethclient.Client) ([]rainbow.Option, error) {
	log.Printf("Processing options on %s  \n", layer)

	options := []rainbow.Option{}

	quoterAddress := ""
	optionMarketViewer := ""
	switch layer {
	case "Optimism":
		quoterAddress = quoterAddressOP
		optionMarketViewer = optionMarketViewerOP
	case "Arbitrum":
		quoterAddress = quoterAddressARB
		optionMarketViewer = optionMarketViewerARB
	default:
		return nil, log.Error("Unknown layer", layer).Err()
	}

	q, err := NewQuoter(common.HexToAddress(quoterAddress), client)
	if err != nil {
		return nil, log.Error("quoter contract", layer, err).Err()
	}

	viewer, err := NewMarketviewer(common.HexToAddress(optionMarketViewer), client)
	if err != nil {
		return nil, log.Error("MarketAddresses", layer, err).Err()
	}

	for i := range len(*markets) {
		marketAddresses, err := viewer.MarketAddresses(&bind.CallOpts{}, (*markets)[i])
		if err != nil {
			return nil, log.Error("MarketAddresses", layer, i, err).Err()
		}
		baseAsset := Asset(marketAddresses.BaseAsset)

		boards, err := viewer.GetLiveBoards(&bind.CallOpts{}, (*markets)[i])
		if err != nil {
			spew.Dump((*markets)[i])
			return nil, log.Error("GetLiveBoards", layer, err).Err()
		}

		for _, b := range boards {
			for ii := range b.Strikes {
				callPut, err := b.process(ii, baseAsset, q, layer)
				if err != nil {
					return nil, log.Error("process", layer, err).Err()
				}
				options = append(options, callPut...)
			}
		}
	}
	log.Printf("Processed  %v options on %s\n", len(options), layer)

	return options, nil
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

func Asset(address common.Address) string {
	// those asset are part of Synthetix so if it's not recognized
	// it is an unknown synthetic asset.

	switch {
	case address.String() == SETH:
		return "sETH"
	case address.String() == SBTC:
		return "sBTC"
	case address.String() == SSOL:
		return "sSOL"
	case address.String() == SLINK:
		return "sLINK"
	case address.String() == AWBTC:
		return "WBTC"
	case address.String() == OWBTC:
		return "WBTC"
	case address.String() == AWETH:
		return "WETH"
	case address.String() == OWETH:
		return "WETH"
	case address.String() == OP:
		return "OP"
	case address.String() == LyraARB:
		return "ARB"
	case address.String() == LINK:
		return "LINK"
	case address.String() == LyraXRP:
		return "XRP"
	default:
		log.Warn("Lyra Unknown token:", address.String())
		return "LLLL"
	}
}

// OptionMarketViewerBoardViewARB copies and slightly adapt OptionMarketViewerBoardView.
func (b *OptionMarketViewerBoardView) process(i int, asset string, quoter *Quoter, layer string) ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	call := rainbow.Option{
		Name:            "",
		Type:            "CALL",
		Asset:           asset,
		Expiry:          rainbow.Expiration(b.Expiry.Int64()),
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
		Expiry:          rainbow.Expiration(b.Expiry.Int64()),
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

	// TODO check properly
	call.URL = url(&call) // , b.Strikes[i].StrikeId)
	put.URL = url(&put)   // , b.Strikes[i].StrikeId)

	// Market IV = board IV (baseIV) * Skew //TODO verify it's still 18 on arbitrum
	call.MarketIV = rainbow.ToFloat(b.BaseIv, rainbow.DefaultEthereumDecimals) *
		rainbow.ToFloat(b.Strikes[i].Skew, rainbow.DefaultEthereumDecimals)
	// keep only 5 decimals (0.XXXXX) and show it as XX.XXX%
	call.MarketIV = math.Floor(call.MarketIV*10_000_000) / 100_000
	put.MarketIV = call.MarketIV

	bidasks, err := getBidsAsks(b.Strikes[i].StrikeId, b.Market, oneOption, quoter)
	if err != nil {
		return options, log.Error("getBidsAsks", layer, err).Err()
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

	return options, nil
}

func getBidsAsks(strikeID *big.Int, market common.Address, amount int, quoter *Quoter) ([]float64, error) {
	// We use 3 iterations (common.Big3) here because  that's what they do in the UI
	premium, _, err := quoter.FullQuotes(&bind.CallOpts{}, market, strikeID, common.Big3,
		rainbow.IntToEthereumUint256(amount, rainbow.DefaultEthereumDecimals))
	if err != nil {
		return nil, log.Error("FullQuotes market=", market, "strikeID=", strikeID, err).Err()
	}
	return rainbow.ToFloats(premium, rainbow.DefaultEthereumDecimals), nil
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
