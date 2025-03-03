package utils

import "regexp"

// Helper function to validate phone number


func isValidPhoneNumber(number string) bool {
	re := regexp.MustCompile(`^[0-9]{10,15}$`)
	return re.MatchString(number)
}