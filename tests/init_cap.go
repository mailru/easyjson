package tests

//easyjson:json
type NoInitCap struct {
	Field []string `json:"field"`
}

//easyjson:json
type InitCap struct {
	Field []string `json:"field" initialCap:"10"`
}

//easyjson:json
type InvalidInitCap struct {
	Field []string `json:"field" initialCap:"-3"`
}

var noInitCap = NoInitCap{Field: []string{"elem"}}
var noInitCapString = `{"field":["elem"]}`

var initCap = InitCap{Field: []string{"elem"}}
var initCapString = `{"field":["elem"]}`

var invalidInitCap = InvalidInitCap{Field: []string{"elem"}}
var invalidInitCapString = `{"field":["elem"]}`
