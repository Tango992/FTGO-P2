package handler

import (
	"encoding/json"
	"net/http"
)

func JsonWriter(w *http.ResponseWriter, statusCode int, data any) {
	(*w).Header().Add("Content-Type", "application/json")
	(*w).WriteHeader(statusCode)
	json.NewEncoder(*w).Encode(data)
}