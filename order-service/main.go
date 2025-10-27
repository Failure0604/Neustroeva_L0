package main

import (
	"encoding/json"
	"log"
	"net/http"

	stan "github.com/nats-io/stan.go"
)
const (
	clusterID = "test-cluster"
	clientID  = "order-service"
	subject   = "orders"
)

func main() {
	initDB()
	restoreCacheFromDB()

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Fatal("Не удалось подключиться к NATS Streaming:", err)
	}
	defer sc.Close()

	_, err = sc.Subscribe(subject, func(msg *stan.Msg) {
		var order Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Printf("Неверный формат сообщения: %v", err)
			return
		}
		if order.ID == "" {
			log.Println("Пропущено сообщение: отсутствует ID")
			return
		}

		if err := saveOrderToDB(order); err != nil {
			log.Printf("Ошибка сохранения в БД: %v", err)
			return
		}
		setCache(order)
		log.Printf("Сохранён и закэширован заказ %s", order.ID)
	}, stan.DurableName("order-durable"))

	if err != nil {
		log.Fatal("Ошибка подписки на канал:", err)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/order/", getOrderHandler)
	log.Println("HTTP-сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}