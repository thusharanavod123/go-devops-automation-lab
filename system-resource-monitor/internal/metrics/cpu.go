package metrics

import (
	"math/rand"
)

func GetCPUUsage() float64 {
	return rand.Float64() * 100
}
