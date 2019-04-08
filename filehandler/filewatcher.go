package main

import (
	"fmt"
	"os"
	"time"
)

//WatchFile monitors file change state
func WatchFile(fname string) {
	filestat1, err := os.Stat(fname)
	PrintFatalError(err)

	for {
		time.Sleep(1 * time.Second)

		filestat2, err := os.Stat(fname)
		PrintFatalError(err)

		if filestat1.ModTime() != filestat2.ModTime() {
			fmt.Println("File was modified at", filestat2.ModTime())
			filestat1, err = os.Stat(fname)
			PrintFatalError(err)
		}
	}
}
