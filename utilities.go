package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// LOGGERS
func logInfo(message string) {
	log.Printf("INFO - %s\n", message)
}

func logError(message string, err error) {
	log.Printf("ERROR - %s: %v\n", message, err)
}

func logFatal(message string, err error) {
	log.Fatalf("FATAL - %s: %v\n", message, err)
}

// HANDLER UTILITIES
func PermissionDeniedRes(w http.ResponseWriter) {
	respondERROR(w, http.StatusForbidden, "Permission Denied")
}

func SomethingWentWrongRes(w http.ResponseWriter) {
	respondERROR(w, http.StatusBadRequest, "Something went wrong")
}

func ErrorParsingJSON(err error, w http.ResponseWriter) {
	logError("Error parsing JSON", err)
	SomethingWentWrongRes(w)
}

// BCRYPT FUNCTIONS
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComparePassword(hashedPwd, pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd)) == nil
}
