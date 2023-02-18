package components

type Driver struct {
	ID                    int64  `gorm:"primaryKey" json:"id"`
	CooperativeIdentifier string `json:"cooperative_identifier"`
	Person                Person `gorm:"foreignKey:DocumentNumber" json:"person"`
	PersonDocumentNumber  string `json:"person_document_number"`
}

func (Driver) TableName() string {
	return "drivers"
}
