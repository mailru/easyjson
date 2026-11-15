package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
)

func TestMarshalEmbeddedPointers(t *testing.T) {
	var unknowns easyjson.UnknownFieldsProxy
	lexer := jlexer.Lexer{Data: []byte("1")}
	unknowns.UnmarshalUnknown(&lexer, "extra")
	if err := lexer.Error(); err != nil {
		t.Fatal(err)
	}

	for _, test := range []struct {
		name string
		in   json.Marshaler
		want string
	}{
		{"nil", EmbeddedPointerOnly{}, `{}`},
		{"zero", EmbeddedPointerOnly{&EmbeddedPointerFields{}}, `{"value":""}`},
		{"populated", EmbeddedPointerOnly{&EmbeddedPointerFields{Value: "v", Optional: "o", Zero: 1, Text: "custom", Dual: 1}}, `{"value":"v","optional":"o","zero":1,"text":"CUSTOM","dual":1}`},
		{"nonzero optional", EmbeddedPointerOnly{&EmbeddedPointerFields{Dual: -1}}, `{"value":"","dual":-1}`},
		{"nil with own field", EmbeddedPointerWithField{Name: "own"}, `{"name":"own"}`},
		{"nil with omitted own field", EmbeddedPointerWithField{}, `{}`},
		{"present with own field", EmbeddedPointerWithField{&EmbeddedPointerFields{Value: "v"}, "own"}, `{"name":"own","value":"v"}`},
		{"present with omitted own field", EmbeddedPointerWithField{EmbeddedPointerFields: &EmbeddedPointerFields{Value: "v"}}, `{"value":"v"}`},
		{"nil outer pointer", NestedEmbeddedPointer{}, `{}`},
		{"nil inner pointer", NestedEmbeddedPointer{&EmbeddedPointerMiddle{Level: "middle"}}, `{"level":"middle"}`},
		{"nil inner pointer and omitted field", NestedEmbeddedPointer{&EmbeddedPointerMiddle{}}, `{}`},
		{"nested populated", NestedEmbeddedPointer{&EmbeddedPointerMiddle{&EmbeddedPointerFields{Value: "v"}, "middle"}}, `{"level":"middle","value":"v"}`},
		{"value embedding nil pointer", ValueEmbeddedPointer{}, `{}`},
		{"value embedding present pointer", ValueEmbeddedPointer{EmbeddedPointerMiddle{&EmbeddedPointerFields{Value: "v"}, ""}}, `{"value":"v"}`},
		{"shadowed nil", ShadowedEmbeddedPointer{Value: "own"}, `{"value":"own"}`},
		{"shadowed populated", ShadowedEmbeddedPointer{&EmbeddedPointerFields{Value: "hidden", Optional: "o"}, "own"}, `{"value":"own","optional":"o"}`},
		{"named nil", NamedEmbeddedPointer{}, `{"inner":null}`},
		{"named populated", NamedEmbeddedPointer{&EmbeddedPointerFields{Value: "v"}}, `{"inner":{"value":"v"}}`},
		{"both nil", MultipleEmbeddedPointers{}, `{}`},
		{"first present", MultipleEmbeddedPointers{EmbeddedPointerFields: &EmbeddedPointerFields{Value: "v"}}, `{"value":"v"}`},
		{"second present", MultipleEmbeddedPointers{EmbeddedPointerOther: &EmbeddedPointerOther{Other: "o"}}, `{"other":"o"}`},
		{"both present", MultipleEmbeddedPointers{&EmbeddedPointerFields{Value: "v"}, &EmbeddedPointerOther{Other: "o"}}, `{"other":"o","value":"v"}`},
		{"private nil", PrivateEmbeddedPointer{}, `{}`},
		{"private present", PrivateEmbeddedPointer{&embeddedPrivatePointerFields{Private: "p"}}, `{"private":"p"}`},
		{"unknowns after nil", EmbeddedPointerWithUnknowns{UnknownFieldsProxy: unknowns}, `{"extra":1}`},
		{"unknowns after present", EmbeddedPointerWithUnknowns{unknowns, &EmbeddedPointerFields{Value: "v"}}, `{"value":"v","extra":1}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Marshal panicked: %v", r)
				}
			}()
			data, err := test.in.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}
			var got, want map[string]interface{}
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatalf("invalid JSON %s: %v", data, err)
			}
			if err := json.Unmarshal([]byte(test.want), &want); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Marshal() = %s; want %s", data, test.want)
			}
		})
	}
}
