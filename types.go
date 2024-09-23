package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Song struct {
	name   string
	path   string
	durSec int
}

type tappableSlider struct {
	widget.Slider
}

func (t *tappableSlider) Tapped(pe *fyne.PointEvent) {
	seekTime(pe.AbsolutePosition.X / 1200)
}

func (t *tappableSlider) SecondaryTapped(_ *fyne.PointEvent) {

}
