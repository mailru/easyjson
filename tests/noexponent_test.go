package tests

import (
	"testing"
)

func TestNoExponentMarshal(t *testing.T) {
	data, err := noExpStruct.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON() error: %v", err)
		return
	}

	got := string(data)
	if got != noExpString {
		t.Errorf("MarshalJSON() = %v, want %v", got, noExpString)
	}
}

func TestNoExponentUnmarshal(t *testing.T) {
	var got NoExponentStruct
	err := got.UnmarshalJSON([]byte(noExpString))
	if err != nil {
		t.Errorf("UnmarshalJSON() error: %v", err)
		return
	}

	if got.Scientific != noExpStruct.Scientific {
		t.Errorf("Scientific = %v, want %v", got.Scientific, noExpStruct.Scientific)
	}
	if got.NoExp != noExpStruct.NoExp {
		t.Errorf("NoExp = %v, want %v", got.NoExp, noExpStruct.NoExp)
	}
	if got.Mixed != noExpStruct.Mixed {
		t.Errorf("Mixed = %v, want %v", got.Mixed, noExpStruct.Mixed)
	}
	if got.QuotedNoExp != noExpStruct.QuotedNoExp {
		t.Errorf("QuotedNoExp = %v, want %v", got.QuotedNoExp, noExpStruct.QuotedNoExp)
	}
	if got.Float32Test != noExpStruct.Float32Test {
		t.Errorf("Float32Test = %v, want %v", got.Float32Test, noExpStruct.Float32Test)
	}
}
