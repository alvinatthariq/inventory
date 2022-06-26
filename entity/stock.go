package entity

type Stock struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability int     `json:"availability"`
	IsActive     bool    `json:"is_active"`
}

type CreateStock struct {
	Name         string  `json:"name" validate:"min=1"`
	Price        float64 `json:"price" validate:"gt=0"`
	Availability int     `json:"availability" validate:"gte=0"`
	IsActive     bool    `json:"is_active"`
}

type CreateStockRequest struct {
	Stocks []CreateStock `json:"stocks" validate:"required,dive,required"`
}
