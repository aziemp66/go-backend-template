package util_jwt

import (
	"errors"
	"time"

	util_error "github.com/Final-Project-Azie/e-commerce-be/util/error"
	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	AccessTokenKey []byte
}

func NewJWTManager(accessTokenKey string) *JWTManager {
	return &JWTManager{AccessTokenKey: []byte(accessTokenKey)}
}

func (j JWTManager) GenerateAuthToken(
	identifier string,
	name string,
	role string,
	duration time.Duration,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		Name:       name,
		Identifier: identifier,
		Role:       role,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365 * 10)),
		},
	})

	tokenString, err := token.SignedString(j.AccessTokenKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JWTManager) VerifyAuthToken(tokenString string) (claim *AuthClaims, err error) {
	claim = &AuthClaims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return j.AccessTokenKey, nil
	})
	if err != nil {
		return claim, util_error.NewBadRequest(err, "Invalid token")
	}

	if !tkn.Valid {
		return claim, util_error.NewForbidden(
			errors.New("token does not match secret key"),
			"You are not authorized for this access",
		)
	}

	return claim, nil
}
