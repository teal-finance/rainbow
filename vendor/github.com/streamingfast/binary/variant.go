package bin

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"
)

//
/// Variant (emulates `fc::static_variant` type)
//

type Variant interface {
	Assign(typeID uint, impl interface{})
	Obtain() (typeID uint, impl interface{})
}

type VariantType struct {
	Name string
	Type interface{}
}

type VariantDefinition struct {
	typeIDToType   map[uint32]reflect.Type
	typeIDToName   map[uint32]string
	typeNameToID   map[string]uint32
	typeIDEncoding TypeIDEncoding
}

type TypeIDEncoding uint32

const (
	Uvarint32TypeIDEncoding TypeIDEncoding = iota
	Uint32TypeIDEncoding
	Uint8TypeIDEncoding
)

// NewVariantDefinition creates a variant definition based on the *ordered* provided types.
// It's the ordering that defines the binary variant value just like in native `nodeos` C++
// and in Smart Contract via the `std::variant` type. It's important to pass the entries
// in the right order!
//
// This variant definition can now be passed to functions of `BaseVariant` to implement
// marshal/unmarshaling functionalities for binary & JSON.
func NewVariantDefinition(typeIDEncoding TypeIDEncoding, types []VariantType) (out *VariantDefinition) {
	if len(types) < 0 {
		panic("it's not valid to create a variant definition without any types")
	}

	typeCount := len(types)
	out = &VariantDefinition{
		typeIDEncoding: typeIDEncoding,
		typeIDToType:   make(map[uint32]reflect.Type, typeCount),
		typeIDToName:   make(map[uint32]string, typeCount),
		typeNameToID:   make(map[string]uint32, typeCount),
	}

	for i, typeDef := range types {
		typeID := uint32(i)

		// FIXME: Check how the reflect.Type is used and cache all its usage in the definition.
		//        Right now, on each Unmarshal, we re-compute some expensive stuff that can be
		//        re-used like the `typeGo.Elem()` which is always the same. It would be preferable
		//        to have those already pre-defined here so we can actually speed up the
		//        Unmarshal code.
		out.typeIDToType[typeID] = reflect.TypeOf(typeDef.Type)
		out.typeIDToName[typeID] = typeDef.Name
		out.typeNameToID[typeDef.Name] = typeID
	}

	return out
}

func (d *VariantDefinition) TypeID(name string) uint32 {
	id, found := d.typeNameToID[name]
	if !found {
		knownNames := make([]string, len(d.typeNameToID))
		i := 0
		for name := range d.typeNameToID {
			knownNames[i] = name
			i++
		}

		panic(fmt.Errorf("trying to use an unknown type name %q, known names are %q", name, strings.Join(knownNames, ", ")))
	}

	return id
}

type VariantImplFactory = func() interface{}
type OnVariant = func(impl interface{}) error

type BaseVariant struct {
	TypeID uint32
	Impl   interface{}
}

func (a *BaseVariant) Assign(typeID uint32, impl interface{}) {
	a.TypeID = typeID
	a.Impl = impl
}

func (a *BaseVariant) Obtain(def *VariantDefinition) (typeID uint32, typeName string, impl interface{}) {
	return a.TypeID, def.typeIDToName[a.TypeID], a.Impl
}

func (a *BaseVariant) MarshalJSON(def *VariantDefinition) ([]byte, error) {
	typeName, found := def.typeIDToName[a.TypeID]
	if !found {
		return nil, fmt.Errorf("type %d is not know by variant definition", a.TypeID)
	}

	return json.Marshal([]interface{}{typeName, a.Impl})
}

