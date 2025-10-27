package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	connStr := "user=order_user password=kristina dbname=order_db host=localhost sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Не удалось проверить подключение к БД:", err)
	}
	log.Println("Подключено к PostgreSQL")
}

func saveOrderToDB(order Order) error {
	itemsJSON, _ := json.Marshal(order.Items)
	_, err := db.Exec(`
		INSERT INTO orders (id, customer_name, items, total)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO NOTHING`,
		order.ID, order.CustomerName, itemsJSON, order.Total)
	return err
}

func loadAllOrdersFromDB() (map[string]Order, error) {
	rows, err := db.Query("SELECT id, customer_name, items, total FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cache := make(map[string]Order)
	for rows.Next() {
		var id, customerName string
		var itemsJSON []byte
		var total float64
		if err := rows.Scan(&id, &customerName, &itemsJSON, &total); err != nil {
			return nil, err
		}
		var items []OrderItem
		json.Unmarshal(itemsJSON, &items)
		cache[id] = Order{
			ID:           id,
			CustomerName: customerName,
			Items:        items,
			Total:        total,
		}
	}
	return cache, nil
}