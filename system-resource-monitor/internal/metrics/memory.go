package metrics

import "math/rand"

// GetMemoryUsage returns a fake memory usage in megabytes.
func GetMemoryUsage() int64 {
	// Return a value between 100 and 16000 MB
	return 100 + rand.Int63n(15900)
}
