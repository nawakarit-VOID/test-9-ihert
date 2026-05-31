package ui

import (
	cpuinfo "test9/cpu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func CreateWindow() {

	a := app.New()
	w := a.NewWindow("CPU Info")

	cpuTabs := cpuinfo.CpuTabs()

	tabs := container.NewAppTabs(
		container.NewTabItem("CPU", container.NewScroll(cpuTabs)),
		//container.NewTabItem("cpu", container.NewScroll(cpu)),
		//container.NewTabItem("Features", nil),
		//container.NewTabItem("Security", container.NewScroll(nil)),
		//container.NewTabItem("Virtualization", container.NewScroll(nil)),
	)

	//w.SetContent(container.NewBorder(nil, nil, nil, nil, cpu))
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}
