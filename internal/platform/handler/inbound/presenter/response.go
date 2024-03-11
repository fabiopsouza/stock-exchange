package presenter

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

//goland:noinspection GoUnhandledErrorResult
func Return(w http.ResponseWriter, model interface{}) {
	json.NewEncoder(w).Encode(model)
}
