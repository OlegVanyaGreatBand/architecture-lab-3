package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJsonBadRequest(rw http.ResponseWriter, message string)  {
	writeJson(rw, http.StatusBadRequest, struct {
		error string
	}{
		error: message,
	})
}

func WriteJsonInternalError(rw http.ResponseWriter, message string)  {
	writeJson(rw, http.StatusInternalServerError, struct {
		error string
	}{
		error: message,
	})
}

func WriteJsonResult(rw http.ResponseWriter, res interface{}) {
	writeJson(rw, http.StatusOK, res)
}

func writeJson(rw http.ResponseWriter, status int, res interface{}) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		log.Printf("Error writing response: %s", err)
	}
}
