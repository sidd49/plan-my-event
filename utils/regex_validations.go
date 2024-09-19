package utils

import (
	"net/mail"
	"regexp"
	"unicode"
)

func CheckValidEmail(email string) bool {
	const emailPattern string = "^[a-zA-Z0-9_!#$%&'*+/=?`{|}~^.-]+@[a-zA-Z0-9-]+.[a-zA-Z]+$"
	re := regexp.MustCompile(emailPattern)
	_, err := mail.ParseAddress(email)
	if re.MatchString(email) && err == nil {
		return true
	}
	return false
}

func CheckValidPassword(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 7 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func CheckValidLocation(location string) bool {
	const locationPattern string = "^[A-Za-z0-9,-]+$"
	re := regexp.MustCompile(locationPattern)
	return re.MatchString(location)
}
