package util_password

import util_error "backend-template/util/error"

type PasswordManager interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) *util_error.ClientError
	PasswordValidation(password string) *util_error.ClientError
}
