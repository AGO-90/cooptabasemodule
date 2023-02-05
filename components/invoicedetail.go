package components

type InvoiceDetail struct {
	ID           int     `json:"id"`
	Observations string  `json:"observations"`
	Amount       float64 `json:"amount"`
	ImageURL     string  `json:"image_url"`
	Type         string  `json:"type"`
}
