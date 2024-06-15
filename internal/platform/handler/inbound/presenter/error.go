package presenter

import (
	"net/http"
)

func BadRequest(w http.ResponseWriter, msg string) {
	Error(w, http.StatusBadRequest, msg)
}

func InternalError(w http.ResponseWriter, err error) {
	Error(w, http.StatusInternalServerError, err.Error())
}

func NotFoundError(w http.ResponseWriter) {
	Error(w, http.StatusNotFound, "Not Found")
}

func Error(w http.ResponseWriter, statusCode int, msg string) {
	ResponseMsg(w, statusCode, msg)
}
