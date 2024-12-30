package easyjson

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func BenchmarkNilCheck(b *testing.B) {
	var a *int
	for i := 0; i < b.N; i++ {
		if !isNilInterface(a) {
			b.Fatal("expected it to be nil")
		}
	}
}

func TestMarshallingAny(t *testing.T) {
	testStruct := []struct {
		FieldString   string
		FieldInt      int
		FieldFloat    float64
		FieldBool     bool `json:"test_bool"`
		FieldDuration time.Duration
	}{{
		FieldString:   "HelloWorld",
		FieldInt:      50246,
		FieldFloat:    454654.4646,
		FieldBool:     true,
		FieldDuration: 90 * time.Hour,
	}, {}}

	testStruct0 := testStruct[0]
	testStruct1 := testStruct[1]
	if reflect.DeepEqual(&testStruct0, &testStruct1) {
		t.Fatal("expected it to be not equal")
	}

	b, err := MarshalAny(&testStruct0)
	if err != nil {
		t.Fatal(err)
	}
	err = UnmarshalAny(b, &testStruct1)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(&testStruct0, &testStruct1) {
		fmt.Println(testStruct0, testStruct1)
		t.Fatal("expected it to be equal")
	}
}
