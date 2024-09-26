package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ebitengine/oto/v3"
)

var otoGlobalContext oto.Context
var player oto.Player
var currentSong Song
var songList = []Song{
	{name: "Innocent of All Things", path: "innocent.mp3"},
	{name: "Reality Surf", path: "reality.mp3"},
	{name: "Noblest Strive", path: "noblest.mp3"},
	{name: "I Think...", path: "ithink.mp3"},
	{name: "Blush", path: "blush.mp3"}}

const WINDOW_WIDTH = 1200
const WINDOW_HEIGHT = 700

func main() {
	initMp3()
	a := app.New()
	w := a.NewWindow("AppContainer")
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	w.SetContent(GetGUI())

	go setProg()
	w.ShowAndRun()
}
