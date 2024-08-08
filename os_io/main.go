package main

import (
	"io"
	"log"
	"os"
)

func main() {
	createFile()
	createDir()
	renameFile()
	moveFile()
	copyFile()
}

// create empty file called empty
func createFile() {
	emptyFile, err := os.Create("testing.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	defer emptyFile.Close()
}

// create folder called test
func createDir() {
	_, err := os.Stat("test")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

// rename file from testing  name to new name empty
func renameFile() {
	oldName := "testing.txt"
	newName := "empty.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

// move file
func moveFile() {
	oldLocation := "/home/desireless/Desktop/go/go_study/os_io/empty1.txt"
	newLocation := "/home/desireless/Desktop/go/go_study/testing.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

// copy file
func copyFile() {
	sourceFile, err := os.Open("/home/desireless/Desktop/go/go_study/os_io/empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	//new file
	newFile, err := os.Create("/home/desireless/Desktop/go/go_study/os_io/testing1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Copied %d bytes.", bytesCopied)
}
