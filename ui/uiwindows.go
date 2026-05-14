package ui

import (
	"fmt"
	"time"

	cpuinfo "test9/cpu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateWindow() {
	// ... สร้าง UI
	a := app.New()
	w := a.NewWindow("CPU Info")

	dataCPUInfo := cpuinfo.CPUdata() //ดึงข้อมูลจากไฟล์ cpuinfo.go

	//dxd.SetText(fmt.Sprintf(dx))
	//cpuui()

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
		perCoreStr := " "
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
	cpuinfo += fmt.Sprintf("CPU : %s\n", dataCPUInfo["modelName"])
	cpuinfo += fmt.Sprintf("Vendor : %s\n", dataCPUInfo["vendor"])
	cpuinfo += fmt.Sprintf("Cores : %d\n", dataCPUInfo["physical_cores"])
	cpuinfo += fmt.Sprintf("Thread : %d\n", dataCPUInfo["logical_cores"])
	cpuinfo += fmt.Sprintf("FreqMax : %.2f GHz\n", dataCPUInfo["frequency"])
	cpuinfo += fmt.Sprintf("Family : %s\n", dataCPUInfo["family"])
	cpuinfo += fmt.Sprintf("Modelid : %s\n", dataCPUInfo["modelid"])
	cpuinfo += fmt.Sprintf("Stepping : %d\n", dataCPUInfo["steppingversion"])
	cpuinfo += fmt.Sprintf("Cache : %d MB\n", dataCPUInfo["cacheSizeMB"])
	cpuinfo += fmt.Sprintf("Microcode : %s\n", dataCPUInfo["microcodeVersion"])

	cpuinfolabel.SetText(cpuinfo)

	var flagsStr string
	flagsStr += fmt.Sprintf("%v\n", dataCPUInfo["flagsStr"])

	flagsStrlabel.SetText(flagsStr)

	var detailLabel string
	detailLabel += fmt.Sprintf("%s\n", dataCPUInfo["Hyperthreading"])
	detailLabel += ("\n[  Thread  ] : [ Core ] : [ Socket ]\n")
	detailLabel += fmt.Sprintf("%s\n", dataCPUInfo["cpuThreadCoreSocketresult"])
	//detailLabel += fmt.Sprintf("Cache\nL1D : %d KB\n", dataCPUInfo["l1d_cache"]) //cpuid
	//detailLabel += fmt.Sprintf("L1I : %d KB\n", dataCPUInfo["l1i_cache"])        //cpuid
	//detailLabel += fmt.Sprintf("L2 : %d KB\n", dataCPUInfo["l2_cache"])          //cpuid
	//detailLabel += fmt.Sprintf("L3 : %d KB\n", dataCPUInfo["l3_cache"])          //cpuid
	detailLabel += fmt.Sprintf("[ Cache ]\n%s\n", dataCPUInfo["cache"])

	detail.SetText(detailLabel)

	cpuuse := container.NewScroll(
		container.NewVBox(
			//widget.NewCard("CPU Information", "", container.NewVBox(
			usageTotalLabel,
			usagePerCoreLabel,
		))

	cpu := container.NewAppTabs(

		container.NewTabItem("Overview", container.NewScroll(cpuinfolabel)),

		container.NewTabItem("Detail", container.NewScroll(detail)),
		container.NewTabItem("Flags Feature", container.NewScroll(flagsStrlabel)),
		container.NewTabItem("Usage", container.NewScroll(cpuuse)),

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
