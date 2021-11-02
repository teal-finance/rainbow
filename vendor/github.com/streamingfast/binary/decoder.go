package bin

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
	"unicode/utf8"

	"go.uber.org/zap"
)

var TypeSize = struct {
	Bool int
	Byte int

	Int8  int
	Int16 int

	Uint8   int
	Uint16  int
	Uint32  int
	Uint64  int
	Uint128 int

	Float32 int
	Float64 int

	PublicKey int
	Signature int

	Tstamp         int
	BlockTimestamp int

	CurrencyName int
}{
	Byte: 1,
	Bool: 1,

	Int8:  1,
	Int16: 2,

	Uint8:   1,
	Uint16:  2,
	Uint32:  4,
	Uint64:  8,
	Uint128: 16,

	Float32: 4,
	Float64: 8,
}

// Decoder implements the EOS unpacking, similar to FC_BUFFER
type Decoder struct {
	data []byte
	pos  int

	currentFieldOpt *option
}

func NewDecoder(data []byte) *Decoder {
	return &Decoder{
		data: data,
	}
}

func (d *Decoder) Decode(v interface{}) (err error) {
	return d.decodeWithOption(v, nil)
}

func (d *Decoder) decodeWithOption(v interface{}, option *option) (err error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return &InvalidDecoderError{reflect.TypeOf(v)}
	}

	// We decode rv not rv.Elem because the Unmarshaler interface
	// test must be applied at the top level of the value.
	err = d.decode(rv, option)
	if err != nil {
		return err
	}
	return nil
}

func (d *Decoder) decode(rv reflect.Value, opt *option) (err error) {
	if opt == nil {
		opt = newDefaultOption()
	}
	d.currentFieldOpt = opt

	unmarshaler, rv := indirect(rv, opt.isOptional())

	if traceEnabled {
		zlog.Debug("decode: type",
			zap.Stringer("value_kind", rv.Kind()),
			zap.Bool("has_unmarshaler", (unmarshaler != nil)),
			zap.Reflect("options", opt),
		)
	}

	if opt.isOptional() {
		isPresent, e := d.ReadByte()
		if e != nil {
			err = fmt.Errorf("decode: %t isPresent, %s", rv.Type(), e)
			return
		}

		if isPresent == 0 {
			if traceEnabled {
				zlog.Debug("decode: skipping optional value", zap.Stringer("type", rv.Kind()))
			}

			rv.Set(reflect.Zero(rv.Type()))
			return
		}

		// we have ptr here we should not go get the element
		unmarshaler, rv = indirect(rv, false)
	}

	if unmarshaler != nil {
		if traceEnabled {
			zlog.Debug("decode: using UnmarshalBinary method to decode type")
		}
		return unmarshaler.UnmarshalBinary(d)
	}
	rt := rv.Type()

	switch rv.Kind() {
	case reflect.String:
		s, e := d.ReadString()
		if e != nil {
			err = e
			return
		}
		rv.SetString(s)
		return
	case reflect.Uint8:
		var n byte
		n, err = d.ReadByte()
		rv.SetUint(uint64(n))
		return
	case reflect.Int8:
		var n int8
		n, err = d.ReadInt8()
		rv.SetInt(int64(n))
		return
	case reflect.Int16:
		var n int16
		n, err = d.ReadInt16(opt.Order)
		rv.SetInt(int64(n))
		return
	case reflect.Int32:
		var n int32
		n, err = d.ReadInt32(opt.Order)
		rv.SetInt(int64(n))
		return
	case reflect.Int64:
		var n int64
		n, err = d.ReadInt64(opt.Order)
		rv.SetInt(int64(n))
		return
	case reflect.Uint16:
		var n uint16
		n, err = d.ReadUint16(opt.Order)
		rv.SetUint(uint64(n))
		return
	case reflect.Uint32:
		var n uint32
		n, err = d.ReadUint32(opt.Order)
		rv.SetUint(uint64(n))
		return
	case reflect.Uint64:
		var n uint64
		n, err = d.ReadUint64(opt.Order)
		rv.SetUint(n)
		return
	case reflect.Float32:
		var n float32
		n, err = d.ReadFloat32(opt.Order)
		rv.SetFloat(float64(n))
		return
	case reflect.Float64:
		var n float64
		n, err = d.ReadFloat64(opt.Order)
		rv.SetFloat(n)
		return
	case reflect.Bool:
		var r bool
		r, err = d.ReadBool()
		rv.SetBool(r)
		return
	}
	switch rt.Kind() {
	case reflect.Array:
		length := rt.Len()
		if traceEnabled {
			zlog.Debug("decoding: reading array", zap.Int("length", length))
		}
		for i := 0; i < length; i++ {
			if err = d.decode(rv.Index(i), opt); err != nil {
				return
			}
		}
		return
	case reflect.Slice:
		var l int
		if opt.hasSizeOfSlice() {
			l = opt.getSizeOfSlice()
		} else {
			length, err := d.ReadUvarint64()
			if err != nil {
				return err
			}
			l = int(length)
		}

		if traceEnabled {
			zlog.Debug("reading slice", zap.Int("len", l), typeField("type", rv))
		}

		rv.Set(reflect.MakeSlice(rt, l, l))
		for i := 0; i < l; i++ {
			if err = d.decode(rv.Index(i), opt); err != nil {
				return
			}
		}

	case reflect.Struct:
		if err = d.decodeStruct(rt, rv); err != nil {
			return
		}

	default:
		return fmt.Errorf("decode: unsupported type %q", rt)
	}

	return
}

