package main 

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) { // Respond if the server is alive and running
	respondWithJSON(w, 200, struct{}{})
}