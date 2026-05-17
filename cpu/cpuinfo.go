package cpuinfo1

import (
	"fmt"
	"math"
	"time"

	"github.com/klauspost/cpuid/v2"
	"github.com/shirou/gopsutil/v3/cpu"
)

func CPUdata() map[string]interface{} {
	// gopsutil
	info, _ := cpu.Info()
	physical, _ := cpu.Counts(false)
	logical, _ := cpu.Counts(true)
	//times, _ := cpu.Times(true)

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
	/*
		for _, t := range times {

			fmt.Println("CPU:", t.CPU)

			fmt.Println("User:", t.User)
			fmt.Println("System:", t.System)
			fmt.Println("Idle:", t.Idle)
			fmt.Println("Nice:", t.Nice)
			fmt.Println("Iowait:", t.Iowait)
			fmt.Println("Irq:", t.Irq)
			fmt.Println("Softirq:", t.Softirq)
			fmt.Println("Steal:", t.Steal)
			fmt.Println("Guest:", t.Guest)
			fmt.Println("GuestNice:", t.GuestNice)

			fmt.Println()
		}
	*/
	// cpuid
	cpuInfo := cpuid.CPU

	c1d := cpuInfo.Cache.L1D
	c1i := cpuInfo.Cache.L1I
	c2 := cpuInfo.Cache.L2
	c3 := cpuInfo.Cache.L3

	c1d, xc1d := processValue(c1d)
	c1i, xc1i := processValue(c1i)
	c2, xc2 := processValue(c2)
	c3, xc3 := processValue(c3)

	//fmt.Printf("c1d = %d %s\n", c1d, xc1d)
	//fmt.Printf("c1i = %d %s\n", c1i, xc1i)
	//fmt.Printf("c2 = %d %s\n", c2, xc2)
	//fmt.Printf("c3 = %d %s\n", c3, xc3)

	var cache string
	cache += fmt.Sprintf("L1d : %d %s\n", c1d, xc1d)
	cache += fmt.Sprintf("L1i : %d %s\n", c1i, xc1i)
	cache += fmt.Sprintf("L2 : %d %s\n", c2, xc2)
	cache += fmt.Sprintf("L3 : %d %s\n", c3, xc3)

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
		"cache": cache,
		//"BrandName":          cpuInfo.BrandName, //ชื่อ cpu
		//"l1d_cache": cpuInfo.Cache.L1D,
		//"l1i_cache": cpuInfo.Cache.L1I,
		//"l2_cache":  cpuInfo.Cache.L2,
		//"l3_cache":  cpuInfo.Cache.L3,
		//"has_avx2": cpuInfo.Has(cpuid.AVX2),

	}
}

// ============================================================================
// monitor
// ============================================================================
type CPUDatast struct {
	UsageTotal   float64   // CPU usage รวม
	UsagePerCore []float64 // CPU usage ต่อ core
	Times        []cpu.TimesStat
	//////////////////////
	/*	CpuName        string
		UserTimes      []float64 // ค่า User ของแต่ละ CPU
		SystemTimes    []float64 // ค่า System ของแต่ละ CPU
		IdleTimes      []float64 // ค่า Idle ของแต่ละ CPU
		NiceTimes      []float64
		IowaitTimes    []float64
		IrqTimes       []float64
		SoftirqTimes   []float64
		StealTimes     []float64
		GuestTimes     []float64
		GuestNiceTimes []float64
	*/

}

type CPUMonitor struct {
	ticker   *time.Ticker
	callback func(CPUDatast)
}

// สร้าง instance ใหม่
func NewCPUMonitor(interval time.Duration, callback func(CPUDatast)) *CPUMonitor {
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
			//cpu.Times()
			times, _ := cpu.Times(true)

			// แยกเฉพาะค่า float64
			//cpuData := CPUDatast{}

			for _, d := range times {
				/*		//////////////////
						totalSeconds := 9425
						hours := totalSeconds / 3600
						remainingSeconds := totalSeconds % 3600
						minutes := remainingSeconds / 60
						seconds := remainingSeconds % 60
						fmt.Printf("%d ชั่วโมง %d นาที %d วินาที\n", hours, minutes, seconds)
						fmt.Printf("%02d:%02d:%02d\n", hours, minutes, seconds)
						////////////       */
				//แปลงเป็น int
				//cpuData.UserTimes = append(cpuData.UserTimes, d.User)
				//cpuData.SystemTimes = append(cpuData.SystemTimes, d.System)
				//cpuData.IdleTimes = append(cpuData.IdleTimes, d.Idle)

				//var x int
				var timesLabel string

				nCPU := d.CPU
				tUser := int(math.Round(d.User))

				thUser, tmUser, tsUser := processTimeS(tUser)

				timesLabel += fmt.Sprintf("core [ %s ] # %d ชั่วโมง %d นาที %d วินาที\n", nCPU, thUser, tmUser, tsUser)

				timesLabel += fmt.Sprintf("core [ %s ] # %d ชั่วโมง %d นาที %d วินาที\n", nCPU, thUser, tmUser, tsUser)

				/*
					//fmt.Println("CPU:", t.CPU)
					fmt.Println("USER:", t.User)
						fmt.Println("System:", t.System)
						fmt.Println("Idle:", t.Idle)
						fmt.Println("Nice:", t.Nice)
						fmt.Println("Iowait:", t.Iowait)
						fmt.Println("Irq:", t.Irq)
						fmt.Println("Softirq:", t.Softirq)
						fmt.Println("Steal:", t.Steal)
						fmt.Println("Guest:", t.Guest)
						fmt.Println("GuestNice:", t.GuestNice)

						fmt.Println()
				*/
			}

			if len(percentTotal) > 0 {
				data := CPUDatast{
					UsageTotal:   percentTotal[0],
					UsagePerCore: percentPerCore,
					Times:        times,
					//UserTimes:    cpuData.UserTimes,
				}
				m.callback(data)
			}
		}
	}()
}

// ============================================================================
// cache
// ============================================================================
// ฟังก์ชันประมวลผลค่าด้วย switch case
func processValue(value int) (int, string) {
	// ตัวอักษร flag ที่สัมผัส
	var x string = ""
	// ตรวจสอบเงื่อนไข
	switch {
	case value >= 1099511627776:
		value = value / 1099511627776
		x = "TB"
		//fmt.Printf("%d %s\n", value, v)

	case value >= 1073741824:
		value = value / 1073741824
		x = "GB"
		//fmt.Printf("%d %s\n", value, v)

	case value >= 1048576:
		value = value / 1048576
		x = "MB"
		//fmt.Printf("%d %s\n", value, v)

	case value >= 1000:
		value = value / 1024
		x = "KB"
		//fmt.Printf("%d %s\n", value, v)

	default:
		x = "B"
		//fmt.Printf("%d %s\n", value, v)

	}
	return value, x
}

// ============================================================================
// เวลา
// ============================================================================
var hours int
var remainingSeconds int
var minutes int
var seconds int

func processTimeS(value int) (int, int, int) {
	hours = value / 3600            // หาชั่วโมง และเศษวินาทีที่เหลือ
	remainingSeconds = value % 3600 // (int หาร int จะเป็นการหารไม่เอาเศษโดยอัตโนมัติ)
	minutes = remainingSeconds / 60 //  นำเศษที่เหลือมาหาหน่วยนาที และวินาทีสุดท้าย
	seconds = remainingSeconds % 60
	return hours, minutes, seconds
}

// ============================================================================
// SECTION_NAME
// ============================================================================
