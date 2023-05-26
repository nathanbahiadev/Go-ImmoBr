package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(sub string) (string, error) {
	secretKey := GetEnvVariable("JWT_SECRET_KEY")
	duration := GetEnvVariable("JWT_DURATION")

	if secretKey == "" || duration == "" {
		return secretKey, errors.New("getenvvariable: cannot read JWT_SECRET_KEY or JWT_DURATION values")
	}

	tokenExpMinutes, _ := strconv.Atoi(duration)
	tokenExpiration := time.Now().Add(time.Duration(tokenExpMinutes) * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"exp": tokenExpiration,
	})

	return token.SignedString([]byte(secretKey))
}
