package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type Metrics struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
}

func getMetrics() (Metrics, error) {

	// CPU %
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return Metrics{}, err
	}

	// Memory %
	vm, err := mem.VirtualMemory()
	if err != nil {
		return Metrics{}, err
	}

	// Disk %
	diskStat, err := disk.Usage("/")
	if err != nil {
		return Metrics{}, err
	}

	metrics := Metrics{
		CPUUsage:    cpuPercent[0],
		MemoryUsage: vm.UsedPercent,
		DiskUsage:   diskStat.UsedPercent,
	}

	return metrics, nil
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {

	metrics, err := getMetrics()
	if err != nil {
		http.Error(w, "Unable to get metrics", http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.MarshalIndent(metrics, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)

	fmt.Println("Server running on http://localhost:8080/metrics")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
