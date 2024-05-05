package sqlite

import (
	"database/sql"

	stockDomain "github.com/fabiopsouza/stock-exchange/stock/internal/core/domain/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/core/port/stock"
)

type Handler struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) stock.Repository {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Create(stock stockDomain.Stock) (int64, error) {
	return 0, nil
}

func (h *Handler) Update(stock stockDomain.Stock) error {
	return nil
}

func (h *Handler) Get(id int64) (stockDomain.Stock, error) {
	return stockDomain.Stock{}, nil
}

func (h *Handler) ListAll() ([]stockDomain.Stock, error) {
	return []stockDomain.Stock{}, nil
}
