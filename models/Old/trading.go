package models

type Trading struct {
	ID    uint `json:"id"`
	NSE   int  `json:"nse"`
	BSE   int  `json:"bse"`
	MCX   int  `json:"mcx"`
	NCDEx int  `json:"ncdex"`
}
