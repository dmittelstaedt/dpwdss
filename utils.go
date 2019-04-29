package main

import "net/http"

// respondJSON repsonds with JSON in body and sets content and http status to header.
func respondJSON(w http.ResponseWriter, httpCode int, response []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpCode)
	w.Write(response)
}
