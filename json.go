package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// helper function to respond with JSON

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSONresponse: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/JSON")
	w.WriteHeader(200)
	w.Write(dat)
}
