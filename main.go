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
	// cpu.Info()
	//CPU - หมายเลข CPU
	Vendorid := widget.NewLabel("Vendorid: ...") //VendorID - ผู้ผลิต CPU
	//Family - CPU family
	//Model - model id
	//Stepping - stepping version
	//PhysicalID - socket id
	//CoreID - core id
	coresmain := widget.NewLabel("Coresmain: ...")       //Cores - จำนวน core
	modelNameLabel := widget.NewLabel("CPU: loading...") //ModelName - ชื่อ CPU เต็ม
	freqLabel := widget.NewLabel("Frequency: ...")       //Mhz - ความเร็ว MHz
	//CacheSize - cache size
	//Flags - feature flags
	//Microcode - microcode version

	// cpu.Counts()
	coreCountsLabel := widget.NewLabel("Cores: ...")     //2*cpu.Counts()*core
	threadCountsLabel := widget.NewLabel("Threads: ...") //2*cpu.Counts()*thread
	//cpu.Percent()
	usageLabel := widget.NewLabel("Usage: ...") //3*cpu.Percent()
	//cpu.Times()

	content := container.NewVBox(
		//cpu.Info()
		coresmain, //CPU
		//VendorID	ผู้ผลิต CPU
		//Family	CPU family
		//Model	model id
		//Stepping	stepping version
		//PhysicalID	socket id
		//CoreID	core id
		//Cores	จำนวน core
		modelNameLabel, //ModelName	ชื่อ CPU เต็ม
		//Mhz	ความเร็ว MHz
		//CacheSize	cache size
		//Flags	feature flags
		//Microcode	microcode version

		//cpu.Counts()
		coreCountsLabel,
		threadCountsLabel,
		freqLabel,
		//cpu.Percent()
		usageLabel,
		//cpu.Times()
		Vendorid,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 200))

	// โหลดข้อมูล CPU static
	info, _ := cpu.Info()
	if len(info) > 0 {
		modelNameLabel.SetText("CPU: " + info[0].ModelName)
		freqLabel.SetText(fmt.Sprintf("Frequency: %.2f MHz", info[0].Mhz))
		//cpulabel.SetText(fmt.Sprintf("CCCPU: %.2f ", info[0].VendorID))
		Vendorid.SetText(fmt.Sprintf("Vendor: %s", info[0].VendorID))
		/*for {
			coresmain.SetText(fmt.Sprintf("Coresmain: %d", info[0].Cores))
		}*/
	}

	cores, _ := cpu.Counts(false)  //Physical Cores /false = คอร์จริง
	threads, _ := cpu.Counts(true) //Logical Cores /true = รวมคอร์ที่มี Hyperthreading ด้วย หรือ(threads)

	coreCountsLabel.SetText(fmt.Sprintf("Cores: %d", cores))
	threadCountsLabel.SetText(fmt.Sprintf("Threads: %d", threads))

	// 🔄 loop อัปเดต usage
	go func() {
		for {
			percent, _ := cpu.Percent(1*time.Second, false)
			if len(percent) > 0 {
				usage := percent[0]

				usageLabel.SetText(fmt.Sprintf("Usage: %.2f%%", usage))
			}
		}
	}()

	w.ShowAndRun()
}
