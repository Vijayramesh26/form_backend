package models

type KycTypeFieldRule struct {
	ID       uint   `json:"id"`
	TypeKey  string `json:"type_key"`
	FieldKey string `json:"field_key"`
	Required int    `json:"required"`
}
