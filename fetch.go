package main

import (
	"fmt"
	"os"
	"os/exec"
)

func downloadSC(query string) {
	dlCmd := exec.Command("./ext/sc", "-b", "-p", "./temp", "-f", query)
	err := dlCmd.Run()
	if err != nil {
		fmt.Printf("Failed to download audio file.\n")
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("successfully downloaded file")
		mvTemp()
	}
}

func mvTemp() {
	tempFiles, err := os.ReadDir("./temp")
	if err != nil {
		panic("error reading temp dir")
	}
	for _, t := range tempFiles {
		fmt.Println("Attempting to move file: " + t.Name())
		mvCmd := exec.Command("mv", "./temp/"+t.Name(), "resources/"+t.Name())
		err := mvCmd.Run()
		if err != nil {
			panic("Failed moving " + t.Name())
		}
	}
}

// Deprecated: not tested/working much bc of HTTP 429 on get request to server, switching to sc development
// requires ffmpeg in local path
func downloadYoutubeVideo(query string, fileName string) {
	fmtQ := "\"" + query + "\""
	fmtQ = "https://www.youtube.com/watch?v=jlzgS2jKaIw"

	dlCmd := exec.Command("you-get", "-o", "./resources", "-O", "temp", fmtQ)
	err := dlCmd.Run()
	if err != nil {
		fmt.Printf("Failed to download video.\n")
		return
	} else {
		fmt.Println("successfully downloaded yt video")
	}

	fmtFile := "\"" + fileName + ".mp3\""
	convertCmd := exec.Command("ffmpeg", "-i", "resources/temp.mp4", "-vn", fmtFile)
	err = convertCmd.Run()
	if err != nil {
		panic("Failed to run ffmpeg convert cmd: " + err.Error())
	} else {
		fmt.Println("successfully converted mp4 to mp3")
	}

	rmVideoCmd := exec.Command("rm", "resources/temp.mp4")
	err = rmVideoCmd.Run()
	if err != nil {
		panic("Failed to rm temp.mp4: " + err.Error())
	} else {
		fmt.Println("successfully deleted temp.mp4")
	}

	songList = append(songList, fileName)
}

var proxyPort int16 = 8000
