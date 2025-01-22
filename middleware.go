package main

import (
	"fmt"
	"net/http"
	"strings"
)

type authenticatedHandler func(http.ResponseWriter, *http.Request, User)

func (apiCfg *APIConfig) middlewareAuth(handler authenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInfo("Running: Auth - middlewareAuth")
		auth := r.Header.Get("Authorization")
		if auth == "" {
			logError("Missing Auth Header", nil)
			PermissionDeniedRes(w)
			return
		}
		splitToken := strings.Split(auth, "Bearer ")
		//Check if "Bearer" in auth header
		if len(splitToken) < 2 {
			logError("Incorrect Auth Header", nil)
			PermissionDeniedRes(w)
			return
		}

		tokenString := splitToken[1]
		token, err := ValidateJWT(tokenString, apiCfg.JWTSecret)
		if err != nil {
			logError("Error Validating JWT", err)
			PermissionDeniedRes(w)
			return
		}
		if !token.Valid {
			logError("Invalid JWT", nil)
			PermissionDeniedRes(w)
			return
		}

		//Check whether token has expired
		if CheckJWTExpired(token) {
			logError("JWT IS EXPIRED", nil)
			PermissionDeniedRes(w)
			return
		}

		userId, err := GetUserIDFromJWT(token)
		if err != nil {
			logError("Error Retrieving UserID from JWT", err)
			PermissionDeniedRes(w)
			return
		}
		fmt.Printf("User: [%v]\n", userId) //remove later

		user, err := apiCfg.DB.GetUserByUserID(userId)
		if err != nil {
			logError(fmt.Sprintf("Unable to Get User [%v] using UserID in JWT", userId), err)
			PermissionDeniedRes(w)
			return
		}

		handler(w, r, *user)
	}
}
