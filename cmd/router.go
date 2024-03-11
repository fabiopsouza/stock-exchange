package main

import (
	"log"
	"net/http"

	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/inbound/stock"
	"github.com/gorilla/mux"
)

func StartRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/stock", stock.Create).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r)) // TODO get from config
}
