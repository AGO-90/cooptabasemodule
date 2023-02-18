package components

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	ID                    int64  `gorm:"primaryKey" json:"id"`
	CooperativeIdentifier string `json:"cooperative_identifier"`
	Person                Person `gorm:"foreignKey:people_document_number" json:"person"`
}

func (Driver) TableName() string {
	return "drivers"
}
