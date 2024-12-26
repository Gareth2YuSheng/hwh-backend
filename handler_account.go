package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// REGISTER NEW USER
func (apiCfg *APIConfig) handlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Running handlerCreateAccount")
	decoder := json.NewDecoder(r.Body)
	req := CreateAccountRequest{}
	err := decoder.Decode(&req)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
		respondERROR(w, 400, "Something went wrong")
		return
	}

	//user, err := apiCfg.DB

	respondOK(w, 201, "Account Created Successfully", req)
}
