package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateJWT(tokenString, jwtSecret string) (*jwt.Token, error) {
	logInfo("Running: Auth - ValidateJWT")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
}

func CreateJWT(user *User, jwtSecret string) (string, error) {
	logInfo("Running: Auth - CreateJWT")
	accessTokenExpireTime := time.Now().Add(time.Hour * 48).Unix()
	claims := &jwt.MapClaims{
		"expiresAt": accessTokenExpireTime,
		"userId":    user.UserID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func GetUserIDFromJWT(token *jwt.Token) (uuid.UUID, error) {
	logInfo("Running: Auth - GetUserIDFromJWT")
	claims := token.Claims.(jwt.MapClaims)
	return uuid.Parse(claims["userId"].(string))
}

func CheckJWTExpired(token *jwt.Token) bool {
	logInfo("Running: Auth - CheckJWTExpired")
	claims := token.Claims.(jwt.MapClaims)
	tokenExpiry := claims["expiresAt"].(float64)
	return tokenExpiry < float64(time.Now().Unix())
}
