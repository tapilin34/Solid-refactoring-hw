package send

import "fmt"

type SMSSender struct {}

// реализуем тут интерфейс Notifier
func (s SMSSender) Send(customer string) error {
	fmt.Printf("Уведомление отправлено по sms клиенту %s\n", customer)
	return nil
}