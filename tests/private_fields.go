package tests

//easyjson:json
type PrivateFieldStruct struct {
	name    string
	process bool `json:"private_process"`
	Public  bool
}

var (
	myPrivateFieldStructValue  = PrivateFieldStruct{name: "hello", process: true, Public: true}
	myPrivateFieldStructString = `{"name":"hello","private_process":true,"Public":true}`
)
