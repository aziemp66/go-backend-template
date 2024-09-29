package util_password

import (
	"regexp"

	util_error "github.com/Final-Project-Azie/e-commerce-be/util/error"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return util_error.NewUnauthorized(err, "Wrong email or password")
	}
	return nil
}

func PasswordValidation(password string) bool {
	// Minimum 8 characters, at least one uppercase letter, one lowercase letter,
	// one digit, and one special character
	lowerCase := regexp.MustCompile(`[a-z]`).MatchString(password)
	upperCase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	digit := regexp.MustCompile(`\d`).MatchString(password)
	// specialChar := regexp.MustCompile(`[!@#$%^&*()[]\\;',./{}|:"<>?]`).MatchString(password)

	length := len(password) >= 8
	return lowerCase && upperCase && digit && length
}
