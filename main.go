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
	ppBtn := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaPause), func() { initMp3() })
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
	// ref: https://github.com/ebitengine/oto?tab=readme-ov-file#usage
	fileBytes, err := os.ReadFile("Yee.mp3")
	if err != nil {
		panic("Failed to read mp3 file: " + err.Error())
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("Failed to decode mp3 bytes reader: " + err.Error())
	}

	// oto config
	otoConfig := oto.NewContextOptions{SampleRate: 44100, ChannelCount: 2, Format: oto.FormatSignedInt16LE}

	otoContext, readyChan, err := oto.NewContext(&otoConfig)
	if err != nil {
		panic("Oto new context failed : " + err.Error())
	}

	<-readyChan // this is waiting for hardware audio devices to be ready

	player := otoContext.NewPlayer(decodedMp3)

	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	// This is how we will seek using the scrubber later on!
	//
	// newPos, err := player.(io.Seeker).Seek(0, io.SeekStart)
	// if err != nil{
	//     panic("player.Seek failed: " + err.Error())
	// }
	// println("Player is now at position:", newPos)
	// player.Play()
	//
	err = player.Close()
	if err != nil {
		panic("failed closing player: " + err.Error())
	}
}
