package tests

//easyjson:json
type CamelCasesFunctions struct {
	Field string
}

var camelCasesFunctionsValue = CamelCasesFunctions{Field: "test"}
var camelCasesFunctionsString = `{"Field":"test"}`
