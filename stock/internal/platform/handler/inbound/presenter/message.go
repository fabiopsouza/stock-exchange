package presenter

import (
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func ResponseMsg(w http.ResponseWriter, status int, msg string) {
	Return(w, status, Message{
		Message: msg,
	})
}
