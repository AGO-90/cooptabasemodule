package components

type Driver struct {
	ID                    int64  `gorm:"primaryKey" json:"id"`
	CooperativeIdentifier string `json:"cooperative_identifier"`
	Person                Person `json:"person"`
}

func (Driver) TableName() string {
	return "drivers"
}
