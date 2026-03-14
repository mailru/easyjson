package tests

// easyjson:json
type WrapperType struct {
	Inner any `json:"Inner"`
}

// easyjson:json
type InterfaceType struct {
	Field1 int `json:"Field1"`
}
