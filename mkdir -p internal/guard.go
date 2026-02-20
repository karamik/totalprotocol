package core

import (
	"fmt"
	"os"
	"sync"
)

type SentinelGuard struct {
	mu            sync.Mutex
	isCompromised bool
}

func NewGuard() *SentinelGuard {
	return &SentinelGuard{}
}

func (g *SentinelGuard) EmergencyShutdown(reason string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.isCompromised {
		return
	}
	g.isCompromised = true

	fmt.Printf("\n[!!!] EMERGENCY SHUTDOWN: %s\n", reason)
	fmt.Println("[...] Очистка ключей из памяти...")
	fmt.Println("[!] TOTAL Protocol остановлен.")
	
	os.Exit(1)
}

func (g *SentinelGuard) MonitorSignals(entropyValid bool, nttError bool) {
	if !entropyValid {
		g.EmergencyShutdown("Сбой QRNG")
	}
	if nttError {
		g.EmergencyShutdown("Сбой NTT")
	}
}
