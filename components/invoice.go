package components

type Invoice struct {
	ID              int     `json:"id"`
	ImageURL        string  `json:"image_url"`
	IssuanceDate    string  `json:"issuance_date"`
	AssociateNumber string  `json:"associate_number"`
	Observations    string  `json:"observations"`
	FinalAmount     float64 `json:"final_amount"`
}
