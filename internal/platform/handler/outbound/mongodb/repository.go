package mongodb

import (
	"context"

	stockDomain "github.com/fabiopsouza/stock-exchange/stock/internal/core/domain/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/core/port/stock"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	collection *mongo.Collection
}

func NewRepository(mongoClient *mongo.Client) stock.Repository {
	return &Handler{
		collection: mongoClient.Database("stockdb").Collection("stocks"),
	}
}

func (h *Handler) Create(ctx context.Context, stock stockDomain.Stock) (uuid.UUID, error) {
	_, err := h.collection.InsertOne(ctx, stock)
	if err != nil {
		return uuid.UUID{}, err
	}

	return stock.ID, nil
}

func (h *Handler) Update(ctx context.Context, stock stockDomain.Stock) (stockDomain.Stock, error) {
	return stockDomain.Stock{}, nil
}

func (h *Handler) Get(ctx context.Context, id int64) (stockDomain.Stock, error) {
	return stockDomain.Stock{}, nil
}

func (h *Handler) ListAll(ctx context.Context) ([]stockDomain.Stock, error) {
	return []stockDomain.Stock{}, nil
}
