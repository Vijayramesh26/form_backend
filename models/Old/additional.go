package models

type Additional struct {
	ID                   uint   `json:"id"`
	ContractNotePhysical int    `json:"contract_note_physical"`
	ContractNoteEmail    string `json:"contract_note_email"`
}
