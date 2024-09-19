package main

import (
	"bytes"
	"io"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

var otoGlobalContext oto.Context
var player oto.Player
var currentSong Song
var songList = []Song{{name: "yee", path: "Yee.mp3"}, {name: "noblest strive", path: "bladee.mp3"}, {name: "song of storms", path: "zelda.mp3"}}

func main() {
	initMp3()

	a := app.New()
	w := a.NewWindow("AppContainer")
	w.Resize(fyne.NewSize(1200, 700))

	scrollArea := widget.NewList(
		func() int { return len(songList) },
		func() fyne.CanvasObject { return widget.NewButton("template", func() {}) },
		func(i widget.ListItemID, btn fyne.CanvasObject) {
			btn.(*widget.Button).SetText(songList[i].name)
			btn.(*widget.Button).OnTapped = func() { go playSong(songList[i]) }
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
	w.SetContent(content)

	w.ShowAndRun()
}

func initMp3() {
	// oto config
	otoConfig := oto.NewContextOptions{SampleRate: 44100, ChannelCount: 2, Format: oto.FormatSignedInt16LE}

	otoContext, readyChan, err := oto.NewContext(&otoConfig)
	if err != nil {
		panic("Oto new context failed : " + err.Error())
	}

	otoGlobalContext = *otoContext

	<-readyChan

}

func decodeMp3(song Song) mp3.Decoder {
	fileBytes, err := os.ReadFile("resources/" + song.path)
	if err != nil {
		panic("Failed to read mp3 file: " + err.Error())
	}
	fileBytesReader := bytes.NewReader(fileBytes)
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("Failed to decode mp3 bytes reader: " + err.Error())
	}
	return *decodedMp3
}

func playSong(song Song) {
	currentSong = song
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
			playSong(songList[(i+1)%(len(songList))])
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

type Song struct {
	name string
	path string
}
