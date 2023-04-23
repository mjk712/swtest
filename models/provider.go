package models

type Provider struct {
	id         uint      `json:"id"`
	providerId string    `json:"providerId"`
	name       string    `json:"name"`
	airlines   []Airline `json:"airlines"`
}
