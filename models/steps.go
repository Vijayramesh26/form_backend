package models

type KycStep struct {
	ID        uint   `json:"id"`
	StepKey   string `json:"step_key"`
	StepLabel string `json:"step_label"`
	Sequence  int    `json:"sequence"`
	IsActive  int    `json:"is_active"`
}
