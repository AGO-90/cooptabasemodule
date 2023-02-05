package components

import (
	"testing"
	"time"
)

func TestPerson_GetAge(t *testing.T) {
	// Fija la fecha actual en un momento específico
	now := time.Date(2023, time.February, 5, 0, 0, 0, 0, time.UTC)

	// Caso de prueba: Edad esperada = 25 años
	birthday := time.Date(1997, time.May, 14, 0, 0, 0, 0, time.UTC)
	person := Person{BirthDate: birthday}
	expectedAge := 25
	actualAge := person.GetAge(now)
	if actualAge != expectedAge {
		t.Errorf("Error: esperado %d, obtenido %d", expectedAge, actualAge)
	}
}
