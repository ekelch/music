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
	"github.com/ebitengine/oto/v3"
)

func GetGUI() *fyne.Container {
	controlArea := container.NewHBox(buildBtnGroup(), layout.NewSpacer(), buildVolumeSlider())
	controlArea.Resize(fyne.Size{Width: 300, Height: 150})

	controlGroup := container.NewVBox(controlArea, buildSongProgress(), buildProgLabel())
	content := container.NewBorder(nil, controlGroup, nil, nil, buildSongList())

	return content
}

func buildSongList() *widget.List {
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

func buildSongProgress() *tappableSlider {
	progress := 0.0
	progressBinding = binding.BindFloat(&progress)
	slider := &tappableSlider{}
	slider.ExtendBaseWidget(slider)
	slider.Min = 0
	slider.Max = PROG_MAX
	slider.Bind(progressBinding)
	return slider
}

func buildProgLabel() *widget.Label {
	str := ""
	progLabelBinding = binding.BindString(&str)
	sLabel := widget.NewLabelWithData(progLabelBinding)
	return sLabel
}

func setProg() {
	for {
		if (player == oto.Player{}) || !player.IsPlaying() {
			continue
		}
		timeString := fmt.Sprintf(
			"%02d:%02d / %02d:%02d",
			int(songTime)/60, int(songTime)%60, currentSong.durSec/60, currentSong.durSec%60)

		progLabelBinding.Set(timeString)
		progressBinding.Set(songTime / float64(currentSong.durSec) * PROG_MAX)

		time.Sleep(time.Millisecond * 250)
		songTime += 0.250
	}
}

var progressBinding binding.Float
var progLabelBinding binding.String
var songTime float64 = 0.0

const PROG_MAX float64 = 10000
