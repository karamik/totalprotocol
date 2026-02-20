package main

import (
	"fmt"
	"time"
	"total-protocol/internal/core"
)

func main() {
	fmt.Println("=== TOTAL Protocol Orchestrator v.8.2 ===")
	fmt.Println("[*] Инициализация систем Sentinel Core...")

	// Подключаем защитника
	guard := core.NewGuard()

	fmt.Println("[*] Мониторинг безопасности активен.")
	
	for {
		// Имитация сигналов от железа (в будущем пойдут из Verilog)
		hwEntropyValid := true 
		hwNttError := false    

		// Постоянная проверка безопасности
		guard.MonitorSignals(hwEntropyValid, hwNttError)

		// Индикация жизни системы
		fmt.Print(".") 
		time.Sleep(2 * time.Second)
	}
}
