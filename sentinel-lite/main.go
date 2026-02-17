package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SentinelNode представляет собой структуру высокопроизводительного узла L2
type SentinelNode struct {
	mu          sync.Mutex
	NodeID      string
	TPS         int
	QueueSize   int
	IsActive    bool
}

// Transaction имитирует структуру входящей транзакции
type Transaction struct {
	Hash      string
	Payload   string
	Timestamp time.Time
}

// NewTransaction создает новую транзакцию с хешем
func NewTransaction(data string) *Transaction {
	hash := sha256.Sum256([]byte(data + time.Now().String()))
	return &Transaction{
		Hash:      hex.EncodeToString(hash[:]),
		Payload:   data,
		Timestamp: time.Now(),
	}
}

// VerifyWithHardware имитирует вызов к Sentinel Core (FPGA)
func (n *SentinelNode) VerifyWithHardware(tx *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[Sentinel-Lite] 📡 TX [%.10s...] sent to Sentinel Core (FPGA)\n", tx.Hash)
	
	// Имитация задержки аппаратной проверки (Atomic Finality < 1ms)
	latency := time.Duration(rand.Intn(5)) * time.Millisecond
	time.Sleep(latency)

	fmt.Printf("[Sentinel-Core] ✅ TX [%.10s...] Verified. Latency: %v\n", tx.Hash, latency)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	node := &SentinelNode{
		NodeID:    "SENTINEL-LITE-PROXIMA-1",
		TPS:       1200000,
		QueueSize: 0,
		IsActive:  true,
	}

	fmt.Printf("=== TOTAL Protocol | Sentinel Lite v.1.1 ===\n")
	fmt.Printf("Node Status: ACTIVE | Target Throughput: %d TPS\n\n", node.TPS)

	var wg sync.WaitGroup
	txCount := 5 // Имитируем пачку транзакций

	for i := 0; i < txCount; i++ {
		wg.Add(1)
		tx := NewTransaction(fmt.Sprintf("Transfer-Batch-%d", i))
		go node.VerifyWithHardware(tx, &wg) // Запуск в параллельном потоке (Goroutine)
	}

	wg.Wait()
	fmt.Println("\n--------------------------------------------------")
	fmt.Println("TOTAL Status: SECURE | All Proofs Anchored via Hardware")
	fmt.Println("System Health: 100% | Thermal Guard: STABLE")
}