func (d *Decoder) decodeStruct(rt reflect.Type, rv reflect.Value) (err error) {
	l := rv.NumField()

	if traceEnabled {
		zlog.Debug("decode: struct", zap.Int("fields", l), zap.Stringer("type", rv.Kind()))
	}

	sizeOfMap := map[string]int{}
	seenBinaryExtensionField := false
	for i := 0; i < l; i++ {
		structField := rt.Field(i)
		fieldTag := parseFieldTag(structField.Tag)

		if fieldTag.Skip {
			if traceEnabled {
				zlog.Debug("decode: skipping struct field with skip flag",
					zap.String("struct_field_name", structField.Name),
				)
			}
			continue
		}

		if !fieldTag.BinaryExtension && seenBinaryExtensionField {
			panic(fmt.Sprintf("the `bin:\"binary_extension\"` tags must be packed together at the end of struct fields, problematic field %q", structField.Name))
		}

		if fieldTag.BinaryExtension {
			seenBinaryExtensionField = true
			// FIXME: This works only if what is in `d.data` is the actual full data buffer that
			//        needs to be decoded. If there is for example two structs in the buffer, this
			//        will not work as we would continue into the next struct.
			//
			//        But at the same time, does it make sense otherwise? What would be the inference
			//        rule in the case of extra bytes available? Continue decoding and revert if it's
			//        not working? But how to detect valid errors?
			if len(d.data[d.pos:]) <= 0 {
				continue
			}
		}
		v := rv.Field(i)
		if !v.CanSet() {
			// This means that the field cannot be set, to fix this
			// we need to create a pointer to said field
			if !v.CanAddr() {
				// we cannot create a point to field skipping
				if traceEnabled {
					zlog.Debug("skipping struct field that cannot be addressed",
						zap.String("struct_field_name", structField.Name),
						zap.Stringer("struct_value_type", v.Kind()),
					)
				}
				return fmt.Errorf("unable to decode a none setup struc field %q with type %q", structField.Name, v.Kind())
			}
			v = v.Addr()
		}

		if !v.CanSet() {
			if traceEnabled {
				zlog.Debug("skipping struct field that cannot be addressed",
					zap.String("struct_field_name", structField.Name),
					zap.Stringer("struct_value_type", v.Kind()),
				)
			}
			continue
		}

		option := &option{
			OptionalField: fieldTag.Optional,
			Order:         fieldTag.Order,
		}

		if s, ok := sizeOfMap[structField.Name]; ok {
			option.setSizeOfSlice(s)
		}

		if traceEnabled {
			zlog.Debug("decode: struct field",
				zap.Stringer("struct_field_value_type", v.Kind()),
				zap.String("struct_field_name", structField.Name),
				zap.Reflect("struct_field_tags", fieldTag),
				zap.Reflect("struct_field_option", option),
			)
		}

		if err = d.decode(v, option); err != nil {
			return
		}

		if fieldTag.SizeOf != "" {
			size := sizeof(structField.Type, v)
			if traceEnabled {
				zlog.Debug("setting size of field",
					zap.String("field_name", fieldTag.SizeOf),
					zap.Int("size", size),
				)
			}
			sizeOfMap[fieldTag.SizeOf] = size
		}
	}
	return
}

