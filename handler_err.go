package main

import "net/http"

// Handle error message from json.go function
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
