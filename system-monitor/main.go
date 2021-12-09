package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// define port
const PORT = ":8080"

// define system info type
type systemInfo struct {
	Hostname        string  `json:"hostname"`
	TotalMem        uint64  `json:"total_mem"`
	FreeMem         uint64  `json:"free_mem"`
	MemUsagePercent float64 `json:"mem_usage_percent"`
	Architecture    string  `json:"architecture"`
	OS              string  `json:"os"`
	NumOfCpuCores   int     `json:"num_of_cpu_cores"`
	CpuUsagePercent float64 `json:"cpu_usage_percent"`
}

func main() {
	http.HandleFunc("/", getSystemInfo)

	fmt.Println("Server is listening on port", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getSystemInfo(w http.ResponseWriter, r *http.Request) {
	totalMem, freeMem, memUseagePercent := getMemInfo()
	hostName, architecture, os := getHostInfo()
	numOfCpuCores, cpuUseagePercent := getCpuInfo()

	systemInfo := systemInfo{
		Hostname:        hostName,
		TotalMem:        totalMem,
		FreeMem:         freeMem,
		MemUsagePercent: memUseagePercent,
		Architecture:    architecture,
		OS:              os,
		NumOfCpuCores:   numOfCpuCores,
		CpuUsagePercent: cpuUseagePercent,
	}

	systemInfoBytes, err := json.Marshal(systemInfo)
	checkError(err)

	w.Write(systemInfoBytes)
}
