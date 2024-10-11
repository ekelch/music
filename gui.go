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
	controlArea := container.NewGridWithColumns(5, buildControlBtns(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), buildVolumeSlider())
	controlArea.Resize(fyne.Size{Width: 300, Height: 150})

	controlGroup := container.NewVBox(controlArea, buildSongProgress(), buildProgLabel())
	content := container.NewBorder(buildSearchForm(), controlGroup, nil, nil, buildSongList())

	return content
}

func buildSongList() *widget.List {
	songListBinding = binding.BindStringList(&songList)
	return widget.NewListWithData(
		songListBinding,
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil,
				nil,
				widget.NewButtonWithIcon("", theme.Icon(theme.IconNameMediaPlay), func() {}),
				container.NewHBox(widget.NewButton("rename", func() {}), widget.NewButton("delete", func() {})),
				widget.NewLabel("template"))
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).Bind(i.(binding.String))

			item, err := i.(binding.String).Get()
			if err != nil {
				panic(err)
			}

			gf_playBtn(item, o.(*fyne.Container).Objects[1].(*widget.Button))
			gf_renameBtn(item, o.(*fyne.Container).Objects[2].(*fyne.Container), 0)
			gf_deleteBtn(item, o.(*fyne.Container).Objects[2].(*fyne.Container), 1)
		})
}

func gf_playBtn(item string, btn *widget.Button) {
	btn.OnTapped = func() {
		readSong(item)
	}
}

func replaceRightDiv_mv(item string, hbox *fyne.Container) {
	oldContents := hbox.Objects
	hbox.RemoveAll()
	hbox.Add(getMoveForm(item, hbox, oldContents))
}

func setDivContent(hbox *fyne.Container, newContents []fyne.CanvasObject) {
	hbox.RemoveAll()
	for _, o := range newContents {
		hbox.Add(o)
	}
}

func getMoveForm(oldFileName string, hbox *fyne.Container, oldContents []fyne.CanvasObject) *fyne.Container {
	mvInput := widget.NewEntry()
	mvInput.PlaceHolder = "New File Name..."
	mvBtn := widget.NewButton("Confirm", func() {
		if len(mvInput.Text) > 0 {
			fmt.Printf("Renaming %s to %s\n", oldFileName, mvInput.Text)
			mvResource(oldFileName, mvInput.Text)
			setDivContent(hbox, oldContents)
		}
	})
	return container.NewHBox(mvInput, mvBtn)
}

func gf_renameBtn(item string, contain *fyne.Container, childIndex int) {
	btn := contain.Objects[childIndex].(*widget.Button)
	btn.OnTapped = func() {
		replaceRightDiv_mv(item, contain)
	}
}

func gf_deleteBtn(item string, contain *fyne.Container, childIndex int) {
	btn := contain.Objects[childIndex].(*widget.Button)
	btn.OnTapped = func() {
		fmt.Println("Deleting " + item)
		rmResource(item)
	}
}

func buildSearchForm() *fyne.Container {
	searchInput := widget.NewEntry()
	searchInput.SetPlaceHolder("Search on Soundcloud...")
	searchBtn := widget.NewButton("Search", nil)
	searchBtn.OnTapped = func() {
		fmt.Println(searchInput.Text)
		if len(searchInput.Text) > 0 {
			searchInput.Disable()
			searchBtn.Disable()
			downloadSC(searchInput.Text)
			searchInput.Text = ""
			searchInput.Enable()
			searchBtn.Enable()
		}
	}

	return container.NewBorder(nil, nil, nil, searchBtn, searchInput)
}

func buildControlBtns() *fyne.Container {
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
var songListBinding binding.StringList
var songElapsed float64 = 0.0

const PROG_MAX float64 = 10000
