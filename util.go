package main

import (
	"strings"
)

func isFileType(fileName string, fileType string) bool { // very poor impl lol
	return strings.Contains(fileName, fileType)
}
