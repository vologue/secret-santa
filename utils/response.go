package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, httpStatus int, message string) {
	RespondWithJSON(w, httpStatus, map[string]string{"message": message})
}

func RespondWithJSON(w http.ResponseWriter, httpStatus int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(response)
}
