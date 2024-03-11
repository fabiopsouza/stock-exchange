package stock

import (
	"encoding/json"
	"net/http"

	stockDomain "github.com/fabiopsouza/stock-exchange/stock/internal/core/domain"
	"github.com/fabiopsouza/stock-exchange/stock/internal/platform/handler/inbound/presenter"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var input stockDomain.StockInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		presenter.BadRequest(w, "Error Binding Request")
		return
	}

	stock := input.ToStock()
	presenter.Return(w, stock)
}
