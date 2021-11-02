package bin

import (
	"bytes"
	"fmt"
)

type MarshalerBinary interface {
	MarshalBinary(encoder *Encoder) error
}

type UnmarshalerBinary interface {
	UnmarshalBinary(decoder *Decoder) error
}

func MarshalBinary(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := NewEncoder(buf)
	err := encoder.Encode(v)
	return buf.Bytes(), err
}

type byteCounter struct {
	count uint64
}

func (c *byteCounter) Write(p []byte) (n int, err error) {
	c.count += uint64(len(p))
	return len(p), nil
}

// MustByteCount acts just like ByteCount but panics if it encounters any encoding errors.
func MustByteCount(v interface{}) uint64 {
	count, err := ByteCount(v)
	if err != nil {
		panic(err)
	}

	return count
}

// ByteCount computes the byte count size for the received populated structure. The reported size
// is the one for the populated structure received in arguments. Depending on how serialization of
// your fields is performed, size could vary for different structure.
func ByteCount(v interface{}) (uint64, error) {
	counter := byteCounter{}
	err := NewEncoder(&counter).Encode(v)
	if err != nil {
		return 0, fmt.Errorf("encode %T: %w", v, err)
	}

	return counter.count, nil
}
