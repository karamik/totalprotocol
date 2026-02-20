package security

import (
	"errors"
	"fmt"
	"sync/atomic"
)

// MemoryGuard implements the Tier 2 security echelon for Sentinel Lite.
// It prevents State Poisoning and OOM attacks by enforcing strict memory quotas
// before the execution layer begins data deserialization.
type MemoryGuard struct {
	Limit     uint64 // Hard limit in bytes
	Allocated uint64 // Currently reserved memory in bytes (atomic)
}

// NewMemoryGuard creates a new instance of the guard with a limit in Megabytes.
func NewMemoryGuard(limitMB uint64) *MemoryGuard {
	return &MemoryGuard{
		Limit: limitMB * 1024 * 1024,
	}
}

// Reserve attempts to allocate a quota for an incoming batch.
// Returns an error if the allocation would exceed the hard limit.
func (mg *MemoryGuard) Reserve(bytes uint64) error {
	for {
		current := atomic.LoadUint64(&mg.Allocated)
		next := current + bytes
		
		if next > mg.Limit {
			return errors.New("SECURITY_ALERT: State Poisoning prevented (Memory quota exceeded)")
		}
		
		if atomic.CompareAndSwapUint64(&mg.Allocated, current, next) {
			return nil
		}
		// Retry if another routine updated the value first
	}
}

// Release frees the reserved quota after processing or on error.
func (mg *MemoryGuard) Release(bytes uint64) {
	for {
		current := atomic.LoadUint64(&mg.Allocated)
		var next uint64
		if current > bytes {
			next = current - bytes
		} else {
			next = 0
		}
		
		if atomic.CompareAndSwapUint64(&mg.Allocated, current, next) {
			break
		}
	}
}

// GetStatus returns the current memory usage for the SIEM dashboard.
func (mg *MemoryGuard) GetStatus() string {
	curr := atomic.LoadUint64(&mg.Allocated)
	return fmt.Sprintf("MemoryGuard Usage: %d/%d bytes (%.2f%%)", 
		curr, mg.Limit, float64(curr)/float64(mg.Limit)*100)
}
