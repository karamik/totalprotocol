package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// TransactionRequest структура входящего JSON-запроса
type TransactionRequest struct {
	Data string `json:"data"`
}

// TransactionResponse структура ответа от ноды
type TransactionResponse struct {
	Hash      string `json:"hash"`
	Status    string `json:"status"`
	Finality  string `json:"finality"`
	Timestamp string `json:"timestamp"`
}

type SentinelNode struct {
	Version string
	mu      sync.Mutex
}

// HandleProcess обрабатывает входящие HTTP запросы
func (n *SentinelNode) HandleProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Имитация работы Sentinel Core (Hardware Verification)
	txHash := n.generateHash(req.Data)
	
	fmt.Printf("[Sentinel-Lite] 📡 Incoming Data: %s\n", req.Data)
	fmt.Printf("[Sentinel-Core] ⚡ Hardware Verification for [%.10s...]\n", txHash)

	// Эмуляция задержки аппаратного ускорения
	time.Sleep(50 * time.Millisecond)

	response := TransactionResponse{
		Hash:      txHash,
		Status:    "VERIFIED_BY_HARDWARE",
		Finality:  "ATOMIC",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (n *SentinelNode) generateHash(data string) string {
	hash := sha256.Sum256([]byte(data + time.Now().String()))
	return hex.EncodeToString(hash[:])
}

func main() {
	node := &SentinelNode{Version: "1.2-API"}

	// Роутинг
	http.HandleFunc("/process", node.HandleProcess)

	fmt.Printf("=== TOTAL Protocol | Sentinel Lite v.%s ===\n", node.Version)
	fmt.Println("[System] API Server is running on http://localhost:8080")
	fmt.Println("[System] Endpoint: POST /process")
	fmt.Println("TOTAL Status: ACTIVE | Standing by for transactions...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
