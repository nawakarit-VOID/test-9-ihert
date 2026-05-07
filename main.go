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
	cpunumber := widget.NewLabel("CPUnumber: ...")               //CPU - หมายเลข CPU
	vendorid := widget.NewLabel("Vendorid: ...")                 //VendorID - ผู้ผลิต CPU
	cpufamily := widget.NewLabel("CPUfamily: ...")               //Family - CPU family
	modelid := widget.NewLabel("Modelid: ...")                   //Model - model id
	steppingversion := widget.NewLabel("Stepping version: ...")  //Stepping - stepping version
	socketid := widget.NewLabel("Socketid: ...")                 //PhysicalID - socket id
	coreid := widget.NewLabel("Coreid: ...")                     //CoreID - core id
	coresmain := widget.NewLabel("Coresmain: ...")               //Cores - จำนวน core
	modelName := widget.NewLabel("CPU: loading...")              //ModelName - ชื่อ CPU เต็ม
	freq := widget.NewLabel("Frequency: ...")                    //Mhz - ความเร็ว MHz
	cacheSize := widget.NewLabel("CacheSize: ...")               //CacheSize - cache size
	featureflags := widget.NewLabel("FeatureFlags: ...")         //Flags - feature flags
	microcodeVersion := widget.NewLabel("MicrocodeVersion: ...") //Microcode - microcode version

	// cpu.Counts()
	coreCounts := widget.NewLabel("Cores: ...")     //2*cpu.Counts()*core
	threadCounts := widget.NewLabel("Threads: ...") //2*cpu.Counts()*thread
	//cpu.Percent()
	usageLabel := widget.NewLabel("Usage: ...") //3*cpu.Percent()
	//cpu.Times()

	content := container.NewVBox(
		//cpu.Info()
		cpunumber,        //CPU - หมายเลข CPU
		vendorid,         //VendorID	ผู้ผลิต CPU
		cpufamily,        //Family	CPU family
		modelid,          //Model	model id
		steppingversion,  //Stepping	stepping version
		socketid,         //PhysicalID	socket id
		coreid,           //CoreID	core id
		coresmain,        //Cores	จำนวน core
		modelName,        //ModelName	ชื่อ CPU เต็ม
		freq,             //Mhz	ความเร็ว MHz
		cacheSize,        //CacheSize	cache size
		featureflags,     //Flags	feature flags
		microcodeVersion, //Microcode	microcode version

		//cpu.Counts()
		coreCounts,
		threadCounts,
		//cpu.Percent()
		usageLabel,
		//cpu.Times()

	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 200))

	// โหลดข้อมูล CPU static
	info, _ := cpu.Info()

	if len(info) > 0 { // cpu.Info()
		var cpucpuresult string
		for i, cpucpu := range info {
			cpucpuresult += fmt.Sprintf("info[%d]:  CPU=%d\n", i, cpucpu.CPU)
		}
		cpunumber.SetText(cpucpuresult) // CPU - หมายเลข CPU

		vendorid.SetText(fmt.Sprintf("Vendor: %s", info[0].VendorID))          //VendorID
		cpufamily.SetText(fmt.Sprintf("Family: %s Family", info[0].Family))    //Family	CPU family
		modelid.SetText(fmt.Sprintf("Model: %s", info[0].Model))               //Model	model id
		steppingversion.SetText(fmt.Sprintf("Stepping: %d", info[0].Stepping)) //Stepping	stepping version
		socketid.SetText(fmt.Sprintf("PhysicalID: %s", info[0].PhysicalID))    //PhysicalID	socket id
		coreid.SetText(fmt.Sprintf("CoreID: %s", info[0].CoreID))              //CoreID	core id

		var cpucoreresult string
		for i, cpucpu := range info {
			cpucoreresult += fmt.Sprintf("info: [%d] ,cpu core: %d\n", i, cpucpu.Cores)
		}
		coresmain.SetText(cpucoreresult) //Cores	จำนวน core

		modelName.SetText("CPU: " + info[0].ModelName)                //ModelName
		freq.SetText(fmt.Sprintf("Frequency: %.2f MHz", info[0].Mhz)) //Mhz
		//CacheSize
		//Flags
		//Microcode

		/*
			cacheSize,        //CacheSize	cache size
			featureflags,     //Flags	feature flags
			microcodeVersion, //Microcode	microcode version
		*/

		/*for {
			coresmain.SetText(fmt.Sprintf("Coresmain: %d", info[0].Cores))
		}*/

	}

	//cpu.Counts()
	cores, _ := cpu.Counts(false) //Physical Cores /false = คอร์จริง
	coreCounts.SetText(fmt.Sprintf("Cores: %d", cores))

	threads, _ := cpu.Counts(true) //Logical Cores /true = รวมคอร์ที่มี Hyperthreading ด้วย หรือ(threads)
	threadCounts.SetText(fmt.Sprintf("Threads: %d", threads))

	//cpu.Percent()
	// 🔄 loop อัปเดต usage
	go func() {
		for {

			percent, _ := cpu.Percent(1*time.Second, false)
			if len(percent) > 0 {
				usage := percent[0]

				fyne.Do(func() {
					usageLabel.SetText(fmt.Sprintf("Usage: %.2f%%", usage))
				})
			}
		}
	}()

	//cpu.Times()

	w.ShowAndRun()
}
