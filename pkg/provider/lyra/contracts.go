package lyra

const optionMarketBTC = "0x47B5BB79F06F06db3D77C6cc4DB1ad6E84faF1f4"
const optionMarketViewerBTC = "0x22c39cE1C3A49224Aea6D8c2AAa0019828E1b5E4"
const optionMarketETH = "0x1f6D98638Eee9f689684767C3021230Dd68df419"
const optionMarketViewerETH = "0x43592bffCF14f1e0A096091E125f023B2ccC2525"
const optionMarketLINK = "0xBCB01210BD1c0790ca45Cc4C49d9a183be99824d"
const optionMarketViewerLINK = "0xE34AF32082aa2Be75f0a80dC179D597E852AF6a1"
const optionMarketSOL = "0xA3562CAc1c39f4D4166CE31005Fc080AB41120aC"
const optionMarketViewerSOL = "0xeAcB3f1a4f3137807f4e901167a999eF6f8312fb"

var OptionMarkets []string = []string{
	optionMarketBTC,
	optionMarketETH,
	optionMarketLINK,
	optionMarketSOL}

var OptionMarketViewers []string = []string{
	optionMarketViewerBTC,
	optionMarketViewerETH,
	optionMarketViewerLINK,
	optionMarketViewerSOL}

//https://github.com/lyra-finance/lyra-protocol/blob/master/deployments/mainnet-ovm/lyra.json#L31-L321
/*"markets": {
	"sETH": {
	  "OptionMarket": {
		"contractName": "OptionMarket",
		"source": "OptionMarket",
		"address": "0x1f6D98638Eee9f689684767C3021230Dd68df419",
		"txn": "0xfcd6164e7b8b2c1d3bb406bd47fb5ccd8979144127b4e4ce033dcd3035451d31",
		"blockNumber": 2482583
	  },
	  "OptionMarketPricer": {
		"contractName": "OptionMarketPricer",
		"source": "OptionMarketPricer",
		"address": "0x39A023FDe14d44c01bcb43993B3A51117174F336",
		"txn": "0x2fcc1886e9e8af5f3e4f1b7d5a0302885d86d1d52a8e3d0692d55949fa551eba",
		"blockNumber": 2482605
	  },
	  "OptionGreekCache": {
		"contractName": "OptionGreekCache",
		"source": "OptionGreekCache",
		"address": "0xbBaa63Be0E3feACb3a3798CB101962B8608fCa99",
		"txn": "0x34fe061e8c72bafb52a6d5003804e6731105f55f3f933983afae9e01752b974d",
		"blockNumber": 2482618
	  },
	  "OptionToken": {
		"contractName": "OptionToken",
		"source": "OptionToken",
		"address": "0xfB1951b7EeF8E7613D3b09424fB4aEf805c16267",
		"txn": "0x5b5066e34c560c7494194d76eecef612d372d4957149569e8ba1f31166d93992",
		"blockNumber": 2482641
	  },
	  "LiquidityPool": {
		"contractName": "LiquidityPool",
		"source": "LiquidityPool",
		"address": "0x2935CD347B79C319A6464fe3b1087170f142418C",
		"txn": "0x94c57bde53f03148a30cda1a78c961eaeceba889eb1e43e631236368e05c0738",
		"blockNumber": 2482673
	  },
	  "LiquidityCertificate": {
		"contractName": "LiquidityCertificate",
		"source": "LiquidityCertificate",
		"address": "0x8e9c516c202277D47Ccd33acEb081b90be2Fd5DC",
		"txn": "0x22109422af24123febf263bf949eca5cddb1e7ee4e65986acc022a9a480a86f0",
		"blockNumber": 2482683
	  },
	  "ShortCollateral": {
		"contractName": "ShortCollateral",
		"source": "ShortCollateral",
		"address": "0xE722F9aee66F649FBfc8CB0d4F906cb55803553c",
		"txn": "0x818bd6f653b6c47154f50517d761ddaa61e707ecd6206d57a4947518195d419c",
		"blockNumber": 2482704
	  },
	  "OptionMarketViewer": {
		"contractName": "OptionMarketViewer",
		"source": "OptionMarketViewer",
		"address": "0x43592bffCF14f1e0A096091E125f023B2ccC2525",
		"txn": "0xc6fb7096d18a3bd5d66fc9dd680b27f3fc509259a38f71c8bd16e4cba8ce1292",
		"blockNumber": 2922087
	  },
	  "PoolHedger": {
		"contractName": "PoolHedger",
		"source": "PoolHedger",
		"address": "0x9093F670c780d8D09a0c88D03d60AAA0B5a9E7A1",
		"txn": "0x2999e9f16f3684a0402448b1e4a5441542bc1706a2d024bafef5a853e243ca20",
		"blockNumber": 2482734
	  },
	  "OptionMarketSafeSlippage": {
		"contractName": "OptionMarketSafeSlippage",
		"source": "OptionMarketSafeSlippage",
		"address": "0xb921E22DE51f17B4227787A6d264F37399eB7D8C",
		"txn": "0x7093e95f38e833874f94cdafd6adedfc29e63ada142e8659e48ef4fee972d07b",
		"blockNumber": 2482752
	  }
	},
	"sLINK": {
	  "OptionMarket": {
		"contractName": "OptionMarket",
		"source": "OptionMarket",
		"address": "0xBCB01210BD1c0790ca45Cc4C49d9a183be99824d",
		"txn": "0x71e65299dd9b602219363b193ca0af8be2ce7390ecce4e46f9868527fdc47357",
		"blockNumber": 2482839
	  },
	  "OptionMarketPricer": {
		"contractName": "OptionMarketPricer",
		"source": "OptionMarketPricer",
		"address": "0x36Ab8de9F177B70fBdbC4050e47d0D578Adf3cFe",
		"txn": "0x7e9295afa641f3771ac26e88e69d185aace3c4af0e82308185574ff1f42e7f1e",
		"blockNumber": 2482854
	  },
	  "OptionGreekCache": {
		"contractName": "OptionGreekCache",
		"source": "OptionGreekCache",
		"address": "0xa5A20F41F367a7eD71d7a617a708De965cC80dB5",
		"txn": "0xe9a1ad6b134f83d736a5644f87e57d6c485433e54c4849e2af9fc76390f6050e",
		"blockNumber": 2482938
	  },
	  "OptionToken": {
		"contractName": "OptionToken",
		"source": "OptionToken",
		"address": "0x48A157f01cCF4Bf24d42185cBCD9Fc0C376292F0",
		"txn": "0x982d9990e6760acbc4fc7c9ac59df1a50c7fbe2c67346c555f67f6d38b80573a",
		"blockNumber": 2482944
	  },
	  "LiquidityPool": {
		"contractName": "LiquidityPool",
		"source": "LiquidityPool",
		"address": "0x69B4B35504a8c1d6179fef7AdDCDB37A8c663BC9",
		"txn": "0xe613be3ad570ab678686e107bf92c22e92981449a823351871a7a72b3fd8e68f",
		"blockNumber": 2482955
	  },
	  "LiquidityCertificate": {
		"contractName": "LiquidityCertificate",
		"source": "LiquidityCertificate",
		"address": "0x4341A26F8B32a6F79572351fAA4EC91964cf63eb",
		"txn": "0x0c0770ab0422b2c27d5cb02593ffdb3f3b48dc5f9709726809c6d56ebf7a4021",
		"blockNumber": 2482988
	  },
	  "ShortCollateral": {
		"contractName": "ShortCollateral",
		"source": "ShortCollateral",
		"address": "0x585a72ccecde68dDFE5327B23134723a305D70F3",
		"txn": "0x034ee0aa795dd845af63ea083271d72e1a019a893c7085f490ca34cc8f965fee",
		"blockNumber": 2483019
	  },
	  "OptionMarketViewer": {
		"contractName": "OptionMarketViewer",
		"source": "OptionMarketViewer",
		"address": "0xE34AF32082aa2Be75f0a80dC179D597E852AF6a1",
		"txn": "0x85402b1f5d2ae9dd6ac6e39bf7a6ad7806a84d9432e4a77c9dff947a5eb0a05f",
		"blockNumber": 2922118
	  },
	  "PoolHedger": {
		"contractName": "PoolHedger",
		"source": "PoolHedger",
		"address": "0x96e91f8dAfC58975025dD6163a41EB593e6e2C62",
		"txn": "0xdbdf2c7d0add7d08b3dc32c9966216194e996cd48d060bb5c9130c4c59933d7a",
		"blockNumber": 2483065
	  },
	  "OptionMarketSafeSlippage": {
		"contractName": "OptionMarketSafeSlippage",
		"source": "OptionMarketSafeSlippage",
		"address": "0x0dEC3173683e4B1E587d466FD0e3A8ad81a39a19",
		"txn": "0xb07d32082c34fd0d61b4608a77b008b0813892a4675d33926cc5eb67190ca0f8",
		"blockNumber": 2483085
	  }
	},
	"sBTC": {
	  "OptionMarket": {
		"contractName": "OptionMarket",
		"source": "OptionMarket",
		"address": "0x47B5BB79F06F06db3D77C6cc4DB1ad6E84faF1f4",
		"txn": "0x17619a77e3d19456a5d0ddd21b21f29b901849bb5a8a7b5c7d8eefe2c949173f",
		"blockNumber": 3753142
	  },
	  "OptionMarketPricer": {
		"contractName": "OptionMarketPricer",
		"source": "OptionMarketPricer",
		"address": "0xC1729854837578dC867C6d644304407F0EBc5A3D",
		"txn": "0x8759098cc4e8f42d168b56d13641b50742035142eaf387d5578734a8b3a73ebf",
		"blockNumber": 3753165
	  },
	  "OptionGreekCache": {
		"contractName": "OptionGreekCache",
		"source": "OptionGreekCache",
		"address": "0x1F80F640A9D53D8Cc339B1314061b6D14A24c78C",
		"txn": "0x32c5f6f5d180b49b0f0efc80ee2c8e849a554b4a41232e6cea4e4ab0db54dcc5",
		"blockNumber": 3753178
	  },
	  "OptionToken": {
		"contractName": "OptionToken",
		"source": "OptionToken",
		"address": "0x6C1d7B56A266752c611c022f4130E72952BF1327",
		"txn": "0x0e106e5aa08ae47934e7fe1e8fc0d01bd99fb08de9b796895fb4ee201d5f675e",
		"blockNumber": 3753184
	  },
	  "LiquidityPool": {
		"contractName": "LiquidityPool",
		"source": "LiquidityPool",
		"address": "0x788843DE0Be1598155bFFaAB7Cfa2eCBd542E7f1",
		"txn": "0x0c26acd2799d88c1fc17114c8a7ee39a0fa79d403cde94cb17bda6e518a466dd",
		"blockNumber": 3753196
	  },
	  "LiquidityCertificate": {
		"contractName": "LiquidityCertificate",
		"source": "LiquidityCertificate",
		"address": "0x5287d03553E736C732Bc71b20248954F2e70F714",
		"txn": "0xaab6893f262650c5839f33169f8a19b837563e64bc12825ecaa0573066e2eef4",
		"blockNumber": 3753198
	  },
	  "ShortCollateral": {
		"contractName": "ShortCollateral",
		"source": "ShortCollateral",
		"address": "0x0A68E15f8E289b9f1Ad1BCAD524FeA30C6125c2D",
		"txn": "0x94dd4deab413ce2bd6d9cb1947fddaa3c35e6865272165ad457cdbae10e75991",
		"blockNumber": 3753208
	  },
	  "OptionMarketViewer": {
		"contractName": "OptionMarketViewer",
		"source": "OptionMarketViewer",
		"address": "0x22c39cE1C3A49224Aea6D8c2AAa0019828E1b5E4",
		"txn": "0x541240f51164a4f355779fdcb7791d2e41a6d919a05ce2a3b902fb849b9d75a9",
		"blockNumber": 3753214
	  },
	  "PoolHedger": {
		"contractName": "PoolHedger",
		"source": "PoolHedger",
		"address": "0x307d49a289Df36023240dFA6746Fd8C181C2971C",
		"txn": "0xef6e1ee1ed82bec229a90890fe2d7f339384fce67638177d8593a1ec1fc3972e",
		"blockNumber": 3753220
	  },
	  "OptionMarketSafeSlippage": {
		"contractName": "OptionMarketSafeSlippage",
		"source": "OptionMarketSafeSlippage",
		"address": "0x644e589fd7d6AeC00dDC4212d17D4D71ce947b8a",
		"txn": "0xf4eab8da7902f4ec56286972a98fab53009171574806a5c627024c27289e4b25",
		"blockNumber": 3753224
	  }
	},
	"sSOL": {
	  "OptionMarket": {
		"contractName": "OptionMarket",
		"source": "OptionMarket",
		"address": "0xA3562CAc1c39f4D4166CE31005Fc080AB41120aC",
		"txn": "0x479bf06ee7df9ddab07540ef6cc2a90526a6f5ffae5c9f1dc7ca9c89842b737e",
		"blockNumber": 3626409
	  },
	  "OptionMarketPricer": {
		"contractName": "OptionMarketPricer",
		"source": "OptionMarketPricer",
		"address": "0xab8dF5c84c6aEEF86b241a91D86A11147b6240e8",
		"txn": "0x846a3aaf737301df6529cc99224d58da0a359ef6db6ec377083f03b2e06a960c",
		"blockNumber": 3626433
	  },
	  "OptionGreekCache": {
		"contractName": "OptionGreekCache",
		"source": "OptionGreekCache",
		"address": "0x5ADAcfE0F7A0a4E69ba5cD25f13c37cDFBa13E32",
		"txn": "0x4f79ebd5fec027a09dc2983a9ae7abe3681f7befb8cf080fe9c8234983b4cd46",
		"blockNumber": 3626470
	  },
	  "OptionToken": {
		"contractName": "OptionToken",
		"source": "OptionToken",
		"address": "0x3cCA343CBdDc9923a19212Ff5B3ee3f8D989DD47",
		"txn": "0x1ca8472171b1d4294f606241f5a8bdc1cc092bf1c07b671201cd243b9f5d384b",
		"blockNumber": 3626507
	  },
	  "LiquidityPool": {
		"contractName": "LiquidityPool",
		"source": "LiquidityPool",
		"address": "0xF492Cf8C61103752A31F182ebA372241e0DE6a1E",
		"txn": "0xa3081a761ff1555bf10b9186c5e59755aaf9321ddbc676c767ea1ea68b3071b7",
		"blockNumber": 3626538
	  },
	  "LiquidityCertificate": {
		"contractName": "LiquidityCertificate",
		"source": "LiquidityCertificate",
		"address": "0xb5ee566F3738355a59Bb5821fF7c7b5B956F0dAB",
		"txn": "0xd9ccbf844cd83cf50acb51ec3fe0069dd88b752c18071ed6f730f307ddac6300",
		"blockNumber": 3626566
	  },
	  "ShortCollateral": {
		"contractName": "ShortCollateral",
		"source": "ShortCollateral",
		"address": "0x8e1cA02b7147Aff9C964EB698c7c7d7Cd62a188f",
		"txn": "0xcf5128c94bc6089355a6e213f10cbce50ec8a51f6097dc4a9854e68d94b4210d",
		"blockNumber": 3626589
	  },
	  "OptionMarketViewer": {
		"contractName": "OptionMarketViewer",
		"source": "OptionMarketViewer",
		"address": "0xeAcB3f1a4f3137807f4e901167a999eF6f8312fb",
		"txn": "0xd99792fe4c976459a7ab83326aa073f2e3c10d66f59480f86bd4442ae15f505d",
		"blockNumber": 3626620
	  },
	  "PoolHedger": {
		"contractName": "PoolHedger",
		"source": "PoolHedger",
		"address": "0x46bcA9208B95886724b11799A41F9125EF7f057c",
		"txn": "0x9d1f9ab5e4d9c867ff0e51dc34063f758ba20d33b7653b0d076adce9f2051d95",
		"blockNumber": 3626656
	  },
	  "OptionMarketSafeSlippage": {
		"contractName": "OptionMarketSafeSlippage",
		"source": "OptionMarketSafeSlippage",
		"address": "0xF31932b99f5e21a55adBd23f90f09ea3E53a514B",
		"txn": "0x34aa1944947bf76dada9ca07b442f1622481b6fa4efd5799091ba55b028b31ba",
		"blockNumber": 3626671
	  }
	}
  }
},
*/
