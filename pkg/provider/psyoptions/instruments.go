// copied from https://github.com/mithraiclabs/PsyOptions-ts/blob/master/packages/market-meta/src/markets/mainnet.ts
package psyoptions

import (
	"fmt"
	"strconv"
	"time"
)

type instrument struct {
	PsyOptionsProgramID          string
	OptionMarketAddress          string //account
	OptionContractMintAddress    string //OptionMint
	OptionWriterTokenMintAddress string //WriterTokenMint
	QuoteAssetMint               string //QuoteAssetMint
	QuoteAssetPoolAddress        string //QuoteAssetPool
	UnderlyingAssetMint          string //UnderlyingAssetMint
	UnderlyingAssetPoolAddress   string //UnderlyingAssetPool
	UnderlyingAssetPerContract   string //UnderlyingAmountPerContract
	QuoteAssetPerContract        string //QuoteAmountPerContract
	SerumMarketAddress           string
	SerumProgramID               string
	Expiry                       int64 //ExpirationUnixTimestamp
}

// Expiration will be used in .Name() when fixed.
func (i instrument) expiration() string {
	seconds := i.Expiry
	expiryTime := time.Unix(seconds, 0).UTC()

	return expiryTime.Format("2006-01-02 15:04:05")
}

func (i instrument) asset() string {
	switch {
	case i.QuoteAssetMint == ETHAddress || i.UnderlyingAssetMint == ETHAddress:
		return "ETH"
	case i.QuoteAssetMint == BTCAddress || i.UnderlyingAssetMint == BTCAddress:
		return "BTC"
	case i.QuoteAssetMint == SOLAddress || i.UnderlyingAssetMint == SOLAddress:
		return "SOL"
	default:
		return "???"
	}
}

