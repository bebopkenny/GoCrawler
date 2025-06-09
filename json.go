package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) { // Format the object as a JSON format everytime
	if code > 499 { // Client side errors
		log.Println("Responding with 5XX error: ", msg)
	}
	type errResponse struct {
		Error string `json:"error"` // Key will be error
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	}) 
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json") // 
	w.WriteHeader(code)
	w.Write(dat)
}