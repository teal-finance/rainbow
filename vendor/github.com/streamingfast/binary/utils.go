package bin

import (
	"math/big"
	"reflect"
)

func isZero(rv reflect.Value) (b bool) {
	return rv.Kind() == 0
}

func twosComplement(v []byte) []byte {
	buf := make([]byte, len(v))
	for i, b := range v {
		buf[i] = b ^ byte(0xff)
	}
	one := big.NewInt(1)
	value := (&big.Int{}).SetBytes(buf)
	return value.Add(value, one).Bytes()
}
