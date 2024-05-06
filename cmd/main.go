package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/inbound/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/outbound/mongodb"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Starting server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Clients
	mongoClient := createMongo(ctx)

	// Repositories
	repository := mongodb.NewRepository(mongoClient)

	// Handlers
	handler := stock.NewHandler(ctx, repository)

	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/stock", handler.Create).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func createMongo(ctx context.Context) *mongo.Client {
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Could not connect to mongodb", err)
	}
	defer closeMongo(mongoClient)

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not ping mongodb", err)
	}

	return mongoClient
}

func closeMongo(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Could not disconnect mongodb", err)
	}
}
