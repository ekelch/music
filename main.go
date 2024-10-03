package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ebitengine/oto/v3"
)

var otoGlobalContext oto.Context
var player oto.Player
var currentSong Song
var songList []string

const WINDOW_WIDTH = 1200
const WINDOW_HEIGHT = 700

func main() {
	initMp3()
	loadResources()

	a := app.New()
	w := a.NewWindow("AppContainer")
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	w.SetContent(GetGUI())

	go setProg()
	go watch() //go to go
	w.ShowAndRun()
}
