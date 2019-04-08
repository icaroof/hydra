package main

import (
	"fmt"
	"os"
)

func main() {
	GenerateFileStatusReport("testfile.txt")
}

//GenerateFileStatusReport prints the file status
func GenerateFileStatusReport(fname string) {
	filestats, err := os.Stat(fname)
	PrintFatalError(err)

	fmt.Println("What's the file name?", filestats.Name())
	fmt.Println("Am I a directory?", filestats.IsDir())
	fmt.Println("What are the permissions?", filestats.Mode())
	fmt.Println("What's the file size?", filestats.Size())
	fmt.Println("When was the last time the file was modified?", filestats.ModTime())
}
