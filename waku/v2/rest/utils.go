package rest

import (
	"encoding/json"
	"net/http"
)

func writeErrOrResponse(w http.ResponseWriter, err error, value interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func writeResponse(w http.ResponseWriter, value interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResponse, err := json.Marshal(value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
}
