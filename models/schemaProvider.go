package models

type SchemaProvider struct {
	Id         uint     `json:"id"`
	SchemaId   Schema   `json:"schemaid"`
	ProviderId Provider `json:"providerid"`
}
