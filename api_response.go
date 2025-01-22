package main

import (
	"encoding/json"
	"fmt"
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
		logError(fmt.Sprintf("Failed to marshal JSON response: %v", res), err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

func respondERROR(w http.ResponseWriter, statusCode int, message string) {
	respondJSON(w, statusCode, false, message, nil)
}

func respondOK(w http.ResponseWriter, statusCode int, message string, body any) {
	respondJSON(w, statusCode, true, message, body)
}
