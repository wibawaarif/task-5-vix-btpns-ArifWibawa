package app

import "github.com/golang-jwt/jwt"

var JWT_KEY = []byte("90a514ab27e2c32fdd1018154b26a199") // JWT Key for token

type JWTuser struct {
	Email string
	jwt.StandardClaims
}