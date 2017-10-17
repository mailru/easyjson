package tests

import (
	"os"
	"testing"

	"github.com/mailru/easyjson/gen"
)

type IntMap map[int]string
type IntMapSlice []IntMap
type IntMapArray [2]IntMap
type IntMapPtr *IntMap
type IntMapMap map[string]IntMap

func TestNonStringKeyedtMapEncoder(t *testing.T) {
	f := "non_string_keyed_map_easyjson.go"
	for _, test := range []struct {
		Data interface{}
	}{
		{
			Data: IntMap{},
		},
		{
			Data: IntMapSlice{},
		},
		{
			Data: IntMapArray{},
		},
		{
			Data: IntMapPtr(nil),
		},
		{
			Data: IntMapMap{},
		},
	} {
		g := gen.NewGenerator(f)
		g.Add(test.Data)
		e := g.Run(os.Stdout)
		if e == nil {
			t.Errorf("generation for %#v should have errored", test.Data)
		}
	}
}
