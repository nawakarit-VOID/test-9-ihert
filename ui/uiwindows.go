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

	dataCPUInfo := cpuinfo.CPUdata() //ดึงข้อมูลจากไฟล์ cpuinfo.go

	overviewlabel := widget.NewLabel("Overviewlabel...") //Overview
	detailLabel := widget.NewLabel("detailLabel...")
	//update cpu usage
	usageTotalLabel := widget.NewLabel("CPU Avg...")
	usagePerCoreLabel := widget.NewLabel("CPU...")
	totalavgLabel := widget.NewLabel("TotalavgLabel...")

	timesStrLabel := widget.NewLabel("timestimesStrLabel...")
	timesLabel := widget.NewLabel("timesLabel...")
	meanLabel := widget.NewLabel("meanLabel...")

	flagsStrlabel := widget.NewLabel("flagsStrlabel...") //flagfeature

	// สร้าง monitor
	monitor := cpuinfo.NewCPUMonitor(1*time.Second, func(data cpuinfo.StCPUData) {

		/*		fyne.Do(func() {
					usageTotalLabel.SetText(fmt.Sprintf("Usage Avg : %.2f%%", data.UsageTotal))
				})
		*/
		// แสดง usage ต่อ core
		var perCoreStr string = ""
		for i, usage := range data.UsagePerCore {
			perCoreStr += fmt.Sprintf("Core [ %d ] : %.1f%%\n", i, usage)
		}

		var timesStr string = ""
		for c, v := range data.Times {
			//for _, t := range data.Times {
			timesStr += fmt.Sprintf(
				"CPU: [ %d ] | User: %.2f s | System: %.2f s | Idle: %.2f s | Nice: %.2f s | Iowait: %.2f s | Irq %.2f s | Softirq %.2f s | Steal %.2f s | Guest %.2f s | GuestNice %.2f s\n",
				c, v.User, v.System, v.Idle, v.Nice, v.Iowait, v.Irq, v.Softirq, v.Steal, v.Guest, v.GuestNice)

		}

		fyne.Do(func() {
			usageTotalLabel.SetText(fmt.Sprintf("Usage Avg : %.2f%%", data.UsageTotal)) // แสดง usage รวม
			usagePerCoreLabel.SetText(perCoreStr)                                       // แสดง usage ต่อ core
			timesStrLabel.SetText(timesStr)                                             //cpu.Times rang
			totalavgLabel.SetText(fmt.Sprintf("%s", data.TotalavgLabel))                //แสดง timeUse Avg all core
			timesLabel.SetText(fmt.Sprintf("%s", data.TimesLabel))                      //แสดง timeUse all core
			meanLabel.SetText(fmt.Sprintln(`***
User : โปรแกรมของผู้ใช้
System : ระบบ
Idle : ไม่ได้ทำอะไร
Nice : เวลาที่ใช้กับ process ที่ถูกปรับ priority (nice)
Iowait : เวลาที่ CPU รอ I/O เช่น disk หรือ network
Irq : เวลาที่ใช้จัดการ Hardware ที่ขัดจังหวะ
Softirq : เวลาที่ใช้จัดการ Software ที่ขัดจังหวะ
Steal : เวลาที่ VM ถูก hypervisor แย่ง CPU ไป
Guest : เวลาที่ CPU ใช้งาน guest virtual machine
GuestNice : เวลาที่ guest VM ใช้งานแบบ nice priority`))

		})
	})

	monitor.Start() // เริ่ม monitoring

	//var overview string
	overviewlabel.SetText(fmt.Sprintf("%s\n", dataCPUInfo["Overview"]))
	detailLabel.SetText(fmt.Sprintf("%s\n", dataCPUInfo["Detail"]))
	flagsStrlabel.SetText(fmt.Sprintf("%v\n", dataCPUInfo["FlagsStr"]))

	//จัดหน้า
	cpuuse := container.NewScroll(
		container.NewVBox(
			//widget.NewCard("CPU Information", "", container.NewVBox(
			usageTotalLabel,
			usagePerCoreLabel,
		))

	cputimeusage := container.NewScroll(
		container.NewVBox(
			totalavgLabel,
			timesLabel,
			timesStrLabel,
			meanLabel,
		))

	cpu := container.NewAppTabs(
		//container.NewTabItem("TEST", container.NewScroll(cputimeuse)),
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
