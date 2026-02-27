package validator

import (
	"errors"
	"regexp"
)

var pnrRegex = regexp.MustCompile(`^[A-Z0-9]{6}$`)

// ValidatePNR checks if the transformed string matches Airline PNR standards
func ValidatePNR(pnr string) error {
	if !pnrRegex.MatchString(pnr) {
		return errors.New("invalid PNR format: must be 6 alphanumeric characters")
	}
	return nil
}

// ValidatePayload ensures crucial fields aren't missing post-translation
func ValidatePayload(data map[string]interface{}) error {
	required := []string{"pnr", "passenger_name", "flight_id"}
	for _, field := range required {
		if _, ok := data[field]; !ok {
			return errors.New("missing required field: " + field)
		}
	}
	return nil
}