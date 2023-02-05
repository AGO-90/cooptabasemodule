package components

type Fine struct {
	Date              string  `json:"date"`
	Amount            float64 `json:"amount"`
	Observations      string  `json:"observations"`
	URL               string  `json:"url"`
	IntendencyPayment bool    `json:"intendency_payment"`
	PaymentDriver     bool    `json:"payment_driver"`
	IMMPaymentDate    string  `json:"imm_payment_date"`
}
