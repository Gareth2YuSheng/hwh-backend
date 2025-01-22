package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// LOGGERS
func logInfo(message string) {
	log.Printf("INFO - %s\n", message)
}

func logError(message string, err error) {
	log.Printf("ERROR - %s: %v\n", message, err)
	//If Error is that DB has too many clients, cut the server
	if strings.Contains(err.Error(), "too many clients") || strings.Contains(err.Error(), "remaining connection slots are reserved") {
		os.Exit(1)
		//Combine with script on EC2 instance to restart the server on exit
	}
}

func logFatal(message string, err error) {
	log.Fatalf("FATAL - %s: %v\n", message, err)
}

// HANDLER UTILITIES
func PermissionDeniedRes(w http.ResponseWriter) {
	respondERROR(w, http.StatusForbidden, "Permission Denied")
}

func ErrorParsingJSON(err error, w http.ResponseWriter) {
	logError("Error parsing JSON", err)
	respondERROR(w, http.StatusBadRequest, "Invalid Request Body")
}

func InvalidURLQuery(message string, err error, w http.ResponseWriter) {
	logError(message, err)
	respondERROR(w, http.StatusBadRequest, "Invalid URL Query")
}

// BCRYPT FUNCTIONS
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComparePassword(hashedPwd, pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd)) == nil
}

// TIME FUNCTIONS
func getTimeNow() time.Time {
	return time.Now().Local().UTC()
}
