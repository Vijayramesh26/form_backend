package models

type KycField struct {
	ID           uint   `json:"id"`
	StepKey      string `json:"step_key"`
	FieldKey     string `json:"field_key"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	Items        string `json:"items"`
	BaseRequired int    `json:"base_required"`
	Sequence     int    `json:"sequence"`
}
