package rainbow

const listMarketsURL = "wss://api.psyoptions.io/v1/graphql"

func GetListMarkets(coin string) []string {
	if coin == "ETH" {
		return []string{
			"8A493gU55NfS4fCjDoLAiN57zPzWf6QQw31QQf1fd6iX",
			"5pHcU2Gz8eCMwynLvz1AHSFoFbKkeTc7ufeqeG4spb99",
			"8qgAVVE6eeJ9u32LvJupgW6NyNWPPFFK4xnXfGtDNeP4",
			"FfkVR6ha6N9fcefGxfDFP97xL54Pc5ipiwnt1uuTpuyX",
			"TLAzx53rb6pEDT3oWitTzg6YNe7BTERLDzujMdU8RQx",
			"9A1m1JXgnMrdq3JkqEYx9rK3CcYm9EHDLeA6sn1hH3PP",
			"HYmPvo8szh62QVaAfUAXR1eppvCfouUPpH68yE87UYmy",
			"8fFcWuVaZSKoge4DCpMcNrR5nNXF2pbXCfBUxkMomgr5",
		}
	}

	return []string{
		"8fhiAYm41RwtiX8WusCSpY617GWPt2LwUnCQcEeer78o",
		"6at26sVk8vTYtLh4YDKvje4PDdgFJsNHHyoGw87WNszP",
		"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
		"7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
		"9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
		"ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28",
		"2q5f1H8xT3tsBzQhwZC3BKnbKMb44fTuDGamZ6xUdZz2",
		"DvohGwDZR9Z2siWBj2Xhgxd1qRScVCpywL3EoRbpon3p",
	}
}
