package main

import (
	"log"
	"net/http"
)

// respondJSON repsonds with JSON in body and sets content and http status to header.
func respondJSON(w http.ResponseWriter, httpCode int, response []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpCode)
	w.Write(response)
}

// respondHeader responds httpCode and empty body.
func respondHeader(w http.ResponseWriter, httpCode int) {
	w.WriteHeader(httpCode)
}

// logRequest logs a request.
func logRequest(r *http.Request) {
	if r.URL.RawQuery == "" {
		log.Println("Request: " + "[" + r.Method + "] " + r.URL.Path)
	} else {
		log.Println("Request: " + "[" + r.Method + "] " + r.URL.Path + "?" + r.URL.RawQuery)
	}
}
