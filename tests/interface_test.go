package tests

import (
	"bytes"
	"testing"
)

func TestInterfaceMarshal(t *testing.T) {
	inner := InterfaceType{
		Field1: 1,
	}
	wrapper := WrapperType{
		Inner: inner,
	}

	data, err := wrapper.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}

	expected := []byte(`{"Inner":{"Field1":1}}`)
	if !bytes.Equal(data, expected) {
		t.Fatalf("MarshalJSON failed: got=%s, expected=%s", string(data), string(expected))
	}
}

func TestInterfaceUnmarshal(t *testing.T) {
	data := []byte(`{"Inner":{"Field1":1}}`)

	var inner InterfaceType
	wrapper := WrapperType{Inner: &inner}
	err := wrapper.UnmarshalJSON(data)
	if err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	if inner.Field1 != 1 {
		t.Fatalf("UnmarshalJSON failed: got=%d, expected=%d", inner.Field1, 1)
	}
}

func TestInterfaceMarshalNil(t *testing.T) {
	var inner *InterfaceType
	wrapper := WrapperType{
		Inner: inner,
	}

	data, err := wrapper.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}

	expected := []byte(`{"Inner":null}`)
	if !bytes.Equal(data, expected) {
		t.Fatalf("MarshalJSON failed: got=%s, expected=%s", string(data), string(expected))
	}
}

func TestInterfaceUnmarshalNil(t *testing.T) {
	data := []byte(`{"Inner":null}`)

	var inner InterfaceType
	wrapper := WrapperType{Inner: &inner}
	err := wrapper.UnmarshalJSON(data)
	if err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	if inner.Field1 != 0 {
		t.Fatalf("UnmarshalJSON failed: got=%d, expected=0", inner.Field1)
	}
}