func (a *BaseVariant) UnmarshalJSON(data []byte, def *VariantDefinition) error {
	typeResult := gjson.GetBytes(data, "0")
	implResult := gjson.GetBytes(data, "1")

	if !typeResult.Exists() || !implResult.Exists() {
		return fmt.Errorf("invalid format, expected '[<type>, <impl>]' pair, got %q", string(data))
	}

	typeName := typeResult.String()
	typeID, found := def.typeNameToID[typeName]
	if !found {
		return fmt.Errorf("type %q is not know by variant definition", typeName)
	}

	typeGo := def.typeIDToType[typeID]
	if typeGo == nil {
		return fmt.Errorf("no known type for %q", typeName)
	}

	a.TypeID = typeID

	if typeGo.Kind() == reflect.Ptr {
		a.Impl = reflect.New(typeGo.Elem()).Interface()
		if err := json.Unmarshal([]byte(implResult.Raw), a.Impl); err != nil {
			return err
		}
	} else {
		// This is not the most optimal way of doing things for "value"
		// types (over "pointer" types) as we always allocate a new pointer
		// element, unmarshal it and then either keep the pointer type or turn
		// it into a value type.
		//
		// However, in non-reflection based code, one would do like this and
		// avoid an `new` memory allocation:
		//
		// ```
		// name := eos.Name("")
		// json.Unmarshal(data, &name)
		// ```
		//
		// This would work without a problem. In reflection code however, I
		// did not find how one can go from `reflect.Zero(typeGo)` (which is
		// the equivalence of doing `name := eos.Name("")`) and take the
		// pointer to it so it can be unmarshalled correctly.
		//
		// A played with various iteration, and nothing got it working. Maybe
		// the next step would be to explore the `unsafe` package and obtain
		// an unsafe pointer and play with it.
		value := reflect.New(typeGo)
		if err := json.Unmarshal([]byte(implResult.Raw), value.Interface()); err != nil {
			return err
		}

		a.Impl = value.Elem().Interface()
	}

	return nil
}

func (a *BaseVariant) UnmarshalBinaryVariant(decoder *Decoder, def *VariantDefinition) (err error) {
	var typeID uint32
	switch def.typeIDEncoding {
	case Uvarint32TypeIDEncoding:
		typeID, err = decoder.ReadUvarint32()
		if err != nil {
			return fmt.Errorf("uvarint32: unable to read variant type id: %s", err)
		}
	case Uint32TypeIDEncoding:
		typeID, err = decoder.ReadUint32(LE())
		if err != nil {
			return fmt.Errorf("uint32: unable to read variant type id: %s", err)
		}
	case Uint8TypeIDEncoding:
		id, err := decoder.ReadUint8()
		if err != nil {
			return fmt.Errorf("uint8: unable to read variant type id: %s", err)
		}
		typeID = uint32(id)
	}

	a.TypeID = typeID

	typeGo := def.typeIDToType[typeID]
	if typeGo == nil {
		return fmt.Errorf("no known type for type %d", typeID)
	}

	if typeGo.Kind() == reflect.Ptr {
		a.Impl = reflect.New(typeGo.Elem()).Interface()
		if err = decoder.Decode(a.Impl); err != nil {
			return fmt.Errorf("unable to decode variant type %d: %s", typeID, err)
		}
	} else {
		// This is not the most optimal way of doing things for "value"
		// types (over "pointer" types) as we always allocate a new pointer
		// element, unmarshal it and then either keep the pointer type or turn
		// it into a value type.
		//
		// However, in non-reflection based code, one would do like this and
		// avoid an `new` memory allocation:
		//
		// ```
		// name := eos.Name("")
		// json.Unmarshal(data, &name)
		// ```
		//
		// This would work without a problem. In reflection code however, I
		// did not find how one can go from `reflect.Zero(typeGo)` (which is
		// the equivalence of doing `name := eos.Name("")`) and take the
		// pointer to it so it can be unmarshalled correctly.
		//
		// A played with various iteration, and nothing got it working. Maybe
		// the next step would be to explore the `unsafe` package and obtain
		// an unsafe pointer and play with it.
		value := reflect.New(typeGo)
		if err = decoder.Decode(value.Interface()); err != nil {
			return fmt.Errorf("unable to decode variant type %d: %s", typeID, err)
		}

		a.Impl = value.Elem().Interface()
	}
	return nil
}
