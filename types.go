package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Song struct {
	name   string
	path   string
	audio  []byte
	durSec int
}

type progressSlider struct {
	widget.Slider
}

func (t *progressSlider) Tapped(pe *fyne.PointEvent) {
	seekTime(pe.Position.X / t.Size().Width)
}

func (t *progressSlider) SecondaryTapped(_ *fyne.PointEvent) {

}

type volumeSlider struct {
	widget.Slider
}

func (t *volumeSlider) Tapped(pe *fyne.PointEvent) {
	setVolume(float64(pe.Position.X / t.Size().Width))
}

func (t *volumeSlider) SecondaryTapped(_ *fyne.PointEvent) {

}
