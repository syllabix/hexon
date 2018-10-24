package hexon

/* This file contains the definition for
* the expected json payload to hexon for
* updating and creating vehicles */

type attributename struct {
	Name string `json:"name"`
}

type identification struct {
	Vin          string `json:"vin"`
	LicensePlate string `json:"license_plate"`
	Location     string `json:"location"`
}

type carinfo struct {
	Make      attributename `json:"make"`
	Model     attributename `json:"model"`
	BodyStyle string        `json:"bodystyle"`
	Category  string        `json:"category"`
}

type pricedetails struct {
	Value       float64 `json:"value"`
	IncludesVAT bool    `json:"incl_vat"`
}

type price struct {
	Currency string       `json:"currency"`
	Consumer pricedetails `json:"consumer"`
}

type salescondition struct {
	Pricing price `json:"pricing"`
}

type odometer struct {
	Reading int    `json:"reading"`
	Unit    string `json:"unit"`
}

type condition struct {
	Used     bool     `json:"used"`
	Odometer odometer `json:"odometer"`
}

type history struct {
	ArrivalDate string `json:"arrival_date"`
}

type color struct {
	Primary string `json:"primary"`
}

type body struct {
	Color color `json:"colour"`
}

type hexonvehicle struct {
	StockNumber    string         `json:"stocknumber"`
	Identification identification `json:"identification"`
	GeneralInfo    carinfo        `json:"general"`
	SalesCondition salescondition `json:"sales_conditions"`
	Condition      condition      `json:"condition"`
	History        history        `json:"history"`
	Body           body           `json:"body"`
}

func payloadify(vehicle Vehicle) hexonvehicle {
	return hexonvehicle{
		StockNumber: vehicle.Vin,
		Identification: identification{
			Vin:          vehicle.Vin,
			LicensePlate: vehicle.LicenseNumber,
			Location:     vehicle.LocationCode,
		},
		GeneralInfo: carinfo{
			Make:      attributename{Name: vehicle.Make},
			Model:     attributename{Name: vehicle.Model},
			BodyStyle: vehicle.BodyStyle,
			Category:  vehicle.HexonCategory,
		},
		SalesCondition: salescondition{
			Pricing: price{
				Currency: vehicle.Currency,
				Consumer: pricedetails{
					Value:       vehicle.PriceIncludingVat,
					IncludesVAT: true,
				},
			},
		},
		Condition: condition{
			Used: !vehicle.IsNewCar,
			Odometer: odometer{
				Reading: vehicle.ExpectedMileage,
				Unit:    vehicle.MileageUnit,
			},
		},
		Body: body{
			Color: color{
				Primary: vehicle.ExteriorColor,
			},
		},
		History: history{
			ArrivalDate: vehicle.ExpectedDateAvailable.Format("2006-01-02"),
		},
	}
}
