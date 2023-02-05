package components

import "time"

type InternalAbsence struct {
	ID           int       `json:"id"`
	Date         time.Time `json:"date"`
	Observations string    `json:"observations"`
}
