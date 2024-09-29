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
	controlArea := container.NewGridWithColumns(5, buildBtnGroup(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), buildVolumeSlider())
	controlArea.Resize(fyne.Size{Width: 300, Height: 150})

	controlGroup := container.NewVBox(controlArea, buildSongProgress(), buildProgLabel())
	content := container.NewBorder(buildSearchForm(), controlGroup, nil, nil, buildSongList())

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

func buildSearchForm() *fyne.Container {
	searchInput := widget.NewEntry()
	searchInput.SetPlaceHolder("Search for a video...")

	fnInput := widget.NewEntry()
	fnInput.SetPlaceHolder("file name...")

	form := container.NewGridWithColumns(2, searchInput, fnInput)

	searchContainer := container.NewBorder(nil, nil, nil, widget.NewButton("Search", func() { downloadSC(searchInput.Text) }), form)
	return container.NewGridWithColumns(3, layout.NewSpacer(), searchContainer, layout.NewSpacer())
}

func buildBtnGroup() *fyne.Container {
	previousBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaSkipPrevious), func() { restartSong() })
	ppBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaPlay), func() { ppSong() })
	nextBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaSkipNext), func() { skipSong() })
	return container.NewHBox(previousBtn, ppBtn, nextBtn)
}

func buildVolumeSlider() *volumeSlider {
	volume := 69.0
	volumeBinding = binding.BindFloat(&volume)
	slider := &volumeSlider{}
	slider.ExtendBaseWidget(slider)
	slider.Min = 0
	slider.Max = 100
	slider.Bind(volumeBinding)
	return slider
}

func buildSongProgress() *progressSlider {
	progress := 0.0
	progressBinding = binding.BindFloat(&progress)
	slider := &progressSlider{}
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
			time.Sleep(time.Millisecond * 250)
			continue
		}
		timeString := fmt.Sprintf(
			"%02d:%02d / %02d:%02d",
			int(songElapsed)/60, int(songElapsed)%60, currentSong.durSec/60, currentSong.durSec%60)

		progLabelBinding.Set(timeString)
		progressBinding.Set(songElapsed / float64(currentSong.durSec) * PROG_MAX)

		time.Sleep(time.Millisecond * 250)
		songElapsed += 0.250
	}
}

var progressBinding binding.Float
var progLabelBinding binding.String
var volumeBinding binding.Float
var songElapsed float64 = 0.0

const PROG_MAX float64 = 10000
