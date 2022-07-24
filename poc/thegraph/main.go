// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spewerspew/spew"
	"github.com/teal-finance/rainbow/pkg/provider/thales"
	"github.com/teal-finance/rainbow/pkg/provider/zerox"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

func main() {
	options := zerox.QueryTheGraph()
	spew.Dump(len(options))
	//const url = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
	//id := "0x8cb6c58a63a568f8155dc3242f3b830c57dba5eb"

	const url = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
	id := "0xb349f4f62c92b7deb4ee7fadc6022c0830612aa4" //"0xe8dd2d01bc36babe1eecbbe863ad294bbc5c15df"
	opt := thales.QueryMarket(id, url)
	spew.Dump(opt)

	/*allLive, _ := thales.QueryAllLiveMarkets(url)
	spew.Dump(len(allLive))

	markets, _ := thales.QueryAllMarkets(url)
	fmt.Println("test")
	spew.Dump(len(markets))*/
	fmt.Println(thales.Underlying(opt.CurrencyKey))
	rpc, _, amm, decimals := thales.LayerInfo("Polygon") // "Polygon") //"Optimism")
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(amm, "\t", rpc, "\t", decimals)
	//AMM := "0x5ae7454827D83526261F3871C1029792644Ef1B1"

	address := common.HexToAddress(amm)
	instance, err := thales.NewThales(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	const UP = 0
	const DOWN = 1
	amount := rainbow.IntToEthereumUint256(1, rainbow.DefaultEthereumDecimals)
	k, err := instance.CapPerMarket(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Testing setup")

	spew.Dump(k)
	spew.Dump(rainbow.ToFloat(k, rainbow.DefaultEthereumDecimals))
	spew.Dump(rainbow.ToFloat(amount, rainbow.DefaultEthereumDecimals))
	spew.Dump(decimals)
	spew.Dump(amount)

	fmt.Println("quote buy UP")
	Mistake := fmt.Errorf("execution reverted: uint overflow from multiplication")
	quote, err := instance.BuyFromAmmQuote(&bind.CallOpts{}, common.HexToAddress(id), UP, amount)
	if err == Mistake {
		spew.Dump(err)
		log.Fatal(err)

	}
	spew.Dump(quote)
	spew.Dump(rainbow.ToFloat(quote, decimals))
	spew.Dump(rainbow.ToFloat(amount, rainbow.DefaultEthereumDecimals))

	fmt.Println("quote sell UP")
	quote, err = instance.SellToAmmQuote(&bind.CallOpts{}, common.HexToAddress(id), UP, amount)
	if err != nil {
		log.Fatal(err)

	}
	spew.Dump(quote)

	spew.Dump(rainbow.ToFloat(quote, decimals))
	spew.Dump(rainbow.ToFloat(amount, rainbow.DefaultEthereumDecimals))

	fmt.Println("quote buy DOWN")

	quote, err = instance.BuyFromAmmQuote(&bind.CallOpts{}, common.HexToAddress(id), DOWN, amount)
	if err != nil {
		log.Fatal(err)

	}
	spew.Dump(quote)

	spew.Dump(rainbow.ToFloat(quote, decimals))
	spew.Dump(rainbow.ToFloat(amount, rainbow.DefaultEthereumDecimals))

	fmt.Println("quote sell DOWN")

	quote, err = instance.SellToAmmQuote(&bind.CallOpts{}, common.HexToAddress(id), DOWN, amount)
	if err != nil {
		log.Fatal(err)

	}
	spew.Dump(quote)

	spew.Dump(rainbow.ToFloat(quote, decimals))
	spew.Dump(rainbow.ToFloat(amount, rainbow.DefaultEthereumDecimals))

	t, _ := rainbow.TimeStringConvert(opt.MaturityDate)
	fmt.Println("expiration ", t)
	/*
		var t int64 = 1654041600 //01-JUN-2022
		response, err := thales.Markets(context.TODO(), graphqlClient, 0, 0, t)
		if err != nil {
			log.Print("ERR: ", err)
		}

		if resp == nil {
			log.Print("ERR: resp=nil")
			return
		}
		spew.Dump(len(response.Markets))
		//name("0x4c554e4100000000000000000000000000000000000000000000000000000000")
		for _, m := range response.Markets {
			if m.Timestamp != 0 { //name(m.CurrencyKey) == "BTC" && m.ExpiryDate == int64(1672992000) {
				spew.Dump(m)
			}

		}
	*/
}
