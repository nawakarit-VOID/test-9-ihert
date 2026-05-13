package ui

import (
	"fmt"
	"time"

	cpuinfo "test9/cpu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shirou/gopsutil/v3/cpu"
)

func CreateWindow() {
	// ... สร้าง UI
	a := app.New()
	w := a.NewWindow("CPU Info")

	datagopsutil := cpuinfo.CPUdata() //ดึงข้อมูลจากไฟล์ cpuinfo.go

	x := widget.NewLabel("x...")
	y := widget.NewLabel("y...")

	dxd := widget.NewLabel("dxd ...")

	//dxd.SetText(fmt.Sprintf(dx))
	//cpuui()

	datagopsutillabel := widget.NewLabel("datagopsutillabel...")

	// Labels
	//overview := widget.NewLabel("overview...")
	detail := widget.NewLabel("detail...")

	//update cpu usage

	usageTotalLabel := widget.NewLabel("Total Usage: 0%")
	usagePerCoreLabel := widget.NewLabel("Per Core: -")

	// สร้าง monitor
	monitor := cpuinfo.NewCPUMonitor(1*time.Second, func(data cpuinfo.CPUData) {
		// แสดง usage รวม
		fyne.Do(func() {
			usageTotalLabel.SetText(fmt.Sprintf("Usage Avg : %.2f%%", data.UsageTotal))
		})

		// แสดง usage ต่อ core
		perCoreStr := ""
		for i, usage := range data.UsagePerCore {
			perCoreStr += fmt.Sprintf("Core [ %d ] : %.1f%%\n ", i, usage)
		}
		fyne.Do(func() {
			usagePerCoreLabel.SetText(perCoreStr)
		})
	})

	// เริ่ม monitoring
	monitor.Start()

	//

	cpuinfolabel := widget.NewLabel("cpuinfolabel...") //Overview

	flagsStrlabel := widget.NewLabel("flagsStrlabel...") //flagfeature

	//cpu.Times()

	var cpuinfo string
	cpuinfo += fmt.Sprintf("CPU : %s\n", datagopsutil["modelName"])
	cpuinfo += fmt.Sprintf("Vendor : %s\n", datagopsutil["vendor"])
	cpuinfo += fmt.Sprintf("Cores : [ %d ]\n", datagopsutil["physical_cores"])
	cpuinfo += fmt.Sprintf("Thread : [ %d ]\n", datagopsutil["logical_cores"])
	cpuinfo += fmt.Sprintf("FrequencyMax : [ %.2f GHz ]\n", datagopsutil["frequency"])
	cpuinfo += fmt.Sprintf("Family : [ %s ]\n", datagopsutil["family"])
	cpuinfo += fmt.Sprintf("Modelid : [ %s ]\n", datagopsutil["modelid"])
	cpuinfo += fmt.Sprintf("SteppingVersion : [ %d ]\n", datagopsutil["steppingversion"])
	cpuinfo += fmt.Sprintf("CacheSize : [ %d MB ]\n", datagopsutil["cacheSizeMB"])
	cpuinfo += fmt.Sprintf("MicrocodeVersion : [ %s ]\n", datagopsutil["microcodeVersion"])

	cpuinfolabel.SetText(cpuinfo)

	var flagsStr string
	flagsStr += fmt.Sprintf("%v\n", datagopsutil["flagsStr"])

	flagsStrlabel.SetText(flagsStr)
	/*
		fyne.Do(func() {
			usageLabel.SetText(fmt.Sprintf("CPU Avg: %.2f%%", usage))
		})
	*/
	var cachs string
	//cpuinfo += fmt.Sprintf("Brand: %s\n", datagopsutil["brand"]) //cpuid
	cachs += fmt.Sprintf("L3: %d KB", datagopsutil["l3_cache"]) //cpuid

	info, _ := cpu.Info()

	if len(info) > 0 {

		/*
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

				}
			}()

			fyne.Do(func() {

				//usagePercentLabel.SetText(builder.String())
				usagePercentLabel.SetText(cpuusagePercentresult)
			})
		*/
	}
	//cpu.Counts()
	cores, _ := cpu.Counts(false) //Physical Cores /false = คอร์จริง
	//coreCounts.SetText(fmt.Sprintf("Cores: %d", cores))

	threads, _ := cpu.Counts(true) //Logical Cores /true = รวมคอร์ที่มี Hyperthreading ด้วย หรือ(threads)
	//threadCounts.SetText(fmt.Sprintf("Threads: %d", threads))

	var cpuThreadCoreSocketresult string
	//cpuThreadCoreSocketresult += fmt.Sprintf("Physical Cores: %d\n", physical)
	//cpuThreadCoreSocketresult += fmt.Sprintf("Logical Cores: %d\n", logical)
	//cpuThreadCoreSocketresult += fmt.Sprintf("Hyperthreading: %v\nDetails: ═════════════════╗\n", threads > cores)

	// แสดงรายละเอียดแต่ละ thread
	//cpuThreadCoreSocketresult += fmt.Sprint("\nDetails: ------------------------\n")
	for i, cpu := range info {
		cpuThreadCoreSocketresult += fmt.Sprintf("Thread [%d] : Core [%s] : Socket [%s]\n",
			i, cpu.CoreID, cpu.PhysicalID)
	}
	//coresthread.SetText(cpucoreresult)

	//cpu.Times()

	var detailLabel string
	detailLabel += fmt.Sprintf("Hyperthreading: [ %v ]", threads > cores)
	detailLabel += ("\n\n[  Thread  ] : [ Core ] : [ Socket ]\n")
	detailLabel += fmt.Sprintf("%s", cpuThreadCoreSocketresult)

	detail.SetText(detailLabel)
	//flagsLabel.SetText(fmt.Sprintf("%s", flagsStr))

	cpuuse := container.NewScroll(
		container.NewVBox(
			//widget.NewCard("CPU Information", "", container.NewVBox(
			usageTotalLabel,
			usagePerCoreLabel,
		))

	cpu := container.NewAppTabs(
		//container.NewScroll(container.NewVBox(

		//widget.NewRichTextFromMarkdown("# CPU Overview"),
		//cpu.Info()
		container.NewTabItem("Overview", container.NewScroll(cpuinfolabel)),
		//InfoLabel,
		//coresthread,
		//cpu.Percent()
		//container.NewTabItem("Cache", container.NewScroll(nil)),

		container.NewTabItem("Detail", container.NewScroll(detail)),

		container.NewTabItem("Flags Feature", container.NewScroll(flagsStrlabel)),
		//flagsLabel,

		container.NewTabItem("Usage", container.NewScroll(cpuuse)),
		//usageLabel,
		container.NewTabItem("x", container.NewScroll(x)),
		container.NewTabItem("y--", container.NewScroll(y)),

		container.NewTabItem("y--", container.NewScroll(dxd)),

		container.NewTabItem("datagopsutillabel", container.NewScroll(datagopsutillabel)),

		//container.NewTabItem("CPU", container.NewScroll(nil)),
		//usagePercentLabel,

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

	)

	tabs := container.NewAppTabs(
		container.NewTabItem("CPU", container.NewScroll(cpu)),
		//container.NewTabItem("Cache", container.NewScroll(nil)),
		//container.NewTabItem("Features", nil),
		//container.NewTabItem("Security", container.NewScroll(nil)),
		//container.NewTabItem("Virtualization", container.NewScroll(nil)),
	)

	//w.SetContent(container.NewBorder(nil, nil, nil, nil, cpu))
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}
