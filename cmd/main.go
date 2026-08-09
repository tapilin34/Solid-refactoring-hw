package main

import (
	"fmt"

	"example/solid/cmd/example"
	"example/solid/internal/repository/model"
	"example/solid/internal/service/order"
	"example/solid/internal/service/send"
)

func main() {
	r := example.TestRepository{}

	newOrder := model.Order{
		Customer: "customer@example.com",
		Products: []string{"apple", "banana"},
		Total:    10.5,
		Status:   "pending",
	}

	// Email-отправка
	emailService := order.NewOrderService(
		r,
		&send.EmailSender{},
	)

	if err := emailService.CreateOrder(newOrder); err != nil {
		fmt.Println("Ошибка при создании заказа через email:", err)
	}

	// SMS-отправка
	smsService := order.NewOrderService(
		r,
		&send.SMSSender{},
	)

	if err := smsService.CreateOrder(newOrder); err != nil {
		fmt.Println("Ошибка при создании заказа через SMS:", err)
	}
}