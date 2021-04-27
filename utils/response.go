package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string
	Data    interface{}
}

func RespondWithError(w http.ResponseWriter, httpStatus int, message string, data interface{}) {
	var resp response
	resp.Message = message
	if data != nil {
		resp.Data = data
	}
	RespondWithJSON(w, httpStatus, resp)

}

func RespondWithJSON(w http.ResponseWriter, httpStatus int, payload response) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(response)
}
