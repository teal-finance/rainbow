Psyoptions uses Anchor, a framework for Solana's Sealevel runtime providing several convenient developer tools for writing smart contracts.
https://github.com/project-serum/anchor

We want to automcatically query all the available options from Psyoptions. To do so, we will use anchor-go which generates Go clients for Solana programs (smart contracts) written using the anchor framework.
```
go get github.com/gagliardetto/anchor-go
go run github.com/gagliardetto/anchor-go --src=idl.json
```
We actually only need the structure:
````go
type OptionMarket struct {
	OptionMint                  ag_solanago.PublicKey
	WriterTokenMint             ag_solanago.PublicKey
	UnderlyingAssetMint         ag_solanago.PublicKey
	QuoteAssetMint              ag_solanago.PublicKey
	UnderlyingAmountPerContract uint64
	QuoteAmountPerContract      uint64
	ExpirationUnixTimestamp     int64
	UnderlyingAssetPool         ag_solanago.PublicKey
	QuoteAssetPool              ag_solanago.PublicKey
	MintFeeAccount              ag_solanago.PublicKey
	ExerciseFeeAccount          ag_solanago.PublicKey
	Expired                     bool
	BumpSeed                    uint8
}
`````
form the file `accounts.go`

idl file is copied from :
````
curl https://raw.githubusercontent.com/mithraiclabs/psyoptions-ts/master/packages/psy-american/src/idl.json -o idl.json
````

But it can also be generated using the anchor cli doing something like:
```
anchor idl parse -f lib.rs -o idl.json
```

Note:
There are 2 Solana go libraries and they are pretty similar:

github.com/gagliardetto/solana-go

github.com/streamingfast/solana-go

We were using the second one but since `anchor-go` is from the same authors as the first one, we will switch all the code to use that lib.


