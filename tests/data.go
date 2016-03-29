package tests

import (
	"fmt"
	"math"
	"time"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/opt"
)

type PrimitiveTypes struct {
	String string
	Bool   bool

	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint   uint
	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64

	Float32 float32
	Float64 float64

	Ptr    *string
	PtrNil *string
}

var str = "bla"

var primitiveTypesValue = PrimitiveTypes{
	String: "test", Bool: true,

	Int:   math.MinInt32,
	Int8:  math.MinInt8,
	Int16: math.MinInt16,
	Int32: math.MinInt32,
	Int64: math.MinInt64,

	Uint:   math.MaxUint32,
	Uint8:  math.MaxUint8,
	Uint16: math.MaxUint16,
	Uint32: math.MaxUint32,
	Uint64: math.MaxUint64,

	Float32: 1.5,
	Float64: math.MaxFloat64,

	Ptr: &str,
}

var primitiveTypesString = "{" +
	`"String":"test","Bool":true,` +

	`"Int":` + fmt.Sprint(math.MinInt32) + `,` +
	`"Int8":` + fmt.Sprint(math.MinInt8) + `,` +
	`"Int16":` + fmt.Sprint(math.MinInt16) + `,` +
	`"Int32":` + fmt.Sprint(math.MinInt32) + `,` +
	`"Int64":` + fmt.Sprint(int64(math.MinInt64)) + `,` +

	`"Uint":` + fmt.Sprint(math.MaxUint32) + `,` +
	`"Uint8":` + fmt.Sprint(math.MaxUint8) + `,` +
	`"Uint16":` + fmt.Sprint(math.MaxUint16) + `,` +
	`"Uint32":` + fmt.Sprint(math.MaxUint32) + `,` +
	`"Uint64":` + fmt.Sprint(uint64(math.MaxUint64)) + `,` +

	`"Float32":` + fmt.Sprint(1.5) + `,` +
	`"Float64":` + fmt.Sprint(math.MaxFloat64) + `,` +

	`"Ptr":"bla",` +
	`"PtrNil":null` +

	"}"

type (
	NamedString string
	NamedBool   bool

	NamedInt   int
	NamedInt8  int8
	NamedInt16 int16
	NamedInt32 int32
	NamedInt64 int64

	NamedUint   uint
	NamedUint8  uint8
	NamedUint16 uint16
	NamedUint32 uint32
	NamedUint64 uint64

	NamedFloat32 float32
	NamedFloat64 float64

	NamedStrPtr *string
)

type NamedPrimitiveTypes struct {
	String NamedString
	Bool   NamedBool

	Int   NamedInt
	Int8  NamedInt8
	Int16 NamedInt16
	Int32 NamedInt32
	Int64 NamedInt64

	Uint   NamedUint
	Uint8  NamedUint8
	Uint16 NamedUint16
	Uint32 NamedUint32
	Uint64 NamedUint64

	Float32 NamedFloat32
	Float64 NamedFloat64

	Ptr    NamedStrPtr
	PtrNil NamedStrPtr
}

var namedPrimitiveTypesValue = NamedPrimitiveTypes{
	String: "test",
	Bool:   true,

	Int:   math.MinInt32,
	Int8:  math.MinInt8,
	Int16: math.MinInt16,
	Int32: math.MinInt32,
	Int64: math.MinInt64,

	Uint:   math.MaxUint32,
	Uint8:  math.MaxUint8,
	Uint16: math.MaxUint16,
	Uint32: math.MaxUint32,
	Uint64: math.MaxUint64,

	Float32: 1.5,
	Float64: math.MaxFloat64,

	Ptr: NamedStrPtr(&str),
}

var namedPrimitiveTypesString = "{" +
	`"String":"test",` +
	`"Bool":true,` +

	`"Int":` + fmt.Sprint(math.MinInt32) + `,` +
	`"Int8":` + fmt.Sprint(math.MinInt8) + `,` +
	`"Int16":` + fmt.Sprint(math.MinInt16) + `,` +
	`"Int32":` + fmt.Sprint(math.MinInt32) + `,` +
	`"Int64":` + fmt.Sprint(int64(math.MinInt64)) + `,` +

	`"Uint":` + fmt.Sprint(math.MaxUint32) + `,` +
	`"Uint8":` + fmt.Sprint(math.MaxUint8) + `,` +
	`"Uint16":` + fmt.Sprint(math.MaxUint16) + `,` +
	`"Uint32":` + fmt.Sprint(math.MaxUint32) + `,` +
	`"Uint64":` + fmt.Sprint(uint64(math.MaxUint64)) + `,` +

	`"Float32":` + fmt.Sprint(1.5) + `,` +
	`"Float64":` + fmt.Sprint(math.MaxFloat64) + `,` +

	`"Ptr":"bla",` +
	`"PtrNil":null` +
	"}"

