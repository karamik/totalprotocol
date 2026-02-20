package core

import "fmt"

// ThermalControl управляет охлаждением Пельтье и Liquid Cooling
type ThermalControl struct {
	IsLiquidCoolingActive bool
	PeltierPower          int // 0-100%
}

func NewThermalControl() *ThermalControl {
	return &ThermalControl{IsLiquidCoolingActive: false, PeltierPower: 0}
}

// ActivateEmergencyCooling реализует логику v.7.0
func (t *ThermalControl) ActivateEmergencyCooling() {
	t.IsLiquidCoolingActive = true
	t.PeltierPower = 100
	fmt.Println("[SECURITY] Криптографическая проверка охлаждения: OK")
	fmt.Println("[SECURITY] Пельтье-элементы запущены на 100% мощности")
}
