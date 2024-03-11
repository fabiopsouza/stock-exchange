package presenter

import (
	"encoding/json"
	"net/http"
)

//goland:noinspection GoUnhandledErrorResult
func Message(w http.ResponseWriter, message string) {
	json.NewEncoder(w).Encode(Response{
		StatusCode: http.StatusOK,
		Message:    message,
	})
}
