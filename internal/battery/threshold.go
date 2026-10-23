// Package battery contains small policies for reporting mouse battery state.
package battery

import "sync"

const LowThreshold = 10

// LowBatteryGate returns true once when a battery level first reaches the low
// threshold. It rearms after the mouse reports a charge above that threshold.
type LowBatteryGate struct {
	mu       sync.Mutex
	notified bool
}

func (g *LowBatteryGate) ShouldNotify(level int) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if level > LowThreshold {
		g.notified = false
		return false
	}
	if g.notified {
		return false
	}
	g.notified = true
	return true
}
