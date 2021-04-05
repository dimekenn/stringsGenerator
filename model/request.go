package model

//request model
type NumberRequest struct {
	NumberOfStrings int `json:"num"`
}

//reguest model to encryptor service
type StringsRequest struct {
	Strings []string `json:"strings"`
}
