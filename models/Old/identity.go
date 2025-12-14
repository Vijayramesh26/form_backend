package models

type Identity struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	IncorporationDate  string `json:"incorporation_date"`
	IncorporationPlace string `json:"incorporation_place"`
	RegistrationNo     string `json:"registration_no"`
	Status             string `json:"status"`
}
