package main

import (
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)

}

func readJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decodedVal := json.NewDecoder(r.Body)
	decodedVal.DisallowUnknownFields()
	return decodedVal.Decode(data)
}

func writeJsonError(w http.ResponseWriter, statusCode int, errMsg string) {
	type envelope struct {
		Error string `json:"error"`
	}
	writeJson(w, statusCode, &envelope{Error: errMsg})
}
