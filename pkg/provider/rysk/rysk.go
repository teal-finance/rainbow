// Copyright 2023 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package rysk

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/teal-finance/emo"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	name         = "Rysk"
	lens         = "0xa306C00e08ebC84a5F4F67b561B8F6EDeb77600D"
	rpc          = "https://arb-mainnet.g.alchemy.com/v2/4TQ_6stSP__V97XUQQC07AV23f_XOemr"
	url          = "https://app.rysk.finance/options"
	USDCDecimals = 6
)

var log = emo.NewZone(name)

type Provider struct {
}

func (Provider) Name() string {
	return name
}

func (p *Provider) Options() ([]rainbow.Option, error) {
	options := []rainbow.Option{}
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return []rainbow.Option{}, log.Error(name, " ethclient error", err).Err()
	}

	address := common.HexToAddress(lens)
	dhvLens, err := NewRysk(address, client)
	if err != nil {
		return []rainbow.Option{}, log.Error(name, " dhvlens contract error", err).Err()
	}
	exp, err := dhvLens.GetExpirations(&bind.CallOpts{})
	if err != nil {
		return []rainbow.Option{}, log.Error(name, " GetExpirations error", err).Err()
	}
	for _, e := range exp {

		expStr := rainbow.Expiration(int64(e))
		drill, err := dhvLens.GetOptionExpirationDrill(&bind.CallOpts{}, e)
		if err != nil {
			return []rainbow.Option{}, log.Error(name, " GetOptionExpirationDrill error", err).Err()
		}
		calls, err := process(expStr, drill.CallOptionDrill, "CALL")
		if err != nil {
			return []rainbow.Option{}, log.Error(name, " calls processing ", expStr, " error", err).Err()
		}
		puts, err := process(expStr, drill.PutOptionDrill, "PUT")
		if err != nil {
			return []rainbow.Option{}, log.Error(name, " puts processing ", expStr, " error", err).Err()
		}
		options = append(options, *calls...)
		options = append(options, *puts...)

	}
	return options, nil
}

func process(expStr string, drills []DHVLensMK1OptionStrikeDrill, optionType string) (*[]rainbow.Option, error) {
	options := []rainbow.Option{}

	for _, drill := range drills {
		delta := rainbow.ToFloat(drill.Delta, rainbow.DefaultEthereumDecimals)
		//spew.Dump(delta)
		//exposure := rainbow.ToFloat(drill.Exposure, rainbow.DefaultEthereumDecimals)
		//spew.Dump(exposure)

		strike := rainbow.ToFloat(drill.Strike, rainbow.DefaultEthereumDecimals)
		_, sell, sellfee, selldisable, sellpremiumtoosmall := convertQuotes(drill.Sell)
		priceSell := sell + sellfee
		if selldisable || sellpremiumtoosmall {
			priceSell = 0
		}
		iv, buy, buyfee, buydisable, buypremiumtoosmall := convertQuotes(drill.Buy)
		priceBuy := buy + buyfee

		if buydisable || buypremiumtoosmall {
			priceBuy = 0
		}

		o :=
			rainbow.Option{
				Type:            optionType,
				Asset:           "ETH",
				UnderlyingAsset: "WETH",
				Strike:          strike,
				Expiry:          expStr,
				ExchangeType:    "DEX",
				Chain:           "Ethereum",
				Layer:           "L2",
				LayerName:       "Arbitrum",
				Provider:        name,
				QuoteCurrency:   "USD",
				UnderlyingQuote: "USDC",
				URL:             url,
				MarketIV:        iv,
				/*
				  OpenInterest:
				  ProtocolID:*/
			}
		o.Name = o.OptionName()
		o.Greeks = rainbow.TheGreeks{
			Delta: delta,
		}
		o.Bid = append(o.Bid, rainbow.Order{
			Price: priceSell,
			Size:  1,
		})
		o.Ask = append(o.Ask, rainbow.Order{
			Price: priceBuy,
			Size:  1,
		})

		options = append(options, o)

	}

	return &options, nil

}

// convertQuotes converst DHVLensMK1TradingSpec to floats mostly
// return iv, quote, fee, disabled, premiumtoosmall
func convertQuotes(t DHVLensMK1TradingSpec) (float64, float64, float64, bool, bool) {
	iv := rainbow.ToFloat(t.Iv, rainbow.DefaultEthereumDecimals-2) // just to store it as XX.XXX% instead of .XXXXX
	quote := rainbow.ToFloat(t.Quote, USDCDecimals)
	fee := rainbow.ToFloat(t.Fee, USDCDecimals)

	return iv, quote, fee, t.Disabled, t.PremiumTooSmall
}

// TODO remove
// this is just to test and learn
func test() {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return //nil, &ethclient.Client{}, log.Error("Lyra ethclient", err).Err()
	}

	address := common.HexToAddress(lens)
	dhvLens, err := NewRysk(address, client)
	if err != nil {
		return // nil, &ethclient.Client{}, log.Error("Lyra registry", err).Err()
	}
	exp, _ := dhvLens.GetExpirations(&bind.CallOpts{})
	spew.Dump(exp)
	drill, err := dhvLens.GetOptionExpirationDrill(&bind.CallOpts{}, exp[0])
	if err != nil {
		return // nil, &ethclient.Client{}, log.Error("Lyra registry", err).Err()
	}
	cs := rainbow.ToFloats(drill.CallStrikes, rainbow.DefaultEthereumDecimals)
	co := drill.CallOptionDrill[3]
	ps := drill.PutStrikes
	po := drill.PutOptionDrill[7]
	spew.Dump(cs)
	spew.Dump(co)
	spew.Dump(ps)
	spew.Dump(po)
	spew.Dump(convertQuotes(co.Buy))
	spew.Dump(convertQuotes(co.Sell))
	spew.Dump(rainbow.Expiration(int64(exp[0])))

}
