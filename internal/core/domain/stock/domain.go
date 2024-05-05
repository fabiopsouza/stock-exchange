package stock

import "github.com/google/uuid"

type StockInput struct {
	Code string
}

type Stock struct {
	ID   uuid.UUID
	Code string
}

func (in StockInput) ToStock() Stock {
	return Stock{
		ID:   uuid.New(),
		Code: in.Code,
	}
}
