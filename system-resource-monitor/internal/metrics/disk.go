package metrics

import "math/rand"

// GetDiskUsage returns a fake disk usage percentage (0.0 - 100.0).
func GetDiskUsage() float64 {
	return rand.Float64() * 100
}
