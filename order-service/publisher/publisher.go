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
		ID:           "ORD-555",
		CustomerName: "Виктория Неустроева",
		Items: []OrderItem{
			{Name: "Ноутбук", Price: 1000, Count: 1},
			{Name: "Мышь", Price: 100, Count: 2},
		},
		Total: 1200,
	}

	data, _ := json.Marshal(order)
	err = sc.Publish("orders", data)
	if err != nil {
		log.Fatal("Ошибка публикации:", err)
	}
	log.Println("Заказ опубликован")

	time.Sleep(1 * time.Second)

}
