package tests

import (
	"testing"

	"github.com/mailru/easyjson/jlexer"
)

func TestSemanticErrorsInt(t *testing.T) {
	for i, test := range []struct {
		Data     []byte
		ErrorNum int
	}{
		{
			Data:     []byte(`[1, 2, 3, "4", "5"]`),
			ErrorNum: 2,
		},
		{
			Data:     []byte(`[1, {"2":"3"}, 3, "4"]`),
			ErrorNum: 2,
		},
		{
			Data:     []byte(`[1, "2", "3", "4", "5", "6"]`),
			ErrorNum: 5,
		},
		{
			Data:     []byte(`[1, 2, 3, 4, "5"]`),
			ErrorNum: 1,
		},
		{
			Data:     []byte(`[{"1": "2"}]`),
			ErrorNum: 1,
		},
	} {
		l := jlexer.Lexer{
			Data:              test.Data,
			UseMultipleErrors: true,
		}

		var v ErrorIntSlice

		v.UnmarshalEasyJSON(&l)

		if len(l.GetSemanticErrors()) != test.ErrorNum {
			t.Errorf("[%d] TestSemanticErrorsInt(): errornum: want: %d, got %d", i, test.ErrorNum, len(l.GetSemanticErrors()))
			t.Errorf("%v", l.GetSemanticErrors())
			t.Errorf("%v", l.Error())
		}
	}
}

func TestSemanticErrorsBool(t *testing.T) {
	for i, test := range []struct {
		Data     []byte
		ErrorNum int
	}{
		{
			Data: []byte(`[true, false, true, false]`),
		},
		{
			Data:     []byte(`["test", "value", "lol", "1"]`),
			ErrorNum: 4,
		},
		{
			Data:     []byte(`[true, 42, {"a":"b", "c":"d"}, false]`),
			ErrorNum: 2,
		},
	} {
		l := jlexer.Lexer{
			Data:              test.Data,
			UseMultipleErrors: true,
		}

		var v ErrorBoolSlice
		v.UnmarshalEasyJSON(&l)

		if len(l.GetSemanticErrors()) != test.ErrorNum {
			t.Errorf("[%d] TestSemanticErrorsBool(): errornum: want: %d, got %d", i, test.ErrorNum, len(l.GetSemanticErrors()))
			t.Errorf("%v", l.Error())
		}
	}
}

func TestSemanticErrorsUint(t *testing.T) {
	for i, test := range []struct {
		Data     []byte
		ErrorNum int
	}{
		{
			Data: []byte(`[42, 42, 42]`),
		},
		{
			Data:     []byte(`[17, "42", 32]`),
			ErrorNum: 1,
		},
		{
			Data:     []byte(`["zz", "zz"]`),
			ErrorNum: 2,
		},
		{
			Data:     []byte(`[{}, 42]`),
			ErrorNum: 1,
		},
	} {
		l := jlexer.Lexer{
			Data:              test.Data,
			UseMultipleErrors: true,
		}

		var v ErrorUintSlice
		v.UnmarshalEasyJSON(&l)

		if len(l.GetSemanticErrors()) != test.ErrorNum {
			t.Errorf("[%d] TestSemanticErrorsUint(): errornum: want: %d, got %d", i, test.ErrorNum, len(l.GetSemanticErrors()))
		}
	}
}

func TestSemanticErrorsStruct(t *testing.T) {
	for i, test := range []struct {
		Data     []byte
		ErrorNum int
	}{
		{
			Data: []byte(`{"string": "test", "slice":[42, 42, 42], "int_slice":[1, 2, 3]}`),
		},
		{
			Data:     []byte(`{"string": {"test": "test"}, "slice":[42, 42, 42], "int_slice":["1", 2, 3]}`),
			ErrorNum: 2,
		},
		{
			Data:     []byte(`{"slice": [42, 42], "string": {"test": "test"}, "int_slice":["1", "2", 3]}`),
			ErrorNum: 3,
		},
		{
			Data:     []byte(`{"string": "test", "slice": {}}`),
			ErrorNum: 1,
		},
		{
			Data:     []byte(`{"slice":5, "string" : "test"}`),
			ErrorNum: 1,
		},
		{
			Data:     []byte(`{"slice" : "test", "string" : "test"}`),
			ErrorNum: 1,
		},
		{
			Data:     []byte(`{"slice": "", "string" : {}, "int":{}}`),
			ErrorNum: 3,
		},
	} {
		l := jlexer.Lexer{
			Data:              test.Data,
			UseMultipleErrors: true,
		}
		var v ErrorStruct
		v.UnmarshalEasyJSON(&l)

		if len(l.GetSemanticErrors()) != test.ErrorNum {
			t.Errorf("[%d] TestSemanticErrorsStruct(): errornum: want: %d, got %d", i, test.ErrorNum, len(l.GetSemanticErrors()))
		}
	}
}

func TestSemanticErrorsNestedStruct(t *testing.T) {
	for i, test := range []struct {
		Data     []byte
		ErrorNum int
	}{
		{
			Data: []byte(`{"error_struct":{}}`),
		},
		{
			Data:     []byte(`{"error_struct":5}`),
			ErrorNum: 1,
		},
		{
			Data:     []byte(`{"error_struct":[]}`),
			ErrorNum: 1,
		},
		{
			Data:     []byte(`{"error_struct":{"int":{}}}`),
			ErrorNum: 1,
		},
		{
			Data:     []byte(`{"error_struct":{"int_slice":{}}, "int":4}`),
			ErrorNum: 1,
		},
		{
			Data:     []byte(`{"error_struct":{"int_slice":["1", 2, "3"]}, "int":[]}`),
			ErrorNum: 3,
		},
	} {
		l := jlexer.Lexer{
			Data:              test.Data,
			UseMultipleErrors: true,
		}
		var v ErrorNestedStruct
		v.UnmarshalEasyJSON(&l)

		if len(l.GetSemanticErrors()) != test.ErrorNum {
			t.Errorf("[%d] TestSemanticErrorsNestedStruct(): errornum: want: %d, got %d", i, test.ErrorNum, len(l.GetSemanticErrors()))
		}
	}
}
