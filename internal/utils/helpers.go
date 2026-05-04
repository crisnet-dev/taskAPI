package utils

import (
	"encoding/json"
	"net/http"
)

func HttpError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"message": message,
	})
}

func HttpResponse(w http.ResponseWriter, data map[string]any, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// func HttpResponse(w http.ResponseWriter, message string, statusCode int) {
// 	w.WriteHeader(statusCode)
// 	json.NewEncoder(w).Encode(message)
// }
