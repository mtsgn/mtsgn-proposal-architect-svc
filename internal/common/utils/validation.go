// pkg/utils/validation.go
package utils

import "regexp"

// IsValidPhoneNumber validates the phone number format
func IsValidPhoneNumber(phoneNumber string) bool {
	// Simple phone number validation (using regex)
	// The regex pattern here allows digits and expects at least 10 digits
	re := regexp.MustCompile(`^\d{10}$`)
	return re.MatchString(phoneNumber)
}
