package tinytest

import (
	"testing"
)

// Encode and decode types
func TestEnvEncoding(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected Env
		output   string
	}{
		"empty": {
			input:    `{"contract":{}}`,
			expected: Env{},
			output:   `{"block":{"height":0,"time":"0"},"contract":{"address":""}}`,
		},
		"full": {
			input: `{"contract":{"address":"my-contract"},"block":{"time":"1234567890","height":42}}`,
			expected: Env{
				Contract: ContractInfo{Address: "my-contract"},
				Block:    BlockInfo{Time: 1234567890, Height: 42},
			},
			output: `{"block":{"height":42,"time":"1234567890"},"contract":{"address":"my-contract"}}`,
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

			// check proper output
			bz, err := loaded.MarshalJSON()
			if err != nil {
				t.Fatalf("Marshalling: %s", err)
			}
			if string(bz) != tc.output {
				t.Fatalf("Unexpected output: '%s'", string(bz))
			}
		})
	}
}

func TestInfoEncoding(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected MessageInfo
		output   string
	}{
		"empty": {
			input:    `{}`,
			expected: MessageInfo{},
			output:   `{"signer":"","funds":null}`,
		},
		"top fields": {
			input: `{"signer":"cosmos1234","funds":[]}`,
			expected: MessageInfo{
				Signer: "cosmos1234",
				Funds:  []Coin{},
			},
			output: `{"signer":"cosmos1234","funds":[]}`,
		},
		"multiple coins": {
			input: `{"signer":"cosmos1234","funds":[{"amount":"12345","denom":"uatom"},{"amount":"76543","denom":"utgd"}]}`,
			expected: MessageInfo{
				Signer: "cosmos1234",
				Funds: []Coin{{
					Amount: "12345",
					Denom:  "uatom",
				}, {
					Amount: "76543",
					Denom:  "utgd",
				}},
			},
			output: `{"signer":"cosmos1234","funds":[{"denom":"uatom","amount":"12345"},{"denom":"utgd","amount":"76543"}]}`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// check it matches expectations
			var loaded MessageInfo
			err := loaded.UnmarshalJSON([]byte(tc.input))
			if err != nil {
				t.Fatalf("Unmarshaling: %s", err)
			}
			if !loaded.Equals(&tc.expected) {
				t.Fatalf("Unmarshaled doesn't match expected.\nGot %#v\nWanted %#v", loaded, tc.expected)
			}

			// check it matches itself
			bz, err := loaded.MarshalJSON()
			if err != nil {
				t.Fatalf("Marshalling: %s", err)
			}
			if string(bz) != tc.output {
				t.Fatalf("Unexpected output: '%s'", string(bz))
			}
		})
	}
}

func TestMessageEncoding(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected ExecuteMsg
		output   string
	}{
		"deposit": {
			input:    `{"deposit":{"to_account":"cosmos1234567","amount":"1865"}}`,
			expected: ExecuteMsg{Deposit: &DepositMsg{ToAccount: "cosmos1234567", Amount: "1865"}},
			output:   `{"deposit":{"to_account":"cosmos1234567","amount":"1865"}}`,
		},
		"withdraw": {
			input:    `{"withdraw":{"from_account":"wasm1542"}}`,
			expected: ExecuteMsg{Withdraw: &WithdrawMsg{FromAccount: "wasm1542"}},
			output:   `{"withdraw":{"from_account":"wasm1542"}}`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// check it matches expectations
			var loaded ExecuteMsg
			err := loaded.UnmarshalJSON([]byte(tc.input))
			if err != nil {
				t.Fatalf("Unmarshaling: %s", err)
			}
			if !loaded.Equals(&tc.expected) {
				t.Fatalf("Unmarshaled doesn't match expected.\nGot %#v\nWanted %#v", loaded, tc.expected)
			}

			// check it matches itself
			bz, err := loaded.MarshalJSON()
			if err != nil {
				t.Fatalf("Marshalling: %s", err)
			}
			if string(bz) != tc.output {
				t.Fatalf("Unexpected output: '%s'", string(bz))
			}
		})
	}
}
