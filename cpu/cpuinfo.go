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
type StCPUData struct {
	UsageTotal   float64   // CPU usage รวม
	UsagePerCore []float64 // CPU usage ต่อ core
	Times        []cpu.TimesStat
	//////////////////////
	//CpuName        string
	//UserTimes      []float64 // ค่า User ของแต่ละ CPU
	//SystemTimes    []float64 // ค่า System ของแต่ละ CPU
	TimesAVGLabel string // ค่า Idle ของแต่ละ CPU
	TimesLabel    string

	//NiceTimes      []float64
	//IowaitTimes    []float64
	//IrqTimes       []float64
	//SoftirqTimes   []float64
	//StealTimes     []float64
	//GuestTimes     []float64
	//GuestNiceTimes []float64

}

type CPUMonitor struct {
	ticker   *time.Ticker
	callback func(StCPUData)
}

// สร้าง instance ใหม่
func NewCPUMonitor(interval time.Duration, callback func(StCPUData)) *CPUMonitor {
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

			//cpuData := StCPUData{}
			var timesLabel string
			var timesAVGLabel string
			var thAvg int
			var tmAvg int
			var tsAvg int

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

				nCPU := d.CPU
				/*				//แปลงเป็นตัวเลข
								tUser := int(math.Round(d.User))
								tSystem := int(math.Round(d.System))
								tIdle := int(math.Round(d.Idle))
								tNice := int(math.Round(d.Nice))
								tIowait := int(math.Round(d.Iowait))
								tIrq := int(math.Round(d.Irq))
								tSoftirq := int(math.Round(d.Softirq))
								tSteal := int(math.Round(d.Steal))
								tGuest := int(math.Round(d.Guest))
								tGuestNice := int(math.Round(d.GuestNice))

								//เอาไปคิดเวลา
								thUser, tmUser, tsUser := processTimeS(tUser)
								thSystem, tmSystem, tsSystem := processTimeS(tSystem)
								thIdle, tmIdle, tsIdle := processTimeS(tIdle)
								thNice, tmNice, tsNice := processTimeS(tNice)
								thIowait, tmIowait, tsIowait := processTimeS(tIowait)
								thIrq, tmIrq, tsIrq := processTimeS(tIrq)
								thSoftirq, tmSoftirq, tsSoftirq := processTimeS(tSoftirq)
								thSteal, tmSteal, tsSteal := processTimeS(tSteal)
								thGuest, tmGuest, tsGuest := processTimeS(tGuest)
								thGuestNice, tmGuestNice, tsGuestNice := processTimeS(tGuestNice)
				*/
				thUser, tmUser, tsUser := processTimeS(d.User)
				thSystem, tmSystem, tsSystem := processTimeS(d.System)
				thIdle, tmIdle, tsIdle := processTimeS(d.Idle)
				thNice, tmNice, tsNice := processTimeS(d.Nice)
				thIowait, tmIowait, tsIowait := processTimeS(d.Iowait)
				thIrq, tmIrq, tsIrq := processTimeS(d.Irq)
				thSoftirq, tmSoftirq, tsSoftirq := processTimeS(d.Softirq)
				thSteal, tmSteal, tsSteal := processTimeS(d.Steal)
				thGuest, tmGuest, tsGuest := processTimeS(d.Guest)
				thGuestNice, tmGuestNice, tsGuestNice := processTimeS(d.GuestNice)

				timesLabel += fmt.Sprintf("core [ %s ]\n	User # %d ชั่วโมง %d นาที %d วินาที\n", nCPU, thUser, tmUser, tsUser)
				timesLabel += fmt.Sprintf("	System # %d ชั่วโมง %d นาที %d วินาที\n", thSystem, tmSystem, tsSystem)
				timesLabel += fmt.Sprintf("	Idle # %d ชั่วโมง %d นาที %d วินาที\n", thIdle, tmIdle, tsIdle)
				timesLabel += fmt.Sprintf("	Nice # %d ชั่วโมง %d นาที %d วินาที\n", thNice, tmNice, tsNice)
				timesLabel += fmt.Sprintf("	Iowait # %d ชั่วโมง %d นาที %d วินาที\n", thIowait, tmIowait, tsIowait)
				timesLabel += fmt.Sprintf("	Irq # %d ชั่วโมง %d นาที %d วินาที\n", thIrq, tmIrq, tsIrq)
				timesLabel += fmt.Sprintf("	Softirq # %d ชั่วโมง %d นาที %d วินาที\n", thSoftirq, tmSoftirq, tsSoftirq)
				timesLabel += fmt.Sprintf("	Steal # %d ชั่วโมง %d นาที %d วินาที\n", thSteal, tmSteal, tsSteal)
				timesLabel += fmt.Sprintf("	Guest # %d ชั่วโมง %d นาที %d วินาที\n", thGuest, tmGuest, tsGuest)
				timesLabel += fmt.Sprintf("	GuestNice# %d ชั่วโมง %d นาที %d วินาที\n", thGuestNice, tmGuestNice, tsGuestNice)
				fmt.Print(timesLabel)

				//AVG
				thAvgscores := []int{thUser, thSystem, thIdle, thNice, thIowait, thIrq, thSoftirq, thSteal, thGuest, thGuestNice}
				thsumAvg := 0
				thvalidCount := 0 // สร้างตัวแปรมาไว้นับเฉพาะคนที่มีคะแนน

				for _, thscore := range thAvgscores {
					thsumAvg += thscore

					// ถ้าคะแนนมากกว่า 0 ให้นับเพิ่ม
					if thscore > 0 {
						thvalidCount++
					}
				}
				// หารด้วยจำนวนเฉพาะคนที่มีคะแนน (ไม่รวมเลข 0)
				// ป้องกันเคสที่ validtCount เป็น 0 ด้วยการเช็คเงื่อนไขก่อนหาร
				var avgWithoutZeros float64
				if thvalidCount > 0 {
					avgWithoutZeros = float64(thsumAvg) / float64(thvalidCount)
				}
				fmt.Printf("ผลรวมคะแนน: %d\n", thsumAvg)
				fmt.Printf("2. ค่าเฉลี่ยแบบนับเฉพาะคนมีคะแนน (หาร %d): %.2f\n", thvalidCount, avgWithoutZeros)

				/*
					tmAvg = (tmUser + tmSystem + tmIdle + tmNice + tmIowait + tmIrq + tmSoftirq + tmSteal + tmGuest + tmGuestNice) / 10
					tsAvg = (tsUser + tsSystem + tsIdle + tsNice + tsIowait + tsIrq + tsSoftirq + tsSteal + tsGuest + tsGuestNice) / 10
				*/
				//fmt.Println( thAvg, "ชั่วโมง", tmAvg, "นาที", tsAvg, "วินาที")
				//timesAVGLabel += fmt.Sprintf("Core [ AVG ]\n	User # %d ชั่วโมง %d นาที %d วินาที\n", thAvg, tmAvg, tsAvg)
				timesAVGLabel += fmt.Sprintf("Core [ %s AVG ]\n	User # %d ชั่วโมง %d นาที %d วินาที\n", nCPU, thAvg, tmAvg, tsAvg)
				//cpuData.TimesAVGLabel = append(cpuData.TimesAVGLabel, timesAVGLabel)
			}

			if len(percentTotal) > 0 {
				data := StCPUData{
					UsageTotal:   percentTotal[0],
					UsagePerCore: percentPerCore,
					Times:        times,
					//
					TimesLabel:    timesLabel,
					TimesAVGLabel: timesAVGLabel,
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

func processTimeS(value float64) (int, int, int) {

	hours = int(value) / 3600            // หาชั่วโมง  (int หาร int จะเป็นการหารไม่เอาเศษโดยอัตโนมัติ) *หารไม่เอาเศษ
	remainingSeconds = int(value) % 3600 //หาเศษวินาทีที่เหลือ *% หารเพื่อเอาเศษ
	minutes = remainingSeconds / 60      //  นำเศษที่เหลือมาหาหน่วยนาที *แบบไม่เอาเศษและวินาทีสุดท้าย
	seconds = remainingSeconds % 60      //และวินาทีสุดท้าย *หารเอาเศษ

	return hours, minutes, seconds
}

// ============================================================================
// SECTION_NAME
// ============================================================================
