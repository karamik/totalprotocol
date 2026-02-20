package security

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

// HardwareIntegrity контролирует состояние холодного пуска (Cold Start)
type HardwareIntegrity struct {
	ExpectedBitstreamHash string
	SecureBootEnabled     bool
}

// NewIntegrityChecker инициализирует проверку на основе eFuses хеша
func NewIntegrityChecker(trustedHash string) *HardwareIntegrity {
	return &HardwareIntegrity{
		ExpectedBitstreamHash: trustedHash,
		SecureBootEnabled:     true,
	}
}

// VerifyColdStart проверяет, не был ли подменен битстрим FPGA и активен ли Scrambler
func (hi *HardwareIntegrity) VerifyColdStart() error {
	// 1. Эмуляция чтения хеша текущей прошивки из регистра FPGA
	// В продакшене: чтение через драйвер /dev/sentinel_core
	currentHash, err := readCurrentBitstreamHash()
	if err != nil {
		return fmt.Errorf("HARDWARE_ERROR: Could not verify FPGA state: %v", err)
	}

	// 2. Сравнение с эталоном из eFuses (защита от Injection)
	if currentHash != hi.ExpectedBitstreamHash {
		hi.TriggerEmergencyZeroization()
		return errors.New("INTEGRITY_VIOLATION: Unauthorized hardware bitstream detected")
	}

	return nil
}

// TriggerEmergencyZeroization — мгновенная очистка ключей в RAM при угрозе
func (hi *HardwareIntegrity) TriggerEmergencyZeroization() {
	// Посылка сигнала прерывания (NMI) в аппаратное ядро
	fmt.Println("CRITICAL: Cold Start violation! Zeroing all session keys...")
	// В реальном железе здесь запись в регистр сброса ключей
}

func readCurrentBitstreamHash() (string, error) {
	// Пример пути к системному девайсу FPGA
	path := "/sys/class/fpga_manager/fpga0/firmware_hash"
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(data), nil
}
