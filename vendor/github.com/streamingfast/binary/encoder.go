package bin

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"reflect"

	"go.uber.org/zap"
)

type Encoder struct {
	output io.Writer
	count  int

	currentFieldOpt *option
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		output: w,
		count:  0,
	}
}

func (e *Encoder) Encode(v interface{}) (err error) {
	return e.encode(reflect.ValueOf(v), nil)
}

func (e *Encoder) encode(rv reflect.Value, opt *option) (err error) {
	if opt == nil {
		opt = newDefaultOption()
	}
	e.currentFieldOpt = opt

	if traceEnabled {
		zlog.Debug("encode: type",
			zap.Stringer("value_kind", rv.Kind()),
			zap.Reflect("options", opt),
		)
	}

	if opt.isOptional() {
		if rv.IsZero() {
			if traceEnabled {
				zlog.Debug("encode: skipping optional value with", zap.Stringer("type", rv.Kind()))
			}
			e.WriteBool(false)
			return nil
		}
		e.WriteBool(true)
	}

	if isZero(rv) {
		return nil
	}

	if marshaler, ok := rv.Interface().(MarshalerBinary); ok {
		if traceEnabled {
			zlog.Debug("encode: using MarshalerBinary method to encode type")
		}
		return marshaler.MarshalBinary(e)
	}

	switch rv.Kind() {
	case reflect.String:
		return e.WriteString(rv.String())
	case reflect.Uint8, reflect.Int8:
		return e.WriteByte(byte(rv.Uint()))
	case reflect.Int16:
		return e.WriteInt16(int16(rv.Int()), opt.Order)
	case reflect.Uint16:
		return e.WriteUint16(uint16(rv.Uint()), opt.Order)
	case reflect.Int32:
		return e.WriteInt32(int32(rv.Int()), opt.Order)
	case reflect.Uint32:
		return e.WriteUint32(uint32(rv.Uint()), opt.Order)
	case reflect.Uint64:
		return e.WriteUint64(rv.Uint(), opt.Order)
	case reflect.Int64:
		return e.WriteInt64(rv.Int(), opt.Order)
	case reflect.Float32:
		return e.WriteFloat32(float32(rv.Float()), opt.Order)
	case reflect.Float64:
		return e.WriteFloat64(rv.Float(), opt.Order)
	case reflect.Bool:
		return e.WriteBool(rv.Bool())
	case reflect.Ptr:
		return e.encode(rv.Elem(), opt)
	}

	rv = reflect.Indirect(rv)
	rt := rv.Type()
	switch rt.Kind() {
	case reflect.Array:
		l := rt.Len()
		if traceEnabled {
			defer func(prev *zap.Logger) { zlog = prev }(zlog)
			zlog = zlog.Named("array")
			zlog.Debug("encode: array", zap.Int("length", l), zap.Stringer("type", rv.Kind()))
		}
		for i := 0; i < l; i++ {
			if err = e.encode(rv.Index(i), nil); err != nil {
				return
			}
		}
	case reflect.Slice:
		var l int
		if opt.hasSizeOfSlice() {
			l = opt.getSizeOfSlice()
			if traceEnabled {
				zlog.Debug("encode: slice with sizeof set", zap.Int("size_of", l))
			}
		} else {
			l = rv.Len()
			if err = e.WriteUVarInt(l); err != nil {
				return
			}
		}
		if traceEnabled {
			defer func(prev *zap.Logger) { zlog = prev }(zlog)
			zlog = zlog.Named("slice")
			zlog.Debug("encode: slice", zap.Int("length", l), zap.Stringer("type", rv.Kind()))
		}

		// we would want to skip to the correct head_offset

		for i := 0; i < l; i++ {
			if err = e.encode(rv.Index(i), opt); err != nil {
				return
			}
		}
	case reflect.Struct:
		if err = e.encodeStruct(rt, rv); err != nil {
			return
		}

	case reflect.Map:
		keyCount := len(rv.MapKeys())

		if traceEnabled {
			zlog.Debug("encode: map",
				zap.Int("key_count", keyCount),
				zap.String("key_type", rt.String()),
				typeField("value_type", rv.Elem()),
			)
			defer func(prev *zap.Logger) { zlog = prev }(zlog)
			zlog = zlog.Named("struct")
		}

		if err = e.WriteUVarInt(keyCount); err != nil {
			return
		}

		for _, mapKey := range rv.MapKeys() {
			if err = e.Encode(mapKey.Interface()); err != nil {
				return
			}

			if err = e.Encode(rv.MapIndex(mapKey).Interface()); err != nil {
				return
			}
		}

	default:
		return fmt.Errorf("encode: unsupported type %q", rt)
	}
	return
}

func (e *Encoder) toWriter(bytes []byte) (err error) {
	e.count += len(bytes)

	if traceEnabled {
		zlog.Debug("	> encode: appending", zap.Stringer("hex", HexBytes(bytes)), zap.Int("pos", e.count))
	}

	_, err = e.output.Write(bytes)
	return
}

func (e *Encoder) WriteByteArray(b []byte, writeLength bool) error {
	if traceEnabled {
		zlog.Debug("encode: write byte array", zap.Int("len", len(b)))
	}
	if writeLength {
		if err := e.WriteUVarInt(len(b)); err != nil {
			return err
		}
	}
	return e.toWriter(b)
}

func (e *Encoder) WriteUVarInt(v int) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write uvarint", zap.Int("val", v))
	}

	buf := make([]byte, 8)
	l := binary.PutUvarint(buf, uint64(v))
	return e.toWriter(buf[:l])
}

