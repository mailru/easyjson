package tests

import (
	"reflect"
	"testing"
)

func TestUnmarshalExtraFields(t *testing.T) {
	golden := []struct {
		Serial string
		Object ExtraFieldsType
	}{
		{`{"Field1":1}`, ExtraFieldsType{Field1: 1,
			Extra: nil,
		}},
		{`{"Field1":1,"FieldX":"2"}`, ExtraFieldsType{Field1: 1,
			Extra: map[string]interface{}{"FieldX": "2"},
		}},
		{`{"FieldX":3}`, ExtraFieldsType{Field1: 0,
			Extra: map[string]interface{}{"FieldX": 3.0},
		}},
	}

	for _, gold := range golden {
		var o ExtraFieldsType
		err := o.UnmarshalJSON([]byte(gold.Serial))
		if err != nil {
			t.Errorf("unmarshal %s error: %s", gold.Serial, err)
			continue
		}
		if !reflect.DeepEqual(o, gold.Object) {
			t.Errorf("unmarshal %s got %#v, want %#v", gold.Serial, o, gold.Object)
		}
	}

	for _, gold := range golden {
		bytes, err := gold.Object.MarshalJSON()
		if err != nil {
			t.Errorf("marshal %#v error: %s", gold.Object, err)
			continue
		}
		if string(bytes) != gold.Serial {
			t.Errorf("unmarshal %#v got %s\nwant %s", gold.Object, bytes, gold.Serial)
		}
	}
}

func TestExtraFieldsOverride(t *testing.T) {
	o := ExtraFieldsType{
		Extra: map[string]interface{}{
			"Field1": 42,
			"f2":     99, // override custom name Field2
			"Field3": 100,
		},
	}
	bytes, err := o.MarshalJSON()
	if err != nil {
		t.Fatal("marshal error:", err)
	}
	if want := `{"Field3":100}`; string(bytes) != want {
		t.Errorf("marshal %#v got %s\nwant %s", o, bytes, want)
	}
}
