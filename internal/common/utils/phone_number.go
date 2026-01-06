package utils

import "strings"

func FormatDialCode(dialCode string) string {
	if strings.HasPrefix(dialCode, "+") {
		return dialCode[1:]
	}
	return dialCode
}

func FormatPhoneNumber(phoneNumber string) string {
	if strings.HasPrefix(phoneNumber, "0") {
		return phoneNumber[1:]
	}
	return phoneNumber
}
