package main

import (
	"fmt"
	"time"
	"total-protocol/internal/core" // Подключаем твою защиту
)

func main() {
	fmt.Println("=== TOTAL Protocol Orchestrator v.8.2 ===")
	fmt.Println("[*] Инициализация систем Sentinel Core...")

	// Запускаем нашего "охранника"
	guard := core.NewGuard()

	fmt.Println("[*] Мониторинг безопасности активен.")
	
	for {
		// Здесь мы будем получать сигналы от твоего Verilog-кода
		hwEntropyValid := true // Если станет false — сработает защита
		hwNttError := false    // Если станет true — система встанет

		// Проверка состояния
		guard.MonitorSignals(hwEntropyValid, hwNttError)

		// Просто индикация работы
		fmt.Print(".") 
		time.Sleep(2 * time.Second)
	}
}
