package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/order/")
	if id == "" {
		http.Error(w, "Требуется ID заказа", http.StatusBadRequest)
		return
	}

	order, ok := getFromCache(id)
	if !ok {
		http.Error(w, "Заказ не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Просмотр заказа</title>
</head>
<body>
  <h2>Поиск заказа по ID</h2>
  <input id="orderId" placeholder="Введите ID заказа (например: ORD-123)" style="width:300px; padding:8px;">
  <button onclick="fetchOrder()">Найти</button>
  <pre id="result" style="margin-top:20px; background:#f5f5f5; padding:10px;"></pre>

  <script>
    async function fetchOrder() {
      const id = document.getElementById('orderId').value.trim();
      if (!id) return;
      
      const res = await fetch('/order/' + encodeURIComponent(id));
      const result = document.getElementById('result');
      
      if (res.ok) {
        const data = await res.json();
        result.textContent = JSON.stringify(data, null, 2);
      } else {
        result.textContent = 'Ошибка: ' + (res.status === 404 ? 'заказ не найден' : 'сервер недоступен');
      }
    }
  </script>
</body>
</html>
	`))
}