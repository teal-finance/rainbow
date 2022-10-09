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
	ZetaID = "ZETAxsqBRek56DhiGXrn75yj2NHU3aYUnxvHXpkf3aD"
	//endpoint = "https://api.mainnet-beta.solana.com" // rpc.MainNetBeta_RPC
	SolanaRPC = "https://solana-mainnet.g.alchemy.com/v2/1NUlFJ7BXSudMEuTM8kns50OXHzDGDjE" //"https://solana-api.projectserum.com" // "https://api.mainnet-beta.solana.com"

)

func Query() ([]Option, error) {
	pubKey := solana.MustPublicKeyFromBase58(ZetaID)

	jsonrpcclient := rpc.NewWithRateLimit(SolanaRPC, 10)

	client := rpc.NewWithCustomRPCClient(jsonrpcclient)

	out, err := client.GetProgramAccounts(context.TODO(), pubKey)
	if err != nil {
		return nil, err
	}

	result := make([]Option, 0, 4*len(out))

	for _, account := range out {
		z := new(zeta.ZetaGroup)

		err = bin.NewBinDecoder(account.Account.Data.GetBinary()).Decode(&z)
		if err != nil {
			// TODO make proper error since this is fixed

			continue
		}

		result = append(result, extractOptions(z, z.Products[:], false)...)
		result = append(result, extractOptions(z, z.ProductsPadding[:], true)...) // extra space that might be used in the future
	}

	return result, nil
}

func extractOptions(z *zeta.ZetaGroup, products []zeta.Product, padding bool) []Option {
	options := make([]Option, 0, len(products))

	counter := 0

	for _, p := range products {
		if p.Strike.IsSet && p.Kind.String() != "Future" {
			if padding {
				options = append(options, Option{z, p, z.ExpirySeriesPadding[counter/23].ExpiryTs})
			} else {
				options = append(options, Option{z, p, z.ExpirySeries[counter/23].ExpiryTs})
			}
		}

		counter++
	}

	return options
}

type Option struct {
	ZG      *zeta.ZetaGroup
	Product zeta.Product
	expiry  uint64
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
	switch {
	case o.ZG.UnderlyingMint == solana.MustPublicKeyFromBase58(SOLAddress):
		return "SOL"
	case o.ZG.UnderlyingMint == solana.MustPublicKeyFromBase58(ETHAddress):
		return "ETH"
	case o.ZG.UnderlyingMint == solana.MustPublicKeyFromBase58(BTCAddress):
		return "BTC"
	default:
		log.Warn("Zeta Unknown token:", o.ZG.UnderlyingMint)

		return "ZZZZ"
	}
}

func (o Option) Name() string {
	return o.Asset() + "-" + o.Expiration() + "-" +
		fmt.Sprintf("%.0f", o.Strike()) + "-" + o.Product.Kind.String()
}

const ContractSize = 1000

func (o Option) Expiration() string {
	seconds := int64(o.expiry)
	expiryTime := time.Unix(seconds, 0).UTC()
	return expiryTime.Format("2006-01-02 15:04:05")
}

func (o Option) OptionType() string {
	return strings.ToUpper(o.Product.Kind.String())
}
