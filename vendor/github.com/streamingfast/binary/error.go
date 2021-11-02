package bin

import "reflect"

// An InvalidDecoderError describes an invalid argument passed to Decoder.
// (The argument to Decoder must be a non-nil pointer.)
type InvalidDecoderError struct {
	Type reflect.Type
}

func (e *InvalidDecoderError) Error() string {
	if e.Type == nil {
		return "decoder: Decode(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "decoder: Decode(non-pointer " + e.Type.String() + ")"
	}
	return "decoder: Decode(nil " + e.Type.String() + ")"
}
