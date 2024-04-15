package check

import (
	"errors"
	"strings"
	"time"
	"unicode"
)

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}

// func Validategmail(gmail string) error {
// 	// Check if the email contains "@gmail.com"
// 	containsGmailCom := strings.Contains(gmail, "@gmail.com")
// 	// Check if the email contains "@gmail.ru"
// 	containsGmailRu := strings.Contains(gmail, "@gmail.ru")

// 	// If the email contains both domains, return an error
// 	if containsGmailCom && containsGmailRu {
// 		return errors.New("email address cannot contain both @gmail.com and @gmail.ru domains")
// 	}

// 	// If the email is valid (doesn't contain both domains), return nil
// 	return nil
// }

func Validategmail(gmail string) error {
	// Check if the email ends with either "@gmail.com" or "@gmail.ru"
	if !strings.HasSuffix(gmail, "@gmail.com") && !strings.HasSuffix(gmail, "@gmail.ru") {
		return errors.New("email address must end with @gmail.com or @gmail.ru")
	}

	return nil
}

// func Validatenumber(number string) error {
// 	containsNumber := strings.Contains(number, "998")

// 	if containsNumber {
// 		return errors.New("number  first three numbers should be: 998for example: 998 700 00 00 ")
// 	}

// 	return nil
// }

func ValidatePassword(newPassword string) error {
	if len(newPassword) < 8 {
		return errors.New("=====password length must be at least 8 characters")
	}

	var hasUppercase, hasLowercase, hasDigit, hasSymbol bool

	for _, char := range newPassword {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsNumber(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSymbol = true
		}
	}

	if !hasUppercase {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLowercase {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasSymbol {
		return errors.New("password must contain at least one symbol")
	}

	return nil
}
