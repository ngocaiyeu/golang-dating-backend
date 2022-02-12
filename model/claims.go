package model

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	UserId string
	jwt.StandardClaims
}
