package stock

import stockDomain "github.com/fabiopsouza/stock-exchange/stock/internal/core/domain/stock"

type Repository interface {
	Create(stock stockDomain.Stock) (int64, error)
	Update(stock stockDomain.Stock) error
	Get(id int64) (stockDomain.Stock, error)
	ListAll() ([]stockDomain.Stock, error)
}
