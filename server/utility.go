package server

import (
	"encoding/json"
	"net/http"
)

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func makeError(w http.ResponseWriter, errCode ErrorCode) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	e := Err{Code: errCode}
	responseJSON(w, e)
}
