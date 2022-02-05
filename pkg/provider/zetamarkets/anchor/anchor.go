package anchor

import (
	"context"
	"fmt"
	"math"
	"time"

	bin "github.com/gagliardetto/binary"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets/anchor/generated/zeta"
)

const ZetaID = "ZETAxsqBRek56DhiGXrn75yj2NHU3aYUnxvHXpkf3aD"
const USDCAddress = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
const USDCDecimals = 6
const SOLAddress = "So11111111111111111111111111111111111111112"

func Query() ([]Option, error) {
	var result []Option
	pubKey := solana.MustPublicKeyFromBase58(ZetaID)
	endpoint := rpc.MainNetBeta_RPC
	jsonrpcclient := rpc.NewWithRateLimit(endpoint, 10)

	client := rpc.NewWithCustomRPCClient(jsonrpcclient)

	out, err := client.GetProgramAccounts(
		context.TODO(),
		pubKey,
	)
	if err != nil {
		return nil, err
	}
	for _, i := range out {
		z := new(zeta.ZetaGroup)
		err = bin.NewBinDecoder(i.Account.Data.GetBinary()).Decode(&z)
		if err != nil {
			continue
		}

		result = append(result, extractOptions(z, z.Products[:])...)
		result = append(result, extractOptions(z, z.ProductsPadding[:])...) //extra space that might be used in the future
	}

	return result, nil
}

func extractOptions(z *zeta.ZetaGroup, products []zeta.Product) []Option {
	var options []Option
	for _, p := range products {
		if p.Strike.IsSet && p.Kind.String() != "Future" {
			options = append(options, Option{z, p})

		}
	}
	return options
}

type Option struct {
	ZG      *zeta.ZetaGroup
	Product zeta.Product
}

func (o Option) SerumAddress() solana.PublicKey {
	return o.Product.Market
}

//watch out if they ever use something else than USDC
func (o Option) Strike() float64 {

	return float64(o.Product.Strike.Value / uint64(math.Pow10(USDCDecimals)))
}

//same comment as previous
func (o Option) Quote() string {
	return "USDC"
}

func (o Option) Asset() string {
	switch {
	case o.ZG.UnderlyingMint == solana.MustPublicKeyFromBase58(SOLAddress):
		return "SOL"
	default:
		return "?"
	}
}
func (o Option) Name() string {
	return o.Asset() + "-" + o.Expiration() + "-" +
		fmt.Sprintf("%.0f", o.Strike()) + "-" + o.Product.Kind.String()
}
func (o Option) ContractSize() float64 {
	switch {
	//SOL is 1
	default:
		return 1000
	}
}
func (o Option) Expiration() string {
	seconds := int64(o.ZG.ExpirySeries[0].ExpiryTs)
	expiryTime := time.Unix(seconds, 0).UTC()

	return expiryTime.Format("2006-01-02 15:04:05")
}
