package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// REGISTER NEW USER
func (apiCfg *APIConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	logInfo("Running handlerCreateUser")
	//decoder := json.NewDecoder(r.Body)
	req := CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	user, err := NewStandardUser(req.Username, req.Password)
	if err != nil {
		logError("Error Creating New Standard User Template", err)
	}
	fmt.Printf("New User: %v\n", user) //remove later

	err = apiCfg.DB.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			logError("User Already Exists", err)
			respondERROR(w, http.StatusBadRequest, "User Already Exists")
			return
		}
		logError("Unable to Create User", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Register User")
		return
	}

	respondOK(w, http.StatusCreated, "User Created Successfully", nil)
}

// LOGIN USER
func (apiCfg *APIConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	logInfo("Running handlerLogin")
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
	fmt.Printf("Found User: %v\n", user) //remove later

	//Validate Password if user has password
	if user.Password != "" {
		fmt.Println("THERE IS A PASSWORD") //remove later
		if !ComparePassword(user.Password, req.Password) {
			logError(fmt.Sprintf("Incorrect Password for User [%s]", req.Username), nil)
			respondERROR(w, http.StatusBadRequest, "Login Failed: Incorrect Username or Password")
			return
		}
	}

	//Generate JWT
	token, err := CreateJWT(user, apiCfg.JwtSecret)
	if err != nil {
		logError("Error Creating JWT", err)
		respondERROR(w, http.StatusInternalServerError, "Login Failed")
		return
	}
	fmt.Println("JWT token ", token) //remove later

	res := LoginResponse{
		UserID:      user.UserID,
		Username:    user.Username,
		Role:        user.Role,
		AccessToken: token,
	}

	respondOK(w, http.StatusOK, "Login Successful", res)
}
