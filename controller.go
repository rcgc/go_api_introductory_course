package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (h *carHandler) getAll(w http.ResponseWriter, r *http.Request){
	defer h.Unlock()
	h.Lock()
	car := Car{}
	q, err := car.getAllCars()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	respondWithJSON(w, http.StatusOK, q)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Add("content-type", "application/json")
	w.Header().Add("authorization", "Access-Control-Allow-Origin")
	w.WriteHeader(code)
	w.Write(response)
}

func idFromUrl(r *http.Request) (string) {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) < 3 {
		return "-1"
	}
	if parts[2] == ""{
		return "-1"
	}
	id :=  parts[2]
	return id
}