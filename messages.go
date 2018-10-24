package hexon

type pubmessage struct {
	StockNumber string `json:"stocknumber"`
	SiteCode    string `json:"site_code"`
}

func makePublishMessage(vin, sitecode string) (pubmessage, error) {
	validationErrs := make([]error, 0, 2)
	if !isValidVin(vin) {
		validationErrs = append(validationErrs, ErrorInvalidVin)
	}

	if sitecode == "" {
		validationErrs = append(validationErrs, ErrorEmptySiteCode)
	}

	if len(validationErrs) > 0 {
		return pubmessage{}, concatErrors("Unable to create valid publish message", validationErrs...)
	}

	return pubmessage{
		StockNumber: vin,
		SiteCode:    sitecode,
	}, nil
}
