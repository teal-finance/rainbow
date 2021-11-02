# StreamingFast binary
[![reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/streamingfast/binary)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Binary Encoder and Decoder library

Usage
----

#### Struct Field Tags
- `_` will skip the field when encoding & deocode a struc
- `sizeof=` indicates this field is a number used to track the length of a another field.
- `little` indicates this field will be encoded as little endian
- `big` indicates this field will be encoded as big endian
- `optional` indicates this field is optional. An optional field will have an initial byte either `0x00` or `0x01` indicating whether the field is present. If the field is present it will be followed by the value.
- Bare values will be parsed as type and little endian when necessary

#### Supported Types
 - `int8`, `int16`, `int32`, `int64`, `Int128`
 - `uint8`, `uint16`,`uint32`,`uint64`, `Uint128`
 - `float32`, `float64`, `Float128`
 - `string`, `bool`
 - `Varint16`, `Varint32`
 - `Varuint16`, `Varuint32`

#### Custom Types
To implement custom types, your types would need to implement the `MarshalerBinary` & `UnmarshalerBinary` interfaces

Example
----

#### Basic Implementation
```Go
type Example struct {
    Var      uint32 `bin:"_"`
    Str      string
    IntCount uint32 `bin:"sizeof=Var"`
    Weird    [8]byte
    Var      []int
}
```


#### Custom Implementation
```Go
type Example struct {
	Prefix byte
	Value  uint32
}

func (e *Example) UnmarshalBinary(decoder *Decoder) (err error) {
	if e.Prefix, err = decoder.ReadByte(); err != nil {
		return err
	}
	if e.Value, err = decoder.ReadUint32(BE()); err != nil {
		return err
	}
	return nil
}

func (e *Example) MarshalBinary(encoder *Encoder) error {
	if err := encoder.WriteByte(e.Prefix); err != nil {
		return err
	}
	return encoder.WriteUint32(e.Value, BE())
}
```

## Contributing

**Issues and PR in this repo related strictly to the binary library**

Report any protocol-specific issues in their
[respective repositories](https://github.com/streamingfast/streamingfast#protocols)

**Please first refer to the general
[StreamingFast contribution guide](https://github.com/streamingfast/streamingfast/blob/master/CONTRIBUTING.md)**,
if you wish to contribute to this code base.

This codebase uses unit tests extensively, please write and run tests.

## License

[Apache 2.0](LICENSE)
