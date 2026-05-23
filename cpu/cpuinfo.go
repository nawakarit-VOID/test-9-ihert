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

	// ============================================================================
	// Overview
	// ============================================================================
	var overview string // gopsutil
	overview += fmt.Sprintf("CPU : %s\n", info[0].ModelName)
	overview += fmt.Sprintf("Vendor : %s\n", info[0].VendorID)
	overview += fmt.Sprintf("Cores : %d\n", physical)
	overview += fmt.Sprintf("Thread : %d\n", logical)
	overview += fmt.Sprintf("FreqMax : %.2f GHz\n", info[0].Mhz/1000)
	overview += fmt.Sprintf("Family : %s\n", info[0].Family)
	overview += fmt.Sprintf("Modelid : %s\n", info[0].Model)
	overview += fmt.Sprintf("Stepping : %d\n", info[0].Stepping)
	overview += fmt.Sprintf("Cache : %d MB\n", info[0].CacheSize/1024)
	overview += fmt.Sprintf("Microcode : %s\n", info[0].Microcode)

	// ============================================================================
	// Detail
	// ============================================================================
	Hyperthreading := fmt.Sprintf("Hyperthreading: [ %v ]\n", logical > physical)

	var cpuThreadCoreSocketresult string //จำนวน thread
	cpuThreadCoreSocketresult += ("\n[  Thread  ] : [ Core ] : [ Socket ]\n")
	for i, cpu := range info {
		cpuThreadCoreSocketresult += fmt.Sprintf("Thread [%d] : Core [%s] : Socket [%s]\n",
			i, cpu.CoreID, cpu.PhysicalID)
	}

	// cpuid
	cpuInfo := cpuid.CPU
	//c1d := cpuInfo.Cache.L1D
	//c1i := cpuInfo.Cache.L1I
	//c2 := cpuInfo.Cache.L2
	//c3 := cpuInfo.Cache.L3

	c1d, xc1d := processValue(cpuInfo.Cache.L1D)
	c1i, xc1i := processValue(cpuInfo.Cache.L1I)
	c2, xc2 := processValue(cpuInfo.Cache.L2)
	c3, xc3 := processValue(cpuInfo.Cache.L3)

	var cache string //cpuid
	cache += "\n[ Cache ]\n"
	cache += fmt.Sprintf("L1d : %d %s\n", c1d, xc1d)
	cache += fmt.Sprintf("L1i : %d %s\n", c1i, xc1i)
	cache += fmt.Sprintf("L2 : %d %s\n", c2, xc2)
	cache += fmt.Sprintf("L3 : %d %s\n", c3, xc3)

	//"BrandName":          cpuInfo.BrandName, //ชื่อ cpu
	//"l1d_cache": cpuInfo.Cache.L1D,
	//"l1i_cache": cpuInfo.Cache.L1I,
	//"l2_cache":  cpuInfo.Cache.L2,
	//"l3_cache":  cpuInfo.Cache.L3,
	//"has_avx2": cpuInfo.Has(cpuid.AVX2),

	var detail string //
	detail += Hyperthreading
	detail += cpuThreadCoreSocketresult
	detail += cache //cpuid

	// ============================================================================
	// Flags Feature
	// ============================================================================
	var flagsfeature string
	for i, flag := range info[0].Flags {
		flagsfeature += flag
		if (i+1)%6 == 0 { // ทีละ 6 flags ต่อบรรทัด
			flagsfeature += "\n"
		} else {
			flagsfeature += " "
		}
	}

	return map[string]interface{}{
		// gopsutil
		"Overview":     overview,
		"Detail":       detail,
		"FlagsFeature": flagsfeature,
	}

}

