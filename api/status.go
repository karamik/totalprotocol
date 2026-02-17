package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	rand.Seed(time.Now().UnixNano())

	// Имитируем данные Sentinel Core
	status := map[string]interface{}{
		"entropy":       fmt.Sprintf("%X %X", rand.Intn(0xFFFF), rand.Intn(0xFFFF)),
		"finality":      0.72 + rand.Float64()*(0.98-0.72),
		"is_processing": rand.Float64() > 0.85,
		"status":        "ACTIVE",
	}

	json.NewEncoder(w).Encode(status)
}
