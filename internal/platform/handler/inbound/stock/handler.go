package stock

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	stockDomain "github.com/fabiopsouza/stock-exchange/stock/internal/core/domain/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/core/service/stock"
	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/inbound/presenter"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	ctx     context.Context
	service stock.Service
}

func NewHandler(ctx context.Context, service stock.Service) *Handler {
	return &Handler{
		ctx:     ctx,
		service: service,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input stockDomain.Stock
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		presenter.BadRequest(w, "Error Binding Request")
		return
	}

	resp, err := h.service.Create(h.ctx, input)
	if err != nil {
		presenter.InternalError(w, err)
		return
	}

	presenter.ResponseData(w, resp)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	symbol, err := getSymbol(r)
	if err != nil {
		presenter.BadRequest(w, err.Error())
		return
	}

	var input stockDomain.Stock
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		presenter.BadRequest(w, "Error Binding Request")
		return
	}

	resp, err := h.service.Update(h.ctx, symbol, input)
	if err != nil {
		presenter.InternalError(w, err)
		return
	}

	presenter.ResponseData(w, resp)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	symbol, err := getSymbol(r)
	if err != nil {
		presenter.BadRequest(w, err.Error())
		return
	}

	resp, err := h.service.Get(h.ctx, symbol)
	if errors.Is(err, mongo.ErrNoDocuments) {
		presenter.NotFoundError(w)
		return
	}

	if err != nil {
		presenter.InternalError(w, err)
		return
	}

	presenter.ResponseData(w, resp)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.ListAll(h.ctx)
	if err != nil {
		presenter.InternalError(w, err)
		return
	}

	presenter.ResponseData(w, resp)
}

func getSymbol(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	symbol, ok := vars["symbol"]
	if !ok {
		return "", errors.New("param symbol required")
	}

	return symbol, nil
}
