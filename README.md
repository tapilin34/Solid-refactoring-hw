# Домашнее задание "Рефакторинг программы по принципам SOLID" - Тапилин Артём
Структура проекта:
- в корне находятся go.mod(файл инициализирующий проект и содержащий основную информацию о наименовании) и go.sum (информация о подключенных зависимостях)
- основной файл для запуска находится в /cmd/main.go
- упрощенная тестовая структура, используемая в main.go - /cmd/example/test_repository.go (просто вывод на экран без работы с БД).
- UNIT тестирование, формировал с помощью ИИ - /cmd/test/order_service_test.go
- /internal/ содержит модель, с описанием основной структуры и определением интерфейсов (/internal/repository/model//model.go)
- сервис создания заказа(/internal/service/order/order_service.go)
- сервисы отправки уведомлений(/internal/service/send/email_sender.go и sms_sender.go)
- подключение структуры SQLiteRepository /internal/repository/sqlite/sqlite_repository.go