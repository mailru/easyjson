package jwriter

import (
	"encoding/json"
	"testing"
)

func TestJsonNumber(t *testing.T) {
	for _, test := range []struct {
		number  json.Number
		want    string
		wantErr bool
	}{
		{number: "123", want: "123"},
		{number: "-0.5", want: "-0.5"},
		{number: "10.00", want: "10.00"},
		{number: "25E-4", want: "25E-4"},
		// encoding/json encodes the empty Number as 0.
		{number: "", want: "0"},
		// Invalid literals set the writer error, like encoding/json errors.
		{number: "abc", wantErr: true},
		{number: "1.", wantErr: true},
		{number: "-", wantErr: true},
	} {
		var w Writer
		w.JsonNumber(test.number)
		if test.wantErr {
			if w.Error == nil {
				t.Errorf("JsonNumber(%q): expected error, got none", test.number)
			}
			continue
		}
		if w.Error != nil {
			t.Errorf("JsonNumber(%q): unexpected error: %v", test.number, w.Error)
			continue
		}
		if got := string(w.Buffer.BuildBytes()); got != test.want {
			t.Errorf("JsonNumber(%q) = %s; want %s", test.number, got, test.want)
		}
	}
}
