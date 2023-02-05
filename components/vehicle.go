package components

type Vehicle struct {
	LicensePlate string  `json:"license_plate"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Year         string  `json:"year"`
	PurchaseYear int8    `json:"purchase_year"`
	Kilometers   float32 `json:"kilometers"`
}
