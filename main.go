package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	a := app.New()
	w := a.NewWindow("CPU Info")

	// Labels
	InfoLabel := widget.NewLabel("...")
	flagsLabel := widget.NewLabel("...")
	//detailLabel := widget.NewLabel("...")
	//coresthread := widget.NewLabel("...")

	// cpu.Info()
	//cpunumber := widget.NewLabel("CPUnumber: ...") //CPU - หมายเลข CPU
	//vendorid := widget.NewLabel("Vendorid: ...")                 //VendorID - ผู้ผลิต CPU
	//cpufamily := widget.NewLabel("CPUfamily: ...")               //Family - CPU family
	//modelid := widget.NewLabel("Modelid: ...")                   //Model - model id
	//steppingversion := widget.NewLabel("Stepping version: ...")  //Stepping - stepping version
	//socketid := widget.NewLabel("Socketid: ...")   //PhysicalID - socket id
	//coreid := widget.NewLabel("Coreid: ...")       //CoreID - core id
	///coresmain := widget.NewLabel("Coresmain: ...") //Cores - จำนวน core
	//modelName := widget.NewLabel("CPU: loading...")              //ModelName - ชื่อ CPU เต็ม
	//freq := widget.NewLabel("Frequency: ...")                    //Mhz - ความเร็ว MHz
	//cacheSize := widget.NewLabel("CacheSize: ...")               //CacheSize - cache size
	//featureflags := widget.NewLabel("FeatureFlags: ...") //Flags - feature flags
	//microcodeVersion := widget.NewLabel("MicrocodeVersion: ...") //Microcode - microcode version

	// cpu.Counts()
	//coreCounts := widget.NewLabel("Cores: ...")     //2*cpu.Counts()*core
	//threadCounts := widget.NewLabel("Threads: ...") //2*cpu.Counts()*thread
	//cpu.Percent()
	usageLabel := widget.NewLabel("Usage: ...") //3*cpu.Percent()
	usagePercentLabel := widget.NewLabel("usagePercentLabel : ...")
	//cpu.Times()

	content := container.NewScroll(container.NewVBox(
		//cpu.Info()
		InfoLabel,
		//coresthread,
		//cpu.Percent()
		usageLabel,
		usagePercentLabel,

		flagsLabel,

		///cpunumber, //CPU - หมายเลข CPU
		//vendorid,         //VendorID	ผู้ผลิต CPU
		//cpufamily,        //Family	CPU family
		//modelid,          //Model	model id
		//steppingversion,  //Stepping	stepping version
		//socketid,  //PhysicalID	socket id
		//coreid,    //CoreID	core id
		//coresmain, //Cores	จำนวน core
		//modelName,        //ModelName	ชื่อ CPU เต็ม
		//freq,             //Mhz	ความเร็ว MHz
		//cacheSize,        //CacheSize	cache size
		//featureflags, //Flags	feature flags
		//microcodeVersion, //Microcode	microcode version

		//cpu.Counts()
		//coreCounts,
		//threadCounts,

		//cpu.Times()

	))

	w.SetContent(container.NewBorder(nil, nil, nil, nil, content))
	w.Resize(fyne.NewSize(600, 600))

	// โหลดข้อมูล CPU static
	info, _ := cpu.Info()

	if len(info) > 0 { // cpu.Info()

		//modelName.SetText("CPU: " + info[0].ModelName) //ModelName
		modelName := info[0].ModelName

		//freq.SetText(fmt.Sprintf("Frequency: %.2f MHz", info[0].Mhz)) //Mhz
		freqSizeGhz := info[0].Mhz / 1000
		//freq.SetText(fmt.Sprintf("Turbo Boost : %.2f GHz", freqSizeGhz)) //Ghz
		/*
			var cpucpuresult string
			for i, cpucpu := range info {
				cpucpuresult += fmt.Sprintf("info[%d]:  CPU=%d\n", i, cpucpu.CPU)
			}
			cpunumber.SetText(cpucpuresult) // CPU - หมายเลข CPU
		*/
		//vendorid.SetText(fmt.Sprintf("Vendor: %s", info[0].VendorID))          //VendorID
		vendorid := info[0].VendorID
		//cpufamily.SetText(fmt.Sprintf("Family: %s", info[0].Family)) //Family	CPU family
		cpufamily := info[0].Family
		//modelid.SetText(fmt.Sprintf("Model: %s", info[0].Model)) //Model	model id
		modelid := info[0].Model
		//steppingversion.SetText(fmt.Sprintf("Stepping: %d", info[0].Stepping)) //Stepping	stepping version
		steppingversion := info[0].Stepping
		// PhysicalID	socket id
		/*
			var socketidresult string
			for i, cpu := range info {
				socketidresult += fmt.Sprintf("Info [%d], PhysicalID:  %s\n", i, cpu.PhysicalID)
			}
			socketid.SetText(socketidresult)
		*/
		/*
			//CoreID	core id
			var coreidresult string
			for i, cpu := range info {
				coreidresult += fmt.Sprintf("Info [%d], CoreID %s\n", i, cpu.CoreID)
			}
			coreid.SetText(coreidresult)
		*/
		/*
			var cpucoreresult string
			for i, cpucpu := range info {
				cpucoreresult += fmt.Sprintf("info: [%d] ,cpu core: %d\n", i, cpucpu.Cores)
			}
			coresmain.SetText(cpucoreresult) //Cores	จำนวน core
		*/

		cacheSizeMB := info[0].CacheSize / 1024
		//cacheSize.SetText(fmt.Sprintf("cacheSize: %d MB", cacheSizeMB)) //CacheSize

		flagsStr := ""
		for i, flag := range info[0].Flags {
			flagsStr += flag
			if (i+1)%6 == 0 { // ทีละ 6 flags ต่อบรรทัด
				flagsStr += "\n"
			} else {
				flagsStr += " "
			}
		}

		//featureflags.SetText(fmt.Sprintf("Flags:\n%s", flagsStr)) //Flags

		//featureflags.SetText(fmt.Sprintf("Flags: %v", info[0].Flags)) //Flags

		microcodeVersion := info[0].Microcode
		//microcodeVersion.SetText(fmt.Sprintf("microcodeVersion: %s", info[0].Microcode)) //Microcode

		/*
			cacheSize,        //CacheSize	cache size
			featureflags,     //Flags	feature flags
			microcodeVersion, //Microcode	microcode version
		*/

		/*for {
			coresmain.SetText(fmt.Sprintf("Coresmain: %d", info[0].Cores))
		}*/

		//cpu.Percent()
		// 🔄 loop อัปเดต usage
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
		//usagePercentLabel

		go func() {
			for {

				var cpuusagePercentresult string
				// ดึง CPU usage ต่อ core
				percent, _ := cpu.Percent(time.Second, true) // true = per core
				//var builder strings.Builder
				// ✅ reset ก่อนใช้
				cpuusagePercentresult = ""
				for i, usage := range percent {
					//builder.WriteString(fmt.Sprintf("Core [%d]: %.2f%%\n", i, usage))
					cpuusagePercentresult += fmt.Sprintf("Core [%d]: %.2f%%\n", i, usage)
				}
				fyne.Do(func() {

					//usagePercentLabel.SetText(builder.String())
					usagePercentLabel.SetText(cpuusagePercentresult)
				})
			}
		}()

		//cpu.Counts()
		cores, _ := cpu.Counts(false) //Physical Cores /false = คอร์จริง
		//coreCounts.SetText(fmt.Sprintf("Cores: %d", cores))

		threads, _ := cpu.Counts(true) //Logical Cores /true = รวมคอร์ที่มี Hyperthreading ด้วย หรือ(threads)
		//threadCounts.SetText(fmt.Sprintf("Threads: %d", threads))

		var cpucoreresult string
		//cpucoreresult += fmt.Sprintf("Physical Cores: %d\n", physical)
		//cpucoreresult += fmt.Sprintf("Logical Cores: %d\n", logical)
		//cpucoreresult += fmt.Sprintf("Hyperthreading: %v\nDetails: ═════════════════╗\n", threads > cores)

		// แสดงรายละเอียดแต่ละ thread
		//cpucoreresult += fmt.Sprint("\nDetails: ------------------------\n")
		for i, cpu := range info {
			cpucoreresult += fmt.Sprintf("Thread [%d] : Core [%s] : Socket [%s]\n",
				i, cpu.CoreID, cpu.PhysicalID)
		}
		//coresthread.SetText(cpucoreresult)

		//cpu.Times()

		var infoLabel string
		infoLabel += fmt.Sprintf("%s | [%.2fGHz]\n", modelName, freqSizeGhz)
		infoLabel += fmt.Sprintf("Core: [%d] Threade: [%d]\n", cores, threads)
		infoLabel += fmt.Sprintf("Vendor: %s\n", vendorid)
		infoLabel += fmt.Sprintf("Family: %s | Model: %s | Stepping: %d\n", cpufamily, modelid, steppingversion)
		infoLabel += fmt.Sprintf("cacheSize: %d MB | microcodeVersion: %s\n", cacheSizeMB, microcodeVersion)
		infoLabel += fmt.Sprintf("\nHyperthreading: %v\nDetails: ════════════════════╗\n", threads > cores)
		infoLabel += fmt.Sprintf("%s", cpucoreresult)

		//infoLabel += fmt.Sprintf("")
		InfoLabel.SetText(infoLabel)

		flagsLabel.SetText(fmt.Sprintf("FeatureFlags: ═══════════════════════════════╗\n%s", flagsStr))

	}

	w.ShowAndRun()
}
