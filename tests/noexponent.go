package tests

//easyjson:json
type NoExponentStruct struct {
	Scientific  float64 `json:"sci"`
	NoExp       float64 `json:"price,noexponent"`
	Mixed       float64 `json:"rate,omitempty,noexponent"`
	QuotedNoExp float64 `json:"str,string,noexponent"`
	Float32Test float32 `json:"f32,noexponent"`
}

var noExpStruct = NoExponentStruct{
	Scientific:  100000000,
	NoExp:       100000000,
	Mixed:       0.000001,
	QuotedNoExp: 1234.5678,
	Float32Test: 999.999,
}

var noExpString = `{"sci":1e+08,"price":100000000,"rate":0.000001,"str":"1234.5678","f32":999.999}`
