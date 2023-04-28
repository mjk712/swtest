package models

type Provider struct {
	Id         uint   `json:"id"`
	ProviderId string `json:"providerid"`
	Name       string `json:"name"`
}

type ProvId struct {
	Id uint `json:"id"`
}
