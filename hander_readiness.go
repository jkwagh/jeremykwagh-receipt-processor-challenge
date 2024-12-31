package main

import "net/http"

func handlerReadiness(w http.ResposneWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