// ============================================================================
// monitor
// ============================================================================
type StCPUData struct {
	Usage      string //
	Timesusage string

	UsagePerCore   []float64 // CPU usage ต่อ core
	PercentPerCore string
	Times          []cpu.TimesStat
	//////////////////////
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
			// แสดง usage ต่อ core
			var PerCore string
			for i, pc := range percentPerCore {
				PerCore += fmt.Sprintf("Core [ %d ] : %.1f%%\n", i, pc)
			}
			//รวม usage
			var usage string
			usage += fmt.Sprintf("Usage Avg : %.2f\n\n", percentTotal[0])
			usage += fmt.Sprintf("%s\n", PerCore)

			//cpu.Times()
			times, _ := cpu.Times(true)

			var timesSec string
			timesSec += "[ ข้อมูลดิบ ]\n\n"
			var timesHms string
			timesHms += "\n[ แปลงเป็นเวลาสากล ]\n"
			var timesTotalAvg string
			timesTotalAvg += "\n[ เฉลี่ย ]\n\n"
			meanLabel := "\n[ ความหมาย ]\n\n" + `User : โปรแกรมของผู้ใช้
System : ระบบ
Idle : ไม่ได้ทำอะไร
Nice : เวลาที่ใช้กับ process ที่ถูกปรับ priority (nice)
Iowait : เวลาที่ CPU รอ I/O เช่น disk หรือ network
Irq : เวลาที่ใช้จัดการ Hardware ที่ขัดจังหวะ
Softirq : เวลาที่ใช้จัดการ Software ที่ขัดจังหวะ
Steal : เวลาที่ VM ถูก hypervisor แย่ง CPU ไป
Guest : เวลาที่ CPU ใช้งาน guest virtual machine
GuestNice : เวลาที่ guest VM ใช้งานแบบ nice priority`

			for _, d := range times {

				nCPU := d.CPU
				//วินาที *ดิบ
				timesSec += fmt.Sprintf(
					"CPU: [ %s ] | User: %.2f s | System: %.2f s | Idle: %.2f s | Nice: %.2f s | Iowait: %.2f s | Irq %.2f s | Softirq %.2f s | Steal %.2f s | Guest %.2f s | GuestNice %.2f s\n",
					nCPU, d.User, d.System, d.Idle, d.Nice, d.Iowait, d.Irq, d.Softirq, d.Steal, d.Guest, d.GuestNice)

				//แปลงเป็นเวลาสากล
				thUser, tmUser, tsUser := processTimeS(d.User)
				//fmt.Println(d.User)
				thSystem, tmSystem, tsSystem := processTimeS(d.System)
				thIdle, tmIdle, tsIdle := processTimeS(d.Idle)
				thNice, tmNice, tsNice := processTimeS(d.Nice)
				thIowait, tmIowait, tsIowait := processTimeS(d.Iowait)
				thIrq, tmIrq, tsIrq := processTimeS(d.Irq)
				thSoftirq, tmSoftirq, tsSoftirq := processTimeS(d.Softirq)
				thSteal, tmSteal, tsSteal := processTimeS(d.Steal)
				thGuest, tmGuest, tsGuest := processTimeS(d.Guest)
				thGuestNice, tmGuestNice, tsGuestNice := processTimeS(d.GuestNice)
				//จัดเรียง

				timesHms += fmt.Sprintf("\n[ %s ]\nUser # [ %d : %d : %d ]\n", nCPU, thUser, tmUser, tsUser)
				timesHms += fmt.Sprintf("System # [ %d : %d : %d ]\n", thSystem, tmSystem, tsSystem)
				timesHms += fmt.Sprintf("Idle # [ %d : %d : %d ]\n", thIdle, tmIdle, tsIdle)
				timesHms += fmt.Sprintf("Nice # [ %d : %d : %d ]\n", thNice, tmNice, tsNice)
				timesHms += fmt.Sprintf("Iowait # [ %d : %d : %d ]\n", thIowait, tmIowait, tsIowait)
				timesHms += fmt.Sprintf("Irq # [ %d : %d : %d ]\n", thIrq, tmIrq, tsIrq)
				timesHms += fmt.Sprintf("Softirq # [ %d : %d : %d ]\n", thSoftirq, tmSoftirq, tsSoftirq)
				timesHms += fmt.Sprintf("Steal # [ %d : %d : %d ]\n", thSteal, tmSteal, tsSteal)
				timesHms += fmt.Sprintf("Guest # [ %d : %d : %d ]\n", thGuest, tmGuest, tsGuest)
				timesHms += fmt.Sprintf("GuestNice # [ %d : %d : %d ]\n", thGuestNice, tmGuestNice, tsGuestNice)
				//fmt.Print(timesHms)

				//AVG//เวลาโดยเฉลี่ย
				thAvgscores := []int{thUser, thSystem, thIdle, thNice, thIowait, thIrq, thSoftirq, thSteal, thGuest, thGuestNice}
				tmAvgscores := []int{tmUser, tmSystem, tmIdle, tmNice, tmIowait, tmIrq, tmSoftirq, tmSteal, tmGuest, tmGuestNice}
				tsAvgscores := []int{tsUser, tsSystem, tsIdle, tsNice, tsIowait, tsIrq, tsSoftirq, tsSteal, tsGuest, tsGuestNice}

				//	thidleAvgscores := []int{}
				//	tmidleAvgscores := []int{}
				//	tsidleAvgscores := []int{}

				//thsumAvg := 0
				//tmsumAvg := 0
				//tssumAvg := 0

				//thvalidCount := 0 // สร้างตัวแปรมาไว้นับเฉพาะคนที่มีคะแนน
				//tmvalidCount := 0
				//tsvalidCount := 0

				/*thsumAvg, thvalidCount, */
				//			 sum, count := numSumAndCount(d.User)
				//thavg := numSumAndCount(d.User)
				//tmavg := numSumAndCount(tmAvgscores)
				//tsavg := numSumAndCount(tsAvgscores)
				/*
					for _, thscore := range thAvgscores {
						thsumAvg += thscore
						if thscore > 0 { // ถ้ามากกว่า 0 ให้นับเพิ่ม
							thvalidCount++
						}
					}

					for _, tmscore := range tmAvgscores {
						tmsumAvg += tmscore
						if tmscore > 0 { // ถ้ามากกว่า 0 ให้นับเพิ่ม
							tmvalidCount++
						}
					}

					for _, tsscore := range tsAvgscores {
						tssumAvg += tsscore
						if tsscore > 0 { // ถ้ามากกว่า 0 ให้นับเพิ่ม
							tsvalidCount++
						}
					}
				*/
				//thsum, thcount := numSumAndCount(thAvgscores)
				thsum, thavg := numSumAndCount(thAvgscores) //ok
				tmsum, tmavg := numSumAndCount(tmAvgscores)
				tssum, tsavg := numSumAndCount(tsAvgscores)

				thAvg := Avg(thsum, thavg)
				tmAvg := Avg(tmsum, tmavg)
				tsAvg := Avg(tssum, tsavg)

				/*		var thavg float64
						if count > 0 {
							thavg = float64(thsum) / float64(thcount)
						}
				*/
				/*
					// หารด้วยจำนวนเฉพาะคนที่มีคะแนน (ไม่รวมเลข 0)
					// ป้องกันเคสที่ validtCount เป็น 0 ด้วยการเช็คเงื่อนไขก่อนหาร
					var thavg float64
					if thvalidCount > 0 {
						thavg = float64(thsumAvg) / float64(thvalidCount)
					}

					var tmavg float64
					if tmvalidCount > 0 {
						tmavg = float64(tmsumAvg) / float64(tmvalidCount)
					}
					var tsavg float64
					if tsvalidCount > 0 {
						tsavg = float64(tssumAvg) / float64(tsvalidCount)
					}
				*/

				//timesTotalAvg += fmt.Sprintf("[ %s ] เฉลี่ย [ %.f : %.f : %.f ]\n", nCPU, thavg /*tmavg, tsavg*/)
				timesTotalAvg += fmt.Sprintf("[ %s ] เฉลี่ย [ %.f : %.f : %.f ]\n", nCPU, thAvg, tmAvg, tsAvg)
				//fmt.Print(timesTotalAvg)

			}
			//รวม timesusage
			var timesusage string
			timesusage += timesSec
			timesusage += timesHms
			timesusage += timesTotalAvg
			timesusage += meanLabel
			var x string
			x += "sssss"
			if len(percentTotal) > 0 {

				data := StCPUData{
					Usage: x,
					//Usage:      usage,
					Timesusage: timesusage,
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

var sum1 int = 0
var count1 int = 0
var avg1 float64
var valueRang int
var i int

func processTimeS(value float64) (int, int, int) {

	hours = int(value) / 3600            // หาชั่วโมง  (int หาร int จะเป็นการหารไม่เอาเศษโดยอัตโนมัติ) *หารไม่เอาเศษ
	remainingSeconds = int(value) % 3600 //หาเศษวินาทีที่เหลือ *% หารเพื่อเอาเศษ
	minutes = remainingSeconds / 60      //  นำเศษที่เหลือมาหาหน่วยนาที *แบบไม่เอาเศษและวินาทีสุดท้าย
	seconds = remainingSeconds % 60      //และวินาทีสุดท้าย *หารเอาเศษ

	//println(value)
	value += value
	//println(value)

	for valueRang := range int(value) {
		valueRang += valueRang

	}

	return hours, minutes, seconds
}

// ============================================================================
// SECTION_NAME
// ============================================================================
var sum int = 0
var count int = 0

func numSumAndCount(value []int) (int, int) {

	for _, x := range value {
		sum += x
		if x > 0 { // ถ้ามากกว่า 0 ให้นับเพิ่ม
			count++
		}
	}

	return sum, count
}

var avg float64

func Avg(valuex int, valuey int) float64 {
	for {
		if valuey > 0 {
			avg = float64(valuex) / float64(valuey)
			return avg
		}
	}
}

/*
	// หารด้วยจำนวนเฉพาะคนที่มีคะแนน (ไม่รวมเลข 0)
	// ป้องกันเคสที่ validtCount เป็น 0 ด้วยการเช็คเงื่อนไขก่อนหาร
	if count > 0 {
		avg = float64(sum) / float64(count)
	}
*/