func sizeof(t reflect.Type, v reflect.Value) int {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n := int(v.Uint())
		// all the builtin array length types are native int
		// so this guards against weird truncation
		if n < 0 {
			return 0
		}
		return n
	default:
		panic(fmt.Sprintf("sizeof field "))
	}
}

var ErrVarIntBufferSize = errors.New("varint: invalid buffer size")

func (d *Decoder) ReadUvarint64() (uint64, error) {
	l, read := binary.Uvarint(d.data[d.pos:])
	if read <= 0 {
		return l, ErrVarIntBufferSize
	}
	if traceEnabled {
		zlog.Debug("decode: read uvarint64", zap.Uint64("val", l))
	}
	d.pos += read
	return l, nil
}

func (d *Decoder) ReadVarint64() (out int64, err error) {
	l, read := binary.Varint(d.data[d.pos:])
	if read <= 0 {
		return l, ErrVarIntBufferSize
	}
	if traceEnabled {
		zlog.Debug("decode: read varint", zap.Int64("val", l))
	}
	d.pos += read
	return l, nil
}

func (d *Decoder) ReadVarint32() (out int32, err error) {
	n, err := d.ReadVarint64()
	if err != nil {
		return out, err
	}
	out = int32(n)
	if traceEnabled {
		zlog.Debug("decode: read varint32", zap.Int32("val", out))
	}
	return
}

func (d *Decoder) ReadUvarint32() (out uint32, err error) {

	n, err := d.ReadUvarint64()
	if err != nil {
		return out, err
	}
	out = uint32(n)
	if traceEnabled {
		zlog.Debug("decode: read uvarint32", zap.Uint32("val", out))
	}
	return
}
func (d *Decoder) ReadVarint16() (out int16, err error) {
	n, err := d.ReadVarint64()
	if err != nil {
		return out, err
	}
	out = int16(n)
	if traceEnabled {
		zlog.Debug("decode: read varint16", zap.Int16("val", out))
	}
	return
}

func (d *Decoder) ReadUvarint16() (out uint16, err error) {

	n, err := d.ReadUvarint64()
	if err != nil {
		return out, err
	}
	out = uint16(n)
	if traceEnabled {
		zlog.Debug("decode: read uvarint16", zap.Uint16("val", out))
	}
	return
}

func (d *Decoder) ReadByteArray() (out []byte, err error) {

	l, err := d.ReadUvarint64()
	if err != nil {
		return nil, err
	}

	if len(d.data) < d.pos+int(l) {
		return nil, fmt.Errorf("byte array: varlen=%d, missing %d bytes", l, d.pos+int(l)-len(d.data))
	}

	out = d.data[d.pos : d.pos+int(l)]
	d.pos += int(l)
	if traceEnabled {
		zlog.Debug("decode: read byte array", zap.Stringer("hex", HexBytes(out)))
	}
	return
}

func (d *Decoder) ReadByte() (out byte, err error) {
	if d.Remaining() < TypeSize.Byte {
		err = fmt.Errorf("required [1] byte, remaining [%d]", d.Remaining())
		return
	}

	out = d.data[d.pos]
	d.pos++
	if traceEnabled {
		zlog.Debug("decode: read byte", zap.Uint8("byte", out), zap.String("hex", hex.EncodeToString([]byte{out})))
	}
	return
}

func (d *Decoder) ReadBool() (out bool, err error) {
	if d.Remaining() < TypeSize.Bool {
		err = fmt.Errorf("bool required [%d] byte, remaining [%d]", TypeSize.Bool, d.Remaining())
		return
	}

	b, err := d.ReadByte()

	if err != nil {
		err = fmt.Errorf("readBool, %s", err)
	}
	out = b != 0
	if traceEnabled {
		zlog.Debug("decode: read bool", zap.Bool("val", out))
	}
	return

}

func (d *Decoder) ReadUint8() (out uint8, err error) {
	out, err = d.ReadByte()
	return
}

func (d *Decoder) ReadInt8() (out int8, err error) {
	b, err := d.ReadByte()
	out = int8(b)
	if traceEnabled {
		zlog.Debug("decode: read int8", zap.Int8("val", out))
	}
	return
}

