package main

import (
	"encoding/json"
	"net/http"
)

// CREATE NEW THREAD
func (apiCfg *APIConfig) handlerCreateThread(w http.ResponseWriter, r *http.Request) {
	logInfo("Running handlerCreateThread")
	//decoder := json.NewDecoder(r.Body)
	req := CreateThreadRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	// user, err := NewStandardUser(req.Username, req.Password)
	// if err != nil {
	// 	logError("Error Creating New Standard User Template", err)
	// }
	// fmt.Printf("New User: %v\n", user) //remove later

	// err = apiCfg.DB.CreateUser(user)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "duplicate key") {
	// 		logError("User Already Exists", err)
	// 		respondERROR(w, http.StatusBadRequest, "User Already Exists")
	// 		return
	// 	}
	// 	logError("Unable to Create User", err)
	// 	respondERROR(w, http.StatusBadRequest, "Failed to Register User")
	// 	return
	// }

	respondOK(w, http.StatusCreated, "Thread Created Successfully", nil)
}
