package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var progressBinding binding.Float

func GetGUI() *fyne.Container {
	controlArea := container.NewHBox(buildBtnGroup(), layout.NewSpacer(), buildVolumeSlider())
	controlArea.Resize(fyne.Size{Width: 300, Height: 150})

	controlGroup := container.NewVBox(controlArea, buildSongProgress())
	content := container.NewBorder(nil, controlGroup, nil, nil, buildScrollArea())

	return content
}

func buildScrollArea() *widget.List {
	return widget.NewList(
		func() int { return len(songList) },
		func() fyne.CanvasObject { return widget.NewButton("template", func() {}) },
		func(i widget.ListItemID, btn fyne.CanvasObject) {
			btn.(*widget.Button).SetText(songList[i].name)
			btn.(*widget.Button).OnTapped = func() { go readSong(songList[i]) }
		})
}

func buildBtnGroup() *fyne.Container {
	previousBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaSkipPrevious), func() { restartSong() })
	ppBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaPlay), func() { ppSong() })
	nextBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaSkipNext), func() { skipSong() })
	return container.NewHBox(previousBtn, ppBtn, nextBtn)
}

func buildVolumeSlider() *widget.Slider {
	volume := 69.0
	volumeBinding := binding.BindFloat(&volume)
	return widget.NewSliderWithData(0.0, 100.0, volumeBinding)
}

func buildSongProgress() *widget.Slider {
	progress := 0.0
	progressBinding = binding.BindFloat(&progress)
	slider := widget.NewSliderWithData(0.0, 100.0, progressBinding)
	return slider
}

func updateSongProgress() {
	for i := 0; i < currentSong.durSec; i++ {
		progressBinding.Set(float64(i) / float64(currentSong.durSec) * 100)
		fmt.Printf("%d / %2d:%2d\n", i, currentSong.durSec/60, currentSong.durSec%60)
		time.Sleep(time.Second)
	}
}
