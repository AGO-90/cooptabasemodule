package components

import "time"

type InternalFault struct {
	ID           int       `json:"id"`
	Date         time.Time `json:"date"`
	Observations string    `json:"observations"`
	Driver       Driver    `json:"driver"`
}

func (InternalFault) TableName() string {
	return "internal_faults"
}
