package tests

import (
	"testing"

	"github.com/mailru/easyjson"
)

func TestStringIntern(t *testing.T) {
	data := []byte(`{"field": "string interning test"}`)

	var i Intern
	allocsPerRun := testing.AllocsPerRun(1000, func() {
		i = Intern{}
		easyjson.Unmarshal(data, &i)
		if i.Field != "string interning test" {
			t.Fatalf("wrong value: %q", i.Field)
		}
	})
	if allocsPerRun != 1 {
		t.Fatalf("expected 1 allocs, got %f", allocsPerRun)
	}

	var n NoIntern
	allocsPerRun = testing.AllocsPerRun(1000, func() {
		n = NoIntern{}
		easyjson.Unmarshal(data, &n)
		if n.Field != "string interning test" {
			t.Fatalf("wrong value: %q", n.Field)
		}
	})
	if allocsPerRun != 2 {
		t.Fatalf("expected 2 allocs, got %f", allocsPerRun)
	}
}
