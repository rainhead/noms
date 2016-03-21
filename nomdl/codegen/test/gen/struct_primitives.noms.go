// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_primitives_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("StructPrimitives",
			[]types.Field{
				types.Field{"uint64", types.MakePrimitiveType(types.Uint64Kind), false},
				types.Field{"uint32", types.MakePrimitiveType(types.Uint32Kind), false},
				types.Field{"uint16", types.MakePrimitiveType(types.Uint16Kind), false},
				types.Field{"uint8", types.MakePrimitiveType(types.Uint8Kind), false},
				types.Field{"int64", types.MakePrimitiveType(types.Int64Kind), false},
				types.Field{"int32", types.MakePrimitiveType(types.Int32Kind), false},
				types.Field{"int16", types.MakePrimitiveType(types.Int16Kind), false},
				types.Field{"int8", types.MakePrimitiveType(types.Int8Kind), false},
				types.Field{"float64", types.MakePrimitiveType(types.Float64Kind), false},
				types.Field{"float32", types.MakePrimitiveType(types.Float32Kind), false},
				types.Field{"bool", types.MakePrimitiveType(types.BoolKind), false},
				types.Field{"string", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"blob", types.MakePrimitiveType(types.BlobKind), false},
				types.Field{"value", types.MakePrimitiveType(types.ValueKind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_primitives_CachedRef = types.RegisterPackage(&p)
}

// StructPrimitives

type StructPrimitives struct {
	_uint64  uint64
	_uint32  uint32
	_uint16  uint16
	_uint8   uint8
	_int64   int64
	_int32   int32
	_int16   int16
	_int8    int8
	_float64 float64
	_float32 float32
	_bool    bool
	_string  string
	_blob    types.Blob
	_value   types.Value

	ref *ref.Ref
}

func NewStructPrimitives() StructPrimitives {
	return StructPrimitives{
		_uint64:  uint64(0),
		_uint32:  uint32(0),
		_uint16:  uint16(0),
		_uint8:   uint8(0),
		_int64:   int64(0),
		_int32:   int32(0),
		_int16:   int16(0),
		_int8:    int8(0),
		_float64: float64(0),
		_float32: float32(0),
		_bool:    false,
		_string:  "",
		_blob:    types.NewEmptyBlob(),
		_value:   types.Bool(false),

		ref: &ref.Ref{},
	}
}

type StructPrimitivesDef struct {
	Uint64  uint64
	Uint32  uint32
	Uint16  uint16
	Uint8   uint8
	Int64   int64
	Int32   int32
	Int16   int16
	Int8    int8
	Float64 float64
	Float32 float32
	Bool    bool
	String  string
	Blob    types.Blob
	Value   types.Value
}

func (def StructPrimitivesDef) New() StructPrimitives {
	return StructPrimitives{
		_uint64:  def.Uint64,
		_uint32:  def.Uint32,
		_uint16:  def.Uint16,
		_uint8:   def.Uint8,
		_int64:   def.Int64,
		_int32:   def.Int32,
		_int16:   def.Int16,
		_int8:    def.Int8,
		_float64: def.Float64,
		_float32: def.Float32,
		_bool:    def.Bool,
		_string:  def.String,
		_blob:    def.Blob,
		_value:   def.Value,
		ref:      &ref.Ref{},
	}
}

func (s StructPrimitives) Def() (d StructPrimitivesDef) {
	d.Uint64 = s._uint64
	d.Uint32 = s._uint32
	d.Uint16 = s._uint16
	d.Uint8 = s._uint8
	d.Int64 = s._int64
	d.Int32 = s._int32
	d.Int16 = s._int16
	d.Int8 = s._int8
	d.Float64 = s._float64
	d.Float32 = s._float32
	d.Bool = s._bool
	d.String = s._string
	d.Blob = s._blob
	d.Value = s._value
	return
}

var __typeForStructPrimitives types.Type

func (m StructPrimitives) Type() types.Type {
	return __typeForStructPrimitives
}

func init() {
	__typeForStructPrimitives = types.MakeType(__genPackageInFile_struct_primitives_CachedRef, 0)
	types.RegisterStruct(__typeForStructPrimitives, builderForStructPrimitives, readerForStructPrimitives)
}

func builderForStructPrimitives(values []types.Value) types.Value {
	i := 0
	s := StructPrimitives{ref: &ref.Ref{}}
	s._uint64 = uint64(values[i].(types.Uint64))
	i++
	s._uint32 = uint32(values[i].(types.Uint32))
	i++
	s._uint16 = uint16(values[i].(types.Uint16))
	i++
	s._uint8 = uint8(values[i].(types.Uint8))
	i++
	s._int64 = int64(values[i].(types.Int64))
	i++
	s._int32 = int32(values[i].(types.Int32))
	i++
	s._int16 = int16(values[i].(types.Int16))
	i++
	s._int8 = int8(values[i].(types.Int8))
	i++
	s._float64 = float64(values[i].(types.Float64))
	i++
	s._float32 = float32(values[i].(types.Float32))
	i++
	s._bool = bool(values[i].(types.Bool))
	i++
	s._string = values[i].(types.String).String()
	i++
	s._blob = values[i].(types.Blob)
	i++
	s._value = values[i]
	i++
	return s
}

func readerForStructPrimitives(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(StructPrimitives)
	values = append(values, types.Uint64(s._uint64))
	values = append(values, types.Uint32(s._uint32))
	values = append(values, types.Uint16(s._uint16))
	values = append(values, types.Uint8(s._uint8))
	values = append(values, types.Int64(s._int64))
	values = append(values, types.Int32(s._int32))
	values = append(values, types.Int16(s._int16))
	values = append(values, types.Int8(s._int8))
	values = append(values, types.Float64(s._float64))
	values = append(values, types.Float32(s._float32))
	values = append(values, types.Bool(s._bool))
	values = append(values, types.NewString(s._string))
	values = append(values, s._blob)
	values = append(values, s._value)
	return values
}

func (s StructPrimitives) Equals(other types.Value) bool {
	return other != nil && __typeForStructPrimitives.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s StructPrimitives) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s StructPrimitives) Chunks() (chunks []types.RefBase) {
	chunks = append(chunks, __typeForStructPrimitives.Chunks()...)
	chunks = append(chunks, s._blob.Chunks()...)
	chunks = append(chunks, s._value.Chunks()...)
	return
}

func (s StructPrimitives) ChildValues() (ret []types.Value) {
	ret = append(ret, types.Uint64(s._uint64))
	ret = append(ret, types.Uint32(s._uint32))
	ret = append(ret, types.Uint16(s._uint16))
	ret = append(ret, types.Uint8(s._uint8))
	ret = append(ret, types.Int64(s._int64))
	ret = append(ret, types.Int32(s._int32))
	ret = append(ret, types.Int16(s._int16))
	ret = append(ret, types.Int8(s._int8))
	ret = append(ret, types.Float64(s._float64))
	ret = append(ret, types.Float32(s._float32))
	ret = append(ret, types.Bool(s._bool))
	ret = append(ret, types.NewString(s._string))
	ret = append(ret, s._blob)
	ret = append(ret, s._value)
	return
}

func (s StructPrimitives) Uint64() uint64 {
	return s._uint64
}

func (s StructPrimitives) SetUint64(val uint64) StructPrimitives {
	s._uint64 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Uint32() uint32 {
	return s._uint32
}

func (s StructPrimitives) SetUint32(val uint32) StructPrimitives {
	s._uint32 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Uint16() uint16 {
	return s._uint16
}

func (s StructPrimitives) SetUint16(val uint16) StructPrimitives {
	s._uint16 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Uint8() uint8 {
	return s._uint8
}

func (s StructPrimitives) SetUint8(val uint8) StructPrimitives {
	s._uint8 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Int64() int64 {
	return s._int64
}

func (s StructPrimitives) SetInt64(val int64) StructPrimitives {
	s._int64 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Int32() int32 {
	return s._int32
}

func (s StructPrimitives) SetInt32(val int32) StructPrimitives {
	s._int32 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Int16() int16 {
	return s._int16
}

func (s StructPrimitives) SetInt16(val int16) StructPrimitives {
	s._int16 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Int8() int8 {
	return s._int8
}

func (s StructPrimitives) SetInt8(val int8) StructPrimitives {
	s._int8 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Float64() float64 {
	return s._float64
}

func (s StructPrimitives) SetFloat64(val float64) StructPrimitives {
	s._float64 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Float32() float32 {
	return s._float32
}

func (s StructPrimitives) SetFloat32(val float32) StructPrimitives {
	s._float32 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Bool() bool {
	return s._bool
}

func (s StructPrimitives) SetBool(val bool) StructPrimitives {
	s._bool = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) String() string {
	return s._string
}

func (s StructPrimitives) SetString(val string) StructPrimitives {
	s._string = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Blob() types.Blob {
	return s._blob
}

func (s StructPrimitives) SetBlob(val types.Blob) StructPrimitives {
	s._blob = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Value() types.Value {
	return s._value
}

func (s StructPrimitives) SetValue(val types.Value) StructPrimitives {
	s._value = val
	s.ref = &ref.Ref{}
	return s
}
