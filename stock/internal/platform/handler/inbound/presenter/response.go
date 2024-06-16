package presenter

import (
	"encoding/json"
	"net/http"
)

func ResponseData(w http.ResponseWriter, value interface{}) {
	Return(w, http.StatusOK, value)
}

//goland:noinspection GoUnhandledErrorResult
func Return(w http.ResponseWriter, status int, model interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(model)
}
