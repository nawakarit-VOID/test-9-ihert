package cpuinfo

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
	overview += fmt.Sprintf("\nCPU : %s\n", info[0].ModelName)
	overview += fmt.Sprintf("Vendor : %s\n", info[0].VendorID)
	overview += fmt.Sprintf("Cores : %d\n", physical)
	overview += fmt.Sprintf("Thread : %d\n", logical)
	overview += fmt.Sprintf("FreqMax : %.2f GHz\n", info[0].Mhz/1000)
	overview += fmt.Sprintf("Family : %s\n", info[0].Family)
	overview += fmt.Sprintf("Modelid : %s\n", info[0].Model)
	overview += fmt.Sprintf("Stepping : %d\n", info[0].Stepping)
	overview += fmt.Sprintf("Cache : %d MB\n", info[0].CacheSize/1024)
	overview += fmt.Sprintf("Microcode : %s\n", info[0].Microcode)

	//แยกส่วน
	var modelName string
	var vendorID string
	var core string
	var thread string
	var freqMax string
	var family string
	var modelid string
	var stepping string
	var cachet string
	var microcode string

	modelName += fmt.Sprintf("\nCPU : %s\n", info[0].ModelName)
	vendorID += fmt.Sprintf("Vendor : %s\n", info[0].VendorID)
	core += fmt.Sprintf("Cores : %d\n", physical)
	thread += fmt.Sprintf("Thread : %d\n", logical)
	freqMax += fmt.Sprintf("FreqMax : %.2f GHz\n", info[0].Mhz/1000)
	family += fmt.Sprintf("Family : %s\n", info[0].Family)
	modelid += fmt.Sprintf("Modelid : %s\n", info[0].Model)
	stepping += fmt.Sprintf("Stepping : %d\n", info[0].Stepping)
	cachet += fmt.Sprintf("Cache : %d MB\n", info[0].CacheSize/1024)
	microcode += fmt.Sprintf("Microcode : %s\n", info[0].Microcode)
	// ============================================================================
	// Detail
	// ============================================================================
	Hyperthreading := fmt.Sprintf("\nHyperthreading: [ %v ]\n", logical > physical)

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
	flagsfeature += "\n"
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

		"ModelName": modelName,
		"VendorID":  vendorID,
		"Core    ":  core,
		"Thread  ":  thread,
		"FreqMax ":  freqMax,
		"Family  ":  family,
		"Modelid ":  modelid,
		"Stepping":  stepping,
		"Cachet ":   cachet,
		"Microcode": microcode,
	}

}

