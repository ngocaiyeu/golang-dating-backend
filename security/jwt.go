package security

import (
	"github.com/golang-jwt/jwt"
	"lienquanMess/model"
	"time"
)

const JWT_KEY = "ucneifbyuewcijwnfu"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
