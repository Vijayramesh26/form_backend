package models

type Payment struct {
	ID              uint   `json:"id"`
	KycDataID       uint   `json:"kyc_data_id"`
	Provider        string `json:"provider"`
	ProviderOrderID string `json:"provider_order_id"`
	Amount          int    `json:"amount"`
	Currency        string `json:"currency"`
	Status          string `json:"status"`
}
