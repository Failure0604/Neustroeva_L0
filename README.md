Перед запуском:
  1. Установить PostgreSQL
  2. Установить GO
  3. Установить NATS Streaming Server

Для создания и публикации заказа:
  1. Открыть \order-service\publisher\publisher.go
  2. Изменить данные заказа в строке order :
  3. Открыть \order-service\publisher\ в PowerShell
  4. go run publisher.go

Для запуска, в PowerShell:
  1. nats-streaming-server
  2. Открыть \order-service
  3. go run .
