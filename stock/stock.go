package stock

import (
	"database/sql"
	"inventory/entity"

	_ "github.com/go-sql-driver/mysql"
)

type stock struct {
	sqlClient *sql.DB
}

type StockItf interface {
	CreateStock(v entity.CreateStock) (entity.Stock, error)
	GetStockByID(id string) (entity.Stock, error)
}

func InitStock(sql *sql.DB) StockItf {
	return &stock{
		sqlClient: sql,
	}
}
