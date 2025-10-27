Перед запуском:
  1. Установить PostgreSQL
  2. Установить GO
  3. Установить NATS Streaming Server

Для создания и публикации заказа:
  1. В PowerShell: nats-streaming-server
  2. Открыть \order-service\publisher\publisher.go
  3. Изменить данные заказа в строке order :
  4. Открыть \order-service\publisher\ в PowerShell
  5. go run publisher.go

Для запуска сервиса:
  1. Открыть \order-service
  2. go run .
