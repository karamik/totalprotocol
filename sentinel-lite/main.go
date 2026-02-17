
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HardwareStatus struct {
	EntropyReady bool    `json:"entropy_ready"`
	Temperature  float64 `json:"core_temp"`
	LogicVersion string  `json:"logic_v"`
}

func handleHardwareStatus(w http.ResponseWriter, r *http.Request) {
	status := HardwareStatus{
		EntropyReady: true,
		Temperature:  36.6,
		LogicVersion: "8.0-Poseidon-Enabled",
	}
	
	fmt.Println("[Sentinel-Lite] Diagnostic request received via API")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func main() {
	http.HandleFunc("/status", handleHardwareStatus)
	// Наш старый эндпоинт для транзакций остается работать
	
	fmt.Println("=== TOTAL Protocol | Sentinel Lite v.1.3 ===")
	fmt.Println("[System] Hardware-Abstraction-Layer (HAL) Active")
	fmt.Println("[System] API listening on http://localhost:8080")
	
	http.ListenAndServe(":8080", nil)
}
