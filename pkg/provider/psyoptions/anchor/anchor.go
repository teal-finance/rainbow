package anchor

import (
	"context"
	"fmt"
	"time"

	bin "github.com/gagliardetto/binary"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	psy "github.com/teal-finance/rainbow/pkg/provider/psyoptions/anchor/generated/psy_american"
)

const PsyOptionsProgramID = "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs"

func Query() ([]Option, error) {
	var result []Option
	pub := solana.MustPublicKeyFromBase58(PsyOptionsProgramID)
	client := rpc.New(rpc.MainNetBeta_RPC)

	out, err := client.GetProgramAccounts(
		context.TODO(),
		pub,
	)
	if err != nil {
		return nil, fmt.Errorf("RPC GetProgramAccounts: %w", err)

	}
	for _, i := range out {
		o := new(Option)
		err = bin.NewBinDecoder(i.Account.Data.GetBinary()).Decode(&o.opt)
		if err != nil {
			return nil, fmt.Errorf("Parsing Options: %w", err)

		}
		result = append(result, *o)
	}
	return result, nil
}

type Option struct {
	opt                psy.OptionMarket
	serumMarketAddress string
}

// Expiration will be used in .Name() when fixed.
func (o Option) Expiration() string {
	seconds := o.opt.ExpirationUnixTimestamp
	expiryTime := time.Unix(seconds, 0).UTC()

	return expiryTime.Format("2006-01-02 15:04:05")
}

func (o Option) SerumMarketAddress() string {
	return o.serumMarketAddress
}

func (o Option) Asset() string {
	switch {
	case o.opt.QuoteAssetMint == solana.MustPublicKeyFromBase58(ETHAddress) || o.opt.UnderlyingAssetMint == solana.MustPublicKeyFromBase58(ETHAddress):
		return "ETH"
	case o.opt.QuoteAssetMint == solana.MustPublicKeyFromBase58(BTCAddress) || o.opt.UnderlyingAssetMint == solana.MustPublicKeyFromBase58(BTCAddress):
		return "BTC"
	case o.opt.QuoteAssetMint == solana.MustPublicKeyFromBase58(SOLAddress) || o.opt.UnderlyingAssetMint == solana.MustPublicKeyFromBase58(SOLAddress):
		return "SOL"
	case o.opt.QuoteAssetMint == solana.MustPublicKeyFromBase58(MSOLAddress) || o.opt.UnderlyingAssetMint == solana.MustPublicKeyFromBase58(MSOLAddress):
		return "mSOL"
	default:
		return "ZZZZ"
	}
}
func (o Option) Quote() string {
	switch {
	case o.opt.QuoteAssetMint == solana.MustPublicKeyFromBase58(USDCAddress) || o.opt.UnderlyingAssetMint == solana.MustPublicKeyFromBase58(USDCAddress):
		return "USDC"
	case o.opt.QuoteAssetMint == solana.MustPublicKeyFromBase58(PAIAddress) || o.opt.UnderlyingAssetMint == solana.MustPublicKeyFromBase58(PAIAddress):
		return "PAI"
	default:
		return "USD"
	}
}
func (o Option) ContractSize() float64 {
	switch {
	case o.Asset() == "ETH":
		return 0.1
	case o.Asset() == "BTC":
		return 0.01
	//SOL & mSOL is 1
	default:
		return 1
	}
}

func (o Option) UnderlyingPerContract() float64 {
	return float64(o.opt.UnderlyingAmountPerContract)
}

func (o Option) QuotePerContract() float64 {
	return float64(o.opt.QuoteAmountPerContract)
}

func (o Option) OptionType() string {
	if o.UnderlyingPerContract() < o.QuotePerContract() {
		return "CALL"
	}

	return "PUT"
}

func (o Option) Strike() float64 {
	if o.OptionType() == "PUT" {
		return o.UnderlyingPerContract() / o.QuotePerContract()
	}

	return o.QuotePerContract() / o.UnderlyingPerContract()
}

func (o Option) Name() string {
	return o.Asset() + "-" + o.Expiration() + "-" +
		fmt.Sprintf("%.0f", o.Strike()) + "-" + o.OptionType()
}
