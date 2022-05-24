package main

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
)

func main() {
	array := []string{
		"GTiGjDqa5NVTzDduegkjEHTSitjA9B54y7kdSuJ1fpw2",
		"EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
		"So11111111111111111111111111111111111111112",
		"R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
	}

	/*arr := []string{
		"R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs",
		"2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk",
		"EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
		"8BfKADHfhguduy1Na7ntqkE15Aaw4NAq9xjbUea8MfM4",
	}*/
	//b(array)
	f(array)
}

func b(array []string) {
	for _, i := range array {
		for _, j := range array {
			for _, k := range array {

				m(solana.MustPublicKeyFromBase58(i),
					solana.MustPublicKeyFromBase58(j), solana.MustPublicKeyFromBase58(k))

			}

		}
	}
}

func m(id, quote, programid solana.PublicKey) {
	bumpSeed := 256
	seed := [][]byte{
		//[]byte("serumMarket"),
		id[:],
		quote[:],
		[]byte("requestQueue"), //serumMarket
		//quote[:],
	}
	for bumpSeed > -1 {
		address, err := solana.CreateProgramAddress(append(seed, []byte{byte(bumpSeed)}), programid)
		if err == nil {
			fmt.Println(address, bumpSeed, err)
		}
		bumpSeed--

	}
}

func f(array []string) {
	for _, i := range array {
		for _, j := range array {
			for _, k := range array {
				fmt.Print("input ", i, " ", j, " ", k, " ")
				fmt.Println(deriveSerumMarketAddress(solana.MustPublicKeyFromBase58(i),
					solana.MustPublicKeyFromBase58(j), solana.MustPublicKeyFromBase58(k)))

			}

		}
	}
}

func deriveSerumMarketAddress(id, quote, programid solana.PublicKey) (solana.PublicKey, uint8, error) {
	seed := [][]byte{
		id[:],
		quote[:],
		[]byte("serumMarket"),
	}
	return solana.FindProgramAddress(seed, programid)
}
