package hexon

import (
	"regexp"
)

var (
	vincheck = regexp.MustCompile("[A-Z0-9]{17}")
)

func isValidVin(vin string) bool {
	return vincheck.MatchString(vin)
}
