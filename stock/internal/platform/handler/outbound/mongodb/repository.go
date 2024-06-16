package mongodb

import (
	"context"
	"errors"

	stockDomain "github.com/fabiopsouza/stock-exchange/stock/stock/internal/core/domain/stock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, stock stockDomain.Stock) (stockDomain.Stock, error)
	Update(ctx context.Context, currentSymbol string, stock stockDomain.Stock) (stockDomain.Stock, error)
	Get(ctx context.Context, symbol string) (stockDomain.Stock, error)
	ListAll(ctx context.Context) ([]stockDomain.Stock, error)
}

type handler struct {
	collection *mongo.Collection
}

func NewRepository(mongoClient *mongo.Client) Repository {
	return &handler{
		collection: mongoClient.Database("stockdb").Collection("stocks"),
	}
}

func (h *handler) Create(ctx context.Context, stock stockDomain.Stock) (stockDomain.Stock, error) {
	_, err := h.collection.InsertOne(ctx, stock)
	if err != nil {
		return stock.Stock{}, err
	}

	return stock, nil
}

func (h *handler) Update(ctx context.Context, currentSymbol string, stock stockDomain.Stock) (stockDomain.Stock, error) {
	filter := bson.D{{"symbol", currentSymbol}}

	result, err := h.collection.UpdateOne(ctx, filter, bson.D{
		{"$set", stock},
	})
	if err != nil {
		return stock.Stock{}, err
	}

	if result.ModifiedCount == 0 {
		return stock.Stock{}, errors.New("0 documents modified")
	}

	return stock, nil
}

func (h *handler) Get(ctx context.Context, symbol string) (stockDomain.Stock, error) {
	filter := bson.D{{"symbol", symbol}}

	var stock stockDomain.Stock
	err := h.collection.FindOne(ctx, filter).Decode(&stock)
	if err != nil {
		return stock.Stock{}, err
	}

	return stock, nil
}

func (h *handler) ListAll(ctx context.Context) ([]stockDomain.Stock, error) {
	cur, err := h.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var results []stockDomain.Stock
	for cur.Next(ctx) {
		var item stockDomain.Stock
		err = cur.Decode(&item)

		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
