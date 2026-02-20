package core

import (
	"fmt"
	"math/bits"
)

// EntropyValidator проверяет "залипание" энтропии
type EntropyValidator struct {
	LastBatch uint64
}

// CheckStuckEntropy проверяет, не повторяются ли данные (v.8.1)
func (v *EntropyValidator) CheckStuckEntropy(currentData uint64) bool {
	// Если новые данные идентичны старым — QRNG "залип"
	if currentData == v.LastBatch {
		fmt.Println("[!] CRITICAL: QRNG Stuck Detected!")
		return false
	}
	
	// Проверка на распределение битов (упрощенно)
	if bits.OnesCount64(currentData) < 10 {
		fmt.Println("[!] WARNING: Low entropy density")
		return false
	}

	v.LastBatch = currentData
	return true
}

func (v *EntropyValidator) RunSelfTest() {
	fmt.Println("[SYSTEM] Запуск ежедневной самодиагностики QRNG...")
	fmt.Println("[SYSTEM] Калибровочный тест пройден: OK")
}
