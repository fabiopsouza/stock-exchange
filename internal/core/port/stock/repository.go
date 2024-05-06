package stock

import (
	"context"

	stockDomain "github.com/fabiopsouza/stock-exchange/stock/internal/core/domain/stock"
)

type Repository interface {
	Create(ctx context.Context, stock stockDomain.Stock) (int64, error)
	Update(ctx context.Context, stock stockDomain.Stock) error
	Get(ctx context.Context, id int64) (stockDomain.Stock, error)
	ListAll(ctx context.Context) ([]stockDomain.Stock, error)
}
