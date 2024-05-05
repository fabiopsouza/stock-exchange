package stock

import (
	"encoding/json"
	"net/http"

	stockDomain "github.com/fabiopsouza/stock-exchange/stock/internal/core/domain/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/core/port/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/inbound/presenter"
)

type Handler struct {
	repository stock.Repository
}

func NewHandler(repository stock.Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input stockDomain.StockInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		presenter.BadRequest(w, "Error Binding Request")
		return
	}

	inputStock := input.ToStock()
	id, err := h.repository.Create(inputStock)
	if err != nil {
		presenter.InternalError(w, err)
		return
	}

	presenter.Created(w, id)
}
