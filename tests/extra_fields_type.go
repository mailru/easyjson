package tests

type ExtraFieldsType struct {
	Field1 int                    `json:",omitempty"`
	Field2 int                    `json:"f2,omitempty"`
	Extra  map[string]interface{} `json:"-,extra"`
}
