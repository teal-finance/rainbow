// copied from https://github.com/mithraiclabs/PsyOptions-ts/blob/master/packages/market-meta/src/markets/mainnet.ts
package psyoptions

import (
	"fmt"
	"strconv"
	"time"
)

type instrument struct {
	PsyOptionsProgramID          string
	OptionMarketAddress          string
	OptionContractMintAddress    string
	OptionWriterTokenMintAddress string
	QuoteAssetMint               string
	QuoteAssetPoolAddress        string
	UnderlyingAssetMint          string
	UnderlyingAssetPoolAddress   string
	UnderlyingAssetPerContract   string
	QuoteAssetPerContract        string
	SerumMarketAddress           string
	SerumProgramID               string
	Expiry                       int64
}

// Expiration will be used in .Name() when fixed.
func (i instrument) expiration() string {
	seconds := i.Expiry // 1000
	//ns := (i.Expiry % 1000) * 1000_000
	expiryTime := time.Unix(seconds, 0).UTC()

	return expiryTime.Format("2006-01-02 15:04:05")
}

func (i instrument) asset() string {
	switch {
	case i.QuoteAssetMint == ETHAddress || i.UnderlyingAssetMint == ETHAddress:
		return "ETH"
	case i.QuoteAssetMint == BTCAddress || i.UnderlyingAssetMint == BTCAddress:
		return "BTC"
	case i.QuoteAssetMint == SolAddress || i.UnderlyingAssetMint == SolAddress:
		return "SOL"
	default:
		return "???"
	}
}

func (i instrument) underlyingPerContract() float64 {
	f, _ := strconv.ParseFloat(i.UnderlyingAssetPerContract, 64)
	return f
}

func (i instrument) quotePerContract() float64 {
	f, _ := strconv.ParseFloat(i.QuoteAssetPerContract, 64)
	return f
}

func (i instrument) optionType() string {
	if i.underlyingPerContract() < i.quotePerContract() {
		return "CALL"
	}

	return "PUT"
}

func (i instrument) strike() float64 {
	if i.optionType() == "PUT" {
		return i.underlyingPerContract() / i.quotePerContract()
	}

	return i.quotePerContract() / i.underlyingPerContract()
}

func (i instrument) name() string {
	return i.asset() + "-" + i.expiration() + "-" +
		fmt.Sprintf("%.0f", i.strike()) + "-" + i.optionType()
}

