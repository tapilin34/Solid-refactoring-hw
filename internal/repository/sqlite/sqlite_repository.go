package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"example/solid/internal/repository/model"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) CreateOrder(order model.Order) error {
	products, err := json.Marshal(order.Products)
	if err != nil {
		return fmt.Errorf("ошибка преобразования товаров: %w", err)
	}

	_, err = r.db.Exec(
		`INSERT INTO orders
			(customer, products, total, status)
		 VALUES (?, ?, ?, ?)`,
		order.Customer,
		string(products),
		order.Total,
		order.Status,
	)
	if err != nil {
		return fmt.Errorf("ошибка сохранения заказа: %w", err)
	}

	return nil
}