func (d *Decoder) ReadUint16(order binary.ByteOrder) (out uint16, err error) {
	if d.Remaining() < TypeSize.Uint16 {
		err = fmt.Errorf("uint16 required [%d] bytes, remaining [%d]", TypeSize.Uint16, d.Remaining())
		return
	}

	out = order.Uint16(d.data[d.pos:])
	d.pos += TypeSize.Uint16
	if traceEnabled {
		zlog.Debug("decode: read uint16", zap.Uint16("val", out))
	}
	return
}

func (d *Decoder) ReadInt16(order binary.ByteOrder) (out int16, err error) {
	n, err := d.ReadUint16(order)
	out = int16(n)
	if traceEnabled {
		zlog.Debug("decode: read int16", zap.Int16("val", out))
	}
	return
}

func (d *Decoder) ReadInt64(order binary.ByteOrder) (out int64, err error) {
	n, err := d.ReadUint64(order)
	out = int64(n)
	if traceEnabled {
		zlog.Debug("decode: read int64", zap.Int64("val", out))
	}
	return
}

func (d *Decoder) ReadUint32(order binary.ByteOrder) (out uint32, err error) {
	if d.Remaining() < TypeSize.Uint32 {
		err = fmt.Errorf("uint32 required [%d] bytes, remaining [%d]", TypeSize.Uint32, d.Remaining())
		return
	}

	out = order.Uint32(d.data[d.pos:])
	d.pos += TypeSize.Uint32
	if traceEnabled {
		zlog.Debug("decode: read uint32", zap.Uint32("val", out))
	}
	return
}

func (d *Decoder) ReadInt32(order binary.ByteOrder) (out int32, err error) {
	n, err := d.ReadUint32(order)
	out = int32(n)
	if traceEnabled {
		zlog.Debug("decode: read int32", zap.Int32("val", out))
	}
	return
}

func (d *Decoder) ReadUint64(order binary.ByteOrder) (out uint64, err error) {
	if d.Remaining() < TypeSize.Uint64 {
		err = fmt.Errorf("decode: uint64 required [%d] bytes, remaining [%d]", TypeSize.Uint64, d.Remaining())
		return
	}

	data := d.data[d.pos : d.pos+TypeSize.Uint64]
	out = order.Uint64(data)
	d.pos += TypeSize.Uint64
	if traceEnabled {
		zlog.Debug("decode: read uint64", zap.Uint64("val", out), zap.Stringer("hex", HexBytes(data)))
	}
	return
}

func (d *Decoder) ReadInt128(order binary.ByteOrder) (out Int128, err error) {
	v, err := d.ReadUint128(order)
	if err != nil {
		return
	}

	return Int128(v), nil
}

func (d *Decoder) ReadUint128(order binary.ByteOrder) (out Uint128, err error) {
	if d.Remaining() < TypeSize.Uint128 {
		err = fmt.Errorf("uint128 required [%d] bytes, remaining [%d]", TypeSize.Uint128, d.Remaining())
		return
	}

	data := d.data[d.pos : d.pos+TypeSize.Uint128]

	if order == binary.LittleEndian {
		out.Lo = order.Uint64(data)
		out.Hi = order.Uint64(data[8:])
	} else {
		out.Lo = order.Uint64(data[8:])
		out.Hi = order.Uint64(data)
	}

	d.pos += TypeSize.Uint128
	if traceEnabled {
		zlog.Debug("decode: read uint128", zap.Stringer("hex", out), zap.Uint64("hi", out.Hi), zap.Uint64("lo", out.Lo))
	}
	return
}

func (d *Decoder) ReadFloat32(order binary.ByteOrder) (out float32, err error) {
	if d.Remaining() < TypeSize.Float32 {
		err = fmt.Errorf("float32 required [%d] bytes, remaining [%d]", TypeSize.Float32, d.Remaining())
		return
	}

	n := order.Uint32(d.data[d.pos:])
	out = math.Float32frombits(n)
	d.pos += TypeSize.Float32
	if traceEnabled {
		zlog.Debug("decode: read float32", zap.Float32("val", out))
	}
	return
}

