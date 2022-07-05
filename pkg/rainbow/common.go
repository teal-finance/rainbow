package rainbow

import (
	"log"
	"math/big"
	"strconv"
	"time"
)

// IntToEthereumFormat convert an int to the Ethereum format.
// Ethereum use 18 decimals in their VM.
// To quote 1 option, you thus need to send 1000000000000000000
func IntToEthereumFormat(i int) *big.Int {
	a := big.NewInt(10)
	a.Exp(big.NewInt(10), big.NewInt(18), nil)
	a.Mul(a, big.NewInt(int64(i)))
	return a
}

// ToFloat
// because Ethereum use 10^18 to represent 1
// what we do is that we cut 10^13 to keep 5 decimals
// (because IV is a percentage an we want to be accurate)
// then convert the remainder to float64 and divide by 1000.
func ToFloat(n *big.Int) float64 {
	q := big.NewInt(0)
	q.Quo(n, big.NewInt(10000000000000)) // divided by 10^13
	return float64(q.Int64()) / 100000.0
}

// TimeStringConvert convert a Unix string to a UTC one
func TimeStringConvert(s string) (expiry string, err error) {
	seconds, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("ERR TimeStringConvert: %v from %+v", err, s)
	} else {
		expiry = time.Unix(seconds, 0).UTC().Format("2006-01-02 15:04:05")
	}
	return expiry, err
}
