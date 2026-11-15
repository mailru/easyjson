package tests

import "github.com/mailru/easyjson"

//easyjson:json
type EmbeddedType struct {
	EmbeddedInnerType
	Inner struct {
		EmbeddedInnerType
	}
	Field2             int
	EmbeddedInnerType2 `json:"named"`
}

type EmbeddedInnerType struct {
	Field1 int
}

type EmbeddedInnerType2 struct {
	Field3 int
}

var embeddedTypeValue EmbeddedType

func init() {
	embeddedTypeValue.Field1 = 1
	embeddedTypeValue.Field2 = 2
	embeddedTypeValue.Inner.Field1 = 3
	embeddedTypeValue.Field3 = 4
}

var embeddedTypeValueString = `{"Inner":{"Field1":3},"Field2":2,"named":{"Field3":4},"Field1":1}`

type EmbeddedPointerFields struct {
	Value    string                  `json:"value"`
	Optional string                  `json:"optional,omitempty"`
	Zero     int                     `json:"zero,omitzero"`
	Text     namedWithTextMarshaler  `json:"text,omitempty"`
	Dual     embeddedPointerOptional `json:"dual,omitempty,omitzero"`
}

type embeddedPointerOptional int

func (v embeddedPointerOptional) IsDefined() bool { return v > 0 }
func (v embeddedPointerOptional) IsZero() bool    { return v == 0 }

//easyjson:json
type EmbeddedPointerOnly struct {
	*EmbeddedPointerFields
}

//easyjson:json
type EmbeddedPointerWithField struct {
	*EmbeddedPointerFields
	Name string `json:"name,omitempty"`
}

type EmbeddedPointerMiddle struct {
	*EmbeddedPointerFields
	Level string `json:"level,omitempty"`
}

//easyjson:json
type NestedEmbeddedPointer struct {
	*EmbeddedPointerMiddle
}

//easyjson:json
type ValueEmbeddedPointer struct {
	EmbeddedPointerMiddle
}

//easyjson:json
type ShadowedEmbeddedPointer struct {
	*EmbeddedPointerFields
	Value string `json:"value"`
}

//easyjson:json
type NamedEmbeddedPointer struct {
	*EmbeddedPointerFields `json:"inner"`
}

type EmbeddedPointerOther struct {
	Other string `json:"other"`
}

//easyjson:json
type MultipleEmbeddedPointers struct {
	*EmbeddedPointerFields
	*EmbeddedPointerOther
}

type embeddedPrivatePointerFields struct {
	Private string `json:"private"`
}

//easyjson:json
type PrivateEmbeddedPointer struct {
	*embeddedPrivatePointerFields
}

//easyjson:json
type EmbeddedPointerWithUnknowns struct {
	easyjson.UnknownFieldsProxy
	*EmbeddedPointerFields
}
