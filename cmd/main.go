package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/inbound/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/outbound/sqlite"
	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "./stock.db"

func main() {
	fmt.Println("Starting server...")

	// Clients
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal("Could not load db file")
	}

	// Repositories
	repository := sqlite.NewRepository(db)

	// Handlers
	handler := stock.NewHandler(repository)

	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/stock", handler.Create).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
