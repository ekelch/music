package main

import (
	"bytes"
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

func main() {
	initMp3()

	a := app.New()
	w := a.NewWindow("AppContainer")
	w.Resize(fyne.NewSize(1200, 700))

	var songList = []Song{Song{name: "yee", path: "Yee.mp3"}, Song{name: "noblest strive", path: "bladee.mp3"}, Song{name: "niki", path: "niki.mp3"}}

	scrollArea := widget.NewList(
		func() int { return len(songList) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(i widget.ListItemID, o fyne.CanvasObject) { o.(*widget.Label).SetText(songList[i].name) })

	testSong := Song{name: "Yee", path: "Yee.mp3"}
	previousBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaSkipPrevious), func() {})
	ppBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaPlay), func() { playSong(testSong) })
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

func playSong(song Song) {
	fileBytes, err := os.ReadFile(song.path)
	if err != nil {
		panic("Failed to read mp3 file: " + err.Error())
	}
	fileBytesReader := bytes.NewReader(fileBytes)
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("Failed to decode mp3 bytes reader: " + err.Error())
	}

	player := otoGlobalContext.NewPlayer(decodedMp3)
	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	err = player.Close()
	if err != nil {
		panic("failed closing player: " + err.Error())
	}
}

type Song struct {
	name string
	path string
}
