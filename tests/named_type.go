package tests

//easyjson:json
type NamedType struct {
	Inner struct {
		// easyjson is mistakenly naming the type of this field 'tests.MyString' in the generated output
		// something about a named type inside an anonmymous type is triggering this bug
		Field MyString
	}
}

type MyString string

var namedTypeValue = NamedType{Inner: struct{ Field MyString }{Field: "test"}}
var namedTypeValueString = `{"Inner":{"Field":"test"}}`
