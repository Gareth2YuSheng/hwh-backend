package main

import (
	"fmt"
	"net/http"
	"strings"
)

type authenticatedHandler func(http.ResponseWriter, *http.Request, User)

func (apiCfg *APIConfig) middlewareAuth(handler authenticatedHandler) http.HandlerFunc {
	logInfo("Running middlewareAuth")
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			PermissionDeniedRes(w)
			return
		}
		splitToken := strings.Split(auth, "Bearer ")
		//Check if "Bearer" in auth header
		if len(splitToken) < 2 {
			PermissionDeniedRes(w)
			return
		}

		tokenString := splitToken[1]
		fmt.Printf("Token String: %s", tokenString)
		// token, err := validateJWT(tokenString)
		// if err != nil {
		// 	//WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
		// 	permissionDenied(w)
		// 	return
		// }
		// if !token.Valid {
		// 	//WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
		// 	permissionDenied(w)
		// 	return
		// }

		//Change this Part
		// userID, err := getID(r)
		// if err != nil {
		// 	//WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
		// 	permissionDenied(w)
		// 	return
		// }
		// account, err := s.GetAccountByID(userID)
		// if err != nil {
		// 	//WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
		// 	permissionDenied(w)
		// 	return
		// }
		// claims := token.Claims.(jwt.MapClaims)
		// if account.Number != int64(claims["accountNumber"].(float64)) {
		// 	permissionDenied(w)
		// 	return
		// }

		//easiest way is to sign the userID into the JWT and check if the provided data matches
		//with the userID in the JWT
		//Change all error messages to Permission denied

		// user :=

		// handler(w, r, user)
	}
}
