package tests

//easyjson:json
type ErrorIntSlice []int

//easyjson:json
type ErrorBoolSlice []bool

//easyjson:json
type ErrorUintSlice []uint

//easyjson:json
type ErrorStruct struct {
	Int      int    `json:"int"`
	String   string `json:"string"`
	Slice    []int  `json:"slice"`
	IntSlice []int  `json:"int_slice"`
}

type ErrorNestedStruct struct {
	ErrorStruct ErrorStruct `json:"error_struct"`
	Int         int         `json:"int"`
}

//easyjson:json
type ErrorIntMap map[uint32]string

//easyjson:json
type ErrorFloatTypes struct {
	Float64 float64 `json:"float64"`
	Float32 float32 `json:"float32"`
}
