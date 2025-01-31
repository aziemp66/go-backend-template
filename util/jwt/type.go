package util_jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	ID   string
	Name string
	Role ROLE
	jwt.RegisteredClaims
}

type JWTManager interface {
	GenerateAuthToken(
		ID string,
		name string,
		role ROLE,
		duration time.Duration,
	) (string, error)
	VerifyAuthToken(tokenString string) (claims *AuthClaims, err error)
}
