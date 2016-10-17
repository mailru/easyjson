package tests

import (
	"fmt"
	"testing"
)

func TestRequiredField(t *testing.T) {
	cases := []struct{ json, errorMessage string }{
		{`{"first_name":"Foo", "last_name": "Bar"}`, ""},
		{`{"last_name":"Bar"}`, "key 'first_name' is required"},
		{"{}", "key 'first_name' is required"},
	}

	for _, tc := range cases {
		var v RequiredOptionalStruct
		err := v.UnmarshalJSON([]byte(tc.json))
		if tc.errorMessage == "" {
			if err != nil {
				t.Errorf("%s. UnmarshalJSON didn`t expect error: %v", tc.json, err)
			}
		} else {
			if fmt.Sprintf("%v", err) != tc.errorMessage {
				t.Errorf("%s. UnmarshalJSON expected error: %v. got: %v", tc.json, tc.errorMessage, err)
			}
		}
	}
}

func TestDefinable(t *testing.T) {
	cases := []struct {
		json    string
		defined bool
	}{
		{`{"a":"foo"}`, true},
		{`{"a":"foo", "b":"bar"}`, true},
		{"{}", true},
		{"null", false},
	}

	for _, tc := range cases {
		var v Definable
		v.UnmarshalJSON([]byte(tc.json))
		if v.Def != tc.defined {
			t.Errorf("UnmarshalJSON expected: %q. got: %q", tc.defined, v.Def)
		}
	}
}
