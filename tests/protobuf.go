package tests

//easyjson:json
type ProtobufStruct struct {
	JsonTagOnlyField bool   `json:"json_tag_only_field"`
	ProtobufField    string `protobuf:"bytes,4,opt,name=protobuf_field,json=protobufField" json:"protobuf_field,omitempty"`
}

var protobufStructValue = ProtobufStruct{ProtobufField: "some value"}
var protobufStructString = `{"json_tag_only_field":false,"protobufField":"some value"}`
