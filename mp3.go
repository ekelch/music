package main

import (
	"bytes"
	"os"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

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
