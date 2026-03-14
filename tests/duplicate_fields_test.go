package tests

import (
	"testing"

	"github.com/mailru/easyjson"
)

func TestDuplicateFieldError(t *testing.T) {
	var d Dupl
	err := easyjson.Unmarshal([]byte(`{"a": 1, "a": 2}`), &d)
	if err == nil || err.Error() != "duplicate field A" {
		t.Error("want error, got ", err)
	}
}

func TestDuplicateFieldNoError(t *testing.T) {
	var d Dupl
	err := easyjson.Unmarshal([]byte(`{"a": 1, "b": 2}`), &d)
	if err != nil {
		t.Error("unexpected error ", err)
	}
	if d != (Dupl{A: 1, B: 2}) {
		t.Error("unexpected value ", d)
	}
}

func TestDuplicateFieldNoErrorWithoutOption(t *testing.T) {
	// We use the DisallowUnknown struct, which is generated without the duplicate field check enabled,
	// so we expect to get no error and the last given value.
	var d DisallowUnknown
	err := easyjson.Unmarshal([]byte(`{"field_one": "a", "field_one": "b"}`), &d)
	if err != nil {
		t.Error("unexpected error ", err)
	}
	if d != (DisallowUnknown{FieldOne: "b"}) {
		t.Error("unexpected value ", d)
	}
}
