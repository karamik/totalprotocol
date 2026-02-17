package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Response - структура данных для фронтенда
type Response struct {
	Entropy      string  `json:"entropy"`
	Finality     float64 `json:"finality"`
	IsProcessing bool    `json:"is_processing"`
	Status       string  `json:"status"`
	Timestamp    int64   `json:"timestamp"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. НАСТРОЙКА CORS (Разрешаем запросы с любого домена)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// 2. ОБРАБОТКА PREFLIGHT-ЗАПРОСА (Важно для браузеров)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 3. ГЕНЕРАЦИЯ ДАННЫХ
	rand.Seed(time.Now().UnixNano())

	// Имитируем показатели Sentinel Core
	status := Response{
		Entropy:      fmt.Sprintf("%X %X", rand.Intn(0xFFFF), rand.Intn(0xFFFF)),
		Finality:     0.85 + rand.Float64()*(0.99-0.85), // Финальность 85-99%
		IsProcessing: rand.Float64() > 0.80,             // 20% шанс вспышки (транзакции)
		Status:       "ACTIVE",
		Timestamp:    time.Now().Unix(),
	}

	// 4. ОТПРАВКА ОТВЕТА
	w.WriteHeader(http.StatusCreated) // Или http.StatusOK
	json.NewEncoder(w).Encode(status)
}
