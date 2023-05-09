package models

type Schema struct {
	Id        uint     `json:"id"`
	Name      string   `json:"name"`
	Providers []string `json:"providers"`
}

type SchemId struct {
	Id uint `json:"id"`
}
