package main

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func checkError(err error) {
	if err != nil {
		log.Printf("Error - %v", err)
	}
}

// Get memory info return total free memory in MB and used memory percentage
func getMemInfo() (uint64, uint64, float64) {
	v, _ := mem.VirtualMemory()

	totalMem := v.Total / 1000000     // in MB
	freeMem := v.Free / 1000000       // in MB
	memUseagePercent := v.UsedPercent // in %

	return totalMem, freeMem, memUseagePercent
}

// Get host info return hostname, architecture, os
func getHostInfo() (string, string, string) {
	architecture, _ := host.KernelArch()
	hostInfo, _ := host.Info()
	hostName := hostInfo.Hostname
	os := hostInfo.OS
	return hostName, architecture, os
}

// Get cpu info return numOfCpuCores cpuUseagePercent
func getCpuInfo() (int, float64) {
	numOfCpuCores, _ := cpu.Counts(true)
	cpuUseagePercent, _ := cpu.Percent(time.Second, false)
	return numOfCpuCores, cpuUseagePercent[0]
}
