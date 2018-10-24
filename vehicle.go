package hexon

import "time"

//The Vehicle struct models a car with attributes used in hexon
type Vehicle struct {
	Vin                   string
	LicenseNumber         string
	LocationCode          string
	Make                  string
	Model                 string
	BodyStyle             string
	HexonCategory         string
	Currency              string
	ExpectedDateAvailable time.Time
	PriceIncludingVat     float64
	IsNewCar              bool
	ExpectedMileage       int
	MileageUnit           string
	ExteriorColor         string
	FuelType              string
	IsVatDeductible       bool
}
