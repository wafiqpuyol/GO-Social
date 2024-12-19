package main

import (
	"net/http"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "okyish",
	}

	if err := writeJson(w, http.StatusOK, &status); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
	}
}
