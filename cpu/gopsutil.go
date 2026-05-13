package cpuinfo

import (
	"time"

	"github.com/klauspost/cpuid/v2"
	"github.com/shirou/gopsutil/v3/cpu"
)

func DisplayCPUgopsutil() map[string]interface{} {
	// gopsutil
	info, _ := cpu.Info()
	percent, _ := cpu.Percent(time.Second, false)
	physical, _ := cpu.Counts(false)
	logical, _ := cpu.Counts(true)
	cpuInfo := cpuid.CPU

	return map[string]interface{}{
		"brand":          cpuInfo.BrandName,
		"vendor":         info[0].VendorID,
		"physical_cores": physical,
		"logical_cores":  logical,
		"l1_cache":       cpuInfo.Cache.L1D,
		"l2_cache":       cpuInfo.Cache.L2,
		"l3_cache":       cpuInfo.Cache.L3,
		"usage":          percent[0],
		"frequency":      info[0].Mhz,
		"has_avx2":       cpuInfo.Has(cpuid.AVX2),
	}
}

/*
	result := ""
	result += fmt.Sprintf("===== CPU Information =====\n")
	result += fmt.Sprintf("Vendor: %s\n", info[0].VendorID)

	result += fmt.Sprintf("\n===== Cores & Threads =====\n")
	result += fmt.Sprintf("Physical Cores: %d\n", physical)
	result += fmt.Sprintf("Logical Cores: %d\n", logical)

	result += fmt.Sprintf("\n===== Cache =====\n")

	result += fmt.Sprintf("\n===== Performance =====\n")
	result += fmt.Sprintf("CPU Usage: %.2f%%\n", percent[0])
	result += fmt.Sprintf("Frequency: %.2f MHz\n", info[0].Mhz)

	return result

}
*/
