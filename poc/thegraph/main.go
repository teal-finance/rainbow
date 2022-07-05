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
	const url = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
	id := "0x5416c2ab11c7852ed9648aa006ee69d412c735c9"

	//const url = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
	//id := "0x3831f981fe62a081ccf4acb7ed1d3fea1ac761db"
	opt := thales.QueryMarket(id, url)
	spew.Dump(opt)

	allLive, _ := thales.QueryAllLiveMarkets(url)
	spew.Dump(len(allLive))

	markets := thales.QueryAllMarkets(url)
	spew.Dump(len(markets))
	fmt.Println(thales.Underlying(opt.CurrencyKey))

	client, err := ethclient.Dial(thales.OptimismRPC)
	if err != nil {
		log.Fatal(err)
	}
	AMM := "0x5ae7454827D83526261F3871C1029792644Ef1B1"

	address := common.HexToAddress(AMM)
	instance, err := thales.NewThales(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	const UP = 0
	const DOWN = 1
	amount := rainbow.IntToEthereumFormat(1)

	quote, err := instance.BuyFromAmmQuote(&bind.CallOpts{}, common.HexToAddress(id), UP, amount)
	if err != nil {
		log.Fatal(err)

	}
	spew.Dump(rainbow.ToFloat(quote))
	spew.Dump(rainbow.ToFloat(amount))

	quote, err = instance.SellToAmmQuote(&bind.CallOpts{}, common.HexToAddress(id), UP, amount)
	if err != nil {
		log.Fatal(err)

	}
	spew.Dump(rainbow.ToFloat(quote))
	spew.Dump(rainbow.ToFloat(amount))

	quote, err = instance.BuyFromAmmQuote(&bind.CallOpts{}, common.HexToAddress(id), DOWN, amount)
	if err != nil {
		log.Fatal(err)

	}
	spew.Dump(rainbow.ToFloat(quote))
	spew.Dump(rainbow.ToFloat(amount))

	quote, err = instance.SellToAmmQuote(&bind.CallOpts{}, common.HexToAddress(id), DOWN, amount)
	if err != nil {
		log.Fatal(err)

	}
	spew.Dump(rainbow.ToFloat(quote))
	spew.Dump(rainbow.ToFloat(amount))

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
