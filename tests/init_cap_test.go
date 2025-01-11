package tests

import (
	"github.com/mailru/easyjson"
	"testing"
)

func TestTagInitialCapacity(t *testing.T) {
	data := []byte(`{"field":["elem"]}`)

	var n NoInitCap
	err := easyjson.Unmarshal(data, &n)
	if err != nil {
		t.Error(err)
	}

	if cap(n.Field) != 4 {
		t.Fatalf("wrong slice capacity: %d", cap(n.Field))
	}

	var c InitCap
	err = easyjson.Unmarshal(data, &c)
	if err != nil {
		t.Error(err)
	}

	if cap(c.Field) != 10 {
		t.Fatalf("wrong slice capacity: %d", cap(c.Field))
	}

	var i InvalidInitCap
	err = easyjson.Unmarshal(data, &i)
	if err != nil {
		t.Error(err)
	}

	if cap(i.Field) != 4 {
		t.Fatalf("wrong slice capacity: %d", cap(i.Field))
	}
}
