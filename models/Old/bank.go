package models

type Bank struct {
	ID        uint   `json:"id"`
	BankName  string `json:"bank_name"`
	AccountNo string `json:"account_no"`
}
