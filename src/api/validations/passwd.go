package validation

import (
	"regexp"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func PasswdStrengthValidator(fld validator.FieldLevel) bool{
	password := fld.Field().String()
	return IsStrongPassword(password)
}

func IsStrongPassword(password string) bool {
	var minLen, upper, lower, number, special bool

	if len(password) >= 8 {
		minLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upper = true
		case unicode.IsLower(char):
			lower = true
		case unicode.IsDigit(char):
			number = true
		case isSpecialChar(char):
			special = true
		}
	}

	return minLen && upper && lower && number && special
}

func isSpecialChar(c rune) bool {
	match, _ := regexp.MatchString(`[!@#$%^&*()_+-={};':"|,.<>\/?]`, string(c))
	return match
}