func (i instrument) contractSize() float64 {
	switch {
	case i.asset() == "ETH":
		return 0.1
	case i.asset() == "BTC":
		return 0.01

	default:
		return 1
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
			Expiry:                       1640995199,
			OptionMarketAddress:          "8VLQxtmr73fuudpgjmavsBFDUGbw5LdDzFDVwZZdxpqA",
			OptionContractMintAddress:    "6k9eKzLSgPuCgvNSWEwaW5jACEwYgssBC31bB2wS3MxL",
			OptionWriterTokenMintAddress: "Gy5BevUHf84X1iR4H2Gj7Wd98VgRgoquGd9TsmchoTPD",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "2ZeecYdPFqkqS5CihmMg6AqFpaUNshY19NDv9m5Mw8Tn",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "35c3Vb2hbVs89DxU7vfEpQ9bGTGCbuospQri9Y1SZnkp",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "500000000",
			SerumMarketAddress:           "7GwoKwaAPuTbsqbPL51BBpHU5FyW8nhmL3gkkDjak5pt",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "9n4Mcw81Q2TxpjZiACBpceDUELZYg74gu4frE2QXa1nb",
			OptionContractMintAddress:    "8YGUjop1EeNkb3kZkzCCy43Y4rUBMtUiKLKEJ8y6i6EY",
			OptionWriterTokenMintAddress: "2qsT87T534QmF8MPcGnW7gsrSYHUwNh2vsQyCVmymcAG",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "GVbWqc3odxKLrDg1JArdxqcpkuVFkUpnXoVYXvjYCYSD",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "H1aPdrzqbF52DBX6bBMkjP8Nz4LntVB7rX7RAUDFUQxX",
			UnderlyingAssetPerContract:   "500000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "wgGAQeX1F1oEn21u8qcxV7psd2topHdNxVq9Zy3ELbP",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "5rCMMWQtJmDUygCn2F2sVqqTgKz5fqxW9zdL7LBXyJbL",
			OptionContractMintAddress:    "F88rrVS6aNfgbMYV5qaZjpK4KpcExRVLWBUFT48MAvtE",
			OptionWriterTokenMintAddress: "7MSkUgQ5YCUDe9Ui35TiZGMK4Uo7UDoLJeovojG5RDnD",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "GXCuEzDiP7SYgN2isDFJhG8SrsTJSSnzfFY7BXeVwqho",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "HfARJH42vSNzo1N6UgpAonREZr3CVNDt8QLC8PShJcTu",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "550000000",
			SerumMarketAddress:           "CoHMVM7WodLGRGht3UTrjQ87hnDgjbQR1vqpmsf9K1zT",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "68Bp2jcpEiei2PEN6nuBtyjQQiAc4oVJpbzFoFVYwgfG",
			OptionContractMintAddress:    "8qixbZvyGi5JxSXhzA4kWEB5JUuyxT32fXwCzpsYry2y",
			OptionWriterTokenMintAddress: "HdRHh9E1FG8Ru6T9YQVLH9QJLukHSD4u3oWLvFfBxKJq",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "6HPjMtEwxAH6zW6oMyPGEa6eDuWLvZ5Hh15YMWEeWyjV",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "2PYrNQ3acj3pfHoBkBX6xuP77o2W7zG1w63x59ccV4n1",
			UnderlyingAssetPerContract:   "550000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "HBbwPms8ypCTEEqg4ReS8jn2e2b31Z8hNMQHsqYPpS4c",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "8u9nahJQ8XiqWeHY4yxRbtdtRM2e6zbU496jznhH62Vv",
			OptionContractMintAddress:    "9cRHFZhhsyFKo4MzMEVpXD47a8oEGzsKqzd3LMMRDVR9",
			OptionWriterTokenMintAddress: "8SU1kb2bDjMpJzXuQcwfmYgTRBqkpeH7h1jn1dUJUaoR",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "DULFvcDQ6kiyisDNjFgj4w7y9osWLn4tLy9wgsBM6TXN",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "BYAidjXSUHHMxxyTYuoaipGzVj72fwRCbMQJstCrYjvc",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "600000000",
			SerumMarketAddress:           "AE3yXyRP6aWrg1V43c2xcPXomBa7djirSFG8wjd11TdL",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "8g2ZWSMx3czpMoAitrhVUh2uJJwJJhrmJamrzap5oCYD",
			OptionContractMintAddress:    "76mv9uHAnDGFdrnhmAEY9tjionwoCf97RnzowyfM7dk9",
			OptionWriterTokenMintAddress: "66VzyPLKPSBX1Doq6A7v87nft3LaRqdLunW6SE45NSzc",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "4Q3TZ2e4iypQrtgTmUPjghESaPL5D1tjfv6AKkQKKo8w",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "4T2n7TJz2FR9tn6pDfS2Su5frCCV9Skev3fQijEvymtu",
			UnderlyingAssetPerContract:   "600000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "H5tjJKFdjxmerDbS88HuN4Lw7Qdy215DPS2casnpQvK5",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "2v26PK84W4tw4czHNXXUpwwVo9sz8PmK8X7tSztstyDv",
			OptionContractMintAddress:    "DbAyr6XuyreX84tonkgdDQmp6hTCLcBvTUX2i96P1wpH",
			OptionWriterTokenMintAddress: "CqzAmpTebjJcFK2P9AtAFYavS4QB2kioJ8YbXGHPdPn5",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "DonrpXbCPmo1CbHXLBZS9u4qrjwM2frU2UuajNLoUZB",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "5ekPNrySpdX8aTufo3nw88gPkR1eobJYRLpSvLmk6rKe",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "650000000",
			SerumMarketAddress:           "6hWM4BfrUafNRchCaA6j3KcpH92dGfrADrbXogyPgAia",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "312mBUANKFnyeCP7iSQEnyT9x7imsZn4qNC1Gav8y7hw",
			OptionContractMintAddress:    "3N5oBDmwJ98oXxBn7RWkn1EDgjeG758v567fp3buFTRy",
			OptionWriterTokenMintAddress: "9R2aNtFp9fDdFHrrUkfcmDAspxdtrTDZd6wo8wDVTp3T",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "GVKfA3mcDKbSQ6BsxBHtsKCLurk5dwSAsKZ4ckQFLjtN",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "BY5u2RfUnNZiAhyKSLwc3iU3NBUESZet465Wq77KCetZ",
			UnderlyingAssetPerContract:   "650000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "7sNeUCWpCX5qhvsM8dYnWzp8XZ8VYNZ5MkTkrnoLEAoL",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "8EXjnmWkJBFoHnnVGuUph4DNtPvBhFamuWf4RrProJtk",
			OptionContractMintAddress:    "75DArcM4iECP5SomDZtdGRK2hV5sYJjBDFkrhWmevRx",
			OptionWriterTokenMintAddress: "AESFQVgj5emBJGsitNs7HUtcGoPB6ajfJvYCMu8CYgik",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "HyD9btu54tJ6voXhYNBQiph9G27HE4BBnkWChLQwsaJi",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "D3JyZnad5Wbw8mYGZN4FNG1dwDumDHCMmg8tvTVU5Bn8",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "750000000",
			SerumMarketAddress:           "8NFuvnzJkmEJ3X1YFao6cP4pVS9vKoKmmkzttJCB51kF",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "DJKp6zoaWNgyZNu3zBv9ZFERVdcvBVH3g1GzH4PCKLcF",
			OptionContractMintAddress:    "EpNqJeKzHJVujfdHFvZuXzR5gYDbWfBeVZajXtpZn5Ho",
			OptionWriterTokenMintAddress: "2acWiankdobojt2JL99HyZiuw1A9P4MAWQfX62xZjufG",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "2TygpHEHMnv5i49ZQ4JeivRne61wwsexEQseDZDWb7Am",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "AurhS1cyfMmdNxNqiuMSN15s7QcsKkPRMZG22P3oMgpA",
			UnderlyingAssetPerContract:   "750000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "2jGG49JXHHUupWjncTSrHzUwF1rz4fEqL6fHUoPjDZFF",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "6VHHJ38vNSD4vx54CLDaBPw4F1fDrPLH8E9tA78aDVEW",
			OptionContractMintAddress:    "3UKZSGoBcPsUSDt4Q8T6abXky96s8kdtomgWJ3oScGzH",
			OptionWriterTokenMintAddress: "59tdmYAucpyE4XirNbqZ2SMA5QuAm5wuNuZRjkszruke",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "EU4LDzpmncEW5qpAcL66kpYjDzHGgVqLDCqe3Rdea1DK",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "9oeztGArp7fYrRYa9CYWhYvg4FQTDN7jDwA38Dnhzi5i",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "440000000",
			SerumMarketAddress:           "8hmRyvCJpU89mL3rfQor1mydSYXr6FWiqfpp5qdNkiFr",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "BF8c3uQ9M8yfryJC3Fgb72smY3hNoYSA8dg3Yf8YFePM",
			OptionContractMintAddress:    "GgST6tqbDHMmyrZicWE3Eh7Z8hhqftoE5icvikJgBjw3",
			OptionWriterTokenMintAddress: "HgAFizAPHQAVqs2hhk1KvN44Mqb4vFA9P7WM79kAbr7Z",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "5CQb7b5kPpsnU6q1cpZHbx6oEVXv2GfaTrA3BPvKutdJ",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "juX2bEW6yKwAmRWoTd2rjJHPSEzprXn9W3qCF3jCNua",
			UnderlyingAssetPerContract:   "440000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "J4oDasxWCvESuLRRNPBLzcaZvkmNZqBUteqNv28wEag",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "6DHhnnvrQD66Yvdi2KXV9Nra6GnpxsqrKM2YkjNCcEEZ",
			OptionContractMintAddress:    "5wq74YQQugZszJCng6V4HFeZuP72Ea9dx1aVsJytH1Kn",
			OptionWriterTokenMintAddress: "Hz2i6eQvS8wfca8tvCxCQAVsL3DeCwPcowkwgBjqais5",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "CSfWgK7ztN4L1wLehVWy526w2aMFpkVntyE7JRLXwKsw",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "9VHda7ecvWcrJSh73xYYGiiMwDDLNkYLHBGdU6cMksuq",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "420000000",
			SerumMarketAddress:           "DjuHNp6Y951UhAuiboCFSgeS4fYFAZNMjwU7RxFEwKrd",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "G4nETcA6u83uJzLpRUEZ2bLGdG3Ua9zpnYUR7LM5hwn",
			OptionContractMintAddress:    "9iabkQfN7MfQh6qtqKpxCfAaNpDcy3Yt5q4Wyvs7kaPk",
			OptionWriterTokenMintAddress: "9XZkskkk3MaJpa3PM8pLdqSA3a6ECh8ZneKccdEVe7A",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "513oT9LEe9wQnsFBpY7VHhGE3HGTQEexVYMUvdEpmE2T",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "JCw9x2uYAfhhDGp1ntJrtA24wMeayGkk9baVcDNVJWVo",
			UnderlyingAssetPerContract:   "420000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "BxBUBz5u6aym6sDSVL7cffFSvaMYmPhAtmoW4oW4fYxb",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "C3sg27SdPiasXUkZ29egQaeFqPhB5EvQ1qFuUANEnWNZ",
			OptionContractMintAddress:    "6osVC8PSW4EzFexWnzpHmJy8BjkafMQ1hq3tPW18XXQN",
			OptionWriterTokenMintAddress: "DX9XEfTCqELUyoZUqmGqt3LRyVkLsGsVRjVQ3yjEuHvG",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "2HUGYPZrWFGFRkyrefSu2Sx6MxZ1j6bUHTU1LhiWj4ML",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "38ZMUeqMHuhfyi2rhHKHwiVV8LbokpmkR7mqeAp2ouJW",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "380000000",
			SerumMarketAddress:           "3QYJghs6Rux2Xsxh4WJC36t4zU4wEHPN6mgFAwmza1hq",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "GgkxRde6upSiG9m3RU7N6bzaYw9w1tbbEeeatqecgFQP",
			OptionContractMintAddress:    "8H6i8imbzrhCv4uh71bJJeGv7KUCktvQK3hN8Gz1UX6i",
			OptionWriterTokenMintAddress: "EzUCRhSLSJ2ADDZTrbQr6KiDEQ2Gu9UytXnCoa6wsvwA",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "FHudJjoyKwkJzQZiZKFwYXAkfFxJQeMWWAsp1BhBEPyc",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "65B4GEKx62Dmpgs7bt6Ze49pPrWKtMJLTQpnq3z5C8EE",
			UnderlyingAssetPerContract:   "380000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "3jnTzcFY4XKMdEvKmxRdgbAvjBLrLP51bLogs3McX3oE",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "BtqHNeJmJPK36vC5AXJuaLFKgLnXXQywmaJo6NcaNgrU",
			OptionContractMintAddress:    "Gx1up5nJ1NhVNVQpfm5uA2EfY78Y2AvZccQTfKXgrjng",
			OptionWriterTokenMintAddress: "AQdvbqEtggWqtjCNNbHgH63xeDgwdmMTcykJWwC6Ai1K",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "9P4UtqhwvirwdD7NiqfTVhV2S5YUNpsxx43kZfU66vF1",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "9cC1dedW6FLfDWnfukk3UFoM2RwP9hVJp8aSn1qGGngh",
			UnderlyingAssetPerContract:   "400000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "4Xf6QVj9Pr8UAL4BfcvAmV45t7k79MDSwj1W2VxsNBE2",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "MUUxaJPHzZQvjFHJrCBP9HkipSktCj5Ck7NsbdebLsX",
			OptionContractMintAddress:    "EGFfTjye7NfHyPbRPzRLskFvGr9J2DVSsxtRevjHWUJz",
			OptionWriterTokenMintAddress: "7G9D6q3fq3qA7SgG6MPeCTGdVhkv4aGU8iywh44BfiP6",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "71oxK9E9zmxs6BNh1SagZ43yRh1w3G3NUU18XfHkB5LK",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "D8ZFGtzbq7aumZZDhumrzccwWgG6e5PCdi9bYBZnhSd8",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "540000000",
			SerumMarketAddress:           "8aXyy7DM2sTv16DeJr899K274PwTf4zwa1cQnQmypAYY",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1640995199,
			OptionMarketAddress:          "AR4oBuRGFN6hxUJxo1Cbpy9LP5DNAbfWNYzr4iVynmeX",
			OptionContractMintAddress:    "3gpH3fPWjWQXWdvuNsTQENJHbhxGph3NkTCRnMAqXA28",
			OptionWriterTokenMintAddress: "8yi84XMdpey65S9nyBUFuoqN5TH8coiBGCoU2wdAtEAf",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "EnQ7GJ6sR55ZUCN1no3Xi2dGEMjmH3jfzJVZwDvb6wF",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "DM7n619rHoe2KHptet3VfNWMYkkUmrYk7FuCsBkCq9L6",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "580000000",
			SerumMarketAddress:           "F98tArG6k5QhsfwHTSXBFuMXBoJxB1Mk2SNzA7jVQaxS",
			SerumProgramID:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramID:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
	}
}