/*
	: modelName,
	: vendorID,
	: core,
	: thread,
	: freqMax,
	: family,
	: modelid,
	: stepping,
	: cache,
	: microcode,
*/
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
			percentTotal, err := cpu.Percent(100*time.Millisecond, false)
			if err != nil || len(percentTotal) == 0 {
				continue
			}

			// ดึง CPU usage ต่อ core
			percentPerCore, err := cpu.Percent(100*time.Millisecond, true)
			if err != nil {
				continue
			}
			// แสดง usage ต่อ core
			var PerCore string
			for i, pc := range percentPerCore {
				PerCore += fmt.Sprintf("Core [ %d ] : %.1f%%\n", i, pc)
			}
			//จัดเรียง usage
			var usage string
			usage += "\n"
			usage += fmt.Sprintf("Usage Avg : %.2f\n\n", percentTotal[0])
			usage += fmt.Sprintf("%s\n", PerCore)

			//cpu.Times()
			times, err := cpu.Times(true)
			if err != nil || len(times) == 0 {
				continue
			}

			var timesTotalAvglabel string
			timesTotalAvglabel += "\n[ เฉลี่ย ]\n\n"
			var timesTotalAvg string

			var timesSec string
			timesSec += "\n[ ข้อมูลดิบ ]\n\n"

			var timesHms string
			timesHms += "\n[ แปลงเป็นเวลาสากล ]\n\n"

			var meaningLabel string
			meaningLabel += "\nความหมาย\n\n"
			meaningLabel += "[ User : โปรแกรมผู้ใช้ ]\n[ System : ระบบ ]\n[ Idle : ไม่ได้ทำอะไร ]\n[ Nice : เวลาปรับ priority ]\n[ Iowait : CPU รอ I/O ]\n[ Irq : Hardware ขัด ]\n[ Softirq : Software ขัดจังหวะ ]\n[ Steal : VM ถูก hyper แย่ง ]\n[ Guest : ใช้ guest virtual ]\n[ GuestNice : VM ใช้แบบ nice priority ]\n"

			var totalUser float64
			var totalSystem float64
			var totalIdle float64
			var totalNice float64
			var totalIowait float64
			var totalIrq float64
			var totalSoftirq float64
			var totalSteal float64
			var totalGuest float64
			var totalGuestNice float64

			for _, d := range times {

				totalUser += d.User
				totalSystem += d.System
				totalIdle += d.Idle //รวม idle
				totalNice += d.Nice
				totalIowait += d.Iowait
				totalIrq += d.Irq
				totalSoftirq += d.Softirq
				totalSteal += d.Steal
				totalGuest += d.Guest
				totalGuestNice += d.GuestNice

				nCPU := d.CPU
				//วินาที *ดิบ
				timesSec += fmt.Sprintf(
					"[ %s ] | User: %.2f s | System: %.2f s | Idle: %.2f s | Nice: %.2f s | Iowait: %.2f s | Irq %.2f s | Softirq %.2f s | Steal %.2f s | Guest %.2f s | GuestNice %.2f s\n",
					nCPU, d.User, d.System, d.Idle, d.Nice, d.Iowait, d.Irq, d.Softirq, d.Steal, d.Guest, d.GuestNice)

				//แปลงเป็นเวลาสากล
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

				//จัดเรียงเวลาสากล
				timesHms += fmt.Sprintf(
					"[ %s ] | User [ %d : %d : %d ] | System [ %d : %d : %d ] | Idle [ %d : %d : %d ] | Nice [ %d : %d : %d ] | Iowait [ %d : %d : %d ] | Irq [ %d : %d : %d ] | Softirq [ %d : %d : %d ] | Steal [ %d : %d : %d ] | Guest [ %d : %d : %d ] | GuestNice [ %d : %d : %d ]\n",
					nCPU, thUser, tmUser, tsUser, thSystem, tmSystem, tsSystem, thIdle, tmIdle, tsIdle, thNice, tmNice, tsNice, thIowait, tmIowait, tsIowait, thIrq, tmIrq, tsIrq, thSoftirq, tmSoftirq, tsSoftirq, thSteal, tmSteal, tsSteal, thGuest, tmGuest, tsGuest, thGuestNice, tmGuestNice, tsGuestNice)

				//AVG//เวลาโดยเฉลี่ย
				// ***แยก system กับ idle
				hUser, mUser, sUser := Avg(totalUser)
				hSystem, mSystem, sSysteme := Avg(totalSystem)
				hIdle, mIdle, sIdle := Avg(totalIdle)
				hNice, mNice, sNice := Avg(totalNice)
				hIowait, mIowait, sIowait := Avg(totalIowait)
				hIrq, mIrq, sIrq := Avg(totalIrq)
				hSoftirq, mSoftirq, sSoftirq := Avg(totalSoftirq)
				hSteal, mSteal, sSteal := Avg(totalSteal)
				hGuest, mGuest, sGuest := Avg(totalGuest)
				hGuestNice, mGuestNice, sGuestNice := Avg(totalGuestNice)
				//จัดเรียงเวลาโดยเฉลี่ย
				timesTotalAvg = fmt.Sprintf(
					"[ AVG ] | User [ %d : %d : %d ] | System [ %d : %d : %d ] | Idle [ %d : %d : %d ] | Nice [ %d : %d : %d ] | Iowait [ %d : %d : %d ] | Irq [ %d : %d : %d ] | Softirq [ %d : %d : %d ] | Steal [ %d : %d : %d ] | Guest [ %d : %d : %d ] | GuestNice [ %d : %d : %d ]\n",
					hUser, mUser, sUser, hSystem, mSystem, sSysteme, hIdle, mIdle, sIdle, hNice, mNice, sNice, hIowait, mIowait, sIowait, hIrq, mIrq, sIrq, hSoftirq, mSoftirq, sSoftirq, hSteal, mSteal, sSteal, hGuest, mGuest, sGuest, hGuestNice, mGuestNice, sGuestNice)

			}

			//จัดเรียง timesusage
			var timesusage string
			timesusage += timesTotalAvglabel
			timesusage += timesTotalAvg
			timesusage += timesSec
			timesusage += timesHms
			timesusage += meaningLabel

			//	timesusage += meanLabel

			if len(percentTotal) > 0 {

				data := StCPUData{
					Usage:      usage,
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
	case value >= 1073741824:
		value = value / 1073741824
		x = "GB"
	case value >= 1048576:
		value = value / 1048576
		x = "MB"
	case value >= 1000:
		value = value / 1024
		x = "KB"
	default:
		x = "B"
	}
	return value, x
}

// ============================================================================
// เวลา
// ============================================================================

func processTimeS(value float64) (int, int, int) {

	hours := int(value) / 3600            // หาชั่วโมง  (int หาร int จะเป็นการหารไม่เอาเศษโดยอัตโนมัติ) *หารไม่เอาเศษ
	remainingSeconds := int(value) % 3600 //หาเศษวินาทีที่เหลือ *% หารเพื่อเอาเศษ
	minutes := remainingSeconds / 60      //  นำเศษที่เหลือมาหาหน่วยนาที *แบบไม่เอาเศษและวินาทีสุดท้าย
	seconds := remainingSeconds % 60      //และวินาทีสุดท้าย *หารเอาเศษ

	return hours, minutes, seconds
}

// ============================================================================
// หาค่าเฉลี่ย
// ============================================================================
func numSumAndCount(value []int) (int, int) {
	sum := 0
	count := 0

	for _, x := range value {
		sum += x
		if x > 0 { // ถ้ามากกว่า 0 ให้นับเพิ่ม
			count++
		}
	}
	return sum, count
}

func Avg(value float64) (int, int, int) {
	physical, _ := cpu.Counts(false)
	AA1 := int(value) / int(physical)
	AA2 := float64(AA1)
	A1, A2, A3 := processTimeS(AA2)
	return A1, A2, A3
}

// ============================================================================
// รวม + เอาออก
// ============================================================================
func CpuPage() fyne.CanvasObject {
	dataCPUInfo := CPUdata()

	return container.NewVBox(
		widget.NewLabel(fmt.Sprintf("%s\n", dataCPUInfo["Overview"])), //1 แสดง cpu info
		widget.NewLabel(fmt.Sprintf("%s\n", dataCPUInfo["ModelName"])),

		widget.NewLabel("Hello"),
		//widget.NewLabel(overview),
		widget.NewButton("OK", func() {}),
	)
}

/*
	"ModelName": modelName,
		"VendorID":  vendorID,
		"Core    ":  core,
		"Thread  ":  thread,
		"FreqMax ":  freqMax,
		"Family  ":  family,
		"Modelid ":  modelid,
		"Stepping":  stepping,
		"Cachet ":   cachet,
		"Microcode": microcode,
*/
