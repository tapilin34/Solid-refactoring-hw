package main

import (
	"database/sql"
	"log"

	"example/solid/internal/repository/model"
	"example/solid/internal/repository/sqlite"
	"example/solid/internal/service/order"
	"example/solid/internal/service/send"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			customer TEXT NOT NULL,
			products TEXT NOT NULL,
			total REAL NOT NULL,
			status TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	r := sqlite.NewSQLiteRepository(db)

	newOrder := model.Order{
		Customer: "customer@example.com",
		Products: []string{"apple", "banana"},
		Total:    10.5,
		Status:   "pending",
	}

	// Создание заказа с отправкой Email
	emailService := order.NewOrderService(
		r,
		&send.EmailSender{},
	)

	if err := emailService.CreateOrder(newOrder); err != nil {
		log.Fatal("Ошибка при создании заказа через email:", err)
	}

	// Создание заказа с отправкой SMS
	smsService := order.NewOrderService(
		r,
		&send.SMSSender{},
	)

	if err := smsService.CreateOrder(newOrder); err != nil {
		log.Fatal("Ошибка при создании заказа через SMS:", err)
	}
}