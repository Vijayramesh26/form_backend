package models

type KycType struct {
	ID        uint   `json:"id"`
	TypeKey   string `json:"type_key"`
	Label     string `json:"label"`
	FeeAmount int    `json:"fee_amount"`
}
