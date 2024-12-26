package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func respondJSON(w http.ResponseWriter, statusCode int, success bool, message string, body any) {
	res, err := json.Marshal(APIResponse{
		Success: success,
		Data:    body,
		Message: message,
	})
	if err != nil {
		log.Fatalf("Failed to marshal JSON response: %v", res)
		w.WriteHeader(500)
		//return fmt.Errorf("Failed to marshal JSON response: %v", res)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	//return json.NewEncoder(w).Encode(data)
	w.Write(res)
}

func respondERROR(w http.ResponseWriter, statusCode int, message string) {
	respondJSON(w, statusCode, false, message, nil)
}

func respondOK(w http.ResponseWriter, statusCode int, message string, body any) {
	respondJSON(w, statusCode, true, message, body)
}
