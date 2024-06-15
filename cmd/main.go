package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	stockService "github.com/fabiopsouza/stock-exchange/stock/internal/core/service/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/inbound/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/outbound/mongodb"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Starting server...")
	ctx := context.Background()

	// Clients
	mongoClient := createMongo(ctx)
	defer closeMongo(mongoClient)

	// Repositories
	repository := mongodb.NewRepository(mongoClient)

	// Services
	service := stockService.NewService(repository)

	// Handlers
	handler := stock.NewHandler(ctx, service)

	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/stock", handler.Create).Methods("POST")
	r.HandleFunc("/stock/{symbol}", handler.Update).Methods("PUT")
	r.HandleFunc("/stock/{symbol}", handler.Get).Methods("GET")
	r.HandleFunc("/stock", handler.List).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func createMongo(ctx context.Context) *mongo.Client {
	log.Println("Connecting to MongoDB...")
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:12345@127.0.0.1:27017/stockdb?authSource=stockdb"))

	if err != nil {
		log.Fatal("Could not connect to mongodb", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not ping mongodb", err)
	}

	log.Println("MongoDB connected successfully")
	return mongoClient
}

func closeMongo(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Could not disconnect mongodb", err)
	}
	log.Println("MongoDB disconnected")
}
