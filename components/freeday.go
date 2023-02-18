package components

import "time"

type FreeDay struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Date     time.Time `json:"date"`
	Reason   string    `json:"reason"`
	Approver Person    `json:"approver"`
}

func (FreeDay) TableName() string {
	return "free_days"
}
