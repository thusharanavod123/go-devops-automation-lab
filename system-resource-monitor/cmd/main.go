package main

import (
	"fmt"
	"time"

	"example.com/system-resource-monitor/internal/metrics"
)

func main() {
	for {
		cpu := metrics.GetCPUUsage()
		mem := metrics.GetMemoryUsage()
		disk := metrics.GetDiskUsage()

		fmt.Printf("CPU: %.2f%% | Memory: %dMB | Disk: %.2f%%\n",
			cpu, mem, disk)

		time.Sleep(2 * time.Second)
	}
}
