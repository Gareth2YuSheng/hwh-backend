package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// REGISTER NEW USER
func (apiCfg *APIConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	logInfo("Running: Handler - CreateUser")
	req := CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	user, err := NewStandardUser(req.Username, req.Password)
	if err != nil {
		logError("Error Creating New Standard User Template", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Register User, Invalid User Details")
		return
	}

	err = apiCfg.DB.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			logError("User Already Exists", err)
			respondERROR(w, http.StatusBadRequest, "User Already Exists")
			return
		}
		logError("Unable to Create User", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Register User")
		return
	}

	respondOK(w, http.StatusCreated, "User Created Successfully", nil)
}

// LOGIN USER
func (apiCfg *APIConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	logInfo("Running: Handler - Login")
	req := LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	//Check if user exists
	user, err := apiCfg.DB.GetUserByUsername(req.Username)
	if err != nil {
		logError(fmt.Sprintf("User [%s] Not Found", req.Username), err)
		respondERROR(w, http.StatusBadRequest, "Login Failed")
		return
	}

	//Validate Password if user has password
	if user.Password != "" {
		if !ComparePassword(user.Password, req.Password) {
			logError(fmt.Sprintf("Incorrect Password for User [%s]", req.Username), nil)
			respondERROR(w, http.StatusBadRequest, "Login Failed: Incorrect Username or Password")
			return
		}
	}

	//Generate JWT
	token, err := CreateJWT(user, apiCfg.JWTSecret)
	if err != nil {
		logError("Error Creating JWT", err)
		respondERROR(w, http.StatusInternalServerError, "Login Failed")
		return
	}

	respondOK(w, http.StatusOK, "Login Successful", LoginResponse{
		UserID:      user.UserID,
		AccessToken: token,
	})
}

// GET USER DATA
func (apiCfg *APIConfig) handlerGetUserData(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running: Handler - GetUserData")

	respondOK(w, http.StatusOK, "", GetUserDataResponse{
		User: user,
	})
}
