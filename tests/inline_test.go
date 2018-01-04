package tests

import (
	"fmt"
	"testing"
)

func TestInlineMapField(t *testing.T) {
	cases := []struct{ json, errorMessage string }{
		{`{"first_name":"Foo","last_name":"Bar","x-x":"x"}`, ""},
	}

	for _, tc := range cases {
		var v InlineMapStruct
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
		if v.InlineMap == nil || v.InlineMap["x-x"] != "x" {
			t.Errorf("%s. UnmarshalJSON failed InlineMap is empty", tc.json)
		}
		data, err := v.MarshalJSON()
		if err != nil {
			t.Errorf("%v. MarshalJSON didn`t expect error: %v", v, err)
		}
		if tc.json != string(data) {
			t.Errorf("%v. MarshalJSON faield, expect %s. got %v", v, tc.json, data)
		}
	}
}
