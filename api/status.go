package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Структура ответа
type Response struct {
	Entropy      string  `json:"entropy"`
	Finality     float64 `json:"finality"`
	IsProcessing bool    `json:"is_processing"`
	Timestamp    int64   `json:"timestamp"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Разрешаем фронтенду с GitHub Pages обращаться к этому API (CORS)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	rand.Seed(time.Now().UnixNano())

	// Генерируем "живые" данные
	status := Response{
		Entropy:      fmt.Sprintf("%X %X", rand.Intn(0xFFFF), rand.Intn(0xFFFF)),
		Finality:     0.72 + rand.Float64()*(0.98-0.72),
		IsProcessing: rand.Float64() > 0.85, // 15% шанс вспышки
		Timestamp:    time.Now().Unix(),
	}

	json.NewEncoder(w).Encode(status)
}
