package stock

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"inventory/entity"

	"github.com/google/uuid"
)

func (s *stock) CreateStock(v entity.CreateStockRequest) ([]entity.Stock, error) {
	stocks := []entity.Stock{}

	err := s.validate.Struct(v)
	if err != nil {
		return []entity.Stock{}, &entity.Error{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	for _, _v := range v.Stocks {
		stocks = append(stocks, entity.Stock{
			ID:           uuid.New().String(),
			Name:         _v.Name,
			Price:        _v.Price,
			Availability: _v.Availability,
			IsActive:     _v.IsActive,
		})
	}

	_, err = s.createSQLStockBulk(stocks)
	if err != nil {
		return stocks, err
	}

	return stocks, nil
}

func (s *stock) GetStockByID(id string) (entity.Stock, error) {
	return s.getSQLStockByID(id)
}

func (s *stock) getSQLStockByID(id string) (entity.Stock, error) {
	var result entity.Stock

	row := s.sqlClient.QueryRow(GetStockQueryByID, id)

	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Price,
		&result.Availability,
		&result.IsActive,
	)

	if err == sql.ErrNoRows {
		return result, &entity.Error{
			Err:  fmt.Errorf("stock not found"),
			Code: http.StatusNotFound,
		}
	} else if err != nil {
		return result, err
	}

	return result, nil
}

func (s *stock) createSQLStockBulk(v []entity.Stock) ([]entity.Stock, error) {
	// begin tx
	tx, err := s.sqlClient.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return v, err
	}

	// prepare stmt
	stmt, err := tx.Prepare(CreateStockQuery)
	if err != nil {
		return v, err
	}
	defer stmt.Close()

	for _, _v := range v {
		_, err = stmt.Exec(
			_v.ID,
			_v.Name,
			_v.Price,
			_v.Availability,
			_v.IsActive,
		)
		if err != nil {
			_ = tx.Rollback()
			return v, err
		}
	}

	// commit tx
	err = tx.Commit()
	if err != nil {
		return v, err
	}

	return v, nil
}
