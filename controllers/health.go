package controllers

import (
	"io"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) //Set 200 OK
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
