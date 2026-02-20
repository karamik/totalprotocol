package main

import (
	"fmt"
	"time"
	"total-protocol/internal/core"
)

func main() {
	fmt.Println("=== TOTAL Protocol Orchestrator v.8.2 ===")
	fmt.Println("[*] Запуск Oracle Shield и Sentinel Guard...")

	// Инициализируем защиту и датчики
	guard := core.NewGuard()
	oracle := core.NewOracle()

	fmt.Println("[*] Система активна. Мониторинг физических параметров...")
	
	for {
		// Читаем данные из нашего "железного моста"
		entropyStatus := oracle.ReadEntropyStatus()
		temp := oracle.ReadThermalSensors()

		fmt.Printf("[LOG] Temp: %.2f°C | Entropy: %v\n", temp, entropyStatus)

		// Если энтропия упала (сигнал атаки или сбоя) — Guard рубит систему
		if !entropyStatus {
			guard.MonitorSignals(false, false)
		}

		// Если температура выше критической (например, 39°C) — предупреждение
		if temp > 39.0 {
			fmt.Printf("[!] WARNING: Высокая температура! Активация Peltier-элементов...\n")
		}

		time.Sleep(2 * time.Second)
	}
}
