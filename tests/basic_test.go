package tests

import (
	"reflect"
	"testing"

	"encoding/json"
)

type testType interface {
	json.Marshaler
	json.Unmarshaler
}

var testCases = []struct {
	Decoded testType
	Encoded string
}{
	{&primitiveTypesValue, primitiveTypesString},
	{&structsValue, structsString},
	{&omitEmptyValue, omitEmptyString},
	{&snakeStructValue, snakeStructString},
	{&omitEmptyDefaultValue, omitEmptyDefaultString},
	{&optsValue, optsString},
	{&rawValue, rawString},
	{&stdMarshalerValue, stdMarshalerString},
}

func TestMarshal(t *testing.T) {
	for i, test := range testCases {
		data, err := test.Decoded.MarshalJSON()
		if err != nil {
			t.Errorf("[%d, %T] MarshalJSON() error: %v", i, test.Decoded, err)
		}

		got := string(data)
		if got != test.Encoded {
			t.Errorf("[%d, %T] MarshalJSON(): got \n%v\n\t\t want \n%v", i, test.Decoded, got, test.Encoded)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	for i, test := range testCases {
		v1 := reflect.New(reflect.TypeOf(test.Decoded).Elem()).Interface()
		v := v1.(testType)

		err := v.UnmarshalJSON([]byte(test.Encoded))
		if err != nil {
			t.Errorf("[%d, %T] UnmarshalJSON() error: %v", i, test.Decoded, err)
		}

		if !reflect.DeepEqual(v, test.Decoded) {
			t.Errorf("[%d, %T] UnmarshalJSON(): got \n%+v\n\t\t want \n%+v", i, test.Decoded, v, test.Encoded)
		}
	}
}
