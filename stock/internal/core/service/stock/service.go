package stock

import (
	"context"
	"errors"

	stockDomain "github.com/fabiopsouza/stock-exchange/stock/stock/internal/core/domain/stock"
	"github.com/fabiopsouza/stock-exchange/stock/stock/internal/platform/handler/outbound/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	Create(ctx context.Context, stock stockDomain.Stock) (stockDomain.Stock, error)
	Update(ctx context.Context, currentSymbol string, stock stockDomain.Stock) (stockDomain.Stock, error)
	Get(ctx context.Context, symbol string) (stockDomain.Stock, error)
	ListAll(ctx context.Context) ([]stockDomain.Stock, error)
}

type handler struct {
	repository mongodb.Repository
}

func NewService(repository mongodb.Repository) Service {
	return &handler{
		repository: repository,
	}
}

func (h *handler) Create(ctx context.Context, stock stockDomain.Stock) (stockDomain.Stock, error) {
	exist, err := h.exist(ctx, stock.Symbol)
	if err != nil {
		return stock.Stock{}, err
	}

	if exist {
		return stock.Stock{}, errors.New("stock already exists")
	}

	return h.repository.Create(ctx, stock)
}

func (h *handler) Update(ctx context.Context, currentSymbol string, stock stockDomain.Stock) (stockDomain.Stock, error) {
	existCurrentSymbol, err := h.exist(ctx, currentSymbol)
	if err != nil {
		return stock.Stock{}, err
	}
	if !existCurrentSymbol {
		return stock.Stock{}, errors.New("stock not found")
	}

	existNewSymbol, err := h.exist(ctx, stock.Symbol)
	if err != nil {
		return stock.Stock{}, err
	}
	if existNewSymbol {
		return stock.Stock{}, errors.New("symbol to update already exists")
	}

	return h.repository.Update(ctx, currentSymbol, stock)
}

func (h *handler) Get(ctx context.Context, symbol string) (stockDomain.Stock, error) {
	return h.repository.Get(ctx, symbol)
}

func (h *handler) ListAll(ctx context.Context) ([]stockDomain.Stock, error) {
	return h.repository.ListAll(ctx)
}

func (h *handler) exist(ctx context.Context, symbol string) (bool, error) {
	item, err := h.Get(ctx, symbol)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return item.Symbol != "", nil
}
