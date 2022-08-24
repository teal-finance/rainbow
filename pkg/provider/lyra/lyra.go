// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package lyra

import (
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spewerspew/spew"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	optimismrpc        = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m" // "https://mainnet.optimism.io"
	name               = "Lyra"
	lyraRegistry       = "0xF5A0442D4753cA1Ea36427ec071aa5E786dA5916"
	optionMarketViewer = "0xEAf788AD8abd9C98bA05F6802a62B8DbC673D76B"
)

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	client, err := ethclient.Dial(optimismrpc)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(lyraRegistry)
	log.Print(address)
	registry, err := NewLyrar(address, client)
	if err != nil {
		log.Print("ERR: ", err)
		return options, err
	}

	// DO NOT use make() here!
	// we don't want to have 0x00...0 initialize here.
	// and we don't know how many market there will be.
	markets := []common.Address{}
	spew.Dump(markets)
	var i int64
	var market common.Address
	err = nil
	for ; err == nil; i++ {
		market, err = registry.OptionMarkets(&bind.CallOpts{}, big.NewInt(i))
		if err == nil {
			markets = append(markets, market)
			log.Println(market, err)

		} else {
			log.Println("escaped ERR:", err)
		}
	}
	spew.Dump(markets)
	if len(markets) == 0 {
		log.Print("ERR: registry.OptionMarkets return empty array")
	}

	sum := 0
	viewer, err := NewLyrap(common.HexToAddress(optionMarketViewer), client)
	if err != nil {
		log.Print("ERR: ", err)
		return options, err
	}

	for i := 0; i < len(markets); i++ {
		marketAddresses, err := viewer.MarketAddresses(&bind.CallOpts{}, markets[i])
		if err != nil {
			log.Print("ERR: MarketAddresses ", err)

			return nil, err
		}
		baseAsset := Asset(marketAddresses.BaseAsset)

		boards, err := viewer.GetLiveBoards(&bind.CallOpts{}, markets[i])
		if err != nil {
			log.Print("ERR: GetLiveBoards ", err)

			return nil, err
		}

		for _, b := range boards {

			sum += len(b.Strikes)

			for _, s := range b.Strikes {
				callPut := process(s, b, baseAsset)
				options = append(options, callPut...)
			}
		}
	}

	log.Print("INF Lyra total markets ", sum)

	return options, nil
}
func process(s OptionMarketViewerStrikeView, b OptionMarketViewerBoardView, asset string) []rainbow.Option {
	options := []rainbow.Option{}
	call := rainbow.Option{
		Name:          "",
		Type:          "CALL",
		Asset:         asset,
		Expiry:        expiration(b.Expiry),
		ExchangeType:  "DEX",
		Chain:         "Ethereum",
		Layer:         "L2",
		LayerName:     "Optimism",
		Provider:      name,
		QuoteCurrency: "USD", // sUSD but anyway
		Bid:           nil,
		Ask:           nil,
		Strike:        rainbow.ToFloat(s.StrikePrice, rainbow.DefaultEthereumDecimals),
	}
	put := rainbow.Option{
		Name:          "",
		Type:          "PUT",
		Asset:         asset,
		Expiry:        expiration(b.Expiry),
		ExchangeType:  "DEX",
		Chain:         "Ethereum",
		Layer:         "L2",
		LayerName:     "Optimism",
		Provider:      name,
		QuoteCurrency: "USD", // sUSD but anyway
		Bid:           nil,
		Ask:           nil,
		Strike:        rainbow.ToFloat(s.StrikePrice, rainbow.DefaultEthereumDecimals),
	}

	call.Name = call.OptionName()
	put.Name = put.OptionName()

	options = append(options, call, put)
	return options
}
func Asset(address common.Address) string {
	// those asset are part of Synthetix so if it's not recognized
	// it is an unknwow synthetic asset.

	switch {
	case address.String() == sETH:
		return "sETH"
	case address.String() == sBTC:
		return "sBTC"
	case address.String() == sSOL:
		return "sSOL"
	case address.String() == sLINK:
		return "sLINK"
	default:
		log.Print("WRN Lyra Unknown token: ", address.String())
		return "LLLL"
	}
}

