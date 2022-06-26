package stock

import (
	"context"
	"database/sql"
	"fmt"
	"inventory/entity"

	"github.com/google/uuid"
)

func (s *stock) CreateStock(v entity.CreateStock) (entity.Stock, error) {
	stock := entity.Stock{
		ID:           uuid.New().String(),
		Name:         v.Name,
		Price:        v.Price,
		Availability: v.Availability,
		IsActive:     v.IsActive,
	}

	_, err := s.createSQLStock(stock)
	if err != nil {
		return stock, err
	}

	return stock, nil
}

func (s *stock) GetStockByID(id string) (entity.Stock, error) {
	return s.getSQLStockByID(id)
}

func (s *stock) createSQLStock(v entity.Stock) (entity.Stock, error) {
	// begin tx
	tx, err := s.sqlClient.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return v, err
	}

	// insert stock
	_, err = tx.Exec(CreateStockQuery,
		v.ID,
		v.Name,
		v.Price,
		v.Availability,
		v.IsActive,
	)
	if err != nil {
		_ = tx.Rollback()
		return v, err
	}

	// commit tx
	err = tx.Commit()
	if err != nil {
		return v, err
	}

	return v, nil
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
		return result, fmt.Errorf("stock not found")
	} else if err != nil {
		return result, err
	}

	return result, nil
}
