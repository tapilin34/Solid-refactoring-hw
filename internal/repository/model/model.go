package model

type RepositoryWriter  interface {
	CreateOrder(order Order) error
}

type Notifier interface {
	Send(customer string) error
}

type Order struct {
	ID       int
	Customer string
	Products [] string
	Total    float64
	Status   string
}