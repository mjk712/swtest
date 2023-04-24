package models

type SchemaProvider struct {
	id         uint     `json:"id"`
	schemaId   Schema   `json:"schemaId"`
	providerId Provider `json:"providerId"`
}
