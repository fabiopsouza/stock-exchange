package presenter

import (
	"encoding/json"
	"net/http"
)

//goland:noinspection GoUnhandledErrorResult
func Return(w http.ResponseWriter, model interface{}) {
	json.NewEncoder(w).Encode(model)
}
