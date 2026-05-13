package cpuinfo

import (
	"fmt"
	"time"

	"github.com/klauspost/cpuid/v2"
	"github.com/shirou/gopsutil/v3/cpu"
)

func DisplayCPUid() string {
	// gopsutil
	info, _ := cpu.Info()
	percent, _ := cpu.Percent(time.Second, false)
	physical, _ := cpu.Counts(false)
	logical, _ := cpu.Counts(true)

	// cpuid
	cpuInfo := cpuid.CPU

	result := ""
	result += fmt.Sprintf("===== CPU Information =====\n")
	result += fmt.Sprintf("Brand: %s\n", cpuInfo.BrandName)
	result += fmt.Sprintf("Vendor: %s\n", info[0].VendorID)

	result += fmt.Sprintf("\n===== Cores & Threads =====\n")
	result += fmt.Sprintf("Physical Cores: %d\n", physical)
	result += fmt.Sprintf("Logical Cores: %d\n", logical)

	result += fmt.Sprintf("\n===== Cache =====\n")
	result += fmt.Sprintf("L1D Cache: %d KB\n", cpuInfo.Cache.L1D)
	result += fmt.Sprintf("L2 Cache: %d KB\n", cpuInfo.Cache.L2)
	result += fmt.Sprintf("L3 Cache: %d KB\n", cpuInfo.Cache.L3)

	result += fmt.Sprintf("\n===== Performance =====\n")
	result += fmt.Sprintf("CPU Usage: %.2f%%\n", percent[0])
	result += fmt.Sprintf("Frequency: %.2f MHz\n", info[0].Mhz)

	return result
}
