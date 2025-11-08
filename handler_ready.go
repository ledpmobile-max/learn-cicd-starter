package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, _ *http.Request) { // <--- CAMBIO AQUÍ
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
