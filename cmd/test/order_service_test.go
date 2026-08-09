package test

import (
	"errors"
	"reflect"
	"testing"

	"example/solid/internal/repository/model"
	"example/solid/internal/service/order"
)

type MockWriter struct {
	CreatedOrder model.Order
	CreateCalled bool
	Err          error
}

func (m *MockWriter) CreateOrder(order model.Order) error {
	m.CreateCalled = true
	m.CreatedOrder = order

	return m.Err
}

type MockNotifier struct {
	Customer    string
	SendCalled  bool
	Err         error
}

func (m *MockNotifier) Send(customer string) error {
	m.SendCalled = true
	m.Customer = customer

	return m.Err
}

func TestOrderServiceCreateOrder(t *testing.T) {
	writer := &MockWriter{}
	notifier := &MockNotifier{}

	service := order.NewOrderService(writer, notifier)

	expectedOrder := model.Order{
		Customer: "customer@example.com",
		Products: []string{"apple", "banana"},
		Total:    10.5,
		Status:   "pending",
	}

	err := service.CreateOrder(expectedOrder)
	if err != nil {
		t.Fatalf("ожидалась успешная операция, получена ошибка: %v", err)
	}

	if !writer.CreateCalled {
		t.Fatal("CreateOrder репозитория не был вызван")
	}

	if !reflect.DeepEqual(writer.CreatedOrder, expectedOrder) {
		t.Errorf(
			"сохраненный заказ отличается от ожидаемого:\nполучен: %#v\nожидался: %#v",
			writer.CreatedOrder,
			expectedOrder,
		)
	}

	if !notifier.SendCalled {
		t.Fatal("Send уведомления не был вызван")
	}

	if notifier.Customer != expectedOrder.Customer {
		t.Errorf(
			"получатель уведомления: %q, ожидался %q",
			notifier.Customer,
			expectedOrder.Customer,
		)
	}
}

func TestOrderServiceCreateOrderRepositoryError(t *testing.T) {
	expectedError := errors.New("ошибка базы данных")

	writer := &MockWriter{
		Err: expectedError,
	}

	notifier := &MockNotifier{}

	service := order.NewOrderService(writer, notifier)

	orderData := model.Order{
		Customer: "customer@example.com",
		Products: []string{"apple"},
		Total:    5,
		Status:   "pending",
	}

	err := service.CreateOrder(orderData)
	if !errors.Is(err, expectedError) {
		t.Fatalf(
			"получена ошибка %v, ожидалась %v",
			err,
			expectedError,
		)
	}

	if !writer.CreateCalled {
		t.Fatal("CreateOrder репозитория не был вызван")
	}

	if notifier.SendCalled {
		t.Fatal("уведомление не должно отправляться при ошибке репозитория")
	}
}