package helper

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func WriteJson(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)

}

func ReadJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decodedVal := json.NewDecoder(r.Body)
	decodedVal.DisallowUnknownFields()
	return decodedVal.Decode(data)
}

func WriteJsonError(w http.ResponseWriter, statusCode int, errMsg string) {
	type envelope struct {
		Error string `json:"error"`
	}
	WriteJson(w, statusCode, &envelope{Error: errMsg})
}

func JsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}

	return WriteJson(w, status, &envelope{Data: data})
}
