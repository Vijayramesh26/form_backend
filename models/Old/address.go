package models

type Address struct {
	ID                    uint   `json:"id"`
	CorrespondenceAddress string `json:"correspondence_address"`
	City                  string `json:"city"`
	Pincode               string `json:"pincode"`
	State                 string `json:"state"`
	Country               string `json:"country"`
}
