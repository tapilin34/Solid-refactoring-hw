package order

import "example/solid/internal/repository/model"

type OrderService struct{
	repository model.RepositoryWriter
	notifier   model.Notifier
}

func NewOrderService(
	repository model.RepositoryWriter,
	notifier model.Notifier,
) *OrderService {
	return &OrderService{
		repository: repository,
		notifier:   notifier,
	}
}

func (o *OrderService) CreateOrder(order model.Order) error {
	// Создание заказа в БД
	if err := o.repository.CreateOrder(order); err != nil {
		return err
	}

	// Отправка уведомления
	if err := o.notifier.Send(order.Customer); err != nil {
		return err
	}
	return nil
}