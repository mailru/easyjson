package tests

import "encoding/json"

//easyjson:json
type JSONNumberStruct struct {
	Int   json.Number `json:"int"`
	Float json.Number `json:"float"`
}

var jsonNumberValue = JSONNumberStruct{
	Int:   json.Number("123"),
	Float: json.Number("10.00"),
}

// json.Number must be encoded as a raw JSON number, not a quoted string,
// matching encoding/json (see issue #220).
var jsonNumberString = `{"int":123,"float":10.00}`
