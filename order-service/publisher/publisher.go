// publisher/publisher.go
package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

type Order struct {
	ID           string      `json:"id"`
	CustomerName string      `json:"customer_name"`
	Items        []OrderItem `json:"items"`
	Total        float64     `json:"total"`
}

type OrderItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func main() {
	sc, err := stan.Connect("test-cluster", "publisher")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	order := Order{
		ID:           "ORD-123",
		CustomerName: "Иван Петров",
		Items: []OrderItem{
			{Name: "Ноутбук", Price: 999.99, Count: 1},
			{Name: "Мышь", Price: 19.99, Count: 2},
		},
		Total: 1039.97,
	}

	data, _ := json.Marshal(order)
	err = sc.Publish("orders", data)
	if err != nil {
		log.Fatal("Ошибка публикации:", err)
	}
	log.Println("Опубликован заказ ORD-123")

	time.Sleep(1 * time.Second)
}