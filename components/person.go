package components

import "time"

type EmploymentStatus string

var (
	Retired          EmploymentStatus = "retired"
	PartiallyRetired EmploymentStatus = "partially_retired"
	Employed         EmploymentStatus = "employed"
	Undefined        EmploymentStatus = "undefined"
)

type Person struct {
	CI               string           `gorm:"primaryKey" json:"ci"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	BirthDate        time.Time        `json:"birth_date"`
	DocumentType     string           `json:"document_type"`
	DocumentNumber   string           `json:"document_number"`
	PreferredName    string           `json:"preferred_name"`
	EmploymentStatus EmploymentStatus `json:"employment_status"`
	Phone            string           `json:"phone"`
	Email            string           `json:"email"`
	Address          string           `json:"address"`
}

func (Person) TableName() string {
	return "people"
}

func (p *Person) GetAge(now time.Time) int {
	age := now.Year() - p.BirthDate.Year()
	if now.YearDay() < p.BirthDate.YearDay() {
		age--
	}
	return age
}
