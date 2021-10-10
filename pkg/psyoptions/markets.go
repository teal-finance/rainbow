//copied from https://github.com/mithraiclabs/psyoptions-ts/blob/master/packages/market-meta/src/markets/mainnet.ts
package psyoptions

import (
	"fmt"
	"strconv"
	"time"
)

type PsyOptions struct {
	Expiry                       int64
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
	SerumProgramId               string
	PsyOptionsProgramId          string
}

//TODO somehting is wrong here
//use in .Name() when fixed
func (psy PsyOptions) Expiration() string {
	seconds := psy.Expiry / 1000
	ns := (psy.Expiry % 1000) * 1000_000
	expiryTime := time.Unix(seconds, ns).UTC()
	return expiryTime.Format("2006-01-02 15:04:05")

}

func (psy PsyOptions) Asset() string {
	var asset string
	if psy.QuoteAssetMint == ETHAddress || psy.UnderlyingAssetMint == ETHAddress {
		asset = "ETH"
	} else if psy.QuoteAssetMint == BTCAddress || psy.UnderlyingAssetMint == BTCAddress {
		asset = "BTC"
	} else if psy.QuoteAssetMint == SolAddress || psy.UnderlyingAssetMint == SolAddress {
		asset = "SOL"
	}
	return asset

}

func (psy PsyOptions) UnderlyingPerContract() float64 {
	f, _ := strconv.ParseFloat(psy.UnderlyingAssetPerContract, 64)
	return f
}

func (psy PsyOptions) QuotePerContract() float64 {
	f, _ := strconv.ParseFloat(psy.QuoteAssetPerContract, 64)
	return f
}

func (psy PsyOptions) Type() string {
	if psy.UnderlyingPerContract() < psy.QuotePerContract() {
		return "CALL"
	}
	return "PUT"
}

func (psy PsyOptions) Strike() float64 {
	if psy.Type() == "PUT" {
		return psy.UnderlyingPerContract() / psy.QuotePerContract()

	}
	return psy.QuotePerContract() / psy.UnderlyingPerContract()

}

func (psy PsyOptions) Name() string {
	return psy.Asset() + "-" + Expiration + "-" +
		fmt.Sprintf("%.0f", psy.Strike()) + "-" + psy.Type()

}

