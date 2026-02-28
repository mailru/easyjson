package tests

import (
	"testing"

	"github.com/mailru/easyjson"
)

func TestStructWithTextMarshalerMarshal(t *testing.T) {
	tests := []struct {
		name     string
		input    StructWrappedTextMarshaler
		expected string
	}{
		{
			name: "Filled struct",
			input: StructWrappedTextMarshaler{
				Value: namedWithTextMarshaler("test"),
			},
			expected: `{"Value":"TEST"}`,
		},
		{
			name:     "Empty struct",
			input:    StructWrappedTextMarshaler{},
			expected: `{"Value":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			marshaled, err := easyjson.Marshal(test.input)
			if err != nil {
				t.Errorf("easyjson.Marshal failed: %v", err)
			}
			if string(marshaled) != test.expected {
				t.Errorf("easyjson.Marshal output incorrect: got %s, want %s", string(marshaled), test.expected)
			}
		})
	}
}
