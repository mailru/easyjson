package tests

//easyjson:json
type FloatFmtStruct struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

var floatFmtStruct = FloatFmtStruct{
	A: 100000000,
	B: 1000,
}
var floatFmtString = `{"a":100000000,"b":1000}`
