package tests

import (
	"testing"

	"github.com/mailru/easyjson/jlexer"
)

func TestSemanticErrorsInt(t *testing.T) {
	if !*jlexer.UseSemanticErrors {
		return
	}
	for i, test := range []struct {
		Data     []byte
		ErrorNum int
	}{
		{
			Data:     []byte(`[1, 2, 3, "4", "5"]`),
			ErrorNum: 2,
		},
		{
			Data:     []byte(`[1, {"2" : "3"}, 3, "4"`),
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
	} {
		l := jlexer.Lexer{Data: test.Data}

		var v ErrorIntSlice

		v.UnmarshalEasyJSON(&l)

		if len(l.SemanticErrors) != test.ErrorNum {
			t.Errorf("[%d] TestSemanticErrors(): errornum: want: %d, got %d", i, test.ErrorNum, len(l.SemanticErrors))
		}
	}
}

func TestSemanticErrorsBool(t *testing.T) {
	if !*jlexer.UseSemanticErrors {
		return
	}
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
			Data:     []byte(`[true, 42, {"a":"b", "c":"d"}, false`),
			ErrorNum: 2,
		},
	} {
		l := jlexer.Lexer{Data: test.Data}

		var v ErrorBoolSlice
		v.UnmarshalEasyJSON(&l)

		if len(l.SemanticErrors) != test.ErrorNum {
			t.Errorf("[%d] TestSemanticErrors(): errornum: want: %d, got %d", i, test.ErrorNum, len(l.SemanticErrors))
		}
	}
}

func TestSemanticErrorsUint(t *testing.T) {
	if !*jlexer.UseSemanticErrors {
		return
	}

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
	} {
		l := jlexer.Lexer{Data: test.Data}

		var v ErrorUintSlice
		v.UnmarshalEasyJSON(&l)

		if len(l.SemanticErrors) != test.ErrorNum {
			t.Errorf("[%d] TestSemanticErrors(): errornum: want: %d, got %d", i, test.ErrorNum, len(l.SemanticErrors))
		}
	}
}
