package models

type Promoter struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	PAN    string `json:"pan"`
	DIN    string `json:"din"`
	Aadhar string `json:"aadhar"`
}
