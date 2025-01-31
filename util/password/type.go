package util_password

type PasswordManager interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) error
	PasswordValidation(password string) error
}