func (d *Decoder) ReadFloat64(order binary.ByteOrder) (out float64, err error) {
	if d.Remaining() < TypeSize.Float64 {
		err = fmt.Errorf("float64 required [%d] bytes, remaining [%d]", TypeSize.Float64, d.Remaining())
		return
	}

	n := order.Uint64(d.data[d.pos:])
	out = math.Float64frombits(n)
	d.pos += TypeSize.Float64
	if traceEnabled {
		zlog.Debug("decode: read Float64", zap.Float64("val", out))
	}
	return
}

func (d *Decoder) ReadFloat128(order binary.ByteOrder) (out Float128, err error) {
	value, err := d.ReadUint128(order)
	if err != nil {
		return out, fmt.Errorf("float128: %s", err)
	}

	return Float128(value), nil
}

func (d *Decoder) SafeReadUTF8String() (out string, err error) {
	data, err := d.ReadByteArray()
	out = strings.Map(fixUtf, string(data))
	if traceEnabled {
		zlog.Debug("read safe UTF8 string", zap.String("val", out))
	}
	return
}

func fixUtf(r rune) rune {
	if r == utf8.RuneError {
		return '�'
	}
	return r
}

func (d *Decoder) ReadString() (out string, err error) {
	data, err := d.ReadByteArray()
	out = string(data)
	if traceEnabled {
		zlog.Debug("read string", zap.String("val", out))
	}
	return
}

func (d *Decoder) SkipBytes(count uint) error {
	if uint(d.Remaining()) < count {
		return fmt.Errorf("request to skip %d but only %d bytes remain", count, d.Remaining())
	}
	d.pos += int(count)
	return nil
}

func (d *Decoder) SetPosition(idx uint) error {
	if int(idx) < len(d.data) {
		d.pos = int(idx)
		return nil
	}
	return fmt.Errorf("request to set position to %d outsize of buffer (buffer size %d)", idx, len(d.data))
}

func (d *Decoder) Position() uint {
	return uint(d.pos)
}

func (d *Decoder) Remaining() int {
	return len(d.data) - d.pos
}

func (d *Decoder) HasRemaining() bool {
	return d.Remaining() > 0
}

// indirect walks down v allocating pointers as needed,
// until it gets to a non-pointer.
// if it encounters an Unmarshaler, indirect stops and returns that.
// if decodingNull is true, indirect stops at the last pointer so it can be set to nil.
//
// *Note* This is a copy of `encoding/json/decoder.go#indirect` of Golang 1.14.
//
// See here: https://github.com/golang/go/blob/go1.14.2/src/encoding/json/decode.go#L439
func indirect(v reflect.Value, decodingNull bool) (UnmarshalerBinary, reflect.Value) {
	// Issue #24153 indicates that it is generally not a guaranteed property
	// that you may round-trip a reflect.Value by calling Value.Addr().Elem()
	// and expect the value to still be settable for values derived from
	// unexported embedded struct fields.
	//
	// The logic below effectively does this when it first addresses the value
	// (to satisfy possible pointer methods) and continues to dereference
	// subsequent pointers as necessary.
	//
	// After the first round-trip, we set v back to the original value to
	// preserve the original RW flags contained in reflect.Value.
	v0 := v
	haveAddr := false

	// If v is a named type and is addressable,
	// start with its address, so that if the type has pointer methods,
	// we find them.
	if v.Kind() != reflect.Ptr && v.Type().Name() != "" && v.CanAddr() {
		haveAddr = true
		v = v.Addr()
	}
	for {
		// Load value from interface, but only if the result will be
		// usefully addressable.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Ptr && !e.IsNil() && (!decodingNull || e.Elem().Kind() == reflect.Ptr) {
				haveAddr = false
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Ptr {
			break
		}

		if v.Elem().Kind() != reflect.Ptr && decodingNull && v.CanSet() {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v interface{}
		//     v = &v
		if v.Elem().Kind() == reflect.Interface && v.Elem().Elem() == v {
			v = v.Elem()
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		if v.Type().NumMethod() > 0 && v.CanInterface() {
			if u, ok := v.Interface().(UnmarshalerBinary); ok {
				return u, reflect.Value{}
			}
		}

		if haveAddr {
			v = v0 // restore original value after round-trip Value.Addr().Elem()
			haveAddr = false
		} else {
			v = v.Elem()
		}
	}
	return nil, v
}
