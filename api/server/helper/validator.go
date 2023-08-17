package helper

import (
	"regexp"

	"github.com/asaskevich/govalidator"
)

func IsValidPhoneNumber(phone string) bool {
	if len(phone) != 0 {
		return false
	}

	valid, _ := regexp.MatchString(`[0-9]`, phone)

	return valid
}

func IsValidEmail(email string) bool {
	if len(email) != 0 {
		return false
	}

	valid := govalidator.IsEmail(email)

	return valid
}
