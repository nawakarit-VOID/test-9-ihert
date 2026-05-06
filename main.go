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
	modelLabel := widget.NewLabel("CPU: loading...")
	coreLabel := widget.NewLabel("Cores: ...")
	threadLabel := widget.NewLabel("Threads: ...")
	freqLabel := widget.NewLabel("Frequency: ...")
	usageLabel := widget.NewLabel("Usage: ...")

	content := container.NewVBox(
		modelLabel,
		coreLabel,
		threadLabel,
		freqLabel,
		usageLabel,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 200))

	// โหลดข้อมูล CPU static
	info, _ := cpu.Info()
	if len(info) > 0 {
		modelLabel.SetText("CPU: " + info[0].ModelName)
		freqLabel.SetText(fmt.Sprintf("Frequency: %.2f MHz", info[0].Mhz))
	}

	cores, _ := cpu.Counts(false)
	threads, _ := cpu.Counts(true)

	coreLabel.SetText(fmt.Sprintf("Cores: %d", cores))
	threadLabel.SetText(fmt.Sprintf("Threads: %d", threads))

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
