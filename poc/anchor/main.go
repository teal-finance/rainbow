// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/spewerspew/spew"
	sol "github.com/streamingfast/solana-go"

	zeta "github.com/teal-finance/rainbow/pkg/provider/zetamarkets/anchor/generated/zeta"
)

const (
	PsyOptionsProgramID = "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs"
	ZetaID              = "ZETAxsqBRek56DhiGXrn75yj2NHU3aYUnxvHXpkf3aD"
	test                = "BPFLoaderUpgradeab1e11111111111111111111111"
	USDCSolana          = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
)

func main() {
	z := solana.MustPublicKeyFromBase58(ZetaID)
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)

	out, err := client.GetProgramAccounts(
		context.TODO(),
		z,
	)
	if err != nil {
		panic(err)
	}
	// spew.Dump(out[0])

	/*pub := solana.MustPublicKeyFromBase58(PsyOptionsProgramID)

	out, err = client.GetProgramAccounts(
		context.TODO(),
		pub,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)*/
	// spew.Dump(len(out[0:3]))
	for _, i := range out {
		// spew.Dump(i.Pubkey)
		/*out, err := client.GetAccountInfo(
			context.TODO(),
			i.Pubkey,
		)
		spew.Dump(out)
		if err != nil {
			panic(err)
		}*/
		zo := new(zeta.ZetaGroup)
		err = bin.NewBinDecoder(i.Account.Data.GetBinary()).Decode(&zo)

		// opt := new(OptionMarket)
		// err = bin.NewBorshDecoder(i.Account.Data.GetBinary()).Decode(&opt)
		if err != nil {
			continue
		}

		spew.Dump(zo) // opt)
		/*
			a, b, c := deriveSerumMarketAddress(i.Pubkey, solana.PublicKey(opt.QuoteAssetMint), pub)
			spew.Dump(a, b, c)
			a, b, c = deriveSerumMarketAddress(i.Pubkey, solana.PublicKey(opt.UnderlyingAssetMint), pub)
			spew.Dump(a, b, c)*/
	}
	// f()
	// solana.MustPublicKeyFromBase58(test))
	/*res, err := client.GetAccountInfo(
		context.TODO(),
		pub,
	)
	if err != nil {
		panic(err)
	}
	//spew.Dump(len(res))
	spew.Dump(res)*/
	// psy_american.SetProgramID(pub)
	// fmt.Println(psy_american.ProgramName)
}

func deriveSerumMarketAddress(id, quote, programid solana.PublicKey) (solana.PublicKey, uint8, error) {
	seed := [][]byte{
		id[:],
		quote[:],
		[]byte("serumMarket"),
	}
	return solana.FindProgramAddress(seed, programid)
}

func dd(pub, quote, programid solana.PublicKey) (solana.PublicKey, uint8, error) {
	seed := [][]byte{
		[]byte(""),
		{1},
	}
	return solana.FindProgramAddress(seed, solana.MustPublicKeyFromBase58(test))
}

func f() {
	program_id := solana.MustPublicKeyFromBase58("BPFLoader1111111111111111111111111111111111")
	public_key := solana.MustPublicKeyFromBase58("SeedPubey1111111111111111111111111111111111")

	got, _ := solana.CreateProgramAddress([][]byte{
		{},
		{1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte(""),
		{1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte(""),
		// {1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte("☉"),
		{1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte("☉"),
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		public_key[:],
		{1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte("Talking"),
		[]byte("Squirrels"),
	},
		program_id,
	)
	fmt.Println(got)
}

/*var address solana.PublicKey
	var err error
	bumpSeed := uint8(math.MaxUint8)
	for bumpSeed > 0 {
		fmt.Print(bumpSeed)

		address, err = solana.CreateProgramAddress(append(seed, []byte{byte(bumpSeed)}), programID)
		if err == nil {
			fmt.Println(address)
		}
		bumpSeed--

	}
}*/

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
