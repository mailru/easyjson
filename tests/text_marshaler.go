package tests

import (
	"strings"
)

//easyjson:json
type StructWrappedTextMarshaler struct {
	Value namedWithTextMarshaler
}
type namedWithTextMarshaler string

func (n namedWithTextMarshaler) MarshalText() ([]byte, error) {
	return []byte(strings.ToUpper(string(n))), nil
}
