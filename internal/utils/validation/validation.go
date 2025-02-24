package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	reLower := regexp.MustCompile(`[a-z]`)
	reUpper := regexp.MustCompile(`[A-Z]`)
	reDigit := regexp.MustCompile(`\d`)
	reSpecial := regexp.MustCompile(`[^a-zA-Z\d]`)
	reLength := regexp.MustCompile(`.{8,}`)

	return reLower.MatchString(password) &&
		reUpper.MatchString(password) &&
		reDigit.MatchString(password) &&
		reSpecial.MatchString(password) &&
		reLength.MatchString(password)
}
