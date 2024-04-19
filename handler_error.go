package main

import "net/http"

// this is the http handler respoonse if succesfull

func Handlerr_Error(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusBadRequest, "Invalid request, something went wrong or the server is down",400)
}
