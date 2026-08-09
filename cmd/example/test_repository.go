package example

import (
	"fmt"

	"example/solid/internal/repository/model"
)

type TestRepository struct {}
// Сохраняем тестовый заказ
func (r TestRepository) CreateOrder(order model.Order) error {
	fmt.Printf(
		"Заказ сохранен: customer=%s, total=%.2f, status=%s\n",
		order.Customer,
		order.Total,
		order.Status,
	)

	return nil
}