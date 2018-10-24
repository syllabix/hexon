package hexon

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

//Common Errors
var (
	ErrorInvalidVin    = errors.New("The provided vin is invalid. It must be 17 a character alphanumeric string")
	ErrorEmptySiteCode = errors.New("A site code is required to publish the vehicle")
	ErrorEmptyUserName = errors.New("Username is required")
	ErrorEmptyPassword = errors.New("Password is required")
	ErrorHexonAPI      = errors.New("Hexon responded with an http status code signaling an error occurred")
)

var (
	vincheck = regexp.MustCompile("[A-Z0-9]{17}")
)

func isValidVin(vin string) bool {
	return vincheck.MatchString(vin)
}

func concatErrors(msg string, errs ...error) error {
	buidler := strings.Builder{}
	buidler.WriteString(fmt.Sprintf("%s:\n", msg))
	for i, err := range errs {
		buidler.WriteString(fmt.Sprintf("%d. %s\n", i, err))
	}
	return errors.New(buidler.String())
}
