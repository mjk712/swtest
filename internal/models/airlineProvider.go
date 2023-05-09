package models

type AirlineProvider struct {
	AirlineName string `db:"airline_name"`
	ProviderId  string `db:"provider_id"`
}
