// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package anchor

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets/anchor/generated/zeta"
)

var log = emo.NewZone("Zeta")

const (
	ZetaID    = "ZETAxsqBRek56DhiGXrn75yj2NHU3aYUnxvHXpkf3aD"
	SolanaRPC = "https://solana-mainnet.g.alchemy.com/v2/1NUlFJ7BXSudMEuTM8kns50OXHzDGDjE" //"https://solana-api.projectserum.com" // "https://api.mainnet-beta.solana.com"

)

func Query() ([]Option, map[string][]uint64, error) {
	m := make(map[string][]uint64)
	pubKey := solana.MustPublicKeyFromBase58(ZetaID)

	jsonrpcclient := rpc.NewWithRateLimit(SolanaRPC, 10)

	client := rpc.NewWithCustomRPCClient(jsonrpcclient)

	out, err := client.GetProgramAccounts(context.TODO(), pubKey)
	if err != nil {
		return nil, nil, log.Error("GetProgramAccounts", err).Err()
	}

	result := make([]Option, 0, 4*len(out))

	for _, account := range out {
		z := new(zeta.ZetaGroup)

		err = bin.NewBinDecoder(account.Account.Data.GetBinary()).Decode(&z)
		if err != nil {
			// spew.Dump(account.Pubkey)
			// too many accounts with that error to explicitely log. so we just silently skip them
			continue
			// return []Option{}, log.Error("NewBinDecoder", "account=", account, err).Err()

		}
		fillmap(m, z)
		greekInfo, err := client.GetAccountInfo(context.TODO(), z.Greeks)
		if err != nil {
			return []Option{}, nil, log.Error("GetAccountInfo", "greeks=", z.Greeks, err).Err()
		}
		gr := new(zeta.Greeks)
		err = bin.NewBinDecoder(greekInfo.Value.Data.GetBinary()).Decode(&gr)
		if err != nil {
			return []Option{}, nil, log.Error("NewBinDecoder", "greeks=", z.Greeks, err).Err()
		}

		result = append(result, extractOptions(z, gr, false)...)
		result = append(result, extractOptions(z, gr, true)...) // extra space that might be used in the future
	}
	// spew.Dump(m)
	// spew.Dump(&m)

	return result, m, nil
}

func extractOptions(z *zeta.ZetaGroup, g *zeta.Greeks, padding bool) []Option {
	options := make([]Option, 0, len(z.Products))
	if padding {
		for i := range z.ProductsPadding {
			if z.ProductsPadding[i].Strike.IsSet && z.ProductsPadding[i].Kind.String() != "Future" {
				options = append(options, Option{z, &z.ProductsPadding[i], &g.ProductGreeksPadding[i%23], z.ExpirySeriesPadding[i/23].ExpiryTs})
			}
		}
	}
	if !padding {
		for i := range z.Products {
			if z.Products[i].Strike.IsSet && z.Products[i].Kind.String() != "Future" {
				options = append(options, Option{z, &z.Products[i], &g.ProductGreeks[i%23], z.ExpirySeries[i/23].ExpiryTs})
			}
		}
	}

	return options
}

func fillmap(m map[string][]uint64, z *zeta.ZetaGroup) {
	a := Asset(z)
	e := []zeta.ExpirySeries{}
	e = append(e, z.ExpirySeries[:]...)
	e = append(e, z.ExpirySeriesPadding[:]...)

	m[a] = extractExpiries(e)
}

func extractExpiries(ze []zeta.ExpirySeries) []uint64 {
	exp := make([]uint64, 0, len(ze))
	for _, e := range ze {
		if e.ExpiryTs != 0 {
			exp = append(exp, e.ExpiryTs)
		}
	}
	return exp
}

type Option struct {
	ZG      *zeta.ZetaGroup
	Product *zeta.Product
	Greek   *zeta.ProductGreeks
	Expiry  uint64
	// in ZetaGroup the Product & ProductPaddings are by packet of 23 = 11 calls + 11 puts + 1 future.
	// To find the corresponding expiry(padding), we need its index in the array, and divide by 23
	// the quotient is the index of the expiry of interest
	// we also need to know if we are in the padding space or nah
}

func (o Option) SerumAddress() solana.PublicKey {
	return o.Product.Market
}

// watch out if they ever use something else than USDC.
var divisor = math.Pow10(USDCDecimals)

func (o Option) Strike() float64 {
	return float64(o.Product.Strike.Value) / divisor
}

// same comment as previous.
func (o Option) Quote() string {
	return "USDC"
}

func (o Option) Asset() string {
	return Asset(o.ZG)
}

func Asset(zg *zeta.ZetaGroup) string {
	switch {
	case zg.UnderlyingMint == solana.MustPublicKeyFromBase58(SOLAddress):
		return "SOL"
	case zg.UnderlyingMint == solana.MustPublicKeyFromBase58(ETHAddress):
		return "ETH"
	case zg.UnderlyingMint == solana.MustPublicKeyFromBase58(BTCAddress):
		return "BTC"
	default:
		log.Warn("Zeta Unknown token:", zg.UnderlyingMint)

		return "ZZZZ"
	}
}

func (o Option) Vol() float64 {
	return FromAnchorToDecimals(o.Greek.Volatility)
}

func (o Option) Name() string {
	return o.Asset() + "-" + o.Expiration() + "-" +
		fmt.Sprintf("%.0f", o.Strike()) + "-" + o.Product.Kind.String()
}

const ContractSize = 1000

func (o Option) Expiration() string {
	seconds := int64(o.Expiry)
	expiryTime := time.Unix(seconds, 0).UTC()
	return expiryTime.Format("2006-01-02 15:04:05")
}

func (o Option) OptionType() string {
	return strings.ToUpper(o.Product.Kind.String())
}
