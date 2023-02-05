package components

type Repair struct {
	ID           int     `json:"id"`
	DateOfEntry  string  `json:"date_of_entry"`
	DateOfExit   string  `json:"date_of_exit"`
	Cost         float64 `json:"cost"`
	Reason       string  `json:"reason"`
	Observations string  `json:"observations"`
}
