package gen

import (
	"testing"
)

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
