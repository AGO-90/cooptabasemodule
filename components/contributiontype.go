package components

type ContributionType struct {
	ID     int     `json:"id"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}
