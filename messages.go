package hexon

import "errors"

var (
	// ErrorInvalidVin is returned when a string is provided with the intention of being used as a VIN
	ErrorInvalidVin = errors.New("The provided vin is invalid. It must be 17 a character alphanumeric string")

	// ErrorEmptySiteCode when a sitecode is provided as an empty string
	ErrorEmptySiteCode = errors.New("A site code is required to publish the vehicle")
)

type pubmessage struct {
	StockNumber string `json:"stocknumber"`
	SiteCode    string `json:"site_code"`
}

func makePublishMessage(vin, sitecode string) (pubmessage, error) {
	if !isValidVin(vin) {
		return pubmessage{}, ErrorInvalidVin
	}

	if sitecode == "" {
		return pubmessage{}, ErrorEmptySiteCode
	}

	return pubmessage{
		StockNumber: vin,
		SiteCode:    sitecode,
	}, nil
}
