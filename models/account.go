package models

type Account struct {
	id       uint   `json:"id"`
	schemaId Schema `json:"schemaId"`
}
