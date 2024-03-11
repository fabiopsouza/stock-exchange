package presenter

import (
	"encoding/json"
	"net/http"
)

func BadRequest(w http.ResponseWriter, message string) {
	ReturnError(w, http.StatusBadRequest, message)
}

//goland:noinspection GoUnhandledErrorResult
func ReturnError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		StatusCode: statusCode,
		Message:    message,
	})
}
