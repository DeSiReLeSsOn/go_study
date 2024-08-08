package main

import (
	"fmt"
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
	receiveFile()
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

// receive file metadata
func receiveFile() {
	fileStat, err := os.Stat("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()
}
