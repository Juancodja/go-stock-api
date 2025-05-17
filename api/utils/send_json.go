package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func SendError(w http.ResponseWriter, status int, message string) {
	SendJSON(w, status, map[string]string{
		"error": message,
	})
}
