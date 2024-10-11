package main

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("Failed to open file watcher!\n" + err.Error())
	}
	defer watcher.Close()

	// run watcher async
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println(event.String())
				fileName, prefixed := strings.CutPrefix(event.Name, RES_DIR+"/")
				if !prefixed {
					panic("Failing to move files, RESDIR not found")
				}
				if event.Has(fsnotify.Create) {
					addResource(fileName)
				} else if event.Has(fsnotify.Remove) {
					rmResource(fileName)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println(err)
			}
		}
	}()

	//add path to watch
	path := RES_DIR
	err = watcher.Add(path)
	if err != nil {
		fmt.Printf("Failed to add watcher to path: %s", path)
	}

	//blocking main goroutine, not sure what this is doing ! just never returns from here ?
	<-make(chan struct{})
}
