package core

import (
	"math/rand"
	"time"
)

// HardwareOracle имитирует чтение из физических регистров FPGA
type HardwareOracle struct {
	Address uint64
}

func NewOracle() *HardwareOracle {
	return &HardwareOracle{Address: 0x40001000} // Пример адреса регистра
}

// ReadEntropyStatus возвращает состояние QRNG
// В реальной системе здесь будет чтение системного прерывания
func (o *HardwareOracle) ReadEntropyStatus() bool {
	rand.Seed(time.Now().UnixNano())
	// Имитируем, что в 99% случаев энтропия в порядке
	return rand.Float32() > 0.01
}

// ReadThermalSensors имитирует проверку температуры зон
func (o *HardwareOracle) ReadThermalSensors() float64 {
	// Базовая температура 35.0°C с небольшим отклонением
	return 35.0 + rand.Float64()*5.0
}
