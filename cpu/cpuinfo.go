package cpuinfo1

import (
	"fmt"
	"time"

	"github.com/klauspost/cpuid/v2"
	"github.com/shirou/gopsutil/v3/cpu"
)

func CPUdata() map[string]interface{} {
	// gopsutil
	info, _ := cpu.Info()
	physical, _ := cpu.Counts(false)
	logical, _ := cpu.Counts(true)

	flagsStr := ""
	for i, flag := range info[0].Flags {
		flagsStr += flag
		if (i+1)%6 == 0 { // ทีละ 6 flags ต่อบรรทัด
			flagsStr += "\n"
		} else {
			flagsStr += " "
		}
	}
	var Hyperthreading string
	Hyperthreading += fmt.Sprintf("Hyperthreading: [ %v ]", logical > physical)

	var cpuThreadCoreSocketresult string
	for i, cpu := range info {
		cpuThreadCoreSocketresult += fmt.Sprintf("Thread [%d] : Core [%s] : Socket [%s]\n",
			i, cpu.CoreID, cpu.PhysicalID)
	}

	// cpuid
	cpuInfo := cpuid.CPU

	p := cpuInfo.Cache.L3
	fmt.Printf("\np = %d\n", p)
	processValue("p", &p)

	//y:= 9
	//u:= 101
	//i:= 40
	//o:= 85
	//p:= 23

	// ประกาศตัวแปรแยกเก็บตัวอักษร V และ G
	/*
		var B string = ""
		var KB string = ""
		var MB string = ""
		var GB string = ""
		var TB string = ""*/
	/*
	   // ฟังก์ชันประมวลผลสัญญาณตามเงื่อนไข
	   	// ประมวลผล p
	   	fmt.Printf("\np = %d\n", p)

	   	if p > 1000000 {
	   		fmt.Println("  → น้อยกว่า 100 ✓ คูณด้วย 2")
	   		p = p / 1000000
	   		pV = "V"
	   		fmt.Printf("  → ผลลัพธ์: %d, V: %s\n", p, pV)
	   		if p < 70 {
	   			fmt.Println("  → น้อยกว่า 70 ✓ บวก 3")
	   			p = p + 3
	   			pG = "G"
	   			fmt.Printf("  → ผลลัพธ์สุดท้าย: %d, V: %s, G: %s\n", p, pV, pG)
	   		} else {
	   			fmt.Println("  → มากกว่าหรือเท่ากับ 70 → ไม่บวก")
	   		}
	   	} else {
	   		fmt.Println("  → มากกว่าหรือเท่ากับ 100 → ไม่ทำอะไร")
	   		fmt.Printf("  → ผลลัพธ์: %d, V: (ไม่มี), G: (ไม่มี)\n", p)
	   	}

	   	fmt.Println("\n════════════════════════════════════")
	   	fmt.Println("     ผลลัพธ์สุดท้าย")
	   	fmt.Println("════════════════════════════════════")
	   	fmt.Printf("p = %d | V: %s | G: %s\n", t, tV, tG)
	   }
	*/
	///////////////////////////////////////////////////////////////////////

	//if p > 1048576 {

	return map[string]interface{}{
		// gopsutil
		"modelName":                 info[0].ModelName, //ชื่อ cpu
		"vendor":                    info[0].VendorID,
		"physical_cores":            physical,
		"logical_cores":             logical,
		"frequency":                 info[0].Mhz / 1000,
		"family":                    info[0].Family,
		"modelid":                   info[0].Model,
		"steppingversion":           info[0].Stepping,
		"cacheSizeMB":               info[0].CacheSize / 1024,
		"flagsStr":                  flagsStr,
		"microcodeVersion":          info[0].Microcode,
		"cpuThreadCoreSocketresult": cpuThreadCoreSocketresult,
		"Hyperthreading":            Hyperthreading,

		//cpuid
		//"BrandName":          cpuInfo.BrandName, //ชื่อ cpu
		"l1d_cache": cpuInfo.Cache.L1D / 1000,
		"l1i_cache": cpuInfo.Cache.L1I / 1000,
		"l2_cache":  cpuInfo.Cache.L2 / 1000,
		"l3_cache":  cpuInfo.Cache.L3 / 1000,
		"l3_cache1": cpuInfo.Cache.L3 / 1000,
		//"l3_test":   x,
		//"has_avx2": cpuInfo.Has(cpuid.AVX2),
	}
}

// ============================================================================
// monitor
// ============================================================================
type CPUData struct {
	UsageTotal   float64   // CPU usage รวม
	UsagePerCore []float64 // CPU usage ต่อ core
}

type CPUMonitor struct {
	ticker   *time.Ticker
	callback func(CPUData)
}

// สร้าง instance ใหม่
func NewCPUMonitor(interval time.Duration, callback func(CPUData)) *CPUMonitor {
	return &CPUMonitor{
		ticker:   time.NewTicker(interval),
		callback: callback,
	}
}

// เริ่ม monitoring
func (m *CPUMonitor) Start() {
	go func() {
		for range m.ticker.C {
			// ดึง CPU usage รวม
			percentTotal, _ := cpu.Percent(100*time.Millisecond, false)

			// ดึง CPU usage ต่อ core
			percentPerCore, _ := cpu.Percent(100*time.Millisecond, true)

			if len(percentTotal) > 0 {
				data := CPUData{
					UsageTotal:   percentTotal[0],
					UsagePerCore: percentPerCore,
				}
				m.callback(data)
			}
		}
	}()
}

// ฟังก์ชันประมวลผลค่าด้วย switch case
func processValue(name string, value *int) {
	// ตัวอักษร flag ที่สัมผัส
	var v string = ""
	// ตรวจสอบเงื่อนไข
	switch {
	case *value < 1000:
		*value = *value //
		v = "B"
		fmt.Printf("%d %s\n", *value, v)

	case *value >= 1024:
		*value = *value / 1024 //
		v = "KB"
		fmt.Printf("%d %s\n", *value, v)

	case *value >= 1048576:
		*value = *value / 1048576 //
		v = "MB"
		fmt.Printf("%d %s\n", *value, v)

	case *value < 100:
		fmt.Println("  Case 2: น้อยกว่า 100 แต่ >= 70 ✓")
		*value = *value * 2 // คูณ 2
		v = "V"
		fmt.Printf("  → คูณ 2: ได้ %d\n", *value)
		fmt.Printf("  → Flag V: %s, Flag G: (ไม่มี)\n", v)

	// Case 3: มากกว่าหรือเท่ากับ 100
	default:
		fmt.Println("  Case 3: มากกว่าหรือเท่ากับ 100 → ไม่ทำอะไร")
		fmt.Printf("  → ค่าเดิม: %d (ไม่เปลี่ยน)\n", *value)
		fmt.Printf("  → Flag V: (ไม่มี), Flag G: (ไม่มี)\n")
	}
}
