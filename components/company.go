package components

import "time"

type Company struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	RUT              string    `json:"rut"`
	RegistrationDate time.Time `json:"registration_date"`
	Active           bool      `json:"active"`
}
