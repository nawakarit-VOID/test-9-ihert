package ui

import (
	"fmt"

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
	overview := widget.NewLabel("overview...")
	detail := widget.NewLabel("detail...")
	flagsLabel := widget.NewLabel("flags.feature...")

	usageLabel := widget.NewLabel("Usage...") //3*cpu.Percent()
	usagePercentLabel := widget.NewLabel("Usage.PercentLabel...")
	//cpu.Times()

	var datagopsutil1 string
	datagopsutil1 += fmt.Sprintf("Brand: %s\n", datagopsutil["brand"])
	datagopsutil1 += fmt.Sprintf("L3: %d KB", datagopsutil["l3_cache"])

	datagopsutillabel.SetText(datagopsutil1)

	info, _ := cpu.Info()

	if len(info) > 0 {

		/*
			fyne.Do(func() {
				usageLabel.SetText(fmt.Sprintf("CPU Avg: %.2f%%", usage))
			})


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
		/*
			var infoLabel string
			infoLabel += fmt.Sprintf("%s | [ %.2fGHz ]\n\n", modelName, freqSizeGhz)
			infoLabel += fmt.Sprintf("Core: [ %d ]\n", cores)
			infoLabel += fmt.Sprintf("Threade: [ %d ]\n", threads)
			infoLabel += fmt.Sprintf("Vendor: [ %s ]\n", vendorid)
			infoLabel += fmt.Sprintf("Family: [ %s ]\n", cpufamily)
			infoLabel += fmt.Sprintf("Model: [ %s ]\n", modelid)
			infoLabel += fmt.Sprintf("Stepping: [ %d ]\n", steppingversion)
			infoLabel += fmt.Sprintf("CacheSize: [ %d ] MB\n", cacheSizeMB)
			infoLabel += fmt.Sprintf("MicrocodeVersion: [ %s ]\n", microcodeVersion)
			//infoLabel += fmt.Sprintf("")
			overview.SetText(infoLabel)
		*/
		var detailLabel string
		detailLabel += fmt.Sprintf("Hyperthreading: [ %v ]", threads > cores)
		detailLabel += ("\n\n[  Thread  ] : [ Core ] : [ Socket ]\n")
		detailLabel += fmt.Sprintf("%s", cpuThreadCoreSocketresult)

		detail.SetText(detailLabel)

		//flagsLabel.SetText(fmt.Sprintf("%s", flagsStr))

	}
	cpuuse := container.NewScroll(
		container.NewVBox(
			usageLabel,
			usagePercentLabel))

	cpu := container.NewAppTabs(
		//container.NewScroll(container.NewVBox(

		//widget.NewRichTextFromMarkdown("# CPU Overview"),
		//cpu.Info()
		container.NewTabItem("Overview", container.NewScroll(overview)),
		//InfoLabel,
		//coresthread,
		//cpu.Percent()
		//container.NewTabItem("Cache", container.NewScroll(nil)),

		container.NewTabItem("Detail", container.NewScroll(detail)),

		container.NewTabItem("Flags Feature", container.NewScroll(flagsLabel)),
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
