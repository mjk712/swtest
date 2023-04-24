package models

type AirlineProvider struct {
	id         uint     `json:"id"`
	airlineId  Airline  `json:"airlineId"`
	providerId Provider `json:"providerId"`
}