func (e *Encoder) WriteVarInt(v int) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write varint", zap.Int("val", v))
	}

	buf := make([]byte, 8)
	l := binary.PutVarint(buf, int64(v))
	return e.toWriter(buf[:l])
}

func (e *Encoder) WriteByte(b byte) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write byte", zap.Uint8("val", b))
	}
	return e.toWriter([]byte{b})
}

func (e *Encoder) WriteBool(b bool) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write bool", zap.Bool("val", b))
	}
	var out byte
	if b {
		out = 1
	}
	return e.WriteByte(out)
}

func (e *Encoder) WriteUint8(i uint8) (err error) {
	return e.WriteByte(i)
}

func (e *Encoder) WriteUint16(i uint16, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write uint16", zap.Uint16("val", i))
	}
	buf := make([]byte, TypeSize.Uint16)
	order.PutUint16(buf, i)
	return e.toWriter(buf)
}

func (e *Encoder) WriteInt16(i int16, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write int16", zap.Int16("val", i))
	}
	return e.WriteUint16(uint16(i), order)
}

func (e *Encoder) WriteInt32(i int32, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write int32", zap.Int32("val", i))
	}
	return e.WriteUint32(uint32(i), order)
}

func (e *Encoder) WriteUint32(i uint32, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write uint32", zap.Uint32("val", i))
	}
	buf := make([]byte, TypeSize.Uint32)
	order.PutUint32(buf, i)
	return e.toWriter(buf)
}

func (e *Encoder) WriteInt64(i int64, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write int64", zap.Int64("val", i))
	}
	return e.WriteUint64(uint64(i), order)
}

func (e *Encoder) WriteUint64(i uint64, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write uint64", zap.Uint64("val", i))
	}
	buf := make([]byte, TypeSize.Uint64)
	order.PutUint64(buf, i)
	return e.toWriter(buf)
}

func (e *Encoder) WriteUint128(i Uint128, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write uint128", zap.Stringer("hex", i), zap.Uint64("lo", i.Lo), zap.Uint64("hi", i.Hi))
	}
	buf := make([]byte, TypeSize.Uint128)
	order.PutUint64(buf, i.Lo)
	order.PutUint64(buf[TypeSize.Uint64:], i.Hi)
	return e.toWriter(buf)
}

func (e *Encoder) WriteInt128(i Int128, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write int128", zap.Stringer("hex", i), zap.Uint64("lo", i.Lo), zap.Uint64("hi", i.Hi))
	}
	buf := make([]byte, TypeSize.Uint128)
	order.PutUint64(buf, i.Lo)
	order.PutUint64(buf[TypeSize.Uint64:], i.Hi)
	return e.toWriter(buf)
}

func (e *Encoder) WriteFloat32(f float32, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write float32", zap.Float32("val", f))
	}
	i := math.Float32bits(f)
	buf := make([]byte, TypeSize.Uint32)
	order.PutUint32(buf, i)

	return e.toWriter(buf)
}
func (e *Encoder) WriteFloat64(f float64, order binary.ByteOrder) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write float64", zap.Float64("val", f))
	}
	i := math.Float64bits(f)
	buf := make([]byte, TypeSize.Uint64)
	order.PutUint64(buf, i)

	return e.toWriter(buf)
}

func (e *Encoder) WriteString(s string) (err error) {
	if traceEnabled {
		zlog.Debug("encode: write string", zap.String("val", s))
	}
	return e.WriteByteArray([]byte(s), true)
}

func (e *Encoder) encodeStruct(rt reflect.Type, rv reflect.Value) (err error) {
	l := rv.NumField()

	if traceEnabled {
		zlog.Debug("encode: struct", zap.Int("fields", l), zap.Stringer("type", rv.Kind()))
	}

	sizeOfMap := map[string]int{}
	for i := 0; i < l; i++ {
		structField := rt.Field(i)
		fieldTag := parseFieldTag(structField.Tag)

		if fieldTag.Skip {
			if traceEnabled {
				zlog.Debug("encode: skipping struct field with skip flag",
					zap.String("struct_field_name", structField.Name),
				)
			}
			continue
		}

		rv := rv.Field(i)

		if fieldTag.SizeOf != "" {
			if traceEnabled {
				zlog.Debug("encode: struct field has sizeof tag",
					zap.String("sizeof_field_name", fieldTag.SizeOf),
					zap.String("struct_field_name", structField.Name),
				)
			}
			sizeOfMap[fieldTag.SizeOf] = sizeof(structField.Type, rv)
		}

		if !rv.CanInterface() {
			if traceEnabled {
				zlog.Debug("encode:  skipping field: unable to interface field, probably since field is not exported",
					zap.String("sizeof_field_name", fieldTag.SizeOf),
					zap.String("struct_field_name", structField.Name),
				)
			}
			continue
		}

		option := &option{
			OptionalField: fieldTag.Optional,
			Order:         fieldTag.Order,
		}

		if s, ok := sizeOfMap[structField.Name]; ok {
			if traceEnabled {
				zlog.Debug("setting sizeof option", zap.String("of", structField.Name), zap.Int("size", s))
			}
			option.setSizeOfSlice(s)
		}

		if traceEnabled {
			zlog.Debug("encode: struct field",
				zap.Stringer("struct_field_value_type", rv.Kind()),
				zap.String("struct_field_name", structField.Name),
				zap.Reflect("struct_field_tags", fieldTag),
				zap.Reflect("struct_field_option", option),
			)
		}

		if err := e.encode(rv, option); err != nil {
			return err
		}
	}
	return nil
}
