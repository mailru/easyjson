package gen

import (
	"testing"
)

type functionNamerCase struct {
	keepFirst     bool
	parts         []string
	camelCaseOut  string
	underScoreOut string
}

func TestCamelToSnake(t *testing.T) {
	for i, test := range []struct {
		In, Out string
	}{
		{"", ""},
		{"A", "a"},
		{"SimpleExample", "simple_example"},
		{"internalField", "internal_field"},

		{"SomeHTTPStuff", "some_http_stuff"},
		{"WriteJSON", "write_json"},
		{"HTTP2Server", "http2_server"},
		{"Some_Mixed_Case", "some_mixed_case"},
		{"do_nothing", "do_nothing"},

		{"JSONHTTPRPCServer", "jsonhttprpc_server"}, // nothing can be done here without a dictionary
	} {
		got := camelToSnake(test.In)
		if got != test.Out {
			t.Errorf("[%d] camelToSnake(%s) = %s; want %s", i, test.In, got, test.Out)
		}
	}
}

func getFunctionNamerCases() []functionNamerCase {
	return []functionNamerCase{
		functionNamerCase{false, []string{}, "", ""},
		functionNamerCase{false, []string{"a"}, "A", "a"},
		functionNamerCase{false, []string{"simple", "example"}, "SimpleExample", "simple_example"},
		functionNamerCase{true, []string{"first", "example"}, "firstExample", "first_example"},
		functionNamerCase{false, []string{"some", "UPPER", "case"}, "SomeUPPERCase", "some_UPPER_case"},
		functionNamerCase{false, []string{"number", "123"}, "Number123", "number_123"},
	}
}

func TestCamelCaseFunctionNamer(t *testing.T) {
	namer := CamelCaseFunctionNamer{}
	for i, test := range getFunctionNamerCases() {
		got := namer.GetName(test.keepFirst, test.parts...)
		if got != test.camelCaseOut {
			t.Errorf("[%d] CamelCaseFunctionNamer.GetName(%v) = %s; want %s", i, test.parts, got, test.camelCaseOut)
		}
	}
}

func TestUnderScoreFunctionNamer(t *testing.T) {
	namer := UnderScoreFunctionNamer{}
	for i, test := range getFunctionNamerCases() {
		got := namer.GetName(test.keepFirst, test.parts...)
		if got != test.underScoreOut {
			t.Errorf("[%d] UnderScoreFunctionNamer.GetName(%v) = %s; want %s", i, test.parts, got, test.underScoreOut)
		}
	}
}
