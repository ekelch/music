package main

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var supportedExt []string = []string{".mp3", ".ogg"}

func isMusic(fileWithExt string) bool {
	for _, ext := range supportedExt {
		if strings.Contains(fileWithExt, ext) {
			return true
		}
	}
	return false
}

func rmSongAtIndex(index int) {
	songList = append(songList[:index], songList[index+1:]...)
	songListBinding.Set(songList)
}

func rmSongByName(fileName string) {
	for i, v := range songList {
		if v == fileName {
			rmSongAtIndex(i)
		}
	}
}

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

func rmResource(fileName string) {
	path := RES_DIR + "/" + fileName
	for i, v := range songList {
		if v == fileName {
			rmCmd := exec.Command("rm", path)
			err := rmCmd.Run()
			if err != nil {
				panic("Fatal error while removing file " + path + "\n" + err.Error())
			}
			rmSongAtIndex(i)
		}
	}
}

func mvResource(oldName string, newName string) {
	if !isMusic(newName) {
		newName = newName + ".mp3" //todo obvious temp...
	}
	oPath := RES_DIR + "/" + oldName
	nPath := RES_DIR + "/" + newName
	mvCmd := exec.Command("mv", oPath, nPath)
	err := mvCmd.Run()
	if err != nil {
		panic("Fatal error moving files: " + oldName + " to " + newName + "\n" + err.Error())
	}
	rmSongByName(oldName)
}

func isFileType(fileName string, fileType string) bool { // very poor impl lol
	return strings.Contains(fileName, fileType)
}
