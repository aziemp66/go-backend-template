package util_jwt

import jwt "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	Identifier string
	Name       string
	Role       string
	jwt.RegisteredClaims
}
