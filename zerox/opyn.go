package zerox

const OpynQuoteCurrency = "USDC"

type Opyn struct {
	ID                  string //0x id
	Name                string // ASSET-DATE-Strike-OptionsType
	Type                string // CALL or PUT
	Asset               string // ETH, BTC, SOL
	ExpirationTimestamp int64
	Strike              float64
}

func Opynmarket(coin string) []Opyn {
	if coin != "ETH" {
		return []Opyn{}
	}
	return []Opyn{
		{
			ID:                  "0x16aab4738843fb2d9eafc8fd261488797bf0df29",
			Name:                "WETHUSDC 29-October-2021 2800Put USDC",
			Type:                "PUT",
			Asset:               "ETH",
			ExpirationTimestamp: 1635494400,
			Strike:              2800,
		},
		{
			ID:                  "0x1a355ed2821d63f7ece49747fec44a0d17340e5f",
			Name:                "WETHUSDC 29-October-2021 3000Call",
			Type:                "CALL",
			Asset:               "ETH",
			ExpirationTimestamp: 1635494400,
			Strike:              3000,
		},
		{
			ID:                  "0x796091d1840a6bd9e6a6b65bf8ef3985d254c3df",
			Name:                "WETHUSDC 29-October-2021 3800Call",
			Type:                "CALL",
			Asset:               "ETH",
			ExpirationTimestamp: 1635494400,
			Strike:              3800,
		},
	}

}
