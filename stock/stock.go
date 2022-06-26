package stock

import (
	"database/sql"

	"inventory/entity"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

type stock struct {
	sqlClient *sql.DB
	validate  *validator.Validate
}

type StockItf interface {
	CreateStock(v entity.CreateStockRequest) ([]entity.Stock, error)
	GetStockByID(id string) (entity.Stock, error)
}

func InitStock(sql *sql.DB) StockItf {
	return &stock{
		sqlClient: sql,
		validate:  validator.New(),
	}
}
