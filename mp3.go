package main

import (
	"bytes"
	"io"
	"os"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

// samples per second
const SAMPLE_RATE int = 44100

// bits per sample
const BIT_DEPTH int = 16

// number of channels, 2 for stereo
const CHANNEL_COUNT int = 2

// Bits/Second of sample
const BITRATE int64 = int64(SAMPLE_RATE * BIT_DEPTH * CHANNEL_COUNT)

func initMp3() { // only runs once on app start
	otoConfig := oto.NewContextOptions{SampleRate: SAMPLE_RATE, ChannelCount: CHANNEL_COUNT, Format: oto.FormatSignedInt16LE}

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
	song.durSec = int(8 * decodedMp3.Length() / BITRATE)
	return decodedMp3
}

func setVolume(v float64) {
	volumeBinding.Set(v * 100)
	player.SetVolume(v)
}

func readSong(song Song) {
	if (player != oto.Player{}) {
		player.Pause()
	}
	decodedMp3 := decodeMp3(&song)
	currentSong = song

	songElapsed = 0
	temp = decodedMp3.Length()

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

func seekTime(percent float32) {
	bytesOffset := BITRATE / 8 * int64(percent*float32(currentSong.durSec))
	if (player != oto.Player{}) {
		newpos, err := player.Seek(bytesOffset, io.SeekStart)
		songElapsed = float64(8 * newpos / BITRATE)
		if err != nil {
			panic("Failed to seek start of song: " + err.Error())
		}
	}
}

var temp int64
