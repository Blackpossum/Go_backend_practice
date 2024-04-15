package main

import "net/http"

// this is the http handler respoonse if succesfull

func Handlerr_management(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
