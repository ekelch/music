package main

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var supportedExt []string = []string{".mp3", ".ogg"}

func loadResources() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot find user home dir: " + err.Error())
	}
	switch runtime.GOOS {
	case "windows":
		RES_DIR = "%APPDATA%/mp/resources"
		TEMP_DIR = "%APPDATA%/mp/temp"
	case "darwin":
		RES_DIR = homeDir + "/Library/Application Support/mp/resources"
		TEMP_DIR = homeDir + "/Library/Application Support/mp/temp"
	case "linux":
		RES_DIR = "~/mp/resources"
		TEMP_DIR = "~/mp/temp"
	}

	err = os.MkdirAll(RES_DIR, os.ModePerm)
	if err != nil {
		panic("Error creating resource dir")
	}
	err = os.MkdirAll(TEMP_DIR, os.ModePerm)
	if err != nil {
		panic("Error creating temp dir")
	}

	resources, err := os.ReadDir(RES_DIR)
	if err != nil {
		panic("Error reading resource dir: " + err.Error())
	}
	for _, file := range resources {
		if !file.IsDir() && isMusic(file.Name()) {
			songList = append(songList, file.Name())
		}
	}
}

func addResource(fileName string) {
	songList = append(songList, fileName)
	songListBinding.Set(songList)
}

func rmResource(fileName string) error {
	for i, _ := range songList {
		if currentSong.path == fileName {
			rmCmd := exec.Command("rm", fileName)
			err := rmCmd.Run()
			if err != nil {
				return err
			}
			songList = append(songList[:i], songList[i+1:]...)
			songListBinding.Set(songList)
			return nil
		}
	}
	panic("File not found while rm in rmResource")
}

func mvResource(oldName string, newName string) error {
	oPath := RES_DIR + "/" + oldName
	nPath := RES_DIR + "/" + newName
	mvCmd := exec.Command("mv", oPath, nPath)
	return mvCmd.Run()
}

func isFileType(fileName string, fileType string) bool { // very poor impl lol
	return strings.Contains(fileName, fileType)
}

func isMusic(fileWithExt string) bool {
	for _, ext := range supportedExt {
		if strings.Contains(fileWithExt, ext) {
			return true
		}
	}
	return false
}
