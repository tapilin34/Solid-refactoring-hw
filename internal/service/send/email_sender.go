package send
import 	"fmt"

type EmailSender struct {}

// реализуем тут интерфейс Notifier
func (e *EmailSender) Send(customer string) error {
	fmt.Printf("Уведомление отправлено по email клиенту %s\n", customer)
	return nil
}