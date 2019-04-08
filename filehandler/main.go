package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func testingFileHangling() {
	//opens a file for read only
	f1, err := os.Open("test1.txt")
	PrintFatalError(err)
	defer f1.Close()

	//creates a new file
	f2, err := os.Create("test2.txt")
	PrintFatalError(err)
	defer f2.Close()

	//opens a file for read and write
	f3, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, 0666)
	PrintFatalError(err)
	defer f3.Close()

	err = os.Rename("test1.txt", "renamedTest1.txt")
	PrintFatalError(err)

	err = os.Rename("./test1.txt", "./testfolder/test1.txt")
	PrintFatalError(err)

	CopyFile("test3.txt", "./testfolder/test3.txt")

	err = os.Remove("test2.txt")
	PrintFatalError(err)

	bytes, err := ioutil.ReadFile("test3.txt")
	fmt.Println(string(bytes))

	//to read larger files, loading them in memory
	scanner := bufio.NewScanner(f3)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println("Found line:", count, scanner.Text())
	}

	writebuffer := bufio.NewWriter(f3)
	for i := 1; i <= 5; i++ {
		writebuffer.WriteString(fmt.Sprintln("Added line", i))
	}
	writebuffer.Flush()

}

//PrintFatalError Logs error if exists
func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}

//CopyFile Copies file fname1 to fname2
func CopyFile(fname1, fname2 string) {
	fOld, err := os.Open(fname1)
	PrintFatalError(err)
	defer fOld.Close()

	fNew, err := os.Create(fname2)
	PrintFatalError(err)
	defer fNew.Close()

	_, err = io.Copy(fNew, fOld)
	PrintFatalError(err)

	err = fNew.Sync()
	PrintFatalError(err)
}
