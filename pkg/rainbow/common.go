package rainbow

import (
	"math"
	"math/big"
	"strconv"
	"time"
)

// DefaultEthereumDecimals is the default decimals that most token uses, starting with ETH.
const DefaultEthereumDecimals int64 = 18

// IntToEthereumUint256 convert an int to the Ethereum format.
// Ethereum use 18 decimals in their VM.
// To quote 1 option, you thus need to send 1000000000000000000.
func IntToEthereumUint256(i int, decimals int64) *big.Int {
	a := big.NewInt(10)
	a.Exp(big.NewInt(10), big.NewInt(decimals), nil)
	a.Mul(a, big.NewInt(int64(i)))
	return a
}

// ToFloat
// because a token use 10^decimals to represent 1
// what we do is that we cut 10^(decimals-5) to keep 5 decimals
// (because IV is a percentage an we want to be accurate)
// then convert the remainder to float64 and divide by 1000.
func ToFloat(n *big.Int, decimals int64) float64 {
	q := big.NewInt(0)
	q.Quo(n, big.NewInt(int64(math.Pow(10, float64(decimals-5))))) // divided by 10^(decimals-5)
	return float64(q.Int64()) / 100000.0
}

// TimeStringConvert convert a Unix string to a UTC one.
func TimeStringConvert(s string) (expiry string, err error) {
	seconds, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Errorf("TimeStringConvert: %v from %+v", err, s)
	} else {
		expiry = time.Unix(seconds, 0).UTC().Format("2006-01-02 15:04:05")
	}
	return expiry, err
}
