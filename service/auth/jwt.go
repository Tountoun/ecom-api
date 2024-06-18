package auth

import (
	"strconv"
	"time"

	"github.com/Tountoun/ecom-api/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	expirationTime := time.Second * time.Duration(config.Envs.JWTExpirationTimeInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expirationTime).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}