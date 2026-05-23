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

	a := app.New()
	w := a.NewWindow("CPU Info")

	//cpu
	dataCPUInfo := cpuinfo.CPUdata() //ดึงข้อมูลจากไฟล์ cpuinfo.go
	//sub cpu
	overviewlabel := widget.NewLabel("Overviewlabel...") //Overview
	detailLabel := widget.NewLabel("detailLabel...")
	flagsStrlabel := widget.NewLabel("flagsStrlabel...") //flagfeature
	usageLabel := widget.NewLabel("usageLabel...")
	TimesusageLabel := widget.NewLabel("TimesusageLabel...")

	//timesStrLabel := widget.NewLabel("timestimesStrLabel...")
	//timesLabel := widget.NewLabel("timesLabel...")
	//meanLabel := widget.NewLabel("meanLabel...")

	//xLabel := widget.NewLabel("xLabel...")

	//รับ cpu
	fyne.Do(func() {
		overviewlabel.SetText(fmt.Sprintf("%s\n", dataCPUInfo["Overview"]))     //1
		detailLabel.SetText(fmt.Sprintf("%s\n", dataCPUInfo["Detail"]))         //2
		flagsStrlabel.SetText(fmt.Sprintf("%v\n", dataCPUInfo["FlagsFeature"])) //3
	})

	// สร้าง monitor cpu
	monitor := cpuinfo.NewCPUMonitor(1*time.Second, func(data cpuinfo.StCPUData) {
		fyne.Do(func() {
			usageLabel.SetText(fmt.Sprintf("%s", data.Usage))           //4 // แสดง usage รวม
			TimesusageLabel.SetText(fmt.Sprintf("%s", data.Timesusage)) //5
		})
	})
	monitor.Start() // เริ่ม monitoring

	//จัดหน้า
	cpuuse := container.NewScroll(
		container.NewVBox(
			//widget.NewCard("CPU Information", "", container.NewVBox(
			usageLabel,
		))

	cputimeusage := container.NewScroll(
		container.NewVBox(
			TimesusageLabel,
		))

	cpu := container.NewAppTabs(
		//container.NewTabItem("TEST", container.NewScroll(xLabel)),
		container.NewTabItem("Overview", container.NewScroll(overviewlabel)),
		container.NewTabItem("Detail", container.NewScroll(detailLabel)),
		container.NewTabItem("Flags Feature", container.NewScroll(flagsStrlabel)),
		container.NewTabItem("Usage", container.NewScroll(cpuuse)),
		container.NewTabItem("TimeUsage", container.NewScroll(cputimeusage)),
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
	w.Resize(fyne.NewSize(1200, 600))
	w.ShowAndRun()
}
