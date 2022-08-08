// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package lyra

const (
	optionMarketBTC        = "0x47B5BB79F06F06db3D77C6cc4DB1ad6E84faF1f4"
	optionMarketViewerBTC  = "0x22c39cE1C3A49224Aea6D8c2AAa0019828E1b5E4"
	optionMarketETH        = "0x1f6D98638Eee9f689684767C3021230Dd68df419"
	optionMarketViewerETH  = "0x43592bffCF14f1e0A096091E125f023B2ccC2525"
	optionMarketLINK       = "0xBCB01210BD1c0790ca45Cc4C49d9a183be99824d"
	optionMarketViewerLINK = "0xE34AF32082aa2Be75f0a80dC179D597E852AF6a1"
	optionMarketSOL        = "0xA3562CAc1c39f4D4166CE31005Fc080AB41120aC"
	optionMarketViewerSOL  = "0xeAcB3f1a4f3137807f4e901167a999eF6f8312fb"
)

// Assets returns the list of assets supported by Lyra.
// order as to be the same for all the arrays
// TODO switch to maps or smth else.
var Assets = []string{
	"sBTC",
	"sETH",
	"sLINK",
	"sSOL",
}

var OptionMarkets = []string{
	optionMarketBTC,
	optionMarketETH,
	optionMarketLINK,
	optionMarketSOL,
}

var OptionMarketViewers = []string{
	optionMarketViewerBTC,
	optionMarketViewerETH,
	optionMarketViewerLINK,
	optionMarketViewerSOL,
}

//https://raw.githubusercontent.com/lyra-finance/lyra-protocol/master/deployments/mainnet-ovm/lyra.json
/* "markets":{
      "sETH":{
         "OptionMarket":{
            "contractName":"OptionMarket",
            "source":"OptionMarket",
            "address":"0x1d42a98848e022908069c2c545aE44Cc78509Bc8",
            "txn":"0x33e1f0acd79f21e894e28ed7106b28d98b6d4755f26bcd9757c8e8f96732521f",
            "blockNumber":12801071
         },
         "OptionMarketPricer":{
            "contractName":"OptionMarketPricer",
            "source":"OptionMarketPricer",
            "address":"0x73b161f1bcF37048A5173619cda53aaa56A28Be0",
            "txn":"0x574caff8dce286d0bcd06d3fa269b334eb2be27ac2a3de7e978c5ae777b9a62b",
            "blockNumber":12801079
         },
         "OptionGreekCache":{
            "contractName":"OptionGreekCache",
            "source":"OptionGreekCache",
            "address":"0xbb3e8Eac35e649ed1071A9Ec42223d474e67b19A",
            "txn":"0x2bda412287348d02c78f6e272e168e4bf13e65cf93015f89657a9c6b2a20f87a",
            "blockNumber":12801088
         },
         "OptionToken":{
            "contractName":"OptionToken",
            "source":"OptionToken",
            "address":"0xCfDfF4E171133D55dE2e45c66a0E144a135D93f2",
            "txn":"0xe0f147fc465b20f3bdce3e5518b4e53639b96d5d123ea2018c940a8153e010ae",
            "blockNumber":12801183
         },
         "LiquidityPool":{
            "contractName":"LiquidityPool",
            "source":"LiquidityPool",
            "address":"0x5Db73886c4730dBF3C562ebf8044E19E8C93843e",
            "txn":"0x6900e3ae214f72f860788d7b2249d32252287160ef0a77840988708e2ed4f446",
            "blockNumber":12801194
         },
         "LiquidityToken":{
            "contractName":"LiquidityToken",
            "source":"LiquidityToken",
            "address":"0x0d1a91354A387a1e9E8FCD8f576670c4C3b723cA",
            "txn":"0x291e8509f4d0bf7daf92bbbd924ebe70b598c579fee37ccf0a338aed2e06ce14",
            "blockNumber":12801215
         },
         "ShortCollateral":{
            "contractName":"ShortCollateral",
            "source":"ShortCollateral",
            "address":"0x3E86B53e1D7DA7eDbA225c3A218d0b5a7544fDfD",
            "txn":"0x1e06c7d5a380ed60d68571f4b98071073c30f7cd5e6da7967944614e32f033c9",
            "blockNumber":12801290
         },
         "ShortPoolHedger":{
            "contractName":"ShortPoolHedger",
            "source":"ShortPoolHedger",
            "address":"0x60a5159bAfb2198b967021AC77E26C1417081477",
            "txn":"0x017a02a1b54e07c0dd05bd7bcb58a788fd57f4a60095a8c0639715491e5b16f9",
            "blockNumber":12801313
         },
         "KeeperHelper":{
            "contractName":"KeeperHelper",
            "source":"KeeperHelper",
            "address":"0x6BA3bC09F00DE197Ab99E7Cde834dFD9A8E0B99C",
            "txn":"0x306afc9f34c3b488b7bf4033d698e56bb437f60d32751136131f15cacb9042ee",
            "blockNumber":14000326
         },
         "GWAVOracle":{
            "contractName":"GWAVOracle",
            "source":"GWAVOracle",
            "address":"0x9357EB9C4627F7cDdBdfF4f648Df5c2fB42cB69F",
            "txn":"0xcfbf38dae61aaafa85ad8ec78a1c5c925a4eb1a3e6879d45b2b55af6608d18e3",
            "blockNumber":12801362
         }
      }
   }
*/
