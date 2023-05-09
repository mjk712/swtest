package models

type Airline struct {
	Id   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type AirlId struct {
	Id uint `json:"id"`
}