//TODO change this quick and dirty way of fetching the available options from PsyOptions
func query() []instrument {
	return []instrument{
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "3qgRsDezeVXJySFrrXAyfkoG9ryfmJMzpGPAGKoNWv84",
			OptionContractMintAddress:    "5z5vLPghW73todzU6RJmApur46R5d2LGogGLHC2p2dz7",
			OptionWriterTokenMintAddress: "2oFpHVpims3B353YqGKCi6JgsM2KNJnt2w6cX5KVX5cX",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "544UcTLdmevdTmUccnsZVfPF9HYxYNXGt1yY6mrMAorr",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "5Zz7nxkpWKTPp7y2C6G1poYRdjfavzsUWAksTKSQa6iP",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "1000000000",
			SerumMarketAddress:           "E3gsiVi4hENPeopDnZshLmWGFhJED1e8ZZviosAcCnxs",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "DEkLZtRssnek68216YBzzUNJfvqraNhxmaacH8BzTGg2",
			OptionContractMintAddress:    "4AjpsqEUhVVt2rbLADRDx48SKrCW2ZRdh4sa6e4DYTEH",
			OptionWriterTokenMintAddress: "3gvWeWtoxGkxPnGJM96QkSuS5jznmSCqWr1cqSz6gEh4",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "HMF5mB3bPtESfZXM1AWJs1Fxf3958hEpANw3fFQeK6Ch",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "3u6C6xKfvnc3LC7DUzTAiC2AsbibpZZY7vCH2BoN1K3f",
			UnderlyingAssetPerContract:   "400000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "Hae2tgrU7MqYNEz8zFzPshLvZkDgg5UCWrEaNFEgDiU8",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "8iRVx5MAjeZi9LrCFeCVqFVBvWefpJRQA1Tdshh6TMUj",
			OptionContractMintAddress:    "9KqRdq8iAHQswQpjUE29eMZxtqtQYK7Fh1aqvbihcFhK",
			OptionWriterTokenMintAddress: "5w7q4U3bDTtYZcFrxUXwkZMF65Lz1Xjt7yDXBejUSicJ",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "7MNKEfEWY6eLyYps6FTJEQUkhpR1S4Pux2FVMg1N3zGm",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "Du1XMevU2qJTqYZ7LdwtjAqDFoSiKu5FVyk4MgtSriGk",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "400000000",
			SerumMarketAddress:           "6SWFKdSNehwM3fcKXtxW2x2w9SQLnw4DZV2YRRH1J8qS",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "G7KfX621rnf3kcX2ySHd3puRN8Brjvjs5hBTa3ioxBYP",
			OptionContractMintAddress:    "CpWxXk1FWZCDKGWHVyA6mX3FMEQoDiCGwKbRQ1wGgDvH",
			OptionWriterTokenMintAddress: "6Y2xBTuBE9wtHbfhdg36g9rz6pFXD3v56GE1JrwJoGLq",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "Fp8LwTWbLAfpj6UpWr1cnvLVCtSXCR36BQRvrjwJy1dA",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "ByzXre3RJupnnXppYmM3oMVuTGwD2H6UfKaXxThqdWK8",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "500000000",
			SerumMarketAddress:           "5yeUkBisDGdgkBTnuh9mrqiwDnGfSm4D1hxkoBvG219j",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "EGXBpThy8RV73bhcBcpXfmiBtKWkr4JBZoqqHCkuiCGE",
			OptionContractMintAddress:    "4EuHYkYXVZdaiNPBk2fJfXt2FVt4382BwRozCDaBPgqF",
			OptionWriterTokenMintAddress: "F1DEQwHVq7LZ2FxNoNzKhrErF4rK7YvDEjq4HYT4Lbdn",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "EGuDhpAK9SC35S3ZsoHMHAmTLPVv76SEqD1uAKL9H65E",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "HenETiyxPzHRdzjGDNjRhgdLfYoKkf5CyKtb6wE6tzCo",
			UnderlyingAssetPerContract:   "300000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "DmeRtx1A44e98aUW8eMbYssMxLfGUo1ZYs3mWHCJSF62",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "4ku7GoJVvfpeeAhYJPcHHwrfrUtGt8nBASEbkiWU9zdU",
			OptionContractMintAddress:    "9AXUzoFLHxsNz7JLA5oatsByPeXkNmR2ztRwFAf6Bweg",
			OptionWriterTokenMintAddress: "BJmRpGwrrbCSNHtLrggkB4CWBEkwWx6Vnv2sgX1PtCo4",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "HcUqc7KJTm9R2QfrJXALpF1rVEL97SNZkvqKfszq27j3",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "BwFKy8zpQrnUvzr7zG76PW1E3zAXMMGp97Mtc2JRuTEJ",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "700000000",
			SerumMarketAddress:           "FExv4h2tusEbVnTYT7d1WmZLcZmewcpHMDcQxDjWfQG",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "6dAdYCYdbsQWtwXNSA1wh5JmSoBjsHrFqPVnDiGYaULn",
			OptionContractMintAddress:    "DS3RF2zR4h73KmxUQ1rsfeAZQrirxTLykFYWz4uTpuyy",
			OptionWriterTokenMintAddress: "CXg8hFUbyTqZWdqTths9MwTvMDQjMCyzCXKMGYqRtxgb",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "Ey2EAv9wQc91T3cUFcK8Ha1xFVqBwD6LAU3hKt9LrFkF",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "SsXFz1tV8bHHt22HnDMM7HjHfvjzUBac3YpkoUm4jcj",
			UnderlyingAssetPerContract:   "700000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "537sXyxpkszLe9GzHDBScrGxQ9RaNwZP2L1Khr1MnY25",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "D7T6ec62KQGbcTWkSeLR7HHc84PLmyVHBfr6Vm4QTi2N",
			OptionContractMintAddress:    "5oPK6i3Jegw2Li1VAYvGpfaipN3qnK2pVxAHCmB94t38",
			OptionWriterTokenMintAddress: "6bWyycvytbYQE2APeXMT19kzzUatDLXyhRu2nwHFV5KE",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "Bpyo1as38BtLY5xoYdBEs8ZsTUyi4BvHmrH6eF3Scu1p",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "DCHway8H5NM2yHQTQkDcjsHemswX5QdMuhB7hQCusDcR",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "750000000",
			SerumMarketAddress:           "DvwGfJDQSaRvX4YnN7XdDK5YAGPkiybRaDbAyEQXKCkD",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "HeCNKQDddRLdTLG8JASy8fmWfSdd5L9b64xo7c5f9TNc",
			OptionContractMintAddress:    "C6tnWm67qN39XrufVhczzh7X1UajNh2H3JAjJkAmDwbG",
			OptionWriterTokenMintAddress: "HVrNCdPrNSdJpVTvkb1AQwgqWeFGoK6rgYcRooLtgq6o",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "HKNRQq95MNX1Vsv2TnTkf9KmXTFCPUPBteg4JhtWDaUp",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "AL44a55kGpauA72wXa1GGGPRp47xtoVHk5sGhCbeGnSk",
			UnderlyingAssetPerContract:   "750000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "6zimXjHYqXvQLLFSzeCvZ5VUnzVjxhG1mSjoxfMacxoE",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "ZGGSFN9XJVMpQbtSdBjiU9eNUaM4bsLKLiEQy8bHy8B",
			OptionContractMintAddress:    "2rjtaKWBWnfuNCBEY8HobiEUJouUWWB1BpcFhw1oHanM",
			OptionWriterTokenMintAddress: "DpS23Jv1nPAqaXAMoqQUchprp4U8zZ1g8xHbiycP2Dnj",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "CSFF7cp8p6kttis9sQE2rBfcsdwQNSMcPBUoPPLuv6MR",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "CcNKRbvLM12Fn7HcHPMEivdhiPcQdC9HnNZL2Ni9ricE",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "650000000",
			SerumMarketAddress:           "6KYXnVR7JaAnrbWvvaid6djHMicux4S6T8Kw1oxUJUAK",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "38SgygsWDperYtHAnpMLB9dhdq6fJXxxVhQfUKReU3BP",
			OptionContractMintAddress:    "DwGAj2eL1WLjVf8mW8qcaQRZsHdhFWdJDp83Y6aukurr",
			OptionWriterTokenMintAddress: "GBaqg7ahoVZFz2Eag9aaLYcgTv7AQDiwEo8e9pWzSkLj",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "AxZk1RHLnRUtAN4i8F7iwxaH9hZqgHcX7BfvFLESaEf3",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "54Areco63ZC8Wy9GPo3h6nreaM4rWWt56P1yDsPdteGr",
			UnderlyingAssetPerContract:   "650000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "DDifKDGDjTGDSQvprTmdzow7AXi8K6YR3DXySKXNSYec",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "6h5uWwrdyAZQXj62s32SbgNeTngFxLagGBpXJ29xfRkV",
			OptionContractMintAddress:    "BCyqG8Bvvr6bLMzo2W2XdEq1aGRJmFzdePBLzD6skxJT",
			OptionWriterTokenMintAddress: "48GvJACamQ1T5rLEt9GDNk67kLwijm7UTemmmQuoPJJH",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "6bGwWByo43ikiw7LVLi97pbYx85B7GS7q5BnLYqp5o4z",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "Ex8qmTiyEFNUoKwzp3Zz7rubi6X21qV5fn1RDgssAYu8",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "600000000",
			SerumMarketAddress:           "CbMMmBaA2wib8ChRdrZR1NbBvZqna9tQYBz7NryjXK3Y",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "4cjAgQRkRK88GufrVryKCRhAcEEMqF65BNztybVEtAh6",
			OptionContractMintAddress:    "AqPL4BcMSH4ij6WnLkCWqKWsJtBXqdGFyCkDLVeyB65N",
			OptionWriterTokenMintAddress: "5ZTPVhKxi8agDRVrnuPdnM6ALyHJpP1ccd9zaLvQiu1W",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "FrH4PeP8Y9r75CZQByaMneRK4ptYnpwY13qyTns3vqvs",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "4p16gTw9MppmSbUbMDB7b5F8kYzee8FV24pmoeejAKTo",
			UnderlyingAssetPerContract:   "600000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "3L7YsuT8r7MfpMgvbDM4iLVQ6zgFyDku6ReX2PeYoB4h",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "AJTb2v3VHfgse1k5tEEUKNxLEYiaYw9Ef7VttgpTsyj4",
			OptionContractMintAddress:    "AYScneU5HVUAgoQQuwToL6bDSajnugGqoZ9cPgYAKRCT",
			OptionWriterTokenMintAddress: "9CBCR62o2u3s33sPaiEXz2VuPiNB3Nhoi98wG7EqJncf",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "Gt7C1AGcXjxdsdLSJ7AJ3Zy35HcFSipVyF3oNGtgzbpC",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "AnhyZBkPAeuuXiTWdpHpLKXKxrY3wGoySNFYDYLYvi17",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "550000000",
			SerumMarketAddress:           "woJLyMzipfemVeCdCFbL2a4W5FDWQSHg6QhMDJcBtGV",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "EdaCKVVH9JExPLSX5t7goFUGX2mJUMkGFDSRbq8GffvF",
			OptionContractMintAddress:    "7aBjstKweHyrDd8aqbAdCaDeigW95R1gfRQC5roWwMrG",
			OptionWriterTokenMintAddress: "Bnpw6BhTnPK3jRHFuzJt6Hm2DVTai3Nq1ABMe98qvucY",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "EQQmvr5NJfvGxwKxYiftaKzGU5ZVyfqzFwbxkqcdehtR",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "B5QYTKRyd5vcS9BNvaeE8ud8oDsWbxFdKFanPPAm4VsY",
			UnderlyingAssetPerContract:   "550000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "BdG46sZiJhUMTy4r1c43TrF4BuXvuR8urJDP1GSSfqJ6",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "FRQnEjv9NRFBuxkt1rwKahy6cyJRii2d8pVUxWvNJKdE",
			OptionContractMintAddress:    "BbDpE81YzpQUvdnEHUxSEANhg1TL56ZKD7Mnz1GgFQtN",
			OptionWriterTokenMintAddress: "5QaotgPZJjkgdCBLbCcjV55zkP1K1wAymQecwUTvRa8F",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "3tPF5qFoVrkdLUpC51XEeoRk6Gnmuf4wGxzH2rKkkW9r",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "6CZXoyJh9eLtEKrDHoEYdajzAgaqvsdQr6hXYcJFGAr3",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "500000000",
			SerumMarketAddress:           "Ehjv9gX9EyfKZ3ELhEzTsCvAmU6hjAH3vWnjyXFZrJaK",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "CLpVkAWwzRyqRjnr31FU7sJfcQNEFKYGpmacXAha3U65",
			OptionContractMintAddress:    "BuFx3hzs2MG97m9F4ZxWNbgCWBf6rBje9o2UiohrXE2",
			OptionWriterTokenMintAddress: "C1RHu2xPcmNvVUShhDdNX5Vo6iyW6aW7UxeR2R8s4LYc",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "5TZe3GstuC2gq1oo8kV5fvFfDtBUvq8MnKK3BDGHG5Ye",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "A9BUBMFVUvXrKPu91PggmP2mSiKikq7yy4m7JLXix31D",
			UnderlyingAssetPerContract:   "500000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "5mRRX42yk4bek1L3wCqkKRxRm8GZYwRBjX9gmjTL4sB9",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "Cqq6ovWRLvQAzDU3z6TiXuZ4mb2g2FLCGhNjAtmdctuJ",
			OptionContractMintAddress:    "fir64pZFFsEUpwGCGgWFivAgtajbMMQZH1rj47SuaEe",
			OptionWriterTokenMintAddress: "6W945wyrYzMS7aQW8nN2VvTfwVGKGBS32cC5KuYJZpUj",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "4xV27MU5mSwuujGNvG2X611dZMR8uLgUh8q6iLL8H9zx",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "7VHoEgRqY43P9mXZabMnCGC51VY8pExvL7n2Tt26PjTC",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "400000000",
			SerumMarketAddress:           "4FLEMKVUDbXEkac64VjnbcWthQZhDPQBipCJoofKsbx7",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "GWUVAHrRqVswdmHG8af6rzJBLEnA5kizcBzsv1L2yjcm",
			OptionContractMintAddress:    "83xeKmCwchBNBzRvxM5BcQLFfNg1CEH6AycXMxBf4hq9",
			OptionWriterTokenMintAddress: "AeMsvE2Xqu6DyHddKYwVCXpTC8nkinTiopyopEtqVFL2",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "A8bFfmvPL6e4mUhWQqhZ4CoRyoAc1U3MaUCfJHEyw4tG",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "Fhedj7cW2NNfBgUejkjBpi2Bu2L1UNnDVyeZQVZn4GVi",
			UnderlyingAssetPerContract:   "400000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "4kUbyDDnaXxUdWYgZSJfnqg7p7Ws9RentEVJGX3WJNVb",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "81jLJH3uwoQN7wEX2pK4iX8UyjKcdjfb2nv4zXVwByx8",
			OptionContractMintAddress:    "FziBTAq9wDTCeBwtYYpqWojPthr2KRP1ox3epJwRuysP",
			OptionWriterTokenMintAddress: "9KVT4NibKgCszyLBvmbrJWgaQs5v1RehdTFmHmXpok5w",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "7bmYRhtfEpgHCLiusTjW83pDSKXZBYywWFsiu2geGs2k",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "DwBMcf92rsknBt6Kor4K5hXUjG9PQfJdwXGhQSh3dfna",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "440000000",
			SerumMarketAddress:           "FY5JEpJzaTbJ1WVUb7WessPq9YadM7pFhxVqyTpZMCHr",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "BR2zKHaXjF2YkMkb9ZLJhv7VRj3hL8fUBjTjrDTNgy7a",
			OptionContractMintAddress:    "GcafgfvzEkR49FnFN7JKZkfz2xFACEqxJbUG9gmKdrRJ",
			OptionWriterTokenMintAddress: "DjVZxQMSyAs49dDkEWonL1oRLQhnxEwvmwnSmnnmgAZL",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "CY8dGmHiyBJZePzjVGiQUQcUDaQzVoE7tKJaM17Kydoo",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "AdmRUmwpfG67a8WQhXPXHYfGXdccibYkVqojU3oRd9XU",
			UnderlyingAssetPerContract:   "440000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "CUryqJxit554umaAJnKM4BrKQGvPEV1skZLHTeKroRV9",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "ByzcijnLMGxWgLLvR7Lxh9nULe3SMh1i1WX9AYZ1gUsh",
			OptionContractMintAddress:    "CZrF12xWaRdfqQvGsuQUNcXdQxCo2i3ehxAzeTrwi5LP",
			OptionWriterTokenMintAddress: "FMG5uuaYg53tuo2FREFkP9KX6tEpzwEUVeSk2ek5VPHd",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "8hLh3ipzd1VWLhchD86sD6WRANGQfAY9sTWwfuprt5Rd",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "48u3a27rnGKCChXiDF3S1GJ4e6P75pnq3nnJ4Xy23ecT",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "400000000",
			SerumMarketAddress:           "VDCCBnzurodzhp8XoUxd8c4rTcKEHxwGxuTo7WE56Gx",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "2diYqAy3sxBEtbN6idLjXB68aGxbi3quKGGTFxtPt7Gn",
			OptionContractMintAddress:    "HTWaBXDntvDmx4Y2ZqVYE1Q8jAeksNz8YqjhUcrYxMLH",
			OptionWriterTokenMintAddress: "EtHwYhzUjeLToGS9HxKnRSW6v7Lfu3hAFa8Nr6WojPUQ",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "GYZuUpe3aLdx2iNfdbYZ7aVNPUSBnC6mHFexV6JNENeZ",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "DMKb3i2KCdf9H5fXdw3NJcAk8gZZp7Ccxt5V9njWrk2z",
			UnderlyingAssetPerContract:   "400000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "HUrJeidenYw6o3LVB6whWqGbZ8Ye6yP5BiHWzpoWJN9Q",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "VTELwTYK8JYckp7XoYiuLooWVaErcnBsGTk4C9WgEMy",
			OptionContractMintAddress:    "46sTkE4jKqJdYvmKbWXctkcepUXW5RpmatuXbTRWZLbA",
			OptionWriterTokenMintAddress: "GyGc7hX9fdBxrzHtKUQqKBRR6iK6YbMPULfqTBfbS5ME",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "AsY94kVe3yspES7vqgEM3UaPuATfXACyzZPdWb9LfHf",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "7Bz7bQGmxbFjpcDjaANi5mumVwRE9m43JGaBzetfLmHd",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "360000000",
			SerumMarketAddress:           "8Qi3FduzyiSmLgM9UkwDdFQhCmfByXgASET1843jKCFa",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "4L2xm4RZyLUEKG3LHZjQyjXFdizyFRX6cYeZmNgVMMwh",
			OptionContractMintAddress:    "7298moehJuHYWVEL4AZSkadvzdJ5r8VtH6iUULkzQXdo",
			OptionWriterTokenMintAddress: "7o3vBryfZ7rgEGsrp5fd4nEDkvEbF7ANvJhY7dXfHYum",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "BwnXoXfrwN2VZFfmtZjWDRwQeEQMTPnuu1dC5pYH6fjc",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "11UXz53wGkrH6o7UvLjZ5cqKiAKVf4o847aCz8k9tRT",
			UnderlyingAssetPerContract:   "360000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "5czShbWeXjxVmVa9t113CYfeaFCPDF6ziSXAJE6qhh8S",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "979FDxC2aiuNjR53Y5ii3M12xHdxsjtNDBjyeZxjhsTt",
			OptionContractMintAddress:    "BCnqbbKqYHrFUxr5dqPSVnGK7ifT9qJeWdGgAKjKh1QD",
			OptionWriterTokenMintAddress: "9ZM6ffB5VAAR9LBbvPmd9foKXsFZdBdinCHRcY1kezfe",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "SZ8mSFa4dbfR6GfKJ4bAFzQmdudf29Wr769iV7gv1Ff",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "AfzAYHSXn3iy4YBbxZYJC6PV7Ks2zAorMPwZsgVGR8jt",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "480000000",
			SerumMarketAddress:           "D8kcKtXdjLuKxKppRpRxFBDHt57abg9thFqLMdPYa5f5",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1637971199,
			OptionMarketAddress:          "DZHFA1QNmJfGGGxtmHbkUwyNDn5v2yxZGbzmftLv7221",
			OptionContractMintAddress:    "H7V65inkyWy98ATJktbeZ4tC4nKEPguSXEettU4dj9mu",
			OptionWriterTokenMintAddress: "BmnbEBk216pPNGRB6WtF53u6jnqp1t2AqQ4d32Do2d1M",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "B6c4oNDvAeCeJYRXAAyZ8zDRZ7j148RufJBFJTRfRdFx",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "ByLraDaivHECKuTqCRPBXtcZQdQHvdYUNQLY7LMSbeec",
			UnderlyingAssetPerContract:   "480000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "5TRT2ihjfZL9UAQ1touTpffQR4HujnZU1YjTouZXNa4R",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
	}
}
