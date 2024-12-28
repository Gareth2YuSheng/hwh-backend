package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString, jwtSecret string) (*jwt.Token, error) {
	logInfo("Running ValidateJWT")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jwtSecret), nil
	})
}

func CreateJWT(user *User, jwtSecret string) (string, error) {
	logInfo("Running CreateJWT")
	accessTokenExpireTime := time.Now().Add(time.Hour * 48).Unix()

	claims := &jwt.MapClaims{
		"expiresAt": accessTokenExpireTime,
		"userId":    user.UserID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
