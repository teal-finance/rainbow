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
	psyoptionsPubkey := solana.MustPublicKeyFromBase58(PsyOptionsProgramID)
	endpoint := "https://api.mainnet-beta.solana.com" //rpc.MainNetBeta_RPC
	jsonrpcclient := rpc.NewWithRateLimit(endpoint, 10)

	client := rpc.NewWithCustomRPCClient(jsonrpcclient)

	out, err := client.GetProgramAccounts(
		context.TODO(),
		psyoptionsPubkey,
	)
	if err != nil {
		return nil, fmt.Errorf("RPC GetProgramAccounts: %w", err)

	}
	for _, i := range out {
		o := new(Option)
		o.optionMarketAddress = i.Pubkey
		err = bin.NewBorshDecoder(i.Account.Data.GetBinary()).Decode(&o.opt)
		if err != nil {
			return nil, fmt.Errorf("Parsing Options: %w", err)

		}
		if o.IsExpired() {
			continue
		}
		if o.IsCall() {
			o.serumMarketAddress, _, err = deriveSerumMarketAddress(o.optionMarketAddress, o.opt.QuoteAssetMint, psyoptionsPubkey)
		} else {
			o.serumMarketAddress, _, err = deriveSerumMarketAddress(o.optionMarketAddress, o.opt.UnderlyingAssetMint, psyoptionsPubkey)
		}
		if err != nil {
			return nil, fmt.Errorf("Derivation Serum Address: %w", err)

		}

		result = append(result, *o)
	}
	return result, nil
}

type Option struct {
	opt                 psy.OptionMarket
	optionMarketAddress solana.PublicKey
	serumMarketAddress  solana.PublicKey
}

func (o Option) OptionMarketAddress() solana.PublicKey {
	return o.optionMarketAddress
}

// Expiration will be used in .Name() when fixed.
func (o Option) Expiration() string {
	seconds := o.opt.ExpirationUnixTimestamp
	expiryTime := time.Unix(seconds, 0).UTC()

	return expiryTime.Format("2006-01-02 15:04:05")
}

func (o Option) SerumMarketAddress() solana.PublicKey {
	return o.serumMarketAddress
}

//https://github.com/mithraiclabs/psyoptions-ts/blob/87afa7280c33f341198c60676d2302c55bbfab5f/packages/psy-american/src/serumUtils.ts#L161-L174
func deriveSerumMarketAddress(optionMarketAddress, priceCurrencyAddress, programid solana.PublicKey) (solana.PublicKey, uint8, error) {
	seed := [][]byte{
		optionMarketAddress[:],
		priceCurrencyAddress[:],
		[]byte("serumMarket"),
	}
	return solana.FindProgramAddress(seed, programid)
}

//Expired
//the Options have a field "Expired"(bool) but it is not set to false even for expired hence the function
func (o Option) IsExpired() bool {
	seconds := o.opt.ExpirationUnixTimestamp
	expiryTime := time.Unix(seconds, 0).UTC()

	//copy from opyn.go. should make a proper function

	// we keep an option even 2 days after expiry
	// mainly because not all protocol stop at expiry or right before
	// TODO re-check later
	date := time.Now()
	return !expiryTime.After(date.Add(-time.Hour * 48))
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

func (o Option) QuotePublicKey() solana.PublicKey {
	q := o.Quote()
	switch {
	case q == "USDC":
		return solana.MustPublicKeyFromBase58(USDCAddress)
	case q == "PAI":
		return solana.MustPublicKeyFromBase58(PAIAddress)
	default: // should be USD
		return solana.MustPublicKeyFromBase58(USDCAddress)
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

func (o Option) IsCall() bool {
	return o.OptionType() == "CALL"
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
