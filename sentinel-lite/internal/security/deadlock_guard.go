package security

import (
	"fmt"
	"sync"
	"time"
)

// DeadlockGuard мониторит состояние ZK-акселератора
type DeadlockGuard struct {
	mu          sync.Mutex
	PendingOps  map[string]time.Time // ID операции -> Время начала
	Threshold   time.Duration        // Лимит ожидания (500ms)
}

func NewDeadlockGuard(threshold time.Duration) *DeadlockGuard {
	return &DeadlockGuard{
		PendingOps: make(map[string]time.Time),
		Threshold:  threshold,
	}
}

// StartOp фиксирует начало взаимодействия с FPGA
func (dg *DeadlockGuard) StartOp(opID string) {
	dg.mu.Lock()
	defer dg.mu.Unlock()
	dg.PendingOps[opID] = time.Now()
}

// EndOp подтверждает завершение операции (ZK-пруф получен)
func (dg *DeadlockGuard) EndOp(opID string) {
	dg.mu.Lock()
	defer dg.mu.Unlock()
	delete(dg.PendingOps, opID)
}

// Verify проверяет наличие "зависших" процессов
func (dg *DeadlockGuard) Verify() error {
	dg.mu.Lock()
	defer dg.mu.Unlock()

	now := time.Now()
	for id, startTime := range dg.PendingOps {
		if now.Sub(startTime) > dg.Threshold {
			return fmt.Errorf("DEADLOCK_DETECTED: Op [%s] exceeded threshold. Initiating State Reset", id)
		}
	}
	return nil
}
