package main

import (
	"log"
	"sync"
)

var (
	cache = make(map[string]Order)
	mu    sync.RWMutex
)

func setCache(order Order) {
	mu.Lock()
	defer mu.Unlock()
	cache[order.ID] = order
}

func getFromCache(id string) (Order, bool) {
	mu.RLock()
	defer mu.RUnlock()
	order, ok := cache[id]
	return order, ok
}

func restoreCacheFromDB() {
	orders, err := loadAllOrdersFromDB()
	if err != nil {
		log.Printf("Не удалось восстановить кэш из БД: %v", err)
		return
	}
	mu.Lock()
	cache = orders
	mu.Unlock()
	log.Printf("Восстановлено %d заказов из БД", len(orders))
}