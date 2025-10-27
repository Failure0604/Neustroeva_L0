Перед запуском:
  1. Установить PostgreSQL
  2. Установить GO
  3. Установить NATS Streaming Server

Для запуска сервиса:
  1. В PowerShell: nats-streaming-server
  2. Открыть \order-service
  3. go run .

Для создания и публикации заказа:
  1. Открыть файл \order-service\publisher\publisher.go
  2. Изменить данные заказа в строке order :
  3. Открыть в PowerShell \order-service\publisher
  4. go run publisher.go
