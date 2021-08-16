package tinytest

import (
	"testing"
)

// Encode and decode types
func TestEnvEncoding(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected Env
	}{
		"empty": {
			input:    `{"contract":{}}`,
			expected: Env{},
		},
		"full": {
			input: `{"contract":{"address":"my-contract"},"block":{"time":"1234567890","height":42}}`,
			expected: Env{
				Contract: ContractInfo{Address: "my-contract"},
				Block:    BlockInfo{Time: 1234567890, Height: 42},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// check it matches expectations
			var loaded Env
			err := loaded.UnmarshalJSON([]byte(tc.input))
			if err != nil {
				t.Fatalf("Unmarshaling: %s", err)
			}
			if loaded != tc.expected {
				t.Fatalf("Unmarshaled doesn't match expected.\nGot %#v\nWanted %#v", loaded, tc.expected)
			}

			// check it matches itself
			bz, err := loaded.MarshalJSON()
			if err != nil {
				t.Fatalf("Marshalling: %s", err)
			}
			var reloaded Env
			err = reloaded.UnmarshalJSON(bz)
			if err != nil {
				t.Fatalf("Unmarshaling: %s", err)
			}
			if reloaded != loaded {
				t.Fatalf("Full cycle changes data.\nHad %#v\nGot %#v", loaded, reloaded)
			}
		})
	}
}
