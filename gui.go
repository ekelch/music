package main

import (
	"io"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ebitengine/oto/v3"
)

func getGUI() *fyne.Container {
	scrollArea := widget.NewList(
		func() int { return len(songList) },
		func() fyne.CanvasObject { return widget.NewButton("template", func() {}) },
		func(i widget.ListItemID, btn fyne.CanvasObject) {
			btn.(*widget.Button).SetText(songList[i].name)
			btn.(*widget.Button).OnTapped = func() { go readSong(songList[i]) }
		})

	previousBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaSkipPrevious), func() { restartSong() })
	ppBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaPlay), func() { ppSong() })
	nextBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaSkipNext), func() { skipSong() })
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

	return content
}

func readSong(song Song) {
	currentSong = song
	if (player != oto.Player{}) {
		player.Pause()
	}
	decodedMp3 := decodeMp3(song)
	player = *otoGlobalContext.NewPlayer(&decodedMp3)
	player.Play()
}

func ppSong() {
	if player.IsPlaying() {
		player.Pause()
	} else {
		player.Play()
	}
}

func skipSong() {
	ppSong()
	for i, v := range songList {
		if v == currentSong {
			readSong(songList[(i+1)%(len(songList))])
			break
		}
	}
}

func restartSong() {
	_, err := player.Seek(0, io.SeekStart)
	if err != nil {
		panic("Failed to seek start of song: " + err.Error())
	}
}
