package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

const sampleRate int = 44100

func initMp3() { // only runs once on app start
	otoConfig := oto.NewContextOptions{SampleRate: sampleRate, ChannelCount: 2, Format: oto.FormatSignedInt16LE}

	otoContext, readyChan, err := oto.NewContext(&otoConfig)
	if err != nil {
		panic("Oto new context failed : " + err.Error())
	}

	otoGlobalContext = *otoContext

	<-readyChan
}

func decodeMp3(song *Song) *mp3.Decoder {
	fileBytes, err := os.ReadFile("resources/" + song.path)
	if err != nil {
		panic("Failed to read mp3 file: " + err.Error())
	}

	reader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(reader)
	if err != nil {
		panic("Failed to decode mp3 bytes reader: " + err.Error())
	}
	song.durSec = int(decodedMp3.Length() * 8 / (int64(decodedMp3.SampleRate()) * 32))
	return decodedMp3
}

func readSong(song Song) {
	if (player != oto.Player{}) {
		player.Pause()
	}
	decodedMp3 := decodeMp3(&song)
	currentSong = song

	songTime = 0

	player = *otoGlobalContext.NewPlayer(decodedMp3)
	player.Play()
}

func ppSong() {
	if (player == oto.Player{}) {
		return
	}
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

func seekTime(t int64) {
	fmt.Println(t)
	if (player != oto.Player{}) {
		re, err := player.Seek(t, io.SeekStart)
		fmt.Println("response from seek" + string(re))
		fmt.Println("error from seek: " + err.Error())
	}
}
