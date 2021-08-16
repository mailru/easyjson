package tinytest

import (
	"bytes"
	"testing"
)

var sampleEnv = Env{
	Contract: ContractInfo{Address: "wasm18vd8fpwxzck93qlwghaj6arh4p7c5n89k7fvsl"},
	// Time in nanoseconds since epoch start
	Block: BlockInfo{Height: 78000, Time: 1629131191823419000},
}

var sampleEnvText = []byte(`{"block":{"height":78000,"time":"1629131191823419000"},"contract":{"address":"wasm18vd8fpwxzck93qlwghaj6arh4p7c5n89k7fvsl"}}`)

var sampleMsgInfo = MessageInfo{
	Signer: "wasm18vd8fpwxzck93qlwghaj6arh4p7c5n89k7fvsl",
	Funds:  []Coin{{Amount: "123000000", Denom: "utgd"}},
}

var sampleMsgInfoText = []byte(`{"signer":"wasm18vd8fpwxzck93qlwghaj6arh4p7c5n89k7fvsl","funds":[{"denom":"utgd","amount":"123000000"}]}`)

func TestGenerateData(t *testing.T) {
	bz, _ := sampleEnv.MarshalJSON()
	// fmt.Println(string(bz))
	if !bytes.Equal(bz, sampleEnvText) {
		t.Fatal("Update sample text")
	}
	bz, _ = sampleMsgInfo.MarshalJSON()
	// fmt.Println(string(bz))
	if !bytes.Equal(bz, sampleMsgInfoText) {
		t.Fatal("Update sample text")
	}
}
