package util_password

import (
	"errors"
	"regexp"

	util_error "backend-template/util/error"

	"golang.org/x/crypto/bcrypt"
)

type passwordManager struct {
	cost int
}

func NewPasswordManager(cost int) PasswordManager {
	return &passwordManager{
		cost: cost,
	}
}

func (p *passwordManager) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	return string(bytes), err
}

func (p *passwordManager) CheckPasswordHash(password, hash string) *util_error.ClientError {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return util_error.NewUnauthorized(err, "Wrong email or password")
	}
	return nil
}

func (p *passwordManager) PasswordValidation(password string) *util_error.ClientError {
	// Minimum 8 characters, at least one uppercase letter, one lowercase letter,
	// one digit, and one special character
	lowerCase := regexp.MustCompile(`[a-z]`).MatchString(password)
	if !lowerCase {
		return util_error.NewBadRequest(errors.New("password must have lowercase characters"), "Password must have lowercase characters")
	}

	upperCase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !upperCase {
		return util_error.NewBadRequest(errors.New("password must have UPPERCASE characters"), "Password must have UPPERCASE characters")
	}

	digit := regexp.MustCompile(`\d`).MatchString(password)
	if !digit {
		return util_error.NewBadRequest(errors.New("password must have Numerical characters"), "Password must have Numerical characters")
	}

	// specialChar := regexp.MustCompile(`[!@#$%^&*()[]\\;',./{}|:"<>?]`).MatchString(password)

	length := len(password) >= 8 && len(password) <= 72
	if !length {
		return util_error.NewBadRequest(errors.New("password must be between 8 to 72 characters long"), "Password must be between 8 to 72 characters long")
	}

	return nil
}
