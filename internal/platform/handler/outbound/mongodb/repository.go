package mongodb

import (
	"context"
	"fmt"

	stockDomain "github.com/fabiopsouza/stock-exchange/stock/internal/core/domain/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/core/port/stock"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	collection *mongo.Collection
}

func NewRepository(mongoClient *mongo.Client) stock.Repository {
	return &Handler{
		collection: mongoClient.Database("stock").Collection("stocks"),
	}
}

func (h *Handler) Create(ctx context.Context, stock stockDomain.Stock) (int64, error) {
	result, err := h.collection.InsertOne(ctx, stock)
	if err != nil {
		return 0, err
	}

	fmt.Println(result)

	return 0, nil
}

func (h *Handler) Update(ctx context.Context, stock stockDomain.Stock) error {
	return nil
}

func (h *Handler) Get(ctx context.Context, id int64) (stockDomain.Stock, error) {
	return stockDomain.Stock{}, nil
}

func (h *Handler) ListAll(ctx context.Context) ([]stockDomain.Stock, error) {
	return []stockDomain.Stock{}, nil
}
