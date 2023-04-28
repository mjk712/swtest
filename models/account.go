package models

type Account struct {
	Id       uint `json:"id"`
	SchemaId uint `json:"schemaid"`
}

type AccSchmId struct {
	Schemaid uint `json:"schemaid"`
}
