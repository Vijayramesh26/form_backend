package models

type Identity struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	IncorporationDate  string `json:"incorporation_date"`
	IncorporationPlace string `json:"incorporation_place"`
	RegistrationNo     string `json:"registration_no"`
	Status             string `json:"status"`
}

type Address struct {
	ID                    uint   `json:"id"`
	CorrespondenceAddress string `json:"correspondence_address"`
	City                  string `json:"city"`
	Pincode               string `json:"pincode"`
	State                 string `json:"state"`
	Country               string `json:"country"`
}

type Bank struct {
	ID        uint   `json:"id"`
	BankName  string `json:"bank_name"`
	AccountNo string `json:"account_no"`
}

type Promoter struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	PAN    string `json:"pan"`
	DIN    string `json:"din"`
	Aadhar string `json:"aadhar"`
}

type Trading struct {
	ID    uint `json:"id"`
	NSE   int  `json:"nse"`
	BSE   int  `json:"bse"`
	MCX   int  `json:"mcx"`
	NCDEx int  `json:"ncdex"`
}

type Additional struct {
	ID                   uint   `json:"id"`
	ContractNotePhysical int    `json:"contract_note_physical"`
	ContractNoteEmail    string `json:"contract_note_email"`
}
