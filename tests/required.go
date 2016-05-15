package tests

//easyjson:json
type RequiredOptionalStruct struct {
	FirstName string `json:"first_name,required"`
	Lastname  string `json:"last_name"`
}