func GetInstruments() []PsyOptions {
	return []PsyOptions{
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "3cPMm4y392xaqGfeCuRrB7NRFkWiouFAX2q7cMUWVfY8",
			OptionContractMintAddress:    "8n7QxHXPPyTQMZZe3cYFrc2PACDHQENAw5sS1H1rxuwV",
			OptionWriterTokenMintAddress: "65FZSMhwiEpo27AYDnKWoFSvbg54pUNbvgDCL91aeeNG",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "9EJojnVJFJSyAW3x2Ets6oL429P5htKHcXKyaL69n6rr",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "BXeBkjGUC5TwB1UKX68goN5J9fzzEW8LdusHxpDdPBRi",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "280000000",
			SerumMarketAddress:           "8A493gU55NfS4fCjDoLAiN57zPzWf6QQw31QQf1fd6iX",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "EZgszJyCFLdEGdTuvGTeiFk4imkSQsLYtTgRwsXUoX4p",
			OptionContractMintAddress:    "GTtz1AsiYSsQAVsWoHQRvfLTEUtNzzTodeR87f6mPoTj",
			OptionWriterTokenMintAddress: "8hEDn4c4zdFvqhv5bJaNhcoFgQ47YieLumGfC3GK2rEc",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "AGw29ArWvSFgXPPUkieBmemvj1XdbLNFDj6EuivfGC9q",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "BXCrnE4wK6TG7C6D6YJarYqMnncqDVDk2CmQW4qFp5uC",
			UnderlyingAssetPerContract:   "280000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "5pHcU2Gz8eCMwynLvz1AHSFoFbKkeTc7ufeqeG4spb99",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "Ew28EnQ5BD9sype1Eg6usZdLsvkM81UaZRyAxy8UtVpX",
			OptionContractMintAddress:    "6maaSivtLFzT41fVnzRmhTjkU4n7DzHw3ksY6dTXdoSr",
			OptionWriterTokenMintAddress: "ASEyfbJLW7Yc97ULds3UTRFpEviyt9JKns6aPoweD5Tr",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "5bM615Fu3t6VavXNAcigT3AukxX6ffp64sVxqX3H8int",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "BARZtXUFodyPyVP8bW1gU3FCimpkwjZLDTmKn4cBfpYc",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "300000000",
			SerumMarketAddress:           "8qgAVVE6eeJ9u32LvJupgW6NyNWPPFFK4xnXfGtDNeP4",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "CRrJjLfPEhkEymD5sv5MWbwCJjXxsyxi6ovziVtqcoJJ",
			OptionContractMintAddress:    "3bLXfvHC5Ux5k55AW3tUuTTSQqSDCDwGB4QuS2cj7L6o",
			OptionWriterTokenMintAddress: "8iGjcCQPnWV4FJfUdyLYVeJYud74t4s2LRAkQ4BBqxev",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "58Z4N3MUiNMdUPnpstNYJ5KaUYGe3DiKaeaMRAscTpGS",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "3e1A9NyXxQU9BDcc81PWtCK8iDVDHRFM7jwjrawWD3Qb",
			UnderlyingAssetPerContract:   "300000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "FfkVR6ha6N9fcefGxfDFP97xL54Pc5ipiwnt1uuTpuyX",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "2wZAKGgMiy8awywB1Z1Nm5oar1vNyTnwLC5yeLVEqZHk",
			OptionContractMintAddress:    "7HrTmG3fsWXrGcEFBBZGYxSa1dLedxrWr7X67LxJcjok",
			OptionWriterTokenMintAddress: "8pwAV62w3WXkkaCYsZnRg3HNUr88nRCkD2DThWmdp6Uq",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "GpdkGLFFGvXwkh8AzrmVhxc7G6oiZXFgB4LujGEW786M",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "2fXDoVGgKtw99g3ujpGTWew3SfZfZpmF7fFnxFf7JUgC",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "320000000",
			SerumMarketAddress:           "TLAzx53rb6pEDT3oWitTzg6YNe7BTERLDzujMdU8RQx",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "951sCDzV9axYL2LgwMew9oPeoD9mrn4LKphdpUq1XnmC",
			OptionContractMintAddress:    "JACMVE1d5uYY3pzaZf8yXg386iBZtJjoKxcPKVh2SeJu",
			OptionWriterTokenMintAddress: "7JdapxoG1rGn7Cmqf9KJZeA8LjqQjWyDEh5wGN63P2Nd",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "wnY3e7DoJHnQAvG4BLJ5bb5HDjMMFaF82TYQUEto4Vg",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "BUn94ujQPjCFYcEkcgd7tMT3dqDod9Paaz6mMss2kiBd",
			UnderlyingAssetPerContract:   "320000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "9A1m1JXgnMrdq3JkqEYx9rK3CcYm9EHDLeA6sn1hH3PP",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "6f7pgiDZWDGHwsALNCrLyTF7pktsPjsBrwQDfd5VcZSu",
			OptionContractMintAddress:    "4WUsmSARXSisFttodmkhfhDV1CQNsZafJU98CNoEv65X",
			OptionWriterTokenMintAddress: "Q1NtXFkhZewErFb9N6g5TVwfgUQ8j9JzPCCKWKiT5PK",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "4RFbjqWAcgGm3hCmSuQAHpXLwcmR6QVHKJVJqJrfHF1K",
			UnderlyingAssetMint:          "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			UnderlyingAssetPoolAddress:   "AwbPC9E7sSpXKP2AY7iNueizCopzsM6nSC9Nu1s1NBT",
			UnderlyingAssetPerContract:   "100000",
			QuoteAssetPerContract:        "340000000",
			SerumMarketAddress:           "HYmPvo8szh62QVaAfUAXR1eppvCfouUPpH68yE87UYmy",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "GrDjAfpwjhApGzx5dhrrrzwkNdvZeFXNfpuuETLdAuru",
			OptionContractMintAddress:    "HAbip5VkPKdpk21zZWYSEWWuX7B18oyPc26PWtUPjmux",
			OptionWriterTokenMintAddress: "2VxhpjJwZ9wWBMTgHqJ2r5Rzk2n6NT2NqV6X6hTGbmiG",
			QuoteAssetMint:               "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
			QuoteAssetPoolAddress:        "4rPXpjQVs4QSQqsxRTuf91fjREULiqhDrFPNsGJ4yAr6",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "2wzWGWtXiXaHNQvztm93gffe62jLtYhdZQK6EUmkdK8X",
			UnderlyingAssetPerContract:   "340000000",
			QuoteAssetPerContract:        "100000",
			SerumMarketAddress:           "8fFcWuVaZSKoge4DCpMcNrR5nNXF2pbXCfBUxkMomgr5",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "ETETQvUURPZfpbeodXDCUXBUw3XMAJAA61Fv85w2cxGG",
			OptionContractMintAddress:    "2TeierUVapxnyUvTVrgbHBZJTdtANRKa6fqZk3b4FYJw",
			OptionWriterTokenMintAddress: "8kJSxckTRvH4eMD9Dc2GKxMnynNmtdazE74m45ULvzjG",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "A7XKN4EJ1i7XjTcXKCRrzpvGFvHmUUVvhtH7LYumGxSr",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "29UEL4TZBo8YYbTVLiPgPZbJe4Uxxtb55YU3NV8iQFnt",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "350000000",
			SerumMarketAddress:           "8fhiAYm41RwtiX8WusCSpY617GWPt2LwUnCQcEeer78o",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "upb5wZVyZX8M7TZAss6S8vT3W4zHjNz8LGPqMtvoCYB",
			OptionContractMintAddress:    "AzfCsVHMHxL838YLoWJN2fqF1H8Fdj9efRVwrtxefLxP",
			OptionWriterTokenMintAddress: "5CHyeJ2p4gqkNtFrALz33GtsxE5V1faQBX2LMmdyQ6tb",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "FgtNzK8hopwGPEnrC3s39Ebt1awXQW7PcuDiAuP2AexG",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "67d6CPUMyDnm4CQCR78pd59Fo4syKiFpEF3UVPeaNmVM",
			UnderlyingAssetPerContract:   "350000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "6at26sVk8vTYtLh4YDKvje4PDdgFJsNHHyoGw87WNszP",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "DJeYEmuEGNK6cyf5zbNdVZSoV5zkYe6SgtS6JmQUfUDV",
			OptionContractMintAddress:    "67gUUopxFs3QfUYWFWKLYjVnAE6K93ArGzZnHYztdXb4",
			OptionWriterTokenMintAddress: "41EhYmbQRqeDHBfrPBNW9QPhfgSvxesitHq8vjdAU9sh",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "CpZk891zHZJKWG9MpCfBC2EQc3HrLCGcqRqNGENu4d41",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "DBogQr9Vwtw18Yt6Ndjpzfy3tEJfuiVDHJafMxhqdcem",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "450000000",
			SerumMarketAddress:           "9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "4JRPAdVP9C7Xszvd6ECuPy4z8faS4sKEMeosKQY8JSZN",
			OptionContractMintAddress:    "BzdCwaPhwzizkbt51DvWXKzHuF8rCUhzfMP7HaMeTSHG",
			OptionWriterTokenMintAddress: "BTqXz8fKaQqMCkoeEkS1yi8azGhpSHSS7m3NrNLGZcQP",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "4jZgrnkGef2ciEYUHu4uk6FhRjneTEjxFSbYXX5Ygkyw",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "5smGQwvbazFbEMNgWE34gXzBmjiqCGdua92vKk9fwecQ",
			UnderlyingAssetPerContract:   "450000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "FbuxXjSTTfhbCNjBfJJKqxkNHtMoTfVBSE7U6BGiZeJz",
			OptionContractMintAddress:    "AgmYhci5Gyd3CLyGQaMZBkvnpZpFbQUt9HoSF9EDkxMK",
			OptionWriterTokenMintAddress: "J1XmmQVgT3VH9ugfHCsrta9gBuWBprqb3qNFYe9g7XD9",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "37ya2iMVrPqCfSDhpwUsG4KE69jvVASb6wEuxugVtpce",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "A19W56Cc93SM4Ryz273qaNJRzXMn9xjSCT7DM2k243Lr",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "400000000",
			SerumMarketAddress:           "2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "BBaD7X1SvZGJwHVGjBqLdrWyDjoaazpoY3J9uqoRaNHs",
			OptionContractMintAddress:    "58yTy6bYR866Z4n8Dqvu6dwjip8Au7zpVSKL2Bq1YexM",
			OptionWriterTokenMintAddress: "3kgPUUUyRrk9aEKxkijV2reNwkxAr6AsX83DigtpEqx3",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "9bcYkafUf62bYZmRSQ1zcdzSkUmoH1aTMNKhtkSViinr",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "HdhxnVwPxDNKYsYfnxP2DeFKMYzGhtAuFm8LDULzGxXK",
			UnderlyingAssetPerContract:   "400000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "Hbfy1guHFpRdBuC4zVn8x2uMFZtnKQtCYzh2e8DxWeix",
			OptionContractMintAddress:    "BXuj4oTvph7mgqauqVhwTEd44AekNFeccS6y1sVFVd94",
			OptionWriterTokenMintAddress: "8XdUn8AVKNfn9vq8hiS6auE5hcjzPiCa3qLw6ijqxFzV",
			QuoteAssetMint:               "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			QuoteAssetPoolAddress:        "5a29mWbQCux3owm31YEnh9Bs3pFD5mxU4ds3nVvoeQt9",
			UnderlyingAssetMint:          "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			UnderlyingAssetPoolAddress:   "CegJCpmQvpsofcgzApSrHfEeeN1bAWEvyBSSKKbG1awa",
			UnderlyingAssetPerContract:   "10000",
			QuoteAssetPerContract:        "500000000",
			SerumMarketAddress:           "2q5f1H8xT3tsBzQhwZC3BKnbKMb44fTuDGamZ6xUdZz2",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
		{
			Expiry:                       1635551999,
			OptionMarketAddress:          "G9AEiXSQjeyLXdT2paa4Fb81GGbGrRqV3uENEtNjH3at",
			OptionContractMintAddress:    "CQYuDAetUSRNRqiQ2HjyeZgDpu3EjtzCgYVxgsS8mvYd",
			OptionWriterTokenMintAddress: "8bnB85xyGAgrEcSiqyHTG2czWhqV2XagbJNwS3uyv5n4",
			QuoteAssetMint:               "9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E",
			QuoteAssetPoolAddress:        "GGVZsViesBjLtrMhaGtXZdNy7bUcqcWChREuMkRepc6m",
			UnderlyingAssetMint:          "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			UnderlyingAssetPoolAddress:   "4nhtKCKxBxvPco5zZL71tVgnBg4Ai7EkamuAAkZJmem5",
			UnderlyingAssetPerContract:   "500000000",
			QuoteAssetPerContract:        "10000",
			SerumMarketAddress:           "DvohGwDZR9Z2siWBj2Xhgxd1qRScVCpywL3EoRbpon3p",
			SerumProgramId:               "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",
			PsyOptionsProgramId:          "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		},
	}
}
