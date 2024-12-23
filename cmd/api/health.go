package main

import (
	"net/http"

	"github.com/wafiqpuyol/GO-Social/internal/helper"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "okyish",
	}

	if err := helper.WriteJson(w, http.StatusOK, &status); err != nil {
		helper.WriteJsonError(w, http.StatusInternalServerError, err.Error())
	}
}
