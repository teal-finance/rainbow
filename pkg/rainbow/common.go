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

// ToFloat converts a big.Int{decimals} to float64.
// Because a token use 10^decimals to represent 1
// what we do is that we cut 10^(decimals-5) to keep 5 decimals
// (because IV is a percentage an we want to be accurate)
// then convert the remainder to float64 and divide by 1000.
func ToFloat(n *big.Int, decimals int64) float64 {
	if decimals == 0 {
		return float64(n.Int64())
	}

	// TODO remove this fix and add better check upfront
	// divide by 10^(decimals-5)
	divisor := big.NewInt(int64(math.Pow(10, float64(decimals-5))))
	if divisor.Int64() == 0 {
		return 0
	}

	q := big.NewInt(0)
	q.Quo(n, divisor)
	return float64(q.Int64()) / 100000.0
}

// ToFloats converts multiple big.Int{decimals} to float64 values.
// ToFloats is the slice version of ToFloat to apply the conversion to each *big.Int of the slice.
func ToFloats(bigs []*big.Int, decimals int64) []float64 {
	floats := make([]float64, len(bigs))
	for i, b := range bigs {
		floats[i] = ToFloat(b, decimals)
	}
	return floats
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

// IsBigIntNull check if the big Int is ==0
// https://stackoverflow.com/questions/64257065/is-there-another-way-of-testing-if-a-big-int-is-0
// TODO make a function maybe
func IsBigIntNull(b *big.Int) bool {

	return len(b.Bits()) == 0
}