/*
func (v *Lyrap) getBidsAsks(boardListing *big.Int, amount int) ([]OptionMarketViewerTradePremiumView, error) {
	ammOrder := []OptionMarketViewerTradePremiumView{}
	a := rainbow.IntToEthereumUint256(amount, rainbow.DefaultEthereumDecimals)

	// Call BID
	trade, err := v.GetPremiumForOpen(&bind.CallOpts{}, boardListing, 1, a)
	if err != nil {
		return nil, err
	}

	ammOrder = append(ammOrder, trade)

	// Call ASK
	trade, err = v.GetPremiumForOpen(&bind.CallOpts{}, boardListing, 0, a)
	if err != nil {
		return nil, err
	}

	ammOrder = append(ammOrder, trade)

	// Put BID
	trade, err = v.GetPremiumForOpen(&bind.CallOpts{}, boardListing, 3, a)
	if err != nil {
		return nil, err
	}

	ammOrder = append(ammOrder, trade)

	// Put ASK
	trade, err = v.GetPremiumForOpen(&bind.CallOpts{}, boardListing, 2, a)
	if err != nil {
		return nil, err
	}

	ammOrder = append(ammOrder, trade)
	return ammOrder, err
}

func processOption(listing *OptionMarketViewerListingView, ammOrder []OptionMarketViewerTradePremiumView, amount int, asset string) []rainbow.Option {
	options := []rainbow.Option{}
	call := rainbow.Option{
		Name:          "",
		Type:          "CALL",
		Asset:         asset,
		Expiry:        expiration(listing.Expiry),
		ExchangeType:  "DEX",
		Chain:         "Ethereum",
		Layer:         "L2",
		Provider:      name,
		QuoteCurrency: "USD", // sUSD but anyway
		Bid:           nil,
		Ask:           nil,
		Strike:        rainbow.ToFloat(listing.Strike, rainbow.DefaultEthereumDecimals),
	}
	put := rainbow.Option{
		Name:          "",
		Type:          "PUT",
		Asset:         asset,
		Expiry:        expiration(listing.Expiry),
		ExchangeType:  "DEX",
		Chain:         "Ethereum",
		Layer:         "L2",
		Provider:      name,
		QuoteCurrency: "USD", // sUSD but anyway
		Bid:           nil,
		Ask:           nil,
		Strike:        rainbow.ToFloat(listing.Strike, rainbow.DefaultEthereumDecimals),
	}

	call.Name = call.OptionName()
	put.Name = put.OptionName()

	call.Bid = append(call.Bid, rainbow.Order{
		Price: rainbow.ToFloat(ammOrder[0].Premium, rainbow.DefaultEthereumDecimals),
		Size:  float64(amount),
	})

	call.Ask = append(call.Ask, rainbow.Order{
		Price: rainbow.ToFloat(ammOrder[1].Premium, rainbow.DefaultEthereumDecimals),
		Size:  float64(amount),
	})

	put.Bid = append(put.Bid, rainbow.Order{
		Price: rainbow.ToFloat(ammOrder[2].Premium, rainbow.DefaultEthereumDecimals),
		Size:  float64(amount),
	})

	put.Ask = append(put.Ask, rainbow.Order{
		Price: rainbow.ToFloat(ammOrder[3].Premium, rainbow.DefaultEthereumDecimals),
		Size:  float64(amount),
	})

	options = append(options, call, put)
	return options
}
*/

func expiration(e *big.Int) string {
	seconds := e.Int64()
	expiryTime := time.Unix(seconds, 0).UTC()
	return expiryTime.Format("2006-01-02 15:04:05")
}
