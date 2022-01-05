package main

import (
	"context"

	bin "github.com/gagliardetto/binary"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	sol "github.com/streamingfast/solana-go"
)

const PsyOptionsProgramID = "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs"

func main() {
	pub := solana.MustPublicKeyFromBase58(PsyOptionsProgramID)
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)

	out, err := client.GetProgramAccounts(
		context.TODO(),
		pub,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(len(out[0:10]))
	for _, i := range out[0:10] {
		spew.Dump(i)
		opt := new(OptionMarket)
		err = bin.NewBinDecoder(i.Account.Data.GetBinary()).Decode(&opt)
		if err != nil {
			panic(err)
		}
		spew.Dump(opt)
	}

	/*res, err := client.GetAccountInfo(
		context.TODO(),
		pub,
	)
	if err != nil {
		panic(err)
	}
	//spew.Dump(len(res))
	spew.Dump(res)*/
	//psy_american.SetProgramID(pub)
	//fmt.Println(psy_american.ProgramName)
}

type anchorOptions struct {
	Options OptionMarket
}

type OptionMarket struct {
	OptionMint                  sol.PublicKey
	WriterTokenMint             sol.PublicKey
	UnderlyingAssetMint         sol.PublicKey
	QuoteAssetMint              sol.PublicKey
	UnderlyingAmountPerContract uint64
	QuoteAmountPerContract      uint64
	ExpirationUnixTimestamp     int64
	UnderlyingAssetPool         sol.PublicKey
	QuoteAssetPool              sol.PublicKey
	MintFeeAccount              sol.PublicKey
	ExerciseFeeAccount          sol.PublicKey
	Expired                     bool
	BumpSeed                    uint8
}
