package tests

//easyjson:json
type StructWithUnknownsUnmarshaler struct {
	Field1 string

	unknownFields map[string]interface{}
}

func (s *StructWithUnknownsUnmarshaler) UnmarshalUnknown(key string, val interface{}) {
	if s.unknownFields == nil {
		s.unknownFields = make(map[string]interface{}, 1)
	}
	s.unknownFields[key] = val
}

func (s StructWithUnknownsUnmarshaler) MarshalUnknowns() map[string]interface{} {
	return s.unknownFields
}

var structWithUnknowns StructWithUnknownsUnmarshaler = StructWithUnknownsUnmarshaler{
	Field1: "123", unknownFields: map[string]interface{}{"Field2": "321"}}
var structWithUnknownsString = `{"Field1":"123","Field2":"321"}`
