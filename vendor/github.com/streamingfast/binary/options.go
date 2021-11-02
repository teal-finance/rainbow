package bin

import "encoding/binary"

type option struct {
	OptionalField bool
	SizeOfSlice   *int
	Order         binary.ByteOrder
}

func LE() binary.ByteOrder { return binary.LittleEndian }
func BE() binary.ByteOrder { return binary.BigEndian }

func newDefaultOption() *option {
	return &option{
		OptionalField: false,
		Order:         LE(),
	}
}

func (o *option) isOptional() bool {
	return o.OptionalField
}

func (o *option) hasSizeOfSlice() bool {
	return o.SizeOfSlice != nil
}

func (o *option) getSizeOfSlice() int {
	return *o.SizeOfSlice
}

func (o *option) setSizeOfSlice(size int) {
	o.SizeOfSlice = &size
}
