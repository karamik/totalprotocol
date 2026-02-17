package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Разрешаем всем (для теста)
}

// TOTALStatus - структура данных, которую видит фронтенд
type TOTALStatus struct {
	Entropy     string  `json:"entropy"`
	Finality    float64 `json:"finality"`
	FPGATemp    float64 `json:"temp"`
	BusIntegrity int     `json:"integrity"`
	IsProcessing bool    `json:"is_processing"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	for {
		// Имитируем работу Sentinel Core
		status := TOTALStatus{
			Entropy:      fmt.Sprintf("%X %X", rand.Intn(0xFFFF), rand.Intn(0xFFFF)),
			Finality:     0.70 + rand.Float64()*(0.95-0.70),
			FPGATemp:     31.0 + rand.Float64()*2.0,
			BusIntegrity: 100,
			IsProcessing: rand.Float64() > 0.8, // 20% шанс "транзакции"
		}

		message, _ := json.Marshal(status)
		ws.WriteMessage(websocket.TextMessage, message)
		time.Sleep(200 * time.Millisecond) // Частота обновления 5Hz
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("TOTAL Heartbeat Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