type SubStruct struct {
	Value     string
	Value2    string
	unexpored bool
}

type SubStructAlias SubStruct

type Structs struct {
	SubStruct
	Value2 int

	Sub1   SubStruct `json:"substruct"`
	Sub2   *SubStruct
	SubNil *SubStruct

	SubA1 SubStructAlias
	SubA2 *SubStructAlias

	Anonymous struct {
		V string
		I int
	}
	Anonymous1 *struct {
		V string
	}
	unexported bool
}

var structsValue = Structs{
	SubStruct: SubStruct{Value: "test"},
	Value2:    5,

	Sub1: SubStruct{Value: "test1", Value2: "v"},
	Sub2: &SubStruct{Value: "test2", Value2: "v2"},

	SubA1: SubStructAlias{Value: "test3", Value2: "v3"},
	SubA2: &SubStructAlias{Value: "test4", Value2: "v4"},

	Anonymous: struct {
		V string
		I int
	}{V: "bla", I: 5},
	Anonymous1: &struct {
		V string
	}{V: "bla1"},
}

var structsString = "{" +
	`"Value2":5,` +

	`"substruct":{"Value":"test1","Value2":"v"},` +
	`"Sub2":{"Value":"test2","Value2":"v2"},` +
	`"SubNil":null,` +

	`"SubA1":{"Value":"test3","Value2":"v3"},` +
	`"SubA2":{"Value":"test4","Value2":"v4"},` +

	`"Anonymous":{"V":"bla","I":5},` +
	`"Anonymous1":{"V":"bla1"},` +

	// Embedded fields go last.
	`"Value":"test"` +
	"}"

type OmitEmpty struct {
	// NOTE: first field is empty to test comma printing.

	StrE, StrNE string  `json:",omitempty"`
	PtrE, PtrNE *string `json:",omitempty"`

	IntNE int `json:"intField,omitempty"`
	IntE  int `json:",omitempty"`

	// NOTE: omitempty has no effect on non-pointer struct fields.
	SubE, SubNE   SubStruct  `json:",omitempty"`
	SubPE, SubPNE *SubStruct `json:",omitempty"`
}

var omitEmptyValue = OmitEmpty{
	StrNE:  "str",
	PtrNE:  &str,
	IntNE:  6,
	SubNE:  SubStruct{Value: "1", Value2: "2"},
	SubPNE: &SubStruct{Value: "3", Value2: "4"},
}

var omitEmptyString = "{" +
	`"StrNE":"str",` +
	`"PtrNE":"bla",` +
	`"intField":6,` +
	`"SubE":{"Value":"","Value2":""},` +
	`"SubNE":{"Value":"1","Value2":"2"},` +
	`"SubPNE":{"Value":"3","Value2":"4"}` +
	"}"

type Opts struct {
	StrNull      opt.String
	StrEmpty     opt.String
	Str          opt.String
	StrOmitempty opt.String `json:",omitempty"`

	IntNull opt.Int
	IntZero opt.Int
	Int     opt.Int
}

var optsValue = Opts{
	StrEmpty: opt.OString(""),
	Str:      opt.OString("test"),

	IntZero: opt.OInt(0),
	Int:     opt.OInt(5),
}

var optsString = `{` +
	`"StrNull":null,` +
	`"StrEmpty":"",` +
	`"Str":"test",` +
	`"IntNull":null,` +
	`"IntZero":0,` +
	`"Int":5` +
	`}`

type Raw struct {
	Field  easyjson.RawMessage
	Field2 string
}

var rawValue = Raw{
	Field:  []byte(`{"a" : "b"}`),
	Field2: "test",
}

var rawString = `{` +
	`"Field":{"a" : "b"},` +
	`"Field2":"test"` +
	`}`

type StdMarshaler struct {
	T time.Time
}

var stdMarshalerValue = StdMarshaler{T: time.Date(2016, 01, 02, 14, 15, 10, 0, time.UTC)}
var stdMarshalerString = `{"T":"2016-01-02T14:15:10Z"}`
