package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("AppContainer")
	w.Resize(fyne.NewSize(1200, 700))

	var songList = []string{"song1", "song2", "song3", "song4", "song5", "song6", "song7", "song8", "song9", "song10", "song11", "song12", "song13", "song14", "song15", "song16", "song17", "song18"}

	scrollArea := widget.NewList(
		func() int { return len(songList) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(i widget.ListItemID, o fyne.CanvasObject) { o.(*widget.Label).SetText(songList[i]) })

	previousBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaSkipPrevious), func() {})
	ppBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaPause), func() {})
	nextBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaSkipNext), func() {})
	controlArea := container.NewHBox(previousBtn, ppBtn, nextBtn)
	controlArea.Resize(fyne.Size{Width: 300, Height: 150})

	progressBar := widget.NewProgressBar()
	go func() {
		songDur := 180.0
		for i := 1.0; i < songDur; i++ {
			time.Sleep(time.Second)
			progressBar.SetValue(i / songDur)
		}
	}()

	controlGroup := container.NewVBox(controlArea, progressBar)

	content := container.NewBorder(nil, controlGroup, nil, nil, scrollArea)
	w.SetContent(content)

	w.ShowAndRun()
}
