package cpuinfo

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/klauspost/cpuid/v2"
	"github.com/shirou/gopsutil/v3/cpu"
)

func CPUdata() map[string]interface{} {
	// gopsutil
	info, _ := cpu.Info()
	//percent, _ := cpu.Percent(time.Second, false)
	physical, _ := cpu.Counts(false)
	logical, _ := cpu.Counts(true)

	usageLabel := widget.NewLabel("usaelabel...")

	flagsStr := ""
	for i, flag := range info[0].Flags {
		flagsStr += flag
		if (i+1)%6 == 0 { // ทีละ 6 flags ต่อบรรทัด
			flagsStr += "\n"
		} else {
			flagsStr += " "
		}
	}

	//cpu.Percent()
	// loop อัปเดต usage
	//
	go func() {
		for {
			percent, _ := cpu.Percent(1*time.Second, false)
			if len(percent) > 0 {
				usage := percent[0]

				fyne.Do(func() {
					usageLabel.SetText(fmt.Sprintf("CPU Avg: %.2f%%", usage))
				})
			}
		}
	}()

	// cpuid
	cpuInfo := cpuid.CPU

	return map[string]interface{}{
		// gopsutil
		"modelName":      info[0].ModelName, //ชื่อ cpu
		"vendor":         info[0].VendorID,
		"physical_cores": physical,
		"logical_cores":  logical,
		//"usage":            percent[0],
		"frequency":        info[0].Mhz / 1000,
		"family":           info[0].Family,
		"modelid":          info[0].Model,
		"steppingversion":  info[0].Stepping,
		"cacheSizeMB":      info[0].CacheSize / 1024,
		"flagsStr":         flagsStr,
		"microcodeVersion": info[0].Microcode,
		//"usage":            usage,
		//"percent":          cpu.Percent(1*time.Second, false),
		"usageLabel": usageLabel,

		//cpuid
		//"BrandName":          cpuInfo.BrandName, //ชื่อ cpu
		"l1_cache": cpuInfo.Cache.L1D,
		"l2_cache": cpuInfo.Cache.L2,
		"l3_cache": cpuInfo.Cache.L3,
		"has_avx2": cpuInfo.Has(cpuid.AVX2),
	}
}
