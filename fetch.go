package main

import (
	"fmt"
	"os"
	"os/exec"
)

func downloadSC(query string) {
	dlCmd := exec.Command("./ext/sc", "-b", "-p", TEMP_DIR, "-f", query)
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
	tempFiles, err := os.ReadDir(TEMP_DIR)
	if err != nil {
		panic("error reading temp dir")
	}
	for _, t := range tempFiles {
		fmt.Println("mv file: " + t.Name())
		fPath := TEMP_DIR + "/" + t.Name()
		rPath := RES_DIR + "/" + t.Name()
		mvCmd := exec.Command("mv", fPath, rPath)
		err := mvCmd.Run()
		if err != nil {
			panic("Failed moving " + t.Name())
		}
	}
}
