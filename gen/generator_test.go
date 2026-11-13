package gen

import (
	"bytes"
	"errors"
	"go/parser"
	"go/token"
	"testing"
)

func TestRunWritesCompleteSourceToWriter(t *testing.T) {
	g := NewGenerator("output.go")
	g.SetPkg("output", "example.com/output")
	var out bytes.Buffer
	if err := g.Run(&out); err != nil {
		t.Fatal(err)
	}
	if _, err := parser.ParseFile(token.NewFileSet(), "output.go", out.Bytes(), parser.AllErrors); err != nil {
		t.Fatalf("writer did not receive a complete Go file: %v", err)
	}
}

type nonEmptyErrorWriter struct {
	err error
}

func (w nonEmptyErrorWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	return 0, w.err
}

func TestRunReturnsHeaderWriteError(t *testing.T) {
	g := NewGenerator("output.go")
	g.SetPkg("output", "example.com/output")
	want := errors.New("header write failed")
	if err := g.Run(nonEmptyErrorWriter{err: want}); !errors.Is(err, want) {
		t.Fatalf("Run returned %v, want %v", err, want)
	}
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

func TestCamelToLowerCamel(t *testing.T) {
	for i, test := range []struct {
		In, Out string
	}{
		{"", ""},
		{"A", "a"},
		{"SimpleExample", "simpleExample"},
		{"internalField", "internalField"},

		{"SomeHTTPStuff", "someHTTPStuff"},
		{"WriteJSON", "writeJSON"},
		{"HTTP2Server", "http2Server"},

		{"JSONHTTPRPCServer", "jsonhttprpcServer"}, // nothing can be done here without a dictionary
	} {
		got := lowerFirst(test.In)
		if got != test.Out {
			t.Errorf("[%d] lowerFirst(%s) = %s; want %s", i, test.In, got, test.Out)
		}
	}
}

func TestJoinFunctionNameParts(t *testing.T) {
	for i, test := range []struct {
		keepFirst bool
		parts     []string
		out       string
	}{
		{false, []string{}, ""},
		{false, []string{"a"}, "A"},
		{false, []string{"simple", "example"}, "SimpleExample"},
		{true, []string{"first", "example"}, "firstExample"},
		{false, []string{"some", "UPPER", "case"}, "SomeUPPERCase"},
		{false, []string{"number", "123"}, "Number123"},
	} {
		got := joinFunctionNameParts(test.keepFirst, test.parts...)
		if got != test.out {
			t.Errorf("[%d] joinFunctionNameParts(%v) = %s; want %s", i, test.parts, got, test.out)
		}
	}
}

func TestFixVendorPath(t *testing.T) {
	for i, test := range []struct {
		In, Out string
	}{
		{"", ""},
		{"time", "time"},
		{"project/vendor/subpackage", "subpackage"},
	} {
		got := fixPkgPathVendoring(test.In)
		if got != test.Out {
			t.Errorf("[%d] fixPkgPathVendoring(%s) = %s; want %s", i, test.In, got, test.Out)
		}
	}

}
