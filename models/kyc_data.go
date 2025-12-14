package models

type KycData struct {
	ID        uint   `json:"id"`
	TypeKey   string `json:"type_key"`
	JsonData  string `json:"json_data"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
