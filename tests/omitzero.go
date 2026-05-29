package tests

import "time"

//easyjson:json
type OmitZeroDefault struct {
	Field   string
	Str     string
	Str1    string               `json:"s,!omitzero"`
	Str2    string               `json:",!omitzero"`
	Time    time.Time            `json:",omitzero"` // implements IsZero() bool
	Array1  [2]int               `json:",omitzero"` // array filled with zero values is omitted
	Array2  [2]int               `json:",omitzero"` // array filled with non-zero values is outputed
	Array3  [2]OmitZeroSubstruct `json:",omitzero"`
	Array4  [2]OmitZeroSubstruct `json:",omitzero"`
	Struct1 OmitZeroSubstruct    `json:",omitzero"`
	Struct2 OmitZeroSubstruct    `json:",!omitzero"` // struct where all the fields are tagged omitzero
}

var omitZeroDefaultValue = OmitZeroDefault{Field: "test", Array2: [2]int{0, 1}, Array4: [2]OmitZeroSubstruct{{}, {F: 1}}}
var omitZeroDefaultString = `{"Field":"test","s":"","Str2":"","Array2":[0,1],"Array4":[{},{"F":1}],"Struct2":{}}`

type OmitZeroSubstruct struct {
	F float32   `json:",omitzero"`
	T time.Time `json:",omitzero"`
}
