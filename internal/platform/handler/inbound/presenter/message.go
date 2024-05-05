package presenter

import "net/http"

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type IDMessage struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	ID         int64  `json:"id"`
}

func Created(w http.ResponseWriter, id int64) {
	Return(w, IDMessage{
		StatusCode: http.StatusCreated,
		Message:    "Created successfully",
		ID:         id,
	})
}

func Updated(w http.ResponseWriter, id int64) {
	Return(w, Message{
		StatusCode: http.StatusOK,
		Message:    "Updated successfully",
	})
}

func Deleted(w http.ResponseWriter, id int64) {
	Return(w, Message{
		StatusCode: http.StatusOK,
		Message:    "Deleted successfully",
	})
}
