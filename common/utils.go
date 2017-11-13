package common

import (
	"net/http"
	"log"
	"encoding/json"
	"os"
)

type configuration struct {
	Server, SqlDatabaseHost, DatabaseUser, DatabasePassword, Database string
}

/*
	AppConfig holds the configuration values from
	config.json file var AppConfig configuration
*/

var AppConfig configuration

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s \n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

func DisplayAppError(w http.ResponseWriter, handleError error, message string, code int) {
	errObj := appError{
		Error:      handleError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handleError)
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}
