package components

import "time"

type EmploymentStatus string

var (
	Retired          EmploymentStatus = "retired"
	PartiallyRetired EmploymentStatus = "partially_retired"
	Employed         EmploymentStatus = "employed"
)

type Person struct {
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	BirthDate        time.Time        `json:"birth_date"`
	DocumentType     string           `json:"document_type"`
	DocumentNumber   string           `json:"document_number"`
	PreferredName    string           `json:"preferred_name"`
	EmploymentStatus EmploymentStatus `json:"employment_status"`
	Phones           []string         `json:"phones"`
	Email            string           `json:"email"`
	Address          string           `json:"address"`
}

func (p *Person) GetAge(now time.Time) int {
	age := now.Year() - p.BirthDate.Year()
	if now.YearDay() < p.BirthDate.YearDay() {
		age--
	}
	return age
}
