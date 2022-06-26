package stock

import (
	"inventory/entity"

	"github.com/google/uuid"
)

var inventory = map[string]entity.Stock{}

func CreateStock(v entity.CreateStock) (entity.Stock, error) {
	stock := entity.Stock{
		ID:           uuid.New().String(),
		Name:         v.Name,
		Price:        v.Price,
		Availability: v.Availability,
		IsActive:     v.IsActive,
	}

	inventory[stock.ID] = stock

	return stock, nil
}

func GetStockByID(id string) (entity.Stock, error) {
	return inventory[id], nil
}
