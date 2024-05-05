package presenter

import (
	"net/http"
)

func BadRequest(w http.ResponseWriter, message string) {
	ReturnError(w, http.StatusBadRequest, message)
}

func InternalError(w http.ResponseWriter, err error) {
	ReturnError(w, http.StatusInternalServerError, err.Error())
}

func ReturnError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	Return(w, Message{
		StatusCode: statusCode,
		Message:    message,
	})
}
